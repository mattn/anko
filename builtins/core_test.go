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
	types      map[string]interface{}
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

func TestKeys(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = {}; b = keys(a)", runOutput: []string{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"a\": nil}; b = keys(a)", runOutput: []string{"a"}, ouput: map[string]interface{}{"a": map[string]interface{}{"a": nil}}},
		{script: "a = {\"a\": 1}; b = keys(a)", runOutput: []string{"a"}, ouput: map[string]interface{}{"a": map[string]interface{}{"a": int64(1)}}},
	}
	runTests(t, tests)
}

func TestKindOf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "kindOf(a)", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: "nil", ouput: map[string]interface{}{"a": reflect.Value{}}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": nil}, runOutput: "nil", ouput: map[string]interface{}{"a": nil}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": interface{}(nil)}, runOutput: "nil", ouput: map[string]interface{}{"a": interface{}(nil)}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": true}, runOutput: "bool", ouput: map[string]interface{}{"a": true}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": int32(1)}, runOutput: "int32", ouput: map[string]interface{}{"a": int32(1)}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": int64(1)}, runOutput: "int64", ouput: map[string]interface{}{"a": int64(1)}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": float32(1.1)}, runOutput: "float32", ouput: map[string]interface{}{"a": float32(1.1)}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": float64(1.1)}, runOutput: "float64", ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": "a"}, runOutput: "string", ouput: map[string]interface{}{"a": "a"}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": 'a'}, runOutput: "int32", ouput: map[string]interface{}{"a": 'a'}},

		{script: "kindOf(a)", input: map[string]interface{}{"a": []interface{}{}}, runOutput: "slice", ouput: map[string]interface{}{"a": []interface{}{}}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: "slice", ouput: map[string]interface{}{"a": []interface{}{nil}}},

		{script: "kindOf(a)", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: "map", ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "kindOf(a)", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "map", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},
	}
	runTests(t, tests)
}

func TestStrconv(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "strconv = import(\"strconv\"); a = true; b = strconv.FormatBool(a)", runOutput: "true", ouput: map[string]interface{}{"a": true, "b": "true"}},
		{script: "strconv = import(\"strconv\"); a = 1.1; b = strconv.FormatFloat(a, toRune(\"f\"), -1, 64)", runOutput: "1.1", ouput: map[string]interface{}{"a": float64(1.1), "b": "1.1"}},
		{script: "strconv = import(\"strconv\"); a = 1; b = strconv.FormatInt(a, 10)", runOutput: "1", ouput: map[string]interface{}{"a": int64(1), "b": "1"}},
		{script: "strconv = import(\"strconv\"); b = strconv.FormatInt(a, 10)", input: map[string]interface{}{"a": uint64(1)}, runOutput: "1", ouput: map[string]interface{}{"a": uint64(1), "b": "1"}},

		{script: "strconv = import(\"strconv\"); a = \"true\"; b, err = strconv.ParseBool(a); err = toString(err)", runOutput: "<nil>", ouput: map[string]interface{}{"a": "true", "b": true, "err": "<nil>"}},
		{script: "strconv = import(\"strconv\"); a = \"2\"; b, err = strconv.ParseBool(a); err = toString(err)", runOutput: "strconv.ParseBool: parsing \"2\": invalid syntax", ouput: map[string]interface{}{"a": "2", "b": false, "err": "strconv.ParseBool: parsing \"2\": invalid syntax"}},
		{script: "strconv = import(\"strconv\"); a = \"1.1\"; b, err = strconv.ParseFloat(a, 64); err = toString(err)", runOutput: "<nil>", ouput: map[string]interface{}{"a": "1.1", "b": float64(1.1), "err": "<nil>"}},
		{script: "strconv = import(\"strconv\"); a = \"a\"; b, err = strconv.ParseFloat(a, 64); err = toString(err)", runOutput: "strconv.ParseFloat: parsing \"a\": invalid syntax", ouput: map[string]interface{}{"a": "a", "b": float64(0), "err": "strconv.ParseFloat: parsing \"a\": invalid syntax"}},
		{script: "strconv = import(\"strconv\"); a = \"1\"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)", runOutput: "<nil>", ouput: map[string]interface{}{"a": "1", "b": int64(1), "err": "<nil>"}},
		{script: "strconv = import(\"strconv\"); a = \"1.1\"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)", runOutput: "strconv.ParseInt: parsing \"1.1\": invalid syntax", ouput: map[string]interface{}{"a": "1.1", "b": int64(0), "err": "strconv.ParseInt: parsing \"1.1\": invalid syntax"}},
		{script: "strconv = import(\"strconv\"); a = \"a\"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)", runOutput: "strconv.ParseInt: parsing \"a\": invalid syntax", ouput: map[string]interface{}{"a": "a", "b": int64(0), "err": "strconv.ParseInt: parsing \"a\": invalid syntax"}},
		{script: "strconv = import(\"strconv\"); a = \"1\"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)", runOutput: "<nil>", ouput: map[string]interface{}{"a": "1", "b": uint64(1), "err": "<nil>"}},
		{script: "strconv = import(\"strconv\"); a = \"a\"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)", runOutput: "strconv.ParseUint: parsing \"a\": invalid syntax", ouput: map[string]interface{}{"a": "a", "b": uint64(0), "err": "strconv.ParseUint: parsing \"a\": invalid syntax"}},
	}
	runTests(t, tests)
}

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "strings = import(\"strings\"); a = \" one two \"; b = strings.TrimSpace(a)", runOutput: "one two", ouput: map[string]interface{}{"a": " one two ", "b": "one two"}},
		{script: "strings = import(\"strings\"); a = \"a b c d\"; b = strings.Split(a, \" \")", runOutput: []string{"a", "b", "c", "d"}, ouput: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c", "d"}}},
		{script: "strings = import(\"strings\"); a = \"a b c d\"; b = strings.SplitN(a, \" \", 3)", runOutput: []string{"a", "b", "c d"}, ouput: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c d"}}},
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
		} else if err != test.runError {
			t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
			continue
		}
		if !value.IsValid() || !value.CanInterface() {
			if !reflect.DeepEqual(value, test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		} else if value.Kind() == reflect.Func {
			if !reflect.DeepEqual(value, reflect.ValueOf(test.runOutput)) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		} else {
			if !reflect.DeepEqual(value.Interface(), test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, reflect.ValueOf(test.runOutput), test.script)
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
			} else if value.Kind() == reflect.Func {
				if !reflect.DeepEqual(value, reflect.ValueOf(outputValue)) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			} else {
				if !reflect.DeepEqual(value.Interface(), outputValue) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, reflect.ValueOf(outputValue), test.script)
					continue
				}
			}
		}

		env.Destroy()
	}
}
