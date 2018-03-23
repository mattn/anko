package vm

import (
	"fmt"
	"io"
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

func TestNumbers(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `1..1`, runError: fmt.Errorf(`strconv.ParseFloat: parsing "1..1": invalid syntax`)},
		{script: `0x1g`, parseError: fmt.Errorf("syntax error")},
		{script: `9223372036854775808`, runError: fmt.Errorf(`strconv.ParseInt: parsing "9223372036854775808": value out of range`)},

		{script: `1`, runOutput: int64(1)},
		{script: `-1`, runOutput: int64(-1)},
		{script: `9223372036854775807`, runOutput: int64(9223372036854775807)},
		{script: `-9223372036854775807`, runOutput: int64(-9223372036854775807)},
		{script: `1.1`, runOutput: float64(1.1)},
		{script: `-1.1`, runOutput: float64(-1.1)},
		{script: `1e1`, runOutput: float64(10)},
		{script: `-1e1`, runOutput: float64(-10)},
		{script: `1e-1`, runOutput: float64(0.1)},
		{script: `-1e-1`, runOutput: float64(-0.1)},
		{script: `0x1`, runOutput: int64(1)},
		{script: `0xc`, runOutput: int64(12)},
		// TOFIX:
		{script: `0xe`, runError: fmt.Errorf(`strconv.ParseFloat: parsing "0xe": invalid syntax`)},
		{script: `0xf`, runOutput: int64(15)},
		{script: `-0x1`, runOutput: int64(-1)},
		{script: `-0xc`, runOutput: int64(-12)},
		{script: `-0xf`, runOutput: int64(-15)},
	}
	runTests(t, tests)
}

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `a`, input: map[string]interface{}{"a": 'a'}, runOutput: 'a', output: map[string]interface{}{"a": 'a'}},
		{script: `a.b`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support member operation"), output: map[string]interface{}{"a": 'a'}},
		{script: `a[0]`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support index operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: `a[0:1]`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support slice operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},

		{script: `a.b = "a"`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support member operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: `a[0] = "a"`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support index operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},
		{script: `a[0:1] = "a"`, input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support slice operation"), runOutput: nil, output: map[string]interface{}{"a": 'a'}},

		{script: `a.b = "a"`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support member operation"), output: map[string]interface{}{"a": "test"}},
		{script: `a[0] = "a"`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support index operation for assignment"), output: map[string]interface{}{"a": "test"}},
		{script: `a[0:1] = "a"`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("type string does not support slice operation for assignment"), output: map[string]interface{}{"a": "test"}},

		{script: `a`, input: map[string]interface{}{"a": "test"}, runOutput: "test", output: map[string]interface{}{"a": "test"}},
		{script: `a["a"]`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": "test"}},
		{script: `a[0]`, input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},
		{script: `a[-1]`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test"}},
		{script: `a[0]`, input: map[string]interface{}{"a": "test"}, runOutput: "t", output: map[string]interface{}{"a": "test"}},
		{script: `a[1]`, input: map[string]interface{}{"a": "test"}, runOutput: "e", output: map[string]interface{}{"a": "test"}},
		{script: `a[3]`, input: map[string]interface{}{"a": "test"}, runOutput: "t", output: map[string]interface{}{"a": "test"}},
		{script: `a[4]`, input: map[string]interface{}{"a": "test"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test"}},

		{script: `a`, input: map[string]interface{}{"a": `"a"`}, runOutput: `"a"`, output: map[string]interface{}{"a": `"a"`}},
		{script: `a[0]`, input: map[string]interface{}{"a": `"a"`}, runOutput: "\"", output: map[string]interface{}{"a": `"a"`}},
		{script: `a[1]`, input: map[string]interface{}{"a": `"a"`}, runOutput: "a", output: map[string]interface{}{"a": `"a"`}},

		{script: `a = "\"a\""`, runOutput: `"a"`, output: map[string]interface{}{"a": `"a"`}},
		{script: `a = "\"a\""; a`, runOutput: `"a"`, output: map[string]interface{}{"a": `"a"`}},
		{script: `a = "\"a\""; a[0]`, runOutput: `"`, output: map[string]interface{}{"a": `"a"`}},
		{script: `a = "\"a\""; a[1]`, runOutput: "a", output: map[string]interface{}{"a": `"a"`}},

		{script: `a`, input: map[string]interface{}{"a": "a\\b"}, runOutput: "a\\b", output: map[string]interface{}{"a": "a\\b"}},
		{script: `a`, input: map[string]interface{}{"a": "a\\\\b"}, runOutput: "a\\\\b", output: map[string]interface{}{"a": "a\\\\b"}},
		{script: `a = "a\b"`, runOutput: "a\b", output: map[string]interface{}{"a": "a\b"}},
		{script: `a = "a\\b"`, runOutput: "a\\b", output: map[string]interface{}{"a": "a\\b"}},

		{script: `a[:]`, input: map[string]interface{}{"a": "test data"}, parseError: fmt.Errorf("syntax error"), output: map[string]interface{}{"a": "test data"}},

		{script: `a[0:]`, input: map[string]interface{}{"a": ""}, runOutput: "", output: map[string]interface{}{"a": ""}},
		{script: `a[1:]`, input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},
		{script: `a[:0]`, input: map[string]interface{}{"a": ""}, runOutput: "", output: map[string]interface{}{"a": ""}},
		{script: `a[:1]`, input: map[string]interface{}{"a": ""}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": ""}},

		{script: `a[1:0]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": "test data"}},
		{script: `a[-1:2]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:-2]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: `a[-1:]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
		{script: `a[:-2]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: `a[0:0]`, input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:1]`, input: map[string]interface{}{"a": "test data"}, runOutput: "t", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:2]`, input: map[string]interface{}{"a": "test data"}, runOutput: "te", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:3]`, input: map[string]interface{}{"a": "test data"}, runOutput: "tes", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:7]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test da", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:8]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test dat", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[0:10]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: `a[1:1]`, input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:2]`, input: map[string]interface{}{"a": "test data"}, runOutput: "e", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:3]`, input: map[string]interface{}{"a": "test data"}, runOutput: "es", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:7]`, input: map[string]interface{}{"a": "test data"}, runOutput: "est da", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:8]`, input: map[string]interface{}{"a": "test data"}, runOutput: "est dat", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:10]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: `a[0:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[2:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "st data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[3:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "t data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[7:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "ta", output: map[string]interface{}{"a": "test data"}},
		{script: `a[8:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "a", output: map[string]interface{}{"a": "test data"}},
		{script: `a[9:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},

		{script: `a[:0]`, input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:1]`, input: map[string]interface{}{"a": "test data"}, runOutput: "t", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:2]`, input: map[string]interface{}{"a": "test data"}, runOutput: "te", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:3]`, input: map[string]interface{}{"a": "test data"}, runOutput: "tes", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:7]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test da", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:8]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test dat", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:9]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[:10]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},

		{script: `a[0:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "test data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[1:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "est data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[2:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "st data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[3:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "t data", output: map[string]interface{}{"a": "test data"}},
		{script: `a[7:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "ta", output: map[string]interface{}{"a": "test data"}},
		{script: `a[8:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "a", output: map[string]interface{}{"a": "test data"}},
		{script: `a[9:]`, input: map[string]interface{}{"a": "test data"}, runOutput: "", output: map[string]interface{}{"a": "test data"}},
		{script: `a[10:]`, input: map[string]interface{}{"a": "test data"}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": "test data"}},
	}
	runTests(t, tests)
}

func TestVar(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `var a = 1++`, runError: fmt.Errorf("Invalid operation")},

		{script: `var a = 1`, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: `var a, b = 1, 2`, runOutput: int64(2), output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: `var a, b, c = 1, 2, 3`, runOutput: int64(3), output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},

		{script: `var a = 1, 2`, runOutput: int64(2), output: map[string]interface{}{"a": int64(1)}},
		{script: `var a, b = 1`, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
	}
	runTests(t, tests)
}

func TestModule(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `module a.b { }`, parseError: fmt.Errorf("syntax error")},
		{script: `module a { 1++ }`, runError: fmt.Errorf("Invalid operation")},
		{script: `module a { }; a.b`, runError: fmt.Errorf("Invalid operation 'b'")},

		{script: `module a { b = nil }; a.b`, runOutput: nil},
		{script: `module a { b = true }; a.b`, runOutput: true},
		{script: `module a { b = 1 }; a.b`, runOutput: int64(1)},
		{script: `module a { b = 1.1 }; a.b`, runOutput: float64(1.1)},
		{script: `module a { b = "a" }; a.b`, runOutput: "a"},
	}
	runTests(t, tests)
}

func TestNew(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `new(foo)`, runError: fmt.Errorf("Undefined type 'foo'")},
		{script: `new(nilT)`, types: map[string]interface{}{"nilT": nil}, runError: fmt.Errorf("type cannot be nil for new")},

		{script: `a = new(bool); *a`, runOutput: false},
		{script: `a = new(int32); *a`, runOutput: int32(0)},
		{script: `a = new(int64); *a`, runOutput: int64(0)},
		{script: `a = new(float32); *a`, runOutput: float32(0)},
		{script: `a = new(float64); *a`, runOutput: float64(0)},
		{script: `a = new(string); *a`, runOutput: ""},
	}
	runTests(t, tests)
}

func TestMake(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `make(foo)`, runError: fmt.Errorf("Undefined type 'foo'")},
		{script: `make(a.b)`, types: map[string]interface{}{"a": true}, runError: fmt.Errorf("no namespace called: a")},
		{script: `make(a.b)`, types: map[string]interface{}{"b": true}, runError: fmt.Errorf("no namespace called: a")},

		{script: `make(nilT)`, types: map[string]interface{}{"nilT": nil}, runError: fmt.Errorf("type cannot be nil for make")},

		{script: `make(bool)`, runOutput: false},
		{script: `make(int32)`, runOutput: int32(0)},
		{script: `make(int64)`, runOutput: int64(0)},
		{script: `make(float32)`, runOutput: float32(0)},
		{script: `make(float64)`, runOutput: float64(0)},
		{script: `make(string)`, runOutput: ""},
	}
	runTests(t, tests)
}

func TestMakeType(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `make(type a, 1++)`, runError: fmt.Errorf("Invalid operation")},

		{script: `make(type a, true)`, runOutput: reflect.TypeOf(true)},
		{script: `a = make(type a, true)`, runOutput: reflect.TypeOf(true), output: map[string]interface{}{"a": reflect.TypeOf(true)}},
		{script: `make(type a, true); a = make([]a)`, runOutput: []bool{}, output: map[string]interface{}{"a": []bool{}}},
		{script: `make(type a, make([]bool))`, runOutput: reflect.TypeOf([]bool{})},
		{script: `make(type a, make([]bool)); a = make(a)`, runOutput: []bool{}, output: map[string]interface{}{"a": []bool{}}},
	}
	runTests(t, tests)
}

func TestLen(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `len(1++)`, runError: fmt.Errorf("Invalid operation")},
		{script: `len(true)`, runError: fmt.Errorf("type bool does not support len operation")},

		{script: `a = ""; len(a)`, runOutput: int64(0)},
		{script: `a = "test"; len(a)`, runOutput: int64(4)},
		{script: `a = []; len(a)`, runOutput: int64(0)},
		{script: `a = [nil]; len(a)`, runOutput: int64(1)},
		{script: `a = [true]; len(a)`, runOutput: int64(1)},
		{script: `a = ["test"]; len(a)`, runOutput: int64(1)},
		{script: `a = [1]; len(a)`, runOutput: int64(1)},
		{script: `a = [1.1]; len(a)`, runOutput: int64(1)},

		{script: `a = [[]]; len(a)`, runOutput: int64(1)},
		{script: `a = [[nil]]; len(a)`, runOutput: int64(1)},
		{script: `a = [[true]]; len(a)`, runOutput: int64(1)},
		{script: `a = [["test"]]; len(a)`, runOutput: int64(1)},
		{script: `a = [[1]]; len(a)`, runOutput: int64(1)},
		{script: `a = [[1.1]]; len(a)`, runOutput: int64(1)},

		{script: `a = [[]]; len(a[0])`, runOutput: int64(0)},
		{script: `a = [[nil]]; len(a[0])`, runOutput: int64(1)},
		{script: `a = [[true]]; len(a[0])`, runOutput: int64(1)},
		{script: `a = [["test"]]; len(a[0])`, runOutput: int64(1)},
		{script: `a = [[1]]; len(a[0])`, runOutput: int64(1)},
		{script: `a = [[1.1]]; len(a[0])`, runOutput: int64(1)},

		{script: `len(a)`, input: map[string]interface{}{"a": "a"}, runOutput: int64(1), output: map[string]interface{}{"a": "a"}},
		{script: `len(a)`, input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: int64(0), output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: `len(a)`, input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},
		{script: `len(a["test"])`, input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, runOutput: int64(4), output: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},

		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{}}, runOutput: int64(0), output: map[string]interface{}{"a": []interface{}{}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{true}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{true}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{"test"}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{int32(1)}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{int32(1)}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{int64(1)}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{float32(1.1)}}},
		{script: `len(a)`, input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, runOutput: int64(1), output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: `len(a[0])`, input: map[string]interface{}{"a": []interface{}{"test"}}, runOutput: int64(4), output: map[string]interface{}{"a": []interface{}{"test"}}},

		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: int64(0), output: map[string]interface{}{"a": [][]interface{}{}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{nil}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{nil}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{nil}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{true}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{"test"}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{"test"}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}},
		{script: `len(a)`, input: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},

		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{nil}}, runOutput: int64(0), output: map[string]interface{}{"a": [][]interface{}{nil}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{nil}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{true}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{"test"}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{"test"}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}},
		{script: `len(a[0])`, input: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},

		{script: `len(a[0][0])`, input: map[string]interface{}{"a": [][]interface{}{{"test"}}}, runOutput: int64(4), output: map[string]interface{}{"a": [][]interface{}{{"test"}}}},
	}
	runTests(t, tests)
}

func TestReferencingAndDereference(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		// TOFIX:
		// {script: `a = 1; b = &a; *b = 2; *b`, runOutput: int64(2), output: map[string]interface{}{"a": int64(2)}},
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
b = 0
closeWaitChan()
for {
	b = 1
}
`,
		`
b = 0
closeWaitChan()
for {
	for {
		b = 1
	}
}
`,
		`
a = []
for i = 0; i < 10000; i++ {
	a += 1
}
b = 0
closeWaitChan()
for i in a {
	b = i
}
`,
		`
a = []
for i = 0; i < 10000; i++ {
	a += 1
}
b = 0
closeWaitChan()
for i in a {
	for j in a {
		b = j
	}
}
`,
		`
closeWaitChan()
b = 0
for i = 0; true; nil {
}
`,
		`
b = 0
closeWaitChan()
for i = 0; true; nil {
	for j = 0; true; nil {
		b = 1
	}
}
`,
		`
a = {}
for i = 0; i < 10000; i++ {
	a[toString(i)] = 1
}
b = 0
closeWaitChan()
for i in a {
	b = 1
}
`,
		`
a = {}
for i = 0; i < 10000; i++ {
	a[toString(i)] = 1
}
b = 0
closeWaitChan()
for i in a {
	for j in a {
		b = 1
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
	if err == nil || err.Error() != ErrInterrupt.Error() {
		t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
}

func TestInterruptConcurrency(t *testing.T) {
	var waitGroup sync.WaitGroup
	env := NewEnv()

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := env.Execute("for { }")
			if err == nil || err.Error() != ErrInterrupt.Error() {
				t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()

	_, err := env.Execute("for { }")
	if err == nil || err.Error() != ErrInterrupt.Error() {
		t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}

	ClearInterrupt(env)

	_, err = env.Execute("for i = 0; i < 1000; i ++ {}")
	if err != nil {
		t.Errorf("Execute error - received: %v - expected: %v", err, nil)
	}

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := env.Execute("for { }")
			if err == nil || err.Error() != ErrInterrupt.Error() {
				t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()
}

func TestRunSingleStmt(t *testing.T) {
	stmts, err := parser.ParseSrc("a = 1")
	if err != nil {
		t.Errorf("ParseSrc error - received: %v - expected: %v", err, nil)
	}

	env := NewEnv()
	value, err := RunSingleStmt(stmts[0], env)
	if err != nil {
		t.Errorf("RunSingleStmt error - received: %v - expected: %v", err, nil)
	}
	if value != int64(1) {
		t.Errorf("RunSingleStmt value - received: %v - expected: %v", value, int64(1))
	}
}

func TestCallFunctionWithVararg(t *testing.T) {
	env := NewEnv()
	err := env.Define("X", func(args ...string) []string {
		return args
	})
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	want := []string{"foo", "bar", "baz"}
	err = env.Define("a", want)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	got, err := env.Execute(`X(a...)`)
	if err != nil {
		t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute error - received %#v - expected: %#v", got, want)
	}
}

func TestAssignToInterface(t *testing.T) {
	env := NewEnv()
	X := new(struct {
		Stdout io.Writer
	})
	err := env.Define("X", X)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = env.Define("a", new(os.File))
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = env.Execute(`X.Stdout = a`)
	if err != nil {
		t.Errorf("Execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
}
