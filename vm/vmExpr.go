package vm

import (
	"reflect"

	"github.com/mattn/anko/ast"
)

// invokeExpr evaluates one expression.
func (runInfo *runInfoStruct) invokeExpr() {
	switch expr := runInfo.expr.(type) {

	// OpExpr
	case *ast.OpExpr:
		runInfo.operator = expr.Op
		runInfo.invokeOperator()

	// IdentExpr
	case *ast.IdentExpr:
		runInfo.rv, runInfo.err = runInfo.env.get(expr.Lit)
		if runInfo.err != nil {
			runInfo.err = newError(expr, runInfo.err)
		}

	// LiteralExpr
	case *ast.LiteralExpr:
		runInfo.rv = expr.Literal

	// ArrayExpr
	case *ast.ArrayExpr:
		a := make([]interface{}, len(expr.Exprs))
		var i int
		for i, runInfo.expr = range expr.Exprs {
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			a[i] = runInfo.rv.Interface()
		}
		runInfo.rv = reflect.ValueOf(a)

	// MapExpr
	case *ast.MapExpr:
		var valueExpr ast.Expr
		var key reflect.Value
		m := make(map[interface{}]interface{}, len(expr.MapExpr))
		// TODO: overhead of map is not needed for MapExpr, can turn into two slices
		for runInfo.expr, valueExpr = range expr.MapExpr {
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			key = runInfo.rv

			runInfo.expr = valueExpr
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}

			m[key.Interface()] = runInfo.rv.Interface()
		}
		runInfo.rv = reflect.ValueOf(m)

	// DerefExpr
	case *ast.DerefExpr:
		runInfo.expr = expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if runInfo.rv.Kind() != reflect.Ptr {
			runInfo.err = newStringError(expr.Expr, "cannot deference non-pointer")
			return
		}
		runInfo.rv = runInfo.rv.Elem()

	// AddrExpr
	case *ast.AddrExpr:
		runInfo.expr = expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if runInfo.rv.CanAddr() {
			runInfo.rv = runInfo.rv.Addr()
		} else {
			i := runInfo.rv.Interface()
			runInfo.rv = reflect.ValueOf(&i)
		}

	// UnaryExpr
	case *ast.UnaryExpr:
		runInfo.expr = expr.Expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		switch expr.Operator {
		case "-":
			switch runInfo.rv.Kind() {
			case reflect.Int64:
				runInfo.rv = reflect.ValueOf(-runInfo.rv.Int())
			case reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Bool:
				runInfo.rv = reflect.ValueOf(-toInt64(runInfo.rv))
			case reflect.Float64:
				runInfo.rv = reflect.ValueOf(-runInfo.rv.Float())
			default:
				runInfo.rv = reflect.ValueOf(-toFloat64(runInfo.rv))
			}
		case "^":
			runInfo.rv = reflect.ValueOf(^toInt64(runInfo.rv))
		case "!":
			if toBool(runInfo.rv) {
				runInfo.rv = falseValue
			} else {
				runInfo.rv = trueValue
			}
		default:
			runInfo.err = newStringError(expr, "unknown operator")
		}

	// ParenExpr
	case *ast.ParenExpr:
		runInfo.expr = expr.SubExpr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

	// MemberExpr
	case *ast.MemberExpr:
		runInfo.expr = expr.Expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if runInfo.rv.Kind() == reflect.Interface {
			runInfo.rv = runInfo.rv.Elem()
		}
		if !runInfo.rv.IsValid() {
			runInfo.err = newStringError(expr, "type invalid does not support member operation")
		}
		if runInfo.rv.CanInterface() {
			if env, ok := runInfo.rv.Interface().(*Env); ok {
				runInfo.rv, runInfo.err = env.get(expr.Name)
				if runInfo.err != nil {
					runInfo.err = newError(expr, runInfo.err)
				}
				return
			}
		}

		method, found := runInfo.rv.Type().MethodByName(expr.Name)
		if found {
			runInfo.rv = runInfo.rv.Method(method.Index)
		}

		if runInfo.rv.Kind() == reflect.Ptr {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch runInfo.rv.Kind() {
		case reflect.Struct:
			field, found := runInfo.rv.Type().FieldByName(expr.Name)
			if found {
				runInfo.rv = runInfo.rv.FieldByIndex(field.Index)
			}
			if runInfo.rv.CanAddr() {
				runInfo.rv = runInfo.rv.Addr()
				method, found := runInfo.rv.Type().MethodByName(expr.Name)
				if found {
					runInfo.rv = runInfo.rv.Method(method.Index)
				}
			}
			runInfo.err = newStringError(expr, "no member named '"+expr.Name+"' for struct")
		case reflect.Map:
			runInfo.rv = getMapIndex(reflect.ValueOf(expr.Name), runInfo.rv)
			// Note if the map is of reflect.Value, it will incorrectly return nil when zero value
			if runInfo.rv == zeroValue {
				runInfo.rv = nilValue
			}
		default:
			runInfo.err = newStringError(expr, "type "+runInfo.rv.Kind().String()+" does not support member operation")
		}

	// ItemExpr
	case *ast.ItemExpr:
		runInfo.expr = expr.Item
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		item := runInfo.rv

		runInfo.expr = expr.Index
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}

		switch item.Kind() {
		case reflect.String, reflect.Slice, reflect.Array:
			var index int
			index, runInfo.err = tryToInt(runInfo.rv)
			if runInfo.err != nil {
				runInfo.err = newStringError(expr, "index must be a number")
				return
			}
			if index < 0 || index >= item.Len() {
				runInfo.err = newStringError(expr, "index out of range")
				return
			}
			if item.Kind() != reflect.String {
				runInfo.rv = item.Index(index)
			} else {
				// String
				runInfo.rv = item.Index(index).Convert(stringType)
			}
		case reflect.Map:
			runInfo.rv = getMapIndex(runInfo.rv, item)
			// Note if the map is of reflect.Value, it will incorrectly return nil when zero value
			if runInfo.rv == zeroValue {
				runInfo.rv = nilValue
			}
		default:
			runInfo.err = newStringError(expr, "type "+item.Kind().String()+" does not support index operation")
		}

	// SliceExpr
	case *ast.SliceExpr:
		runInfo.expr = expr.Item
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		item := runInfo.rv

		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}

		switch item.Kind() {
		case reflect.String, reflect.Slice, reflect.Array:
			var beginIndex, endIndex int
			if expr.Begin == nil {
				beginIndex = 0
			} else {
				runInfo.expr = expr.Begin
				runInfo.invokeExpr()
				if runInfo.err != nil {
					return
				}

				beginIndex, runInfo.err = tryToInt(runInfo.rv)
				if runInfo.err != nil {
					runInfo.err = newStringError(expr, "index must be a number")
					return
				}
				if beginIndex < 0 || beginIndex > item.Len() {
					runInfo.err = newStringError(expr, "index out of range")
					return
				}
			}
			if expr.End == nil {
				endIndex = item.Len()
			} else {
				runInfo.expr = expr.End
				runInfo.invokeExpr()
				if runInfo.err != nil {
					return
				}
				endIndex, runInfo.err = tryToInt(runInfo.rv)
				if runInfo.err != nil {
					runInfo.err = newStringError(expr, "index must be a number")
					return
				}
				if endIndex < 0 || endIndex > item.Len() {
					runInfo.err = newStringError(expr, "index out of range")
					return
				}
			}
			if beginIndex > endIndex {
				runInfo.err = newStringError(expr, "invalid slice index")
				return
			}
			runInfo.rv = item.Slice(beginIndex, endIndex)
		default:
			runInfo.err = newStringError(expr, "type "+item.Kind().String()+" does not support slice operation")
		}

	// LetsExpr
	case *ast.LetsExpr:
		var i int
		for i, runInfo.expr = range expr.RHSS {
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
				runInfo.rv = runInfo.rv.Elem()
			}
			if i < len(expr.LHSS) {
				runInfo.expr = expr.LHSS[i]
				runInfo.invokeLetExpr()
				if runInfo.err != nil {
					return
				}
			}

		}

	// TernaryOpExpr
	case *ast.TernaryOpExpr:
		runInfo.expr = expr.Expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if toBool(runInfo.rv) {
			runInfo.expr = expr.LHS
		} else {
			runInfo.expr = expr.RHS
		}
		runInfo.invokeExpr()

	// NilCoalescingOpExpr
	case *ast.NilCoalescingOpExpr:
		// if left side has no error and is not nil, returns left side
		// otherwise returns right side
		runInfo.expr = expr.LHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			if !isNil(runInfo.rv) {
				return
			}
		}
		runInfo.expr = expr.RHS
		runInfo.invokeExpr()

	// LenExpr
	case *ast.LenExpr:
		runInfo.expr = expr.Expr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch runInfo.rv.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
			runInfo.rv = reflect.ValueOf(int64(runInfo.rv.Len()))
		default:
			runInfo.err = newStringError(expr, "type "+runInfo.rv.Kind().String()+" does not support len operation")
		}

	// NewExpr
	case *ast.NewExpr:
		var t reflect.Type
		t, runInfo.err = getTypeFromString(runInfo.env, expr.Type)
		if runInfo.err != nil {
			return
		}
		if t == nil {
			runInfo.err = newErrorf(expr, "type cannot be nil for new")
			return
		}
		runInfo.rv = reflect.New(t)

	// MakeExpr
	case *ast.MakeExpr:
		var t reflect.Type
		t, runInfo.err = getTypeFromString(runInfo.env, expr.Type)
		if runInfo.err != nil {
			return
		}
		if t == nil {
			runInfo.err = newErrorf(expr, "type cannot be nil for make")
			return
		}

		for i := 1; i < expr.Dimensions; i++ {
			t = reflect.SliceOf(t)
		}
		if expr.Dimensions < 1 {
			runInfo.rv, runInfo.err = makeValue(t)
			return
		}

		var alen int
		if expr.LenExpr != nil {
			runInfo.expr = expr.LenExpr
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			alen = toInt(runInfo.rv)
		}

		var acap int
		if expr.CapExpr != nil {
			runInfo.expr = expr.CapExpr
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			acap = toInt(runInfo.rv)
		} else {
			acap = alen
		}

		runInfo.rv = reflect.MakeSlice(reflect.SliceOf(t), alen, acap)

	// MakeTypeExpr
	case *ast.MakeTypeExpr:
		runInfo.expr = expr.Type
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if !runInfo.rv.IsValid() || runInfo.rv.Type() == nil {
			runInfo.err = newErrorf(expr, "type cannot be nil for make type")
		}

		// if expr.Name has a dot in it, it should give a syntax error, so no needs to check err
		runInfo.env.DefineReflectType(expr.Name, runInfo.rv.Type())

		runInfo.rv = reflect.ValueOf(runInfo.rv.Type())

	// MakeChanExpr
	case *ast.MakeChanExpr:
		var t reflect.Type
		t, runInfo.err = getTypeFromString(runInfo.env, expr.Type)
		if runInfo.err != nil {
			return
		}
		if t == nil {
			runInfo.err = newErrorf(expr, "type cannot be nil for make chan")
			return
		}

		var size int
		if expr.SizeExpr != nil {
			runInfo.expr = expr.SizeExpr
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
			size = toInt(runInfo.rv)
		}

		runInfo.rv = reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), size)

	// ChanExpr
	case *ast.ChanExpr:
		runInfo.expr = expr.RHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		rhs := runInfo.rv

		if expr.LHS != nil {
			runInfo.expr = expr.LHS
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
		}

		if expr.LHS == nil || runInfo.rv.Kind() != reflect.Chan {
			// lhs is not channel
			if rhs.Kind() != reflect.Chan {
				// rhs is not channel
				runInfo.err = newStringError(expr, "invalid operation for chan")
				return
			}
			// receive from rhs channel
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
				return
			}
			if !ok {
				runInfo.rv = nilValue
			}
			// TOFIX: this runs the expr.LHS again, will sometimes cause bugs, for example with i++
			// most likely need a syntax change to expr = <- expr for let expr
			runInfo.expr = expr.LHS
			runInfo.invokeLetExpr()
			return
		}

		// send to lhs channel
		rhs, runInfo.err = convertReflectValueToType(rhs, runInfo.rv.Type().Elem())
		if runInfo.err != nil {
			runInfo.err = newStringError(expr, "cannot use type "+rhs.Type().String()+" as type "+runInfo.rv.Type().Elem().String()+" to send to chan")
			return
		}
		cases := []reflect.SelectCase{{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(runInfo.ctx.Done()),
		}, {
			Dir:  reflect.SelectSend,
			Chan: runInfo.rv,
			Send: rhs,
		}}
		if chosen, _, _ := reflect.Select(cases); chosen == 0 {
			runInfo.err = ErrInterrupt
			return
		}

	// FuncExpr
	case *ast.FuncExpr:
		runInfo.expr = expr
		runInfo.funcExpr()

	// AnonCallExpr
	case *ast.AnonCallExpr:
		runInfo.expr = expr
		runInfo.anonCallExpr()

	// CallExpr
	case *ast.CallExpr:
		runInfo.expr = expr
		runInfo.callExpr()

	// DeleteExpr
	case *ast.DeleteExpr:
		runInfo.expr = expr.WhatExpr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		whatExpr := runInfo.rv

		if expr.KeyExpr != nil {
			runInfo.expr = expr.KeyExpr
			runInfo.invokeExpr()
			if runInfo.err != nil {
				return
			}
		}

		if whatExpr.Kind() == reflect.Interface && !whatExpr.IsNil() {
			whatExpr = whatExpr.Elem()
		}

		switch whatExpr.Kind() {
		case reflect.String:
			if expr.KeyExpr == nil {
				runInfo.err = runInfo.env.Delete(whatExpr.String())
				if runInfo.err != nil {
					runInfo.err = newError(expr, runInfo.err)
				}
				return
			}
			if runInfo.rv.Kind() == reflect.Bool && runInfo.rv.Bool() {
				runInfo.err = runInfo.env.DeleteGlobal(whatExpr.String())
				if runInfo.err != nil {
					runInfo.err = newError(expr, runInfo.err)
				}
				return
			}
			runInfo.err = runInfo.env.Delete(whatExpr.String())
			if runInfo.err != nil {
				runInfo.err = newError(expr, runInfo.err)
			}

		case reflect.Map:
			if expr.KeyExpr == nil {
				runInfo.err = newStringError(expr, "second argument to delete cannot be nil for map")
				return
			}
			if whatExpr.IsNil() {
				return
			}
			runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, whatExpr.Type().Key())
			if runInfo.err != nil {
				runInfo.err = newStringError(expr, "cannot use type "+whatExpr.Type().Key().String()+" as type "+runInfo.rv.Type().String()+" in delete")
				return
			}
			whatExpr.SetMapIndex(runInfo.rv, reflect.Value{})
		default:
			runInfo.err = newStringError(expr, "first argument to delete cannot be type "+whatExpr.Kind().String())
		}

	// IncludeExpr
	case *ast.IncludeExpr:
		runInfo.expr = expr.ItemExpr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		itemExpr := runInfo.rv

		runInfo.expr = expr.ListExpr
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}

		if runInfo.rv.Kind() != reflect.Slice && runInfo.rv.Kind() != reflect.Array {
			runInfo.err = newStringError(expr, "second argument must be slice or array; but have "+runInfo.rv.Kind().String())
			return
		}

		for i := 0; i < runInfo.rv.Len(); i++ {
			if equal(itemExpr, runInfo.rv.Index(i)) {
				runInfo.rv = trueValue
				return
			}
		}
		runInfo.rv = falseValue

	default:
		runInfo.err = newStringError(expr, "unknown expression")
	}

}
