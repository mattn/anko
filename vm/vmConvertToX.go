package vm

import (
	"context"
	"fmt"
	"reflect"
)

// reflectValueSlicetoInterfaceSlice convert from a slice of reflect.Value to a interface slice
// returned in normal reflect.Value form
func reflectValueSlicetoInterfaceSlice(valueSlice []reflect.Value) reflect.Value {
	interfaceSlice := make([]interface{}, 0, len(valueSlice))
	for _, value := range valueSlice {
		if value.Kind() == reflect.Interface && !value.IsNil() {
			value = value.Elem()
		}
		if value.CanInterface() {
			interfaceSlice = append(interfaceSlice, value.Interface())
		} else {
			interfaceSlice = append(interfaceSlice, nil)
		}
	}
	return reflect.ValueOf(interfaceSlice)
}

// convertReflectValueToType trys to covert the reflect.Value to the reflect.Type
// if it can not, it returns the original rv and an error
func convertReflectValueToType(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	if rt == interfaceType || rv.Type() == rt {
		// if reflect.Type is interface or the types match, return the provided reflect.Value
		return rv, nil
	}
	if rt == rv.Type() {
		newRv := reflect.New(rt).Elem()
		if newRv.CanSet() {
			newRv.Set(rv)
			return newRv, nil
		} else {
			return rv, nil
		}
	}
	if rv.Type().ConvertibleTo(rt) {
		// if reflect can covert, do that conversion and return
		return rv.Convert(rt), nil
	}
	if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) &&
		(rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array) {
		// covert slice or array
		return convertSliceOrArray(rv, rt)
	}
	if rv.Kind() == rt.Kind() {
		// kind matches
		switch rv.Kind() {
		case reflect.Map:
			// convert map
			return convertMap(rv, rt)
		case reflect.Func:
			// for runVMFunction conversions, call convertVMFunctionToType
			return convertVMFunctionToType(rv, rt)
		case reflect.Ptr:
			// both rv and rt are pointers, convert what they are pointing to
			value, err := convertReflectValueToType(rv.Elem(), rt.Elem())
			if err != nil {
				return rv, err
			}
			// need to make a new value to be able to set it
			ptrV, err := makeValue(rt)
			if err != nil {
				return rv, err
			}
			// set value and return new pointer
			ptrV.Elem().Set(value)
			return ptrV, nil
		}
	}
	if rv.Type() == interfaceType {
		if rv.IsNil() {
			// return nil of correct type
			return reflect.Zero(rt), nil
		}
		// try to convert the element
		return convertReflectValueToType(rv.Elem(), rt)
	}

	if rv.Type() == stringType {
		if rt == byteType {
			aString := rv.String()
			if len(aString) < 1 {
				return reflect.Zero(rt), nil
			}
			if len(aString) > 1 {
				return rv, errInvalidTypeConversion
			}
			return reflect.ValueOf(aString[0]), nil
		}
		if rt == runeType {
			aString := rv.String()
			if len(aString) < 1 {
				return reflect.Zero(rt), nil
			}
			if len(aString) > 1 {
				return rv, errInvalidTypeConversion
			}
			return reflect.ValueOf(rune(aString[0])), nil
		}
	}

	if rt.Kind() == reflect.Ptr {
		newRv := reflect.New(rt).Elem()
		if newRv.CanSet() {
			prtRv, err := convertReflectValueToType(rv, rt.Elem())
			if err != nil {
				return rv, err
			}
			newRv.Set(prtRv.Addr())
			return newRv, nil
		}
	}

	if rt.Kind() == reflect.Struct && rv.Type() == mapIiType {
		rvIiMap := rv.Interface().(map[interface{}]interface{})
		rvSiMap := make(map[string]interface{}, len(rvIiMap))
		for k, v := range rvIiMap {
			kStr, ok := k.(string)
			if !ok {
				return rv, fmt.Errorf("invalid type conversion")
			}
			rvSiMap[kStr] = v
		}

		newRv := reflect.New(rt).Elem()
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)

			fieldV, ok := rvSiMap[field.Name]
			if !ok {
				continue
			}
			if newRv.Field(i).CanSet() {
				newFieldV, err := convertReflectValueToType(reflect.ValueOf(fieldV), newRv.Field(i).Type())
				if err != nil {
					return rv, err
				}
				newRv.Field(i).Set(newFieldV)
			}
		}
		return newRv, nil
	}

	if rt.Kind() == reflect.Map && rv.Type() == mapIiType {
		keys := rv.MapKeys()
		isSameKeyType := rt.Key() == rv.Type().Key()
		newRv := reflect.MakeMap(reflect.MapOf(rt.Key(), rt.Elem()))

		for i := 0; i < len(keys); i++ {
			key := keys[i]
			value := rv.MapIndex(key)

			if isSameKeyType {
				newValue, err := convertReflectValueToType(value, rt.Elem())
				if err != nil {
					return rv, err
				}
				newRv.SetMapIndex(key, newValue)
				continue
			}

			// only deal with rv and rt map key is of type string
			if key.Kind() == reflect.Interface {
				keyStr, ok := key.Interface().(string)
				if !ok {
					return rv, fmt.Errorf("invalid type conversion")
				}

				if rt.Key().Kind() != reflect.String {
					return rv, fmt.Errorf("invalid type conversion")
				}

				newKey := reflect.New(rt.Key()).Elem()
				if !newKey.CanSet() {
					continue
				}
				newKey.Set(reflect.ValueOf(keyStr))

				newValue, err := convertReflectValueToType(value, rt.Elem())
				if err != nil {
					return rv, err
				}
				newRv.SetMapIndex(newKey, newValue)
			}
		}
		return newRv, nil
	}

	return rv, errInvalidTypeConversion
}

// convertSliceOrArray trys to covert the reflect.Value slice or array to the slice or array reflect.Type
func convertSliceOrArray(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	rtElemType := rt.Elem()

	// try to covert elements to new slice/array
	var value reflect.Value
	if rt.Kind() == reflect.Slice {
		// make slice
		value = reflect.MakeSlice(rt, rv.Len(), rv.Len())
	} else {
		// make array
		value = reflect.New(rt).Elem()
	}

	var err error
	var v reflect.Value
	for i := 0; i < rv.Len(); i++ {
		v, err = convertReflectValueToType(rv.Index(i), rtElemType)
		if err != nil {
			return rv, err
		}
		value.Index(i).Set(v)
	}

	// return new converted slice or array
	return value, nil
}

// convertVMFunctionToType is for translating a runVMFunction into the correct type
// so it can be passed to a Go function argument with the correct static types
// it creates a translate function runVMConvertFunction
func convertVMFunctionToType(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	// only translates runVMFunction type
	if !checkIfRunVMFunction(rv.Type()) {
		return rv, errInvalidTypeConversion
	}

	// create runVMConvertFunction to match reflect.Type
	// this function is being called by the Go function
	runVMConvertFunction := func(in []reflect.Value) []reflect.Value {
		// note: this function is being called by another reflect Call
		// only way to pass along any errors is by panic

		// make the reflect.Value slice of each of the VM reflect.Value
		args := make([]reflect.Value, 0, rt.NumIn()+1)
		// for runVMFunction first arg is always context
		// TOFIX: use normal context
		args = append(args, reflect.ValueOf(context.Background()))
		for i := 0; i < rt.NumIn(); i++ {
			// have to do the double reflect.ValueOf that runVMFunction expects
			args = append(args, reflect.ValueOf(in[i]))
		}

		// Call runVMFunction
		rvs := rv.Call(args)

		// call processCallReturnValues to process runVMFunction return values
		// returns normal VM reflect.Value form
		rv, err := processCallReturnValues(rvs, true, false)
		if err != nil {
			panic(err)
		}

		if rt.NumOut() < 1 {
			// Go function does not want any return values, so give it none
			return []reflect.Value{}
		}
		if rt.NumOut() < 2 {
			// Go function wants one return value
			// will try to covert to reflect.Value correct type and return
			rv, err = convertReflectValueToType(rv, rt.Out(0))
			if err != nil {
				panic("function wants return type " + rt.Out(0).String() + " but received type " + rv.Type().String())
			}
			return []reflect.Value{rv}
		}

		// Go function wants more than one return value
		// make sure we have a slice/array with enought values

		if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
			panic(fmt.Sprintf("function wants %v return values but received %v", rt.NumOut(), rv.Kind().String()))
		}
		if rv.Len() < rt.NumOut() {
			panic(fmt.Sprintf("function wants %v return values but received %v values", rt.NumOut(), rv.Len()))
		}

		// try to covert each value in slice to wanted type and put into a reflect.Value slice
		rvs = make([]reflect.Value, rt.NumOut())
		for i := 0; i < rv.Len(); i++ {
			rvs[i], err = convertReflectValueToType(rv.Index(i), rt.Out(i))
			if err != nil {
				panic("function wants return type " + rt.Out(i).String() + " but received type " + rvs[i].Type().String())
			}
		}

		// return created reflect.Value slice
		return rvs
	}

	// make the reflect.Value function that calls runVMConvertFunction
	return reflect.MakeFunc(rt, runVMConvertFunction), nil
}
