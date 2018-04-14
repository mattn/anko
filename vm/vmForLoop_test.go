package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestForLoop(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `break`, RunError: fmt.Errorf("unexpected break statement")},
		{Script: `continue`, RunError: fmt.Errorf("unexpected continue statement")},
		{Script: `for 1++ { }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `for { 1++ }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `for a in 1++ { }`, RunError: fmt.Errorf("invalid operation")},

		{Script: `for { break }`, RunOutput: nil},
		{Script: `for {a = 1; if a == 1 { break } }`, RunOutput: nil},
		{Script: `a = 1; for { if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for { if a == 1 { break }; a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a = 1; for { a++; if a == 2 { continue } else { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; for { a++; if a == 2 { continue }; if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `for a in [1] { if a == 1 { break } }`, RunOutput: nil},
		{Script: `for a in [1, 2] { if a == 2 { break } }`, RunOutput: nil},
		{Script: `for a in [1, 2, 3] { if a == 3 { break } }`, RunOutput: nil},

		{Script: `a = [1]; for b in a { if b == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = [1, 2]; for b in a { if b == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1, 2, 3]; for b in a { if b == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{Script: `a = [1, 2]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": int64(2)}},
		{Script: `a = [1, 2, 3]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}, "b": int64(3)}},

		{Script: `a = 1; for a < 2 { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for a < 3 { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for nil { a++; if a > 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for nil { a++; if a > 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for true { a++; if a > 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; for true { a++; if a > 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(4)}},

		{Script: `func x() { return [1] }; for b in x() { if b == 1 { break } }`, RunOutput: nil},
		{Script: `func x() { return [1, 2] }; for b in x() { if b == 2 { break } }`, RunOutput: nil},
		{Script: `func x() { return [1, 2, 3] }; for b in x() { if b == 3 { break } }`, RunOutput: nil},

		{Script: `func x() { a = 1; for { if a == 1 { return } } }; x()`, RunOutput: nil},
		{Script: `func x() { a = 1; for { if a == 1 { return nil } } }; x()`, RunOutput: nil},
		{Script: `func x() { a = 1; for { if a == 1 { return true } } }; x()`, RunOutput: true},
		{Script: `func x() { a = 1; for { if a == 1 { return 1 } } }; x()`, RunOutput: int64(1)},
		{Script: `func x() { a = 1; for { if a == 1 { return 1.1 } } }; x()`, RunOutput: float64(1.1)},
		{Script: `func x() { a = 1; for { if a == 1 { return "a" } } }; x()`, RunOutput: "a"},

		{Script: `func x() { for a in [1, 2, 3] { if a == 3 { return } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a in [1, 2, 3] { if a == 1 { continue } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a in [1, 2, 3] { if a == 1 { continue };  if a == 3 { return } } }; x()`, RunOutput: nil},

		{Script: `func x() { return [1, 2] }; a = 1; for i in x() { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `func x() { return [1, 2, 3] }; a = 1; for i in x() { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(4)}},

		// TOFIX:
		// {Script: `for a = 1, 2; nil; nil { return }`},
		// {Script: `for a, b = 1; nil; nil { return }`},
		// {Script: `for a, b = 1, 2; nil; nil { return }`},

		{Script: `for a.b = 1; nil; nil { return }`, RunError: fmt.Errorf("undefined symbol 'a'")},

		{Script: `for a = 1; nil; nil { if a == 1 { break } }`, RunOutput: nil},
		{Script: `for a = 1; nil; nil { if a == 2 { break }; a++ }`, RunOutput: nil},
		{Script: `for a = 1; nil; nil { a++; if a == 3 { break } }`, RunOutput: nil},

		{Script: `for a = 1; a < 1; nil { }`, RunOutput: nil},
		{Script: `for a = 1; a > 1; nil { }`, RunOutput: nil},
		{Script: `for a = 1; a == 1; nil { break }`, RunOutput: nil},

		{Script: `for a = 1; a == 1; a++ { }`, RunOutput: nil},
		{Script: `for a = 1; a < 2; a++ { }`, RunOutput: nil},
		{Script: `for a = 1; a < 3; a++ { }`, RunOutput: nil},

		{Script: `a = 1; for b = 1; a < 1; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = ["123", "456", "789"]; b = ""; for i = 0; i < len(a); i++ { b += a[i][len(a[i]) - 2:]; b += a[i][:len(a[i]) - 2] }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{"123", "456", "789"}, "b": "231564897"}},
		{Script: `a = [[["123"], ["456"]], [["789"]]]; b = ""; for i = 0; i < len(a); i++ { for j = 0; j < len(a[i]); j++ {  for k = 0; k < len(a[i][j]); k++ { for l = 0; l < len(a[i][j][k]); l++ { b += a[i][j][k][l] + "-" } } } }`,
			RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{[]interface{}{[]interface{}{"123"}, []interface{}{"456"}}, []interface{}{[]interface{}{"789"}}}, "b": "1-2-3-4-5-6-7-8-9-"}},

		{Script: `func x() { for a = 1; a < 3; a++ { if a == 1 { return a } } }; x()`, RunOutput: int64(1)},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 2 { return a } } }; x()`, RunOutput: int64(2)},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 3 { return a } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 4 { return a } } }; x()`, RunOutput: nil},

		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 1 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 2 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 3 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 4 { continue } }; return a }; x()`, RunOutput: int64(3)},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{reflect.Value{}}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{reflect.Value{}}, "b": reflect.Value{}}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{nil}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{nil}, "b": nil}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{true}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{true}, "b": true}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{int32(1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int32(1)}, "b": int32(1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{int64(1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": float32(1.1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": float64(1.1)}},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}, "b": interface{}(reflect.Value{})}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(nil)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(nil)}, "b": interface{}(nil)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(true)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(true)}, "b": interface{}(true)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}, "b": interface{}(int32(1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}, "b": interface{}(int64(1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}, "b": interface{}(float32(1.1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}, "b": interface{}(float64(1.1))}},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": interface{}([]interface{}{nil})}, RunOutput: nil, Output: map[string]interface{}{"a": interface{}([]interface{}{nil}), "b": nil}},

		{Script: `for i in nil { }`, RunError: fmt.Errorf("for cannot loop over type interface")},
		{Script: `for i in true { }`, RunError: fmt.Errorf("for cannot loop over type bool")},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": reflect.Value{}}, RunError: fmt.Errorf("for cannot loop over type struct"), Output: map[string]interface{}{"a": reflect.Value{}}},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": interface{}(nil)}, RunError: fmt.Errorf("for cannot loop over type interface"), Output: map[string]interface{}{"a": interface{}(nil)}},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": interface{}(true)}, RunError: fmt.Errorf("for cannot loop over type bool"), Output: map[string]interface{}{"a": interface{}(true)}},
		{Script: `for i in [1, 2, 3] { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; a < 3; a++ { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = b; a < 3; a++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; b < 3; a++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; a < 3; b++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `a = 1; b = [{"c": "c"}]; for i in b { a = i }`, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"c": "c"}, "b": []interface{}{map[string]interface{}{"c": "c"}}}},
		{Script: `a = 1; b = {"x": [{"y": "y"}]};  for i in b.x { a = i }`, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"y": "y"}, "b": map[string]interface{}{"x": []interface{}{map[string]interface{}{"y": "y"}}}}},

		{Script: `a = {}; b = 1; for i in a { b = i }; b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{}, "b": int64(1)}},
		{Script: `a = {"x": 2}; b = 1; for i in a { b = i }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": "x"}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { b++ }; b`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(2)}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},

		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { continue }; b = i }; b`, RunOutput: "y", Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "y"}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { continue }; b = i }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": "x"}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "x" { return 1 } }`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "y" { return 2 } }`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { break }; b++ }; if b > 1 { return false } else { return true }`, RunOutput: true, Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { break }; b++ }; if b > 1 { return false } else { return true }`, RunOutput: true, Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 1; for i in a { if (i ==  "x" || i ==  "y") { break }; b++ }; b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(1)}},

		{Script: `a = ["123", "456", "789"]; b = ""; for v in a { b += v[len(v) - 2:]; b += v[:len(v) - 2] }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{"123", "456", "789"}, "b": "231564897"}},
		{Script: `a = [[["123"], ["456"]], [["789"]]]; b = ""; for x in a { for y in x  {  for z in y { for i = 0; i < len(z); i++ { b += z[i] + "-" } } } }`,
			RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{[]interface{}{[]interface{}{"123"}, []interface{}{"456"}}, []interface{}{[]interface{}{"789"}}}, "b": "1-2-3-4-5-6-7-8-9-"}},

		{Script: `a = {"x": 2}; b = 0; for k, v in a { b = k }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": "x"}},
		{Script: `a = {"x": 2}; b = 0; for k, v in a { b = v }; b`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[string]interface{}{"x": int64(2)}, "b": int64(2)}},
	}
	RunTests(t, tests, nil)
}
