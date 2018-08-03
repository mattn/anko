package vm

import (
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// Run executes statements in the specified environment.
func Run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	rv, err := run(stmts, env)
	if err == ErrReturn {
		err = nil
	}
	if !rv.IsValid() || !rv.CanInterface() {
		return nil, err
	}
	return rv.Interface(), err
}

// run executes statements in the specified environment.
func run(stmts []ast.Stmt, env *Env) (reflect.Value, error) {
	rv := nilValue
	var err error
	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.BreakStmt:
			return nilValue, ErrBreak
		case *ast.ContinueStmt:
			return nilValue, ErrContinue
		case *ast.ReturnStmt:
			rv, err = runSingleStmt(stmt, env)
			if err != nil {
				return rv, err
			}
			return rv, ErrReturn
		default:
			rv, err = runSingleStmt(stmt, env)
			if err != nil {
				return rv, err
			}
		}
	}
	return rv, nil
}

// RunSingleStmt executes one statement in the specified environment.
func RunSingleStmt(stmt ast.Stmt, env *Env) (interface{}, error) {
	rv, err := runSingleStmt(stmt, env)
	if !rv.IsValid() || !rv.CanInterface() {
		return nil, err
	}
	return rv.Interface(), err
}

// runSingleStmt executes one statement in the specified environment.
func runSingleStmt(stmt ast.Stmt, env *Env) (reflect.Value, error) {
	if *(env.interrupt) {
		return nilValue, ErrInterrupt
	}
	switch stmt := stmt.(type) {

	// ExprStmt
	case *ast.ExprStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, newError(stmt.Expr, err)
		}
		return rv, nil

	// VarStmt
	case *ast.VarStmt:
		var err error

		// get right side expression values
		rvs := make([]reflect.Value, len(stmt.Exprs))
		for i, expr := range stmt.Exprs {
			rvs[i], err = invokeExpr(expr, env)
			if err != nil {
				return nilValue, newError(expr, err)
			}
		}

		if len(rvs) == 1 && len(stmt.Names) > 1 {
			// only one right side value but many left side names
			value := rvs[0]
			if value.Kind() == reflect.Interface && !value.IsNil() {
				value = value.Elem()
			}
			if (value.Kind() == reflect.Slice || value.Kind() == reflect.Array) && value.Len() > 0 {
				// value is slice/array, add each value to left side names
				for i := 0; i < value.Len() && i < len(stmt.Names); i++ {
					env.defineValue(stmt.Names[i], value.Index(i))
				}
				// return last value of slice/array
				return value.Index(value.Len() - 1), nil
			}
		}

		// define all names with right side values
		for i := 0; i < len(rvs) && i < len(stmt.Names); i++ {
			env.defineValue(stmt.Names[i], rvs[i])
		}

		// return last right side value
		return rvs[len(rvs)-1], nil

	// LetsStmt
	case *ast.LetsStmt:
		var err error

		// get right side expression values
		rvs := make([]reflect.Value, len(stmt.Rhss))
		for i, rhs := range stmt.Rhss {
			rvs[i], err = invokeExpr(rhs, env)
			if err != nil {
				return nilValue, newError(rhs, err)
			}
		}

		if len(rvs) == 1 && len(stmt.Lhss) > 1 {
			// only one right side value but many left side expressions
			value := rvs[0]
			if value.Kind() == reflect.Interface && !value.IsNil() {
				value = value.Elem()
			}
			if (value.Kind() == reflect.Slice || value.Kind() == reflect.Array) && value.Len() > 0 {
				// value is slice/array, add each value to left side expression
				for i := 0; i < value.Len() && i < len(stmt.Lhss); i++ {
					_, err = invokeLetExpr(stmt.Lhss[i], value.Index(i), env)
					if err != nil {
						return nilValue, newError(stmt.Lhss[i], err)
					}
				}
				// return last value of slice/array
				return value.Index(value.Len() - 1), nil
			}
		}

		// invoke all left side expressions with right side values
		for i := 0; i < len(rvs) && i < len(stmt.Lhss); i++ {
			value := rvs[i]
			if value.Kind() == reflect.Interface && !value.IsNil() {
				value = value.Elem()
			}
			_, err = invokeLetExpr(stmt.Lhss[i], value, env)
			if err != nil {
				return nilValue, newError(stmt.Lhss[i], err)
			}
		}

		// return last right side value
		return rvs[len(rvs)-1], nil

	// LetMapItemStmt
	case *ast.LetMapItemStmt:
		rv, err := invokeExpr(stmt.Rhs, env)
		if err != nil {
			return nilValue, newError(stmt, err)
		}
		var rvs []reflect.Value
		if isNil(rv) {
			rvs = []reflect.Value{nilValue, falseValue}
		} else {
			rvs = []reflect.Value{rv, trueValue}
		}
		for i, lhs := range stmt.Lhss {
			v := rvs[i]
			if v.Kind() == reflect.Interface && !v.IsNil() {
				v = v.Elem()
			}
			_, err = invokeLetExpr(lhs, v, env)
			if err != nil {
				return nilValue, newError(lhs, err)
			}
		}
		return rvs[0], nil

	// IfStmt
	case *ast.IfStmt:
		// if
		rv, err := invokeExpr(stmt.If, env)
		if err != nil {
			return rv, newError(stmt.If, err)
		}

		if toBool(rv) {
			// then
			newenv := env.NewEnv()
			rv, err = run(stmt.Then, newenv)
			if err != nil {
				return rv, newError(stmt, err)
			}
			return rv, nil
		}

		for _, statement := range stmt.ElseIf {
			elseIf := statement.(*ast.IfStmt)
			// else if - if
			newenv := env.NewEnv()
			rv, err = invokeExpr(elseIf.If, newenv)
			if err != nil {
				return rv, newError(elseIf.If, err)
			}
			if !toBool(rv) {
				continue
			}

			// else if - then
			newenv = env.NewEnv()
			rv, err = run(elseIf.Then, newenv)
			if err != nil {
				return rv, newError(elseIf, err)
			}
			return rv, nil
		}

		if len(stmt.Else) > 0 {
			// else
			newenv := env.NewEnv()
			rv, err = run(stmt.Else, newenv)
			if err != nil {
				return rv, newError(stmt, err)
			}
		}

		return rv, nil

	// TryStmt
	case *ast.TryStmt:
		newenv := env.NewEnv()
		_, err := run(stmt.Try, newenv)
		if err != nil {
			// Catch
			cenv := env.NewEnv()
			if stmt.Var != "" {
				cenv.defineValue(stmt.Var, reflect.ValueOf(err))
			}
			_, e1 := run(stmt.Catch, cenv)
			if e1 != nil {
				err = newError(stmt.Catch[0], e1)
			} else {
				err = nil
			}
		}
		if len(stmt.Finally) > 0 {
			// Finally
			_, e2 := run(stmt.Finally, newenv)
			if e2 != nil {
				err = newError(stmt.Finally[0], e2)
			}
		}
		return nilValue, newError(stmt, err)

	// LoopStmt
	case *ast.LoopStmt:
		newenv := env.NewEnv()
		for {
			if *(env.interrupt) {
				return nilValue, ErrInterrupt
			}
			if stmt.Expr != nil {
				ev, ee := invokeExpr(stmt.Expr, newenv)
				if ee != nil {
					return ev, ee
				}
				if !toBool(ev) {
					break
				}
			}

			rv, err := run(stmt.Stmts, newenv)
			if err != nil && err != ErrContinue {
				if err == ErrBreak {
					break
				}
				if err == ErrReturn {
					return rv, err
				}
				return nilValue, newError(stmt, err)
			}
		}
		return nilValue, nil

	// ForStmt
	case *ast.ForStmt:
		val, ee := invokeExpr(stmt.Value, env)
		if ee != nil {
			return val, ee
		}
		if val.Kind() == reflect.Interface && !val.IsNil() {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			newenv := env.NewEnv()

			for i := 0; i < val.Len(); i++ {
				if *(env.interrupt) {
					return nilValue, ErrInterrupt
				}
				iv := val.Index(i)
				if iv.Kind() == reflect.Ptr || (iv.Kind() == reflect.Interface && !iv.IsNil()) {
					iv = iv.Elem()
				}
				newenv.defineValue(stmt.Vars[0], iv)
				rv, err := run(stmt.Stmts, newenv)
				if err != nil && err != ErrContinue {
					if err == ErrBreak {
						break
					}
					if err == ErrReturn {
						return rv, err
					}
					return nilValue, newError(stmt, err)
				}
			}
			return nilValue, nil
		case reflect.Map:
			newenv := env.NewEnv()

			keys := val.MapKeys()
			for i := 0; i < len(keys); i++ {
				if *(env.interrupt) {
					return nilValue, ErrInterrupt
				}
				newenv.defineValue(stmt.Vars[0], keys[i])
				if len(stmt.Vars) > 1 {
					newenv.defineValue(stmt.Vars[1], val.MapIndex(keys[i]))
				}
				rv, err := run(stmt.Stmts, newenv)
				if err != nil && err != ErrContinue {
					if err == ErrBreak {
						break
					}
					if err == ErrReturn {
						return rv, err
					}
					return nilValue, newError(stmt, err)
				}
			}
			return nilValue, nil
		case reflect.Chan:
			newenv := env.NewEnv()

			for {
				if *(env.interrupt) {
					return nilValue, ErrInterrupt
				}
				iv, ok := val.Recv()
				if !ok {
					break
				}
				if iv.Kind() == reflect.Interface || iv.Kind() == reflect.Ptr {
					iv = iv.Elem()
				}
				newenv.defineValue(stmt.Vars[0], iv)
				rv, err := run(stmt.Stmts, newenv)
				if err != nil && err != ErrContinue {
					if err == ErrBreak {
						break
					}
					if err == ErrReturn {
						return rv, err
					}
					return nilValue, newError(stmt, err)
				}
			}
			return nilValue, nil
		default:
			return nilValue, newStringError(stmt, "for cannot loop over type "+val.Kind().String())
		}

	// CForStmt
	case *ast.CForStmt:
		newenv := env.NewEnv()
		_, err := runSingleStmt(stmt.Stmt1, newenv)
		if err != nil {
			return nilValue, err
		}
		for {
			if *(env.interrupt) {
				return nilValue, ErrInterrupt
			}
			fb, err := invokeExpr(stmt.Expr2, newenv)
			if err != nil {
				return nilValue, err
			}
			if !toBool(fb) {
				break
			}

			rv, err := run(stmt.Stmts, newenv)
			if err != nil && err != ErrContinue {
				if err == ErrBreak {
					break
				}
				if err == ErrReturn {
					return rv, err
				}
				return nilValue, newError(stmt, err)
			}
			_, err = invokeExpr(stmt.Expr3, newenv)
			if err != nil {
				return nilValue, err
			}
		}
		return nilValue, nil

	// ReturnStmt
	case *ast.ReturnStmt:
		var err error
		rv := nilValue
		switch len(stmt.Exprs) {
		case 0:
			return rv, nil
		case 1:
			rv, err = invokeExpr(stmt.Exprs[0], env)
			if err != nil {
				return rv, newError(stmt, err)
			}
			return rv, nil
		}
		rvs := make([]interface{}, len(stmt.Exprs))
		for i, expr := range stmt.Exprs {
			rv, err = invokeExpr(expr, env)
			if err != nil {
				return rv, newError(stmt, err)
			}
			if !rv.IsValid() || !rv.CanInterface() {
				rvs[i] = nil
			} else {
				rvs[i] = rv.Interface()
			}
		}
		return reflect.ValueOf(rvs), nil

	// ThrowStmt
	case *ast.ThrowStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, newError(stmt, err)
		}
		if !rv.IsValid() {
			return nilValue, newError(stmt, err)
		}
		return rv, newStringError(stmt, fmt.Sprint(rv.Interface()))

	// ModuleStmt
	case *ast.ModuleStmt:
		newenv := env.NewEnv()
		newenv.SetName(stmt.Name)
		rv, err := run(stmt.Stmts, newenv)
		if err != nil {
			return rv, newError(stmt, err)
		}
		env.defineGlobalValue(stmt.Name, reflect.ValueOf(newenv))
		return rv, nil

	// SwitchStmt
	case *ast.SwitchStmt:
		newenv := env.NewEnv()
		rv, err := invokeExpr(stmt.Expr, newenv)
		if err != nil {
			return rv, newError(stmt, err)
		}

		var caseValue reflect.Value
		body := stmt.Body.(*ast.SwitchBodyStmt)
		var statements []ast.Stmt
		if body.Default != nil {
			statements = body.Default
		}

	Loop:
		for _, switchCaseStmt := range body.Cases {
			caseStmt := switchCaseStmt.(*ast.SwitchCaseStmt)
			for _, expr := range caseStmt.Exprs {
				caseValue, err = invokeExpr(expr, newenv)
				if err != nil {
					return nilValue, newError(expr, err)
				}
				if equal(rv, caseValue) {
					statements = caseStmt.Stmts
					break Loop
				}
			}
		}

		if statements != nil {
			rv, err = run(statements, newenv)
			if err != nil {
				return rv, err
			}
		}

		return rv, nil

	// GoroutineStmt
	case *ast.GoroutineStmt:
		return invokeExpr(stmt.Expr, env)

	// default
	default:
		return nilValue, newStringError(stmt, "unknown statement")
	}
}
