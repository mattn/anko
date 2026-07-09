package vm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// funcExpr creates a function that reflect Call can use.
// When called, it will run runVMFunction, to run the function statements
func (runInfo *runInfoStruct) funcExpr() {
	funcExpr := runInfo.expr.(*ast.FuncExpr)

	// for adding env into saved function
	envFunc := runInfo.env
	options := runInfo.options

	// runVMFunc runs the function statements with args as the parameter values.
	// Returns the VM value and a reflect.Value holding the error of the run,
	// which is of type error or *Error and nil valued when the run succeeded.
	runVMFunc := func(ctx context.Context, args []reflect.Value) (reflect.Value, reflect.Value) {
		runInfo := runInfoStruct{ctx: ctx, options: options, env: envFunc.NewEnv(), stmt: funcExpr.Stmt, rv: nilValue}

		// add Params to newEnv
		for i, param := range funcExpr.Params {
			runInfo.env.DefineValue(param, args[i])
		}

		// run function statements
		runInfo.runSingleStmt()
		if runInfo.err != nil && runInfo.err != ErrReturn {
			return nilValue, reflect.ValueOf(newError(funcExpr, runInfo.err))
		}

		return runInfo.rv, errorNilValue
	}

	if !funcExpr.VarArg {
		// non-variadic functions with few parameters are created with a
		// concrete signature, so callExpr can call them directly without
		// going through the reflect.Call machinery.
		// the signatures match what reflect.MakeFunc would have created, so
		// checkIfRunVMFunction and reflect.Call still work on these values.
		switch len(funcExpr.Params) {
		case 0:
			runInfo.rv = reflect.ValueOf(func(ctx context.Context) (reflect.Value, reflect.Value) {
				return runVMFunc(ctx, nil)
			})
		case 1:
			runInfo.rv = reflect.ValueOf(func(ctx context.Context, arg0 reflect.Value) (reflect.Value, reflect.Value) {
				return runVMFunc(ctx, []reflect.Value{arg0})
			})
		case 2:
			runInfo.rv = reflect.ValueOf(func(ctx context.Context, arg0, arg1 reflect.Value) (reflect.Value, reflect.Value) {
				return runVMFunc(ctx, []reflect.Value{arg0, arg1})
			})
		case 3:
			runInfo.rv = reflect.ValueOf(func(ctx context.Context, arg0, arg1, arg2 reflect.Value) (reflect.Value, reflect.Value) {
				return runVMFunc(ctx, []reflect.Value{arg0, arg1, arg2})
			})
		case 4:
			runInfo.rv = reflect.ValueOf(func(ctx context.Context, arg0, arg1, arg2, arg3 reflect.Value) (reflect.Value, reflect.Value) {
				return runVMFunc(ctx, []reflect.Value{arg0, arg1, arg2, arg3})
			})
		}
		if len(funcExpr.Params) <= 4 {
			// if function name is not empty, define it in the env
			if funcExpr.Name != "" {
				runInfo.env.DefineValue(funcExpr.Name, runInfo.rv)
			}
			return
		}
	}

	// variadic functions and functions with many parameters go through
	// reflect.MakeFunc

	// create the inTypes needed by reflect.FuncOf
	inTypes := make([]reflect.Type, len(funcExpr.Params)+1)
	// for runVMFunction first arg is always context
	inTypes[0] = contextType
	for i := 1; i < len(inTypes); i++ {
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
		args := make([]reflect.Value, len(funcExpr.Params))

		// get Params, except last Params
		for i := 0; i < len(funcExpr.Params)-1; i++ {
			args[i] = in[i+1].Interface().(reflect.Value)
		}
		// get last Params
		if len(funcExpr.Params) > 0 {
			if funcExpr.VarArg {
				// function is variadic, get last Params without convert to Interface and then reflect.Value
				args[len(funcExpr.Params)-1] = in[len(funcExpr.Params)]
			} else {
				args[len(funcExpr.Params)-1] = in[len(funcExpr.Params)].Interface().(reflect.Value)
			}
		}

		rv, errV := runVMFunc(in[0].Interface().(context.Context), args)
		if !errV.IsNil() {
			// return nil value and error
			// need to do double reflect.ValueOf of the error in order to match
			return []reflect.Value{reflectValueNilValue, reflect.ValueOf(errV)}
		}

		// the reflect.ValueOf of rv is needed to work in the reflect.Value slice
		// reflectValueErrorNilValue is already a double reflect.ValueOf
		return []reflect.Value{reflect.ValueOf(rv), reflectValueErrorNilValue}
	}

	// make the reflect.Value function that calls runVMFunction
	runInfo.rv = reflect.MakeFunc(funcType, runVMFunction)

	// if function name is not empty, define it in the env
	if funcExpr.Name != "" {
		runInfo.env.DefineValue(funcExpr.Name, runInfo.rv)
	}
}

// anonCallExpr handles ast.AnonCallExpr which calls a function anonymously
func (runInfo *runInfoStruct) anonCallExpr() {
	anonCallExpr := runInfo.expr.(*ast.AnonCallExpr)

	runInfo.expr = anonCallExpr.Expr
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return
	}

	if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
		runInfo.rv = runInfo.rv.Elem()
	}
	if runInfo.rv.Kind() != reflect.Func {
		runInfo.err = newStringError(anonCallExpr, "cannot call type "+runInfo.rv.Kind().String())
		runInfo.rv = nilValue
		return
	}

	runInfo.expr = &ast.CallExpr{Func: runInfo.rv, SubExprs: anonCallExpr.SubExprs, VarArg: anonCallExpr.VarArg, Go: anonCallExpr.Go}
	runInfo.expr.SetPosition(anonCallExpr.Expr.Position())
	runInfo.invokeExpr()
}

// callExpr handles *ast.CallExpr which calls a function
func (runInfo *runInfoStruct) callExpr() {
	// Note that if the function type looks the same as the VM function type, the returned values will probably be wrong

	callExpr := runInfo.expr.(*ast.CallExpr)

	f := callExpr.Func
	if !f.IsValid() {
		// if function is not valid try to get by function name
		f, runInfo.err = runInfo.env.GetValue(callExpr.Name)
		if runInfo.err != nil {
			runInfo.err = newError(callExpr, runInfo.err)
			runInfo.rv = nilValue
			return
		}
	}

	if f.Kind() == reflect.Interface && !f.IsNil() {
		f = f.Elem()
	}
	if f.Kind() != reflect.Func {
		runInfo.err = newStringError(callExpr, "cannot call type "+f.Kind().String())
		runInfo.rv = nilValue
		return
	}

	var rvs []reflect.Value
	var args []reflect.Value
	var useCallSlice bool
	fType := f.Type()
	// check if this is a runVMFunction type
	isRunVMFunction := checkIfRunVMFunction(fType)

	// fast path: VM functions created with a concrete signature can be
	// called directly, without going through the reflect.Call machinery
	if isRunVMFunction && !callExpr.VarArg && !fType.IsVariadic() && fType.NumIn()-1 == len(callExpr.SubExprs) {
		if runInfo.callVMFunctionDirect(f, callExpr) {
			return
		}
	}

	// create/convert the args to the function
	args, useCallSlice = runInfo.makeCallArgs(fType, isRunVMFunction, callExpr)
	if runInfo.err != nil {
		return
	}

	if !runInfo.options.Debug {
		// captures panic
		defer recoverFunc(runInfo)
	}

	runInfo.rv = nilValue

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
					runInfo.rv = args[i].Elem()
					runInfo.expr = identExpr
					runInfo.invokeLetExpr()
				}
			}
		}
	}

	// processCallReturnValues to get/convert return values to normal rv form
	runInfo.rv, runInfo.err = processCallReturnValues(rvs, isRunVMFunction, true)
}

// callVMFunctionDirect calls a VM function that was created with a concrete
// signature in funcExpr directly, skipping the reflect.Call machinery.
// Returns false without evaluating any arguments if the function does not
// have one of the concrete signatures, so callExpr can fall back to the
// reflect based path. The caller must check that the number of arguments
// matches the function signature.
// handled is a named return so it stays true when recoverFunc recovers a
// panic from the called function.
func (runInfo *runInfoStruct) callVMFunctionDirect(f reflect.Value, callExpr *ast.CallExpr) (handled bool) {
	// check the concrete signature before evaluating the arguments, so the
	// fallback path does not evaluate expressions twice
	var fn0 func(context.Context) (reflect.Value, reflect.Value)
	var fn1 func(context.Context, reflect.Value) (reflect.Value, reflect.Value)
	var fn2 func(context.Context, reflect.Value, reflect.Value) (reflect.Value, reflect.Value)
	var fn3 func(context.Context, reflect.Value, reflect.Value, reflect.Value) (reflect.Value, reflect.Value)
	var fn4 func(context.Context, reflect.Value, reflect.Value, reflect.Value, reflect.Value) (reflect.Value, reflect.Value)
	switch fn := f.Interface().(type) {
	case func(context.Context) (reflect.Value, reflect.Value):
		fn0 = fn
	case func(context.Context, reflect.Value) (reflect.Value, reflect.Value):
		fn1 = fn
	case func(context.Context, reflect.Value, reflect.Value) (reflect.Value, reflect.Value):
		fn2 = fn
	case func(context.Context, reflect.Value, reflect.Value, reflect.Value) (reflect.Value, reflect.Value):
		fn3 = fn
	case func(context.Context, reflect.Value, reflect.Value, reflect.Value, reflect.Value) (reflect.Value, reflect.Value):
		fn4 = fn
	default:
		return false
	}

	handled = true

	// evaluate the arguments
	var argsBuf [4]reflect.Value
	args := argsBuf[:0]
	for _, subExpr := range callExpr.SubExprs {
		runInfo.expr = subExpr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return true
		}
		args = append(args, runInfo.rv)
	}

	if !runInfo.options.Debug {
		// captures panic
		defer recoverFunc(runInfo)
	}

	runInfo.rv = nilValue

	if callExpr.Go {
		switch {
		case fn0 != nil:
			go fn0(runInfo.ctx)
		case fn1 != nil:
			go fn1(runInfo.ctx, args[0])
		case fn2 != nil:
			go fn2(runInfo.ctx, args[0], args[1])
		case fn3 != nil:
			go fn3(runInfo.ctx, args[0], args[1], args[2])
		case fn4 != nil:
			go fn4(runInfo.ctx, args[0], args[1], args[2], args[3])
		}
		return true
	}

	var rv, errV reflect.Value
	switch {
	case fn0 != nil:
		rv, errV = fn0(runInfo.ctx)
	case fn1 != nil:
		rv, errV = fn1(runInfo.ctx, args[0])
	case fn2 != nil:
		rv, errV = fn2(runInfo.ctx, args[0], args[1])
	case fn3 != nil:
		rv, errV = fn3(runInfo.ctx, args[0], args[1], args[2])
	case fn4 != nil:
		rv, errV = fn4(runInfo.ctx, args[0], args[1], args[2], args[3])
	}

	if !errV.IsNil() {
		// same error conversion as processCallReturnValues
		if errV.Type() == vmErrorType {
			runInfo.err = errV.Interface().(*Error)
		} else {
			runInfo.err = errV.Interface().(error)
		}
		return true
	}

	runInfo.rv = rv
	return true
}

// checkIfRunVMFunction checking the number and types of the reflect.Type.
// If it matches the types for a runVMFunction this will return true, otherwise false
func checkIfRunVMFunction(rt reflect.Type) bool {
	if rt.NumIn() < 1 || rt.NumOut() != 2 || rt.In(0) != contextType || rt.Out(0) != reflectValueType || rt.Out(1) != reflectValueType {
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
		for i := 1; i < rt.NumIn()-1; i++ {
			if rt.In(i) != reflectValueType {
				return false
			}
		}
	}
	return true
}

// makeCallArgs creates the arguments reflect.Value slice for the four different kinds of functions.
// Also returns true if CallSlice should be used on the arguments, or false if Call should be used.
func (runInfo *runInfoStruct) makeCallArgs(rt reflect.Type, isRunVMFunction bool, callExpr *ast.CallExpr) ([]reflect.Value, bool) {
	// number of arguments
	numInReal := rt.NumIn()
	numIn := numInReal
	if isRunVMFunction {
		// for runVMFunction the first arg is context so does not count against number of SubExprs
		numIn--
	}
	if numIn < 1 {
		// no arguments needed
		if isRunVMFunction {
			// for runVMFunction first arg is always context
			return []reflect.Value{reflect.ValueOf(runInfo.ctx)}, false
		}
		return []reflect.Value{}, false
	}

	// number of expressions
	numExprs := len(callExpr.SubExprs)
	// checks to short circuit wrong number of arguments
	if (!rt.IsVariadic() && !callExpr.VarArg && numIn != numExprs) ||
		(rt.IsVariadic() && callExpr.VarArg && (numIn < numExprs || numIn > numExprs+1)) ||
		(rt.IsVariadic() && !callExpr.VarArg && numIn > numExprs+1) ||
		(!rt.IsVariadic() && callExpr.VarArg && numIn < numExprs) {
		runInfo.err = newStringError(callExpr, fmt.Sprintf("function wants %v arguments but received %v", numIn, numExprs))
		runInfo.rv = nilValue
		return nil, false
	}
	if rt.IsVariadic() && rt.In(numInReal-1).Kind() != reflect.Slice && rt.In(numInReal-1).Kind() != reflect.Array {
		runInfo.err = newStringError(callExpr, "function is variadic but last parameter is of type "+rt.In(numInReal-1).String())
		runInfo.rv = nilValue
		return nil, false
	}

	var args []reflect.Value
	indexIn := 0
	indexInReal := 0
	indexExpr := 0

	if numInReal > numExprs {
		args = make([]reflect.Value, 0, numInReal)
	} else {
		args = make([]reflect.Value, 0, numExprs)
	}
	if isRunVMFunction {
		// for runVMFunction first arg is always context
		args = append(args, reflect.ValueOf(runInfo.ctx))
		indexInReal++
	}

	// create arguments except the last one
	for indexInReal < numInReal-1 && indexExpr < numExprs-1 {
		runInfo.expr = callExpr.SubExprs[indexExpr]
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return nil, false
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(runInfo.rv))
		} else {
			runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, rt.In(indexInReal))
			if runInfo.err != nil {
				runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
				runInfo.rv = nilValue
				return nil, false
			}
			args = append(args, runInfo.rv)
		}
		indexIn++
		indexInReal++
		indexExpr++
	}

	if !rt.IsVariadic() && !callExpr.VarArg {
		// function is not variadic and call is not variadic
		// add last arguments and return
		runInfo.expr = callExpr.SubExprs[indexExpr]
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return nil, false
		}
		if runInfo.err != nil {
			return nil, false
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(runInfo.rv))
		} else {
			runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, rt.In(indexInReal))
			if runInfo.err != nil {
				runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
				runInfo.rv = nilValue
				return nil, false
			}
			args = append(args, runInfo.rv)
		}
		return args, false
	}

	if !rt.IsVariadic() && callExpr.VarArg {
		// function is not variadic and call is variadic
		runInfo.expr = callExpr.SubExprs[indexExpr]
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return nil, false
		}
		if runInfo.rv.Kind() != reflect.Slice && runInfo.rv.Kind() != reflect.Array {
			runInfo.err = newStringError(callExpr, "call is variadic but last parameter is of type "+runInfo.rv.Type().String())
			runInfo.rv = nilValue
			return nil, false
		}
		if runInfo.rv.Len() < numIn-indexIn {
			runInfo.err = newStringError(callExpr, fmt.Sprintf("function wants %v arguments but received %v", numIn, numExprs+runInfo.rv.Len()-1))
			runInfo.rv = nilValue
			return nil, false
		}

		indexSlice := 0
		for indexInReal < numInReal {
			if isRunVMFunction {
				args = append(args, reflect.ValueOf(runInfo.rv.Index(indexSlice)))
			} else {
				runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv.Index(indexSlice), rt.In(indexInReal))
				if runInfo.err != nil {
					runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
						"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
					runInfo.rv = nilValue
					return nil, false
				}
				args = append(args, runInfo.rv)
			}
			indexIn++
			indexInReal++
			indexSlice++
		}
		return args, false
	}

	// function is variadic and call may or may not be variadic

	if indexExpr == numExprs {
		// no more expressions, return what we have and let reflect Call handle if call is variadic or not
		return args, false
	}

	if numIn > numExprs {
		// there are more arguments after this one, so does not matter if call is variadic or not
		// add the last argument then return what we have and let reflect Call handle if call is variadic or not
		runInfo.expr = callExpr.SubExprs[indexExpr]
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return nil, false
		}
		if isRunVMFunction {
			args = append(args, reflect.ValueOf(runInfo.rv))
		} else {
			runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, rt.In(indexInReal))
			if runInfo.err != nil {
				runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
				runInfo.rv = nilValue
				return nil, false
			}
			args = append(args, runInfo.rv)
		}
		return args, false
	}

	if rt.IsVariadic() && !callExpr.VarArg {
		// function is variadic and call is not variadic
		sliceType := rt.In(numInReal - 1).Elem()
		for indexExpr < numExprs {
			runInfo.expr = callExpr.SubExprs[indexExpr]
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return nil, false
			}
			runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, sliceType)
			if runInfo.err != nil {
				runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
					"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
				runInfo.rv = nilValue
				return nil, false
			}
			args = append(args, runInfo.rv)
			indexExpr++
		}
		return args, false

	}

	// function is variadic and call is variadic
	// the only time we return CallSlice is true
	sliceType := rt.In(numInReal - 1)
	if sliceType.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
		sliceType = sliceType.Elem()
	}
	runInfo.expr = callExpr.SubExprs[indexExpr]
	runInfo.invokeExpr()
	if runInfo.err != nil {
		return nil, false
	}
	runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, sliceType)
	if runInfo.err != nil {
		runInfo.err = newStringError(callExpr.SubExprs[indexExpr],
			"function wants argument type "+rt.In(indexInReal).String()+" but received type "+runInfo.rv.Type().String())
		runInfo.rv = nilValue
		return nil, false
	}
	args = append(args, runInfo.rv)

	return args, true
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
	if rvs[0].Type() != reflectValueType {
		return nilValue, fmt.Errorf("VM function value 1 did not return reflect value type but returned %v type", rvs[0].Type().String())
	}
	if rvs[1].Type() != reflectValueType {
		return nilValue, fmt.Errorf("VM function value 2 did not return reflect value type but returned %v type", rvs[1].Type().String())
	}

	rvError := rvs[1].Interface().(reflect.Value)
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
