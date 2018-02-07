package vm

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/mattn/anko/parser"
)

type testStruct struct {
	script     string
	parseError error
	types      map[string]interface{}
	input      map[string]interface{}
	runError   error
	runOutput  interface{}
	output     map[string]interface{}
}

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
)

func TestIf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "if true {}", input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if true {}", input: map[string]interface{}{"a": true}, runOutput: nil, output: map[string]interface{}{"a": true}},
		{script: "if true {}", input: map[string]interface{}{"a": int64(1)}, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "if true {}", input: map[string]interface{}{"a": float64(1.1)}, runOutput: nil, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "if true {}", input: map[string]interface{}{"a": "a"}, runOutput: nil, output: map[string]interface{}{"a": "a"}},

		{script: "if true {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if true {a = true}", input: map[string]interface{}{"a": int64(2)}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "if true {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "if true {a = 1.1}", input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: "if true {a = \"a\"}", input: map[string]interface{}{"a": int64(2)}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: "if a == 1 {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 2 {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "if a == 1 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 2 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: "if a == 1 {a = 1} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), output: map[string]interface{}{"a": int64(3)}},
		{script: "if a == 2 {a = 1} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 1 {a = 1} else if a == 2 {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), output: map[string]interface{}{"a": int64(3)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else {a = 4}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), output: map[string]interface{}{"a": int64(4)}},

		{script: "if a == 1 {a = 1} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if a == 2 {a = nil} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = nil} else if a == 3 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 1 {a = 1} else if a == 2 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(5), output: map[string]interface{}{"a": int64(5)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = 4} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), output: map[string]interface{}{"a": int64(4)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = nil} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
	}
	runTests(t, tests)
}

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a", input: map[string]interface{}{"a": 'a'}, runOutput: 'a', output: map[string]interface{}{"a": 'a'}},
		{script: "a.b", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support member operation"), output: map[string]interface{}{"a": 'a'}},
		{script: "a[0]", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support index operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: "a[0:1]", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support slice operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},

		{script: "a.b = \"a\"", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support member operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: "a[0] = \"a\"", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support index operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: "a[0:1] = \"a\"", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support slice operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},

		{script: "a.b = \"a\"", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support member operation"), output: map[string]interface{}{"a": "test"}},
		{script: "a[0] = \"a\"", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support index operation for assignment"), output: map[string]interface{}{"a": "test"}},
		{script: "a[0:1] = \"a\"", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support slice operation for assignment"), output: map[string]interface{}{"a": "test"}},

		{script: "a", input: map[string]interface{}{"a": "test"}, runOutput: "test", output: map[string]interface{}{"a": "test"}},
		{script: "a[\"a\"]", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": "test"}},
		{script: "a[0]", input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},
		{script: "a[-1]", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test"}},
		{script: "a[0]", input: map[string]interface{}{"a": "test"}, runOutput: "t", output: map[string]interface{}{"a": "test"}},
		{script: "a[1]", input: map[string]interface{}{"a": "test"}, runOutput: "e", output: map[string]interface{}{"a": "test"}},
		{script: "a[3]", input: map[string]interface{}{"a": "test"}, runOutput: "t", output: map[string]interface{}{"a": "test"}},
		{script: "a[4]", input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test"}},

		{script: "a", input: map[string]interface{}{"a": "\"a\""}, runOutput: "\"a\"", output: map[string]interface{}{"a": "\"a\""}},
		{script: "a[0]", input: map[string]interface{}{"a": "\"a\""}, runOutput: "\"", output: map[string]interface{}{"a": "\"a\""}},
		{script: "a[1]", input: map[string]interface{}{"a": "\"a\""}, runOutput: "a", output: map[string]interface{}{"a": "\"a\""}},

		{script: `a = "\"a\""`, runOutput: "\"a\"", output: map[string]interface{}{"a": "\"a\""}},
		{script: `a = "\"a\""; a`, runOutput: "\"a\"", output: map[string]interface{}{"a": "\"a\""}},
		{script: `a = "\"a\""; a[0]`, runOutput: "\"", output: map[string]interface{}{"a": "\"a\""}},
		{script: `a = "\"a\""; a[1]`, runOutput: "a", output: map[string]interface{}{"a": "\"a\""}},

		{script: "a", input: map[string]interface{}{"a": "a\\b"}, runOutput: "a\\b", output: map[string]interface{}{"a": "a\\b"}},
		{script: "a", input: map[string]interface{}{"a": "a\\\\b"}, runOutput: "a\\\\b", output: map[string]interface{}{"a": "a\\\\b"}},
		{script: "a = \"a\\b\"", runOutput: "a\b", output: map[string]interface{}{"a": "a\b"}},
		{script: "a = \"a\\\\b\"", runOutput: "a\\b", output: map[string]interface{}{"a": "a\\b"}},

		{script: "a[:]", input: map[string]interface{}{"a": "test data"}, parseError: fmt.Errorf("syntax error"), output: map[string]interface{}{"a": "test data"}},

		{script: "a[0:]", input: map[string]interface{}{"a": ""}, runOutput: "", output: map[string]interface{}{"a": ""}},
		{script: "a[1:]", input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},
		{script: "a[:0]", input: map[string]interface{}{"a": ""}, runOutput: "", output: map[string]interface{}{"a": ""}},
		{script: "a[:1]", input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},

		{script: "a[1:0]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": "test data"}},
		{script: "a[-1:2]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:-2]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: "a[-1:]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: "a[:-2]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: "a[0:0]", input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:1]", input: map[string]interface{}{"a": "test data"}, runOutput: "t", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:2]", input: map[string]interface{}{"a": "test data"}, runOutput: "te", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:3]", input: map[string]interface{}{"a": "test data"}, runOutput: "tes", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:7]", input: map[string]interface{}{"a": "test data"}, runOutput: "test da", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:8]", input: map[string]interface{}{"a": "test data"}, runOutput: "test dat", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[0:10]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: "a[1:1]", input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:2]", input: map[string]interface{}{"a": "test data"}, runOutput: "e", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:3]", input: map[string]interface{}{"a": "test data"}, runOutput: "es", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:7]", input: map[string]interface{}{"a": "test data"}, runOutput: "est da", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:8]", input: map[string]interface{}{"a": "test data"}, runOutput: "est dat", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:10]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: "a[0:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[2:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "st data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[3:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "t data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[7:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "ta", output: map[string]interface{}{"a": "test data"}},
		{script: "a[8:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "a", output: map[string]interface{}{"a": "test data"}},
		{script: "a[9:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},

		{script: "a[:0]", input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:1]", input: map[string]interface{}{"a": "test data"}, runOutput: "t", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:2]", input: map[string]interface{}{"a": "test data"}, runOutput: "te", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:3]", input: map[string]interface{}{"a": "test data"}, runOutput: "tes", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:7]", input: map[string]interface{}{"a": "test data"}, runOutput: "test da", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:8]", input: map[string]interface{}{"a": "test data"}, runOutput: "test dat", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:9]", input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[:10]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: "a[0:]", input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[1:]", input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[2:]", input: map[string]interface{}{"a": "test data"}, runOutput: "st data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[3:]", input: map[string]interface{}{"a": "test data"}, runOutput: "t data", output: map[string]interface{}{"a": "test data"}},
		{script: "a[7:]", input: map[string]interface{}{"a": "test data"}, runOutput: "ta", output: map[string]interface{}{"a": "test data"}},
		{script: "a[8:]", input: map[string]interface{}{"a": "test data"}, runOutput: "a", output: map[string]interface{}{"a": "test data"}},
		{script: "a[9:]", input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: "a[10:]", input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
	}
	runTests(t, tests)
}

func TestVar(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "var a = 1", runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "var a, b = 1, 2", runOutput: int64(2), output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "var a, b, c = 1, 2, 3", runOutput: int64(3), output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
	}
	runTests(t, tests)
}

func TestModule(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "module a { }; a.b", runError: fmt.Errorf("Invalid operation 'b'"), runOutput: nil},
		{script: "module a { b = nil }; a.b", runOutput: nil},
		{script: "module a { b = true }; a.b", runOutput: true},
		{script: "module a { b = 1 }; a.b", runOutput: int64(1)},
		{script: "module a { b = 1.1 }; a.b", runOutput: float64(1.1)},
		{script: "module a { b = \"a\" }; a.b", runOutput: "a"},
	}
	runTests(t, tests)
}

func TestMake(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "make(nilT)", types: map[string]interface{}{"nilT": nil}, runError: fmt.Errorf("invalid type for make")},

		{script: "make(bool)", types: map[string]interface{}{"bool": true}, runOutput: false},
		{script: "make(int32)", types: map[string]interface{}{"int32": int32(1)}, runOutput: int32(0)},
		{script: "make(int64)", types: map[string]interface{}{"int64": int64(1)}, runOutput: int64(0)},
		{script: "make(float32)", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: float32(0)},
		{script: "make(float64)", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: float64(0)},
		{script: "make(string)", types: map[string]interface{}{"string": "a"}, runOutput: ""},

		{script: "make(array2x)", types: map[string]interface{}{"array2x": [][]interface{}{}}, runOutput: [][]interface{}(nil)},
	}
	runTests(t, tests)
}

func TestForLoop(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "break", runError: fmt.Errorf("Unexpected break statement")},
		{script: "continue", runError: fmt.Errorf("Unexpected continue statement")},

		{script: "for { break }", runOutput: nil},
		{script: "for {a = 1; if a == 1 { break } }", runOutput: nil},
		{script: "a = 1; for { if a == 1 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for { if a == 1 { break }; a++ }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},

		{script: "a = 1; for { a++; if a == 2 { continue } else { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: "a = 1; for { a++; if a == 2 { continue }; if a == 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: "for a in [1] { if a == 1 { break } }", runOutput: nil},
		{script: "for a in [1, 2] { if a == 2 { break } }", runOutput: nil},
		{script: "for a in [1, 2, 3] { if a == 3 { break } }", runOutput: nil},

		{script: "a = [1]; for b in a { if b == 1 { break } }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1, 2]; for b in a { if b == 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{script: "a = [1, 2, 3]; for b in a { if b == 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1]; b = 0; for c in a { b = c }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: "a = [1, 2]; b = 0; for c in a { b = c }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": int64(2)}},
		{script: "a = [1, 2, 3]; b = 0; for c in a { b = c }", runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}, "b": int64(3)}},

		{script: "a = 1; for a < 2 { a++ }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for a < 3 { a++ }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: "a = 1; for nil { a++; if a > 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for nil { a++; if a > 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for true { a++; if a > 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: "a = 1; for true { a++; if a > 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(4)}},

		{script: "func x() { return [1] }; for b in x() { if b == 1 { break } }", runOutput: nil},
		{script: "func x() { return [1, 2] }; for b in x() { if b == 2 { break } }", runOutput: nil},
		{script: "func x() { return [1, 2, 3] }; for b in x() { if b == 3 { break } }", runOutput: nil},

		{script: "func x() { a = 1; for { if a == 1 { return } } }; x()", runOutput: nil},
		{script: "func x() { a = 1; for { if a == 1 { return nil } } }; x()", runOutput: nil},
		{script: "func x() { a = 1; for { if a == 1 { return true } } }; x()", runOutput: true},
		{script: "func x() { a = 1; for { if a == 1 { return 1 } } }; x()", runOutput: int64(1)},
		{script: "func x() { a = 1; for { if a == 1 { return 1.1 } } }; x()", runOutput: float64(1.1)},
		{script: "func x() { a = 1; for { if a == 1 { return \"a\" } } }; x()", runOutput: "a"},

		{script: "func x() { for a in [1, 2, 3] { if a == 3 { return } } }; x()", runOutput: nil},
		{script: "func x() { for a in [1, 2, 3] { if a == 1 { continue } } }; x()", runOutput: nil},
		{script: "func x() { for a in [1, 2, 3] { if a == 1 { continue };  if a == 3 { return } } }; x()", runOutput: nil},

		{script: "func x() { return [1, 2] }; a = 1; for i in x() { a++ }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: "func x() { return [1, 2, 3] }; a = 1; for i in x() { a++ }", runOutput: nil, output: map[string]interface{}{"a": int64(4)}},

		{script: "for a = 1; nil; nil { if a == 1 { break } }", runOutput: nil},
		{script: "for a = 1; nil; nil { if a == 2 { break }; a++ }", runOutput: nil},
		{script: "for a = 1; nil; nil { a++; if a == 3 { break } }", runOutput: nil},

		{script: "for a = 1; a < 1; nil { }", runOutput: nil},
		{script: "for a = 1; a > 1; nil { }", runOutput: nil},
		{script: "for a = 1; a == 1; nil { break }", runOutput: nil},

		{script: "for a = 1; a == 1; a++ { }", runOutput: nil},
		{script: "for a = 1; a < 2; a++ { }", runOutput: nil},
		{script: "for a = 1; a < 3; a++ { }", runOutput: nil},

		{script: "a = 1; for b = 1; a < 1; a++ { }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ { }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ { }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 1 { continue } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 1 { continue } }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 1 { continue } }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 1 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 1 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 1 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 2 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 3 { break } }", runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: "func x() { for a = 1; a < 3; a++ { if a == 1 { return a } } }; x()", runOutput: int64(1)},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 2 { return a } } }; x()", runOutput: int64(2)},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 3 { return a } } }; x()", runOutput: nil},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 4 { return a } } }; x()", runOutput: nil},

		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 1 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 2 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 3 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 4 { continue } }; return a }; x()", runOutput: int64(3)},

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{reflect.Value{}}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{reflect.Value{}}, "b": reflect.Value{}}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{nil}, "b": nil}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{true}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{true}, "b": true}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{int32(1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int32(1)}, "b": int32(1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{int64(1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": float32(1.1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": float64(1.1)}},

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}, "b": interface{}(reflect.Value{})}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(nil)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(nil)}, "b": interface{}(nil)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(true)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(true)}, "b": interface{}(true)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}, "b": interface{}(int32(1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}, "b": interface{}(int64(1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}, "b": interface{}(float32(1.1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}, "b": interface{}(float64(1.1))}},

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": interface{}([]interface{}{nil})}, runOutput: nil, output: map[string]interface{}{"a": interface{}([]interface{}{nil}), "b": nil}},

		{script: "for i in nil { }", runError: fmt.Errorf("for cannot loop over type interface")},
		{script: "for i in true { }", runError: fmt.Errorf("for cannot loop over type bool")},
		{script: "for i in a { }", input: map[string]interface{}{"a": reflect.Value{}}, runError: fmt.Errorf("for cannot loop over type struct"), output: map[string]interface{}{"a": reflect.Value{}}},
		{script: "for i in a { }", input: map[string]interface{}{"a": interface{}(nil)}, runError: fmt.Errorf("for cannot loop over type interface"), output: map[string]interface{}{"a": interface{}(nil)}},
		{script: "for i in a { }", input: map[string]interface{}{"a": interface{}(true)}, runError: fmt.Errorf("for cannot loop over type bool"), output: map[string]interface{}{"a": interface{}(true)}},
		{script: "for i in [1, 2, 3] { b++ }", runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: "for a = 1; a < 3; a++ { b++ }", runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: "for a = b; a < 3; a++ { }", runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: "for a = 1; b < 3; a++ { }", runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: "for a = 1; a < 3; b++ { }", runError: fmt.Errorf("Undefined symbol 'b'")},

		{script: "a = 1; b = [{\"c\": \"c\"}]; for i in b { a = i }", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"c": "c"}, "b": []interface{}{map[string]interface{}{"c": "c"}}}},
		{script: "a = 1; b = {\"x\": [{\"y\": \"y\"}]};  for i in b.x { a = i }", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"y": "y"}, "b": map[string]interface{}{"x": []interface{}{map[string]interface{}{"y": "y"}}}}},

		{script: "a = {}; b = 1; for i in a { b = i }; b", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{}, "b": int64(1)}},
		{script: "a = {\"x\": 2}; b = 1; for i in a { b = i }; b", runOutput: "x", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": "x"}},
		{script: "a = {\"x\": 2, \"y\": 3}; b = 0; for i in a { b++ }; b", runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(2)}},
		{script: "a = {\"x\": 2, \"y\": 3}; for i in a { b++ }", runError: fmt.Errorf("Undefined symbol 'b'"), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},

		{script: "a = {\"x\": 2, \"y\": 3}; b = 0; for i in a { if i ==  \"x\" { continue }; b = i }; b", runOutput: "y", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "y"}},
		{script: "a = {\"x\": 2, \"y\": 3}; b = 0; for i in a { if i ==  \"y\" { continue }; b = i }; b", runOutput: "x", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "x"}},
		{script: "a = {\"x\": 2, \"y\": 3}; for i in a { if i ==  \"x\" { return 1 } }", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: "a = {\"x\": 2, \"y\": 3}; for i in a { if i ==  \"y\" { return 2 } }", runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: "a = {\"x\": 2, \"y\": 3}; b = 0; for i in a { if i ==  \"x\" { break }; b++ }; if b > 1 { return false } else { return true }", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: "a = {\"x\": 2, \"y\": 3}; b = 0; for i in a { if i ==  \"y\" { break }; b++ }; if b > 1 { return false } else { return true }", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: "a = {\"x\": 2, \"y\": 3}; b = 1; for i in a { if (i ==  \"x\" || i ==  \"y\") { break }; b++ }; b", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(1)}},
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
				t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.parseError, test.script)
				continue
			}
		} else if err != test.parseError {
			t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.parseError, test.script)
			continue
		}

		env := NewEnv()

		for inputName, inputValue := range test.input {
			err = env.Define(inputName, inputValue)
			if err != nil {
				t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.script)
				continue loop
			}
		}

		for typeName, typeValue := range test.types {
			err = env.DefineType(typeName, typeValue)
			if err != nil {
				t.Errorf("DefineType error: %v - typeName: %v - script: %v", err, typeName, test.script)
				continue loop
			}
		}

		value, err = Run(stmts, env)
		if err != nil && test.runError != nil {
			if err.Error() != test.runError.Error() {
				t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.runError, test.script)
				continue
			}
		} else if err != test.runError {
			t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.runError, test.script)
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

func TestInterrupts(t *testing.T) {
	scripts := []string{
		`
closeWaitChan()
for {
}
`,
		`
closeWaitChan()
for {
	for {
	}
}
`,
		`
a = []
for i = 0; i < 10000; i++ {
	a += 1
}
closeWaitChan()
for i in a {
}
`,
		`
a = []
for i = 0; i < 10000; i++ {
	a += 1
}
closeWaitChan()
for i in a {
	for j in a {
	}
}
`,
		`
closeWaitChan()
for i = 0; true; nil {
}
`,
		`
closeWaitChan()
for i = 0; true; nil {
	for j = 0; true; nil {
	}
}
`,
		`
a = {}
for i = 0; i < 10000; i++ {
	a[toString(i)] = 1
}
closeWaitChan()
for i in a {
}
`,
		`
a = {}
for i = 0; i < 10000; i++ {
	a[toString(i)] = 1
}
closeWaitChan()
for i in a {
	for j in a {
	}
}
`,
	}
	for _, script := range scripts {
		runInterruptTest(t, script)
	}
}

func runInterruptTest(t *testing.T, script string) {
	waitChan := make(chan struct{}, 1)
	closeWaitChan := func() {
		close(waitChan)
	}
	toString := func(value interface{}) string {
		return fmt.Sprintf("%v", value)
	}
	env := NewEnv()
	err := env.Define("closeWaitChan", closeWaitChan)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = env.Define("toString", toString)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}

	go func() {
		<-waitChan
		Interrupt(env)
	}()

	_, err = env.Execute(script)
	if err == nil {
		t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
	} else if err.Error() != InterruptError.Error() {
		t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
	}
}

func TestInterruptConcurrency(t *testing.T) {
	var waitGroup sync.WaitGroup
	env := NewEnv()

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := env.Execute("for { }")
			if err == nil {
				t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
			} else if err.Error() != InterruptError.Error() {
				t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()

	_, err := env.Execute("for { }")
	if err == nil {
		t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
	} else if err.Error() != InterruptError.Error() {
		t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
	}

	ClearInterrupt(env)

	_, err = env.Execute("for i = 0; i < 1000; i ++ {}")
	if err != nil {
		t.Errorf("Execute error - received %v expected: %v", err, nil)
	}

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := env.Execute("for { }")
			if err == nil {
				t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
			} else if err.Error() != InterruptError.Error() {
				t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()
}
