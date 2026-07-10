package vm

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/mattn/anko/ast"
	"github.com/mattn/anko/env"
)

// Options provides options to run VM with
type Options struct {
	Debug bool // run in Debug mode
}

type (
	// Error is a VM run error.
	Error struct {
		Message string
		Pos     ast.Position
	}

	// runInfo provides run incoming and outgoing information
	runInfoStruct struct {
		// incoming
		ctx      context.Context
		env      *env.Env
		options  *Options
		stmt     ast.Stmt
		expr     ast.Expr
		operator ast.Operator

		// deferred function calls of the current function invocation,
		// run in LIFO order when the function returns
		defers []capturedFunc

		// outgoing
		rv  reflect.Value
		err error
	}

	// capturedFunc holds a deferred function with its arguments already
	// evaluated at the defer statement
	capturedFunc struct {
		fn        reflect.Value
		args      []reflect.Value
		callSlice bool
	}
)

var (
	nilType            = reflect.TypeOf(nil)
	stringType         = reflect.TypeOf("a")
	byteType           = reflect.TypeOf(byte('a'))
	runeType           = reflect.TypeOf('a')
	interfaceType      = reflect.ValueOf([]interface{}{int64(1)}).Index(0).Type()
	interfaceSliceType = reflect.TypeOf([]interface{}{})
	reflectValueType   = reflect.TypeOf(reflect.Value{})
	errorType          = reflect.ValueOf([]error{nil}).Index(0).Type()
	vmErrorType        = reflect.TypeOf(&Error{})
	contextType        = reflect.TypeOf((*context.Context)(nil)).Elem()

	nilValue                  = reflect.New(reflect.TypeOf((*interface{})(nil)).Elem()).Elem()
	trueValue                 = reflect.ValueOf(true)
	falseValue                = reflect.ValueOf(false)
	int64Type                 = reflect.TypeOf(int64(0))
	float64Type               = reflect.TypeOf(float64(0))
	zeroValue                 = reflect.Value{}
	errorNilValue             = reflect.New(errorType).Elem()
	reflectValueNilValue      = reflect.ValueOf(nilValue)
	reflectValueErrorNilValue = reflect.ValueOf(reflect.New(errorType).Elem())

	errInvalidTypeConversion = fmt.Errorf("invalid type conversion")

	// ErrBreak when there is an unexpected break statement
	ErrBreak = errors.New("unexpected break statement")
	// ErrContinue when there is an unexpected continue statement
	ErrContinue = errors.New("unexpected continue statement")
	// ErrReturn when there is an unexpected return statement
	ErrReturn = errors.New("unexpected return statement")
	// ErrInterrupt when execution has been interrupted
	ErrInterrupt = errors.New("execution interrupted")
)

// Error returns the VM error message.
func (e *Error) Error() string {
	return e.Message
}

// newError makes VM error from error
func newError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if pos == nil {
		return &Error{Message: err.Error(), Pos: ast.Position{Line: 1, Column: 1}}
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

// newStringError makes VM error from string
func newStringError(pos ast.Pos, err string) error {
	if err == "" {
		return nil
	}
	if pos == nil {
		return &Error{Message: err, Pos: ast.Position{Line: 1, Column: 1}}
	}
	return &Error{Message: err, Pos: pos.Position()}
}

// recoverFunc generic recover function
func recoverFunc(runInfo *runInfoStruct) {
	recoverInterface := recover()
	if recoverInterface == nil {
		return
	}
	switch value := recoverInterface.(type) {
	case *Error:
		runInfo.err = value
	case error:
		runInfo.err = value
	default:
		runInfo.err = fmt.Errorf("%v", recoverInterface)
	}
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		// from reflect IsNil:
		// Note that IsNil is not always equivalent to a regular comparison with nil in Go.
		// For example, if v was created by calling ValueOf with an uninitialized interface variable i,
		// i==nil will be true but v.IsNil will panic as v will be the zero Value.
		return v.IsNil()
	default:
		return false
	}
}

const int64CacheMin = -1
const int64CacheMax = 4095

var int64Cache [int64CacheMax - int64CacheMin + 1]reflect.Value

func init() {
	for i := int64CacheMin; i <= int64CacheMax; i++ {
		int64Cache[i-int64CacheMin] = reflect.ValueOf(int64(i))
	}
}

// int64Value returns a reflect.Value for the given int64, using a cache for small values.
func int64Value(v int64) reflect.Value {
	if v >= int64CacheMin && v <= int64CacheMax {
		return int64Cache[v-int64CacheMin]
	}
	return reflect.ValueOf(v)
}

// float64Value returns a reflect.Value for the given float64.
func float64Value(v float64) reflect.Value {
	return reflect.ValueOf(v)
}

// numToString converts a numeric reflect.Value to its string representation.
// For float32, it uses 32-bit precision to match fmt.Sprintf("%v") behavior.
func numToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'G', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'G', -1, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	default:
		return ""
	}
}

// isIntKind returns true if the value is a signed integer kind.
func isIntKind(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isNum(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// equal returns true when lhsV and rhsV is same value.
func equal(lhsV, rhsV reflect.Value) bool {
	lhsIsNil, rhsIsNil := isNil(lhsV), isNil(rhsV)
	if lhsIsNil && rhsIsNil {
		return true
	}
	if (!lhsIsNil && rhsIsNil) || (lhsIsNil && !rhsIsNil) {
		return false
	}
	if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
		rhsV = rhsV.Elem()
	}

	// Compare a string and a number.
	// This will attempt to convert the string to a number,
	// while leaving the other side alone. Code further
	// down takes care of converting ints and floats as needed.
	if isNum(lhsV) && rhsV.Kind() == reflect.String {
		rhsF, err := tryToFloat64(rhsV)
		if err != nil {
			// Couldn't convert RHS to a float, they can't be compared.
			return false
		}
		rhsV = reflect.ValueOf(rhsF)
	} else if lhsV.Kind() == reflect.String && isNum(rhsV) {
		// If the LHS is a string formatted as an int, try that before trying float
		lhsI, err := tryToInt64(lhsV)
		if err != nil {
			// if LHS is a float, e.g. "1.2", we need to set lhsV to a float64
			lhsF, err := tryToFloat64(lhsV)
			if err != nil {
				return false
			}
			lhsV = reflect.ValueOf(lhsF)
		} else {
			lhsV = reflect.ValueOf(lhsI)
		}
	}

	if isNum(lhsV) && isNum(rhsV) {
		lhsKind := lhsV.Kind()
		rhsKind := rhsV.Kind()
		lhsIsFloat := lhsKind == reflect.Float32 || lhsKind == reflect.Float64
		rhsIsFloat := rhsKind == reflect.Float32 || rhsKind == reflect.Float64
		if !lhsIsFloat && !rhsIsFloat {
			return toInt64(lhsV) == toInt64(rhsV)
		}
		// when both are same kind, direct comparison is safe
		if lhsKind == rhsKind {
			return toFloat64(lhsV) == toFloat64(rhsV)
		}
		// mixed types: use string representation for compatibility
		// (e.g. float32(1.1) should equal float64(1.1))
		return numToString(lhsV) == numToString(rhsV)
	}

	// Try to compare bools to strings and numbers
	if lhsV.Kind() == reflect.Bool || rhsV.Kind() == reflect.Bool {
		lhsB, err := tryToBool(lhsV)
		if err != nil {
			return false
		}
		rhsB, err := tryToBool(rhsV)
		if err != nil {
			return false
		}
		return lhsB == rhsB
	}

	return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
}

// isHashable returns true if the value can be used as a map key without
// panicking. A nil interface is a valid map key even though
// reflect.Value.Comparable reports false for it.
func isHashable(v reflect.Value) bool {
	if v.Kind() == reflect.Interface && v.IsNil() {
		return true
	}
	return v.Comparable()
}

// hashableTypeString returns the type string of the value for error messages,
// unwrapping interfaces to show the dynamic type.
func hashableTypeString(v reflect.Value) string {
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}
	return v.Type().String()
}

func getMapIndex(key reflect.Value, aMap reflect.Value) reflect.Value {
	if aMap.IsNil() {
		return nilValue
	}

	var err error
	key, err = convertReflectValueToType(key, aMap.Type().Key())
	if err != nil {
		return nilValue
	}
	if !isHashable(key) {
		// an unhashable key can never be in a map
		return nilValue
	}

	// From reflect MapIndex:
	// It returns the zero Value if key is not found in the map or if v represents a nil map.
	value := aMap.MapIndex(key)
	if !value.IsValid() {
		return nilValue
	}

	if aMap.Type().Elem() == interfaceType && !value.IsNil() {
		value = reflect.ValueOf(value.Interface())
	}

	return value
}

// appendSlice appends rhs to lhs
// function assumes lhsV and rhsV are slice or array
func appendSlice(expr ast.Expr, lhsV reflect.Value, rhsV reflect.Value) (reflect.Value, error) {
	lhsT := lhsV.Type().Elem()
	rhsT := rhsV.Type().Elem()

	if lhsT == rhsT {
		return reflect.AppendSlice(lhsV, rhsV), nil
	}

	if rhsT.ConvertibleTo(lhsT) {
		for i := 0; i < rhsV.Len(); i++ {
			lhsV = reflect.Append(lhsV, rhsV.Index(i).Convert(lhsT))
		}
		return lhsV, nil
	}

	leftHasSubArray := lhsT.Kind() == reflect.Slice || lhsT.Kind() == reflect.Array
	rightHasSubArray := rhsT.Kind() == reflect.Slice || rhsT.Kind() == reflect.Array

	if leftHasSubArray != rightHasSubArray && lhsT != interfaceType && rhsT != interfaceType {
		return nilValue, newStringError(expr, "invalid type conversion")
	}

	if !leftHasSubArray && !rightHasSubArray {
		for i := 0; i < rhsV.Len(); i++ {
			value := rhsV.Index(i)
			if rhsT == interfaceType {
				value = value.Elem()
			}
			if lhsT == value.Type() {
				lhsV = reflect.Append(lhsV, value)
			} else if value.Type().ConvertibleTo(lhsT) {
				lhsV = reflect.Append(lhsV, value.Convert(lhsT))
			} else {
				return nilValue, newStringError(expr, "invalid type conversion")
			}
		}
		return lhsV, nil
	}

	if (leftHasSubArray || lhsT == interfaceType) && (rightHasSubArray || rhsT == interfaceType) {
		for i := 0; i < rhsV.Len(); i++ {
			value := rhsV.Index(i)
			if rhsT == interfaceType {
				value = value.Elem()
				if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
					return nilValue, newStringError(expr, "invalid type conversion")
				}
			}
			newSlice, err := appendSlice(expr, reflect.MakeSlice(lhsT, 0, value.Len()), value)
			if err != nil {
				return nilValue, err
			}
			lhsV = reflect.Append(lhsV, newSlice)
		}
		return lhsV, nil
	}

	return nilValue, newStringError(expr, "invalid type conversion")
}

func makeType(runInfo *runInfoStruct, typeStruct *ast.TypeStruct) reflect.Type {
	switch typeStruct.Kind {
	case ast.TypeDefault:
		return getTypeFromEnv(runInfo, typeStruct)
	case ast.TypePtr:
		var t reflect.Type
		if typeStruct.SubType != nil {
			t = makeType(runInfo, typeStruct.SubType)
		} else {
			t = getTypeFromEnv(runInfo, typeStruct)
		}
		if runInfo.err != nil {
			return nil
		}
		if t == nil {
			return nil
		}
		return reflect.PtrTo(t)
	case ast.TypeSlice:
		var t reflect.Type
		if typeStruct.SubType != nil {
			t = makeType(runInfo, typeStruct.SubType)
		} else {
			t = getTypeFromEnv(runInfo, typeStruct)
		}
		if runInfo.err != nil {
			return nil
		}
		if t == nil {
			return nil
		}
		for i := 1; i < typeStruct.Dimensions; i++ {
			t = reflect.SliceOf(t)
		}
		return reflect.SliceOf(t)
	case ast.TypeMap:
		key := makeType(runInfo, typeStruct.Key)
		if runInfo.err != nil {
			return nil
		}
		if key == nil {
			return nil
		}
		t := makeType(runInfo, typeStruct.SubType)
		if runInfo.err != nil {
			return nil
		}
		if t == nil {
			return nil
		}
		if !runInfo.options.Debug {
			// captures panic
			defer recoverFunc(runInfo)
		}
		t = reflect.MapOf(key, t)
		return t
	case ast.TypeChan:
		var t reflect.Type
		if typeStruct.SubType != nil {
			t = makeType(runInfo, typeStruct.SubType)
		} else {
			t = getTypeFromEnv(runInfo, typeStruct)
		}
		if runInfo.err != nil {
			return nil
		}
		if t == nil {
			return nil
		}
		return reflect.ChanOf(reflect.BothDir, t)
	case ast.TypeStructType:
		var t reflect.Type
		fields := make([]reflect.StructField, 0, len(typeStruct.StructNames))
		for i := 0; i < len(typeStruct.StructNames); i++ {
			t = makeType(runInfo, typeStruct.StructTypes[i])
			if runInfo.err != nil {
				return nil
			}
			if t == nil {
				return nil
			}
			fields = append(fields, reflect.StructField{Name: typeStruct.StructNames[i], Type: t})
		}
		if !runInfo.options.Debug {
			// captures panic
			defer recoverFunc(runInfo)
		}
		t = reflect.StructOf(fields)
		return t
	default:
		runInfo.err = fmt.Errorf("unknown kind")
		return nil
	}
}

func getTypeFromEnv(runInfo *runInfoStruct, typeStruct *ast.TypeStruct) reflect.Type {
	var e *env.Env
	e, runInfo.err = runInfo.env.GetEnvFromPath(typeStruct.Env)
	if runInfo.err != nil {
		return nil
	}

	var t reflect.Type
	t, runInfo.err = e.Type(typeStruct.Name)
	return t
}

func makeValue(t reflect.Type) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Chan:
		return reflect.MakeChan(t, 0), nil
	case reflect.Func:
		return reflect.MakeFunc(t, nil), nil
	case reflect.Map:
		// note creating slice as work around to create map
		// just doing MakeMap can give incorrect type for defined types
		value := reflect.MakeSlice(reflect.SliceOf(t), 0, 1)
		value = reflect.Append(value, reflect.MakeMap(reflect.MapOf(t.Key(), t.Elem())))
		return value.Index(0), nil
	case reflect.Ptr:
		ptrV := reflect.New(t.Elem())
		v, err := makeValue(t.Elem())
		if err != nil {
			return nilValue, err
		}
		ptrV.Elem().Set(v)
		return ptrV, nil
	case reflect.Slice:
		return reflect.MakeSlice(t, 0, 0), nil
	}
	return reflect.New(t).Elem(), nil
}

// precedenceOfKinds returns the greater of two kinds
// string > float > int
func precedenceOfKinds(kind1 reflect.Kind, kind2 reflect.Kind) reflect.Kind {
	if kind1 == kind2 {
		return kind1
	}
	switch kind1 {
	case reflect.String:
		return kind1
	case reflect.Float64, reflect.Float32:
		switch kind2 {
		case reflect.String:
			return kind2
		}
		return kind1
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch kind2 {
		case reflect.String, reflect.Float64, reflect.Float32:
			return kind2
		}
	}
	return kind1
}
