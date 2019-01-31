package vm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

// invokeExpr evaluates one expression.
func invokeExpr(ctx context.Context, expr ast.Expr, env *Env) (reflect.Value, error) {
	switch e := expr.(type) {

	// OpExpr
	case *ast.OpExpr:
		rv, err := invokeOperator(ctx, e.Op, env)
		if err != nil {
			return rv, newError(e.Op, err)
		}
		return rv, nil

	// IdentExpr
	case *ast.IdentExpr:
		return env.get(e.Lit)

	// LiteralExpr
	case *ast.LiteralExpr:
		return e.Literal, nil

	// ArrayExpr
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		for i, expr := range e.Exprs {
			arg, err := invokeExpr(ctx, expr, env)
			if err != nil {
				return nilValue, newError(expr, err)
			}
			a[i] = arg.Interface()
		}
		return reflect.ValueOf(a), nil

	// MapExpr
	case *ast.MapExpr:
		var err error
		var key reflect.Value
		var value reflect.Value
		m := make(map[interface{}]interface{}, len(e.MapExpr))
		for keyExpr, valueExpr := range e.MapExpr {
			key, err = invokeExpr(ctx, keyExpr, env)
			if err != nil {
				return nilValue, newError(keyExpr, err)
			}
			value, err = invokeExpr(ctx, valueExpr, env)
			if err != nil {
				return nilValue, newError(valueExpr, err)
			}
			m[key.Interface()] = value.Interface()
		}
		return reflect.ValueOf(m), nil

	// DerefExpr
	case *ast.DerefExpr:
		v := nilValue
		var err error
		switch ee := e.Expr.(type) {

		case *ast.IdentExpr:
			v, err = env.get(ee.Lit)
			if err != nil {
				return nilValue, newError(e, err)
			}

		case *ast.MemberExpr:
			v, err := invokeExpr(ctx, ee.Expr, env)
			if err != nil {
				return nilValue, newError(ee.Expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.get(ee.Name)
					if !m.IsValid() || err != nil {
						return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", ee.Name))
					}
					return m, nil
				}
			}

			m := v.MethodByName(ee.Name)
			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					field, found := v.Type().FieldByName(ee.Name)
					if !found {
						return nilValue, newStringError(e, "no member named '"+ee.Name+"' for struct")
					}
					return v.FieldByIndex(field.Index), nil
				} else if v.Kind() == reflect.Map {
					// From reflect MapIndex:
					// It returns the zero Value if key is not found in the map or if v represents a nil map.
					m = v.MapIndex(reflect.ValueOf(ee.Name))
				} else {
					return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return nilValue, newStringError(e, "invalid operation for the value")
		}
		if v.Kind() != reflect.Ptr {
			return nilValue, newStringError(e, "cannot deference for the value")
		}
		return v.Elem(), nil

	// AddrExpr
	case *ast.AddrExpr:
		v := nilValue
		var err error
		switch ee := e.Expr.(type) {

		case *ast.IdentExpr:
			v, err = env.get(ee.Lit)
			if err != nil {
				return nilValue, newError(e, err)
			}

		case *ast.MemberExpr:
			v, err := invokeExpr(ctx, ee.Expr, env)
			if err != nil {
				return nilValue, newError(ee.Expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.get(ee.Name)
					if !m.IsValid() || err != nil {
						return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", ee.Name))
					}
					return m, nil
				}
			}

			m := v.MethodByName(ee.Name)
			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					m = v.FieldByName(ee.Name)
					if !m.IsValid() {
						return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", ee.Name))
					}
				} else if v.Kind() == reflect.Map {
					// From reflect MapIndex:
					// It returns the zero Value if key is not found in the map or if v represents a nil map.
					m = v.MapIndex(reflect.ValueOf(ee.Name))
				} else {
					return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return nilValue, newStringError(e, "invalid operation for the value")
		}
		if !v.CanAddr() {
			i := v.Interface()
			return reflect.ValueOf(&i), nil
		}
		return v.Addr(), nil

	// UnaryExpr
	case *ast.UnaryExpr:
		v, err := invokeExpr(ctx, e.Expr, env)
		if err != nil {
			return nilValue, newError(e.Expr, err)
		}
		switch e.Operator {
		case "-":
			if v.Kind() == reflect.Int64 {
				return reflect.ValueOf(-v.Int()), nil
			}
			if v.Kind() == reflect.Float64 {
				return reflect.ValueOf(-v.Float()), nil
			}
			return reflect.ValueOf(-toFloat64(v)), nil
		case "^":
			return reflect.ValueOf(^toInt64(v)), nil
		case "!":
			if toBool(v) {
				return falseValue, nil
			}
			return trueValue, nil
		default:
			return nilValue, newStringError(e, "unknown operator")
		}

	// ParenExpr
	case *ast.ParenExpr:
		v, err := invokeExpr(ctx, e.SubExpr, env)
		if err != nil {
			return nilValue, newError(e.SubExpr, err)
		}
		return v, nil

	// MemberExpr
	case *ast.MemberExpr:
		v, err := invokeExpr(ctx, e.Expr, env)
		if err != nil {
			return nilValue, newError(e.Expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if !v.IsValid() {
			return nilValue, newStringError(e, "type invalid does not support member operation")
		}
		if v.CanInterface() {
			if vme, ok := v.Interface().(*Env); ok {
				m, err := vme.get(e.Name)
				if err != nil || !m.IsValid() {
					return nilValue, newStringError(e, fmt.Sprintf("invalid operation '%s'", e.Name))
				}
				return m, nil
			}
		}

		method, found := v.Type().MethodByName(e.Name)
		if found {
			return v.Method(method.Index), nil
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Struct:
			field, found := v.Type().FieldByName(e.Name)
			if found {
				return v.FieldByIndex(field.Index), nil
			}
			if v.CanAddr() {
				v = v.Addr()
				method, found := v.Type().MethodByName(e.Name)
				if found {
					return v.Method(method.Index), nil
				}
			}
			return nilValue, newStringError(e, "no member named '"+e.Name+"' for struct")
		case reflect.Map:
			v = getMapIndex(reflect.ValueOf(e.Name), v)
			// Note if the map is of reflect.Value, it will incorrectly return nil when zero value
			if v == zeroValue {
				return nilValue, nil
			}
			return v, nil
		default:
			return nilValue, newStringError(e, "type "+v.Kind().String()+" does not support member operation")
		}

	// ItemExpr
	case *ast.ItemExpr:
		v, err := invokeExpr(ctx, e.Value, env)
		if err != nil {
			return nilValue, newError(e.Value, err)
		}
		i, err := invokeExpr(ctx, e.Index, env)
		if err != nil {
			return nilValue, newError(e.Index, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.String, reflect.Slice, reflect.Array:
			ii, err := tryToInt(i)
			if err != nil {
				return nilValue, newStringError(e, "index must be a number")
			}
			if ii < 0 || ii >= v.Len() {
				return nilValue, newStringError(e, "index out of range")
			}
			if v.Kind() != reflect.String {
				return v.Index(ii), nil
			}
			// String
			return v.Index(ii).Convert(stringType), nil
		case reflect.Map:
			v = getMapIndex(i, v)
			// Note if the map is of reflect.Value, it will incorrectly return nil when zero value
			if v == zeroValue {
				return nilValue, nil
			}
			return v, nil
		default:
			return nilValue, newStringError(e, "type "+v.Kind().String()+" does not support index operation")
		}

	// SliceExpr
	case *ast.SliceExpr:
		v, err := invokeExpr(ctx, e.Value, env)
		if err != nil {
			return nilValue, newError(e.Value, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.String, reflect.Slice, reflect.Array:
			var rbi, rei int
			if e.Begin != nil {
				rb, err := invokeExpr(ctx, e.Begin, env)
				if err != nil {
					return nilValue, newError(e.Begin, err)
				}
				rbi, err = tryToInt(rb)
				if err != nil {
					return nilValue, newStringError(e, "index must be a number")
				}
				if rbi < 0 || rbi > v.Len() {
					return nilValue, newStringError(e, "index out of range")
				}
			} else {
				rbi = 0
			}
			if e.End != nil {
				re, err := invokeExpr(ctx, e.End, env)
				if err != nil {
					return nilValue, newError(e.End, err)
				}
				rei, err = tryToInt(re)
				if err != nil {
					return nilValue, newStringError(e, "index must be a number")
				}
				if rei < 0 || rei > v.Len() {
					return nilValue, newStringError(e, "index out of range")
				}
			} else {
				rei = v.Len()
			}
			if rbi > rei {
				return nilValue, newStringError(e, "invalid slice index")
			}
			return v.Slice(rbi, rei), nil
		default:
			return nilValue, newStringError(e, "type "+v.Kind().String()+" does not support slice operation")
		}

	// LetsExpr
	case *ast.LetsExpr:
		var err error
		rvs := make([]reflect.Value, len(e.RHSS))
		for i, rhs := range e.RHSS {
			rvs[i], err = invokeExpr(ctx, rhs, env)
			if err != nil {
				return nilValue, newError(rhs, err)
			}
		}
		for i, lhs := range e.LHSS {
			if i >= len(rvs) {
				break
			}
			v := rvs[i]
			if v.Kind() == reflect.Interface && !v.IsNil() {
				v = v.Elem()
			}
			_, err = invokeLetExpr(ctx, lhs, v, env)
			if err != nil {
				return nilValue, newError(lhs, err)
			}
		}
		return rvs[len(rvs)-1], nil

	// TernaryOpExpr
	case *ast.TernaryOpExpr:
		rv, err := invokeExpr(ctx, e.Expr, env)
		if err != nil {
			return nilValue, newError(e.Expr, err)
		}
		if toBool(rv) {
			lhsV, err := invokeExpr(ctx, e.LHS, env)
			if err != nil {
				return nilValue, newError(e.LHS, err)
			}
			return lhsV, nil
		}
		rhsV, err := invokeExpr(ctx, e.RHS, env)
		if err != nil {
			return nilValue, newError(e.RHS, err)
		}
		return rhsV, nil

	// NilCoalescingOpExpr
	case *ast.NilCoalescingOpExpr:
		var err error
		rv, _ := invokeExpr(ctx, e.LHS, env)
		if toBool(rv) {
			return rv, nil
		}
		rv, err = invokeExpr(ctx, e.RHS, env)
		if err != nil {
			return nilValue, newError(e.RHS, err)
		}
		return rv, nil

	// LenExpr
	case *ast.LenExpr:
		rv, err := invokeExpr(ctx, e.Expr, env)
		if err != nil {
			return nilValue, newError(e.Expr, err)
		}

		if rv.Kind() == reflect.Interface && !rv.IsNil() {
			rv = rv.Elem()
		}

		switch rv.Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
			return reflect.ValueOf(int64(rv.Len())), nil
		}
		return nilValue, newStringError(e, "type "+rv.Kind().String()+" does not support len operation")

	// NewExpr
	case *ast.NewExpr:
		t, err := getTypeFromString(env, e.Type)
		if err != nil {
			return nilValue, newError(e, err)
		}
		if t == nil {
			return nilValue, newErrorf(expr, "type cannot be nil for new")
		}

		return reflect.New(t), nil

	// MakeExpr
	case *ast.MakeExpr:
		t, err := getTypeFromString(env, e.Type)
		if err != nil {
			return nilValue, newError(e, err)
		}
		if t == nil {
			return nilValue, newErrorf(expr, "type cannot be nil for make")
		}

		for i := 1; i < e.Dimensions; i++ {
			t = reflect.SliceOf(t)
		}
		if e.Dimensions < 1 {
			v, err := makeValue(t)
			if err != nil {
				return nilValue, newError(e, err)
			}
			return v, nil
		}

		var alen int
		if e.LenExpr != nil {
			rv, err := invokeExpr(ctx, e.LenExpr, env)
			if err != nil {
				return nilValue, newError(e.LenExpr, err)
			}
			alen = toInt(rv)
		}

		var acap int
		if e.CapExpr != nil {
			rv, err := invokeExpr(ctx, e.CapExpr, env)
			if err != nil {
				return nilValue, newError(e.CapExpr, err)
			}
			acap = toInt(rv)
		} else {
			acap = alen
		}

		return reflect.MakeSlice(reflect.SliceOf(t), alen, acap), nil

	// MakeTypeExpr
	case *ast.MakeTypeExpr:
		rv, err := invokeExpr(ctx, e.Type, env)
		if err != nil {
			return nilValue, newError(e, err)
		}
		if !rv.IsValid() || rv.Type() == nil {
			return nilValue, newErrorf(expr, "type cannot be nil for make type")
		}

		// if e.Name has a dot in it, it should give a syntax error
		// so no needs to check err
		env.DefineReflectType(e.Name, rv.Type())

		return reflect.ValueOf(rv.Type()), nil

	// MakeChanExpr
	case *ast.MakeChanExpr:
		t, err := getTypeFromString(env, e.Type)
		if err != nil {
			return nilValue, newError(e, err)
		}
		if t == nil {
			return nilValue, newErrorf(expr, "type cannot be nil for make chan")
		}

		var size int
		if e.SizeExpr != nil {
			rv, err := invokeExpr(ctx, e.SizeExpr, env)
			if err != nil {
				return nilValue, newError(e.SizeExpr, err)
			}
			size = int(toInt64(rv))
		}

		return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), size), nil

	// ChanExpr
	case *ast.ChanExpr:
		rhs, err := invokeExpr(ctx, e.RHS, env)
		if err != nil {
			return nilValue, newError(e.RHS, err)
		}

		if e.LHS == nil {
			if rhs.Kind() == reflect.Chan {
				cases := []reflect.SelectCase{{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ctx.Done()),
					Send: zeroValue,
				}, {
					Dir:  reflect.SelectRecv,
					Chan: rhs,
					Send: zeroValue,
				}}
				chosen, rv, _ := reflect.Select(cases)
				if chosen == 0 {
					return nilValue, ErrInterrupt
				}
				return rv, nil
			}
		} else {
			lhs, err := invokeExpr(ctx, e.LHS, env)
			if err != nil {
				return nilValue, newError(e.LHS, err)
			}
			if lhs.Kind() == reflect.Chan {
				chanType := lhs.Type().Elem()
				if chanType == interfaceType || (rhs.IsValid() && rhs.Type() == chanType) {
					cases := []reflect.SelectCase{{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(ctx.Done()),
						Send: zeroValue,
					}, {
						Dir:  reflect.SelectSend,
						Chan: lhs,
						Send: rhs,
					}}
					if chosen, _, _ := reflect.Select(cases); chosen == 0 {
						return nilValue, ErrInterrupt
					}
				} else {
					rhs, err = convertReflectValueToType(rhs, chanType)
					if err != nil {
						return nilValue, newStringError(e, "cannot use type "+rhs.Type().String()+" as type "+chanType.String()+" to send to chan")
					}
					cases := []reflect.SelectCase{{
						Dir:  reflect.SelectRecv,
						Chan: reflect.ValueOf(ctx.Done()),
						Send: zeroValue,
					}, {
						Dir:  reflect.SelectSend,
						Chan: lhs,
						Send: rhs,
					}}
					if chosen, _, _ := reflect.Select(cases); chosen == 0 {
						return nilValue, ErrInterrupt
					}
				}
				return nilValue, nil
			} else if rhs.Kind() == reflect.Chan {
				cases := []reflect.SelectCase{{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ctx.Done()),
					Send: zeroValue,
				}, {
					Dir:  reflect.SelectRecv,
					Chan: rhs,
					Send: zeroValue,
				}}
				chosen, rv, ok := reflect.Select(cases)
				if chosen == 0 {
					return nilValue, ErrInterrupt
				}
				if !ok {
					return nilValue, newErrorf(expr, "failed to send to channel")
				}
				return invokeLetExpr(ctx, e.LHS, rv, env)
			}
		}

		return nilValue, newStringError(e, "invalid operation for chan")

	// FuncExpr
	case *ast.FuncExpr:
		return funcExpr(ctx, e, env)

	// AnonCallExpr
	case *ast.AnonCallExpr:
		return anonCallExpr(ctx, e, env)

	// CallExpr
	case *ast.CallExpr:
		return callExpr(ctx, e, env)

	// DeleteExpr
	case *ast.DeleteExpr:
		whatExpr, err := invokeExpr(ctx, e.WhatExpr, env)
		if err != nil {
			return nilValue, newError(e.WhatExpr, err)
		}
		var keyExpr reflect.Value
		if e.KeyExpr != nil {
			keyExpr, err = invokeExpr(ctx, e.KeyExpr, env)
			if err != nil {
				return nilValue, newError(e.KeyExpr, err)
			}
		}
		if whatExpr.Kind() == reflect.Interface && !whatExpr.IsNil() {
			whatExpr = whatExpr.Elem()
		}

		switch whatExpr.Kind() {
		case reflect.String:
			if e.KeyExpr == nil {
				return nilValue, env.Delete(whatExpr.String())
			}
			if keyExpr.Kind() == reflect.Bool && keyExpr.Bool() {
				return nilValue, env.DeleteGlobal(whatExpr.String())
			}
			return nilValue, env.Delete(whatExpr.String())

		case reflect.Map:
			if whatExpr.IsNil() {
				return nilValue, nil
			}
			if whatExpr.Type().Key() != keyExpr.Type() {
				keyExpr, err = convertReflectValueToType(keyExpr, whatExpr.Type().Key())
				if err != nil {
					return nilValue, newStringError(e, "cannot use type "+whatExpr.Type().Key().String()+" as type "+keyExpr.Type().String()+" in delete")
				}
			}
			whatExpr.SetMapIndex(keyExpr, reflect.Value{})
			return nilValue, nil
		}

		return nilValue, newStringError(e, "first argument to delete cannot be type "+whatExpr.Kind().String())

	// IncludeExpr
	case *ast.IncludeExpr:
		itemExpr, err := invokeExpr(ctx, e.ItemExpr, env)
		if err != nil {
			return nilValue, newError(e.ItemExpr, err)
		}
		listExpr, err := invokeExpr(ctx, e.ListExpr.(*ast.SliceExpr), env)
		if err != nil {
			return nilValue, newError(e.ListExpr.(*ast.SliceExpr), err)
		}

		if listExpr.Kind() != reflect.Slice && listExpr.Kind() != reflect.Array {
			return nilValue, newStringError(e, "second argument must be slice or array; but have "+listExpr.Kind().String())
		}

		for i := 0; i < listExpr.Len(); i++ {
			// todo: https://github.com/mattn/anko/issues/192
			if equal(itemExpr, listExpr.Index(i)) {
				return trueValue, nil
			}
		}
		return falseValue, nil

	default:
		return nilValue, newStringError(e, "unknown expression")
	}
}
