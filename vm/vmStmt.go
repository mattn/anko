package vm

import (
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// Run executes statements in the specified environment.
func Run(stmts []ast.Stmt, env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.BreakStmt); ok {
			return NilValue, BreakError
		}
		if _, ok := stmt.(*ast.ContinueStmt); ok {
			return NilValue, ContinueError
		}
		rv, err = RunSingleStmt(stmt, env)
		if err != nil {
			return rv, err
		}
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return reflect.ValueOf(rv), ReturnError
		}
	}
	return rv, nil
}

// RunSingleStmt executes one statement in the specified environment.
func RunSingleStmt(stmt ast.Stmt, env *Env) (reflect.Value, error) {
	env.Lock()
	if *(env.interrupt) {
		*(env.interrupt) = false
		env.Unlock()

		return NilValue, InterruptError
	}
	env.Unlock()

	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, nil
	case *ast.VarStmt:
		rv := NilValue
		var err error
		rvs := []reflect.Value{}
		for _, expr := range stmt.Exprs {
			rv, err = invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(expr, err)
			}
			rvs = append(rvs, rv)
		}
		result := []interface{}{}
		for i, name := range stmt.Names {
			if i < len(rvs) {
				env.Define(name, rvs[i])
				result = append(result, rvs[i].Interface())
			}
		}
		return reflect.ValueOf(result), nil
	case *ast.LetsStmt:
		rv := NilValue
		var err error
		vs := []interface{}{}
		for _, rhs := range stmt.Rhss {
			rv, err = invokeExpr(rhs, env)
			if err != nil {
				return rv, NewError(rhs, err)
			}
			if rv == NilValue {
				vs = append(vs, nil)
			} else if rv.IsValid() && rv.CanInterface() {
				vs = append(vs, rv.Interface())
			} else {
				vs = append(vs, nil)
			}
		}
		rvs := reflect.ValueOf(vs)
		if len(stmt.Lhss) > 1 && rvs.Len() == 1 {
			item := rvs.Index(0)
			if item.Kind() == reflect.Interface {
				item = item.Elem()
			}
			if item.Kind() == reflect.Slice {
				rvs = item
			}
		}
		for i, lhs := range stmt.Lhss {
			if i >= rvs.Len() {
				break
			}
			v := rvs.Index(i)
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			_, err = invokeLetExpr(lhs, v, env)
			if err != nil {
				return rvs, NewError(lhs, err)
			}
		}
		if rvs.Len() == 1 {
			return rvs.Index(0), nil
		}
		return rvs, nil
	case *ast.IfStmt:
		// If
		rv, err := invokeExpr(stmt.If, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if toBool(rv) {
			// Then
			newenv := env.NewEnv()
			defer newenv.Destroy()
			rv, err = Run(stmt.Then, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		done := false
		if len(stmt.ElseIf) > 0 {
			for _, stmt := range stmt.ElseIf {
				stmt_if := stmt.(*ast.IfStmt)
				// ElseIf
				rv, err = invokeExpr(stmt_if.If, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				if !toBool(rv) {
					continue
				}
				// ElseIf Then
				done = true
				rv, err = Run(stmt_if.Then, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				break
			}
		}
		if !done && len(stmt.Else) > 0 {
			// Else
			newenv := env.NewEnv()
			defer newenv.Destroy()
			rv, err = Run(stmt.Else, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	case *ast.TryStmt:
		newenv := env.NewEnv()
		defer newenv.Destroy()
		_, err := Run(stmt.Try, newenv)
		if err != nil {
			// Catch
			cenv := env.NewEnv()
			defer cenv.Destroy()
			if stmt.Var != "" {
				cenv.Define(stmt.Var, reflect.ValueOf(err))
			}
			_, e1 := Run(stmt.Catch, cenv)
			if e1 != nil {
				err = NewError(stmt.Catch[0], e1)
			} else {
				err = nil
			}
		}
		if len(stmt.Finally) > 0 {
			// Finally
			fenv := env.NewEnv()
			defer fenv.Destroy()
			_, e2 := Run(stmt.Finally, newenv)
			if e2 != nil {
				err = NewError(stmt.Finally[0], e2)
			}
		}
		return NilValue, NewError(stmt, err)
	case *ast.LoopStmt:
		newenv := env.NewEnv()
		defer newenv.Destroy()
		for {
			if stmt.Expr != nil {
				ev, ee := invokeExpr(stmt.Expr, newenv)
				if ee != nil {
					return ev, ee
				}
				if !toBool(ev) {
					break
				}
			}

			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
	case *ast.ForStmt:
		val, ee := invokeExpr(stmt.Value, env)
		if ee != nil {
			return val, ee
		}
		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}
		if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
			newenv := env.NewEnv()
			defer newenv.Destroy()

			for i := 0; i < val.Len(); i++ {
				iv := val.Index(i)
				if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
					iv = iv.Elem()
				}
				newenv.Define(stmt.Var, iv)
				rv, err := Run(stmt.Stmts, newenv)
				if err != nil {
					if err == BreakError {
						err = nil
						break
					}
					if err == ContinueError {
						err = nil
						continue
					}
					if err == ReturnError {
						return rv, err
					}
					return rv, NewError(stmt, err)
				}
			}
			return NilValue, nil
		} else if val.Kind() == reflect.Chan {
			newenv := env.NewEnv()
			defer newenv.Destroy()

			for {
				iv, ok := val.Recv()
				if !ok {
					break
				}
				if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
					iv = iv.Elem()
				}
				newenv.Define(stmt.Var, iv)
				rv, err := Run(stmt.Stmts, newenv)
				if err != nil {
					if err == BreakError {
						err = nil
						break
					}
					if err == ContinueError {
						err = nil
						continue
					}
					if err == ReturnError {
						return rv, err
					}
					return rv, NewError(stmt, err)
				}
			}
			return NilValue, nil
		} else {
			return NilValue, NewStringError(stmt, "Invalid operation for non-array value")
		}
	case *ast.CForStmt:
		newenv := env.NewEnv()
		defer newenv.Destroy()
		_, err := invokeExpr(stmt.Expr1, newenv)
		if err != nil {
			return NilValue, err
		}
		for {
			fb, err := invokeExpr(stmt.Expr2, newenv)
			if err != nil {
				return NilValue, err
			}
			if !toBool(fb) {
				break
			}

			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					return rv, err
				}
				return rv, NewError(stmt, err)
			}
			_, err = invokeExpr(stmt.Expr3, newenv)
			if err != nil {
				return NilValue, err
			}
		}
		return NilValue, nil
	case *ast.ReturnStmt:
		rvs := []interface{}{}
		switch len(stmt.Exprs) {
		case 0:
			return NilValue, nil
		case 1:
			rv, err := invokeExpr(stmt.Exprs[0], env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		for _, expr := range stmt.Exprs {
			rv, err := invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if isNil(rv) {
				rvs = append(rvs, nil)
			} else if rv.IsValid() {
				rvs = append(rvs, rv.Interface())
			} else {
				rvs = append(rvs, nil)
			}
		}
		return reflect.ValueOf(rvs), nil
	case *ast.ThrowStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if !rv.IsValid() {
			return NilValue, NewError(stmt, err)
		}
		return rv, NewStringError(stmt, fmt.Sprint(rv.Interface()))
	case *ast.ModuleStmt:
		newenv := env.NewEnv()
		newenv.SetName(stmt.Name)
		rv, err := Run(stmt.Stmts, newenv)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		env.DefineGlobal(stmt.Name, reflect.ValueOf(newenv))
		return rv, nil
	case *ast.SwitchStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		done := false
		var default_stmt *ast.DefaultStmt
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				default_stmt = ssd
				continue
			}
			case_stmt := ss.(*ast.CaseStmt)
			cv, err := invokeExpr(case_stmt.Expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if !equal(rv, cv) {
				continue
			}
			rv, err = Run(case_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			done = true
			break
		}
		if !done && default_stmt != nil {
			rv, err = Run(default_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	default:
		return NilValue, NewStringError(stmt, "unknown statement")
	}
}
