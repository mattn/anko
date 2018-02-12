package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReturns(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "return nil++", runError: fmt.Errorf("Invalid operation")},

		{script: "return", runOutput: nil},
		{script: "return nil", runOutput: nil},
		{script: "return true", runOutput: true},
		{script: "return 1", runOutput: int64(1)},
		{script: "return 1.1", runOutput: float64(1.1)},
		{script: "return \"a\"", runOutput: "a"},

		{script: "b()", input: map[string]interface{}{"b": func() {}}, runOutput: nil},
		{script: "b()", input: map[string]interface{}{"b": func() reflect.Value { return reflect.Value{} }}, runOutput: reflect.Value{}},
		{script: "b()", input: map[string]interface{}{"b": func() interface{} { return nil }}, runOutput: nil},
		{script: "b()", input: map[string]interface{}{"b": func() bool { return true }}, runOutput: true},
		{script: "b()", input: map[string]interface{}{"b": func() int32 { return int32(1) }}, runOutput: int32(1)},
		{script: "b()", input: map[string]interface{}{"b": func() int64 { return int64(1) }}, runOutput: int64(1)},
		{script: "b()", input: map[string]interface{}{"b": func() float32 { return float32(1.1) }}, runOutput: float32(1.1)},
		{script: "b()", input: map[string]interface{}{"b": func() float64 { return float64(1.1) }}, runOutput: float64(1.1)},
		{script: "b()", input: map[string]interface{}{"b": func() string { return "a" }}, runOutput: "a"},

		{script: "b(a)", input: map[string]interface{}{"a": reflect.Value{}, "b": func(c reflect.Value) reflect.Value { return c }}, runOutput: reflect.Value{}, output: map[string]interface{}{"a": reflect.Value{}}},
		{script: "b(a)", input: map[string]interface{}{"a": nil, "b": func(c interface{}) interface{} { return c }}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "b(a)", input: map[string]interface{}{"a": true, "b": func(c bool) bool { return c }}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "b(a)", input: map[string]interface{}{"a": int32(1), "b": func(c int32) int32 { return c }}, runOutput: int32(1), output: map[string]interface{}{"a": int32(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": int64(1), "b": func(c int64) int64 { return c }}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": float32(1.1), "b": func(c float32) float32 { return c }}, runOutput: float32(1.1), output: map[string]interface{}{"a": float32(1.1)}},
		{script: "b(a)", input: map[string]interface{}{"a": float64(1.1), "b": func(c float64) float64 { return c }}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: "b(a)", input: map[string]interface{}{"a": "a", "b": func(c string) string { return c }}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: "b(a)", input: map[string]interface{}{"a": "a", "b": func(c bool) bool { return c }}, runError: fmt.Errorf("argument type string cannot be used for function argument type bool"), output: map[string]interface{}{"a": "a"}},
		{script: "b(a)", input: map[string]interface{}{"a": int64(1), "b": func(c int32) int32 { return c }}, runOutput: int32(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": int32(1), "b": func(c int64) int64 { return c }}, runOutput: int64(1), output: map[string]interface{}{"a": int32(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": float64(1.25), "b": func(c float32) float32 { return c }}, runOutput: float32(1.25), output: map[string]interface{}{"a": float64(1.25)}},
		{script: "b(a)", input: map[string]interface{}{"a": float32(1.25), "b": func(c float64) float64 { return c }}, runOutput: float64(1.25), output: map[string]interface{}{"a": float32(1.25)}},
		{script: "b(a)", input: map[string]interface{}{"a": true, "b": func(c string) string { return c }}, runError: fmt.Errorf("argument type bool cannot be used for function argument type string"), output: map[string]interface{}{"a": true}},

		{script: "func aFunc() {}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return}; a = aFunc()", runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: "func aFunc() {return nil}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return true}; aFunc()", runOutput: true},
		{script: "func aFunc() {return 1}; aFunc()", runOutput: int64(1)},
		{script: "func aFunc() {return 1.1}; aFunc()", runOutput: float64(1.1)},
		{script: "func aFunc() {return \"a\"}; aFunc()", runOutput: "a"},

		{script: "func aFunc() {return 1 + 2}; aFunc()", runOutput: int64(3)},
		{script: "func aFunc() {return 1.25 + 2.25}; aFunc()", runOutput: float64(3.5)},
		{script: "func aFunc() {return \"a\" + \"b\"}; aFunc()", runOutput: "ab"},

		{script: "func aFunc() {return 1 + 2, 3 + 4}; aFunc()", runOutput: []interface{}{int64(3), int64(7)}},
		{script: "func aFunc() {return 1.25 + 2.25, 3.25 + 4.25}; aFunc()", runOutput: []interface{}{float64(3.5), float64(7.5)}},
		{script: "func aFunc() {return \"a\" + \"b\", \"c\" + \"d\"}; aFunc()", runOutput: []interface{}{"ab", "cd"}},

		{script: "func aFunc() {return nil, nil}; aFunc()", runOutput: []interface{}{nil, nil}},
		{script: "func aFunc() {return true, false}; aFunc()", runOutput: []interface{}{true, false}},
		{script: "func aFunc() {return 1, 2}; aFunc()", runOutput: []interface{}{int64(1), int64(2)}},
		{script: "func aFunc() {return 1.1, 2.2}; aFunc()", runOutput: []interface{}{float64(1.1), float64(2.2)}},
		{script: "func aFunc() {return \"a\", \"b\"}; aFunc()", runOutput: []interface{}{"a", "b"}},

		{script: "func aFunc() {return [nil]}; aFunc()", runOutput: []interface{}{nil}},
		{script: "func aFunc() {return [nil, nil]}; aFunc()", runOutput: []interface{}{nil, nil}},
		{script: "func aFunc() {return [nil, nil, nil]}; aFunc()", runOutput: []interface{}{nil, nil, nil}},
		{script: "func aFunc() {return [nil, nil], [nil, nil]}; aFunc()", runOutput: []interface{}{[]interface{}{nil, nil}, []interface{}{nil, nil}}},

		{script: "func aFunc() {return [true]}; aFunc()", runOutput: []interface{}{true}},
		{script: "func aFunc() {return [true, false]}; aFunc()", runOutput: []interface{}{true, false}},
		{script: "func aFunc() {return [true, false, true]}; aFunc()", runOutput: []interface{}{true, false, true}},
		{script: "func aFunc() {return [true, false], [false, true]}; aFunc()", runOutput: []interface{}{[]interface{}{true, false}, []interface{}{false, true}}},

		{script: "func aFunc() {return []}; aFunc()", runOutput: []interface{}{}},
		{script: "func aFunc() {return [1]}; aFunc()", runOutput: []interface{}{int64(1)}},
		{script: "func aFunc() {return [1, 2]}; aFunc()", runOutput: []interface{}{int64(1), int64(2)}},
		{script: "func aFunc() {return [1, 2, 3]}; aFunc()", runOutput: []interface{}{int64(1), int64(2), int64(3)}},
		{script: "func aFunc() {return [1, 2], [3, 4]}; aFunc()", runOutput: []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}},

		{script: "func aFunc() {return [1.1]}; aFunc()", runOutput: []interface{}{float64(1.1)}},
		{script: "func aFunc() {return [1.1, 2.2]}; aFunc()", runOutput: []interface{}{float64(1.1), float64(2.2)}},
		{script: "func aFunc() {return [1.1, 2.2, 3.3]}; aFunc()", runOutput: []interface{}{float64(1.1), float64(2.2), float64(3.3)}},
		{script: "func aFunc() {return [1.1, 2.2], [3.3, 4.4]}; aFunc()", runOutput: []interface{}{[]interface{}{float64(1.1), float64(2.2)}, []interface{}{float64(3.3), float64(4.4)}}},

		{script: "func aFunc() {return [\"a\"]}; aFunc()", runOutput: []interface{}{"a"}},
		{script: "func aFunc() {return [\"a\", \"b\"]}; aFunc()", runOutput: []interface{}{"a", "b"}},
		{script: "func aFunc() {return [\"a\", \"b\", \"c\"]}; aFunc()", runOutput: []interface{}{"a", "b", "c"}},
		{script: "func aFunc() {return [\"a\", \"b\"], [\"c\", \"d\"]}; aFunc()", runOutput: []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}}},

		{script: "func aFunc() {return nil, nil}; aFunc()", runOutput: []interface{}{interface{}(nil), interface{}(nil)}},
		{script: "func aFunc() {return true, false}; aFunc()", runOutput: []interface{}{true, false}},
		{script: "func aFunc() {return 1, 2}; aFunc()", runOutput: []interface{}{int64(1), int64(2)}},
		{script: "func aFunc() {return 1.1, 2.2}; aFunc()", runOutput: []interface{}{float64(1.1), float64(2.2)}},
		{script: "func aFunc() {return \"a\", \"b\"}; aFunc()", runOutput: []interface{}{"a", "b"}},

		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: reflect.Value{}, output: map[string]interface{}{"a": reflect.Value{}}},

		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: []interface{}{reflect.Value{}, reflect.Value{}}, output: map[string]interface{}{"a": reflect.Value{}}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: []interface{}{nil, nil}, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: []interface{}{true, true}, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": int32(1)}, runOutput: []interface{}{int32(1), int32(1)}, output: map[string]interface{}{"a": int32(1)}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: []interface{}{int64(1), int64(1)}, output: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": float32(1.1)}, runOutput: []interface{}{float32(1.1), float32(1.1)}, output: map[string]interface{}{"a": float32(1.1)}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: []interface{}{float64(1.1), float64(1.1)}, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: []interface{}{"a", "a"}, output: map[string]interface{}{"a": "a"}},

		{script: "func a(x) { return x}; a(nil)", runOutput: nil},
		{script: "func a(x) { return x}; a(true)", runOutput: true},
		{script: "func a(x) { return x}; a(1)", runOutput: int64(1)},
		{script: "func a(x) { return x}; a(1.1)", runOutput: float64(1.1)},
		{script: "func a(x) { return x}; a(\"a\")", runOutput: "a"},

		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": true}, runOutput: nil, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": int64(1)}, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": float64(1.1)}, runOutput: nil, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": "a"}, runOutput: nil, output: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {return nil, nil}; a, b = aFunc()", runOutput: nil, output: map[string]interface{}{"a": nil, "b": nil}},
		{script: "func aFunc() {return true, false}; a, b = aFunc()", runOutput: false, output: map[string]interface{}{"a": true, "b": false}},
		{script: "func aFunc() {return 1, 2}; a, b = aFunc()", runOutput: int64(2), output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "func aFunc() {return 1.1, 2.2}; a, b = aFunc()", runOutput: float64(2.2), output: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{script: "func aFunc() {return \"a\", \"b\"}; a, b = aFunc()", runOutput: "b", output: map[string]interface{}{"a": "a", "b": "b"}},
	}
	runTests(t, tests)
}

func TestFunctions(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = nil; a()", runError: fmt.Errorf("can not call type interface {}"), output: map[string]interface{}{"a": nil}},
		{script: "a = true; a()", runError: fmt.Errorf("can not call type bool"), output: map[string]interface{}{"a": true}},
		{script: "a = nil; b = func c(d) { return d == nil }; c = nil; c(a)", runError: fmt.Errorf("can not call type interface {}"), output: map[string]interface{}{"a": nil}},
		{script: "a = [true]; func b(c) { return c() }; b(a)", runError: fmt.Errorf("can not call type []interface {}")},
		{script: "a = [func () { return nil}]; func b(c) { return c() }; b(a[1])", runError: fmt.Errorf("index out of range")},
		{script: "a = nil; func b(c) { }; b()", runError: fmt.Errorf("expected 1 function arguments but received 0"), output: map[string]interface{}{"a": nil}},
		{script: "a = nil; func b(c) { }; b(a, a)", runError: fmt.Errorf("expected 1 function arguments but received 2"), output: map[string]interface{}{"a": nil}},

		{script: "a", input: map[string]interface{}{"a": testVarFunc}, runOutput: testVarFunc, output: map[string]interface{}{"a": testVarFunc}},
		{script: "a()", input: map[string]interface{}{"a": testVarFunc}, runOutput: int64(1), output: map[string]interface{}{"a": testVarFunc}},
		{script: "a", input: map[string]interface{}{"a": testVarFuncP}, runOutput: testVarFuncP, output: map[string]interface{}{"a": testVarFuncP}},
		// TOFIX:
		// {script: "a()", input: map[string]interface{}{"a": testVarFuncP}, runOutput: int64(1), output: map[string]interface{}{"a": testVarFuncP}},

		{script: "module a { func b() { return } }; a.b()", runOutput: nil},
		{script: "module a { func b() { return nil} }; a.b()", runOutput: nil},
		{script: "module a { func b() { return true} }; a.b()", runOutput: true},
		{script: "module a { func b() { return 1} }; a.b()", runOutput: int64(1)},
		{script: "module a { func b() { return 1.1} }; a.b()", runOutput: float64(1.1)},
		{script: "module a { func b() { return \"a\"} }; a.b()", runOutput: "a"},

		{script: "if true { module a { func b() { return } } }; a.b()", runOutput: nil},
		{script: "if true { module a { func b() { return nil} } }; a.b()", runOutput: nil},
		{script: "if true { module a { func b() { return true} } }; a.b()", runOutput: true},
		{script: "if true { module a { func b() { return 1} } }; a.b()", runOutput: int64(1)},
		{script: "if true { module a { func b() { return 1.1} } }; a.b()", runOutput: float64(1.1)},
		{script: "if true { module a { func b() { return \"a\"} } }; a.b()", runOutput: "a"},

		{script: "if true { module a { func b() { return 1} } }; a.b()", runOutput: int64(1)},

		{script: "a = 1; func b() { a = 2 }; b()", runOutput: int64(2), output: map[string]interface{}{"a": int64(2)}},
		{script: "b(a); a", input: map[string]interface{}{"a": int64(1), "b": func(c interface{}) { c = int64(2) }}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: "func b() { }; go b()", runOutput: nil},

		{script: "b(a)", input: map[string]interface{}{"a": reflect.Value{}, "b": func(c reflect.Value) bool { return c == reflect.Value{} }}, runOutput: true, output: map[string]interface{}{"a": reflect.Value{}}},
		{script: "b(a)", input: map[string]interface{}{"a": nil, "b": func(c interface{}) bool { return c == nil }}, runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "b(a)", input: map[string]interface{}{"a": true, "b": func(c bool) bool { return c == true }}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "b(a)", input: map[string]interface{}{"a": int32(1), "b": func(c int32) bool { return c == 1 }}, runOutput: true, output: map[string]interface{}{"a": int32(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": int64(1), "b": func(c int64) bool { return c == 1 }}, runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "b(a)", input: map[string]interface{}{"a": float32(1.1), "b": func(c float32) bool { return c == 1.1 }}, runOutput: true, output: map[string]interface{}{"a": float32(1.1)}},
		{script: "b(a)", input: map[string]interface{}{"a": float64(1.1), "b": func(c float64) bool { return c == 1.1 }}, runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "b(a)", input: map[string]interface{}{"a": "a", "b": func(c string) bool { return c == "a" }}, runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "x(a, b, c, d, e, f, g)", input: map[string]interface{}{"a": nil, "b": true, "c": int32(1), "d": int64(2), "e": float32(1.1), "f": float64(2.2), "g": "g",
			"x": func(a interface{}, b bool, c int32, d int64, e float32, f float64, g string) bool {
				return a == nil && b == true && c == 1 && d == 2 && e == 1.1 && f == 2.2 && g == "g"
			}}, runOutput: true, output: map[string]interface{}{"a": nil, "b": true, "c": int32(1), "d": int64(2), "e": float32(1.1), "f": float64(2.2), "g": "g"}},
		{script: "x(a, b, c, d, e, f, g)", input: map[string]interface{}{"a": nil, "b": true, "c": int32(1), "d": int64(2), "e": float32(1.1), "f": float64(2.2), "g": "g",
			"x": func(a interface{}, b bool, c int32, d int64, e float32, f float64, g string) (interface{}, bool, int32, int64, float32, float64, string) {
				return a, b, c, d, e, f, g
			}}, runOutput: []interface{}{nil, true, int32(1), int64(2), float32(1.1), float64(2.2), "g"}, output: map[string]interface{}{"a": nil, "b": true, "c": int32(1), "d": int64(2), "e": float32(1.1), "f": float64(2.2), "g": "g"}},

		{script: "a = nil; b(a)", input: map[string]interface{}{"b": func(c interface{}) bool { return c == nil }}, runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; b(a)", input: map[string]interface{}{"b": func(c bool) bool { return c == true }}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; b(a)", input: map[string]interface{}{"b": func(c int64) bool { return c == 1 }}, runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; b(a)", input: map[string]interface{}{"b": func(c float64) bool { return c == 1.1 }}, runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; b(a)", input: map[string]interface{}{"b": func(c string) bool { return c == "a" }}, runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "func b(c) { return c == nil }; b(a)", input: map[string]interface{}{"a": nil}, runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "func b(c) { return c == true }; b(a)", input: map[string]interface{}{"a": true}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "func b(c) { return c == 1 }; b(a)", input: map[string]interface{}{"a": int32(1)}, runOutput: true, output: map[string]interface{}{"a": int32(1)}},
		{script: "func b(c) { return c == 1 }; b(a)", input: map[string]interface{}{"a": int64(1)}, runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "func b(c) { return c == 1.1 }; b(a)", input: map[string]interface{}{"a": float32(1.1)}, runOutput: true, output: map[string]interface{}{"a": float32(1.1)}},
		{script: "func b(c) { return c == 1.1 }; b(a)", input: map[string]interface{}{"a": float64(1.1)}, runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "func b(c) { return c == \"a\" }; b(a)", input: map[string]interface{}{"a": "a"}, runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "a = nil; func b(c) { return c == nil }; b(a)", runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; func b(c) { return c == true }; b(a)", runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; func b(c) { return c == 1 }; b(a)", input: map[string]interface{}{"a": int64(1)}, runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; func b(c) { return c == 1.1 }; b(a)", input: map[string]interface{}{"a": float64(1.1)}, runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; func b(c) { return c == \"a\" }; b(a)", input: map[string]interface{}{"a": "a"}, runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{nil}, "b": func(c interface{}) bool { return c == nil }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{true}, "b": func(c interface{}) bool { return c == true }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{int32(1)}, "b": func(c interface{}) bool { return c == int32(1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{int32(1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{int64(1)}, "b": func(c interface{}) bool { return c == int64(1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": func(c interface{}) bool { return c == float32(1.1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{float32(1.1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": func(c interface{}) bool { return c == float64(1.1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{"a"}, "b": func(c interface{}) bool { return c == "a" }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{"a"}}},

		// TOFIX:
		//		{script: "b(a)",
		//			input: map[string]interface{}{"a": []bool{true, false, true}, "b": func(c ...bool) bool {
		//				for i := 0; i < len(c); i++ {
		//					if !c[i] {
		//						return false
		//					}
		//				}
		//				return true
		//			}}, runOutput: true, output: map[string]interface{}{"a": true}},

		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{true}, "b": func(c bool) bool { return c == true }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{int32(1)}, "b": func(c int32) bool { return c == int32(1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{int32(1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{int64(1)}, "b": func(c int64) bool { return c == int64(1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": func(c float32) bool { return c == float32(1.1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{float32(1.1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": func(c float64) bool { return c == float64(1.1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{script: "b(a[0])", input: map[string]interface{}{"a": []interface{}{"a"}, "b": func(c string) bool { return c == "a" }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{"a"}}},

		{script: "a = [nil]; b(a[0])", input: map[string]interface{}{"b": func(c interface{}) bool { return c == nil }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = [true]; b(a[0])", input: map[string]interface{}{"b": func(c bool) bool { return c == true }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = [1]; b(a[0])", input: map[string]interface{}{"b": func(c int64) bool { return c == int64(1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1.1]; b(a[0])", input: map[string]interface{}{"b": func(c float64) bool { return c == float64(1.1) }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{script: "a = [\"a\"]; b(a[0])", input: map[string]interface{}{"b": func(c string) bool { return c == "a" }}, runOutput: true, output: map[string]interface{}{"a": []interface{}{"a"}}},

		{script: "a = [nil]; func b(c) { c == nil }; b(a[0])", runOutput: true, output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = [true]; func b(c) { c == true }; b(a[0])", runOutput: true, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = [1]; func b(c) { c == 1 }; b(a[0])", runOutput: true, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1.1]; func b(c) { c == 1.1 }; b(a[0])", runOutput: true, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{script: "a = [\"a\"]; func b(c) { c == \"a\" }; b(a[0])", runOutput: true, output: map[string]interface{}{"a": []interface{}{"a"}}},

		{script: "a = nil; b = func (d) { return d == nil }; b(a)", runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; b = func (d) { return d == true }; b(a)", runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; b = func (d) { return d == 1 }; b(a)", runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; b = func (d) { return d == 1.1 }; b(a)", runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; b = func (d) { return d == \"a\" }; b(a)", runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "a = nil; b = func c(d) { return d == nil }; b(a)", runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; b = func c(d) { return d == true }; b(a)", runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; b = func c(d) { return d == 1 }; b(a)", runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; b = func c(d) { return d == 1.1 }; b(a)", runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; b = func c(d) { return d == \"a\" }; b(a)", runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "a = nil; b = func c(d) { return d == nil }; c(a)", runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; b = func c(d) { return d == true }; c(a)", runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; b = func c(d) { return d == 1 }; c(a)", runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; b = func c(d) { return d == 1.1 }; c(a)", runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; b = func c(d) { return d == \"a\" }; c(a)", runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "a = nil; func b() { return func c(d) { d == nil } }; e = b(); e(a)", runOutput: true, output: map[string]interface{}{"a": nil}},
		{script: "a = true; func b() { return func c(d) { d == true } }; e = b(); e(a)", runOutput: true, output: map[string]interface{}{"a": true}},
		{script: "a = 1; func b() { return func c(d) { d == 1 } }; e = b(); e(a)", runOutput: true, output: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1.1; func b() { return func c(d) { d == 1.1 } }; e = b(); e(a)", runOutput: true, output: map[string]interface{}{"a": float64(1.1)}},
		{script: "a = \"a\"; func b() { return func c(d) { d == \"a\" } }; e = b(); e(a)", runOutput: true, output: map[string]interface{}{"a": "a"}},

		{script: "a = func () { return nil }; func b(c) { return c() }; b(a)", runOutput: nil},
		{script: "a = func () { return true }; func b(c) { return c() }; b(a)", runOutput: true},
		{script: "a = func () { return 1 }; func b(c) { return c() }; b(a)", runOutput: int64(1)},
		{script: "a = func () { return 1.1 }; func b(c) { return c() }; b(a)", runOutput: float64(1.1)},
		{script: "a = func () { return \"a\" }; func b(c) { return c() }; b(a)", runOutput: "a"},

		{script: "a = [nil]; func c(d) { return d[0] }; c(a)", runOutput: nil},
		{script: "a = [true]; func c(d) { return d[0] }; c(a)", runOutput: true},
		{script: "a = [1]; func c(d) { return d[0] }; c(a)", runOutput: int64(1)},
		{script: "a = [1.1]; func c(d) { return d[0] }; c(a)", runOutput: float64(1.1)},
		{script: "a = [\"a\"]; func c(d) { return d[0] }; c(a)", runOutput: "a"},

		{script: "a = {\"b\": nil}; func c(d) { return d.b }; c(a)", runOutput: nil},
		{script: "a = {\"b\": true}; func c(d) { return d.b }; c(a)", runOutput: true},
		{script: "a = {\"b\": 1}; func c(d) { return d.b }; c(a)", runOutput: int64(1)},
		{script: "a = {\"b\": 1.1}; func c(d) { return d.b }; c(a)", runOutput: float64(1.1)},
		{script: "a = {\"b\": \"a\"}; func c(d) { return d.b }; c(a)", runOutput: "a"},
	}
	runTests(t, tests)
}

func TestFunctionsInArraysAndMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = [func () { return nil }]; a[0]()", runOutput: nil},
		{script: "a = [func () { return true }]; a[0]()", runOutput: true},
		{script: "a = [func () { return 1 }]; a[0]()", runOutput: int64(1)},
		{script: "a = [func () { return 1.1 }]; a[0]()", runOutput: float64(1.1)},
		{script: "a = [func () { return \"a\" }]; a[0]()", runOutput: "a"},

		{script: "a = [func () { return nil }]; b = a[0]; b()", runOutput: nil},
		{script: "a = [func () { return true }]; b = a[0]; b()", runOutput: true},
		{script: "a = [func () { return 1 }]; b = a[0]; b()", runOutput: int64(1)},
		{script: "a = [func () { return 1.1 }]; b = a[0]; b()", runOutput: float64(1.1)},
		{script: "a = [func () { return \"a\" }]; b = a[0]; b()", runOutput: "a"},

		{script: "a = [func () { return nil}]; func b(c) { return c() }; b(a[0])", runOutput: nil},
		{script: "a = [func () { return true}]; func b(c) { return c() }; b(a[0])", runOutput: true},
		{script: "a = [func () { return 1}]; func b(c) { return c() }; b(a[0])", runOutput: int64(1)},
		{script: "a = [func () { return 1.1}]; func b(c) { return c() }; b(a[0])", runOutput: float64(1.1)},
		{script: "a = [func () { return \"a\"}]; func b(c) { return c() }; b(a[0])", runOutput: "a"},

		{script: "a = {\"b\": func () { return nil }}; a[\"b\"]()", runOutput: nil},
		{script: "a = {\"b\": func () { return true }}; a[\"b\"]()", runOutput: true},
		{script: "a = {\"b\": func () { return 1 }}; a[\"b\"]()", runOutput: int64(1)},
		{script: "a = {\"b\": func () { return 1.1 }}; a[\"b\"]()", runOutput: float64(1.1)},
		{script: "a = {\"b\": func () { return \"a\" }}; a[\"b\"]()", runOutput: "a"},

		{script: "a = {\"b\": func () { return nil }}; a.b()", runOutput: nil},
		{script: "a = {\"b\": func () { return true }}; a.b()", runOutput: true},
		{script: "a = {\"b\": func () { return 1 }}; a.b()", runOutput: int64(1)},
		{script: "a = {\"b\": func () { return 1.1 }}; a.b()", runOutput: float64(1.1)},
		{script: "a = {\"b\": func () { return \"a\" }}; a.b()", runOutput: "a"},

		{script: "a = {\"b\": func () { return nil }}; func c(d) { return d() }; c(a.b)", runOutput: nil},
		{script: "a = {\"b\": func () { return true }}; func c(d) { return d() }; c(a.b)", runOutput: true},
		{script: "a = {\"b\": func () { return 1 }}; func c(d) { return d() }; c(a.b)", runOutput: int64(1)},
		{script: "a = {\"b\": func () { return 1.1 }}; func c(d) { return d() }; c(a.b)", runOutput: float64(1.1)},
		{script: "a = {\"b\": func () { return \"a\" }}; func c(d) { return d() }; c(a.b)", runOutput: "a"},
	}
	runTests(t, tests)
}
