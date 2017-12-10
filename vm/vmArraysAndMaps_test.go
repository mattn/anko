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
