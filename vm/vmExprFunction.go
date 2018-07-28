package vm

import (
	"fmt"
	"os"
	"reflect"

	"github.com/mattn/anko/ast"
)

// funcExpr creates a function that reflect Call can use.
// When called, it will run runVMFunction, to run the function statements
func funcExpr(funcExpr *ast.FuncExpr, env *Env) (reflect.Value, error) {
	// create the inTypes needed by reflect.FuncOf
	inTypes := make([]reflect.Type, len(funcExpr.Params), len(funcExpr.Params))
	for i := 0; i < len(inTypes); i++ {
		inTypes[i] = reflectValueType
	}
	if funcExpr.VarArg {
		inTypes[len(inTypes)-1] = interfaceSliceType
	}
	// create funcType, output is always slice of reflect.Type with two values
	funcType := reflect.FuncOf(inTypes, []reflect.Type{reflectValueType, reflectValueType}, funcExpr.VarArg)

	// create a function that can be used by reflect.MakeFunc
	// this function is a translator that converts a function call into a vm run
	// returns slice of reflect.Type with two values:
	// return value of the function and error value of the run
	runVMFunction := func(in []reflect.Value) []reflect.Value {
		var err error
		var rv reflect.Value

		// create newEnv for run
		newEnv := env.NewEnv()
		// add Params to newEnv, except last Params
		for i := 0; i < len(funcExpr.Params)-1; i++ {
			rv = in[i].Interface().(reflect.Value)
			err = newEnv.defineValue(funcExpr.Params[i], rv)
			if err != nil {
				return []reflect.Value{reflect.ValueOf(nilValue), reflect.ValueOf(reflect.ValueOf(newError(funcExpr, err)))}
			}
		}
		// add last Params to newEnv
		if len(funcExpr.Params) > 0 {
			if funcExpr.VarArg {
				// function is variadic, add last Params to newEnv without convert to Interface and then reflect.Value
				rv = in[len(funcExpr.Params)-1]
				err = newEnv.defineValue(funcExpr.Params[len(funcExpr.Params)-1], rv)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(nilValue), reflect.ValueOf(reflect.ValueOf(newError(funcExpr, err)))}
				}
			} else {
				// function is not variadic, add last Params to newEnv
				rv = in[len(funcExpr.Params)-1].Interface().(reflect.Value)
				err = newEnv.defineValue(funcExpr.Params[len(funcExpr.Params)-1], rv)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(nilValue), reflect.ValueOf(reflect.ValueOf(newError(funcExpr, err)))}
				}
			}
		}

		// run function statements
		rv, err = run(funcExpr.Stmts, newEnv)
		if err != nil && err != ErrReturn {
			err = newError(funcExpr, err)
			// return nil value and error
			// need to do single reflect.ValueOf because nilValue is already reflect.Value of nil
			// need to do double reflect.ValueOf of newError in order to match
			return []reflect.Value{reflect.ValueOf(nilValue), reflect.ValueOf(reflect.ValueOf(newError(funcExpr, err)))}
		}

		// the reflect.ValueOf of rv is needed to work in the reflect.Value slice
		// reflectValueErrorNilValue is already a double reflect.ValueOf
		return []reflect.Value{reflect.ValueOf(rv), reflectValueErrorNilValue}
	}

	// make the reflect.Value function that calls runVMFunction
	rv := reflect.MakeFunc(funcType, runVMFunction)

	// if function name is not empty, define it in the env
	if funcExpr.Name != "" {
		err := env.defineValue(funcExpr.Name, rv)
		if err != nil {
			return nilValue, newError(funcExpr, err)
		}
	}

	// return the reflect.Value created
	return rv, nil
}

// anonCallExpr handles ast.AnonCallExpr which calls a function anonymously
func anonCallExpr(e *ast.AnonCallExpr, env *Env) (reflect.Value, error) {
	f, err := invokeExpr(e.Expr, env)
	if err != nil {
		return nilValue, newError(e, err)
	}
	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	if f.Kind() == reflect.Func {
		return invokeExpr(&ast.CallExpr{Func: f, SubExprs: e.SubExprs, VarArg: e.VarArg, Go: e.Go}, env)
	}
	if !f.IsValid() {
		return nilValue, newStringError(e, "cannot call type invalid")
	}
	return nilValue, newStringError(e, "cannot call type "+f.Type().String())
}

// callExpr handles *ast.CallExpr which calls a function
func callExpr(callExpr *ast.CallExpr, env *Env) (rv reflect.Value, err error) {
	// Note that if the function type looks the same as the VM function type, the returned values will probably be wrong

	rv = nilValue

	f := callExpr.Func
	if !f.IsValid() {
		// if function is not valid try to get by function name
		f, err = env.get(callExpr.Name)
		if err != nil {
			err = newError(callExpr, err)
			return
		}
	}

	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	if !f.IsValid() {
		err = newStringError(callExpr, "cannot call type invalid")
		return
	}
	if f.Kind() != reflect.Func {
		err = newStringError(callExpr, "cannot call type "+f.Type().String())
		return
	}

	var rvs []reflect.Value
	var args []reflect.Value
	var useCallSlice bool
	fType := f.Type()
	// check if this is a runVMFunction type
	isRunVMFunction := checkIfRunVMFunction(fType)
	// create/convert the args to the function
	args, useCallSlice, err = makeCallArgs(fType, isRunVMFunction, callExpr, env)
	if err != nil {
		return
	}

	// capture panics if not in debug mode
	defer func() {
		if os.Getenv("ANKO_DEBUG") == "" {
			if recoverResult := recover(); recoverResult != nil {
				err = fmt.Errorf("%v", recoverResult)
			}
		}
	}()

	// useCallSlice lets us know to use CallSlice instead of Call because of the format of the args
	if useCallSlice {
		if callExpr.Go {
			go f.CallSlice(args)
			return
		}
		rvs = f.CallSlice(args)
	} else {
		if callExpr.Go {
			go f.Call(args)
			return
		}
		rvs = f.Call(args)
	}

	// TOFIX: how VM pointers/addressing work
	// Until then, this is a work around to set pointers back to VM variables
	// This will probably panic for some functions and/or calls that are variadic
	if !isRunVMFunction {
		for i, expr := range callExpr.SubExprs {
			if addrExpr, ok := expr.(*ast.AddrExpr); ok {
				if identExpr, ok := addrExpr.Expr.(*ast.IdentExpr); ok {
					invokeLetExpr(identExpr, args[i].Elem(), env)
				}
			}
		}
	}

	// processCallReturnValues to get/convert return values to normal rv form
	rv, err = processCallReturnValues(rvs, isRunVMFunction, true)

	return
}

// checkIfRunVMFunction checking the number and types of the reflect.Type.
// If it matches the types for a runVMFunction this will return true, otherwise false
func checkIfRunVMFunction(rt reflect.Type) bool {
	if rt.NumOut() != 2 || rt.Out(0) != reflectValueType || rt.Out(1) != reflectValueType {
		return false
	}
	if rt.NumIn() > 1 {
		if rt.IsVariadic() {
			if rt.In(rt.NumIn()-1) != interfaceSliceType {
				return false
			}
		} else {
			if rt.In(rt.NumIn()-1) != reflectValueType {
				return false
			}
		}
		for i := 0; i < rt.NumIn()-1; i++ {
			if rt.In(i) != reflectValueType {
				return false
			}
		}
	}
	return true
}

// makeCallArgs creates the arguments reflect.Value slice for the four diffrent kinds of functions.
// Also returns true if CallSlice should be used on the arguments, or false if Call should be used.
func makeCallArgs(rt reflect.Type, isRunVMFunction bool, callExpr *ast.CallExpr, env *Env) ([]reflect.Value, bool, error) {
	// number of arguments
	numIn := rt.NumIn()
	if numIn < 1 {
		// no arguments needed
		return []reflect.Value{}, false, nil
	}

	// number of expressions
	numExprs := len(callExpr.SubExprs)
	// checks to short circuit wrong number of arguments
	if (!rt.IsVariadic() && !callExpr.VarArg && numIn != numExprs) ||
		(rt.IsVariadic() && callExpr.VarArg && (numIn < numExprs || numIn > numExprs+1)) ||
		(rt.IsVariadic() && !callExpr.VarArg && numIn > numExprs+1) ||
		(!rt.IsVariadic() && callExpr.VarArg && numIn < numExprs) {
		return []reflect.Value{}, false, newStringError(callExpr, fmt.Sprintf("function wants %v arguments but received %v", numIn, numExprs))
	}
	if rt.IsVariadic() && rt.In(numIn-1).Kind() != reflect.Slice && rt.In(numIn-1).Kind() != reflect.Array {
		return []reflect.Value{}, false, newStringError(callExpr, "function is variadic but last parameter is of type "+rt.In(numIn-1).String())
	}

	var err error
	var arg reflect.Value
	var args []reflect.Value
	if numIn > numExprs {
		args = make([]reflect.Value, 0, numIn)
	} else {
		args = make([]reflect.Value, 0, numExprs)
	}
	indexIn := 0
	indexExpr := 0

	// create arguments except the last one
	for indexIn < numIn-1 && indexExpr < numExprs-1 {
		arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
		if err != nil {
			return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(arg))
		} else {
			arg, err = convertReflectValueToType(arg, rt.In(indexIn))
			if err != nil {
				return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
			}
			args = append(args, arg)
		}
		indexIn++
		indexExpr++
	}

	if !rt.IsVariadic() && !callExpr.VarArg {
		// function is not variadic and call is not variadic
		// add last arguments and return
		arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
		if err != nil {
			return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(arg))
		} else {
			arg, err = convertReflectValueToType(arg, rt.In(indexIn))
			if err != nil {
				return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
			}
			args = append(args, arg)
		}
		return args, false, nil
	}

	if !rt.IsVariadic() && callExpr.VarArg {
		// function is not variadic and call is variadic
		arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
		if err != nil {
			return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
		}
		if arg.Kind() != reflect.Slice && arg.Kind() != reflect.Array {
			return []reflect.Value{}, false, newStringError(callExpr, "call is variadic but last parameter is of type "+arg.Type().String())
		}
		if arg.Len() < numIn-indexIn {
			return []reflect.Value{}, false, newStringError(callExpr, fmt.Sprintf("function wants %v arguments but received %v", numIn, numExprs+arg.Len()-1))
		}

		indexSlice := 0
		for indexIn < numIn {
			if isRunVMFunction {
				args = append(args, reflect.ValueOf(arg.Index(indexSlice)))
			} else {
				arg, err = convertReflectValueToType(arg.Index(indexSlice), rt.In(indexIn))
				if err != nil {
					return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
						"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
				}
				args = append(args, arg)
			}
			indexIn++
			indexSlice++
		}
		return args, false, nil
	}

	// function is variadic and call may or may not be variadic

	if indexExpr == numExprs {
		// no more expressions, return what we have and let reflect Call handle if call is variadic or not
		return args, false, nil
	}

	if numIn > numExprs {
		// there are more arguments after this one, so does not matter if call is variadic or not
		// add the last argument then return what we have and let reflect Call handle if call is variadic or not
		arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
		if err != nil {
			return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(arg))
		} else {
			arg, err = convertReflectValueToType(arg, rt.In(indexIn))
			if err != nil {
				return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
			}
			args = append(args, arg)
		}
		return args, false, nil
	}

	if rt.IsVariadic() && !callExpr.VarArg {
		// function is variadic and call is not variadic
		sliceType := rt.In(numIn - 1).Elem()
		for indexExpr < numExprs {
			arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
			if err != nil {
				return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
			}
			arg, err = convertReflectValueToType(arg, sliceType)
			if err != nil {
				return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
			}
			args = append(args, arg)
			indexExpr++
		}
		return args, false, nil

	}

	// function is variadic and call is variadic
	// the only time we return CallSlice is true
	sliceType := rt.In(numIn - 1)
	if sliceType.Kind() == reflect.Interface && !arg.IsNil() {
		sliceType = sliceType.Elem()
	}
	arg, err = invokeExpr(callExpr.SubExprs[indexExpr], env)
	if err != nil {
		return []reflect.Value{}, false, newError(callExpr.SubExprs[indexExpr], err)
	}
	arg, err = convertReflectValueToType(arg, sliceType)
	if err != nil {
		return []reflect.Value{}, false, newStringError(callExpr.SubExprs[indexExpr],
			"function wants argument type "+rt.In(indexIn).String()+" but received type "+arg.Type().String())
	}
	args = append(args, arg)

	return args, true, nil
}

// processCallReturnValues get/converts the values returned from a function call into our normal reflect.Value, error
func processCallReturnValues(rvs []reflect.Value, isRunVMFunction bool, convertToInterfaceSlice bool) (reflect.Value, error) {
	// check if it is not runVMFunction
	if !isRunVMFunction {
		// the function was a Go function, convert to our normal reflect.Value, error
		switch len(rvs) {
		case 0:
			// no return values so return nil reflect.Value and nil error
			return nilValue, nil
		case 1:
			// one return value but need to add nil error
			return rvs[0], nil
		}
		if convertToInterfaceSlice {
			// need to convert from a slice of reflect.Value to slice of interface
			return reflectValueSlicetoInterfaceSlice(rvs), nil
		}
		// need to keep as slice of reflect.Value
		return reflect.ValueOf(rvs), nil
	}

	// is a runVMFunction, expect return in the runVMFunction format
	// convertToInterfaceSlice is ignored
	// some of the below checks probably can be removed because they are done in checkIfRunVMFunction

	if len(rvs) != 2 {
		return nilValue, fmt.Errorf("VM function did not return 2 values but returned %v values", len(rvs))
	}
	if !rvs[0].IsValid() {
		return nilValue, fmt.Errorf("VM function value 1 did not return reflect value type but returned invalid type")
	}
	if !rvs[1].IsValid() {
		return nilValue, fmt.Errorf("VM function value 2 did not return reflect value type but returned invalid type")
	}
	if rvs[0].Type() != reflectValueType {
		return nilValue, fmt.Errorf("VM function value 1 did not return reflect value type but returned %v type", rvs[0].Type().String())
	}
	if rvs[1].Type() != reflectValueType {
		return nilValue, fmt.Errorf("VM function value 2 did not return reflect value type but returned %v type", rvs[1].Type().String())
	}

	rvError := rvs[1].Interface().(reflect.Value)
	if !rvError.IsValid() {
		return nilValue, fmt.Errorf("VM function error type is invalid")
	}
	if rvError.Type() != errorType && rvError.Type() != vmErrorType {
		return nilValue, fmt.Errorf("VM function error type is %v", rvError.Type())
	}

	if rvError.IsNil() {
		// no error, so return the normal VM reflect.Value form
		return rvs[0].Interface().(reflect.Value), nil
	}

	// VM returns two types of errors, check to see which type
	if rvError.Type() == vmErrorType {
		// convert to VM *Error
		return nilValue, rvError.Interface().(*Error)
	}
	// convert to error
	return nilValue, rvError.Interface().(error)
}
