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

	testVarValueBool    = reflect.ValueOf(true)
	testVarValueInt32   = reflect.ValueOf(int32(1))
	testVarValueInt64   = reflect.ValueOf(int64(1))
	testVarValueFloat32 = reflect.ValueOf(float32(1.1))
	testVarValueFloat64 = reflect.ValueOf(float64(1.1))
	testVarValueString  = reflect.ValueOf("a")
)

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
		{script: "a = b", input: map[string]interface{}{"b": 'a'}, runOutput: 'a', ouput: map[string]interface{}{"a": 'a', "b": 'a'}},

		{script: "a = b", input: map[string]interface{}{"b": testVarValue}, runOutput: testVarValue, ouput: map[string]interface{}{"a": testVarValue, "b": testVarValue}},
		{script: "a = b", input: map[string]interface{}{"b": testVarBoolP}, runOutput: testVarBoolP, ouput: map[string]interface{}{"a": testVarBoolP, "b": testVarBoolP}},
		{script: "a = b", input: map[string]interface{}{"b": testVarInt32P}, runOutput: testVarInt32P, ouput: map[string]interface{}{"a": testVarInt32P, "b": testVarInt32P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarInt64P}, runOutput: testVarInt64P, ouput: map[string]interface{}{"a": testVarInt64P, "b": testVarInt64P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarFloat32P}, runOutput: testVarFloat32P, ouput: map[string]interface{}{"a": testVarFloat32P, "b": testVarFloat32P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarFloat64P}, runOutput: testVarFloat64P, ouput: map[string]interface{}{"a": testVarFloat64P, "b": testVarFloat64P}},
		{script: "a = b", input: map[string]interface{}{"b": testVarStringP}, runOutput: testVarStringP, ouput: map[string]interface{}{"a": testVarStringP, "b": testVarStringP}},

		{script: "a = b", input: map[string]interface{}{"b": testVarValueBool}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueInt32}, runOutput: int32(1), ouput: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueInt64}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueFloat32}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueFloat64}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{script: "a = b", input: map[string]interface{}{"b": testVarValueString}, runOutput: "a", ouput: map[string]interface{}{"a": "a", "b": "a"}},

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

		{script: "a == b", input: map[string]interface{}{"a": false, "b": false}, runOutput: true, ouput: map[string]interface{}{"a": false, "b": false}},
		{script: "a == b", input: map[string]interface{}{"a": false, "b": true}, runOutput: false, ouput: map[string]interface{}{"a": false, "b": true}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": false}, runOutput: false, ouput: map[string]interface{}{"a": true, "b": false}},
		{script: "a == b", input: map[string]interface{}{"a": true, "b": true}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},

		{script: "a == b", input: map[string]interface{}{"a": int32(1), "b": int32(1)}, runOutput: true, ouput: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(1), "b": int32(2)}, runOutput: false, ouput: map[string]interface{}{"a": int32(1), "b": int32(2)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(2), "b": int32(1)}, runOutput: false, ouput: map[string]interface{}{"a": int32(2), "b": int32(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int32(2), "b": int32(2)}, runOutput: true, ouput: map[string]interface{}{"a": int32(2), "b": int32(2)}},

		{script: "a == b", input: map[string]interface{}{"a": int64(1), "b": int64(1)}, runOutput: true, ouput: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(1), "b": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2), "b": int64(1)}},
		{script: "a == b", input: map[string]interface{}{"a": int64(2), "b": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},

		{script: "a == b", input: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}, runOutput: true, ouput: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}, runOutput: false, ouput: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}, runOutput: false, ouput: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}, runOutput: true, ouput: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}},

		{script: "a == b", input: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}, runOutput: true, ouput: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}},
		{script: "a == b", input: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}},

		{script: "a == b", input: map[string]interface{}{"a": 'a', "b": 'a'}, runOutput: true, ouput: map[string]interface{}{"a": 'a', "b": 'a'}},
		{script: "a == b", input: map[string]interface{}{"a": 'a', "b": 'b'}, runOutput: false, ouput: map[string]interface{}{"a": 'a', "b": 'b'}},
		{script: "a == b", input: map[string]interface{}{"a": 'b', "b": 'a'}, runOutput: false, ouput: map[string]interface{}{"a": 'b', "b": 'a'}},
		{script: "a == b", input: map[string]interface{}{"a": 'b', "b": 'b'}, runOutput: true, ouput: map[string]interface{}{"a": 'b', "b": 'b'}},

		{script: "a == b", input: map[string]interface{}{"a": "a", "b": "a"}, runOutput: true, ouput: map[string]interface{}{"a": "a", "b": "a"}},
		{script: "a == b", input: map[string]interface{}{"a": "a", "b": "b"}, runOutput: false, ouput: map[string]interface{}{"a": "a", "b": "b"}},
		{script: "a == b", input: map[string]interface{}{"a": "b", "b": "a"}, runOutput: false, ouput: map[string]interface{}{"a": "b", "b": "a"}},
		{script: "a == b", input: map[string]interface{}{"a": "b", "b": "b"}, runOutput: true, ouput: map[string]interface{}{"a": "b", "b": "b"}},
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
		{script: "func aFunc() {return}; aFunc()", runOutput: nil},
		{script: "func aFunc() {return}; a = aFunc()", runOutput: nil, ouput: map[string]interface{}{"a": nil}},

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

		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": float64(1.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "func aFunc() {return a}; aFunc()", input: map[string]interface{}{"a": "a"}, runOutput: "a", ouput: map[string]interface{}{"a": "a"}},

		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": reflect.Value{}}, runOutput: []interface{}{nil, nil}, ouput: map[string]interface{}{"a": reflect.Value{}}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": nil}, runOutput: []interface{}{nil, nil}, ouput: map[string]interface{}{"a": nil}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": true}, runOutput: []interface{}{true, true}, ouput: map[string]interface{}{"a": true}},
		{script: "func aFunc() {return a, a}; aFunc()", input: map[string]interface{}{"a": int32(1)}, runOutput: []interface{}{int32(1), int32(1)}, ouput: map[string]interface{}{"a": int32(1)}},

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

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a", input: map[string]interface{}{"a": 'a'}, runOutput: 'a', ouput: map[string]interface{}{"a": 'a'}},
		{script: "a[0]", input: map[string]interface{}{"a": 'a'}, runError: fmt.Errorf("type int32 does not support indexing"), runOutput: nil, ouput: map[string]interface{}{"a": 'a'}},

		{script: "a", input: map[string]interface{}{"a": "test"}, runOutput: "test", ouput: map[string]interface{}{"a": "test"}},
		{script: "a[0]", input: map[string]interface{}{"a": "test"}, runOutput: 't', ouput: map[string]interface{}{"a": "test"}},
		{script: "a[1]", input: map[string]interface{}{"a": "test"}, runOutput: 'e', ouput: map[string]interface{}{"a": "test"}},
		{script: "a[3]", input: map[string]interface{}{"a": "test"}, runOutput: 't', ouput: map[string]interface{}{"a": "test"}},
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
		{script: "a", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{}, ouput: map[string]interface{}{"a": []int32{}}},
		{script: "a", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{}, ouput: map[string]interface{}{"a": []int64{}}},
		{script: "a", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{}, ouput: map[string]interface{}{"a": []float32{}}},
		{script: "a", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{}, ouput: map[string]interface{}{"a": []float64{}}},
		{script: "a", input: map[string]interface{}{"a": []string{}}, runOutput: []string{}, ouput: map[string]interface{}{"a": []string{}}},

		{script: "a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []bool{true, false}, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []int32{1, 2}, ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []int64{1, 2}, ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []float32{1.1, 2.2}, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []float64{1.1, 2.2}, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []string{"a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a += true", input: map[string]interface{}{"a": []bool{}}, runOutput: []bool{true}, ouput: map[string]interface{}{"a": []bool{true}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, ouput: map[string]interface{}{"a": []int32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, ouput: map[string]interface{}{"a": []int32{1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, ouput: map[string]interface{}{"a": []int64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, ouput: map[string]interface{}{"a": []int64{1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1}, ouput: map[string]interface{}{"a": []float32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1.1}, ouput: map[string]interface{}{"a": []float32{1.1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1}, ouput: map[string]interface{}{"a": []float64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1.1}, ouput: map[string]interface{}{"a": []float64{1.1}}},
		{script: "a += \"a\"", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, ouput: map[string]interface{}{"a": []string{"a"}}},
		{script: "a += 97", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, ouput: map[string]interface{}{"a": []string{"a"}}},

		{script: "a[0]", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: true, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(1), ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(1), ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "a", ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[1]", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: false, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[1]", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(2), ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(2), ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(2.2), ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(2.2), ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[2]", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[0]", input: map[string]interface{}{"a": []bool{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []bool{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int32{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []int32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int64{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []int64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float32{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []float32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float64{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []float64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []string{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, ouput: map[string]interface{}{"a": []string{}}},

		{script: "a = []; a[0]", runError: fmt.Errorf("index out of range"), runOutput: nil},
		{script: "a = []; a[-1]", runError: fmt.Errorf("index out of range"), runOutput: nil},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": nil}, runError: fmt.Errorf("index should be a number"), runOutput: nil, ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": true}, runError: fmt.Errorf("index should be a number"), runOutput: nil, ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int32(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float32(1.1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float64(1.1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "1"}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("index should be a number"), runOutput: nil, ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarBoolP}, runError: fmt.Errorf("index should be a number"), runOutput: nil, ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt32P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt64P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat32P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat64P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarStringP}, runError: fmt.Errorf("index should be a number"), runOutput: nil, ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
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

func TestMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = {}", runOutput: map[string]interface{}{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}", runOutput: map[string]interface{}{"b": nil}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}", runOutput: map[string]interface{}{"b": true}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}", runOutput: map[string]interface{}{"b": int64(1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}", runOutput: map[string]interface{}{"b": float64(1.1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}", runOutput: map[string]interface{}{"b": "b"}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a = {\"b\": {}}", runOutput: map[string]interface{}{"b": map[string]interface{}{}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{}}}},
		{script: "a = {\"b\": {\"c\": nil}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": nil}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": nil}}}},
		{script: "a = {\"b\": {\"c\": true}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": true}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": true}}}},
		{script: "a = {\"b\": {\"c\": 1}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}}},
		{script: "a = {\"b\": {\"c\": 1.1}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}}},
		{script: "a = {\"b\": {\"c\": \"c\"}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": "c"}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": "c"}}}},

		{script: "a = {\"b\": {}}; a.b", runOutput: map[string]interface{}{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{}}}},
		{script: "a = {\"b\": {\"c\": nil}}; a.b", runOutput: map[string]interface{}{"c": nil}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": nil}}}},
		{script: "a = {\"b\": {\"c\": true}}; a.b", runOutput: map[string]interface{}{"c": true}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": true}}}},
		{script: "a = {\"b\": {\"c\": 1}}; a.b", runOutput: map[string]interface{}{"c": int64(1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}}},
		{script: "a = {\"b\": {\"c\": 1.1}}; a.b", runOutput: map[string]interface{}{"c": float64(1.1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}}},
		{script: "a = {\"b\": {\"c\": \"c\"}}; a.b", runOutput: map[string]interface{}{"c": "c"}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": "c"}}}},

		{script: "a = {\"b\": []}", runOutput: map[string]interface{}{"b": []interface{}{}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}", runOutput: map[string]interface{}{"b": []interface{}{nil}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}", runOutput: map[string]interface{}{"b": []interface{}{true}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}", runOutput: map[string]interface{}{"b": []interface{}{int64(1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}", runOutput: map[string]interface{}{"b": []interface{}{float64(1.1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"c\"]}", runOutput: map[string]interface{}{"b": []interface{}{"c"}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"c"}}}},

		{script: "a = {}; a.b", runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}; a.b", runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}; a.b", runOutput: true, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}; a.b", runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}; a.b", runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}; a.b", runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a = {}; a[\"b\"]", runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}; a[\"b\"]", runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}; a[\"b\"]", runOutput: true, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}; a[\"b\"]", runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}; a[\"b\"]", runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}; a[\"b\"]", runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: map[string]interface{}{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: map[string]interface{}{"b": nil}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: map[string]interface{}{"b": true}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: map[string]interface{}{"b": int32(1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: map[string]interface{}{"b": int64(1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: map[string]interface{}{"b": float32(1.1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: map[string]interface{}{"b": float64(1.1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: map[string]interface{}{"b": "b"}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: int32(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: int32(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}},
	}
	runTests(t, tests)
}

func TestArraysAndMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = [{\"b\": nil}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(nil)}}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(nil)}}}},
		{script: "a = [{\"b\": true}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(true)}}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(true)}}}},
		{script: "a = [{\"b\": 1}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}}},
		{script: "a = [{\"b\": 1.1}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}}},
		{script: "a = [{\"b\": \"b\"}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}("b")}}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}("b")}}}},

		{script: "a = [{\"b\": nil}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(nil)}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(nil)}}}},
		{script: "a = [{\"b\": true}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(true)}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(true)}}}},
		{script: "a = [{\"b\": 1}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(int64(1))}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}}},
		{script: "a = [{\"b\": 1.1}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(float64(1.1))}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}}},
		{script: "a = [{\"b\": \"b\"}]; a[0]", runOutput: map[string]interface{}{"b": interface{}("b")}, ouput: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}("b")}}}},

		{script: "a = {\"b\": []}", runOutput: map[string]interface{}{"b": []interface{}{}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}", runOutput: map[string]interface{}{"b": []interface{}{nil}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}", runOutput: map[string]interface{}{"b": []interface{}{true}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}", runOutput: map[string]interface{}{"b": []interface{}{int64(1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}", runOutput: map[string]interface{}{"b": []interface{}{float64(1.1)}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"b\"]}", runOutput: map[string]interface{}{"b": []interface{}{"b"}}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"b"}}}},

		{script: "a = {\"b\": []}; a.b", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}; a.b", runOutput: []interface{}{nil}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}; a.b", runOutput: []interface{}{true}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}; a.b", runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}; a.b", runOutput: []interface{}{float64(1.1)}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"b\"]}; a.b", runOutput: []interface{}{"b"}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"b"}}}},
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

		{script: "a = 1; for nil { a++; if a > 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for nil { a++; if a > 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a = 1; for true { a++; if a > 2 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a = 1; for true { a++; if a > 3 { break } }", runOutput: nil, ouput: map[string]interface{}{"a": int64(4)}},

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

		{script: "func x() { return [1, 2] }; a = 1; for i in x() { a++ }", runOutput: nil, ouput: map[string]interface{}{"a": int64(3)}},
		{script: "func x() { return [1, 2, 3] }; a = 1; for i in x() { a++ }", runOutput: nil, ouput: map[string]interface{}{"a": int64(4)}},

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

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{reflect.Value{}}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{reflect.Value{}}, "b": reflect.Value{}}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{nil}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{nil}, "b": nil}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{true}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{true}, "b": true}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{int32(1)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int32(1)}, "b": int32(1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{int64(1)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": float32(1.1)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": float64(1.1)}},

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}, "b": interface{}(reflect.Value{})}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(nil)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(nil)}, "b": interface{}(nil)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(true)}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(true)}, "b": interface{}(true)}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}, "b": interface{}(int32(1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}, "b": interface{}(int64(1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}, "b": interface{}(float32(1.1))}},
		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}}, runOutput: nil, ouput: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}, "b": interface{}(float64(1.1))}},

		{script: "b = 0; for i in a { b = i }", input: map[string]interface{}{"a": interface{}([]interface{}{nil})}, runOutput: nil, ouput: map[string]interface{}{"a": interface{}([]interface{}{nil}), "b": nil}},

		{script: "for i in nil { }", runError: fmt.Errorf("for cannot loop over type interface")},
		{script: "for i in true { }", runError: fmt.Errorf("for cannot loop over type bool")},
		{script: "a = {}; for i in a { }", runError: fmt.Errorf("for cannot loop over type map")},
		{script: "for i in a { }", input: map[string]interface{}{"a": reflect.Value{}}, runError: fmt.Errorf("for cannot loop over type invalid"), ouput: map[string]interface{}{"a": reflect.Value{}}},
		{script: "for i in a { }", input: map[string]interface{}{"a": interface{}(nil)}, runError: fmt.Errorf("for cannot loop over type interface"), ouput: map[string]interface{}{"a": interface{}(nil)}},
		{script: "for i in a { }", input: map[string]interface{}{"a": interface{}(true)}, runError: fmt.Errorf("for cannot loop over type bool"), ouput: map[string]interface{}{"a": interface{}(true)}},

		{script: "a = 1; b = [{\"c\": \"c\"}]; for i in b { a = i }", runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"c": "c"}, "b": []interface{}{map[string]interface{}{"c": "c"}}}},
		{script: "a = 1; b = {\"x\": [{\"y\": \"y\"}]};  for i in b.x { a = i }", runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"y": "y"}, "b": map[string]interface{}{"x": []interface{}{map[string]interface{}{"y": "y"}}}}},
	}
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
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, reflect.ValueOf(test.runOutput), test.script)
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
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, reflect.ValueOf(outputValue), test.script)
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
	env := NewEnv()
	err := env.Define("closeWaitChan", closeWaitChan)
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
