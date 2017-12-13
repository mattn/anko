package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

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

		{script: "a[2]", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[0]", input: map[string]interface{}{"a": []bool{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []bool{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int32{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int64{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float32{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float64{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []string{}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []string{}}},

		{script: "a[1] = true", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: true, ouput: map[string]interface{}{"a": []bool{true, true}}},
		// TOFIX:
		//		{script: "a[1] = 3", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(3), ouput: map[string]interface{}{"a": []int32{1, 3}}},
		{script: "a[1] = 3", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(3), ouput: map[string]interface{}{"a": []int64{1, 3}}},
		// TOFIX:
		//		{script: "a[1] = 3.3", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(3.3), ouput: map[string]interface{}{"a": []float32{1.1, 3.3}}},
		{script: "a[1] = 3.3", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(3.3), ouput: map[string]interface{}{"a": []float64{1.1, 3.3}}},
		{script: "a[1] = \"c\"", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "c", ouput: map[string]interface{}{"a": []string{"a", "c"}}},

		{script: "a[2] = true", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[2] = 3", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2] = 3", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[2] = \"c\"", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a = []; a[0]", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[-1]", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[0] = 1", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[-1] = 1", runError: fmt.Errorf("index out of range")},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": nil}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": true}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int32(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float32(1.1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float64(1.1)}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "1"}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarBoolP}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt32P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt64P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat32P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat64P}, runOutput: int64(2), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarStringP}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": nil}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": true}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int(1)}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int32(1)}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": float32(1.1)}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": float64(1.1)}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": "1"}, runOutput: int64(3), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "make([]bool, nil)", runError: fmt.Errorf("Undefined type 'bool'")},
		{script: "make([]nilT, nil)", types: map[string]interface{}{"nilT": nil}, runError: fmt.Errorf("invalid type for make array")},

		// remove ", nil" when/if syntax support it
		{script: "make([]bool, nil)", types: map[string]interface{}{"bool": true}, runOutput: []bool{}},
		{script: "make([]int32, nil)", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{}},
		{script: "make([]int64, nil)", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{}},
		{script: "make([]float32, nil)", types: map[string]interface{}{"float32": float32(1)}, runOutput: []float32{}},
		{script: "make([]float64, nil)", types: map[string]interface{}{"float64": float64(1)}, runOutput: []float64{}},
		{script: "make([]string, nil)", types: map[string]interface{}{"string": "a"}, runOutput: []string{}},

		{script: "make([]bool, 0)", types: map[string]interface{}{"bool": true}, runOutput: []bool{}},
		{script: "make([]int32, 0)", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{}},
		{script: "make([]int64, 0)", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{}},
		{script: "make([]float32, 0)", types: map[string]interface{}{"float32": float32(1)}, runOutput: []float32{}},
		{script: "make([]float64, 0)", types: map[string]interface{}{"float64": float64(1)}, runOutput: []float64{}},
		{script: "make([]string, 0)", types: map[string]interface{}{"string": "a"}, runOutput: []string{}},

		{script: "make([]bool, 2)", types: map[string]interface{}{"bool": true}, runOutput: []bool{false, false}},
		{script: "make([]int32, 2)", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{int32(0), int32(0)}},
		{script: "make([]int64, 2)", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{int64(0), int64(0)}},
		{script: "make([]float32, 2)", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: []float32{float32(0), float32(0)}},
		{script: "make([]float64, 2)", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: []float64{float64(0), float64(0)}},
		{script: "make([]string, 2)", types: map[string]interface{}{"string": "a"}, runOutput: []string{"", ""}},

		{script: "make([]bool, 0, 2)", types: map[string]interface{}{"bool": true}, runOutput: []bool{}},
		{script: "make([]int32, 0, 2)", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{}},
		{script: "make([]int64, 0, 2)", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{}},
		{script: "make([]float32, 0, 2)", types: map[string]interface{}{"float32": float32(1)}, runOutput: []float32{}},
		{script: "make([]float64, 0, 2)", types: map[string]interface{}{"float64": float64(1)}, runOutput: []float64{}},
		{script: "make([]string, 0, 2)", types: map[string]interface{}{"string": "a"}, runOutput: []string{}},

		{script: "make([]bool, 2, 2)", types: map[string]interface{}{"bool": true}, runOutput: []bool{false, false}},
		{script: "make([]int32, 2, 2)", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{int32(0), int32(0)}},
		{script: "make([]int64, 2, 2)", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{int64(0), int64(0)}},
		{script: "make([]float32, 2, 2)", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: []float32{float32(0), float32(0)}},
		{script: "make([]float64, 2, 2)", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: []float64{float64(0), float64(0)}},
		{script: "make([]string, 2, 2)", types: map[string]interface{}{"string": "a"}, runOutput: []string{"", ""}},

		{script: "a = make([]bool, 0); a += true; a += false", types: map[string]interface{}{"bool": true}, runOutput: []bool{true, false}, ouput: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a = make([]int32, 0); a += 1; a += 2", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{int32(1), int32(2)}, ouput: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{script: "a = make([]int64, 0); a += 1; a += 2", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		{script: "a = make([]float32, 0); a += 1.1; a += 2.2", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: []float32{float32(1.1), float32(2.2)}, ouput: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{script: "a = make([]float64, 0); a += 1.1; a += 2.2", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: []float64{float64(1.1), float64(2.2)}, ouput: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{script: "a = make([]string, 0); a += \"a\"; a += \"b\"", types: map[string]interface{}{"string": "a"}, runOutput: []string{"a", "b"}, ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a = make([]bool, 2); a[0] = true; a[1] = false", types: map[string]interface{}{"bool": true}, runOutput: false, ouput: map[string]interface{}{"a": []bool{true, false}}},
		// TOFIX:
		// {script: "a = make([]int32, 2); a[0] = 1; a[1] = 2", types: map[string]interface{}{"int32": int32(1)}, runOutput: int32(2), ouput: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{script: "a = make([]int64, 2); a[0] = 1; a[1] = 2", types: map[string]interface{}{"int64": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		// TOFIX:
		// {script: "a = make([]float32, 2); a[0] = 1.1; a[1] = 2.2", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: float32(2.2), ouput: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{script: "a = make([]float64, 2); a[0] = 1.1; a[1] = 2.2", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: float64(2.2), ouput: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{script: "a = make([]string, 2); a[0] = \"a\"; a[1] = \"b\"", types: map[string]interface{}{"string": "a"}, runOutput: "b", ouput: map[string]interface{}{"a": []string{"a", "b"}}},

		// remove ", nil" when/if syntax support it
		{script: "make([]boolA, nil)", types: map[string]interface{}{"boolA": []bool{}}, runOutput: [][]bool{}},
		{script: "make([]int32A, nil)", types: map[string]interface{}{"int32A": []int32{}}, runOutput: [][]int32{}},
		{script: "make([]int64A, nil)", types: map[string]interface{}{"int64A": []int64{}}, runOutput: [][]int64{}},
		{script: "make([]float32A, nil)", types: map[string]interface{}{"float32A": []float32{}}, runOutput: [][]float32{}},
		{script: "make([]float64A, nil)", types: map[string]interface{}{"float64A": []float64{}}, runOutput: [][]float64{}},
		{script: "make([]stringA, nil)", types: map[string]interface{}{"stringA": []string{}}, runOutput: [][]string{}},
	}
	runTests(t, tests)
}

func TestArraySlice(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = [1, 2, 3]; a[true:2]", runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:true]", runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[0:0]", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:1]", runOutput: []interface{}{int64(1)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:2]", runOutput: []interface{}{int64(1), int64(2)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:3]", runOutput: []interface{}{int64(1), int64(2), int64(3)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:4]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[1:0]", runError: fmt.Errorf("invalid slice index"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:1]", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:2]", runOutput: []interface{}{int64(2)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:3]", runOutput: []interface{}{int64(2), int64(3)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:4]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[2:1]", runError: fmt.Errorf("invalid slice index"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:2]", runOutput: []interface{}{}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:3]", runOutput: []interface{}{int64(3)}, ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:4]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[3:2]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[3:3]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[3:4]", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[true:2] = 4", runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:true] = 4", runError: fmt.Errorf("index must be a number"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:0] = 4", runError: fmt.Errorf("slice cannot be assigned"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:1] = 4", runError: fmt.Errorf("slice cannot be assigned"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:4] = 4", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:0] = 4", runError: fmt.Errorf("invalid slice index"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:4] = 4", runError: fmt.Errorf("index out of range"), ouput: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
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

		{script: "a.b", input: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(2), ouput: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(2), ouput: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(2.2), ouput: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(2.2), ouput: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}},

		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: int32(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runError: fmt.Errorf("type map does not support slice operation"), ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}, runOutput: reflect.Value{}, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}},

		{script: "a.b = true", input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		// TOFIX:
		//		{script: "a.b = 3", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(3), ouput: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{script: "a.b = 3", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(3), ouput: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		// TOFIX:
		//		{script: "a.b = 3.3", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(3.3), ouput: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{script: "a.b = 3.3", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(3.3), ouput: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{script: "a.b = \"c\"", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "c", ouput: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{script: "a[\"b\"] = true", input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		// TOFIX:
		//		{script: "a[\"b\"] = 3", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(3), ouput: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{script: "a[\"b\"] = 3", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(3), ouput: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		// TOFIX:
		//		{script: "a[\"b\"] = 3.3", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(3.3), ouput: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{script: "a[\"b\"] = 3.3", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(3.3), ouput: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{script: "a[\"b\"] = \"c\"", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "c", ouput: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{script: "a[c] = true", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, runError: fmt.Errorf("map key must be string type"), runOutput: nil, ouput: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},

		{script: "make(mapStringBool)", types: map[string]interface{}{"mapStringBool": map[string]bool{}}, runOutput: map[string]bool{}},
		{script: "make(mapStringInt32)", types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, runOutput: map[string]int32{}},
		{script: "make(mapStringInt64)", types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, runOutput: map[string]int64{}},
		{script: "make(mapStringFloat32)", types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, runOutput: map[string]float32{}},
		{script: "make(mapStringFloat64)", types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, runOutput: map[string]float64{}},
		{script: "make(mapStringString)", types: map[string]interface{}{"mapStringString": map[string]string{}}, runOutput: map[string]string{}},

		{script: "a = make(mapStringBool)", types: map[string]interface{}{"mapStringBool": map[string]bool{}}, runOutput: map[string]bool{}, ouput: map[string]interface{}{"a": map[string]bool{}}},
		{script: "a = make(mapStringInt32)", types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, runOutput: map[string]int32{}, ouput: map[string]interface{}{"a": map[string]int32{}}},
		{script: "a = make(mapStringInt64)", types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, runOutput: map[string]int64{}, ouput: map[string]interface{}{"a": map[string]int64{}}},
		{script: "a = make(mapStringFloat32)", types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, runOutput: map[string]float32{}, ouput: map[string]interface{}{"a": map[string]float32{}}},
		{script: "a = make(mapStringFloat64)", types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, runOutput: map[string]float64{}, ouput: map[string]interface{}{"a": map[string]float64{}}},
		{script: "a = make(mapStringString)", types: map[string]interface{}{"mapStringString": map[string]string{}}, runOutput: map[string]string{}, ouput: map[string]interface{}{"a": map[string]string{}}},

		{script: "a = make(mapStringBool); a[\"b\"] = true", types: map[string]interface{}{"mapStringBool": map[string]bool{"b": true}}, runOutput: true, ouput: map[string]interface{}{"a": map[string]bool{"b": true}}},
		// TOFIX:
		//		{script: "a = make(mapStringInt32); a[\"b\"] = 1", types: map[string]interface{}{"mapStringInt32": map[string]int32{"b": int32(1)}}, runOutput: int32(1), ouput: map[string]interface{}{"a": map[string]int32{"b": int32(1)}}},
		{script: "a = make(mapStringInt64); a[\"b\"] = 1", types: map[string]interface{}{"mapStringInt64": map[string]int64{"b": int64(1)}}, runOutput: int64(1), ouput: map[string]interface{}{"a": map[string]int64{"b": int64(1)}}},
		// TOFIX:
		//		{script: "a = make(mapStringFloat32); a[\"b\"] = 1.1", types: map[string]interface{}{"mapStringFloat32": map[string]float32{"b": float32(1.1)}}, runOutput: float32(1.1), ouput: map[string]interface{}{"a": map[string]float32{"b": float32(1.1)}}},
		{script: "a = make(mapStringFloat64); a[\"b\"] = 1.1", types: map[string]interface{}{"mapStringFloat64": map[string]float64{"b": float64(1.1)}}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": map[string]float64{"b": float64(1.1)}}},
		{script: "a = make(mapStringString); a[\"b\"] = \"b\"", types: map[string]interface{}{"mapStringString": map[string]string{"b": "b"}}, runOutput: "b", ouput: map[string]interface{}{"a": map[string]string{"b": "b"}}},
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
