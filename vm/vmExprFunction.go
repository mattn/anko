package vm

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/mattn/anko/ast"
)

func FuncExpr(e *ast.FuncExpr, env *Env) (reflect.Value, error) {
	f := reflect.ValueOf(
		func(e *ast.FuncExpr, env *Env) Func {
			return func(args ...reflect.Value) (reflect.Value, error) {
				if !e.VarArg {
					if len(e.Params) != len(args) {
						return nilValue, NewStringError(e, fmt.Sprintf("expected %v function arguments but received %v", len(e.Params), len(args)))
					}
				}
				newenv := env.NewEnv()
				if e.VarArg {
					newenv.defineValue(e.Params[0], reflect.ValueOf(args))
				} else {
					for i, arg := range e.Params {
						newenv.defineValue(arg, args[i])
					}
				}
				rv, err := run(e.Stmts, newenv)
				if err == ReturnError {
					err = nil
				}
				return rv, err
			}
		}(e, env),
	)
	env.defineValue(e.Name, f)
	return f, nil
}

func AnonCallExpr(e *ast.AnonCallExpr, env *Env) (reflect.Value, error) {
	f, err := invokeExpr(e.Expr, env)
	if err != nil {
		return nilValue, NewError(e, err)
	}
	if f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		return nilValue, NewStringError(e, "cannot call type "+f.Type().String())
	}
	return invokeExpr(&ast.CallExpr{Func: f, SubExprs: e.SubExprs, VarArg: e.VarArg, Go: e.Go}, env)
}

func CallExpr(e *ast.CallExpr, env *Env) (reflect.Value, error) {
	var err error
	f := e.Func

	if !f.IsValid() {
		f, err = env.get(e.Name)
		if err != nil {
			return nilValue, err
		}
	}

	if !f.IsValid() {
		return nilValue, NewStringError(e, "cannot call type "+f.Type().String())
	}
	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}

	if f.Kind() != reflect.Func {
		return nilValue, NewStringError(e, "cannot call type "+f.Type().String())
	}

	_, isReflect := f.Interface().(Func)

	var arg reflect.Value
	args := []reflect.Value{}
	l := len(e.SubExprs)
	for i, expr := range e.SubExprs {

		arg, err = invokeExpr(expr, env)
		if err != nil {
			return nilValue, NewError(expr, err)
		}

		if i < f.Type().NumIn() {
			if !f.Type().IsVariadic() {
				iType := f.Type().In(i)
				if arg.Kind() == reflect.Interface && !arg.IsNil() {
					arg = arg.Elem()
				}
				if arg.Type() == unsafePointerType {
					arg = reflect.New(iType).Elem()
				}
				if !arg.IsValid() {
					arg = reflect.Zero(iType)
				} else if arg.Kind() == reflect.Func {
					if _, isFunc := arg.Interface().(Func); isFunc {
						rfunc := arg
						arg = reflect.MakeFunc(iType, func(args []reflect.Value) []reflect.Value {
							for i := range args {
								args[i] = reflect.ValueOf(args[i])
							}
							if e.Go {
								go func() {
									rfunc.Call(args)
								}()
								return []reflect.Value{}
							}
							var rets []reflect.Value
							for _, v := range rfunc.Call(args)[:iType.NumOut()] {
								rets = append(rets, v.Interface().(reflect.Value))
							}
							return rets
						})
					}
				} else if iType != interfaceType && arg.Type() != iType {
					if arg.Type().ConvertibleTo(iType) {
						arg = arg.Convert(iType)
					} else {
						return nilValue, NewStringError(expr, "argument type "+arg.Type().String()+" cannot be used for function argument type "+iType.String())
					}
				}
			}
		}

		if !arg.IsValid() {
			arg = nilValue
		}

		if isReflect {
			if arg.Kind() == reflect.Interface && !arg.IsNil() {
				arg = arg.Elem()
			}
			if e.VarArg && i == l-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, reflect.ValueOf(arg.Index(j).Elem()))
				}
			} else {
				args = append(args, reflect.ValueOf(arg))
			}
		} else {
			if e.VarArg && i == l-1 {
				for j := 0; j < arg.Len(); j++ {
					args = append(args, arg.Index(j).Elem())
				}
			} else {
				args = append(args, arg)
			}
		}

	}

	ret := nilValue
	fnc := func() {
		defer func() {
			if os.Getenv("ANKO_DEBUG") == "" {
				if ex := recover(); ex != nil {
					if e, ok := ex.(error); ok {
						err = e
					} else {
						err = errors.New(fmt.Sprint(ex))
					}
				}
			}
		}()
		if f.Kind() == reflect.Interface {
			f = f.Elem()
		}
		rets := f.Call(args)
		if isReflect {
			ev := rets[1].Interface()
			if ev != nil {
				err = ev.(error)
			}
			ret = rets[0].Interface().(reflect.Value)
		} else {
			for i, expr := range e.SubExprs {
				if ae, ok := expr.(*ast.AddrExpr); ok {
					if id, ok := ae.Expr.(*ast.IdentExpr); ok {
						invokeLetExpr(id, args[i].Elem().Elem(), env)
					}
				}
			}
			if f.Type().NumOut() == 0 {
			} else if f.Type().NumOut() == 1 {
				ret = rets[0]
			} else {
				var result []interface{}
				for _, r := range rets {
					result = append(result, r.Interface())
				}
				ret = reflect.ValueOf(result)
			}
		}
	}

	if e.Go {
		go fnc()
		return nilValue, nil
	}
	fnc()
	if err != nil {
		return nilValue, NewError(e, err)
	}
	return ret, nil
}
