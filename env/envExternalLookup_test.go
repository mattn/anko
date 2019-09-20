package env

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type TestExternalLookup struct {
	values map[string]reflect.Value
	types  map[string]reflect.Type
}

func NewTestExternalLookup() *TestExternalLookup {
	return &TestExternalLookup{
		values: make(map[string]reflect.Value),
		types:  make(map[string]reflect.Type),
	}
}

func (testExternalLookup *TestExternalLookup) SetValue(symbol string, value interface{}) error {
	if strings.Contains(symbol, ".") {
		return ErrSymbolContainsDot
	}

	if value == nil {
		testExternalLookup.values[symbol] = NilValue
	} else {
		testExternalLookup.values[symbol] = reflect.ValueOf(value)
	}

	return nil
}

func (testExternalLookup *TestExternalLookup) Get(symbol string) (reflect.Value, error) {
	if value, ok := testExternalLookup.values[symbol]; ok {
		return value, nil
	}
	return NilValue, fmt.Errorf("undefined symbol '%s'", symbol)
}

func (testExternalLookup *TestExternalLookup) DefineType(symbol string, aType interface{}) error {
	if strings.Contains(symbol, ".") {
		return ErrSymbolContainsDot
	}

	var reflectType reflect.Type
	if aType == nil {
		reflectType = NilType
	} else {
		var ok bool
		reflectType, ok = reflectType.(reflect.Type)
		if !ok {
			reflectType = reflect.TypeOf(aType)
		}
	}

	testExternalLookup.types[symbol] = reflectType
	return nil
}

func (testExternalLookup *TestExternalLookup) Type(symbol string) (reflect.Type, error) {
	if value, ok := testExternalLookup.types[symbol]; ok {
		return value, nil
	}
	return NilType, fmt.Errorf("undefined symbol '%s'", symbol)
}

func TestExternalLookupValueAndGet(t *testing.T) {
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
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface},
		{testInfo: "bool", varName: "a", varDefineValue: true, varGetValue: true, varKind: reflect.Bool},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1), varGetValue: int16(1), varKind: reflect.Int16},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), varGetValue: int32(1), varKind: reflect.Int32},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), varGetValue: uint32(1), varKind: reflect.Uint32},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), varGetValue: uint64(1), varKind: reflect.Uint64},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), varGetValue: float32(1), varKind: reflect.Float32},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64},
		{testInfo: "string", varName: "a", varDefineValue: "a", varGetValue: "a", varKind: reflect.String},

		{testInfo: "string with dot", varName: "a.a", varDefineValue: "a", varGetValue: nil, varKind: reflect.String, defineError: ErrSymbolContainsDot, getError: fmt.Errorf("undefined symbol 'a.a'")},
		{testInfo: "string with quotes", varName: "a", varDefineValue: `"a"`, varGetValue: `"a"`, varKind: reflect.String},
	}

	// ExternalLookup set And get
	for _, test := range tests {
		testExternalLookup := NewTestExternalLookup()
		env := NewEnv()
		env.SetExternalLookup(testExternalLookup)

		err = testExternalLookup.SetValue(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestExternalLookupValueAndGet %v - SetValue error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestExternalLookupValueAndGet %v - SetValue error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}
		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("TestExternalLookupValueAndGet %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("TestExternalLookupValueAndGet %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("TestExternalLookupValueAndGet %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}
	}
}

func TestExternalLookupTypeAndGet(t *testing.T) {
	var err error
	var valueType reflect.Type
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		typeError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil},
		{testInfo: "bool", varName: "a", varDefineValue: true},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1)},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1)},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1)},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1)},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1)},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1)},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1)},
		{testInfo: "string", varName: "a", varDefineValue: "a"},

		{testInfo: "string with dot", varName: "a.a", varDefineValue: nil, defineError: ErrSymbolContainsDot, typeError: fmt.Errorf("undefined type 'a.a'")},
	}

	// DefineType
	for _, test := range tests {
		testExternalLookup := NewTestExternalLookup()
		env := NewEnv()
		env.SetExternalLookup(testExternalLookup)

		err = testExternalLookup.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestExternalLookupTypeAndGet %v - DefineType error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestExternalLookupTypeAndGet %v - DefineType error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = env.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("TestExternalLookupTypeAndGet %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("TestExternalLookupTypeAndGet %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("TestExternalLookupTypeAndGet %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("TestExternalLookupTypeAndGet %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
	}

}

func TestExternalLookupAddr(t *testing.T) {
	var err error
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		addrError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, addrError: nil},
		{testInfo: "bool", varName: "a", varDefineValue: true, addrError: fmt.Errorf("unaddressable")},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), addrError: fmt.Errorf("unaddressable")},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), addrError: fmt.Errorf("unaddressable")},
		{testInfo: "string", varName: "a", varDefineValue: "a", addrError: fmt.Errorf("unaddressable")},
	}

	for _, test := range tests {
		envParent := NewEnv()
		testExternalLookup := NewTestExternalLookup()
		envParent.SetExternalLookup(testExternalLookup)
		envChild := envParent.NewEnv()

		err = testExternalLookup.SetValue(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestExternalLookupAddr %v - SetValue error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestExternalLookupAddr %v - SetValue error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		_, err = envChild.Addr(test.varName)
		if err != nil && test.addrError != nil {
			if err.Error() != test.addrError.Error() {
				t.Errorf("TestExternalLookupAddr %v - Addr error - received: %v - expected: %v", test.testInfo, err, test.addrError)
				continue
			}
		} else if err != test.addrError {
			t.Errorf("TestExternalLookupAddr %v - Addr error - received: %v - expected: %v", test.testInfo, err, test.addrError)
			continue
		}
	}
}
