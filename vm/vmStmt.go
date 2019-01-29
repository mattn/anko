package vm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// Run executes statement in the specified environment.
func Run(stmt ast.Stmt, env *Env) (interface{}, error) {
	return RunContext(context.Background(), stmt, env)
}

// RunContext executes statement in the specified environment with context.
func RunContext(ctx context.Context, stmt ast.Stmt, env *Env) (interface{}, error) {
	runInfo := runInfoStruct{ctx: ctx, env: env, stmt: stmt, rv: nilValue}
	runInfo.runSingleStmt()
	if runInfo.err == ErrReturn {
		runInfo.err = nil
	}
	if !runInfo.rv.IsValid() || !runInfo.rv.CanInterface() {
		return nil, runInfo.err
	}
	return runInfo.rv.Interface(), runInfo.err
}

// runSingleStmt executes statement in the specified environment with context.
func (runInfo *runInfoStruct) runSingleStmt() {
	select {
	case <-runInfo.ctx.Done():
		runInfo.err = ErrInterrupt
		return
	default:
	}

	switch stmt := runInfo.stmt.(type) {

	// nil
	case nil:

	// StmtsStmt
	case *ast.StmtsStmt:
		for _, stmt := range stmt.Stmts {
			switch stmt.(type) {
			case *ast.BreakStmt:
				runInfo.err = ErrBreak
				return
			case *ast.ContinueStmt:
				runInfo.err = ErrContinue
				return
			case *ast.ReturnStmt:
				runInfo.stmt = stmt
				runInfo.runSingleStmt()
				if runInfo.err != nil {
					return
				}
				runInfo.err = ErrReturn
				return
			default:
				runInfo.stmt = stmt
				runInfo.runSingleStmt()
				if runInfo.err != nil {
					return
				}
			}
		}

	// ExprStmt
	case *ast.ExprStmt:
		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr, runInfo.env)

	// VarStmt
	case *ast.VarStmt:
		// get right side expression values
		rvs := make([]reflect.Value, len(stmt.Exprs))
		for i, expr := range stmt.Exprs {
			rvs[i], runInfo.err = invokeExpr(runInfo.ctx, expr, runInfo.env)
			if runInfo.err != nil {
				return
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
					runInfo.err = runInfo.env.defineValue(stmt.Names[i], value.Index(i))
					if runInfo.err != nil {
						return
					}
				}
				// return last value of slice/array
				runInfo.rv = value.Index(value.Len() - 1)
				return
			}
		}

		// define all names with right side values
		for i := 0; i < len(rvs) && i < len(stmt.Names); i++ {
			runInfo.err = runInfo.env.defineValue(stmt.Names[i], rvs[i])
			if runInfo.err != nil {
				return
			}
		}

		// return last right side value
		runInfo.rv = rvs[len(rvs)-1]

	// LetsStmt
	case *ast.LetsStmt:
		// get right side expression values
		rvs := make([]reflect.Value, len(stmt.RHSS))
		for i, rhs := range stmt.RHSS {
			rvs[i], runInfo.err = invokeExpr(runInfo.ctx, rhs, runInfo.env)
			if runInfo.err != nil {
				return
			}
		}

		if len(rvs) == 1 && len(stmt.LHSS) > 1 {
			// only one right side value but many left side expressions
			value := rvs[0]
			if value.Kind() == reflect.Interface && !value.IsNil() {
				value = value.Elem()
			}
			if (value.Kind() == reflect.Slice || value.Kind() == reflect.Array) && value.Len() > 0 {
				// value is slice/array, add each value to left side expression
				for i := 0; i < value.Len() && i < len(stmt.LHSS); i++ {
					_, runInfo.err = invokeLetExpr(runInfo.ctx, stmt.LHSS[i], value.Index(i), runInfo.env)
					if runInfo.err != nil {
						return
					}
				}
				// return last value of slice/array
				runInfo.rv = value.Index(value.Len() - 1)
				return
			}
		}

		// invoke all left side expressions with right side values
		for i := 0; i < len(rvs) && i < len(stmt.LHSS); i++ {
			value := rvs[i]
			if value.Kind() == reflect.Interface && !value.IsNil() {
				value = value.Elem()
			}
			_, runInfo.err = invokeLetExpr(runInfo.ctx, stmt.LHSS[i], value, runInfo.env)
			if runInfo.err != nil {
				return
			}
		}

		// return last right side value
		runInfo.rv = rvs[len(rvs)-1]

	// LetMapItemStmt
	case *ast.LetMapItemStmt:
		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.RHS, runInfo.env)
		if runInfo.err != nil {
			return
		}
		var rvs []reflect.Value
		if isNil(runInfo.rv) {
			rvs = []reflect.Value{nilValue, falseValue}
		} else {
			rvs = []reflect.Value{runInfo.rv, trueValue}
		}
		for i, lhs := range stmt.LHSS {
			v := rvs[i]
			if v.Kind() == reflect.Interface && !v.IsNil() {
				v = v.Elem()
			}
			_, runInfo.err = invokeLetExpr(runInfo.ctx, lhs, v, runInfo.env)
			if runInfo.err != nil {
				return
			}
		}
		runInfo.rv = rvs[0]

	// IfStmt
	case *ast.IfStmt:
		// if
		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.If, runInfo.env)
		if runInfo.err != nil {
			return
		}

		env := runInfo.env

		if toBool(runInfo.rv) {
			// then
			runInfo.stmt = stmt.Then
			runInfo.env = env.NewEnv()
			runInfo.runSingleStmt()
			runInfo.env = env
			return
		}

		for _, statement := range stmt.ElseIf {
			elseIf := statement.(*ast.IfStmt)

			// else if - if
			runInfo.env = env.NewEnv()
			runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, elseIf.If, runInfo.env)
			if runInfo.err != nil {
				runInfo.env = env
				return
			}

			if !toBool(runInfo.rv) {
				continue
			}

			// else if - then
			runInfo.stmt = elseIf.Then
			runInfo.env = env.NewEnv()
			runInfo.runSingleStmt()
			runInfo.env = env
			return
		}

		if stmt.Else != nil {
			// else
			runInfo.stmt = stmt.Else
			runInfo.env = env.NewEnv()
			runInfo.runSingleStmt()
		}

		runInfo.env = env

	// TryStmt
	case *ast.TryStmt:
		// only the try statment will ignore any error except ErrInterrupt
		// all other parts will return the error

		env := runInfo.env
		runInfo.env = env.NewEnv()

		runInfo.stmt = stmt.Try
		runInfo.runSingleStmt()

		if runInfo.err != nil {
			if runInfo.err == ErrInterrupt {
				runInfo.env = env
				return
			}

			// Catch
			runInfo.stmt = stmt.Catch
			runInfo.env = env.NewEnv()
			if stmt.Var != "" {
				runInfo.err = runInfo.env.defineValue(stmt.Var, reflect.ValueOf(runInfo.err))
				if runInfo.err != nil {
					runInfo.env = env
					return
				}
			}
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				runInfo.env = env
				return
			}
		}

		if stmt.Finally != nil {
			// Finally
			runInfo.stmt = stmt.Finally
			runInfo.env = env.NewEnv()
			runInfo.runSingleStmt()
		}

		runInfo.env = env

	// LoopStmt
	case *ast.LoopStmt:
		env := runInfo.env
		runInfo.env = env.NewEnv()

		for {
			select {
			case <-runInfo.ctx.Done():
				runInfo.err = ErrInterrupt
				runInfo.env = env
				return
			default:
			}

			if stmt.Expr != nil {
				runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr, runInfo.env)
				if runInfo.err != nil {
					break
				}
				if !toBool(runInfo.rv) {
					break
				}
			}

			runInfo.stmt = stmt.Stmt
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				if runInfo.err == ErrBreak || runInfo.err == ErrReturn {
					runInfo.err = nil
				}
				break
			}
		}

		runInfo.env = env

	// ForStmt
	case *ast.ForStmt:
		var value reflect.Value
		value, runInfo.err = invokeExpr(runInfo.ctx, stmt.Value, runInfo.env)
		if runInfo.err != nil {
			return
		}
		if value.Kind() == reflect.Interface && !value.IsNil() {
			value = value.Elem()
		}

		env := runInfo.env
		runInfo.env = env.NewEnv()

		switch value.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < value.Len(); i++ {
				select {
				case <-runInfo.ctx.Done():
					runInfo.err = ErrInterrupt
					runInfo.env = env
					return
				default:
				}

				iv := value.Index(i)
				if iv.Kind() == reflect.Ptr || (iv.Kind() == reflect.Interface && !iv.IsNil()) {
					iv = iv.Elem()
				}
				runInfo.err = runInfo.env.defineValue(stmt.Vars[0], iv)
				if runInfo.err != nil {
					break
				}

				runInfo.stmt = stmt.Stmt
				runInfo.runSingleStmt()
				if runInfo.err != nil {
					if runInfo.err == ErrBreak || runInfo.err == ErrReturn {
						runInfo.err = nil
					}
					break
				}
			}
			runInfo.env = env

		case reflect.Map:
			keys := value.MapKeys()
			for i := 0; i < len(keys); i++ {
				select {
				case <-runInfo.ctx.Done():
					runInfo.err = ErrInterrupt
					runInfo.env = env
					return
				default:
				}

				runInfo.err = runInfo.env.defineValue(stmt.Vars[0], keys[i])
				if runInfo.err != nil {
					break
				}

				if len(stmt.Vars) > 1 {
					runInfo.err = runInfo.env.defineValue(stmt.Vars[1], value.MapIndex(keys[i]))
					if runInfo.err != nil {
						break
					}
				}

				runInfo.stmt = stmt.Stmt
				runInfo.runSingleStmt()
				if runInfo.err != nil {
					if runInfo.err == ErrBreak || runInfo.err == ErrReturn {
						runInfo.err = nil
					}
					break
				}
			}
			runInfo.env = env

		case reflect.Chan:
			var chosen int
			var ok bool
			for {
				cases := []reflect.SelectCase{{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(runInfo.ctx.Done()),
					Send: zeroValue,
				}, {
					Dir:  reflect.SelectRecv,
					Chan: value,
					Send: zeroValue,
				}}
				chosen, runInfo.rv, ok = reflect.Select(cases)
				if chosen == 0 {
					runInfo.err = ErrInterrupt
					break
				}
				if !ok {
					break
				}

				if runInfo.rv.Kind() == reflect.Ptr || (runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil()) {
					runInfo.rv = runInfo.rv.Elem()
				}
				runInfo.err = runInfo.env.defineValue(stmt.Vars[0], runInfo.rv)
				if runInfo.err != nil {
					break
				}

				runInfo.stmt = stmt.Stmt
				runInfo.runSingleStmt()
				if runInfo.err != nil {
					if runInfo.err == ErrBreak || runInfo.err == ErrReturn {
						runInfo.err = nil
					}
					break
				}
			}
			runInfo.env = env

		default:
			runInfo.env = env
			runInfo.err = newStringError(stmt, "for cannot loop over type "+value.Kind().String())
		}

	// CForStmt
	case *ast.CForStmt:
		env := runInfo.env
		runInfo.env = env.NewEnv()

		if stmt.Stmt1 != nil {
			runInfo.stmt = stmt.Stmt1
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				runInfo.env = env
				return
			}
		}

		for {
			select {
			case <-runInfo.ctx.Done():
				runInfo.err = ErrInterrupt
				runInfo.env = env
				return
			default:
			}

			if stmt.Expr2 != nil {
				runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr2, runInfo.env)
				if runInfo.err != nil {
					break
				}
				if !toBool(runInfo.rv) {
					break
				}
			}

			runInfo.stmt = stmt.Stmt
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				if runInfo.err == ErrBreak || runInfo.err == ErrReturn {
					runInfo.err = nil
				}
				break
			}

			if stmt.Expr3 != nil {
				runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr3, runInfo.env)
				if runInfo.err != nil {
					break
				}
			}
		}
		runInfo.env = env

	// ReturnStmt
	case *ast.ReturnStmt:
		switch len(stmt.Exprs) {
		case 0:
			return
		case 1:
			runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Exprs[0], runInfo.env)
			return
		}
		rvs := make([]interface{}, len(stmt.Exprs))
		for i, expr := range stmt.Exprs {
			runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, expr, runInfo.env)
			if runInfo.err != nil {
				return
			}
			if !runInfo.rv.IsValid() || !runInfo.rv.CanInterface() {
				rvs[i] = nil
			} else {
				rvs[i] = runInfo.rv.Interface()
			}
		}
		runInfo.rv = reflect.ValueOf(rvs)

	// ThrowStmt
	case *ast.ThrowStmt:
		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr, runInfo.env)
		if runInfo.err != nil {
			return
		}
		if !runInfo.rv.IsValid() {
			runInfo.err = newStringError(stmt.Expr, "can not throw invalid type")
			return
		}
		runInfo.err = newStringError(stmt, fmt.Sprint(runInfo.rv.Interface()))

	// ModuleStmt
	case *ast.ModuleStmt:
		env := runInfo.env
		runInfo.env = env.NewEnv()
		runInfo.env.SetName(stmt.Name)

		runInfo.stmt = stmt.Stmt
		runInfo.runSingleStmt()
		if runInfo.err != nil {
			runInfo.env = env
			return
		}

		runInfo.err = env.defineGlobalValue(stmt.Name, reflect.ValueOf(runInfo.env))
		runInfo.env = env

	// SwitchStmt
	case *ast.SwitchStmt:
		env := runInfo.env
		runInfo.env = env.NewEnv()

		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr, runInfo.env)
		if runInfo.err != nil {
			runInfo.env = env
			return
		}

		body := stmt.Body.(*ast.SwitchBodyStmt)
		if body.Default != nil {
			runInfo.stmt = body.Default
		}

		// TOFIX: make this so it is not runing the Exprs over and over again. This will cause bug in some cases, like i++.
	Loop:
		for _, switchCaseStmt := range body.Cases {
			caseStmt := switchCaseStmt.(*ast.SwitchCaseStmt)
			for _, expr := range caseStmt.Exprs {
				runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, expr, runInfo.env)
				if runInfo.err != nil {
					runInfo.env = env
					return
				}
				if equal(runInfo.rv, runInfo.rv) {
					runInfo.stmt = caseStmt.Stmt
					break Loop
				}
			}
		}

		if runInfo.stmt != nil {
			runInfo.runSingleStmt()
		}

		runInfo.env = env

	// GoroutineStmt
	case *ast.GoroutineStmt:
		runInfo.rv, runInfo.err = invokeExpr(runInfo.ctx, stmt.Expr, runInfo.env)

	// default
	default:
		runInfo.err = newStringError(stmt, "unknown statement")
	}

}
