package env

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasicType(t *testing.T) {
	env := NewEnv()
	aType, err := env.Type("string")
	if err != nil {
		t.Fatalf("Type error - %v", err)
	}
	if aType != reflect.TypeOf("a") {
		t.Errorf("Type - received: %v - expected: %v", aType, reflect.TypeOf("a"))
	}

	aType, err = env.Type("int64")
	if err != nil {
		t.Fatal("Type error:", err)
	}
	if aType != reflect.TypeOf(int64(1)) {
		t.Errorf("Type - received: %v - expected: %v", aType, reflect.TypeOf(int64(1)))
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
		env := NewEnv()

		err = env.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = env.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
	}

	// DefineType NewEnv
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType NewEnv %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType NewEnv %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType NewEnv %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType NewEnv %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType NewEnv %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType NewEnv %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
	}

	// DefineType NewModule
	for _, test := range tests {
		envParent := NewEnv()
		envChild, err := envParent.NewModule("envChild")
		if err != nil {
			t.Fatalf("DefineType NewModule %v - NewModule error - received: %v - expected: %v", test.testInfo, err, nil)
		}

		err = envParent.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineType NewModule %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineType NewModule %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineType NewModule %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineType NewModule %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineType NewModule %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineType NewModule %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
	}

	// DefineGlobalType
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envChild.DefineGlobalType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineGlobalType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineGlobalType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envParent.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineGlobalType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineGlobalType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineGlobalType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineGlobalType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineGlobalType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineGlobalType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineGlobalType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineGlobalType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
	}

	// DefineGlobalReflectType
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envChild.DefineGlobalReflectType(test.varName, reflect.TypeOf(test.varDefineValue))
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineGlobalReflectType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineGlobalReflectType %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		valueType, err = envParent.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineGlobalReflectType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineGlobalReflectType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineGlobalReflectType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineGlobalReflectType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}

		valueType, err = envChild.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("DefineGlobalReflectType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("DefineGlobalReflectType %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
			continue
		}
		if valueType == nil || test.varDefineValue == nil {
			if valueType != reflect.TypeOf(test.varDefineValue) {
				t.Errorf("DefineGlobalReflectType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
			}
		} else if valueType.String() != reflect.TypeOf(test.varDefineValue).String() {
			t.Errorf("DefineGlobalReflectType %v - Type check - received: %v - expected: %v", test.testInfo, valueType, reflect.TypeOf(test.varDefineValue))
		}
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
		{testInfo: "nil", varName: "a", varDefineValue: nil, typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "bool", varName: "a", varDefineValue: true, typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), typeError: fmt.Errorf("undefined type 'a'")},
		{testInfo: "string", varName: "a", varDefineValue: "a", typeError: fmt.Errorf("undefined type 'a'")},
	}

	// DefineTypeFail
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envChild.DefineType(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestDefineTypeFail %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestDefineTypeFail %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		_, err = envParent.Type(test.varName)
		if err != nil && test.typeError != nil {
			if err.Error() != test.typeError.Error() {
				t.Errorf("TestDefineTypeFail %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
				continue
			}
		} else if err != test.typeError {
			t.Errorf("TestDefineTypeFail %v - Type error - received: %v - expected: %v", test.testInfo, err, test.typeError)
		}
	}
}

func TestGetTypeSymbols(t *testing.T) {
	var symbols []string
	values := map[string]interface{}{
		"a": int64(1),
		"b": float64(1),
		"c": true,
		"d": "a",
	}

	env := NewEnv()
	for s, v := range values {
		env.DefineType(s, v)
	}

	symbols = env.GetTypeSymbols()
	if len(symbols) != len(values) {
		t.Errorf("Expected %d symbols, received %d", len(values), len(symbols))
	}

	for _, symbol := range symbols {
		_, ok := values[symbol]
		if !ok {
			t.Errorf("Missing %s symbol", symbol)
		}
	}
}
