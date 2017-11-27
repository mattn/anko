package core

import (
	"os"
	"reflect"
	"testing"

	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

type testStruct struct {
	script     string
	parseError error
	input      map[string]interface{}
	runError   error
	runOutput  interface{}
	ouput      map[string]interface{}
}

func TestLen(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = \"\"; len(a)", runOutput: int64(0)},
		{script: "a = \"test\"; len(a)", runOutput: int64(4)},
		{script: "a = []; len(a)", runOutput: int64(0)},
		{script: "a = [nil]; len(a)", runOutput: int64(1)},
		{script: "a = [true]; len(a)", runOutput: int64(1)},
		{script: "a = [\"test\"]; len(a)", runOutput: int64(1)},
		{script: "a = [1]; len(a)", runOutput: int64(1)},
		{script: "a = [1.1]; len(a)", runOutput: int64(1)},

		{script: "a = [[]]; len(a)", runOutput: int64(1)},
		{script: "a = [[nil]]; len(a)", runOutput: int64(1)},
		{script: "a = [[true]]; len(a)", runOutput: int64(1)},
		{script: "a = [[\"test\"]]; len(a)", runOutput: int64(1)},
		{script: "a = [[1]]; len(a)", runOutput: int64(1)},
		{script: "a = [[1.1]]; len(a)", runOutput: int64(1)},

		{script: "a = [[]]; len(a[0])", runOutput: int64(0)},
		{script: "a = [[nil]]; len(a[0])", runOutput: int64(1)},
		{script: "a = [[true]]; len(a[0])", runOutput: int64(1)},
		{script: "a = [[\"test\"]]; len(a[0])", runOutput: int64(1)},
		{script: "a = [[1]]; len(a[0])", runOutput: int64(1)},
		{script: "a = [[1.1]]; len(a[0])", runOutput: int64(1)},

		{script: "len(a)", input: map[string]interface{}{"a": "a"}, runOutput: int64(1), ouput: map[string]interface{}{"a": "a"}},
		{script: "len(a)", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: int64(0), ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "len(a)", input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},
		{script: "len(a[\"test\"])", input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, runOutput: int64(4), ouput: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},

		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{}}, runOutput: int64(0), ouput: map[string]interface{}{"a": []interface{}{}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{true}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{true}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{"test"}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{int32(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{int32(1)}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{int64(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{float32(1.1)}}},
		{script: "len(a)", input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "len(a[0])", input: map[string]interface{}{"a": []interface{}{"test"}}, runOutput: int64(4), ouput: map[string]interface{}{"a": []interface{}{"test"}}},

		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: int64(0), ouput: map[string]interface{}{"a": [][]interface{}{}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{nil}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{nil}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{int32(1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{int32(1)}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{float32(1.1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{float32(1.1)}}}},
		{script: "len(a)", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}},

		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{nil}}, runOutput: int64(0), ouput: map[string]interface{}{"a": [][]interface{}{nil}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{int32(1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{int32(1)}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{float32(1.1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{float32(1.1)}}}},
		{script: "len(a[0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}, runOutput: int64(1), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}},

		{script: "len(a[0][0])", input: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}, runOutput: int64(4), ouput: map[string]interface{}{"a": [][]interface{}{[]interface{}{"test"}}}},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests []testStruct) {
	var value reflect.Value
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
		LoadAllBuiltins(env)

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
		} else if err != test.runError {
			t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
			continue
		}
		if !value.IsValid() || !value.CanInterface() {
			if !reflect.DeepEqual(value, test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		} else {
			if !reflect.DeepEqual(value.Interface(), test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		}

		for outputName, outputValue := range test.ouput {
			value, err = env.Get(outputName)
			if err != nil {
				t.Errorf("Get error: %v - outputName: %v - script: %v", err, outputName, test.script)
				continue loop
			}
			if !value.IsValid() || !value.CanInterface() {
				if !reflect.DeepEqual(value, outputValue) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			} else {
				if !reflect.DeepEqual(value.Interface(), outputValue) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			}
		}

		env.Destroy()
	}
}
