package vm

import (
	"fmt"
	"reflect"
)

// reflectValueSlicetoInterfaceSlice convert from a slice of reflect.Value to a interface slice
// returned in normal reflect.Value form
func reflectValueSlicetoInterfaceSlice(valueSlice []reflect.Value) reflect.Value {
	interfaceSlice := make([]interface{}, 0, len(valueSlice))
	for _, value := range valueSlice {
		if !value.IsValid() {
			interfaceSlice = append(interfaceSlice, nil)
			continue
		}
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
// if it can not, it returns the orginal rv and an error
func convertReflectValueToType(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	if !rv.IsValid() {
		// if not valid return a valid reflect.Value of the reflect.Type
		return makeValue(rt)
	}
	if rt == interfaceType || rv.Type() == rt {
		// if reflect.Type is interface or the types match, return the provided reflect.Value
		return rv, nil
	}
	if rv.Type().ConvertibleTo(rt) {
		// if reflect can covert, do that convertion and return
		return rv.Convert(rt), nil
	}
	if rv.Kind() == reflect.Func && rt.Kind() == reflect.Func {
		// for function convertions, call convertVMFunctionToType
		return convertVMFunctionToType(rv, rt)
	}
	if rv.Kind() == reflect.Ptr && rt.Kind() == reflect.Ptr {
		// both rv and rt are pointers
		// convert what they are pointing to
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
	if rv.Type() == interfaceType {
		// reflect.Value is an interface, so try to convert the element
		return convertReflectValueToType(rv.Elem(), rt)
	}

	// TOFIX: need to handle the case where rv or rt are a pointer but not both

	return rv, fmt.Errorf("invalid type conversion")
}

func convertVMFunctionToType(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	if !checkIfRunVMFunction(rv.Type()) {
		return rv, fmt.Errorf("invalid type conversion")
	}

	runVMConvertFunction := func(in []reflect.Value) []reflect.Value {
		args := make([]reflect.Value, 0, rt.NumIn())
		for i := 0; i < rt.NumIn(); i++ {
			args = append(args, reflect.ValueOf(in[i]))
		}

		rvs := rv.Call(args)

		rv, err := processCallReturnValues(rvs, true, false)
		if err != nil {
			panic("function run error: " + err.Error())
		}

		if rt.NumOut() < 1 {
			return []reflect.Value{}
		}
		if rt.NumOut() < 2 {
			rv, err = convertReflectValueToType(rv, rt.Out(0))
			if err != nil {
				panic("function wants return type " + rt.Out(0).String() + " but received type " + rv.Type().String())
			}
			return []reflect.Value{rv}
		}

		if rv.Type() != reflectValueSliceType {
			panic("bad function return type: " + rv.Type().String())
		}

		outValues := rv.Interface().([]reflect.Value)
		if len(outValues) < rt.NumOut() {
			panic(fmt.Sprintf("function wants %v return values but received %v values", rt.NumOut(), len(outValues)))
		}

		rvs = make([]reflect.Value, 0, rt.NumOut())
		for i := 0; i < rt.NumOut(); i++ {
			rv, err = convertReflectValueToType(outValues[i], rt.Out(i))
			if err != nil {
				panic("function wants return type " + rt.Out(i).String() + " but received type " + rv.Type().String())
			}
			rvs = append(rvs, rv)
		}
		return rvs
	}

	return reflect.MakeFunc(rt, runVMConvertFunction), nil
}
