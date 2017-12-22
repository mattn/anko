package vm

import (
	"os"
	"reflect"
	"testing"
)

func TestReturns(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "return", runOutput: nil},
		{script: "return nil", runOutput: nil},
		{script: "return true", runOutput: true},
		{script: "return 1", runOutput: int64(1)},
		{script: "return 1.1", runOutput: float64(1.1)},
		{script: "return \"a\"", runOutput: "a"},

		{script: "func aFunc() {return}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return}; a = aFunc()", runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: "func aFunc() {return nil}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return true}; aFunc()", runOutput: true},
		{script: "func aFunc() {return 1}; aFunc()", runOutput: int64(1)},
		{script: "func aFunc() {return 1.1}; aFunc()", runOutput: float64(1.1)},
		{script: "func aFunc() {return \"a\"}; aFunc()", runOutput: "a"},

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

		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: []interface{}{nil, nil}, output: map[string]interface{}{"a": reflect.Value{}}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: []interface{}{nil, nil}, output: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: []interface{}{true, true}, output: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": int32(1)}, runOutput: []interface{}{int32(1), int32(1)}, output: map[string]interface{}{"a": int32(1)}},

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
		{script: "a", input: map[string]interface{}{"a": testVarFunc}, runOutput: testVarFunc, output: map[string]interface{}{"a": testVarFunc}},
		{script: "a()", input: map[string]interface{}{"a": testVarFunc}, runOutput: int64(1), output: map[string]interface{}{"a": testVarFunc}},
		{script: "a", input: map[string]interface{}{"a": testVarFuncP}, runOutput: testVarFuncP, output: map[string]interface{}{"a": testVarFuncP}},
		// TOFIX:
		// {script: "a()", input: map[string]interface{}{"a": testVarFuncP}, runOutput: int64(1), output: map[string]interface{}{"a": testVarFuncP}},

		{script: "module a { func b() { return } }; a.b()", runOutput: nil},
		{script: "module a { func b() { return nil} }; a.b()", runOutput: nil},
		{script: "module a { func b() { return 1} }; a.b()", runOutput: int64(1)},
		{script: "module a { func b() { return 1.1} }; a.b()", runOutput: float64(1.1)},
		{script: "module a { func b() { return \"a\"} }; a.b()", runOutput: "a"},
	}
	runTests(t, tests)
}
