package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var (
	testArrayEmpty []interface{}
	testArray      = []interface{}{nil, true, int64(1), float64(1.1), "a"}
	testMapEmpty   map[string]interface{}
	testMap        = map[string]interface{}{"a": nil, "b": true, "c": int64(1), "d": float64(1.1), "e": "e"}
)

func TestArrays(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = []", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{}}},
		{script: "a = [nil]", runOutput: []interface{}{interface{}(nil)}, output: map[string]interface{}{"a": []interface{}{interface{}(nil)}}},
		{script: "a = [true]", runOutput: []interface{}{true}, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = [\"test\"]", runOutput: []interface{}{"test"}, output: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = [1]", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = [1.1]", runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a = [[]]", runOutput: []interface{}{[]interface{}{}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{}}}},
		{script: "a = [[nil]]", runOutput: []interface{}{[]interface{}{interface{}(nil)}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{interface{}(nil)}}}},
		{script: "a = [[true]]", runOutput: []interface{}{[]interface{}{true}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{true}}}},
		{script: "a = [[\"test\"]]", runOutput: []interface{}{[]interface{}{"test"}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{"test"}}}},
		{script: "a = [[1]]", runOutput: []interface{}{[]interface{}{int64(1)}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1)}}}},
		{script: "a = [[1.1]]", runOutput: []interface{}{[]interface{}{float64(1.1)}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{float64(1.1)}}}},

		{script: "a = []; a += nil", runOutput: []interface{}{nil}, output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = []; a += true", runOutput: []interface{}{true}, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = []; a += \"test\"", runOutput: []interface{}{"test"}, output: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = []; a += 1", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = []; a += 1.1", runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a = []; a += []", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{}}},
		{script: "a = []; a += [nil]", runOutput: []interface{}{nil}, output: map[string]interface{}{"a": []interface{}{nil}}},
		{script: "a = []; a += [true]", runOutput: []interface{}{true}, output: map[string]interface{}{"a": []interface{}{true}}},
		{script: "a = []; a += [\"test\"]", runOutput: []interface{}{"test"}, output: map[string]interface{}{"a": []interface{}{"test"}}},
		{script: "a = []; a += [1]", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{script: "a = []; a += [1.1]", runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},

		{script: "a", input: map[string]interface{}{"a": []bool{}}, runOutput: []bool{}, output: map[string]interface{}{"a": []bool{}}},
		{script: "a", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{}, output: map[string]interface{}{"a": []int32{}}},
		{script: "a", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{}, output: map[string]interface{}{"a": []int64{}}},
		{script: "a", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{}, output: map[string]interface{}{"a": []float32{}}},
		{script: "a", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{}, output: map[string]interface{}{"a": []float64{}}},
		{script: "a", input: map[string]interface{}{"a": []string{}}, runOutput: []string{}, output: map[string]interface{}{"a": []string{}}},

		{script: "a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []bool{true, false}, output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []int32{1, 2}, output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []int64{1, 2}, output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []float32{1.1, 2.2}, output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []float64{1.1, 2.2}, output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []string{"a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a += true", input: map[string]interface{}{"a": []bool{}}, runOutput: []bool{true}, output: map[string]interface{}{"a": []bool{true}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, output: map[string]interface{}{"a": []int32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int32{}}, runOutput: []int32{1}, output: map[string]interface{}{"a": []int32{1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, output: map[string]interface{}{"a": []int64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []int64{}}, runOutput: []int64{1}, output: map[string]interface{}{"a": []int64{1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1}, output: map[string]interface{}{"a": []float32{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float32{}}, runOutput: []float32{1.1}, output: map[string]interface{}{"a": []float32{1.1}}},
		{script: "a += 1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1}, output: map[string]interface{}{"a": []float64{1}}},
		{script: "a += 1.1", input: map[string]interface{}{"a": []float64{}}, runOutput: []float64{1.1}, output: map[string]interface{}{"a": []float64{1.1}}},
		{script: "a += \"a\"", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, output: map[string]interface{}{"a": []string{"a"}}},
		{script: "a += 97", input: map[string]interface{}{"a": []string{}}, runOutput: []string{"a"}, output: map[string]interface{}{"a": []string{"a"}}},

		{script: "a[0]", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: true, output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(1), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(1), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(1.1), output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(1.1), output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[0]", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "a", output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[1]", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: false, output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[1]", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(2), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(2), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(2.2), output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(2.2), output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[1]", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "b", output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[2]", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[2]", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[0]", input: map[string]interface{}{"a": []bool{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []bool{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int32{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []int64{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float32{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float32{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []float64{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float64{}}},
		{script: "a[0]", input: map[string]interface{}{"a": []string{}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []string{}}},

		{script: "a[1] = true", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: true, output: map[string]interface{}{"a": []bool{true, true}}},
		// TOFIX:
		//		{script: "a[1] = 3", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int32(3), output: map[string]interface{}{"a": []int32{1, 3}}},
		{script: "a[1] = 3", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(3), output: map[string]interface{}{"a": []int64{1, 3}}},
		// TOFIX:
		//		{script: "a[1] = 3.3", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float32(3.3), output: map[string]interface{}{"a": []float32{1.1, 3.3}}},
		{script: "a[1] = 3.3", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(3.3), output: map[string]interface{}{"a": []float64{1.1, 3.3}}},
		{script: "a[1] = \"c\"", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "c", output: map[string]interface{}{"a": []string{"a", "c"}}},

		{script: "a[2] = true", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: true, output: map[string]interface{}{"a": []bool{true, false, true}}},
		{script: "a[2] = 3", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: int64(3), output: map[string]interface{}{"a": []int32{1, 2, 3}}},
		{script: "a[2] = 3", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: int64(3), output: map[string]interface{}{"a": []int64{1, 2, 3}}},
		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: float64(3.3), output: map[string]interface{}{"a": []float32{1.1, 2.2, 3.3}}},
		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: float64(3.3), output: map[string]interface{}{"a": []float64{1.1, 2.2, 3.3}}},
		{script: "a[2] = \"c\"", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: "c", output: map[string]interface{}{"a": []string{"a", "b", "c"}}},

		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: float64(3.3), output: map[string]interface{}{"a": []int32{1, 2, 3}}},
		{script: "a[2] = 3.3", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: float64(3.3), output: map[string]interface{}{"a": []int64{1, 2, 3}}},

		{script: "a[3] = true", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[3] = 4", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[3] = 4", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[3] = 4.4", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[3] = 4.4", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[3] = \"d\"", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a[2] = nil", input: map[string]interface{}{"a": []bool{true, false}}, runError: fmt.Errorf("type interface {} cannot be assigned to type bool for array index"), output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a[2] = nil", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("type interface {} cannot be assigned to type int32 for array index"), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2] = \"a\"", input: map[string]interface{}{"a": []int32{1, 2}}, runError: fmt.Errorf("type string cannot be assigned to type int32 for array index"), output: map[string]interface{}{"a": []int32{1, 2}}},
		{script: "a[2] = true", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("type bool cannot be assigned to type int64 for array index"), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2] = \"a\"", input: map[string]interface{}{"a": []int64{1, 2}}, runError: fmt.Errorf("type string cannot be assigned to type int64 for array index"), output: map[string]interface{}{"a": []int64{1, 2}}},
		{script: "a[2] = true", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runError: fmt.Errorf("type bool cannot be assigned to type float32 for array index"), output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{script: "a[2] = \"a\"", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runError: fmt.Errorf("type string cannot be assigned to type float64 for array index"), output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{script: "a[2] = nil", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("type interface {} cannot be assigned to type string for array index"), output: map[string]interface{}{"a": []string{"a", "b"}}},
		{script: "a[2] = true", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("type bool cannot be assigned to type string for array index"), output: map[string]interface{}{"a": []string{"a", "b"}}},
		{script: "a[2] = 1.1", input: map[string]interface{}{"a": []string{"a", "b"}}, runError: fmt.Errorf("type float64 cannot be assigned to type string for array index"), output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a = []; a[0]", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[-1]", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[1] = 1", runError: fmt.Errorf("index out of range")},
		{script: "a = []; a[-1] = 1", runError: fmt.Errorf("index out of range")},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": nil}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": true}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int(1)}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int32(1)}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float32(1.1)}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": float64(1.1)}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "1"}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarBoolP}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt32P}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarInt64P}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat32P}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarFloat64P}, runOutput: int64(2), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a]", input: map[string]interface{}{"a": testVarStringP}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": nil}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": true}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int(1)}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int32(1)}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": int64(1)}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": float32(1.1)}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": float64(1.1)}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": "1"}, runOutput: int64(3), output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{script: "b = [1, 2]; b[a] = 3", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{script: "a", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: [][]interface{}{}, output: map[string]interface{}{"a": [][]interface{}{}}},
		{script: "a[0] = []", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{}}}},
		{script: "a[0] = [nil]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{nil}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}},
		{script: "a[0] = [true]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{true}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}},
		{script: "a[0] = [1]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}},
		{script: "a[0] = [1.1]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}},
		{script: "a[0] = [\"b\"]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: []interface{}{"b"}, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{"b"}}}},

		{script: "a[0] = [nil]; a[0][0]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{nil}}}},
		{script: "a[0] = [true]; a[0][0]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: true, output: map[string]interface{}{"a": [][]interface{}{[]interface{}{true}}}},
		{script: "a[0] = [1]; a[0][0]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: int64(1), output: map[string]interface{}{"a": [][]interface{}{[]interface{}{int64(1)}}}},
		{script: "a[0] = [1.1]; a[0][0]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: float64(1.1), output: map[string]interface{}{"a": [][]interface{}{[]interface{}{float64(1.1)}}}},
		{script: "a[0] = [\"b\"]; a[0][0]", input: map[string]interface{}{"a": [][]interface{}{}}, runOutput: "b", output: map[string]interface{}{"a": [][]interface{}{[]interface{}{"b"}}}},

		{script: "a", input: map[string]interface{}{"a": testArrayEmpty}, runOutput: testArrayEmpty, output: map[string]interface{}{"a": testArrayEmpty}},
		{script: "a += []", input: map[string]interface{}{"a": testArrayEmpty}, runOutput: []interface{}(nil), output: map[string]interface{}{"a": testArrayEmpty}},

		{script: "a", input: map[string]interface{}{"a": testArray}, runOutput: testArray, output: map[string]interface{}{"a": testArray}},
		{script: "a[0]", input: map[string]interface{}{"a": testArray}, runOutput: nil, output: map[string]interface{}{"a": testArray}},
		{script: "a[0] = 1", input: map[string]interface{}{"a": testArray}, runOutput: int64(1), output: map[string]interface{}{"a": testArray}},
		{script: "a[0]", input: map[string]interface{}{"a": testArray}, runOutput: int64(1), output: map[string]interface{}{"a": testArray}},
		{script: "a[0] = nil", input: map[string]interface{}{"a": testArray}, runOutput: nil, output: map[string]interface{}{"a": testArray}},
		{script: "a[0]", input: map[string]interface{}{"a": testArray}, runOutput: nil, output: map[string]interface{}{"a": testArray}},

		{script: "a[1]", input: map[string]interface{}{"a": testArray}, runOutput: true, output: map[string]interface{}{"a": testArray}},
		{script: "a[1] = false", input: map[string]interface{}{"a": testArray}, runOutput: false, output: map[string]interface{}{"a": testArray}},
		{script: "a[1]", input: map[string]interface{}{"a": testArray}, runOutput: false, output: map[string]interface{}{"a": testArray}},
		{script: "a[1] = true", input: map[string]interface{}{"a": testArray}, runOutput: true, output: map[string]interface{}{"a": testArray}},
		{script: "a[1]", input: map[string]interface{}{"a": testArray}, runOutput: true, output: map[string]interface{}{"a": testArray}},

		{script: "a[2]", input: map[string]interface{}{"a": testArray}, runOutput: int64(1), output: map[string]interface{}{"a": testArray}},
		{script: "a[2] = 2", input: map[string]interface{}{"a": testArray}, runOutput: int64(2), output: map[string]interface{}{"a": testArray}},
		{script: "a[2]", input: map[string]interface{}{"a": testArray}, runOutput: int64(2), output: map[string]interface{}{"a": testArray}},
		{script: "a[2] = 1", input: map[string]interface{}{"a": testArray}, runOutput: int64(1), output: map[string]interface{}{"a": testArray}},
		{script: "a[2]", input: map[string]interface{}{"a": testArray}, runOutput: int64(1), output: map[string]interface{}{"a": testArray}},

		{script: "a[3]", input: map[string]interface{}{"a": testArray}, runOutput: float64(1.1), output: map[string]interface{}{"a": testArray}},
		{script: "a[3] = 2.2", input: map[string]interface{}{"a": testArray}, runOutput: float64(2.2), output: map[string]interface{}{"a": testArray}},
		{script: "a[3]", input: map[string]interface{}{"a": testArray}, runOutput: float64(2.2), output: map[string]interface{}{"a": testArray}},
		{script: "a[3] = 1.1", input: map[string]interface{}{"a": testArray}, runOutput: float64(1.1), output: map[string]interface{}{"a": testArray}},
		{script: "a[3]", input: map[string]interface{}{"a": testArray}, runOutput: float64(1.1), output: map[string]interface{}{"a": testArray}},

		{script: "a[4]", input: map[string]interface{}{"a": testArray}, runOutput: "a", output: map[string]interface{}{"a": testArray}},
		{script: "a[4] = \"x\"", input: map[string]interface{}{"a": testArray}, runOutput: "x", output: map[string]interface{}{"a": testArray}},
		{script: "a[4]", input: map[string]interface{}{"a": testArray}, runOutput: "x", output: map[string]interface{}{"a": testArray}},
		{script: "a[4] = \"a\"", input: map[string]interface{}{"a": testArray}, runOutput: "a", output: map[string]interface{}{"a": testArray}},
		{script: "a[4]", input: map[string]interface{}{"a": testArray}, runOutput: "a", output: map[string]interface{}{"a": testArray}},
	}
	runTests(t, tests)
}

func TestMakeArrays(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{

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

		{script: "a = make([]bool, 0); a += true; a += false", types: map[string]interface{}{"bool": true}, runOutput: []bool{true, false}, output: map[string]interface{}{"a": []bool{true, false}}},
		{script: "a = make([]int32, 0); a += 1; a += 2", types: map[string]interface{}{"int32": int32(1)}, runOutput: []int32{int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{script: "a = make([]int64, 0); a += 1; a += 2", types: map[string]interface{}{"int64": int64(1)}, runOutput: []int64{int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		{script: "a = make([]float32, 0); a += 1.1; a += 2.2", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: []float32{float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{script: "a = make([]float64, 0); a += 1.1; a += 2.2", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: []float64{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{script: "a = make([]string, 0); a += \"a\"; a += \"b\"", types: map[string]interface{}{"string": "a"}, runOutput: []string{"a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}}},

		{script: "a = make([]bool, 2); a[0] = true; a[1] = false", types: map[string]interface{}{"bool": true}, runOutput: false, output: map[string]interface{}{"a": []bool{true, false}}},
		// TOFIX:
		// {script: "a = make([]int32, 2); a[0] = 1; a[1] = 2", types: map[string]interface{}{"int32": int32(1)}, runOutput: int32(2), output: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{script: "a = make([]int64, 2); a[0] = 1; a[1] = 2", types: map[string]interface{}{"int64": int64(1)}, runOutput: int64(2), output: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		// TOFIX:
		// {script: "a = make([]float32, 2); a[0] = 1.1; a[1] = 2.2", types: map[string]interface{}{"float32": float32(1.1)}, runOutput: float32(2.2), output: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{script: "a = make([]float64, 2); a[0] = 1.1; a[1] = 2.2", types: map[string]interface{}{"float64": float64(1.1)}, runOutput: float64(2.2), output: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{script: "a = make([]string, 2); a[0] = \"a\"; a[1] = \"b\"", types: map[string]interface{}{"string": "a"}, runOutput: "b", output: map[string]interface{}{"a": []string{"a", "b"}}},

		// remove ", nil" when/if syntax support it
		{script: "make([]boolA, nil)", types: map[string]interface{}{"boolA": []bool{}}, runOutput: [][]bool{}},
		{script: "make([]int32A, nil)", types: map[string]interface{}{"int32A": []int32{}}, runOutput: [][]int32{}},
		{script: "make([]int64A, nil)", types: map[string]interface{}{"int64A": []int64{}}, runOutput: [][]int64{}},
		{script: "make([]float32A, nil)", types: map[string]interface{}{"float32A": []float32{}}, runOutput: [][]float32{}},
		{script: "make([]float64A, nil)", types: map[string]interface{}{"float64A": []float64{}}, runOutput: [][]float64{}},
		{script: "make([]stringA, nil)", types: map[string]interface{}{"stringA": []string{}}, runOutput: [][]string{}},

		{script: "make([]array, 0)", types: map[string]interface{}{"array": []interface{}{}}, runOutput: [][]interface{}{}},
		{script: "a = make([]array, 0)", types: map[string]interface{}{"array": []interface{}{}}, runOutput: [][]interface{}{}, output: map[string]interface{}{"a": [][]interface{}{}}},
	}
	runTests(t, tests)
}

func TestArraySlice(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = [1, 2, 3]; a[:]", parseError: fmt.Errorf("syntax error")},

		{script: "a = [1, 2, 3]; a[true:2]", runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:true]", runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[-1:0]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:0]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:1]", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:2]", runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:3]", runOutput: []interface{}{int64(1), int64(2), int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:4]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[1:0]", runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:1]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:2]", runOutput: []interface{}{int64(2)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:3]", runOutput: []interface{}{int64(2), int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:4]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[2:1]", runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:2]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:3]", runOutput: []interface{}{int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:4]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[3:2]", runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[3:3]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[3:4]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[4:4]", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[-1:]", runError: fmt.Errorf("index out of range"), runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:]", runOutput: []interface{}{int64(1), int64(2), int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:]", runOutput: []interface{}{int64(2), int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[2:]", runOutput: []interface{}{int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[3:]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[4:]", runError: fmt.Errorf("index out of range"), runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[:-1]", runError: fmt.Errorf("index out of range"), runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:0]", runOutput: []interface{}{}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:1]", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:2]", runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:3]", runOutput: []interface{}{int64(1), int64(2), int64(3)}, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:4]", runError: fmt.Errorf("index out of range"), runOutput: nil, output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[true:2] = 4", runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:true] = 4", runError: fmt.Errorf("index must be a number"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[0:0] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:1] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:4] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:0] = 4", runError: fmt.Errorf("invalid slice index"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:4] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[-1:] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[0:] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[1:] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[4:] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{script: "a = [1, 2, 3]; a[:-1] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:0] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:1] = 4", runError: fmt.Errorf("slice cannot be assigned"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{script: "a = [1, 2, 3]; a[:4] = 4", runError: fmt.Errorf("index out of range"), output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
	}
	runTests(t, tests)
}

func TestArrayAppendArrays(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []bool{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true}, output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, false}, output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false}}},

		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{true}, output: map[string]interface{}{"a": []bool{}, "b": []interface{}{true}}},
		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true, true}, output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, true}}},
		{script: "b = [true]; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, true, false}, output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, true, false}}},

		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{}}, runOutput: []interface{}{true, false}, output: map[string]interface{}{"a": []bool{}, "b": []interface{}{true, false}}},
		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{true}}, runOutput: []interface{}{true, false, true}, output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, false, true}}},
		{script: "b = [true, false]; b += a", input: map[string]interface{}{"a": []bool{true, false}}, runOutput: []interface{}{true, false, true, false}, output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false, true, false}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int32(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int32(1), int32(2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int32(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int32(1), int32(2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), int64(2), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int64(2), int32(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), int64(2), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int64(2), int32(1), int32(2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int32(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int32(1), int32(2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), float64(2.2), int32(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{int64(1), float64(2.2)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{int64(1), float64(2.2), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), float64(2.2), int32(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), float64(2.2), int32(1), int32(2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{}}, runOutput: []interface{}{float64(1.1), int64(2)}, output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{1}}, runOutput: []interface{}{float64(1.1), int64(2), int32(1)}, output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int64(2), int32(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int32{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), int32(1), int32(2)}, output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int64(2), int32(1), int32(2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(1), int64(2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), int64(2), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(2), int64(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), int64(2), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2), int64(1), int64(2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(1), int64(2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), float64(2.2), int64(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{int64(1), float64(2.2)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{int64(1), float64(2.2), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), float64(2.2), int64(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), float64(2.2), int64(1), int64(2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{}}, runOutput: []interface{}{float64(1.1), int64(2)}, output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{1}}, runOutput: []interface{}{float64(1.1), int64(2), int64(1)}, output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(2), int64(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []int64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), int64(1), int64(2)}, output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(2), int64(1), int64(2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float32(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float32(1), float32(2)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float32(1.1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float32(1.1), float32(2.2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float32(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float32(1), float32(2)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float32(1.1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float32(1.1), float32(2.2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), int64(2), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), int64(2), float32(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), int64(2), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), int64(2), float32(1), float32(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), int64(2), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), int64(2), float32(1.1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float32(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float32(1), float32(2)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float32(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{int64(1), float64(2.2)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float64(2.2), float32(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float64(2.2), float32(1), float32(2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float64(2.2), float32(1.1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{}}, runOutput: []interface{}{float64(1.1), int64(2)}, output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1)}, output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), int64(2), float32(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1), float32(2)}, output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), int64(2), float32(1), float32(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1.1)}, output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), int64(2), float32(1.1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float32{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}, output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1), float64(2)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1)}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2)}}},

		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(1), float64(2)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(1.1)}}},
		{script: "b = [1]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(1.1), float64(2.2)}}},

		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1), int64(2)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), int64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), int64(2), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), int64(2), float64(1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), int64(2), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), int64(2), float64(1), float64(2)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), int64(2), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), int64(2), float64(1.1)}}},
		{script: "b = [1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(1), float64(2)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(1.1)}}},
		{script: "b = [1.1]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1)}}},
		{script: "b = [1.1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}}},

		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{int64(1), float64(2.2)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(2.2), float64(1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(2.2), float64(1), float64(2)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(2.2), float64(1.1)}}},
		{script: "b = [1, 2.2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}}},

		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{}}, runOutput: []interface{}{float64(1.1), int64(2)}, output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1)}, output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), int64(2), float64(1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1, 2}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1), float64(2)}, output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), int64(2), float64(1), float64(2)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1.1)}, output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), int64(2), float64(1.1)}}},
		{script: "b = [1.1, 2]; b += a", input: map[string]interface{}{"a": []float64{1.1, 2.2}}, runOutput: []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}, output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}}},

		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": []string{}, "b": []interface{}{}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a"}, output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a"}}},
		{script: "b = []; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b"}}},

		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{"a"}, output: map[string]interface{}{"a": []string{}, "b": []interface{}{"a"}}},
		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a", "a"}, output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "a"}}},
		{script: "b = [\"a\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "a", "b"}}},

		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{"a", "b"}, output: map[string]interface{}{"a": []string{}, "b": []interface{}{"a", "b"}}},
		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{"a", "b", "a"}, output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "b", "a"}}},
		{script: "b = [\"a\", \"b\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{"a", "b", "a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b", "a", "b"}}},

		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{}}, runOutput: []interface{}{int64(1), "a"}, output: map[string]interface{}{"a": []string{}, "b": []interface{}{int64(1), "a"}}},
		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{"a"}}, runOutput: []interface{}{int64(1), "a", "a"}, output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{int64(1), "a", "a"}}},
		{script: "b = [1, \"a\"]; b += a", input: map[string]interface{}{"a": []string{"a", "b"}}, runOutput: []interface{}{int64(1), "a", "a", "b"}, output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{int64(1), "a", "a", "b"}}},

		{script: "a = []; a += [nil, nil]; a += [nil, nil]", runOutput: []interface{}{interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil)}, output: map[string]interface{}{"a": []interface{}{interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil)}}},
		{script: "a = []; a += [true, false]; a += [false, true]", runOutput: []interface{}{interface{}(true), interface{}(false), interface{}(false), interface{}(true)}, output: map[string]interface{}{"a": []interface{}{interface{}(true), interface{}(false), interface{}(false), interface{}(true)}}},
		{script: "a = []; a += [1, 2]; a += [3, 4]", runOutput: []interface{}{interface{}(int64(1)), interface{}(int64(2)), interface{}(int64(3)), interface{}(int64(4))}, output: map[string]interface{}{"a": []interface{}{interface{}(int64(1)), interface{}(int64(2)), interface{}(int64(3)), interface{}(int64(4))}}},
		{script: "a = []; a += [1.1, 2.2]; a += [3.3, 4.4]", runOutput: []interface{}{interface{}(float64(1.1)), interface{}(float64(2.2)), interface{}(float64(3.3)), interface{}(float64(4.4))}, output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1)), interface{}(float64(2.2)), interface{}(float64(3.3)), interface{}(float64(4.4))}}},
		{script: "a = []; a += [\"a\", \"b\"]; a += [\"c\", \"d\"]", runOutput: []interface{}{interface{}("a"), interface{}("b"), interface{}("c"), interface{}("d")}, output: map[string]interface{}{"a": []interface{}{interface{}("a"), interface{}("b"), interface{}("c"), interface{}("d")}}},

		{script: "a = []; a += [[nil, nil]]; a += [[nil, nil]]", runOutput: []interface{}{[]interface{}{nil, nil}, []interface{}{nil, nil}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{nil, nil}, []interface{}{nil, nil}}}},
		{script: "a = []; a += [[true, false]]; a += [[false, true]]", runOutput: []interface{}{[]interface{}{true, false}, []interface{}{false, true}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{true, false}, []interface{}{false, true}}}},
		{script: "a = []; a += [[1, 2]]; a += [[3, 4]]", runOutput: []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{script: "a = []; a += [[1.1, 2.2]]; a += [[3.3, 4.4]]", runOutput: []interface{}{[]interface{}{float64(1.1), float64(2.2)}, []interface{}{float64(3.3), float64(4.4)}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{float64(1.1), float64(2.2)}, []interface{}{float64(3.3), float64(4.4)}}}},
		{script: "a = []; a += [[\"a\", \"b\"]]; a += [[\"c\", \"d\"]]", runOutput: []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}}, output: map[string]interface{}{"a": []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}}}},

		{script: "a = make(array2x); a += [[1, 2]]", types: map[string]interface{}{"array2x": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1), int64(2)}}},
		{script: "a = make(array2x); b = [1, 2]; a += [b]", types: map[string]interface{}{"array2x": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1), int64(2)}}},
		{script: "a = make(array2x); a += [[1, 2],[3, 4]]", types: map[string]interface{}{"array2x": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}},
		{script: "a = make(array2x); b = [1, 2]; a += [b]; b = [3, 4]; a += [b]", types: map[string]interface{}{"array2x": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}},

		{script: "a = make(array3x); a += [[[1, 2]]]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}}}},
		{script: "a = make(array3x); b = [[1, 2]]; a += [b]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}}}},
		{script: "a = make(array3x); b = [1, 2]; a += [[b]]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}}}},
		{script: "a = make(array3x); a += [[[1, 2],[3, 4]]]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{script: "a = make(array3x); b = [[1, 2],[3, 4]]; a += [b]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{script: "a = make(array3x); b = [1, 2]; c = [b]; b = [3, 4]; c += [b]; a += [c]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{script: "a = make(array3x); b = [1, 2]; c = []; c += [b]; b = [3, 4]; c += [b]; a += [c]", types: map[string]interface{}{"array3x": [][][]interface{}{}}, runOutput: [][][]interface{}{[][]interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
	}
	runTests(t, tests)
}

func TestMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = {}", runOutput: map[string]interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}", runOutput: map[string]interface{}{"b": nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}", runOutput: map[string]interface{}{"b": true}, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}", runOutput: map[string]interface{}{"b": int64(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}", runOutput: map[string]interface{}{"b": float64(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}", runOutput: map[string]interface{}{"b": "b"}, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a = {\"b\": {}}", runOutput: map[string]interface{}{"b": map[string]interface{}{}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{}}}},
		{script: "a = {\"b\": {\"c\": nil}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": nil}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": nil}}}},
		{script: "a = {\"b\": {\"c\": true}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": true}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": true}}}},
		{script: "a = {\"b\": {\"c\": 1}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}}},
		{script: "a = {\"b\": {\"c\": 1.1}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}}},
		{script: "a = {\"b\": {\"c\": \"c\"}}", runOutput: map[string]interface{}{"b": map[string]interface{}{"c": "c"}}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": "c"}}}},

		{script: "a = {\"b\": {}}; a.b", runOutput: map[string]interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{}}}},
		{script: "a = {\"b\": {\"c\": nil}}; a.b", runOutput: map[string]interface{}{"c": nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": nil}}}},
		{script: "a = {\"b\": {\"c\": true}}; a.b", runOutput: map[string]interface{}{"c": true}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": true}}}},
		{script: "a = {\"b\": {\"c\": 1}}; a.b", runOutput: map[string]interface{}{"c": int64(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": int64(1)}}}},
		{script: "a = {\"b\": {\"c\": 1.1}}; a.b", runOutput: map[string]interface{}{"c": float64(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": float64(1.1)}}}},
		{script: "a = {\"b\": {\"c\": \"c\"}}; a.b", runOutput: map[string]interface{}{"c": "c"}, output: map[string]interface{}{"a": map[string]interface{}{"b": map[string]interface{}{"c": "c"}}}},

		{script: "a = {\"b\": []}", runOutput: map[string]interface{}{"b": []interface{}{}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}", runOutput: map[string]interface{}{"b": []interface{}{nil}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}", runOutput: map[string]interface{}{"b": []interface{}{true}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}", runOutput: map[string]interface{}{"b": []interface{}{int64(1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}", runOutput: map[string]interface{}{"b": []interface{}{float64(1.1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"c\"]}", runOutput: map[string]interface{}{"b": []interface{}{"c"}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"c"}}}},

		{script: "a = {}; a.b", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}; a.b", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}; a.b", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}; a.b", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}; a.b", runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}; a.b", runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a = {}; a[\"b\"]", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a = {\"b\": nil}; a[\"b\"]", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a = {\"b\": true}; a[\"b\"]", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a = {\"b\": 1}; a[\"b\"]", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a = {\"b\": 1.1}; a[\"b\"]", runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a = {\"b\": \"b\"}; a[\"b\"]", runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: map[string]interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: map[string]interface{}{"b": nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: map[string]interface{}{"b": true}, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: map[string]interface{}{"b": int32(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: map[string]interface{}{"b": int64(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: map[string]interface{}{"b": float32(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: map[string]interface{}{"b": float64(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: map[string]interface{}{"b": "b"}, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: int32(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: float32(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a.b", input: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}, runOutput: true, output: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(2), output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(2), output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(2.2), output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(2.2), output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}},
		{script: "a.b", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "b", output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}},

		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, runOutput: reflect.Value{}, output: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runOutput: int32(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runOutput: float32(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a[0:1]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, runError: fmt.Errorf("type map does not support slice operation"), output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}, runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}},
		{script: "a[c]", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}},

		{script: "a.b = nil", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{script: "a.b = true", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{script: "a.b = 1", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a.b = 1.1", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{script: "a.b = \"b\"", input: map[string]interface{}{"a": map[string]interface{}{}}, runOutput: "b", output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{script: "a.b = true", input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, runOutput: true, output: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		// TOFIX:
		//		{script: "a.b = 3", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(3), output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{script: "a.b = 3", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(3), output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		// TOFIX:
		//		{script: "a.b = 3.3", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(3.3), output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{script: "a.b = 3.3", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(3.3), output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{script: "a.b = \"c\"", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "c", output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{script: "a[\"b\"] = true", input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, runOutput: true, output: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		// TOFIX:
		//		{script: "a[\"b\"] = 3", input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, runOutput: int32(3), output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{script: "a[\"b\"] = 3", input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, runOutput: int64(3), output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		// TOFIX:
		//		{script: "a[\"b\"] = 3.3", input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, runOutput: float32(3.3), output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{script: "a[\"b\"] = 3.3", input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, runOutput: float64(3.3), output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{script: "a[\"b\"] = \"c\"", input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, runOutput: "c", output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{script: "a[c] = \"x\"", input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, runError: fmt.Errorf("index type bool cannot be used for map index type string"), runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{script: "a[c] = \"x\"", input: map[string]interface{}{"a": map[bool]interface{}{true: "b"}, "c": true}, runOutput: "x", output: map[string]interface{}{"a": map[bool]interface{}{true: "x"}, "c": true}},

		// note if passed an uninitialized map there does not seem to be a way to update that map
		{script: "a", input: map[string]interface{}{"a": testMapEmpty}, runOutput: testMapEmpty, output: map[string]interface{}{"a": testMapEmpty}},
		{script: "a.b", input: map[string]interface{}{"a": testMapEmpty}, runOutput: nil, output: map[string]interface{}{"a": testMapEmpty}},
		{script: "a.b = 1", input: map[string]interface{}{"a": testMapEmpty}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a.b", input: map[string]interface{}{"a": testMapEmpty}, runOutput: nil, output: map[string]interface{}{"a": testMapEmpty}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": testMapEmpty}, runOutput: nil, output: map[string]interface{}{"a": testMapEmpty}},
		{script: "a[\"b\"] = 1", input: map[string]interface{}{"a": testMapEmpty}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{script: "a[\"b\"]", input: map[string]interface{}{"a": testMapEmpty}, runOutput: nil, output: map[string]interface{}{"a": testMapEmpty}},

		{script: "a", input: map[string]interface{}{"a": testMap}, runOutput: testMap, output: map[string]interface{}{"a": testMap}},
		{script: "a.a", input: map[string]interface{}{"a": testMap}, runOutput: nil, output: map[string]interface{}{"a": testMap}},
		{script: "a.a = true", input: map[string]interface{}{"a": testMap}, runOutput: true, output: map[string]interface{}{"a": testMap}},
		{script: "a.a", input: map[string]interface{}{"a": testMap}, runOutput: true, output: map[string]interface{}{"a": testMap}},
		{script: "a.a = nil", input: map[string]interface{}{"a": testMap}, runOutput: nil, output: map[string]interface{}{"a": testMap}},

		{script: "a", input: map[string]interface{}{"a": testMap}, runOutput: testMap, output: map[string]interface{}{"a": testMap}},
		{script: "a.b", input: map[string]interface{}{"a": testMap}, runOutput: true, output: map[string]interface{}{"a": testMap}},
		{script: "a.b = false", input: map[string]interface{}{"a": testMap}, runOutput: false, output: map[string]interface{}{"a": testMap}},
		{script: "a.b", input: map[string]interface{}{"a": testMap}, runOutput: false, output: map[string]interface{}{"a": testMap}},
		{script: "a.b = true", input: map[string]interface{}{"a": testMap}, runOutput: true, output: map[string]interface{}{"a": testMap}},

		{script: "a", input: map[string]interface{}{"a": testMap}, runOutput: testMap, output: map[string]interface{}{"a": testMap}},
		{script: "a.c", input: map[string]interface{}{"a": testMap}, runOutput: int64(1), output: map[string]interface{}{"a": testMap}},
		{script: "a.c = 2", input: map[string]interface{}{"a": testMap}, runOutput: int64(2), output: map[string]interface{}{"a": testMap}},
		{script: "a.c", input: map[string]interface{}{"a": testMap}, runOutput: int64(2), output: map[string]interface{}{"a": testMap}},
		{script: "a.c = 1", input: map[string]interface{}{"a": testMap}, runOutput: int64(1), output: map[string]interface{}{"a": testMap}},

		{script: "a", input: map[string]interface{}{"a": testMap}, runOutput: testMap, output: map[string]interface{}{"a": testMap}},
		{script: "a.d", input: map[string]interface{}{"a": testMap}, runOutput: float64(1.1), output: map[string]interface{}{"a": testMap}},
		{script: "a.d = 2.2", input: map[string]interface{}{"a": testMap}, runOutput: float64(2.2), output: map[string]interface{}{"a": testMap}},
		{script: "a.d", input: map[string]interface{}{"a": testMap}, runOutput: float64(2.2), output: map[string]interface{}{"a": testMap}},
		{script: "a.d = 1.1", input: map[string]interface{}{"a": testMap}, runOutput: float64(1.1), output: map[string]interface{}{"a": testMap}},

		{script: "a", input: map[string]interface{}{"a": testMap}, runOutput: testMap, output: map[string]interface{}{"a": testMap}},
		{script: "a.e", input: map[string]interface{}{"a": testMap}, runOutput: "e", output: map[string]interface{}{"a": testMap}},
		{script: "a.e = \"x\"", input: map[string]interface{}{"a": testMap}, runOutput: "x", output: map[string]interface{}{"a": testMap}},
		{script: "a.e", input: map[string]interface{}{"a": testMap}, runOutput: "x", output: map[string]interface{}{"a": testMap}},
		{script: "a.e = \"e\"", input: map[string]interface{}{"a": testMap}, runOutput: "e", output: map[string]interface{}{"a": testMap}},

		{script: "a = {\"b\": 1, \"c\": nil}", runOutput: map[string]interface{}{"b": int64(1), "c": nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.b", runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.c", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.d", runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},

		{script: "a = {\"b\": 1, \"c\": nil}; a == nil", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a != nil", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.b == nil", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.b != nil", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.c == nil", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.c != nil", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.d == nil", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.d != nil", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},

		{script: "a = {\"b\": 1, \"c\": nil}; a == 1", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a != 1", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.b == 1", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.b != 1", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.c == 1", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.c != 1", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.d == 1", runOutput: false, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
		{script: "a = {\"b\": 1, \"c\": nil}; a.d != 1", runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1), "c": nil}}},
	}
	runTests(t, tests)
}

func TestMakeMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "make(mapStringBool)", types: map[string]interface{}{"mapStringBool": map[string]bool{}}, runOutput: map[string]bool{}},
		{script: "make(mapStringInt32)", types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, runOutput: map[string]int32{}},
		{script: "make(mapStringInt64)", types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, runOutput: map[string]int64{}},
		{script: "make(mapStringFloat32)", types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, runOutput: map[string]float32{}},
		{script: "make(mapStringFloat64)", types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, runOutput: map[string]float64{}},
		{script: "make(mapStringString)", types: map[string]interface{}{"mapStringString": map[string]string{}}, runOutput: map[string]string{}},

		{script: "a = make(mapStringBool)", types: map[string]interface{}{"mapStringBool": map[string]bool{}}, runOutput: map[string]bool{}, output: map[string]interface{}{"a": map[string]bool{}}},
		{script: "a = make(mapStringInt32)", types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, runOutput: map[string]int32{}, output: map[string]interface{}{"a": map[string]int32{}}},
		{script: "a = make(mapStringInt64)", types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, runOutput: map[string]int64{}, output: map[string]interface{}{"a": map[string]int64{}}},
		{script: "a = make(mapStringFloat32)", types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, runOutput: map[string]float32{}, output: map[string]interface{}{"a": map[string]float32{}}},
		{script: "a = make(mapStringFloat64)", types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, runOutput: map[string]float64{}, output: map[string]interface{}{"a": map[string]float64{}}},
		{script: "a = make(mapStringString)", types: map[string]interface{}{"mapStringString": map[string]string{}}, runOutput: map[string]string{}, output: map[string]interface{}{"a": map[string]string{}}},

		{script: "a = make(mapStringBool); a[\"b\"] = true", types: map[string]interface{}{"mapStringBool": map[string]bool{"b": true}}, runOutput: true, output: map[string]interface{}{"a": map[string]bool{"b": true}}},
		// TOFIX:
		//		{script: "a = make(mapStringInt32); a[\"b\"] = 1", types: map[string]interface{}{"mapStringInt32": map[string]int32{"b": int32(1)}}, runOutput: int32(1), output: map[string]interface{}{"a": map[string]int32{"b": int32(1)}}},
		{script: "a = make(mapStringInt64); a[\"b\"] = 1", types: map[string]interface{}{"mapStringInt64": map[string]int64{"b": int64(1)}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]int64{"b": int64(1)}}},
		// TOFIX:
		//		{script: "a = make(mapStringFloat32); a[\"b\"] = 1.1", types: map[string]interface{}{"mapStringFloat32": map[string]float32{"b": float32(1.1)}}, runOutput: float32(1.1), output: map[string]interface{}{"a": map[string]float32{"b": float32(1.1)}}},
		{script: "a = make(mapStringFloat64); a[\"b\"] = 1.1", types: map[string]interface{}{"mapStringFloat64": map[string]float64{"b": float64(1.1)}}, runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]float64{"b": float64(1.1)}}},
		{script: "a = make(mapStringString); a[\"b\"] = \"b\"", types: map[string]interface{}{"mapStringString": map[string]string{"b": "b"}}, runOutput: "b", output: map[string]interface{}{"a": map[string]string{"b": "b"}}},
	}
	runTests(t, tests)
}

func TestArraysAndMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a = [{\"b\": nil}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(nil)}}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(nil)}}}},
		{script: "a = [{\"b\": true}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(true)}}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(true)}}}},
		{script: "a = [{\"b\": 1}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}}},
		{script: "a = [{\"b\": 1.1}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}}},
		{script: "a = [{\"b\": \"b\"}]", runOutput: []interface{}{map[string]interface{}{"b": interface{}("b")}}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}("b")}}}},

		{script: "a = [{\"b\": nil}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(nil)}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(nil)}}}},
		{script: "a = [{\"b\": true}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(true)}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(true)}}}},
		{script: "a = [{\"b\": 1}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(int64(1))}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(int64(1))}}}},
		{script: "a = [{\"b\": 1.1}]; a[0]", runOutput: map[string]interface{}{"b": interface{}(float64(1.1))}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}(float64(1.1))}}}},
		{script: "a = [{\"b\": \"b\"}]; a[0]", runOutput: map[string]interface{}{"b": interface{}("b")}, output: map[string]interface{}{"a": []interface{}{map[string]interface{}{"b": interface{}("b")}}}},

		{script: "a = {\"b\": []}", runOutput: map[string]interface{}{"b": []interface{}{}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}", runOutput: map[string]interface{}{"b": []interface{}{nil}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}", runOutput: map[string]interface{}{"b": []interface{}{true}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}", runOutput: map[string]interface{}{"b": []interface{}{int64(1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}", runOutput: map[string]interface{}{"b": []interface{}{float64(1.1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"b\"]}", runOutput: map[string]interface{}{"b": []interface{}{"b"}}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"b"}}}},

		{script: "a = {\"b\": []}; a.b", runOutput: []interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{}}}},
		{script: "a = {\"b\": [nil]}; a.b", runOutput: []interface{}{nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{nil}}}},
		{script: "a = {\"b\": [true]}; a.b", runOutput: []interface{}{true}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{true}}}},
		{script: "a = {\"b\": [1]}; a.b", runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a = {\"b\": [1.1]}; a.b", runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a = {\"b\": [\"b\"]}; a.b", runOutput: []interface{}{"b"}, output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{"b"}}}},

		{script: "a.b = []", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{}}}},
		{script: "a.b = [nil]", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{interface{}(nil)}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{interface{}(nil)}}}},
		{script: "a.b = [true]", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{true}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{true}}}},
		{script: "a.b = [1]", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{int64(1)}}}},
		{script: "a.b = [1.1]", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{float64(1.1)}}}},
		{script: "a.b = [\"b\"]", input: map[string]interface{}{"a": map[string][]interface{}{}}, runOutput: []interface{}{"b"}, output: map[string]interface{}{"a": map[string][]interface{}{"b": []interface{}{"b"}}}},

		{script: "b[0] = [nil]; a.b = b", input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{interface{}(nil)}}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{interface{}(nil)}}}, "b": [][]interface{}{[]interface{}{interface{}(nil)}}}},
		{script: "b[0] = [true]; a.b = b", input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{true}}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{true}}}, "b": [][]interface{}{[]interface{}{true}}}},
		{script: "b[0] = [1]; a.b = b", input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1)}}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}, "b": [][]interface{}{[]interface{}{int64(1)}}}},
		{script: "b[0] = [1.1]; a.b = b", input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{float64(1.1)}}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{float64(1.1)}}}, "b": [][]interface{}{[]interface{}{float64(1.1)}}}},
		{script: "b[0] = [\"b\"]; a.b = b", input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{"b"}}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{"b"}}}, "b": [][]interface{}{[]interface{}{"b"}}}},

		{script: "a", input: map[string]interface{}{"a": map[string][][]interface{}{}}, runOutput: map[string][][]interface{}{}, output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{script: "a.b = 1", input: map[string]interface{}{"a": map[string][][]interface{}{}}, runError: fmt.Errorf("type int64 cannot be assigned to type [][]interface {} for map"), runOutput: nil, output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{script: "a[\"b\"] = 1", input: map[string]interface{}{"a": map[string][][]interface{}{}}, runError: fmt.Errorf("type int64 cannot be assigned to type [][]interface {} for map"), runOutput: nil, output: map[string]interface{}{"a": map[string][][]interface{}{}}},

		{script: "a = {}; a.b = []; a.b += 1; a.b[0] = {}; a.b[0].c = []; a.b[0].c += 2; a.b[0].c[0]", runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"b": []interface{}{map[string]interface{}{"c": []interface{}{int64(2)}}}}}},

		{script: "a = {}; a.b = b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runError: fmt.Errorf("index out of range"), runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},

		{script: "a = {}; a.b = b; a.b[0] = []", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [nil]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{nil}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [nil]; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{nil}}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{nil}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [true]; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{true}}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{true}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1]; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{int64(1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1.1]; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{float64(1.1)}}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{float64(1.1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [\"c\"]; a.b", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{[]interface{}{"c"}}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{"c"}}}, "b": [][]interface{}{}}},

		{script: "a = {}; a.b = b; a.b[0] = [nil]; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{nil}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{nil}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [true]; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{true}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{true}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1]; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1.1]; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{float64(1.1)}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{float64(1.1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [\"c\"]; a.b[0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: []interface{}{"c"}, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{"c"}}}, "b": [][]interface{}{}}},

		{script: "a = {}; a.b = b; a.b[0] = [nil]; a.b[0][0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{nil}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [true]; a.b[0][0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{true}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1]; a.b[0][0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1.1]; a.b[0][0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: float64(1.1), output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{float64(1.1)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [\"c\"]; a.b[0][0]", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: "c", output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{"c"}}}, "b": [][]interface{}{}}},

		{script: "a = {}; a.b = b; a.b[0] = [nil]; a.b[0][1] = nil", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: nil, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{nil, nil}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [true]; a.b[0][1] = true", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: true, output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{true, true}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1]; a.b[0][1] = 2", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: int64(2), output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{int64(1), int64(2)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [1.1]; a.b[0][1] = 2.2", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: float64(2.2), output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{float64(1.1), float64(2.2)}}}, "b": [][]interface{}{}}},
		{script: "a = {}; a.b = b; a.b[0] = [\"c\"]; a.b[0][1] = \"d\"", input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: "d", output: map[string]interface{}{"a": map[string]interface{}{"b": [][]interface{}{[]interface{}{"c", "d"}}}, "b": [][]interface{}{}}},
	}
	runTests(t, tests)
}

func TestMakeArraysAndMaps(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "make(mapArray2x)", types: map[string]interface{}{"mapArray2x": map[string][][]interface{}{}}, runOutput: map[string][][]interface{}{}},
		{script: "a = make(mapArray2x)", types: map[string]interface{}{"mapArray2x": map[string][][]interface{}{}}, runOutput: map[string][][]interface{}{}, output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{script: "a = make(mapArray2x); a", types: map[string]interface{}{"mapArray2x": map[string][][]interface{}{}}, runOutput: map[string][][]interface{}{}, output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{script: "a = make(mapArray2x); a.b = b", types: map[string]interface{}{"mapArray2x": map[string][][]interface{}{}}, input: map[string]interface{}{"b": [][]interface{}{}}, runOutput: [][]interface{}{}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},

		{script: "b = make(array2x); a.b = b", types: map[string]interface{}{"array2x": [][]interface{}{}}, input: map[string]interface{}{"a": map[string][][]interface{}{}}, runOutput: [][]interface{}(nil), output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}(nil)}}},
		{script: "b = make(array2x); a.b = b; a.b[0] = [1]", types: map[string]interface{}{"array2x": [][]interface{}{}}, input: map[string]interface{}{"a": map[string][][]interface{}{}}, runOutput: []interface{}{int64(1)}, output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}}},
		{script: "b = make(array2x); a.b = b; a.b[0] = []; a.b[0][0] = 1", types: map[string]interface{}{"array2x": [][]interface{}{}}, input: map[string]interface{}{"a": map[string][][]interface{}{}}, runOutput: int64(1), output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{int64(1)}}}}},
		{script: "b = make(array2x); a.b = b; a.b[0] = [1]; a.b[0][1] = 2", types: map[string]interface{}{"array2x": [][]interface{}{}}, input: map[string]interface{}{"a": map[string][][]interface{}{}}, runOutput: int64(2), output: map[string]interface{}{"a": map[string][][]interface{}{"b": [][]interface{}{[]interface{}{int64(1), int64(2)}}}}},
	}
	runTests(t, tests)
}
