package vm

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/mattn/anko/ast"
	"github.com/mattn/anko/ast/astutil"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
)

func TestNumbers(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: ``},
		{Script: `;`},
		{Script: `
`},
		{Script: `
1
`, RunOutput: int64(1)},

		{Script: `1..1`, ParseError: fmt.Errorf("invalid number: 1..1")},
		{Script: `1e.1`, ParseError: fmt.Errorf("invalid number: 1e.1")},
		{Script: `1ee1`, ParseError: fmt.Errorf("syntax error")},
		{Script: `1e+e1`, ParseError: fmt.Errorf("syntax error")},
		{Script: `0x1g`, ParseError: fmt.Errorf("syntax error")},
		{Script: `0b2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `9223372036854775808`, ParseError: fmt.Errorf("invalid number: 9223372036854775808")},
		{Script: `-9223372036854775809`, ParseError: fmt.Errorf("invalid number: -9223372036854775809")},

		{Script: `1`, RunOutput: int64(1)},
		{Script: `-1`, RunOutput: int64(-1)},
		{Script: `9223372036854775807`, RunOutput: int64(9223372036854775807)},
		{Script: `-9223372036854775808`, RunOutput: int64(-9223372036854775808)},
		{Script: `-9223372036854775807-1`, RunOutput: int64(-9223372036854775808)},
		{Script: `-9223372036854775807 -1`, RunOutput: int64(-9223372036854775808)},
		{Script: `-9223372036854775807 - 1`, RunOutput: int64(-9223372036854775808)},
		{Script: `1.1`, RunOutput: float64(1.1)},
		{Script: `-1.1`, RunOutput: float64(-1.1)},

		{Script: `1e1`, RunOutput: float64(10)},
		{Script: `1.5e1`, RunOutput: float64(15)},
		{Script: `1e-1`, RunOutput: float64(0.1)},

		{Script: `-1e1`, RunOutput: float64(-10)},
		{Script: `-1.5e1`, RunOutput: float64(-15)},
		{Script: `-1e-1`, RunOutput: float64(-0.1)},

		{Script: `0x1`, RunOutput: int64(1)},
		{Script: `0xa`, RunOutput: int64(10)},
		{Script: `0xb`, RunOutput: int64(11)},
		{Script: `0xc`, RunOutput: int64(12)},
		{Script: `0xe`, RunOutput: int64(14)},
		{Script: `0xf`, RunOutput: int64(15)},
		{Script: `0Xf`, RunOutput: int64(15)},
		{Script: `0XF`, RunOutput: int64(15)},
		{Script: `0x7FFFFFFFFFFFFFFF`, RunOutput: int64(9223372036854775807)},

		{Script: `-0x1`, RunOutput: int64(-1)},
		{Script: `-0xc`, RunOutput: int64(-12)},
		{Script: `-0xe`, RunOutput: int64(-14)},
		{Script: `-0xf`, RunOutput: int64(-15)},
		{Script: `-0Xf`, RunOutput: int64(-15)},
		{Script: `-0x7FFFFFFFFFFFFFFF`, RunOutput: int64(-9223372036854775807)},

		{Script: `0b1`, RunOutput: int64(1)},
		{Script: `0b001100`, RunOutput: int64(12)},
		{Script: `0b111111111111111111111111111111111111111111111111111111111111111`, RunOutput: int64(9223372036854775807)},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestStrings(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a`, Input: map[string]interface{}{"a": 'a'}, RunOutput: 'a', Output: map[string]interface{}{"a": 'a'}},
		{Script: `a.b`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support member operation"), Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support index operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support slice operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},

		{Script: `a.b = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support member operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0] = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support index operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},
		{Script: `a[0:1] = "a"`, Input: map[string]interface{}{"a": 'a'}, RunError: fmt.Errorf("type int32 does not support slice operation"), RunOutput: nil, Output: map[string]interface{}{"a": 'a'}},

		{Script: `a.b = "a"`, Input: map[string]interface{}{"a": "test"}, RunError: fmt.Errorf("type string does not support member operation"), Output: map[string]interface{}{"a": "test"}},
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

		{Script: `a[1:0]`, Input: map[string]interface{}{"a": "test data"}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "test data"}},
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

		// index assignment - len 0
		{Script: `a = ""; a[0] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "x"}},
		{Script: `a = ""; a[1] = "x"`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": ""}},

		// index assignment - len 1
		{Script: `a = "a"; a[0] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "x"}},
		{Script: `a = "a"; a[1] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "ax"}},
		{Script: `a = "a"; a[2] = "x"`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "a"}},

		// index assignment - len 2
		{Script: `a = "ab"; a[0] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "xb"}},
		{Script: `a = "ab"; a[1] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "ax"}},
		{Script: `a = "ab"; a[2] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "abx"}},
		{Script: `a = "ab"; a[3] = "x"`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "ab"}},

		// index assignment - len 3
		{Script: `a = "abc"; a[0] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "xbc"}},
		{Script: `a = "abc"; a[1] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[2] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "abx"}},
		{Script: `a = "abc"; a[3] = "x"`, RunOutput: "x", Output: map[string]interface{}{"a": "abcx"}},
		{Script: `a = "abc"; a[4] = "x"`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": "abc"}},

		// index assignment - vm types
		{Script: `a = "abc"; a[1] = nil`, RunOutput: nil, Output: map[string]interface{}{"a": "ac"}},
		{Script: `a = "abc"; a[1] = true`, RunError: fmt.Errorf("type bool cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = 120`, RunOutput: int64(120), Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[1] = 2.2`, RunError: fmt.Errorf("type float64 cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = ["a"]`, RunError: fmt.Errorf("type []interface {} cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},

		// index assignment - Go types
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": reflect.Value{}}, RunError: fmt.Errorf("type reflect.Value cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": "ac"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": true}, RunError: fmt.Errorf("type bool cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": int32(120)}, RunOutput: int32(120), Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": int64(120)}, RunOutput: int64(120), Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": float32(1.1)}, RunError: fmt.Errorf("type float32 cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": float64(2.2)}, RunError: fmt.Errorf("type float64 cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": "x"}, RunOutput: "x", Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": 'x'}, RunOutput: 'x', Output: map[string]interface{}{"a": "axc"}},
		{Script: `a = "abc"; a[1] = b`, Input: map[string]interface{}{"b": struct{}{}}, RunError: fmt.Errorf("type struct {} cannot be assigned to type string"), Output: map[string]interface{}{"a": "abc"}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestVar(t *testing.T) {
	t.Parallel()

	testInput1 := map[string]interface{}{"b": func() {}}
	tests := []Test{
		// simple one variable
		{Script: `1 = 2`, RunError: fmt.Errorf("invalid operation")},
		{Script: `var 1 = 2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a = 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `var a = 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a := 1`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a := 1`, ParseError: fmt.Errorf("syntax error")},
		{Script: `y = z`, RunError: fmt.Errorf("undefined symbol 'z'")},

		{Script: `a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `a = true`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `a = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1.1`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `a = "a"`, RunOutput: "a", Output: map[string]interface{}{"a": "a"}},

		{Script: `var a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `var a = true`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `var a = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a = 1.1`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `var a = "a"`, RunOutput: "a", Output: map[string]interface{}{"a": "a"}},

		{Script: `a = b`, Input: map[string]interface{}{"b": reflect.Value{}}, RunOutput: reflect.Value{}, Output: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{Script: `a = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a = b`, Input: map[string]interface{}{"b": true}, RunOutput: true, Output: map[string]interface{}{"a": true, "b": true}},
		{Script: `a = b`, Input: map[string]interface{}{"b": int32(1)}, RunOutput: int32(1), Output: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": float32(1.1)}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": float64(1.1)}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": "a"}, RunOutput: "a", Output: map[string]interface{}{"a": "a", "b": "a"}},
		{Script: `a = b`, Input: map[string]interface{}{"b": 'a'}, RunOutput: 'a', Output: map[string]interface{}{"a": 'a', "b": 'a'}},
		{Script: `a = b`, Input: map[string]interface{}{"b": struct{}{}}, RunOutput: struct{}{}, Output: map[string]interface{}{"a": struct{}{}, "b": struct{}{}}},

		{Script: `var a = b`, Input: map[string]interface{}{"b": reflect.Value{}}, RunOutput: reflect.Value{}, Output: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": true}, RunOutput: true, Output: map[string]interface{}{"a": true, "b": true}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": int32(1)}, RunOutput: int32(1), Output: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": float32(1.1)}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": float64(1.1)}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": "a"}, RunOutput: "a", Output: map[string]interface{}{"a": "a", "b": "a"}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": 'a'}, RunOutput: 'a', Output: map[string]interface{}{"a": 'a', "b": 'a'}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": struct{}{}}, RunOutput: struct{}{}, Output: map[string]interface{}{"a": struct{}{}, "b": struct{}{}}},

		// simple one variable overwrite
		{Script: `a = true; a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `a = nil; a = true`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `a = 1; a = 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1.1; a = 2.2`, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `a = "a"; a = "b"`, RunOutput: "b", Output: map[string]interface{}{"a": "b"}},

		{Script: `var a = true; var a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `var a = nil; var a = true`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `var a = 1; var a = 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `var a = 1.1; var a = 2.2`, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `var a = "a"; var a = "b"`, RunOutput: "b", Output: map[string]interface{}{"a": "b"}},

		{Script: `a = nil`, Input: map[string]interface{}{"a": true}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `a = true`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `a = 2`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 2`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 2.2`, Input: map[string]interface{}{"a": float32(1.1)}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `a = 2.2`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `a = "b"`, Input: map[string]interface{}{"a": "a"}, RunOutput: "b", Output: map[string]interface{}{"a": "b"}},

		{Script: `var a = nil`, Input: map[string]interface{}{"a": true}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `var a = true`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `var a = 2`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `var a = 2`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `var a = 2.2`, Input: map[string]interface{}{"a": float32(1.1)}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `var a = 2.2`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": float64(2.2)}},
		{Script: `var a = "b"`, Input: map[string]interface{}{"a": "a"}, RunOutput: "b", Output: map[string]interface{}{"a": "b"}},

		// Go variable copy
		{Script: `a = b`, Input: testInput1, RunOutput: testInput1["b"], Output: map[string]interface{}{"a": testInput1["b"], "b": testInput1["b"]}},
		{Script: `var a = b`, Input: testInput1, RunOutput: testInput1["b"], Output: map[string]interface{}{"a": testInput1["b"], "b": testInput1["b"]}},

		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValue}, RunOutput: testVarValue, Output: map[string]interface{}{"a": testVarValue, "b": testVarValue}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarBoolP}, RunOutput: testVarBoolP, Output: map[string]interface{}{"a": testVarBoolP, "b": testVarBoolP}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarInt32P}, RunOutput: testVarInt32P, Output: map[string]interface{}{"a": testVarInt32P, "b": testVarInt32P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarInt64P}, RunOutput: testVarInt64P, Output: map[string]interface{}{"a": testVarInt64P, "b": testVarInt64P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFloat32P}, RunOutput: testVarFloat32P, Output: map[string]interface{}{"a": testVarFloat32P, "b": testVarFloat32P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFloat64P}, RunOutput: testVarFloat64P, Output: map[string]interface{}{"a": testVarFloat64P, "b": testVarFloat64P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarStringP}, RunOutput: testVarStringP, Output: map[string]interface{}{"a": testVarStringP, "b": testVarStringP}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFuncP}, RunOutput: testVarFuncP, Output: map[string]interface{}{"a": testVarFuncP, "b": testVarFuncP}},

		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValue}, RunOutput: testVarValue, Output: map[string]interface{}{"a": testVarValue, "b": testVarValue}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarBoolP}, RunOutput: testVarBoolP, Output: map[string]interface{}{"a": testVarBoolP, "b": testVarBoolP}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarInt32P}, RunOutput: testVarInt32P, Output: map[string]interface{}{"a": testVarInt32P, "b": testVarInt32P}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarInt64P}, RunOutput: testVarInt64P, Output: map[string]interface{}{"a": testVarInt64P, "b": testVarInt64P}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarFloat32P}, RunOutput: testVarFloat32P, Output: map[string]interface{}{"a": testVarFloat32P, "b": testVarFloat32P}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarFloat64P}, RunOutput: testVarFloat64P, Output: map[string]interface{}{"a": testVarFloat64P, "b": testVarFloat64P}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarStringP}, RunOutput: testVarStringP, Output: map[string]interface{}{"a": testVarStringP, "b": testVarStringP}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarFuncP}, RunOutput: testVarFuncP, Output: map[string]interface{}{"a": testVarFuncP, "b": testVarFuncP}},

		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueBool}, RunOutput: testVarValueBool, Output: map[string]interface{}{"a": testVarValueBool, "b": testVarValueBool}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueInt32}, RunOutput: testVarValueInt32, Output: map[string]interface{}{"a": testVarValueInt32, "b": testVarValueInt32}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueInt64}, RunOutput: testVarValueInt64, Output: map[string]interface{}{"a": testVarValueInt64, "b": testVarValueInt64}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueFloat32}, RunOutput: testVarValueFloat32, Output: map[string]interface{}{"a": testVarValueFloat32, "b": testVarValueFloat32}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueFloat64}, RunOutput: testVarValueFloat64, Output: map[string]interface{}{"a": testVarValueFloat64, "b": testVarValueFloat64}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueString}, RunOutput: testVarValueString, Output: map[string]interface{}{"a": testVarValueString, "b": testVarValueString}},

		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueBool}, RunOutput: testVarValueBool, Output: map[string]interface{}{"a": testVarValueBool, "b": testVarValueBool}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueInt32}, RunOutput: testVarValueInt32, Output: map[string]interface{}{"a": testVarValueInt32, "b": testVarValueInt32}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueInt64}, RunOutput: testVarValueInt64, Output: map[string]interface{}{"a": testVarValueInt64, "b": testVarValueInt64}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueFloat32}, RunOutput: testVarValueFloat32, Output: map[string]interface{}{"a": testVarValueFloat32, "b": testVarValueFloat32}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueFloat64}, RunOutput: testVarValueFloat64, Output: map[string]interface{}{"a": testVarValueFloat64, "b": testVarValueFloat64}},
		{Script: `var a = b`, Input: map[string]interface{}{"b": testVarValueString}, RunOutput: testVarValueString, Output: map[string]interface{}{"a": testVarValueString, "b": testVarValueString}},

		// one variable spacing
		{Script: `
a  =  1
`, RunOutput: int64(1)},
		{Script: `
a  =  1;
`, RunOutput: int64(1)},

		// one variable many values
		{Script: `, b = 1, 2`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: int64(2), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `var , b = 1, 2`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: int64(2), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `a,  = 1, 2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a,  = 1, 2`, ParseError: fmt.Errorf("syntax error")},

		// TOFIX: should not error?
		{Script: `a = 1, 2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1)}},
		// TOFIX: should not error?
		{Script: `a = 1, 2, 3`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1)}},

		// two variables many values
		{Script: `a, b  = 1,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = 1,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a, b  = ,1`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: int64(1)},
		{Script: `var a, b  = ,1`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: int64(1)},
		{Script: `a, b  = 1,,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = 1,,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a, b  = ,1,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = ,1,`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a, b  = ,,1`, ParseError: fmt.Errorf("syntax error")},
		{Script: `var a, b  = ,,1`, ParseError: fmt.Errorf("syntax error")},

		{Script: `a.c, b = 1, 2`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a, b.c = 1, 2`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `a, b = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a, b = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a, b = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `var a, b = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `a, b = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `var a, b = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},

		// three variables many values
		{Script: `a, b, c = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a, b, c = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a, b, c = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `var a, b, c = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `a, b, c = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{Script: `var a, b, c = 1, 2, 3`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{Script: `a, b, c = 1, 2, 3, 4`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{Script: `var a, b, c = 1, 2, 3, 4`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},

		// scope
		{Script: `func(){ a = 1 }(); a`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `func(){ var a = 1 }(); a`, RunError: fmt.Errorf("undefined symbol 'a'")},

		{Script: `a = 1; func(){ a = 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `var a = 1; func(){ a = 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; func(){ var a = 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a = 1; func(){ var a = 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1)}},

		// function return
		{Script: `a, 1++ = func(){ return 1, 2 }()`, RunError: fmt.Errorf("invalid operation"), Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a = func(){ return 1 }()`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `var a = func(){ return 1 }()`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a, b = func(){ return 1, 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `var a, b = func(){ return 1, 2 }()`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestModule(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `module a.b { }`, ParseError: fmt.Errorf("syntax error")},
		{Script: `module a { 1++ }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `module a { }; a.b`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `module a { b = 1 }`, RunOutput: nil},

		{Script: `module a { b = nil }; a.b`, RunOutput: nil},
		{Script: `module a { b = true }; a.b`, RunOutput: true},
		{Script: `module a { b = 1 }; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1.5 }; a.b`, RunOutput: float64(1.5)},
		{Script: `module a { b = "a" }; a.b`, RunOutput: "a"},
		{Script: `module a { func b() { return "b"} }; a.b()`, RunOutput: "b"},

		{Script: `module a { _b = "b"; func c() { return _b} }`, RunOutput: nil},
		{Script: `module a { _b = "b"; func c() { return _b} }; a.c()`, RunOutput: "b"},

		{Script: `module a { b = 1 }; var c = a; c.b = 2; c.b`, RunOutput: int64(2)},

		// test module copy
		{Script: `module a { b = 1 }; c = a; d = a; a.b = 2; a.b`, RunOutput: int64(2)},
		{Script: `module a { b = 1 }; c = a; d = a; a.b = 2; c.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; a.b = 2; d.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; c.b = 2; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; c.b = 2; c.b`, RunOutput: int64(2)},
		{Script: `module a { b = 1 }; c = a; d = a; c.b = 2; d.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; d.b = 2; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; d.b = 2; c.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; c = a; d = a; d.b = 2; d.b`, RunOutput: int64(2)},

		{Script: `module a { b = 1 }; var c = a; var d = a; a.b = 2; a.b`, RunOutput: int64(2)},
		{Script: `module a { b = 1 }; var c = a; var d = a; a.b = 2; c.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; a.b = 2; d.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; c.b = 2; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; c.b = 2; c.b`, RunOutput: int64(2)},
		{Script: `module a { b = 1 }; var c = a; var d = a; c.b = 2; d.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; d.b = 2; a.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; d.b = 2; c.b`, RunOutput: int64(1)},
		{Script: `module a { b = 1 }; var c = a; var d = a; d.b = 2; d.b`, RunOutput: int64(2)},

		// test type scope
		{Script: `module b { make(type Duration, a) }; func c() { d = new(time.Duration); return *d }; c()`, Input: map[string]interface{}{"a": time.Duration(0)}, RunError: fmt.Errorf("no namespace called: time")},
		{Script: `module time { make(type Duration, a) }; func c() { d = new(time.Duration); return *d }; c()`, Input: map[string]interface{}{"a": time.Duration(0)}, RunOutput: time.Duration(0)},
		{Script: `module x { module time { make(type Duration, a) } }; func c() { d = new(x.time.Duration); return *d }; c()`, Input: map[string]interface{}{"a": time.Duration(0)}, RunOutput: time.Duration(0)},
		{Script: `module x { module time { make(type Duration, a) } }; func c() { d = new(y.time.Duration); return *d }; c()`, Input: map[string]interface{}{"a": time.Duration(0)}, RunError: fmt.Errorf("no namespace called: y")},
		{Script: `module x { module time { make(type Duration, a) } }; func c() { d = new(x.y.Duration); return *d }; c()`, Input: map[string]interface{}{"a": time.Duration(0)}, RunError: fmt.Errorf("no namespace called: y")},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `new(foo)`, RunError: fmt.Errorf("undefined type 'foo'")},
		{Script: `new(nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},

		// default
		{Script: `a = new(bool); *a`, RunOutput: false},
		{Script: `a = new(int32); *a`, RunOutput: int32(0)},
		{Script: `a = new(int64); *a`, RunOutput: int64(0)},
		{Script: `a = new(float32); *a`, RunOutput: float32(0)},
		{Script: `a = new(float64); *a`, RunOutput: float64(0)},
		{Script: `a = new(string); *a`, RunOutput: ""},

		// ptr
		{Script: `a = new(*string); b = *a; *b`, RunOutput: ""},
		{Script: `a = new(*string); **a`, RunOutput: ""},

		// slice
		{Script: `a = new([]int64); *a`, RunOutput: []int64{}},

		// map
		{Script: `a = new(map[string]int64); *a`, RunOutput: map[string]int64{}},

		// chan
		{Script: `a = new(chan int64); go func(){ (*a) <- 1 }(); <- *a`, RunOutput: int64(1)},
		{Script: `a = new(chan int64); go func(){ *a <- 1 }(); <- *a`, RunOutput: int64(1)},

		// struct
		{Script: `a = new(struct{ A int64 }); *a`, RunOutput: struct{ A int64 }{}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMake(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `make(map[[]string]int64)`, RunError: fmt.Errorf("reflect.MapOf: invalid key type []string")},
	}
	runTests(t, tests, nil, &Options{Debug: false})

	tests = []Test{
		{Script: `make(struct {})`, ParseError: fmt.Errorf("syntax error")},
		{Script: `make(struct { , })`, ParseError: fmt.Errorf("syntax error")},
		{Script: `make(struct { A map })`, ParseError: fmt.Errorf("syntax error")},
		{Script: `make(struct { , A int64})`, ParseError: fmt.Errorf("syntax error")},
		{Script: `make(struct { A int64, })`, ParseError: fmt.Errorf("syntax error")},

		{Script: `make(foo)`, RunError: fmt.Errorf("undefined type 'foo'")},
		{Script: `make(a.b)`, Types: map[string]interface{}{"a": true}, RunError: fmt.Errorf("no namespace called: a")},
		{Script: `make(a.b)`, Types: map[string]interface{}{"b": true}, RunError: fmt.Errorf("no namespace called: a")},
		{Script: `make([]int64, 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `make([]int64, 1, 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `make([]int64, 2, 1)`, RunError: fmt.Errorf("make slice len > cap")},
		{Script: `make(map[foo]int64)`, RunError: fmt.Errorf("undefined type 'foo'")},
		{Script: `make(map[int64]foo)`, RunError: fmt.Errorf("undefined type 'foo'")},
		{Script: `make(chan foo)`, RunError: fmt.Errorf("undefined type 'foo'")},
		{Script: `make(chan int64, 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `make(struct { A foo })`, RunError: fmt.Errorf("undefined type 'foo'")},

		// nill type
		{Script: `make(nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make(*nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make([]nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make(map[nilT]string)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make(map[string]nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make(chan nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make(struct { A nilT })`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},

		// default
		{Script: `make(bool)`, RunOutput: false},
		{Script: `make(int32)`, RunOutput: int32(0)},
		{Script: `make(int64)`, RunOutput: int64(0)},
		{Script: `make(float32)`, RunOutput: float32(0)},
		{Script: `make(float64)`, RunOutput: float64(0)},
		{Script: `make(string)`, RunOutput: ""},

		// ptr
		{Script: `a = make(*int64); *a`, RunOutput: int64(0)},
		{Script: `a = make(**int64); **a`, RunOutput: int64(0)},
		{Script: `a = make(***int64); ***a`, RunOutput: int64(0)},
		{Script: `a = make(*[]int64); *a`, RunOutput: []int64{}},
		{Script: `a = make(*map[string]int64); *a`, RunOutput: map[string]int64{}},
		{Script: `a = make(*chan int64); go func(){ (*a) <- 1 }(); <- *a`, RunOutput: int64(1)},
		{Script: `a = make(*chan int64); go func(){ *a <- 1 }(); <- *a`, RunOutput: int64(1)},

		// slice
		{Script: `make([]int64)`, RunOutput: []int64{}},
		{Script: `a = make([]int64, 1); a[0]`, RunOutput: int64(0)},
		{Script: `a = make([]int64, 1, 2); a[0]`, RunOutput: int64(0)},
		{Script: `make([]*int64)`, RunOutput: []*int64{}},
		{Script: `make([][]int64)`, RunOutput: [][]int64{}},
		{Script: `make([]map[string]int64)`, RunOutput: []map[string]int64{}},

		// map
		{Script: `make(map[string]int64)`, RunOutput: map[string]int64{}},
		{Script: `make(map[string]*int64)`, RunOutput: map[string]*int64{}},
		{Script: `make(map[*string]int64)`, RunOutput: map[*string]int64{}},
		{Script: `make(map[*string]*int64)`, RunOutput: map[*string]*int64{}},
		{Script: `make(map[string][]int64)`, RunOutput: map[string][]int64{}},
		{Script: `make(map[string]chan int64)`, RunOutput: map[string]chan int64{}},
		{Script: `make(map[chan string]int64)`, RunOutput: map[chan string]int64{}},

		// chan
		{Script: `a = make(chan int64); go func(){ a <- 1 }(); <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 1); a <- 1; <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan *int64, 1); b = 1; a <- &b; c = <- a; *c`, RunOutput: int64(1)},
		{Script: `a = make(chan []int64, 1); a <- [1]; <- a`, RunOutput: []int64{1}},
		{Script: `a = make(chan map[string]int64, 1); b = make(map[string]int64); a <- b; <- a`, RunOutput: map[string]int64{}},
		{Script: `a = make(chan int64, 1); b = &a; *b <- 1; <- *b`, RunOutput: int64(1)},

		// struct
		{Script: `make(struct { A int64 })`, RunOutput: struct{ A int64 }{}},
		{Script: `make(struct { A *int64 })`, RunOutput: struct{ A *int64 }{}},
		{Script: `make(struct { A []int64 })`, RunOutput: struct{ A []int64 }{A: []int64{}}},
		{Script: `make(struct { A map[string]int64 })`, RunOutput: struct{ A map[string]int64 }{A: map[string]int64{}}},
		{Script: `a = make(struct { A chan int64 }); go func(){ a.A <- 1 }(); <- a.A`, RunOutput: int64(1)},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeType(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `make(type a, 1++)`, RunError: fmt.Errorf("invalid operation")},

		{Script: `make(type a, true)`, RunOutput: reflect.TypeOf(true)},
		{Script: `a = make(type a, true)`, RunOutput: reflect.TypeOf(true), Output: map[string]interface{}{"a": reflect.TypeOf(true)}},
		{Script: `make(type a, true); a = make([]a)`, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}}},
		{Script: `make(type a, make([]bool))`, RunOutput: reflect.TypeOf([]bool{})},
		{Script: `make(type a, make([]bool)); a = make(a)`, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestReferencingAndDereference(t *testing.T) {
	t.Parallel()

	tests := []Test{
		// TOFIX:
		// {Script: `a = 1; b = &a; *b = 2; *b`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestChan(t *testing.T) {
	t.Parallel()

	tests := []Test{
		// send on closed channel
		{Script: `a = make(chan int64, 2); close(a); a <- 1`, RunError: fmt.Errorf("send on closed channel")},
	}
	runTests(t, tests, nil, &Options{Debug: false})

	tests = []Test{
		{Script: `a = make(chan int64, 2); a <- 1; = <- a`, ParseError: fmt.Errorf("missing expressions on left side of channel operator"), RunError: fmt.Errorf("invalid operation")},

		{Script: `<- 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `1++ <- 1`, RunError: fmt.Errorf("invalid operation")},
		{Script: `1 <- 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = make(chan int64, 2); a <- 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `<- 1`, RunError: fmt.Errorf("receive from non-chan type int64")},
		{Script: `1 <- 1`, RunError: fmt.Errorf("send to non-chan type int64")},
		{Script: `a = make(chan int64, 2); 1 <- a`, RunError: fmt.Errorf("send to non-chan type int64")},
		{Script: `a = make(chan bool, 2); a <- 1`, RunError: fmt.Errorf("cannot use type int64 as type bool to send to chan")},
		{Script: `close(1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `close(1)`, RunError: fmt.Errorf("type cannot be int64 for close")},

		// let channel errors
		{Script: `a = <- c`, RunError: fmt.Errorf("undefined symbol 'c'")},
		{Script: `a, b = <- c`, RunError: fmt.Errorf("undefined symbol 'c'")},
		{Script: `a = <- 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a, b = <- 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `c = 1; a = <- c`, RunError: fmt.Errorf("receive from non-chan type int64")},
		{Script: `c = 1; a, b = <- c`, RunError: fmt.Errorf("receive from non-chan type int64")},
		{Script: `a = make(chan int64, 2); a <- 1; 1++ = <- a`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = make(chan int64, 2); a <- 1; 1++, b = <- a`, RunError: fmt.Errorf("invalid operation")},

		// send to channel
		{Script: `a <- nil`, Input: map[string]interface{}{"a": make(chan interface{}, 2)}},
		{Script: `a <- true`, Input: map[string]interface{}{"a": make(chan bool, 2)}},
		{Script: `a <- 1`, Input: map[string]interface{}{"a": make(chan int32, 2)}},
		{Script: `a <- 2`, Input: map[string]interface{}{"a": make(chan int64, 2)}},
		{Script: `a <- 1.5`, Input: map[string]interface{}{"a": make(chan float32, 2)}},
		{Script: `a <- 2.5`, Input: map[string]interface{}{"a": make(chan float64, 2)}},
		{Script: `a <- "b"`, Input: map[string]interface{}{"a": make(chan string, 2)}},

		{Script: `a = make(chan interface, 2); a <- nil`},
		{Script: `a = make(chan bool, 2); a <- true`},
		{Script: `a = make(chan int32, 2); a <- 1`},
		{Script: `a = make(chan int64, 2); a <- 2`},
		{Script: `a = make(chan float32, 2); a <- 1.5`},
		{Script: `a = make(chan float64, 2); a <- 2.5`},
		{Script: `a = make(chan string, 2); a <- "b"`},

		// send to channel then receive from channel
		{Script: `a <- nil; <- a`, Input: map[string]interface{}{"a": make(chan interface{}, 2)}, RunOutput: nil},
		{Script: `a <- true; <- a`, Input: map[string]interface{}{"a": make(chan bool, 2)}, RunOutput: true},
		{Script: `a <- 1; <- a`, Input: map[string]interface{}{"a": make(chan int32, 2)}, RunOutput: int32(1)},
		{Script: `a <- 2; <- a`, Input: map[string]interface{}{"a": make(chan int64, 2)}, RunOutput: int64(2)},
		{Script: `a <- 1.5; <- a`, Input: map[string]interface{}{"a": make(chan float32, 2)}, RunOutput: float32(1.5)},
		{Script: `a <- 2.5; <- a`, Input: map[string]interface{}{"a": make(chan float64, 2)}, RunOutput: float64(2.5)},
		{Script: `a <- "b"; <- a`, Input: map[string]interface{}{"a": make(chan string, 2)}, RunOutput: "b"},

		{Script: `a = make(chan interface, 2); a <- nil; <- a`, RunOutput: nil},
		{Script: `a = make(chan bool, 2); a <- true; <- a`, RunOutput: true},
		{Script: `a = make(chan int32, 2); a <- 1; <- a`, RunOutput: int32(1)},
		{Script: `a = make(chan int64, 2); a <- 2; <- a`, RunOutput: int64(2)},
		{Script: `a = make(chan float32, 2); a <- 1.5; <- a`, RunOutput: float32(1.5)},
		{Script: `a = make(chan float64, 2); a <- 2.5; <- a`, RunOutput: float64(2.5)},
		{Script: `a = make(chan string, 2); a <- "b"; <- a`, RunOutput: "b"},

		// send to channel, receive from channel, then assign to variable
		{Script: `a <- nil; b = <- a`, Input: map[string]interface{}{"a": make(chan interface{}, 2)}, RunOutput: nil, Output: map[string]interface{}{"b": nil}},
		{Script: `a <- true; b = <- a`, Input: map[string]interface{}{"a": make(chan bool, 2)}, RunOutput: true, Output: map[string]interface{}{"b": true}},
		{Script: `a <- 1; b = <- a`, Input: map[string]interface{}{"a": make(chan int32, 2)}, RunOutput: int32(1), Output: map[string]interface{}{"b": int32(1)}},
		{Script: `a <- 2; b = <- a`, Input: map[string]interface{}{"a": make(chan int64, 2)}, RunOutput: int64(2), Output: map[string]interface{}{"b": int64(2)}},
		{Script: `a <- 1.5; b = <- a`, Input: map[string]interface{}{"a": make(chan float32, 2)}, RunOutput: float32(1.5), Output: map[string]interface{}{"b": float32(1.5)}},
		{Script: `a <- 2.5; b = <- a`, Input: map[string]interface{}{"a": make(chan float64, 2)}, RunOutput: float64(2.5), Output: map[string]interface{}{"b": float64(2.5)}},
		{Script: `a <- "b"; b = <- a`, Input: map[string]interface{}{"a": make(chan string, 2)}, RunOutput: "b", Output: map[string]interface{}{"b": "b"}},

		{Script: `a = make(chan interface, 2); a <- nil; b = <- a`, RunOutput: nil, Output: map[string]interface{}{"b": nil}},
		{Script: `a = make(chan bool, 2); a <- true; b = <- a`, RunOutput: true, Output: map[string]interface{}{"b": true}},
		{Script: `a = make(chan int32, 2); a <- 1; b = <- a`, RunOutput: int32(1), Output: map[string]interface{}{"b": int32(1)}},
		{Script: `a = make(chan int64, 2); a <- 2; b = <- a`, RunOutput: int64(2), Output: map[string]interface{}{"b": int64(2)}},
		{Script: `a = make(chan float32, 2); a <- 1.5; b = <- a`, RunOutput: float32(1.5), Output: map[string]interface{}{"b": float32(1.5)}},
		{Script: `a = make(chan float64, 2); a <- 2.5; b = <- a`, RunOutput: float64(2.5), Output: map[string]interface{}{"b": float64(2.5)}},
		{Script: `a = make(chan string, 2); a <- "b"; b = <- a`, RunOutput: "b", Output: map[string]interface{}{"b": "b"}},

		// receive from closed channel
		{Script: `a = make(chan int64, 2); a <- 1; close(a); <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); a <- 1; close(a); <- a; <- a`, RunOutput: nil},

		// receive & send from same channel
		{Script: `a = make(chan int64, 2); a <- 1; a <- <- a`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); a <- 1; a <- <- a; <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); a <- 1; a <- <- a; b = <- a`, RunOutput: int64(1), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `a = make(chan int64, 2); a <- 1; a <- (<- a)`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); a <- 1; a <- (<- a); <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); a <- 1; a <- (<- a); b = <- a`, RunOutput: int64(1), Output: map[string]interface{}{"b": int64(1)}},

		// 1 then null into a
		{Script: `a = make(chan int64, 2); a <- a <- 1`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); a <- a <- 1; <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); a <- a <- 1; <- a; <- a`, RunOutput: int64(0)},

		// receive & send different channel
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- <- a`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- <- a; <- b`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- <- a; c = <- b`, RunOutput: int64(1), Output: map[string]interface{}{"c": int64(1)}},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- (<- a)`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- (<- a); <- b`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); a <- 1; b <- (<- a); c = <- b`, RunOutput: int64(1), Output: map[string]interface{}{"c": int64(1)}},

		// 1 into a then null into b
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); b <- a <- 1`, RunOutput: nil},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); b <- a <- 1; <- a`, RunOutput: int64(1)},
		{Script: `a = make(chan int64, 2); b = make(chan int64, 2); b <- a <- 1; <- a; <- b`, RunOutput: int64(0)},

		// test ok
		{Script: `a = make(chan int64, 2); a <- 1; b, ok = <- a`, RunOutput: int64(1), Output: map[string]interface{}{"b": int64(1), "ok": true}},
		{Script: `a = make(chan int64, 2); a <- 1; b, 1++ = <- a`, RunOutput: int64(1), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `a = make(chan int64, 2); a <- 1; close(a); b, ok = <- a`, RunOutput: int64(1), Output: map[string]interface{}{"b": int64(1), "ok": true}},
		{Script: `a = make(chan int64, 2); a <- 1; close(a); b = <- a; b, ok = <- a`, RunOutput: nil, Output: map[string]interface{}{"b": int64(1), "ok": false}},

		// test let ++
		{Script: `a = make(chan int64, 2); b = [1, 2, 3, 4]; c = 0; a <- 11; b[c++] = <- a; b[1]`, RunOutput: int64(11)},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestVMDelete(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `delete(1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `delete(1)`, RunError: fmt.Errorf("first argument to delete cannot be type int64")},
		{Script: `a = 1; delete("a"); a`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a = {"b": "b"}; delete(a)`, RunError: fmt.Errorf("second argument to delete cannot be nil for map")},

		{Script: `delete("a", 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `b = []; delete(a, b)`, Input: map[string]interface{}{"a": map[int32]interface{}{2: int32(2)}}, RunError: fmt.Errorf("cannot use type int32 as type []interface {} in delete"), Output: map[string]interface{}{"a": map[int32]interface{}{2: int32(2)}}},

		// test no variable
		{Script: `delete("a")`},
		{Script: `delete("a", false)`},
		{Script: `delete("a", true)`},
		{Script: `delete("a", nil)`},

		// test DeleteGlobal
		{Script: `a = 1; func b() { delete("a") }; b()`, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; func b() { delete("a", false) }; b()`, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; func b() { delete("a", true) }; b(); a`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a = 2; func b() { a = 3; delete("a"); return a }; b()`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 2; func b() { a = 3; delete("a", false); return a }; b()`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 2; func b() { a = 3; delete("a", true); return a }; b()`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a = 2; func b() { a = 3; delete("a") }; b(); a`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 2; func b() { a = 3; delete("a", false) }; b(); a`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 2; func b() { a = 3; delete("a", true) }; b(); a`, RunError: fmt.Errorf("undefined symbol 'a'")},

		// test empty map
		{Script: `delete(a, "a")`, Input: map[string]interface{}{"a": testMapEmpty}, Output: map[string]interface{}{"a": testMapEmpty}},

		// test map
		{Script: `a = {"b": "b"}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": "b"}; delete(a, "b"); a.b`, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": "b", "c":"c"}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": "c"}}},
		{Script: `a = {"b": "b", "c":"c"}; delete(a, "b"); a.b`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": "c"}}},

		// test key convert
		{Script: `delete(a, 2)`, Input: map[string]interface{}{"a": map[int32]interface{}{2: int32(2)}}, Output: map[string]interface{}{"a": map[int32]interface{}{}}},
		{Script: `delete(a, 2); a[2]`, Input: map[string]interface{}{"a": map[int32]interface{}{2: int32(2)}}, Output: map[string]interface{}{"a": map[int32]interface{}{}}},
		{Script: `delete(a, 2)`, Input: map[string]interface{}{"a": map[int32]interface{}{2: int32(2), 3: int32(3)}}, Output: map[string]interface{}{"a": map[int32]interface{}{3: int32(3)}}},
		{Script: `delete(a, 2); a[2]`, Input: map[string]interface{}{"a": map[int32]interface{}{2: int32(2), 3: int32(3)}}, Output: map[string]interface{}{"a": map[int32]interface{}{3: int32(3)}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestComment(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `# 1`},
		{Script: `# 1;`},
		{Script: `# 1 // 2`},
		{Script: `# 1 \n 2`},
		{Script: `# 1 # 2`},

		{Script: `1# 1`, RunOutput: int64(1)},
		{Script: `1# 1;`, RunOutput: int64(1)},
		{Script: `1# 1 // 2`, RunOutput: int64(1)},
		{Script: `1# 1 \n 2`, RunOutput: int64(1)},
		{Script: `1# 1 # 2`, RunOutput: int64(1)},

		{Script: `1
# 1`, RunOutput: int64(1)},
		{Script: `1
# 1;`, RunOutput: int64(1)},
		{Script: `1
# 1 // 2`, RunOutput: int64(1)},
		{Script: `1
# 1 \n 2`, RunOutput: int64(1)},
		{Script: `1
# 1 # 2`, RunOutput: int64(1)},

		{Script: `// 1`},
		{Script: `// 1;`},
		{Script: `// 1 // 2`},
		{Script: `// 1 \n 2`},
		{Script: `// 1 # 2`},

		{Script: `1// 1`, RunOutput: int64(1)},
		{Script: `1// 1;`, RunOutput: int64(1)},
		{Script: `1// 1 // 2`, RunOutput: int64(1)},
		{Script: `1// 1 \n 2`, RunOutput: int64(1)},
		{Script: `1// 1 # 2`, RunOutput: int64(1)},

		{Script: `1
// 1`, RunOutput: int64(1)},
		{Script: `1
// 1;`, RunOutput: int64(1)},
		{Script: `1
// 1 // 2`, RunOutput: int64(1)},
		{Script: `1
// 1 \n 2`, RunOutput: int64(1)},
		{Script: `1
// 1 # 2`, RunOutput: int64(1)},

		{Script: `/* 1 */`},
		{Script: `/* * 1 */`},
		{Script: `/* 1 * */`},
		{Script: `/** 1 */`},
		{Script: `/*** 1 */`},
		{Script: `/**** 1 */`},
		{Script: `/* 1 **/`},
		{Script: `/* 1 ***/`},
		{Script: `/* 1 ****/`},
		{Script: `/** 1 ****/`},
		{Script: `/*** 1 ****/`},
		{Script: `/**** 1 ****/`},

		{Script: `1/* 1 */`, RunOutput: int64(1)},
		{Script: `1/* * 1 */`, RunOutput: int64(1)},
		{Script: `1/* 1 * */`, RunOutput: int64(1)},
		{Script: `1/** 1 */`, RunOutput: int64(1)},
		{Script: `1/*** 1 */`, RunOutput: int64(1)},
		{Script: `1/**** 1 */`, RunOutput: int64(1)},
		{Script: `1/* 1 **/`, RunOutput: int64(1)},
		{Script: `1/* 1 ***/`, RunOutput: int64(1)},
		{Script: `1/* 1 ****/`, RunOutput: int64(1)},
		{Script: `1/** 1 ****/`, RunOutput: int64(1)},
		{Script: `1/*** 1 ****/`, RunOutput: int64(1)},
		{Script: `1/**** 1 ****/`, RunOutput: int64(1)},

		{Script: `/* 1 */1`, RunOutput: int64(1)},
		{Script: `/* * 1 */1`, RunOutput: int64(1)},
		{Script: `/* 1 * */1`, RunOutput: int64(1)},
		{Script: `/** 1 */1`, RunOutput: int64(1)},
		{Script: `/*** 1 */1`, RunOutput: int64(1)},
		{Script: `/**** 1 */1`, RunOutput: int64(1)},
		{Script: `/* 1 **/1`, RunOutput: int64(1)},
		{Script: `/* 1 ***/1`, RunOutput: int64(1)},
		{Script: `/* 1 ****/1`, RunOutput: int64(1)},
		{Script: `/** 1 ****/1`, RunOutput: int64(1)},
		{Script: `/*** 1 ****/1`, RunOutput: int64(1)},
		{Script: `/**** 1 ****/1`, RunOutput: int64(1)},

		{Script: `1
/* 1 */`, RunOutput: int64(1)},
		{Script: `1
/* * 1 */`, RunOutput: int64(1)},
		{Script: `1
/* 1 * */`, RunOutput: int64(1)},
		{Script: `1
/** 1 */`, RunOutput: int64(1)},
		{Script: `1
/*** 1 */`, RunOutput: int64(1)},
		{Script: `1
/**** 1 */`, RunOutput: int64(1)},
		{Script: `1
/* 1 **/`, RunOutput: int64(1)},
		{Script: `1
/* 1 ***/`, RunOutput: int64(1)},
		{Script: `1
/* 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/** 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/*** 1 ****/`, RunOutput: int64(1)},
		{Script: `1
/**** 1 ****/`, RunOutput: int64(1)},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestCancelWithContext(t *testing.T) {
	scripts := []string{
		`
b = 0
close(waitChan)
for {
	b = 1
}
`,
		`
b = 0
close(waitChan)
for {
	for {
		b = 1
	}
}
`,
		`
a = []
for i = 0; i < 20000; i++ {
	a += 1
}
b = 0
close(waitChan)
for {
	for i in a {
		b = i
	}
}
`,
		`
a = []
for i = 0; i < 20000; i++ {
	a += 1
}
b = 0
close(waitChan)
for i in a {
	for j in a {
		b = j
	}
}
`,
		`
close(waitChan)
b = 0
for i = 0; true; nil {
}
`,
		`
b = 0
close(waitChan)
for i = 0; true; nil {
	for j = 0; true; nil {
		b = 1
	}
}
`,
		`
a = {}
for i = 0; i < 20000; i++ {
	a[toString(i)] = 1
}
b = 0
close(waitChan)
for {
	for i in a {
		b = 1
	}
}
`,
		`
a = {}
for i = 0; i < 20000; i++ {
	a[toString(i)] = 1
}
b = 0
close(waitChan)
for i in a {
	for j in a {
		b = 1
	}
}
`,
		`
close(waitChan)
<- make(chan string)
`,
		`
a = ""
close(waitChan)
a = <- make(chan string)
`,
		`
for {
	a = ""
	close(waitChan)
	a = <- make(chan string)
}
`,
		`
a = make(chan int)
close(waitChan)
a <- 1
`,
		`
a = make(chan interface)
close(waitChan)
a <- nil
`,
		`
a = make(chan int64, 1)
close(waitChan)
for v in a { }
`,
		`
close(waitChan)
try {
	for { }
} catch { }
`,
	}
	for _, script := range scripts {
		runCancelTestWithContext(t, script)
	}
}

func runCancelTestWithContext(t *testing.T, script string) {
	waitChan := make(chan struct{}, 1)
	toString := func(value interface{}) string {
		return fmt.Sprintf("%v", value)
	}
	e := env.NewEnv()
	err := e.Define("waitChan", waitChan)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = e.Define("toString", toString)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-waitChan
		time.Sleep(time.Millisecond)
		cancel()
	}()

	_, err = ExecuteContext(ctx, e, nil, script)
	if err == nil || err.Error() != ErrInterrupt.Error() {
		t.Errorf("execute error - received %#v - expected: %#v - script: %v", err, ErrInterrupt, script)
	}
}

func TestContextConcurrency(t *testing.T) {
	var waitGroup sync.WaitGroup
	e := env.NewEnv()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := ExecuteContext(ctx, e, nil, "for { }")
			if err == nil || err.Error() != ErrInterrupt.Error() {
				t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	cancel()
	waitGroup.Wait()
	cancel()

	_, err := ExecuteContext(ctx, e, nil, "for { }")
	if err == nil || err.Error() != ErrInterrupt.Error() {
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}

	ctx, cancel = context.WithCancel(context.Background())
	_, err = ExecuteContext(ctx, e, nil, "for i = 0; i < 1000; i++ {}")
	if err != nil {
		t.Errorf("execute error - received: %v - expected: %v", err, nil)
	}
	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := ExecuteContext(ctx, e, nil, "for i = 0; i < 1000; i++ { }")
			if err != nil {
				t.Errorf("execute error - received: %v - expected: %v", err, nil)
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := ExecuteContext(ctx, e, nil, "for { }")
			if err == nil || err.Error() != ErrInterrupt.Error() {
				t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
			}
			waitGroup.Done()
		}()
	}
	time.Sleep(time.Millisecond)
	cancel()
	waitGroup.Wait()

	waitGroup.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			_, err := Execute(e, nil, "for i = 0; i < 1000; i++ { }")
			if err != nil {
				t.Errorf("execute error - received: %v - expected: %v", err, nil)
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func TestContextFunction(t *testing.T) {
	t.Parallel()

	e := env.NewEnv()
	script := `
		func myFunc(myVar) {
			myVar = 3
		}`
	envModule, err := e.NewModule("a")
	if err != nil {
		t.Fatal("NewModule error:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	_, err = ExecuteContext(ctx, envModule, nil, script)
	if err != nil {
		t.Errorf("execute error - received %#v - expected: %#v", err, nil)
	}
	cancel()

	script = "a.myFunc(2)"

	ctx, cancel = context.WithCancel(context.Background())
	_, err = ExecuteContext(ctx, e, nil, script)
	if err != nil {
		t.Errorf("execute error - received %#v - expected: %#v", err, nil)
	}
	cancel()
}

func TestAssignToInterface(t *testing.T) {
	t.Parallel()

	e := env.NewEnv()
	X := new(struct {
		Stdout io.Writer
	})
	err := e.Define("X", X)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = e.Define("a", new(os.File))
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = Execute(e, nil, "X.Stdout = a")
	if err != nil {
		t.Errorf("execute error - received %#v - expected: %#v", err, ErrInterrupt)
	}
}

// TestValueEqual do some basic ValueEqual tests for coverage
func TestValueEqual(t *testing.T) {
	t.Parallel()

	result := valueEqual(true, true)
	if result != true {
		t.Fatal("ValueEqual")
	}
	result = valueEqual(true, false)
	if result != false {
		t.Fatal("ValueEqual")
	}
	result = valueEqual(false, true)
	if result != false {
		t.Fatal("ValueEqual")
	}
}

// TestUnknownCases tests switch cases that are the unknown cases
func TestUnknownCases(t *testing.T) {
	t.Parallel()

	oneLiteral := &ast.LiteralExpr{Literal: reflect.ValueOf(int64(1))}
	type (
		BadStmt struct {
			ast.StmtImpl
		}
		BadExpr struct {
			ast.ExprImpl
		}
		BadOperator struct {
			ast.OperatorImpl
		}
	)

	stmts := []ast.Stmt{
		&BadStmt{},
		&ast.ExprStmt{Expr: &BadExpr{}},
		&ast.ExprStmt{Expr: &ast.OpExpr{Op: &BadOperator{}}},
		&ast.ExprStmt{Expr: &ast.UnaryExpr{Expr: oneLiteral}},
		&ast.ExprStmt{Expr: &ast.OpExpr{Op: &ast.BinaryOperator{LHS: oneLiteral}}},
		&ast.ExprStmt{Expr: &ast.OpExpr{Op: &ast.ComparisonOperator{LHS: oneLiteral, RHS: oneLiteral}}},
		&ast.ExprStmt{Expr: &ast.OpExpr{Op: &ast.AddOperator{LHS: oneLiteral, RHS: oneLiteral}}},
		&ast.ExprStmt{Expr: &ast.OpExpr{Op: &ast.MultiplyOperator{LHS: oneLiteral, RHS: oneLiteral}}},
	}

	for _, stmt := range stmts {
		e := env.NewEnv()
		_, err := Run(e, nil, stmt)
		if err == nil {
			t.Errorf("no error - stmt: %#v", stmt)
		} else if len(err.Error()) < 9 || err.Error()[:8] != "unknown " {
			t.Errorf("err: %v - stmt: %#v", err, stmt)
		}
	}
}

type Foo struct {
	Value int
}
type Bar struct {
	Foo
	Ref *int
}

func (f Foo) ValueReceiver() int {
	return f.Value
}

func (b *Bar) PointerReceiver() (int, int) {
	b.Value = 0
	*b.Ref = 0
	return b.Value, *b.Ref
}

func TestCallStructMethod(t *testing.T) {
	t.Parallel()

	ref := 10
	ptr := &Bar{
		Foo: Foo{
			Value: 100,
		},
		Ref: &ref,
	}
	val := Bar{
		Foo: Foo{
			Value: 200,
		},
		Ref: &ref,
	}

	// execution in native go
	v := ptr.ValueReceiver()
	if v != 100 {
		t.Fatal("ptr: call value receiver failed, v should equal to 100")
	}
	v, r := ptr.PointerReceiver()
	if v != 0 || r != 0 {
		t.Fatal("ptr: call pointer receiver failed, both should be 0")
	}

	ref = 10
	v = val.ValueReceiver()
	if v != 200 {
		t.Fatal("val: call value receiver failed, v should equal to 200")
	}
	v, r = val.PointerReceiver()
	if v != 0 || r != 0 {
		t.Fatal("val: call pointer receiver failed, both should be 0")
	}

	// reinitialize values before executing script in VM
	ptr.Value = 100
	val.Value = 200

	// Define pointer 'ptr' to struct Bar in VM
	ref = 10
	tests := []Test{
		{Script: "ptr.ValueReceiver()", Input: map[string]interface{}{"ptr": ptr}, RunOutput: 100},
		{Script: "ptr.PointerReceiver()", Input: map[string]interface{}{"ptr": ptr}, RunOutput: []interface{}{0, 0}},
	}
	runTests(t, tests, nil, &Options{Debug: true})

	// Define value 'val' to struct Bar in VM
	ref = 10
	tests = []Test{
		{Script: "val.ValueReceiver()", Input: map[string]interface{}{"val": val}, RunOutput: 200},
		{Script: "val.PointerReceiver()", Input: map[string]interface{}{"val": val}, RunOutput: []interface{}{0, 0}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func BenchmarkFibGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(29)
	}
}

func BenchmarkFibVM(b *testing.B) {
	b.StopTimer()

	e := env.NewEnv()
	a, err := e.NewModule("a")
	if err != nil {
		b.Fatal("NewModule error:", err)
	}

	script := `
fib = func(x) {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}`

	_, err = Execute(a, nil, script)
	if err != nil {
		b.Fatal("Execute error:", err)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err = Execute(e, nil, "a.fib(29)")
		if err != nil {
			b.Fatal("Execute error:", err)
		}
	}
}

func TestLetsStatementPosition(t *testing.T) {
	src := `a, b = 1, 2
`
	stmts, err := parser.ParseSrc(src)
	if err != nil {
		t.Fatal(err)
	}

	var stmtFound bool
	err = astutil.Walk(stmts, func(e interface{}) error {
		switch e := e.(type) {
		case *ast.LetsStmt:
			if len(e.LHSS) == 2 {
				if is, want := e.Position().Line, 1; is != want {
					t.Fatalf("%v != %v", is, want)
				}
				if is, want := e.Position().Column, 1; is != want {
					t.Fatalf("%v != %v", is, want)
				}
			}
			stmtFound = true
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	if is, want := stmtFound, true; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}
