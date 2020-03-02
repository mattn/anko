package vm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
)

// Execute parses script and executes in the specified environment.
func Execute(env *env.Env, options *Options, script string) (interface{}, error) {
	stmt, err := parser.ParseSrc(script)
	if err != nil {
		return nilValue, err
	}

	return RunContext(context.Background(), env, options, stmt)
}

// ExecuteContext parses script and executes in the specified environment with context.
func ExecuteContext(ctx context.Context, env *env.Env, options *Options, script string) (interface{}, error) {
	stmt, err := parser.ParseSrc(script)
	if err != nil {
		return nilValue, err
	}

	return RunContext(ctx, env, options, stmt)
}

// Run executes statement in the specified environment.
func Run(env *env.Env, options *Options, stmt ast.Stmt) (interface{}, error) {
	return RunContext(context.Background(), env, options, stmt)
}

// RunContext executes statement in the specified environment with context.
func RunContext(ctx context.Context, env *env.Env, options *Options, stmt ast.Stmt) (interface{}, error) {
	runInfo := runInfoStruct{ctx: ctx, env: env, options: options, stmt: stmt, rv: nilValue}
	if runInfo.options == nil {
		runInfo.options = &Options{}
	}
	runInfo.runSingleStmt()
	if runInfo.err == ErrReturn {
		runInfo.err = nil
	}
	return runInfo.rv.Interface(), runInfo.err
}

// runSingleStmt executes statement in the specified environment with context.
func (runInfo *runInfoStruct) runSingleStmt() {
	select {
	case <-runInfo.ctx.Done():
		runInfo.rv = nilValue
		runInfo.err = ErrInterrupt
		return
	default:
	}

	switch stmt := runInfo.stmt.(type) {
	case nil:
	case *ast.StmtsStmt:
		runInfo.runStmtsStmt(stmt)
	case *ast.ExprStmt:
		runInfo.runExprStmt(stmt)
	case *ast.VarStmt:
		runInfo.runVarStmt(stmt)
	case *ast.LetsStmt:
		runInfo.runLetsStmt(stmt)
	case *ast.LetMapItemStmt:
		runInfo.runLetMapItemStmt(stmt)
	case *ast.IfStmt:
		runInfo.runIfStmt(stmt)
	case *ast.TryStmt:
		runInfo.runTryStmt(stmt)
	case *ast.LoopStmt:
		runInfo.runLoopStmt(stmt)
	case *ast.ForStmt:
		runInfo.runForStmt(stmt)
	case *ast.CForStmt:
		runInfo.runCForStmt(stmt)
	case *ast.ReturnStmt:
		runInfo.runReturnStmt(stmt)
	case *ast.ThrowStmt:
		runInfo.runThrowStmt(stmt)
	case *ast.ModuleStmt:
		runInfo.runModuleStmt(stmt)
	case *ast.SwitchStmt:
		runInfo.runSwitchStmt(stmt)
	case *ast.GoroutineStmt:
		runInfo.runGoroutineStmt(stmt)
	case *ast.DeleteStmt:
		runInfo.runDeleteStmt(stmt)
	case *ast.CloseStmt:
		runInfo.runCloseStmt(stmt)
	case *ast.ChanStmt:
		runInfo.runChanStmt(stmt)
	default:
		runInfo.err = newStringError(stmt, "unknown statement")
		runInfo.rv = nilValue
	}
}

func (runInfo *runInfoStruct) runStmtsStmt(stmt *ast.StmtsStmt) {
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
}

func (runInfo *runInfoStruct) runExprStmt(stmt *ast.ExprStmt) {
	runInfo.expr = stmt.Expr
	runInfo.invokeExpr()
}

func (runInfo *runInfoStruct) runVarStmt(stmt *ast.VarStmt) {
	// get right side expression values
	rvs := make([]reflect.Value, len(stmt.Exprs))
	var i int
	for i, runInfo.expr = range stmt.Exprs {
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if env, ok := runInfo.rv.Interface().(*env.Env); ok {
			rvs[i] = reflect.ValueOf(env.DeepCopy())
		} else {
			rvs[i] = runInfo.rv
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
				runInfo.env.DefineValue(stmt.Names[i], value.Index(i))
			}
			// return last value of slice/array
			runInfo.rv = value.Index(value.Len() - 1)
			return
		}
	}

	// define all names with right side values
	for i = 0; i < len(rvs) && i < len(stmt.Names); i++ {
		runInfo.env.DefineValue(stmt.Names[i], rvs[i])
	}

	// return last right side value
	runInfo.rv = rvs[len(rvs)-1]
}

func (runInfo *runInfoStruct) runLetsStmt(stmt *ast.LetsStmt) {
	// get right side expression values
	rvs := make([]reflect.Value, len(stmt.RHSS))
	var i int
	for i, runInfo.expr = range stmt.RHSS {
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if env, ok := runInfo.rv.Interface().(*env.Env); ok {
			rvs[i] = reflect.ValueOf(env.DeepCopy())
		} else {
			rvs[i] = runInfo.rv
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
				runInfo.rv = value.Index(i)
				runInfo.expr = stmt.LHSS[i]
				runInfo.invokeLetExpr()
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
	for i = 0; i < len(rvs) && i < len(stmt.LHSS); i++ {
		value := rvs[i]
		if value.Kind() == reflect.Interface && !value.IsNil() {
			value = value.Elem()
		}
		runInfo.rv = value
		runInfo.expr = stmt.LHSS[i]
		runInfo.invokeLetExpr()
		if runInfo.err != nil {
			return
		}
	}

	// return last right side value
	runInfo.rv = rvs[len(rvs)-1]
}

func (runInfo *runInfoStruct) runLetMapItemStmt(stmt *ast.LetMapItemStmt) {
	runInfo.expr = stmt.RHS
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}
	var rvs []reflect.Value
	if isNil(runInfo.rv) {
		rvs = []reflect.Value{nilValue, falseValue}
	} else {
		rvs = []reflect.Value{runInfo.rv, trueValue}
	}
	var i int
	for i, runInfo.expr = range stmt.LHSS {
		runInfo.rv = rvs[i]
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}
		runInfo.invokeLetExpr()
		if runInfo.err != nil {
			return
		}
	}
	runInfo.rv = rvs[0]
}

func (runInfo *runInfoStruct) runIfStmt(stmt *ast.IfStmt) {
	// if
	runInfo.expr = stmt.If
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}

	env := runInfo.env

	if toBool(runInfo.rv) {
		// then
		runInfo.rv = nilValue
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
		runInfo.expr = elseIf.If
		runInfo.invokeExpr()
		if runInfo.err != nil {
			runInfo.env = env
			return
		}

		if !toBool(runInfo.rv) {
			continue
		}

		// else if - then
		runInfo.rv = nilValue
		runInfo.stmt = elseIf.Then
		runInfo.env = env.NewEnv()
		runInfo.runSingleStmt()
		runInfo.env = env
		return
	}

	if stmt.Else != nil {
		// else
		runInfo.rv = nilValue
		runInfo.stmt = stmt.Else
		runInfo.env = env.NewEnv()
		runInfo.runSingleStmt()
	}

	runInfo.env = env
}

func (runInfo *runInfoStruct) runTryStmt(stmt *ast.TryStmt) {
	// only the try statement will ignore any error except ErrInterrupt
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
		if stmt.Var != "" {
			runInfo.env.DefineValue(stmt.Var, reflect.ValueOf(runInfo.err))
		}
		runInfo.err = nil
		runInfo.runSingleStmt()
		if runInfo.err != nil {
			runInfo.env = env
			return
		}
	}

	if stmt.Finally != nil {
		// Finally
		runInfo.stmt = stmt.Finally
		runInfo.runSingleStmt()
	}

	runInfo.env = env
}

func (runInfo *runInfoStruct) runLoopStmt(stmt *ast.LoopStmt) {
	env := runInfo.env
	runInfo.env = env.NewEnv()

	for {
		select {
		case <-runInfo.ctx.Done():
			runInfo.err = ErrInterrupt
			runInfo.rv = nilValue
			runInfo.env = env
			return
		default:
		}

		if stmt.Expr != nil {
			runInfo.expr = stmt.Expr
			runInfo.invokeExpr()
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
			if runInfo.err == ErrContinue {
				runInfo.err = nil
				continue
			}
			if runInfo.err == ErrReturn {
				runInfo.env = env
				return
			}
			if runInfo.err == ErrBreak {
				runInfo.err = nil
			}
			break
		}
	}

	runInfo.rv = nilValue
	runInfo.env = env
}

func (runInfo *runInfoStruct) runForStmt(stmt *ast.ForStmt) {
	runInfo.expr = stmt.Value
	runInfo.invokeExpr()
	value := runInfo.rv
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
				runInfo.rv = nilValue
				runInfo.env = env
				return
			default:
			}

			iv := value.Index(i)
			if iv.Kind() == reflect.Interface && !iv.IsNil() {
				iv = iv.Elem()
			}
			if iv.Kind() == reflect.Ptr {
				iv = iv.Elem()
			}
			runInfo.env.DefineValue(stmt.Vars[0], iv)

			runInfo.stmt = stmt.Stmt
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				if runInfo.err == ErrContinue {
					runInfo.err = nil
					continue
				}
				if runInfo.err == ErrReturn {
					runInfo.env = env
					return
				}
				if runInfo.err == ErrBreak {
					runInfo.err = nil
				}
				break
			}
		}
		runInfo.rv = nilValue
		runInfo.env = env

	case reflect.Map:
		keys := value.MapKeys()
		for i := 0; i < len(keys); i++ {
			select {
			case <-runInfo.ctx.Done():
				runInfo.err = ErrInterrupt
				runInfo.rv = nilValue
				runInfo.env = env
				return
			default:
			}

			runInfo.env.DefineValue(stmt.Vars[0], keys[i])

			if len(stmt.Vars) > 1 {
				runInfo.env.DefineValue(stmt.Vars[1], value.MapIndex(keys[i]))
			}

			runInfo.stmt = stmt.Stmt
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				if runInfo.err == ErrContinue {
					runInfo.err = nil
					continue
				}
				if runInfo.err == ErrReturn {
					runInfo.env = env
					return
				}
				if runInfo.err == ErrBreak {
					runInfo.err = nil
				}
				break
			}
		}
		runInfo.rv = nilValue
		runInfo.env = env

	case reflect.Chan:
		var chosen int
		var ok bool
		for {
			cases := []reflect.SelectCase{{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(runInfo.ctx.Done()),
			}, {
				Dir:  reflect.SelectRecv,
				Chan: value,
			}}
			chosen, runInfo.rv, ok = reflect.Select(cases)
			if chosen == 0 {
				runInfo.err = ErrInterrupt
				runInfo.rv = nilValue
				break
			}
			if !ok {
				break
			}

			if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
				runInfo.rv = runInfo.rv.Elem()
			}
			if runInfo.rv.Kind() == reflect.Ptr {
				runInfo.rv = runInfo.rv.Elem()
			}

			runInfo.env.DefineValue(stmt.Vars[0], runInfo.rv)

			runInfo.stmt = stmt.Stmt
			runInfo.runSingleStmt()
			if runInfo.err != nil {
				if runInfo.err == ErrContinue {
					runInfo.err = nil
					continue
				}
				if runInfo.err == ErrReturn {
					runInfo.env = env
					return
				}
				if runInfo.err == ErrBreak {
					runInfo.err = nil
				}
				break
			}
		}
		runInfo.rv = nilValue
		runInfo.env = env

	default:
		runInfo.err = newStringError(stmt, "for cannot loop over type "+value.Kind().String())
		runInfo.rv = nilValue
		runInfo.env = env
	}
}

func (runInfo *runInfoStruct) runCForStmt(stmt *ast.CForStmt) {
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
			runInfo.rv = nilValue
			runInfo.env = env
			return
		default:
		}

		if stmt.Expr2 != nil {
			runInfo.expr = stmt.Expr2
			runInfo.invokeExpr()
			if runInfo.err != nil {
				break
			}
			if !toBool(runInfo.rv) {
				break
			}
		}

		runInfo.stmt = stmt.Stmt
		runInfo.runSingleStmt()
		if runInfo.err == ErrContinue {
			runInfo.err = nil
		}
		if runInfo.err != nil {
			if runInfo.err == ErrReturn {
				runInfo.env = env
				return
			}
			if runInfo.err == ErrBreak {
				runInfo.err = nil
			}
			break
		}

		if stmt.Expr3 != nil {
			runInfo.expr = stmt.Expr3
			runInfo.invokeExpr()
			if runInfo.err != nil {
				break
			}
		}
	}
	runInfo.rv = nilValue
	runInfo.env = env
}

func (runInfo *runInfoStruct) runReturnStmt(stmt *ast.ReturnStmt) {
	switch len(stmt.Exprs) {
	case 0:
		runInfo.rv = nilValue
		return
	case 1:
		runInfo.expr = stmt.Exprs[0]
		runInfo.invokeExpr()
		return
	}
	rvs := make([]interface{}, len(stmt.Exprs))
	var i int
	for i, runInfo.expr = range stmt.Exprs {
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		rvs[i] = runInfo.rv.Interface()
	}
	runInfo.rv = reflect.ValueOf(rvs)
}

func (runInfo *runInfoStruct) runThrowStmt(stmt *ast.ThrowStmt) {
	runInfo.expr = stmt.Expr
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}
	runInfo.err = newStringError(stmt, fmt.Sprint(runInfo.rv.Interface()))
}

func (runInfo *runInfoStruct) runModuleStmt(stmt *ast.ModuleStmt) {
	e := runInfo.env
	runInfo.env, runInfo.err = e.NewModule(stmt.Name)
	if runInfo.err != nil {
		return
	}
	runInfo.stmt = stmt.Stmt
	runInfo.runSingleStmt()
	runInfo.env = e
	if runInfo.err != nil {
		return
	}
	runInfo.rv = nilValue
}

func (runInfo *runInfoStruct) runSwitchStmt(stmt *ast.SwitchStmt) {
	env := runInfo.env
	runInfo.env = env.NewEnv()

	runInfo.expr = stmt.Expr
	runInfo.invokeExpr()
	if runInfo.err != nil {
		runInfo.env = env
		return
	}
	value := runInfo.rv

	for _, switchCaseStmt := range stmt.Cases {
		caseStmt := switchCaseStmt.(*ast.SwitchCaseStmt)
		for _, runInfo.expr = range caseStmt.Exprs {
			runInfo.invokeExpr()
			if runInfo.err != nil {
				runInfo.env = env
				return
			}
			if equal(runInfo.rv, value) {
				runInfo.stmt = caseStmt.Stmt
				runInfo.runSingleStmt()
				runInfo.env = env
				return
			}
		}
	}

	if stmt.Default == nil {
		runInfo.rv = nilValue
	} else {
		runInfo.stmt = stmt.Default
		runInfo.runSingleStmt()
	}

	runInfo.env = env
}

func (runInfo *runInfoStruct) runGoroutineStmt(stmt *ast.GoroutineStmt) {
	runInfo.expr = stmt.Expr
	runInfo.invokeExpr()
}

func (runInfo *runInfoStruct) runDeleteStmt(stmt *ast.DeleteStmt) {
	runInfo.expr = stmt.Item
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}
	item := runInfo.rv

	if stmt.Key != nil {
		runInfo.expr = stmt.Key
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
	}

	if item.Kind() == reflect.Interface && !item.IsNil() {
		item = item.Elem()
	}

	switch item.Kind() {
	case reflect.String:
		if stmt.Key != nil && runInfo.rv.Kind() == reflect.Bool && runInfo.rv.Bool() {
			runInfo.env.DeleteGlobal(item.String())
			runInfo.rv = nilValue
			return
		}
		runInfo.env.Delete(item.String())
		runInfo.rv = nilValue

	case reflect.Map:
		if stmt.Key == nil {
			runInfo.err = newStringError(stmt, "second argument to delete cannot be nil for map")
			runInfo.rv = nilValue
			return
		}
		if item.IsNil() {
			runInfo.rv = nilValue
			return
		}
		runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, item.Type().Key())
		if runInfo.err != nil {
			runInfo.err = newStringError(stmt, "cannot use type "+item.Type().Key().String()+" as type "+runInfo.rv.Type().String()+" in delete")
			runInfo.rv = nilValue
			return
		}
		item.SetMapIndex(runInfo.rv, reflect.Value{})
		runInfo.rv = nilValue
	default:
		runInfo.err = newStringError(stmt, "first argument to delete cannot be type "+item.Kind().String())
		runInfo.rv = nilValue
	}
}

func (runInfo *runInfoStruct) runCloseStmt(stmt *ast.CloseStmt) {
	runInfo.expr = stmt.Expr
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}
	if runInfo.rv.Kind() == reflect.Chan {
		runInfo.rv.Close()
		runInfo.rv = nilValue
		return
	}
	runInfo.err = newStringError(stmt, "type cannot be "+runInfo.rv.Kind().String()+" for close")
	runInfo.rv = nilValue
}

func (runInfo *runInfoStruct) runChanStmt(stmt *ast.ChanStmt) {
	runInfo.expr = stmt.RHS
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}
	if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
		runInfo.rv = runInfo.rv.Elem()
	}

	if runInfo.rv.Kind() != reflect.Chan {
		// rhs is not channel
		runInfo.err = newStringError(stmt, "receive from non-chan type "+runInfo.rv.Kind().String())
		runInfo.rv = nilValue
		return
	}

	// rhs is channel
	// receive from rhs channel
	rhs := runInfo.rv
	cases := []reflect.SelectCase{{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(runInfo.ctx.Done()),
	}, {
		Dir:  reflect.SelectRecv,
		Chan: rhs,
	}}
	var chosen int
	var ok bool
	chosen, runInfo.rv, ok = reflect.Select(cases)
	if chosen == 0 {
		runInfo.err = ErrInterrupt
		runInfo.rv = nilValue
		return
	}

	rhs = runInfo.rv // store rv in rhs temporarily

	if stmt.OkExpr != nil {
		// set ok to OkExpr
		if ok {
			runInfo.rv = trueValue
		} else {
			runInfo.rv = falseValue
		}
		runInfo.expr = stmt.OkExpr
		runInfo.invokeLetExpr()
		// TODO: ok to ignore error?
	}

	if ok {
		// set rv to lhs
		runInfo.rv = rhs
		runInfo.expr = stmt.LHS
		runInfo.invokeLetExpr()
		if runInfo.err != nil {
			return
		}
	} else {
		runInfo.rv = nilValue
	}
}
