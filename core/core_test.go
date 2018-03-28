package core

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

type testStruct struct {
	script     string
	parseError error
	types      map[string]interface{}
	input      map[string]interface{}
	runError   error
	runOutput  interface{}
	checkError func(error) bool
	output     map[string]interface{}
}

func TestKeys(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `a = {}; b = keys(a)`, runOutput: []string{}, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: `a = {"a": nil}; b = keys(a)`, runOutput: []string{"a"}, output: map[string]interface{}{"a": map[string]interface{}{"a": nil}}},
		{script: `a = {"a": 1}; b = keys(a)`, runOutput: []string{"a"}, output: map[string]interface{}{"a": map[string]interface{}{"a": int64(1)}}},
		{script: `a = [{"a": 1}][0]; b = keys(a)`, runOutput: []string{"a"}, output: map[string]interface{}{"a": map[string]interface{}{"a": int64(1)}}},
	}
	runTests(t, tests)
}

func TestKindOf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `kindOf(a)`, input: map[string]interface{}{"a": reflect.Value{}}, runOutput: "struct", output: map[string]interface{}{"a": reflect.Value{}}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": nil}, runOutput: "nil", output: map[string]interface{}{"a": nil}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": true}, runOutput: "bool", output: map[string]interface{}{"a": true}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": int32(1)}, runOutput: "int32", output: map[string]interface{}{"a": int32(1)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": int64(1)}, runOutput: "int64", output: map[string]interface{}{"a": int64(1)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": float32(1.1)}, runOutput: "float32", output: map[string]interface{}{"a": float32(1.1)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": float64(1.1)}, runOutput: "float64", output: map[string]interface{}{"a": float64(1.1)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": "a"}, runOutput: "string", output: map[string]interface{}{"a": "a"}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": 'a'}, runOutput: "int32", output: map[string]interface{}{"a": 'a'}},

		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(nil)}, runOutput: "nil", output: map[string]interface{}{"a": interface{}(nil)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(true)}, runOutput: "bool", output: map[string]interface{}{"a": interface{}(true)}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(int32(1))}, runOutput: "int32", output: map[string]interface{}{"a": interface{}(int32(1))}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(int64(1))}, runOutput: "int64", output: map[string]interface{}{"a": interface{}(int64(1))}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(float32(1))}, runOutput: "float32", output: map[string]interface{}{"a": interface{}(float32(1))}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}(float64(1))}, runOutput: "float64", output: map[string]interface{}{"a": interface{}(float64(1))}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": interface{}("a")}, runOutput: "string", output: map[string]interface{}{"a": interface{}("a")}},

		{script: `kindOf(a)`, input: map[string]interface{}{"a": []interface{}{}}, runOutput: "slice", output: map[string]interface{}{"a": []interface{}{}}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: "slice", output: map[string]interface{}{"a": []interface{}{nil}}},

		{script: `kindOf(a)`, input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: "map", output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: `kindOf(a)`, input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "map", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: `a = make(interface); kindOf(a)`, runOutput: "nil", output: map[string]interface{}{"a": interface{}(nil)}},
	}
	runTests(t, tests)
}

func TestRange(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	tests := []testStruct{
		{script: `range()`, runError: errors.New("Missing arguments")},
		{script: `range(-1)`, runOutput: []int64{}},
		{script: `range(0)`, runOutput: []int64{}},
		{script: `range(1)`, runOutput: []int64{0}},
		{script: `range(2)`, runOutput: []int64{0, 1}},
		{script: `range(10)`, runOutput: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{script: `range(-1,1)`, runOutput: []int64{-1, 0, 1}},
		{script: `range(-1,0,1)`, runError: errors.New("Too many arguments")},
	}
	runTests(t, tests)
}

func TestLoad(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	tests := []testStruct{
		{script: `load('testdata/test.ank'); X(1)`, runOutput: int64(2)},
		{script: `load('testdata/not-found.ank'); X(1)`, checkError: func(err error) bool { return strings.HasPrefix(err.Error(), "open testdata/not-found.ank:") }},
		{script: `load('testdata/broken.ank'); X(1)`, runError: errors.New("syntax error")},
		{script: `load('testdata/error.ank'); X(1)`, runError: errors.New("type int64 does not support member operation")},
	}
	runTests(t, tests)
}

func TestDefined(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	tests := []testStruct{
		{script: `var a = 1; defined("a")`, runOutput: true},
		{script: `defined("a")`, runOutput: false},
		{script: `func(){ var a = 1 }(); defined("a")`, runOutput: false},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests []testStruct) {
	var value interface{}
loop:
	for _, test := range tests {
		stmts, err := parser.ParseSrc(test.script)
		if err != nil && test.parseError != nil {
			if err.Error() != test.parseError.Error() {
				t.Errorf("ParseSrc error - received: %v expected: %v - script: %v", err, test.parseError, test.script)
				continue
			}
		} else if err != test.parseError {
			t.Errorf("ParseSrc error - received: %v expected: %v - script: %v", err, test.parseError, test.script)
			continue
		}

		env := vm.NewEnv()
		Import(env)

		for typeName, typeValue := range test.types {
			err = env.DefineType(typeName, typeValue)
			if err != nil {
				t.Errorf("DefineType error: %v - typeName: %v - script: %v", err, typeName, test.script)
				continue loop
			}
		}

		for inputName, inputValue := range test.input {
			err = env.Define(inputName, inputValue)
			if err != nil {
				t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.script)
				continue loop
			}
		}

		value, err = vm.Run(stmts, env)
		if err != nil && test.runError != nil {
			if err.Error() != test.runError.Error() {
				t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
				continue
			}
		} else if test.checkError != nil {
			if !test.checkError(err) {
				t.Errorf("Run error - unexpected error: %v - script: %v", err, test.script)
			}
		} else if err != test.runError {
			t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
			continue
		}

		switch reflect.ValueOf(value).Kind() {
		case reflect.Func:
			// This is best effort to check if functions match, but it could be wrong
			valueV := reflect.ValueOf(value)
			runOutputV := reflect.ValueOf(test.runOutput)
			if !valueV.IsValid() || !runOutputV.IsValid() {
				if valueV.IsValid() != runOutputV.IsValid() {
					t.Errorf("Run output - received IsValid: %v - expected IsValid: %v - script: %v", valueV.IsValid(), runOutputV.IsValid(), test.script)
					continue
				}
			} else if valueV.Kind() != runOutputV.Kind() {
				t.Errorf("Run output - received Kind: %v - expected Kind: %v - script: %v", valueV.Kind(), runOutputV.Kind(), test.script)
				continue
			} else if valueV.Type() != runOutputV.Type() {
				t.Errorf("Run output - received Type: %v - expected Type: %v - script: %v", valueV.Type(), runOutputV.Type(), test.script)
				continue
			} else if valueV.Pointer() != runOutputV.Pointer() {
				// From reflect: If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.
				t.Errorf("Run output - received Pointer: %v - expected Pointer: %v - script: %v", valueV.Pointer(), runOutputV.Pointer(), test.script)
				continue
			}
		default:
			if !reflect.DeepEqual(value, test.runOutput) {
				t.Errorf("Run output - received: %#v - expected: %#v - script: %v", value, test.runOutput, test.script)
				t.Errorf("received type: %T - expected: %T", value, test.runOutput)
				continue
			}
		}

		for outputName, outputValue := range test.output {
			value, err = env.Get(outputName)
			if err != nil {
				t.Errorf("Get error: %v - outputName: %v - script: %v", err, outputName, test.script)
				continue loop
			}

			switch reflect.ValueOf(value).Kind() {
			case reflect.Func:
				// This is best effort to check if functions match, but it could be wrong
				valueV := reflect.ValueOf(value)
				outputValueV := reflect.ValueOf(outputValue)
				if !valueV.IsValid() || !outputValueV.IsValid() {
					if valueV.IsValid() != outputValueV.IsValid() {
						t.Errorf("outputName %v - received IsValid: %v - expected IsValid: %v - script: %v", outputName, valueV.IsValid(), outputValueV.IsValid(), test.script)
						continue
					}
				} else if valueV.Kind() != outputValueV.Kind() {
					t.Errorf("outputName %v - received Kind: %v - expected Kind: %v - script: %v", outputName, valueV.Kind(), outputValueV.Kind(), test.script)
					continue
				} else if valueV.Type() != outputValueV.Type() {
					t.Errorf("outputName %v - received Kind: %v - expected Kind: %v - script: %v", outputName, valueV.Type(), outputValueV.Type(), test.script)
					continue
				} else if valueV.Pointer() != outputValueV.Pointer() {
					// From reflect: If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.
					t.Errorf("outputName %v - received Pointer: %v - expected Pointer: %v - script: %v", outputName, valueV.Pointer(), outputValueV.Pointer(), test.script)
					continue
				}
			default:
				if !reflect.DeepEqual(value, outputValue) {
					t.Errorf("outputName %v - received: %#v - expected: %#v - script: %v", outputName, value, outputValue, test.script)
					t.Errorf("received type: %T - expected: %T", value, outputValue)
					continue
				}
			}
		}

		env.Destroy()
	}
}
