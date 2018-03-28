package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestForLoop(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `break`, runError: fmt.Errorf("Unexpected break statement")},
		{script: `continue`, runError: fmt.Errorf("Unexpected continue statement")},
		{script: `for 1++ { }`, runError: fmt.Errorf("Invalid operation")},
		{script: `for { 1++ }`, runError: fmt.Errorf("Invalid operation")},
		{script: `for a in 1++ { }`, runError: fmt.Errorf("Invalid operation")},

		{script: `for { break }`, runOutput: nil},
		{script: `for {a = 1; if a == 1 { break } }`, runOutput: nil},
		{script: `a = 1; for { if a == 1 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for { if a == 1 { break }; a++ }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},

		{script: `a = 1; for { a++; if a == 2 { continue } else { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: `a = 1; for { a++; if a == 2 { continue }; if a == 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: `for a in [1] { if a == 1 { break } }`, runOutput: nil},
		{script: `for a in [1, 2] { if a == 2 { break } }`, runOutput: nil},
		{script: `for a in [1, 2, 3] { if a == 3 { break } }`, runOutput: nil},

		{script: `a = [1]; for b in a { if b == 1 { break } }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: `a = [1, 2]; for b in a { if b == 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{script: `a = [1, 2, 3]; for b in a { if b == 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: `a = [1]; b = 0; for c in a { b = c }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: `a = [1, 2]; b = 0; for c in a { b = c }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": int64(2)}},
		{script: `a = [1, 2, 3]; b = 0; for c in a { b = c }`, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}, "b": int64(3)}},

		{script: `a = 1; for a < 2 { a++ }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for a < 3 { a++ }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: `a = 1; for nil { a++; if a > 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for nil { a++; if a > 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for true { a++; if a > 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: `a = 1; for true { a++; if a > 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(4)}},

		{script: `func x() { return [1] }; for b in x() { if b == 1 { break } }`, runOutput: nil},
		{script: `func x() { return [1, 2] }; for b in x() { if b == 2 { break } }`, runOutput: nil},
		{script: `func x() { return [1, 2, 3] }; for b in x() { if b == 3 { break } }`, runOutput: nil},

		{script: `func x() { a = 1; for { if a == 1 { return } } }; x()`, runOutput: nil},
		{script: `func x() { a = 1; for { if a == 1 { return nil } } }; x()`, runOutput: nil},
		{script: `func x() { a = 1; for { if a == 1 { return true } } }; x()`, runOutput: true},
		{script: `func x() { a = 1; for { if a == 1 { return 1 } } }; x()`, runOutput: int64(1)},
		{script: `func x() { a = 1; for { if a == 1 { return 1.1 } } }; x()`, runOutput: float64(1.1)},
		{script: `func x() { a = 1; for { if a == 1 { return "a" } } }; x()`, runOutput: "a"},

		{script: `func x() { for a in [1, 2, 3] { if a == 3 { return } } }; x()`, runOutput: nil},
		{script: `func x() { for a in [1, 2, 3] { if a == 1 { continue } } }; x()`, runOutput: nil},
		{script: `func x() { for a in [1, 2, 3] { if a == 1 { continue };  if a == 3 { return } } }; x()`, runOutput: nil},

		{script: `func x() { return [1, 2] }; a = 1; for i in x() { a++ }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},
		{script: `func x() { return [1, 2, 3] }; a = 1; for i in x() { a++ }`, runOutput: nil, output: map[string]interface{}{"a": int64(4)}},

		// TOFIX:
		// {script: `for a = 1, 2; nil; nil { return }`},
		// {script: `for a, b = 1; nil; nil { return }`},
		// {script: `for a, b = 1, 2; nil; nil { return }`},

		{script: `for a.b = 1; nil; nil { return }`, runError: fmt.Errorf("Undefined symbol 'a'")},

		{script: `for a = 1; nil; nil { if a == 1 { break } }`, runOutput: nil},
		{script: `for a = 1; nil; nil { if a == 2 { break }; a++ }`, runOutput: nil},
		{script: `for a = 1; nil; nil { a++; if a == 3 { break } }`, runOutput: nil},

		{script: `for a = 1; a < 1; nil { }`, runOutput: nil},
		{script: `for a = 1; a > 1; nil { }`, runOutput: nil},
		{script: `for a = 1; a == 1; nil { break }`, runOutput: nil},

		{script: `for a = 1; a == 1; a++ { }`, runOutput: nil},
		{script: `for a = 1; a < 2; a++ { }`, runOutput: nil},
		{script: `for a = 1; a < 3; a++ { }`, runOutput: nil},

		{script: `a = 1; for b = 1; a < 1; a++ { }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 2; a++ { }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for b = 1; a < 3; a++ { }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { continue } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { continue } }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { continue } }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 1; a++ {  if a == 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 2; a++ {  if a == 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for b = 1; a < 3; a++ {  if a == 2 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for b = 1; a < 1; a++ {  if a == 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `a = 1; for b = 1; a < 2; a++ {  if a == 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(2)}},
		{script: `a = 1; for b = 1; a < 3; a++ {  if a == 3 { break } }`, runOutput: nil, output: map[string]interface{}{"a": int64(3)}},

		{script: `func x() { for a = 1; a < 3; a++ { if a == 1 { return a } } }; x()`, runOutput: int64(1)},
		{script: `func x() { for a = 1; a < 3; a++ { if a == 2 { return a } } }; x()`, runOutput: int64(2)},
		{script: `func x() { for a = 1; a < 3; a++ { if a == 3 { return a } } }; x()`, runOutput: nil},
		{script: `func x() { for a = 1; a < 3; a++ { if a == 4 { return a } } }; x()`, runOutput: nil},

		{script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 1 { continue } }; return a }; x()`, runOutput: int64(3)},
		{script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 2 { continue } }; return a }; x()`, runOutput: int64(3)},
		{script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 3 { continue } }; return a }; x()`, runOutput: int64(3)},
		{script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 4 { continue } }; return a }; x()`, runOutput: int64(3)},

		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{reflect.Value{}}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{reflect.Value{}}, "b": reflect.Value{}}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{nil}, "b": nil}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{true}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{true}, "b": true}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{int32(1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int32(1)}, "b": int32(1)}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{int64(1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": float32(1.1)}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": float64(1.1)}},

		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}, "b": interface{}(reflect.Value{})}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(nil)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(nil)}, "b": interface{}(nil)}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(true)}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(true)}, "b": interface{}(true)}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}, "b": interface{}(int32(1))}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}, "b": interface{}(int64(1))}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}, "b": interface{}(float32(1.1))}},
		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}}, runOutput: nil, output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}, "b": interface{}(float64(1.1))}},

		{script: `b = 0; for i in a { b = i }`, input: map[string]interface{}{"a": interface{}([]interface{}{nil})}, runOutput: nil, output: map[string]interface{}{"a": interface{}([]interface{}{nil}), "b": nil}},

		{script: `for i in nil { }`, runError: fmt.Errorf("for cannot loop over type interface")},
		{script: `for i in true { }`, runError: fmt.Errorf("for cannot loop over type bool")},
		{script: `for i in a { }`, input: map[string]interface{}{"a": reflect.Value{}}, runError: fmt.Errorf("for cannot loop over type struct"), output: map[string]interface{}{"a": reflect.Value{}}},
		{script: `for i in a { }`, input: map[string]interface{}{"a": interface{}(nil)}, runError: fmt.Errorf("for cannot loop over type interface"), output: map[string]interface{}{"a": interface{}(nil)}},
		{script: `for i in a { }`, input: map[string]interface{}{"a": interface{}(true)}, runError: fmt.Errorf("for cannot loop over type bool"), output: map[string]interface{}{"a": interface{}(true)}},
		{script: `for i in [1, 2, 3] { b++ }`, runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: `for a = 1; a < 3; a++ { b++ }`, runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: `for a = b; a < 3; a++ { }`, runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: `for a = 1; b < 3; a++ { }`, runError: fmt.Errorf("Undefined symbol 'b'")},
		{script: `for a = 1; a < 3; b++ { }`, runError: fmt.Errorf("Undefined symbol 'b'")},

		{script: `a = 1; b = [{"c": "c"}]; for i in b { a = i }`, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"c": "c"}, "b": []interface{}{map[string]interface{}{"c": "c"}}}},
		{script: `a = 1; b = {"x": [{"y": "y"}]};  for i in b.x { a = i }`, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"y": "y"}, "b": map[string]interface{}{"x": []interface{}{map[string]interface{}{"y": "y"}}}}},

		{script: `a = {}; b = 1; for i in a { b = i }; b`, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{}, "b": int64(1)}},
		{script: `a = {"x": 2}; b = 1; for i in a { b = i }; b`, runOutput: "x", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": "x"}},
		{script: `a = {"x": 2, "y": 3}; b = 0; for i in a { b++ }; b`, runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(2)}},
		{script: `a = {"x": 2, "y": 3}; for i in a { b++ }`, runError: fmt.Errorf("Undefined symbol 'b'"), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},

		{script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { continue }; b = i }; b`, runOutput: "y", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "y"}},
		{script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { continue }; b = i }; b`, runOutput: "x", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "x"}},
		{script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "x" { return 1 } }`, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "y" { return 2 } }`, runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { break }; b++ }; if b > 1 { return false } else { return true }`, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { break }; b++ }; if b > 1 { return false } else { return true }`, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{script: `a = {"x": 2, "y": 3}; b = 1; for i in a { if (i ==  "x" || i ==  "y") { break }; b++ }; b`, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(1)}},

		{script: `a = {"x": 2}; b = 0; for k, v in a { b = k }; b`, runOutput: "x", output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": "x"}},
		{script: `a = {"x": 2}; b = 0; for k, v in a { b = v }; b`, runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": int64(2)}},
	}
	runTests(t, tests)
}
