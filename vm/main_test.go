package vm

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
)

type (
	testStruct1 struct {
		aInterface interface{}
		aBool      bool
		aInt32     int32
		aInt64     int64
		aFloat32   float32
		aFloat64   float32
		aString    string
		aFunc      func()

		aPtrInterface      *interface{}
		aPtrBool           *bool
		aPtrInt32          *int32
		aPtrInt64          *int64
		aPtrFloat32        *float32
		aPtrFloat64        *float32
		aPtrString         *string
		aPtrSliceInterface *[]interface{}
		aPtrSliceBool      *[]bool
		aPtrSliceInt32     *[]int32
		aPtrSliceInt64     *[]int64
		aPtrSliceFloat32   *[]float32
		aPtrSliceFloat64   *[]float32
		aPtrSliceString    *[]string

		aSliceInterface    []interface{}
		aSliceBool         []bool
		aSliceInt32        []int32
		aSliceInt64        []int64
		aSliceFloat32      []float32
		aSliceFloat64      []float32
		aSliceString       []string
		aSlicePtrInterface []*interface{}
		aSlicePtrBool      []*bool
		aSlicePtrInt32     []*int32
		aSlicePtrInt64     []*int64
		aSlicePtrFloat32   []*float32
		aSlicePtrFloat64   []*float32
		aSlicePtrString    []*string

		aMapInterface    map[string]interface{}
		aMapBool         map[string]bool
		aMapInt32        map[string]int32
		aMapInt64        map[string]int64
		aMapFloat32      map[string]float32
		aMapFloat64      map[string]float32
		aMapString       map[string]string
		aMapPtrInterface map[string]*interface{}
		aMapPtrBool      map[string]*bool
		aMapPtrInt32     map[string]*int32
		aMapPtrInt64     map[string]*int64
		aMapPtrFloat32   map[string]*float32
		aMapPtrFloat64   map[string]*float32
		aMapPtrString    map[string]*string

		aChanInterface    chan interface{}
		aChanBool         chan bool
		aChanInt32        chan int32
		aChanInt64        chan int64
		aChanFloat32      chan float32
		aChanFloat64      chan float32
		aChanString       chan string
		aChanPtrInterface chan *interface{}
		aChanPtrBool      chan *bool
		aChanPtrInt32     chan *int32
		aChanPtrInt64     chan *int64
		aChanPtrFloat32   chan *float32
		aChanPtrFloat64   chan *float32
		aChanPtrString    chan *string

		aPtrStruct *testStruct1
	}
	testStruct2 struct {
		aStruct testStruct1
	}
)

var (
	testVarValue    = reflect.Value{}
	testVarValueP   = &reflect.Value{}
	testVarBool     = true
	testVarBoolP    = &testVarBool
	testVarInt32    = int32(1)
	testVarInt32P   = &testVarInt32
	testVarInt64    = int64(1)
	testVarInt64P   = &testVarInt64
	testVarFloat32  = float32(1)
	testVarFloat32P = &testVarFloat32
	testVarFloat64  = float64(1)
	testVarFloat64P = &testVarFloat64
	testVarString   = "a"
	testVarStringP  = &testVarString
	testVarFunc     = func() int64 { return 1 }
	testVarFuncP    = &testVarFunc

	testVarValueBool    = reflect.ValueOf(true)
	testVarValueInt32   = reflect.ValueOf(int32(1))
	testVarValueInt64   = reflect.ValueOf(int64(1))
	testVarValueFloat32 = reflect.ValueOf(float32(1.1))
	testVarValueFloat64 = reflect.ValueOf(float64(1.1))
	testVarValueString  = reflect.ValueOf("a")

	testSliceEmpty []interface{}
	testSlice      = []interface{}{nil, true, int64(1), float64(1.1), "a"}
	testMapEmpty   map[interface{}]interface{}
	testMap        = map[interface{}]interface{}{"a": nil, "b": true, "c": int64(1), "d": float64(1.1), "e": "e"}
)

// Test is utility struct to make tests easy.
type Test struct {
	Script         string
	ParseError     error
	ParseErrorFunc *func(*testing.T, error)
	EnvSetupFunc   *func(*testing.T, *env.Env)
	Types          map[string]interface{}
	Input          map[string]interface{}
	RunError       error
	RunErrorFunc   *func(*testing.T, error)
	RunOutput      interface{}
	Output         map[string]interface{}
}

// TestOptions is utility struct to pass options to the test.
type TestOptions struct {
	EnvSetupFunc *func(*testing.T, *env.Env)
	Timeout      time.Duration
}

// runTests runs VM tests
func runTests(t *testing.T, tests []Test, testOptions *TestOptions, options *Options) {
	for _, test := range tests {
		runTest(t, test, testOptions, options)
	}
}

// runTest runs VM test
func runTest(t *testing.T, test Test, testOptions *TestOptions, options *Options) {
	timeout := 60 * time.Second

	// parser.EnableErrorVerbose()
	// parser.EnableDebug(8)

	stmt, err := parser.ParseSrc(test.Script)
	if test.ParseErrorFunc != nil {
		(*test.ParseErrorFunc)(t, err)
	} else if err != nil && test.ParseError != nil {
		if err.Error() != test.ParseError.Error() {
			t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.ParseError, test.Script)
		}
		return
	} else if err != test.ParseError {
		t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.ParseError, test.Script)
		return
	}
	// Note: Still want to run the code even after a parse error to see what happens

	envTest := env.NewEnv()
	if testOptions != nil {
		if testOptions.EnvSetupFunc != nil {
			(*testOptions.EnvSetupFunc)(t, envTest)
		}
		if testOptions.Timeout != 0 {
			timeout = testOptions.Timeout
		}
	}
	if test.EnvSetupFunc != nil {
		(*test.EnvSetupFunc)(t, envTest)
	}

	for typeName, typeValue := range test.Types {
		err = envTest.DefineType(typeName, typeValue)
		if err != nil {
			t.Errorf("DefineType error: %v - typeName: %v - script: %v", err, typeName, test.Script)
			return
		}
	}

	for inputName, inputValue := range test.Input {
		err = envTest.Define(inputName, inputValue)
		if err != nil {
			t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.Script)
			return
		}
	}

	var value interface{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	value, err = RunContext(ctx, envTest, options, stmt)
	cancel()
	if test.RunErrorFunc != nil {
		(*test.RunErrorFunc)(t, err)
	} else if err != nil && test.RunError != nil {
		if err.Error() != test.RunError.Error() {
			t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.RunError, test.Script)
			return
		}
	} else if err != test.RunError {
		t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.RunError, test.Script)
		return
	}

	if !valueEqual(value, test.RunOutput) {
		t.Errorf("Run output - received: %#v - expected: %#v - script: %v", value, test.RunOutput, test.Script)
		t.Errorf("received type: %T - expected: %T", value, test.RunOutput)
		return
	}

	for outputName, outputValue := range test.Output {
		value, err = envTest.Get(outputName)
		if err != nil {
			t.Errorf("Get error: %v - outputName: %v - script: %v", err, outputName, test.Script)
			return
		}

		if !valueEqual(value, outputValue) {
			t.Errorf("outputName %v - received: %#v - expected: %#v - script: %v", outputName, value, outputValue, test.Script)
			t.Errorf("received type: %T - expected: %T", value, outputValue)
			continue
		}
	}
}

// valueEqual return true if v1 and v2 is same value. If passed function, does
// extra checks otherwise just doing reflect.DeepEqual
func valueEqual(v1 interface{}, v2 interface{}) bool {
	v1RV := reflect.ValueOf(v1)
	switch v1RV.Kind() {
	case reflect.Func:
		// This is best effort to check if functions match, but it could be wrong
		v2RV := reflect.ValueOf(v2)
		if !v1RV.IsValid() || !v2RV.IsValid() {
			if v1RV.IsValid() != !v2RV.IsValid() {
				return false
			}
			return true
		} else if v1RV.Kind() != v2RV.Kind() {
			return false
		} else if v1RV.Type() != v2RV.Type() {
			return false
		} else if v1RV.Pointer() != v2RV.Pointer() {
			// From reflect: If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.
			return false
		}
		return true
	}
	switch value1 := v1.(type) {
	case error:
		switch value2 := v2.(type) {
		case error:
			return value1.Error() == value2.Error()
		}
	}

	return reflect.DeepEqual(v1, v2)
}
