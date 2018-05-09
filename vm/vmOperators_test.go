package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBasicOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	testInput1 := map[string]interface{}{"b": func() {}}
	tests := []Test{
		{Script: `]`, ParseError: fmt.Errorf("syntax error")},
		{Script: `1 = 2`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a, = 1, 2`, ParseError: fmt.Errorf("syntax error")},
		// TOFIX: should be syntax error
		{Script: `,b = 1, 2`, RunOutput: int64(2), Output: map[string]interface{}{"b": int64(1)}},
		// TOFIX: should be error of some kind
		{Script: `a, b = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a, b = 1,`, ParseError: fmt.Errorf("syntax error")},
		// TOFIX: should be syntax error
		{Script: `a, b = ,2`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},

		{Script: `2 + 1`, RunOutput: int64(3)},
		{Script: `2 - 1`, RunOutput: int64(1)},
		{Script: `2 * 1`, RunOutput: int64(2)},
		{Script: `2 / 1`, RunOutput: float64(2)},
		{Script: `2.1 + 1.1`, RunOutput: float64(3.2)},
		{Script: `2.1 - 1.1`, RunOutput: float64(1)},
		{Script: `2.1 * 2.0`, RunOutput: float64(4.2)},
		{Script: `6.5 / 2.0`, RunOutput: float64(3.25)},

		{Script: `a + b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(3)},
		{Script: `a - b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(1)},
		{Script: `a * b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(2)},
		{Script: `a / b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: float64(2)},
		{Script: `a + b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(3.2)},
		{Script: `a - b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(1)},
		{Script: `a * b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, RunOutput: float64(4.2)},
		{Script: `a / b`, Input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, RunOutput: float64(3.25)},

		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": "b"}, RunOutput: "ab"},
		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": int64(1)}, RunOutput: "a1"},
		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": float64(1.1)}, RunOutput: "a1.1"},
		{Script: `a + z`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `z + b`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},

		{Script: `a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `a = nil; a = nil`, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `a = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a = b; b = nil`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a = true`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `a = false`, Input: map[string]interface{}{"a": true}, RunOutput: false, Output: map[string]interface{}{"a": false}},

		{Script: `a = b`, Input: map[string]interface{}{"b": reflect.Value{}}, RunOutput: reflect.Value{}, Output: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{Script: `a = b`, Input: map[string]interface{}{"b": struct{}{}}, RunOutput: struct{}{}, Output: map[string]interface{}{"a": struct{}{}, "b": struct{}{}}},
		{Script: `a = b`, Input: testInput1, RunOutput: testInput1["b"], Output: map[string]interface{}{"a": testInput1["b"], "b": testInput1["b"]}},

		{Script: `a = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a = b`, Input: map[string]interface{}{"b": true}, RunOutput: true, Output: map[string]interface{}{"a": true, "b": true}},
		{Script: `a = b`, Input: map[string]interface{}{"b": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": float64(2.1)}, RunOutput: float64(2.1), Output: map[string]interface{}{"a": float64(2.1), "b": float64(2.1)}},
		{Script: `a = b`, Input: map[string]interface{}{"b": "a"}, RunOutput: "a", Output: map[string]interface{}{"a": "a", "b": "a"}},
		{Script: `a = b`, Input: map[string]interface{}{"b": 'a'}, RunOutput: 'a', Output: map[string]interface{}{"a": 'a', "b": 'a'}},

		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValue}, RunOutput: testVarValue, Output: map[string]interface{}{"a": testVarValue, "b": testVarValue}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarBoolP}, RunOutput: testVarBoolP, Output: map[string]interface{}{"a": testVarBoolP, "b": testVarBoolP}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarInt32P}, RunOutput: testVarInt32P, Output: map[string]interface{}{"a": testVarInt32P, "b": testVarInt32P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarInt64P}, RunOutput: testVarInt64P, Output: map[string]interface{}{"a": testVarInt64P, "b": testVarInt64P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFloat32P}, RunOutput: testVarFloat32P, Output: map[string]interface{}{"a": testVarFloat32P, "b": testVarFloat32P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFloat64P}, RunOutput: testVarFloat64P, Output: map[string]interface{}{"a": testVarFloat64P, "b": testVarFloat64P}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarStringP}, RunOutput: testVarStringP, Output: map[string]interface{}{"a": testVarStringP, "b": testVarStringP}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarFuncP}, RunOutput: testVarFuncP, Output: map[string]interface{}{"a": testVarFuncP, "b": testVarFuncP}},

		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueBool}, RunOutput: testVarValueBool, Output: map[string]interface{}{"a": testVarValueBool, "b": testVarValueBool}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueInt32}, RunOutput: testVarValueInt32, Output: map[string]interface{}{"a": testVarValueInt32, "b": testVarValueInt32}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueInt64}, RunOutput: testVarValueInt64, Output: map[string]interface{}{"a": testVarValueInt64, "b": testVarValueInt64}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueFloat32}, RunOutput: testVarValueFloat32, Output: map[string]interface{}{"a": testVarValueFloat32, "b": testVarValueFloat32}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueFloat64}, RunOutput: testVarValueFloat64, Output: map[string]interface{}{"a": testVarValueFloat64, "b": testVarValueFloat64}},
		{Script: `a = b`, Input: map[string]interface{}{"b": testVarValueString}, RunOutput: testVarValueString, Output: map[string]interface{}{"a": testVarValueString, "b": testVarValueString}},

		{Script: `a = 1, 2`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a, b = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a.c, b = 1, 2`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a, b.c = 1, 2`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `a, b = 1, 2`, Input: map[string]interface{}{"a": int64(3)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `a, b, c = 1, 2, 3`, Input: map[string]interface{}{"a": int64(3)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{Script: `a, b = [1, 2], [3, 4]`, RunOutput: []interface{}{int64(3), int64(4)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": []interface{}{int64(3), int64(4)}}},

		{Script: `y = z`, RunError: fmt.Errorf("undefined symbol 'z'")},
		{Script: `z.y.x = 1`, RunError: fmt.Errorf("undefined symbol 'z'")},

		{Script: `c = a + b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(3), Output: map[string]interface{}{"c": int64(3)}},
		{Script: `c = a - b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"c": int64(1)}},
		{Script: `c = a * b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"c": int64(2)}},
		{Script: `c = a / b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: float64(2), Output: map[string]interface{}{"c": float64(2)}},
		{Script: `c = a + b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(3.2), Output: map[string]interface{}{"c": float64(3.2)}},
		{Script: `c = a - b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(1), Output: map[string]interface{}{"c": float64(1)}},
		{Script: `c = a * b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, RunOutput: float64(4.2), Output: map[string]interface{}{"c": float64(4.2)}},
		{Script: `c = a / b`, Input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, RunOutput: float64(3.25), Output: map[string]interface{}{"c": float64(3.25)}},

		{Script: `a = nil; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = false; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = true; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1.5; a++`, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a = "1"; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = "a"; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a = nil; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a = false; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a = true; a--`, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a = 2; a--`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2.5; a--`, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a = "2"; a--`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "a"; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},

		{Script: `a++`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a++`, Input: map[string]interface{}{"a": false}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a++`, Input: map[string]interface{}{"a": true}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": float32(1.5)}, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a++`, Input: map[string]interface{}{"a": float64(1.5)}, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a++`, Input: map[string]interface{}{"a": "2"}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a++`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a--`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": false}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": true}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a--`, Input: map[string]interface{}{"a": int32(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": float32(2.5)}, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a--`, Input: map[string]interface{}{"a": float64(2.5)}, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a--`, Input: map[string]interface{}{"a": "2"}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},

		{Script: `1++`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},
		{Script: `1--`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},
		{Script: `z++`, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `z--`, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `!(1++)`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},

		{Script: `a += 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a -= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a *= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `a /= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: float64(1), Output: map[string]interface{}{"a": float64(1)}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(3.1), Output: map[string]interface{}{"a": float64(3.1)}},
		{Script: `a -= 1`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `a *= 2`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(4.2), Output: map[string]interface{}{"a": float64(4.2)}},
		{Script: `a /= 2`, Input: map[string]interface{}{"a": 6.5}, RunOutput: float64(3.25), Output: map[string]interface{}{"a": float64(3.25)}},

		{Script: `a ** 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a ** 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: float64(8), Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a &= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a &= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a &= 1`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a &= 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a |= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a |= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a |= 1`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a |= 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a << 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >> 2`, Input: map[string]interface{}{"a": int64(8)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(8)}},
		{Script: `a << 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >> 2`, Input: map[string]interface{}{"a": float64(8)}, RunOutput: int64(2), Output: map[string]interface{}{"a": float64(8)}},

		{Script: `a % 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a % 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a % 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(0), Output: map[string]interface{}{"a": float64(2.1)}},
		{Script: `a % 3`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `a * 4`, Input: map[string]interface{}{"a": "a"}, RunOutput: "aaaa", Output: map[string]interface{}{"a": "a"}},
		{Script: `a * 4.0`, Input: map[string]interface{}{"a": "a"}, RunOutput: float64(0), Output: map[string]interface{}{"a": "a"}},

		{Script: `-a`, Input: map[string]interface{}{"a": nil}, RunOutput: float64(-0), Output: map[string]interface{}{"a": nil}},
		{Script: `-a`, Input: map[string]interface{}{"a": "a"}, RunOutput: float64(-0), Output: map[string]interface{}{"a": "a"}},
		{Script: `-a`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(-2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `-a`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: float64(-2.1), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `^a`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(-1), Output: map[string]interface{}{"a": nil}},
		{Script: `^a`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(-1), Output: map[string]interface{}{"a": "a"}},
		{Script: `^a`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(-3), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `^a`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(-3), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `!true`, RunOutput: false},
		{Script: `!1`, RunOutput: false},
	}
	RunTests(t, tests, nil)
}

func TestComparisonOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `a == 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a == 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a == nil`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": nil}, RunOutput: false, Output: map[string]interface{}{"a": nil}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": nil}, RunOutput: false, Output: map[string]interface{}{"a": nil}},

		{Script: `1 == 1.0`, RunOutput: true},
		{Script: `1 != 1.0`, RunOutput: false},
		{Script: `"a" != "a"`, RunOutput: false},

		{Script: `a > 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a > 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a >= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a >= 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `1 == 1 && 1  == 1`, RunOutput: true},
		{Script: `1 == 1 && 1  == 2`, RunOutput: false},
		{Script: `1 == 2 && 1  == 1`, RunOutput: false},
		{Script: `1 == 2 && 1  == 2`, RunOutput: false},
		{Script: `1 == 1 || 1  == 1`, RunOutput: true},
		{Script: `1 == 1 || 1  == 2`, RunOutput: true},
		{Script: `1 == 2 || 1  == 1`, RunOutput: true},
		{Script: `1 == 2 || 1  == 2`, RunOutput: false},

		{Script: `true == "true"`, RunOutput: true},
		{Script: `true == "TRUE"`, RunOutput: true},
		{Script: `true == "True"`, RunOutput: true},
		{Script: `true == "false"`, RunOutput: false},
		{Script: `true == "foo"`, RunOutput: false},
		{Script: `false == "false"`, RunOutput: true},
		{Script: `false == "FALSE"`, RunOutput: true},
		{Script: `false == "False"`, RunOutput: true},
		{Script: `false == "true"`, RunOutput: false},
		{Script: `false == "foo"`, RunOutput: false},

		{Script: `0 == "0"`, RunOutput: true},
		{Script: `"1.0" == 1`, RunOutput: true},
		{Script: `1 == "1"`, RunOutput: true},
		{Script: `0.0 == "0"`, RunOutput: true},
		{Script: `0.0 == "0.0"`, RunOutput: true},
		{Script: `1.0 == "1.0"`, RunOutput: true},
		{Script: `1.2 == "1.2"`, RunOutput: true},
		{Script: `"7" == 7.2`, RunOutput: false},
		{Script: `1.2 == "1"`, RunOutput: false},
		{Script: `"1.1" == 1`, RunOutput: false},
		{Script: `0 == "1"`, RunOutput: false},

		{Script: `a == b`, Input: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}, RunOutput: true, Output: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{Script: `a == b`, Input: map[string]interface{}{"a": reflect.Value{}, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": reflect.Value{}, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": reflect.Value{}}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": reflect.Value{}}},

		{Script: `a == b`, Input: map[string]interface{}{"a": nil, "b": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a == b`, Input: map[string]interface{}{"a": nil, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": nil, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": nil}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": nil}},

		{Script: `a == b`, Input: map[string]interface{}{"a": false, "b": false}, RunOutput: true, Output: map[string]interface{}{"a": false, "b": false}},
		{Script: `a == b`, Input: map[string]interface{}{"a": false, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": false, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": false}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": false}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": true}, RunOutput: true, Output: map[string]interface{}{"a": true, "b": true}},

		{Script: `a == b`, Input: map[string]interface{}{"a": int32(1), "b": int32(1)}, RunOutput: true, Output: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(1), "b": int32(2)}, RunOutput: false, Output: map[string]interface{}{"a": int32(1), "b": int32(2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(2), "b": int32(1)}, RunOutput: false, Output: map[string]interface{}{"a": int32(2), "b": int32(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(2), "b": int32(2)}, RunOutput: true, Output: map[string]interface{}{"a": int32(2), "b": int32(2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": int64(1), "b": int64(1)}, RunOutput: true, Output: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(1), "b": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2), "b": int64(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(2), "b": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2), "b": int64(2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}, RunOutput: true, Output: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}, RunOutput: false, Output: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}, RunOutput: false, Output: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}, RunOutput: true, Output: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}, RunOutput: true, Output: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": 'a', "b": 'a'}, RunOutput: true, Output: map[string]interface{}{"a": 'a', "b": 'a'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'a', "b": 'b'}, RunOutput: false, Output: map[string]interface{}{"a": 'a', "b": 'b'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'b', "b": 'a'}, RunOutput: false, Output: map[string]interface{}{"a": 'b', "b": 'a'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'b', "b": 'b'}, RunOutput: true, Output: map[string]interface{}{"a": 'b', "b": 'b'}},

		{Script: `a == b`, Input: map[string]interface{}{"a": "a", "b": "a"}, RunOutput: true, Output: map[string]interface{}{"a": "a", "b": "a"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "a", "b": "b"}, RunOutput: false, Output: map[string]interface{}{"a": "a", "b": "b"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "b", "b": "a"}, RunOutput: false, Output: map[string]interface{}{"a": "b", "b": "a"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "b", "b": "b"}, RunOutput: true, Output: map[string]interface{}{"a": "b", "b": "b"}},

		{Script: `b = "\"a\""; a == b`, Input: map[string]interface{}{"a": `"a"`}, RunOutput: true, Output: map[string]interface{}{"a": `"a"`, "b": `"a"`}},

		{Script: `a = "test"; a == "test"`, RunOutput: true},
		{Script: `a = "test"; a[0:1] == "t"`, RunOutput: true},
		{Script: `a = "test"; a[0:2] == "te"`, RunOutput: true},
		{Script: `a = "test"; a[1:3] == "es"`, RunOutput: true},
		{Script: `a = "test"; a[0:4] == "test"`, RunOutput: true},

		{Script: `a = "a b"; a[1] == ' '`, RunOutput: true},
		{Script: `a = "test"; a[0] == 't'`, RunOutput: true},
		{Script: `a = "test"; a[1] == 'e'`, RunOutput: true},
		{Script: `a = "test"; a[3] == 't'`, RunOutput: true},

		{Script: `a = "a b"; a[1] != ' '`, RunOutput: false},
		{Script: `a = "test"; a[0] != 't'`, RunOutput: false},
		{Script: `a = "test"; a[1] != 'e'`, RunOutput: false},
		{Script: `a = "test"; a[3] != 't'`, RunOutput: false},
	}
	RunTests(t, tests, nil)
}
