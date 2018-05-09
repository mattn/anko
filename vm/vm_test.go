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
	tests := []Test{
		{Script: `1..1`, RunError: fmt.Errorf(`strconv.ParseFloat: parsing "1..1": invalid syntax`)},
		{Script: `0x1g`, ParseError: fmt.Errorf("syntax error")},
		{Script: `9223372036854775808`, RunError: fmt.Errorf(`strconv.ParseInt: parsing "9223372036854775808": value out of range`)},

		{Script: `1`, RunOutput: int64(1)},
		{Script: `-1`, RunOutput: int64(-1)},
		{Script: `9223372036854775807`, RunOutput: int64(9223372036854775807)},
		{Script: `-9223372036854775807`, RunOutput: int64(-9223372036854775807)},
		{Script: `1.1`, RunOutput: float64(1.1)},
		{Script: `-1.1`, RunOutput: float64(-1.1)},
		{Script: `1e1`, RunOutput: float64(10)},
		{Script: `-1e1`, RunOutput: float64(-10)},
		{Script: `1e-1`, RunOutput: float64(0.1)},
		{Script: `-1e-1`, RunOutput: float64(-0.1)},
		{Script: `0x1`, RunOutput: int64(1)},
		{Script: `0xc`, RunOutput: int64(12)},
		// TOFIX:
		{Script: `0xe`, RunError: fmt.Errorf(`strconv.ParseFloat: parsing "0xe": invalid syntax`)},
		{Script: `0xf`, RunOutput: int64(15)},
		{Script: `-0x1`, RunOutput: int64(-1)},
		{Script: `-0xc`, RunOutput: int64(-12)},
		{Script: `-0xf`, RunOutput: int64(-15)},
	}
	RunTests(t, tests, nil)
}

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `a`, Input: map[string]interface{}{"a": 'a'}, RunOutput: 'a', Output: map[string]interface{}{"a": 'a'}},
		{Script: `a.b`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support member operation"), Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support index operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support slice operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},

		{Script: `a.b = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support member operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0] = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support index operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0:1] = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support slice operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},

		{Script: `a.b = "a"`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("type string does not support member operation"), Output: map[string]interface{}{"a": "test"}},
		{Script: `a[0] = "a"`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("type string does not support index operation for assignment"), Output: map[string]interface{}{"a": "test"}},
		{Script: `a[0:1] = "a"`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("type string does not support slice operation for assignment"), Output: map[string]interface{}{"a": "test"}},

		{Script: `a`, Input: map[string]interface{}{"a": "test"}, RunOutput: "test", Output: map[string]interface{}{"a": "test"}},
		{Script: `a["a"]`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"a": "test"}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": ""}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": ""}},
		{Script: `a[-1]`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test"}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": "test"}, RunOutput: "t", Output: map[string]interface{}{"a": "test"}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": "test"}, RunOutput: "e", Output: map[string]interface{}{"a": "test"}},
		{Script: `a[3]`, Input: map[string]interface{}{"a": "test"}, RunOutput: "t", Output: map[string]interface{}{"a": "test"}},
		{Script: `a[4]`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test"}},

		{Script: `a`, Input: map[string]interface{}{"a": `"a"`}, RunOutput: `"a"`, Output: map[string]interface{}{"a": `"a"`}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": `"a"`}, RunOutput: `"`, Output: map[string]interface{}{"a": `"a"`}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": `"a"`}, RunOutput: "a", Output: map[string]interface{}{"a": `"a"`}},

		{Script: `a = "\"a\""`, RunOutput: `"a"`, Output: map[string]interface{}{"a": `"a"`}},
		{Script: `a = "\"a\""; a`, RunOutput: `"a"`, Output: map[string]interface{}{"a": `"a"`}},
		{Script: `a = "\"a\""; a[0]`, RunOutput: `"`, Output: map[string]interface{}{"a": `"a"`}},
		{Script: `a = "\"a\""; a[1]`, RunOutput: "a", Output: map[string]interface{}{"a": `"a"`}},

		{Script: `a`, Input: map[string]interface{}{"a": "a\\b"}, RunOutput: "a\\b", Output: map[string]interface{}{"a": "a\\b"}},
		{Script: `a`, Input: map[string]interface{}{"a": "a\\\\b"}, RunOutput: "a\\\\b", Output: map[string]interface{}{"a": "a\\\\b"}},
		{Script: `a = "a\b"`, RunOutput: "a\b", Output: map[string]interface{}{"a": "a\b"}},
		{Script: `a = "a\\b"`, RunOutput: "a\\b", Output: map[string]interface{}{"a": "a\\b"}},

		{Script: `a[:]`, Input: map[string]interface{}{"a": "test data"}, ParseError: fmt.Errorf("syntax error"), Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[0:]`, Input: map[string]interface{}{"a": ""}, RunOutput: "", Output: map[string]interface{}{"a": ""}},
		{Script: `a[1:]`, Input: map[string]interface{}{"a": ""}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": ""}},
		{Script: `a[:0]`, Input: map[string]interface{}{"a": ""}, RunOutput: "", Output: map[string]interface{}{"a": ""}},
		{Script: `a[:1]`, Input: map[string]interface{}{"a": ""}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": ""}},
		{Script: `a[0:0]`, Input: map[string]interface{}{"a": ""}, RunOutput: "", Output: map[string]interface{}{"a": ""}},

		{Script: `a[1:0]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("invalid slice index"), Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[-1:2]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:-2]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[-1:]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:-2]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[0:0]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "t", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:2]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "te", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:3]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "tes", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:7]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test da", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:8]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test dat", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[0:10]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[1:1]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:2]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "e", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:3]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "es", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:7]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "est da", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:8]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "est dat", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "est data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:10]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[0:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "est data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[2:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "st data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[3:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "t data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[7:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "ta", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[8:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "a", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[9:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "", Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[:0]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:1]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "t", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:2]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "te", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:3]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "tes", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:7]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test da", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:8]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test dat", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:9]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[:10]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},

		{Script: `a[0:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "test data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[1:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "est data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[2:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "st data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[3:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "t data", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[7:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "ta", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[8:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "a", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[9:]`, Input: map[string]interface{}{"a": "test data"}, RunOutput: "", Output: map[string]interface{}{"a": "test data"}},
		{Script: `a[10:]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},
	}
	RunTests(t, tests, nil)
}

func TestVar(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	var parseErrFunc = func(t *testing.T, err error) {
		if !ValueEqual("syntax error", err.Error()) {
			t.Errorf("err: %s, not equal syntax error", err)
		}
	}
	tests := []Test{
		{Script: `a := 1`, ParseErrorFunc: &parseErrFunc},
		{Script: `var a = 1++`, RunError: fmt.Errorf("invalid operation")},

		{Script: `var , b = 1, 2`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: int64(2), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `var a,  = 1, 2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = 1,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = 1,,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = ,2,`, ParseError: fmt.Errorf("syntax error")},

		{Script: `var a = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a, b = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `var a, b, c = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},

		{Script: `var a = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a, b = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
	}
	RunTests(t, tests, nil)
}

func TestModule(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `module a.b { }`, ParseError: fmt.Errorf("syntax error")},
		{Script: `module a { 1++ }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `module a { }; a.b`, RunError: fmt.Errorf("invalid operation 'b'")},

		{Script: `module a { b = nil }; a.b`, RunOutput: nil},
		{Script: `module a { b = true }; a.b`, RunOutput: true},
		{Script: `module a { b = 1 }; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1.1 }; a.b`, RunOutput: float64(1.1)},
		{Script: `module a { b = "a" }; a.b`, RunOutput: "a"},
	}
	RunTests(t, tests, nil)
}

func TestNew(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `new(foo)`, RunError: fmt.Errorf("Undefined type 'foo'")},
		{Script: `new(nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("type cannot be nil for new")},

		{Script: `a = new(bool); *a`, RunOutput: false},
		{Script: `a = new(int32); *a`, RunOutput: int32(0)},
		{Script: `a = new(int64); *a`, RunOutput: int64(0)},
		{Script: `a = new(float32); *a`, RunOutput: float32(0)},
		{Script: `a = new(float64); *a`, RunOutput: float64(0)},
		{Script: `a = new(string); *a`, RunOutput: ""},
	}
	RunTests(t, tests, nil)
}

func TestMake(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `make(foo)`, RunError: fmt.Errorf("Undefined type 'foo'")},
		{Script: `make(a.b)`, Types: map[string]interface{}{"a": true}, RunError: fmt.Errorf("no namespace called: a")},
		{Script: `make(a.b)`, Types: map[string]interface{}{"b": true}, RunError: fmt.Errorf("no namespace called: a")},

		{Script: `make(nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("type cannot be nil for make")},

		{Script: `make(bool)`, RunOutput: false},
		{Script: `make(int32)`, RunOutput: int32(0)},
		{Script: `make(int64)`, RunOutput: int64(0)},
		{Script: `make(float32)`, RunOutput: float32(0)},
		{Script: `make(float64)`, RunOutput: float64(0)},
		{Script: `make(string)`, RunOutput: ""},
	}
	RunTests(t, tests, nil)
}

func TestMakeType(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `make(type a, 1++)`, RunError: fmt.Errorf("invalid operation")},

		{Script: `make(type a, true)`, RunOutput: reflect.TypeOf(true)},
		{Script: `a = make(type a, true)`, RunOutput: reflect.TypeOf(true), Output: map[string]interface{}{"a": reflect.TypeOf(true)}},
		{Script: `make(type a, true); a = make([]a)`, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}}},
		{Script: `make(type a, make([]bool))`, RunOutput: reflect.TypeOf([]bool{})},
		{Script: `make(type a, make([]bool)); a = make(a)`, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}}},
	}
	RunTests(t, tests, nil)
}

func TestLen(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `len(1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `len(true)`, RunError: fmt.Errorf("type bool does not support len operation")},

		{Script: `a = ""; len(a)`, RunOutput: int64(0)},
		{Script: `a = "test"; len(a)`, RunOutput: int64(4)},
		{Script: `a = []; len(a)`, RunOutput: int64(0)},
		{Script: `a = [nil]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [true]; len(a)`, RunOutput: int64(1)},
		{Script: `a = ["test"]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [1]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [1.1]; len(a)`, RunOutput: int64(1)},

		{Script: `a = [[]]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [[nil]]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [[true]]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [["test"]]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [[1]]; len(a)`, RunOutput: int64(1)},
		{Script: `a = [[1.1]]; len(a)`, RunOutput: int64(1)},

		{Script: `a = [[]]; len(a[0])`, RunOutput: int64(0)},
		{Script: `a = [[nil]]; len(a[0])`, RunOutput: int64(1)},
		{Script: `a = [[true]]; len(a[0])`, RunOutput: int64(1)},
		{Script: `a = [["test"]]; len(a[0])`, RunOutput: int64(1)},
		{Script: `a = [[1]]; len(a[0])`, RunOutput: int64(1)},
		{Script: `a = [[1.1]]; len(a[0])`, RunOutput: int64(1)},

		{Script: `len(a)`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(1), Output: map[string]interface{}{"a": "a"}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: int64(0), Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},
		{Script: `len(a["test"])`, Input: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}, RunOutput: int64(4), Output: map[string]interface{}{"a": map[string]interface{}{"test": "test"}}},

		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{}}, RunOutput: int64(0), Output: map[string]interface{}{"a": []interface{}{}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{nil}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{nil}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{true}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{true}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{int32(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{int32(1)}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{float32(1.1)}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": []interface{}{"a"}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{"a"}}},

		{Script: `len(a[0])`, Input: map[string]interface{}{"a": []interface{}{"test"}}, RunOutput: int64(4), Output: map[string]interface{}{"a": []interface{}{"test"}}},

		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: int64(0), Output: map[string]interface{}{"a": [][]interface{}{}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{nil}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{nil}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{nil}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{true}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},
		{Script: `len(a)`, Input: map[string]interface{}{"a": [][]interface{}{{"a"}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{"a"}}}},

		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{nil}}, RunOutput: int64(0), Output: map[string]interface{}{"a": [][]interface{}{nil}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{nil}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{true}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{int32(1)}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{float32(1.1)}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},
		{Script: `len(a[0])`, Input: map[string]interface{}{"a": [][]interface{}{{"a"}}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{"a"}}}},

		{Script: `len(a[0][0])`, Input: map[string]interface{}{"a": [][]interface{}{{"test"}}}, RunOutput: int64(4), Output: map[string]interface{}{"a": [][]interface{}{{"test"}}}},
	}
	RunTests(t, tests, nil)
}

func TestReferencingAndDereference(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		// TOFIX:
		// {Script: `a = 1; b = &a; *b = 2; *b`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
	}
	RunTests(t, tests, nil)
}

func TestMakeChan(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `make(chan foobar, 2)`, RunError: fmt.Errorf("Undefined type 'foobar'")},
		{Script: `make(chan nilT, 2)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("type cannot be nil for make chan")},
		{Script: `make(chan bool, 1++)`, RunError: fmt.Errorf("invalid operation")},

		{Script: `a = make(chan bool); b = func (c) { c <- true }; go b(a); <- a`, RunOutput: true},
	}
	RunTests(t, tests, nil)
}

func TestChan(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `a = make(chan bool, 2); 1++ <- 1`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = make(chan bool, 2); a <- 1++`, RunError: fmt.Errorf("invalid operation")},

		// TODO: move close from core into vm code, then update tests

		{Script: `1 <- 1`, RunError: fmt.Errorf("invalid operation for chan")},
		// TODO: this panics, should we capture the panic in a better way?
		// {Script: `a = make(chan bool, 2); close(a); a <- true`, Input: map[string]interface{}{"close": func(b interface{}) { reflect.ValueOf(b).Close() }}, RunError: fmt.Errorf("channel is close")},
		// TODO: add chan syntax ok
		// {Script: `a = make(chan bool, 2); a <- true; close(a); b, ok <- a; ok`, Input: map[string]interface{}{"close": func(b interface{}) { reflect.ValueOf(b).Close() }}, RunOutput: false, Output: map[string]interface{}{"b": nil}},
		{Script: `a = make(chan bool, 2); a <- true; close(a); b = false; b <- a`, Input: map[string]interface{}{"close": func(b interface{}) { reflect.ValueOf(b).Close() }}, RunOutput: true, Output: map[string]interface{}{"b": true}},
		// TOFIX: add chan syntax ok, do not return error. Also b should be nil
		{Script: `a = make(chan bool, 2); a <- true; close(a); b = false; b <- a; b <- a`, Input: map[string]interface{}{"close": func(b interface{}) { reflect.ValueOf(b).Close() }}, RunError: fmt.Errorf("failed to send to channel"), Output: map[string]interface{}{"b": true}},

		{Script: `a = make(chan bool, 2); a <- 1`, RunError: fmt.Errorf("cannot use type int64 as type bool to send to chan")},

		{Script: `a = make(chan interface, 2); a <- nil; <- a`, RunOutput: nil},
		{Script: `a = make(chan bool, 2); a <- true; <- a`, RunOutput: true},
		{Script: `a = make(chan int32, 2); a <- 1; <- a`, RunOutput: int32(1)},
		{Script: `a = make(chan int64, 2); a <- 1; <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan float32, 2); a <- 1.1; <- a`, RunOutput: float32(1.1)},
		{Script: `a = make(chan float64, 2); a <- 1.1; <- a`, RunOutput: float64(1.1)},

		{Script: `a <- nil; <- a`, Input: map[string]interface{}{"a": make(chan interface{}, 2)}, RunOutput: nil},
		{Script: `a <- true; <- a`, Input: map[string]interface{}{"a": make(chan bool, 2)}, RunOutput: true},
		{Script: `a <- 1; <- a`, Input: map[string]interface{}{"a": make(chan int32, 2)}, RunOutput: int32(1)},
		{Script: `a <- 1; <- a`, Input: map[string]interface{}{"a": make(chan int64, 2)}, RunOutput: int64(1)},
		{Script: `a <- 1.1; <- a`, Input: map[string]interface{}{"a": make(chan float32, 2)}, RunOutput: float32(1.1)},
		{Script: `a <- 1.1; <- a`, Input: map[string]interface{}{"a": make(chan float64, 2)}, RunOutput: float64(1.1)},
		{Script: `a <- "b"; <- a`, Input: map[string]interface{}{"a": make(chan string, 2)}, RunOutput: "b"},

		{Script: `a = make(chan bool, 2); a <- true; a <- <- a`, RunOutput: nil},
		{Script: `a = make(chan bool, 2); a <- true; a <- (<- a)`, RunOutput: nil},
		{Script: `a = make(chan bool, 2); a <- true; a <- <- a; <- a`, RunOutput: true},
		{Script: `a = make(chan bool, 2); a <- true; b = false; b <- a`, RunOutput: true, Output: map[string]interface{}{"b": true}},
		// TOFIX: if variable is not created yet, should make variable instead of error
		// {Script: `a = make(chan bool, 2); a <- true; b <- a`, RunOutput: true, Output: map[string]interface{}{"b": true}},
	}
	RunTests(t, tests, nil)
}

func TestVMDelete(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `delete(1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `delete("a", 1++)`, RunError: fmt.Errorf("invalid operation")},

		{Script: `a = 1; delete("a"); a`, RunError: fmt.Errorf("undefined symbol 'a'")},

		{Script: `delete("a")`},
		{Script: `delete("a", false)`},
		{Script: `delete("a", true)`},
		{Script: `delete("a", nil)`},

		{Script: `a = 1; func b() { delete("a") }; b()`, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; func b() { delete("a", false) }; b()`, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; func b() { delete("a", true) }; b(); a`, RunError: fmt.Errorf("undefined symbol 'a'")},
	}
	RunTests(t, tests, nil)
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
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
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
				t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()

	_, err := env.Execute("for { }")
	if err == nil || err.Error() != ErrInterrupt.Error() {
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}

	ClearInterrupt(env)

	_, err = env.Execute("for i = 0; i < 1000; i ++ {}")
	if err != nil {
		t.Errorf("execute error - received: %v - expected: %v", err, nil)
	}

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := env.Execute("for { }")
			if err == nil || err.Error() != ErrInterrupt.Error() {
				t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	Interrupt(env)
	waitGroup.Wait()
}

func TestRunSingleStmt(t *testing.T) {
	stmts, err := parser.ParseSrc(`a = 1`)
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

	stmts, err = parser.ParseSrc(`return a`)
	if err != nil {
		t.Errorf("ParseSrc error - received: %v - expected: %v", err, nil)
	}

	env = NewEnv()
	err = env.defineValue("a", reflect.Value{})
	if err != nil {
		t.Errorf("defineValue error - received: %v - expected: %v", err, nil)
	}

	value, err = RunSingleStmt(stmts[0], env)
	if err != nil {
		t.Errorf("RunSingleStmt error - received: %v - expected: %v", err, nil)
	}
	if value != nil {
		t.Errorf("RunSingleStmt value - received: %v - expected: %v", value, nil)
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
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("execute error - received %#v - expected: %#v", got, want)
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
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
}
