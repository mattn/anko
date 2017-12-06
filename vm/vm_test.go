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
	input      map[string]interface{}
	runError   error
	runOutput  interface{}
	ouput      map[string]interface{}
}

func TestBasicOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	testInput1 := map[string]interface{}{"b": func() {}}
	tests := []testStruct{
		{script: "]", parseError: fmt.Errorf("syntax error"), runOutput: nil},
		{script: "1 = 2", runError: fmt.Errorf("Invalid operation"), runOutput: nil},

		{script: "2 + 1", runOutput: int64(3)},
		{script: "2 - 1", runOutput: int64(1)},
		{script: "2 * 1", runOutput: int64(2)},
		{script: "2 / 1", runOutput: float64(2)},
		{script: "2.1 + 1.1", runOutput: float64(3.2)},
		{script: "2.1 - 1.1", runOutput: float64(1)},
		{script: "2.1 * 2.0", runOutput: float64(4.2)},
		{script: "6.5 / 2.0", runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3)},
		{script: "a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1)},
		{script: "a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2)},
		{script: "a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2)},
		{script: "a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2)},
		{script: "a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1)},
		{script: "a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2)},
		{script: "a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": "a", "b": "b"}, runOutput: "ab"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": int64(1)}, runOutput: "a1"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": float64(1.1)}, runOutput: "a1.1"},
		{script: "a + z", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z + b", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = nil; a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = b", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = b; b = nil", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = true", runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "a = false", input: map[string]interface{}{"a": true}, runOutput: false, ouput: map[string]interface{}{"a": false}},

		{script: "a = b", input: map[string]interface{}{"b": reflect.Value{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{script: "a = b", input: map[string]interface{}{"b": struct{}{}}, runOutput: struct{}{}, ouput: map[string]interface{}{"a": struct{}{}, "b": struct{}{}}},
		{script: "a = b", input: testInput1, runOutput: testInput1["b"], ouput: map[string]interface{}{"a": testInput1["b"], "b": testInput1["b"]}},

		{script: "a = b", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = b", input: map[string]interface{}{"b": true}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": float64(2.1)}, runOutput: float64(2.1), ouput: map[string]interface{}{"a": float64(2.1), "b": float64(2.1)}},
		{script: "a = b", input: map[string]interface{}{"b": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a", "b": "a"}},

		{script: "a, b = 1, 2", input: map[string]interface{}{"a": int64(3)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "a, b, c = 1, 2, 3", input: map[string]interface{}{"a": int64(3)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
		{script: "a, b = [1, 2], [3, 4]", runOutput: []interface{}{int64(3), int64(4)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": []interface{}{int64(3), int64(4)}}},

		{script: "y = z", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z.y.x = 1", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "c = a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3), ouput: map[string]interface{}{"c": int64(3)}},
		{script: "c = a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"c": int64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"c": int64(2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2), ouput: map[string]interface{}{"c": float64(2)}},
		{script: "c = a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2), ouput: map[string]interface{}{"c": float64(3.2)}},
		{script: "c = a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1), ouput: map[string]interface{}{"c": float64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2), ouput: map[string]interface{}{"c": float64(4.2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25), ouput: map[string]interface{}{"c": float64(3.25)}},

		{script: "a++", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a--", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a++", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},

		{script: "1++", runError: fmt.Errorf("Invalid operation"), runOutput: nil},
		{script: "1--", runError: fmt.Errorf("Invalid operation"), runOutput: nil},
		{script: "z++", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},
		{script: "z--", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: nil},

		{script: "a += 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a -= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), ouput: map[string]interface{}{"a": int64(4)}},
		{script: "a /= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1), ouput: map[string]interface{}{"a": float64(1)}},
		{script: "a += 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a -= 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": 2.1}, runOutput: float64(4.2), ouput: map[string]interface{}{"a": float64(4.2)}},
		{script: "a /= 2", input: map[string]interface{}{"a": 6.5}, runOutput: float64(3.25), ouput: map[string]interface{}{"a": float64(3.25)}},

		{script: "a ** 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a ** 3", input: map[string]interface{}{"a": float64(2)}, runOutput: float64(8), ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a &= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a &= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a |= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a |= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a << 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": int64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(8)}},
		{script: "a << 2", input: map[string]interface{}{"a": float64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": float64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(8)}},

		{script: "a % 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": float64(2.1)}},
		{script: "a % 3", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(2.1)}},

		{script: "-a", input: map[string]interface{}{"a": nil}, runOutput: float64(-0), ouput: map[string]interface{}{"a": nil}},
		{script: "-a", input: map[string]interface{}{"a": "a"}, runOutput: float64(-0), ouput: map[string]interface{}{"a": "a"}},
		{script: "-a", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(-2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "-a", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(-2.1), ouput: map[string]interface{}{"a": float64(2.1)}},
		{script: "^a", input: map[string]interface{}{"a": nil}, runOutput: int64(-1), ouput: map[string]interface{}{"a": nil}},
		{script: "^a", input: map[string]interface{}{"a": "a"}, runOutput: int64(-1), ouput: map[string]interface{}{"a": "a"}},
		{script: "^a", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(-3), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "^a", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(-3), ouput: map[string]interface{}{"a": float64(2.1)}},

		{script: "a * 4", input: map[string]interface{}{"a": "a"}, runOutput: "aaaa", ouput: map[string]interface{}{"a": "a"}},
		{script: "a * 4.0", input: map[string]interface{}{"a": "a"}, runOutput: float64(0), ouput: map[string]interface{}{"a": "a"}},
		{script: "!true", runOutput: false},
		{script: "!1", runOutput: false},
	}
	runTests(t, tests)
}

func TestComparisonOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a == 1", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a == 1", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == nil", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},
		{script: "a == 2.0", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},

		{script: "1 == 1.0", runOutput: true},
		{script: "1 != 1.0", runOutput: false},
		{script: "\"a\" != \"a\"", runOutput: false},

		{script: "a > 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a > 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a >= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 3", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 3", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >= 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a <= 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a >= 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 3", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 3", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >= 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a <= 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "1 == 1 && 1  == 1", runOutput: true},
		{script: "1 == 1 && 1  == 2", runOutput: false},
		{script: "1 == 2 && 1  == 1", runOutput: false},
		{script: "1 == 2 && 1  == 2", runOutput: false},
		{script: "1 == 1 || 1  == 1", runOutput: true},
		{script: "1 == 1 || 1  == 2", runOutput: true},
		{script: "1 == 2 || 1  == 1", runOutput: true},
		{script: "1 == 2 || 1  == 2", runOutput: false},

		{script: `true == "true"`, runOutput: true},
		{script: `true == "TRUE"`, runOutput: true},
		{script: `true == "True"`, runOutput: true},
		{script: `true == "false"`, runOutput: false},
		{script: `true == "foo"`, runOutput: false},
		{script: `false == "false"`, runOutput: true},
		{script: `false == "FALSE"`, runOutput: true},
		{script: `false == "False"`, runOutput: true},
		{script: `false == "true"`, runOutput: false},
		{script: `false == "foo"`, runOutput: false},

		{script: `0 == "0"`, runOutput: true},
		{script: `"1.0" == 1`, runOutput: true},
		{script: `1 == "1"`, runOutput: true},
		{script: `0.0 == "0"`, runOutput: true},
		{script: `0.0 == "0.0"`, runOutput: true},
		{script: `1.0 == "1.0"`, runOutput: true},
		{script: `1.2 == "1.2"`, runOutput: true},
		{script: `"7" == 7.2`, runOutput: true},
		{script: `1.2 == "1"`, runOutput: false},
		{script: `"1.1" == 1`, runOutput: false},
		{script: `0 == "1"`, runOutput: false},
	}
	runTests(t, tests)
}

func TestIf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "if true {}", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if true {}", input: map[string]interface{}{"a": true}, runOutput: nil, ouput: map[string]interface{}{"a": true}},
		{script: "if true {}", input: map[string]interface{}{"a": int64(1)}, runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "if true {}", input: map[string]interface{}{"a": float64(1.1)}, runOutput: nil, ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "if true {}", input: map[string]interface{}{"a": "a"}, runOutput: nil, ouput: map[string]interface{}{"a": "a"}},

		{script: "if true {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if true {a = true}", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "if true {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "if true {a = 1.1}", input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "if true {a = \"a\"}", input: map[string]interface{}{"a": int64(2)}, runOutput: "a", ouput: map[string]interface{}{"a": "a"}},

		{script: "if a == 1 {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 2 {a = 1}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "if a == 1 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 2 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},

		{script: "if a == 1 {a = 1} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "if a == 2 {a = 1} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 1 {a = 1} else if a == 2 {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else {a = 4}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), ouput: map[string]interface{}{"a": int64(4)}},

		{script: "if a == 1 {a = 1} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if a == 2 {a = nil} else {a = 3}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = nil} else if a == 3 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "if a == 1 {a = 1} else if a == 2 {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},

		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(5), ouput: map[string]interface{}{"a": int64(5)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = nil}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = 4} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), ouput: map[string]interface{}{"a": int64(4)}},
		{script: "if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = nil} else {a = 5}", input: map[string]interface{}{"a": int64(2)}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
	}
	runTests(t, tests)
}

func TestReturns(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
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

		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": reflect.Value{}}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: []interface{}{interface{}(nil), interface{}(nil)}, ouput: map[string]interface{}{"a": reflect.Value{}}},

		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a"}},

		{script: "func a(x) { return x}; a(nil)", runOutput: nil},
		{script: "func a(x) { return x}; a(true)", runOutput: true},
		{script: "func a(x) { return x}; a(1)", runOutput: int64(1)},
		{script: "func a(x) { return x}; a(1.1)", runOutput: float64(1.1)},
		{script: "func a(x) { return x}; a(\"a\")", runOutput: "a"},

		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": true}, runOutput: nil, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": int64(1)}, runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": float64(1.1)}, runOutput: nil, ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a}; for {aFunc(); break}", input: map[string]interface{}{"a": "a"}, runOutput: nil, ouput: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {for {return a}}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {for {if true {return a}}}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {return nil, nil}; a, b = aFunc()", runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "func aFunc() {return true, false}; a, b = aFunc()", runOutput: false, ouput: map[string]interface{}{"a": true, "b": false}},
		{script: "func aFunc() {return 1, 2}; a, b = aFunc()", runOutput: int64(2), ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "func aFunc() {return 1.1, 2.2}; a, b = aFunc()", runOutput: float64(2.2), ouput: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{script: "func aFunc() {return \"a\", \"b\"}; a, b = aFunc()", runOutput: "b", ouput: map[string]interface{}{"a": "a", "b": "b"}},
	}
	runTests(t, tests)
}

func TestArrays(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = []", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []interface{}{}}},
		{script: "a = [nil]", runOutput: []interface{}{interface{}(nil)}, ouput: map[string]interface{}{"a": []interface{}{interface{}(nil)}}},
		{script: "a = [true]", runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = [\"test\"]", runOutput: []interface{}{"test"}, ouput: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = [1]", runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1.1]", runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a = [[]]", runOutput: []interface{}{[]interface{}{}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{}}}},
		{script: "a = [[nil]]", runOutput: []interface{}{[]interface{}{interface{}(nil)}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{interface{}(nil)}}}},
		{script: "a = [[true]]", runOutput: []interface{}{[]interface{}{true}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{true}}}},
		{script: "a = [[\"test\"]]", runOutput: []interface{}{[]interface{}{"test"}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{"test"}}}},
		{script: "a = [[1]]", runOutput: []interface{}{[]interface{}{int64(1)}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1)}}}},
		{script: "a = [[1.1]]", runOutput: []interface{}{[]interface{}{float64(1.1)}}, ouput: map[string]interface{}{"a": []interface{}{[]interface{}{float64(1.1)}}}},

		{script: "a = []; a += nil", runOutput: []interface{}{nil}, ouput: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = []; a += true", runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = []; a += \"test\"", runOutput: []interface{}{"test"}, ouput: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = []; a += 1", runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = []; a += 1.1", runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a = []; a += []", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []interface{}{}}},
		{script: "a = []; a += [nil]", runOutput: []interface{}{nil}, ouput: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = []; a += [true]", runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = []; a += [\"test\"]", runOutput: []interface{}{"test"}, ouput: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = []; a += [1]", runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = []; a += [1.1]", runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a", input: map[string]interface{}{"a": []bool{}}, runOutput: []bool{}, ouput: map[string]interface{}{"a": []bool{}}},
		{script: "a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []bool{true, false}, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a += true", input: map[string]interface{}{"a": []bool{}}, runOutput: []bool{true}, ouput: map[string]interface{}{"a": []bool{true}}},

		{script: "a", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{}, ouput: map[string]interface{}{"a": []int32{}}},
		{script: "a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []int32{1, 2}, ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, ouput: map[string]interface{}{"a": []int32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, ouput: map[string]interface{}{"a": []int32{1}}},

		{script: "a", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{}, ouput: map[string]interface{}{"a": []int64{}}},
		{script: "a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []int64{1, 2}, ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, ouput: map[string]interface{}{"a": []int64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, ouput: map[string]interface{}{"a": []int64{1}}},

		{script: "a", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{}, ouput: map[string]interface{}{"a": []float32{}}},
		{script: "a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []float32{1, 2}, ouput: map[string]interface{}{"a": []float32{1, 2}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1}, ouput: map[string]interface{}{"a": []float32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1.1}, ouput: map[string]interface{}{"a": []float32{1.1}}},

		{script: "a", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{}, ouput: map[string]interface{}{"a": []float64{}}},
		{script: "a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []float64{1, 2}, ouput: map[string]interface{}{"a": []float64{1, 2}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1}, ouput: map[string]interface{}{"a": []float64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1.1}, ouput: map[string]interface{}{"a": []float64{1.1}}},

		{script: "a", input: map[string]interface{}{"a": []string{}}, runOutput: []string{}, ouput: map[string]interface{}{"a": []string{}}},
		{script: "a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []string{"a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}}},
		{script: "a += \"a\"", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, ouput: map[string]interface{}{"a": []string{"a"}}},
		{script: "a += 97", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, ouput: map[string]interface{}{"a": []string{"a"}}},
	}
	runTests(t, tests)
}

func TestArrayAppendArrays(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []bool{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, false}, ouput: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false}}},

		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": []bool{}, "b": []interface{}{true}}},
		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true, true}, ouput: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, true}}},
		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, true, false}, ouput: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, true, false}}},

		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{true, false}, ouput: map[string]interface{}{"a": []bool{}, "b": []interface{}{true, false}}},
		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true, false, true}, ouput: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, false, true}}},
		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, false, true, false}, ouput: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false, true, false}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int32(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int32(1), int32(2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int32(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int32(1), int32(2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), int64(2), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int64(2), int32(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), int64(2), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int64(2), int32(1), int32(2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int32(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int32(1), int32(2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), float64(2.2), int32(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1), float64(2.2)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), float64(2.2), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), float64(2.2), int32(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), float64(2.2), int32(1), int32(2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1), int64(2)}, ouput: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), int64(2), int32(1)}, ouput: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int64(2), int32(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int64(2), int32(1), int32(2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(1), int64(2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), int64(2), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(2), int64(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(2), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2), int64(1), int64(2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(1), int64(2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), float64(2.2), int64(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1), float64(2.2)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), float64(2.2), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), float64(2.2), int64(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), float64(2.2), int64(1), int64(2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1), int64(2)}, ouput: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), int64(2), int64(1)}, ouput: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(2), int64(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(2), int64(1), int64(2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float32(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float32(1), float32(2)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float32(1.1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float32(1.1), float32(2.2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float32(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float32(1), float32(2)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float32(1.1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float32(1.1), float32(2.2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), int64(2), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), int64(2), float32(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), int64(2), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), int64(2), float32(1), float32(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), int64(2), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), int64(2), float32(1.1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float32(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float32(1), float32(2)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float32(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1), float64(2.2)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float64(2.2), float32(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float64(2.2), float32(1), float32(2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float64(2.2), float32(1.1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1), int64(2)}, ouput: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1)}, ouput: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), int64(2), float32(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1), float32(2)}, ouput: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), int64(2), float32(1), float32(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1.1)}, ouput: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), int64(2), float32(1.1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1), float64(2)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(1), float64(2)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(1.1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(1.1), float64(2.2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), int64(2), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), int64(2), float64(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), int64(2), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), int64(2), float64(1), float64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), int64(2), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), int64(2), float64(1.1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(1), float64(2)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(2.2), float64(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(2.2), float64(1), float64(2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(2.2), float64(1.1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1), int64(2)}, ouput: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1)}, ouput: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), int64(2), float64(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1), float64(2)}, ouput: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), int64(2), float64(1), float64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1.1)}, ouput: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), int64(2), float64(1.1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []string{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a"}, ouput: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a"}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b"}}},

		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{"a"}, ouput: map[string]interface{}{"a": []string{}, "b": []interface{}{"a"}}},
		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a", "a"}, ouput: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "a"}}},
		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "a", "b"}}},

		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{"a", "b"}, ouput: map[string]interface{}{"a": []string{}, "b": []interface{}{"a", "b"}}},
		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a", "b", "a"}, ouput: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "b", "a"}}},
		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "b", "a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b", "a", "b"}}},

		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{int64(1), "a"}, ouput: map[string]interface{}{"a": []string{}, "b": []interface{}{int64(1), "a"}}},
		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{int64(1), "a", "a"}, ouput: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{int64(1), "a", "a"}}},
		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{int64(1), "a", "a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{int64(1), "a", "a", "b"}}},
	}
	runTests(t, tests)
}

func TestVar(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "var a = 1", runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "var a, b = 1, 2", runOutput: int64(2), ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "var a, b, c = 1, 2, 3", runOutput: int64(3), ouput: map[string]interface{}{"a": int64(1), "b": int64(2), "c": int64(3)}},
	}
	runTests(t, tests)
}

func TestForLoop(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "for { break }", runOutput: nil},
		{script: "for {a = 1; if a == 1 { break } }", runOutput: nil},
		{script: "a = 1; for { if a == 1 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for { if a == 1 { break }; a++ }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},

		{script: "a = 1; for { a++; if a == 2 { continue } else { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a = 1; for { a++; if a == 2 { continue }; if a == 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},

		{script: "for a in [1] { if a == 1 { break } }", runOutput: nil},
		{script: "for a in [1, 2] { if a == 2 { break } }", runOutput: nil},
		{script: "for a in [1, 2, 3] { if a == 3 { break } }", runOutput: nil},

		{script: "a = [1]; for b in a { if b == 1 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1, 2]; for b in a { if b == 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{script: "a = [1, 2, 3]; for b in a { if b == 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1]; b = 0; for c in a { b = c }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: "a = [1, 2]; b = 0; for c in a { b = c }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": int64(2)}},
		{script: "a = [1, 2, 3]; b = 0; for c in a { b = c }", runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}, "b": int64(3)}},

		{script: "a = 1; for a < 2 { a++ }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for a < 3 { a++ }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},

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

		{script: "for a = 1; nil; nil { if a == 1 { break } }", runOutput: nil},
		{script: "for a = 1; nil; nil { if a == 2 { break }; a++ }", runOutput: nil},
		{script: "for a = 1; nil; nil { a++; if a == 3 { break } }", runOutput: nil},

		{script: "for a = 1; a < 1; nil { }", runOutput: nil},
		{script: "for a = 1; a > 1; nil { }", runOutput: nil},
		{script: "for a = 1; a == 1; nil { break }", runOutput: nil},

		{script: "for a = 1; a == 1; a++ { }", runOutput: nil},
		{script: "for a = 1; a < 2; a++ { }", runOutput: nil},
		{script: "for a = 1; a < 3; a++ { }", runOutput: nil},

		{script: "a = 1; for b = 1; a < 1; a++ { }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ { }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ { }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},

		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 1 { continue } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 1 { continue } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 1 { continue } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},

		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 1 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 1 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 1 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 1; a++ {  if a == 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for b = 1; a < 2; a++ {  if a == 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a = 1; for b = 1; a < 3; a++ {  if a == 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},

		{script: "func x() { for a = 1; a < 3; a++ { if a == 1 { return a } } }; x()", runOutput: int64(1)},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 2 { return a } } }; x()", runOutput: int64(2)},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 3 { return a } } }; x()", runOutput: nil},
		{script: "func x() { for a = 1; a < 3; a++ { if a == 4 { return a } } }; x()", runOutput: nil},

		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 1 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 2 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 3 { continue } }; return a }; x()", runOutput: int64(3)},
		{script: "func x() { a = 1; for b = 1; a < 3; a++ { if a == 4 { continue } }; return a }; x()", runOutput: int64(3)},
	}
	runTests(t, tests)
}

func TestTemp(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{}
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

		env := NewEnv()

		for inputName, inputValue := range test.input {
			err = env.Define(inputName, inputValue)
			if err != nil {
				t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.script)
				continue loop
			}
		}

		value, err = Run(stmts, env)
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
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
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
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			}
		}

		env.Destroy()
	}
}

func testInterrupt(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	waitChan := make(chan struct{}, 1)

	env := NewEnv()
	sleepMillisecond := func() { time.Sleep(time.Millisecond) }

	err := env.Define("println", fmt.Println)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = env.Define("sleep", sleepMillisecond)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}

	script := `
# sleep for 10 seconds
for i = 0; i < 10000; i++ {
	sleep()
}
# Should interrupt before printing the next line
println("this line should not be printed")
`

	go func() {
		close(waitChan)
		_, err := env.Execute(script)
		if err == nil {
			t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
		} else if err.Error() != InterruptError.Error() {
			t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
		}
		waitGroup.Done()
	}()

	<-waitChan
	Interrupt(env)

	waitGroup.Wait()
}

func TestInterruptRaces(t *testing.T) {
	// Run testInterrupt many times
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			testInterrupt(t)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}
