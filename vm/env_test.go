package vm

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type externalResolver struct {
	entries map[string]interface{}
}

func NewExternalResolver() *externalResolver {
	return &externalResolver{entries: make(map[string]interface{})}
}

func (er *externalResolver) Set(name string, value interface{}) {
	er.entries[name] = value
}

func (er *externalResolver) Get(name string) (reflect.Value, error) {
	if v, ok := er.entries[name]; ok {
		if v == nil {
			return NilValue, nil
		}
		return reflect.ValueOf(v), nil
	}
	return NilValue, fmt.Errorf("Undefined symbol '%s'", name)
}

func (er *externalResolver) Type(name string) (reflect.Type, error) {
	if v, ok := er.entries[name]; ok {
		if v == nil {
			return NilType, nil
		}
		return reflect.TypeOf(v), nil
	}
	return NilType, fmt.Errorf("Undefined symbol '%s'", name)
}

func TestExternal(t *testing.T) {
	er := NewExternalResolver()

	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface, defineError: nil, getError: nil},
		{testInfo: "string", varName: "b", varDefineValue: "a", varGetValue: "a", varKind: reflect.String, defineError: nil, getError: nil},
		{testInfo: "int16", varName: "c", varDefineValue: int16(1), varGetValue: int16(1), varKind: reflect.Int16, defineError: nil, getError: nil},
		{testInfo: "int32", varName: "d", varDefineValue: int32(1), varGetValue: int32(1), varKind: reflect.Int32, defineError: nil, getError: nil},
		{testInfo: "int64", varName: "e", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64, defineError: nil, getError: nil},
		{testInfo: "uint32", varName: "f", varDefineValue: uint32(1), varGetValue: uint32(1), varKind: reflect.Uint32, defineError: nil, getError: nil},
		{testInfo: "uint64", varName: "g", varDefineValue: uint64(1), varGetValue: uint64(1), varKind: reflect.Uint64, defineError: nil, getError: nil},
		{testInfo: "float32", varName: "h", varDefineValue: float32(1), varGetValue: float32(1), varKind: reflect.Float32, defineError: nil, getError: nil},
		{testInfo: "float64", varName: "i", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64, defineError: nil, getError: nil},
		{testInfo: "bool", varName: "j", varDefineValue: true, varGetValue: true, varKind: reflect.Bool, defineError: nil, getError: nil},
	}

	// First set all values
	for _, test := range tests {
		er.Set(test.varName, test.varDefineValue)
	}

	// Now create an Env and set the external
	env := NewEnv()
	env.SetExternal(er)

	// Now make sure they were stored correctly and can be retrieved from the Env
	// The Env doesn't know about any of the variables in the test array, but the external resolver does.
	for _, test := range tests {
		val, err := env.Get(test.varName)
		if err != nil {
			t.Errorf("Execute error - unable to retrieve variable %v from external resolver: %v", test.varName, err)
		}
		if val != test.varGetValue {
			t.Errorf("Execute error - received unexpected value %v from external resolver, expected %v", val, test.varGetValue)
		}

		varType, err := env.Type(test.varName)
		if err != nil {
			t.Errorf("Execute error - unable to retrieve type of variable %v from external resolver: %v", test.varName, err)
		}
		if varType == nil || test.varDefineValue == nil {
			if varType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("Execute error - type received: %v expected: %v", varType, reflect.TypeOf(test.varDefineValue))
			}
		} else if varType.Kind() != test.varKind {
			t.Errorf("Execute error - received unexpected type %v from external resolver, expected %v", varType, test.varKind)
		}
	}
}

func TestExecuteError(t *testing.T) {
	env := NewEnv()
	script := "a]]"
	_, err := env.Execute(script)
	if err == nil {
		t.Errorf("Execute error - received: %v expected: %v", err, fmt.Errorf("syntax error"))
	} else if err.Error() != "syntax error" {
		t.Errorf("Execute error - received: %v expected: %v", err, fmt.Errorf("syntax error"))
	}
}

func TestSetError(t *testing.T) {
	envParent := NewEnv()
	envChild := envParent.NewEnv()
	err := envChild.Set("a", "a")
	if err == nil {
		t.Errorf("Set error - received: %v expected: %v", err, fmt.Errorf("Unknown symbol 'a'"))
	} else if err.Error() != "Unknown symbol 'a'" {
		t.Errorf("Set error - received: %v expected: %v", err, fmt.Errorf("Unknown symbol 'a'"))
	}
}

func TestAddrError(t *testing.T) {
	envParent := NewEnv()
	envChild := envParent.NewEnv()
	_, err := envChild.Addr("a")
	if err == nil {
		t.Errorf("Addr error - received: %v expected: %v", err, fmt.Errorf("Undefined symbol 'a'"))
	} else if err.Error() != "Undefined symbol 'a'" {
		t.Errorf("Addr error - received: %v expected: %v", err, fmt.Errorf("Undefined symbol 'a'"))
	}
}

func TestDefineAndGet(t *testing.T) {
	var err error
	var value interface{}
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface, defineError: nil, getError: nil},
		{testInfo: "string", varName: "a", varDefineValue: "a", varGetValue: "a", varKind: reflect.String, defineError: nil, getError: nil},
		{testInfo: " int16", varName: "a", varDefineValue: int16(1), varGetValue: int16(1), varKind: reflect.Int16, defineError: nil, getError: nil},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), varGetValue: int32(1), varKind: reflect.Int32, defineError: nil, getError: nil},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64, defineError: nil, getError: nil},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), varGetValue: uint32(1), varKind: reflect.Uint32, defineError: nil, getError: nil},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), varGetValue: uint64(1), varKind: reflect.Uint64, defineError: nil, getError: nil},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), varGetValue: float32(1), varKind: reflect.Float32, defineError: nil, getError: nil},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64, defineError: nil, getError: nil},
		{testInfo: "bool", varName: "a", varDefineValue: true, varGetValue: true, varKind: reflect.Bool, defineError: nil, getError: nil},

		{testInfo: "string with dot", varName: "a.a", varDefineValue: "a", varGetValue: nil, varKind: reflect.Interface, defineError: fmt.Errorf("Unknown symbol 'a.a'"), getError: fmt.Errorf("Undefined symbol 'a.a'")},
		{testInfo: "string with quotes", varName: "a", varDefineValue: "\"a\"", varGetValue: "\"a\"", varKind: reflect.String, defineError: nil, getError: nil},
	}

	// DefineAndGet
	for _, test := range tests {
		env := NewEnv()

		err = env.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		env.Destroy()
	}

	// DefineAndGet NewPackage
	for _, test := range tests {
		env := NewPackage("package")

		err = env.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet NewPackage %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet NewPackage %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet NewPackage %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet NewPackage %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet NewPackage %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		if env.GetName() != "package" {
			t.Errorf("DefineAndGet NewPackage %v - GetName check - received %#v expected: %#v", test.testInfo, env.GetName(), "package")
		}
		if env.String() != "package" {
			t.Errorf("DefineAndGet NewPackage %v - String check - received %#v expected: %#v", test.testInfo, env.GetName(), "package")
		}

		env.Destroy()
	}

	// DefineAndGet NewEnv
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet NewEnv %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet NewEnv %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet NewEnv %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet NewEnv %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet NewEnv %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		envChild.Destroy()
	}

	// DefineAndGet DefineGlobal
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envChild.DefineGlobal(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet DefineGlobal %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet DefineGlobal %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envParent.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet DefineGlobal %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet DefineGlobal %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet DefineGlobal %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		envParent.Destroy()
	}

}

func TestDefineModify(t *testing.T) {
	var err error
	var value interface{}
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface, defineError: nil, getError: nil},
		{testInfo: "string", varName: "a", varDefineValue: "a", varGetValue: "a", varKind: reflect.String, defineError: nil, getError: nil},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64, defineError: nil, getError: nil},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64, defineError: nil, getError: nil},
		{testInfo: "bool", varName: "a", varDefineValue: true, varGetValue: true, varKind: reflect.Bool, defineError: nil, getError: nil},
	}
	changeTests := []struct {
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface, defineError: nil, getError: nil},
		{varDefineValue: "a", varGetValue: "a", varKind: reflect.String, defineError: nil, getError: nil},
		{varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64, defineError: nil, getError: nil},
		{varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64, defineError: nil, getError: nil},
		{varDefineValue: true, varGetValue: true, varKind: reflect.Bool, defineError: nil, getError: nil},
	}

	// DefineModify
	for _, test := range tests {
		env := NewEnv()

		err = env.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		// DefineModify changeTest
		for _, changeTest := range changeTests {
			err = env.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = env.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify changeTest  %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("DefineModify changeTest  %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify changeTest  %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}

		env.Destroy()
	}

	// DefineModify envParent
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify envParent %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify envParent %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify envParent  %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify envParent  %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify envParent  %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		for _, changeTest := range changeTests {
			err = envParent.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify envParent changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify envParent changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = envChild.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify envParent changeTest %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("ChanDefineModify envParent changeTestgeTest %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify envParent changeTest %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}

		envChild.Destroy()
	}

	// DefineModify envChild
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify envChild %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify envChild %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify envChild  %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify envChild  %v - Get error - received: %v expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify envChild  %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		for _, changeTest := range changeTests {
			err = envChild.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify envChild changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify envChild changeTest %v - Set error - received: %v expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = envChild.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify envChild changeTest %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("ChanDefineModify envChild changeTestgeTest %v - Get error - received: %v expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify envChild changeTest %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}

		envChild.Destroy()
	}
}

func TestDefineType(t *testing.T) {
	var err error
	var valueType reflect.Type
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		typeError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, defineError: nil, typeError: nil},
		{testInfo: "string", varName: "a", varDefineValue: "a", defineError: nil, typeError: nil},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1), defineError: nil, typeError: nil},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), defineError: nil, typeError: nil},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), defineError: nil, typeError: nil},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), defineError: nil, typeError: nil},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), defineError: nil, typeError: nil},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), defineError: nil, typeError: nil},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), defineError: nil, typeError: nil},
		{testInfo: "bool", varName: "a", varDefineValue: true, defineError: nil, typeError: nil},
		{testInfo: "string with dot", varName: "a.a", varDefineValue: nil, defineError: fmt.Errorf("Unknown symbol 'a.a'"), typeError: fmt.Errorf("Undefined type 'a.a'")},
	}

	// DefineType
	for _, test := range tests {
		env := NewEnv()

		err = env.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = env.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}

		env.Destroy()
	}

	// DefineType NewEnv
	for _, test := range tests {
		envParent := NewEnv()
		envParent.SetName("parent")
		envChild := envParent.NewEnv()
		envChild.SetName("child")

		err = envParent.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType NewEnv %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType NewEnv %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType NewEnv %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType NewEnv %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType NewEnv %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType NewEnv %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
		envChild.Destroy()
	}

	// DefineType NewPackage
	for _, test := range tests {
		envParent := NewEnv()
		envParent.SetName("parent")
		envChild := envParent.NewPackage("child")

		err = envParent.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType NewPackage  %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType NewPackage %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType NewPackage %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType NewPackage %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType NewPackage %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType NewPackage %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}

		envChild.Destroy()
	}

	// DefineType NewModule
	for _, test := range tests {
		envParent := NewEnv()
		envParent.SetName("parent")
		envChild := envParent.NewModule("child")

		err = envParent.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType NewModule %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType NewModule %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType NewModule %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType NewModule %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType NewModule %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType NewModule %v - Type check - received: %v expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}

		envChild.Destroy()
	}
}

func TestDefineTypeFail(t *testing.T) {
	var err error
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		typeError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "string", varName: "a", varDefineValue: "a", defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
		{testInfo: "bool", varName: "a", varDefineValue: true, defineError: nil, typeError: fmt.Errorf("Undefined type 'a'")},
	}

	// DefineTypeFail
	for _, test := range tests {
		envParent := NewEnv()
		envParent.SetName("parent")
		envChild := envParent.NewEnv()
		envChild.SetName("child")

		err = envChild.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestDefineTypeFail %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestDefineTypeFail %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		_, err = envParent.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("TestDefineTypeFail %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("TestDefineTypeFail %v - Type error - received: %v expected: %v", test.testInfo, err, test.typeError)
		}

		envChild.Destroy()
	}
}

func TestAddr(t *testing.T) {
	var err error
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		addrError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, defineError: nil, addrError: nil},
		{testInfo: "string", varName: "a", varDefineValue: "a", defineError: nil, addrError: fmt.Errorf("Unaddressable")},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), defineError: nil, addrError: fmt.Errorf("Unaddressable")},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), defineError: nil, addrError: fmt.Errorf("Unaddressable")},
		{testInfo: "bool", varName: "a", varDefineValue: true, defineError: nil, addrError: fmt.Errorf("Unaddressable")},
	}

	// TestAddr
	for _, test := range tests {
		envParent := NewEnv()
		envParent.SetName("parent")
		envChild := envParent.NewEnv()
		envChild.SetName("child")

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestAddr %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestAddr %v - Define error - received: %v expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		_, err = envChild.Addr(test.varName)
		if err != nil && test.addrError != nil {
			if err.Error() != test.addrError.Error() {
				t.Errorf("TestAddr %v - Addr error - received: %v expected: %v", test.testInfo, err, test.addrError)
				continue
			}
		} else if err != test.addrError {
			t.Errorf("TestAddr %v - Addr error - received: %v expected: %v", test.testInfo, err, test.addrError)
			continue
		}

		envChild.Destroy()
	}
}

func TestRaceCreateSameVariable(t *testing.T) {
	// Test creating same variable in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Define("a", i)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			_, err = env.Get("a")
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	_, err := env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func TestRaceCreateDifferentVariables(t *testing.T) {
	// Test creating different variables in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Define(fmt.Sprint(i), i)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			_, err = env.Get(fmt.Sprint(i))
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	for i := 0; i < 100; i++ {
		_, err := env.Get(fmt.Sprint(i))
		if err != nil {
			t.Errorf("Get error: %v", err)
		}
	}
}

func TestRaceReadDifferentVariables(t *testing.T) {
	// Test reading different variables in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		err := env.Define(fmt.Sprint(i), i)
		if err != nil {
			t.Errorf("Define error: %v", err)
		}
		_, err = env.Get(fmt.Sprint(i))
		if err != nil {
			t.Errorf("Get error: %v", err)
		}
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			_, err := env.Get(fmt.Sprint(i))
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()
}

func TestRaceSetSameVariable(t *testing.T) {
	// Test setting same variable in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	err := env.Define("a", 0)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Set("a", i)
			if err != nil {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func TestRaceSetSameVariableNewEnv(t *testing.T) {
	// Test setting same variable in parallel with NewEnv

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	err := env.Define("a", 0)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			env = env.NewEnv().NewEnv()
			err := env.Set("a", i)
			if err != nil {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}
}

func TestRaceDefineAndSetSameVariable(t *testing.T) {
	// Test defining and setting same variable in parallel
	for i := 0; i < 100; i++ {
		raceDefineAndSetSameVariable(t)
	}
}

func raceDefineAndSetSameVariable(t *testing.T) {
	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	envParent := NewEnv()
	envChild := envParent.NewEnv()

	for i := 0; i < 2; i++ {
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envParent.Set("a", 1)
			if err != nil && err.Error() != "Unknown symbol 'a'" {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envParent.Define("a", 2)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envChild.Set("a", 3)
			if err != nil && err.Error() != "Unknown symbol 'a'" {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envChild.Define("a", 4)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			waitGroup.Done()
		}()
	}

	close(waitChan)
	waitGroup.Wait()

	_, err := envParent.Get("a") // value of a could be 1, 2, or 3
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	_, err = envChild.Get("a") // value of a could be 3 or 4
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func BenchmarkDefine(b *testing.B) {
	var err error
	env := NewEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := env.Define("a", 1)
		if err != nil {
			b.Errorf("Set error: %v", err)
		}
	}
	b.StopTimer()
	_, err = env.Get("a")
	if err != nil {
		b.Errorf("Get error: %v", err)
	}
}

func BenchmarkSet(b *testing.B) {
	env := NewEnv()
	err := env.Define("a", 1)
	if err != nil {
		b.Errorf("Define error: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = env.Set("a", 1)
		if err != nil {
			b.Errorf("Set error: %v", err)
		}
	}
	b.StopTimer()
	_, err = env.Get("a")
	if err != nil {
		b.Errorf("Get error: %v", err)
	}
}
