package vm

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
)

func TestSlices(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `[1++]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `1++[0]`, RunError: fmt.Errorf("invalid operation")},

		{Script: `[]`, RunOutput: []interface{}{}},
		{Script: `[nil]`, RunOutput: []interface{}{nil}},
		{Script: `[true]`, RunOutput: []interface{}{true}},
		{Script: `["a"]`, RunOutput: []interface{}{"a"}},
		{Script: `[1]`, RunOutput: []interface{}{int64(1)}},
		{Script: `[1.1]`, RunOutput: []interface{}{float64(1.1)}},

		{Script: `a = []; a.b`, RunError: fmt.Errorf("type slice does not support member operation")},
		{Script: `a = []; a.b = 1`, RunError: fmt.Errorf("type slice does not support member operation")},

		{Script: `a = []`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{}}},
		{Script: `a = [nil]`, RunOutput: []interface{}{interface{}(nil)}, Output: map[string]interface{}{"a": []interface{}{interface{}(nil)}}},
		{Script: `a = [true]`, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": []interface{}{true}}},
		{Script: `a = [1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = [1.1]`, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{Script: `a = ["a"]`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []interface{}{"a"}}},

		{Script: `a = [[]]`, RunOutput: []interface{}{[]interface{}{}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{}}}},
		{Script: `a = [[nil]]`, RunOutput: []interface{}{[]interface{}{interface{}(nil)}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{interface{}(nil)}}}},
		{Script: `a = [[true]]`, RunOutput: []interface{}{[]interface{}{true}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{true}}}},
		{Script: `a = [[1]]`, RunOutput: []interface{}{[]interface{}{int64(1)}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1)}}}},
		{Script: `a = [[1.1]]`, RunOutput: []interface{}{[]interface{}{float64(1.1)}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{float64(1.1)}}}},
		{Script: `a = [["a"]]`, RunOutput: []interface{}{[]interface{}{"a"}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{"a"}}}},

		{Script: `a = []; a += nil`, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": []interface{}{nil}}},
		{Script: `a = []; a += true`, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": []interface{}{true}}},
		{Script: `a = []; a += 1`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = []; a += 1.1`, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{Script: `a = []; a += "a"`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []interface{}{"a"}}},

		{Script: `a = []; a += []`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{}}},
		{Script: `a = []; a += [nil]`, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": []interface{}{nil}}},
		{Script: `a = []; a += [true]`, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": []interface{}{true}}},
		{Script: `a = []; a += [1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = []; a += [1.1]`, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []interface{}{float64(1.1)}}},
		{Script: `a = []; a += ["a"]`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []interface{}{"a"}}},

		{Script: `a = [0]; a[0]++`, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = [[0]]; a[0][0]++`, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1)}}}},

		{Script: `a = [2]; a[0]--`, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = [[2]]; a[0][0]--`, RunOutput: int64(1), Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1)}}}},

		{Script: `a`, Input: map[string]interface{}{"a": []bool{}}, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}}},
		{Script: `a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []int32{}, Output: map[string]interface{}{"a": []int32{}}},
		{Script: `a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []int64{}, Output: map[string]interface{}{"a": []int64{}}},
		{Script: `a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []float32{}, Output: map[string]interface{}{"a": []float32{}}},
		{Script: `a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []float64{}, Output: map[string]interface{}{"a": []float64{}}},
		{Script: `a`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []string{}, Output: map[string]interface{}{"a": []string{}}},

		{Script: `a`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []int32{1, 2}, Output: map[string]interface{}{"a": []int32{1, 2}}},
		{Script: `a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []int64{1, 2}, Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []float32{1.1, 2.2}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []float64{1.1, 2.2}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: []string{"a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a += true`, Input: map[string]interface{}{"a": []bool{}}, RunOutput: []bool{true}, Output: map[string]interface{}{"a": []bool{true}}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []int32{1}, Output: map[string]interface{}{"a": []int32{1}}},
		{Script: `a += 1.1`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []int32{1}, Output: map[string]interface{}{"a": []int32{1}}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []int64{1}, Output: map[string]interface{}{"a": []int64{1}}},
		{Script: `a += 1.1`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []int64{1}, Output: map[string]interface{}{"a": []int64{1}}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []float32{1}, Output: map[string]interface{}{"a": []float32{1}}},
		{Script: `a += 1.1`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []float32{1.1}, Output: map[string]interface{}{"a": []float32{1.1}}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []float64{1}, Output: map[string]interface{}{"a": []float64{1}}},
		{Script: `a += 1.1`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []float64{1.1}, Output: map[string]interface{}{"a": []float64{1.1}}},
		{Script: `a += "a"`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []string{"a"}, Output: map[string]interface{}{"a": []string{"a"}}},
		{Script: `a += 97`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []string{"a"}, Output: map[string]interface{}{"a": []string{"a"}}},

		{Script: `a[0]`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: true, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: int32(1), Output: map[string]interface{}{"a": []int32{1, 2}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: int64(1), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: "a", Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a[1]`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: false, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: int32(2), Output: map[string]interface{}{"a": []int32{1, 2}}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: int64(2), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: float32(2.2), Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a[0]`, Input: map[string]interface{}{"a": []bool{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []bool{}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []int32{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int32{}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []int64{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int64{}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []float32{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float32{}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []float64{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float64{}}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": []string{}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []string{}}},

		{Script: `a[1] = true`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: true, Output: map[string]interface{}{"a": []bool{true, true}}},
		{Script: `a[1] = 3`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: int64(3), Output: map[string]interface{}{"a": []int32{1, 3}}},
		{Script: `a[1] = 3`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: int64(3), Output: map[string]interface{}{"a": []int64{1, 3}}},
		{Script: `a[1] = 3.3`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []float32{1.1, 3.3}}},
		{Script: `a[1] = 3.3`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []float64{1.1, 3.3}}},
		{Script: `a[1] = "c"`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: "c", Output: map[string]interface{}{"a": []string{"a", "c"}}},

		{Script: `a = []; a[0]`, RunError: fmt.Errorf("index out of range")},
		{Script: `a = []; a[-1]`, RunError: fmt.Errorf("index out of range")},
		{Script: `a = []; a[1] = 1`, RunError: fmt.Errorf("index out of range")},
		{Script: `a = []; a[-1] = 1`, RunError: fmt.Errorf("index out of range")},

		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": nil}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": true}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": int(1)}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": float32(1.1)}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": "1"}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarBoolP}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarInt32P}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarInt64P}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarFloat32P}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarFloat64P}, RunOutput: int64(2), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a]`, Input: map[string]interface{}{"a": testVarStringP}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": nil}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": true}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": int(1)}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": float32(1.1)}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": "1"}, RunOutput: int64(3), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(3)}}},
		{Script: `b = [1, 2]; b[a] = 3`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"b": []interface{}{int64(1), int64(2)}}},

		{Script: `a`, Input: map[string]interface{}{"a": testSliceEmpty}, RunOutput: testSliceEmpty, Output: map[string]interface{}{"a": testSliceEmpty}},
		{Script: `a += []`, Input: map[string]interface{}{"a": testSliceEmpty}, RunOutput: []interface{}(nil), Output: map[string]interface{}{"a": testSliceEmpty}},

		{Script: `a`, Input: map[string]interface{}{"a": testSlice}, RunOutput: testSlice, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: nil, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[0] = 1`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[0] = nil`, Input: map[string]interface{}{"a": testSlice}, RunOutput: nil, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[0]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: nil, Output: map[string]interface{}{"a": testSlice}},

		{Script: `a[1]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: true, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[1] = false`, Input: map[string]interface{}{"a": testSlice}, RunOutput: false, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: false, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[1] = true`, Input: map[string]interface{}{"a": testSlice}, RunOutput: true, Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[1]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: true, Output: map[string]interface{}{"a": testSlice}},

		{Script: `a[2]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[2] = 2`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(2), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(2), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[2] = 1`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: int64(1), Output: map[string]interface{}{"a": testSlice}},

		{Script: `a[3]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[3] = 2.2`, Input: map[string]interface{}{"a": testSlice}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[3]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[3] = 1.1`, Input: map[string]interface{}{"a": testSlice}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[3]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": testSlice}},

		{Script: `a[4]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: "a", Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[4] = "x"`, Input: map[string]interface{}{"a": testSlice}, RunOutput: "x", Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[4]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: "x", Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[4] = "a"`, Input: map[string]interface{}{"a": testSlice}, RunOutput: "a", Output: map[string]interface{}{"a": testSlice}},
		{Script: `a[4]`, Input: map[string]interface{}{"a": testSlice}, RunOutput: "a", Output: map[string]interface{}{"a": testSlice}},

		{Script: `a[0][0] = true`, Input: map[string]interface{}{"a": []interface{}{[]string{"a"}}}, RunError: fmt.Errorf("type bool cannot be assigned to type string for slice index"), Output: map[string]interface{}{"a": []interface{}{[]string{"a"}}}},
		{Script: `a[0][0] = "a"`, Input: map[string]interface{}{"a": []interface{}{[]bool{true}}}, RunError: fmt.Errorf("type string cannot be assigned to type bool for slice index"), Output: map[string]interface{}{"a": []interface{}{[]bool{true}}}},

		{Script: `a[0][0] = b[0][0]`, Input: map[string]interface{}{"a": []interface{}{[]bool{true}}, "b": []interface{}{[]string{"b"}}}, RunError: fmt.Errorf("type string cannot be assigned to type bool for slice index"), Output: map[string]interface{}{"a": []interface{}{[]bool{true}}}},
		{Script: `b[0][0] = a[0][0]`, Input: map[string]interface{}{"a": []interface{}{[]bool{true}}, "b": []interface{}{[]string{"b"}}}, RunError: fmt.Errorf("type bool cannot be assigned to type string for slice index"), Output: map[string]interface{}{"a": []interface{}{[]bool{true}}}},

		{Script: `a = make([][]bool); a[0] = make([]bool); a[0] = [true, 1]`, RunError: fmt.Errorf("type []interface {} cannot be assigned to type []bool for slice index"), Output: map[string]interface{}{"a": [][]bool{{}}}},
		{Script: `a = make([][]bool); a[0] = make([]bool); a[0] = [true, false]`, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": [][]bool{{true, false}}}},

		{Script: `a = make([][][]bool); a[0] = make([][]bool); a[0][0] = make([]bool); a[0] = [[true, 1]]`, RunError: fmt.Errorf("type []interface {} cannot be assigned to type [][]bool for slice index"), Output: map[string]interface{}{"a": [][][]bool{{{}}}}},
		{Script: `a = make([][][]bool); a[0] = make([][]bool); a[0][0] = make([]bool); a[0] = [[true, false]]`, RunOutput: []interface{}{[]interface{}{true, false}}, Output: map[string]interface{}{"a": [][][]bool{{{true, false}}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestSlicesAutoAppend(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a[2]`, Input: map[string]interface{}{"a": []bool{true, false}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int32{1, 2}}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a[2]`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a[2] = true`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: true, Output: map[string]interface{}{"a": []bool{true, false, true}}},
		{Script: `a[2] = 3`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: int64(3), Output: map[string]interface{}{"a": []int32{1, 2, 3}}},
		{Script: `a[2] = 3`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: int64(3), Output: map[string]interface{}{"a": []int64{1, 2, 3}}},
		{Script: `a[2] = 3.3`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []float32{1.1, 2.2, 3.3}}},
		{Script: `a[2] = 3.3`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []float64{1.1, 2.2, 3.3}}},
		{Script: `a[2] = "c"`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: "c", Output: map[string]interface{}{"a": []string{"a", "b", "c"}}},

		{Script: `a[2] = 3.3`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []int32{1, 2, 3}}},
		{Script: `a[2] = 3.3`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": []int64{1, 2, 3}}},

		{Script: `a[3] = true`, Input: map[string]interface{}{"a": []bool{true, false}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a[3] = 4`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int32{1, 2}}},
		{Script: `a[3] = 4`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[3] = 4.4`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a[3] = 4.4`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a[3] = "d"`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []bool{true, false}}, Output: map[string]interface{}{"a": []bool{true, false, false}}},
		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []int32{1, 2}}, Output: map[string]interface{}{"a": []int32{1, 2, 0}}},
		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []int64{1, 2}}, Output: map[string]interface{}{"a": []int64{1, 2, 0}}},
		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []float32{1.5, 2.5}}, Output: map[string]interface{}{"a": []float32{1.5, 2.5, 0}}},
		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []float64{1.5, 2.5}}, Output: map[string]interface{}{"a": []float64{1.5, 2.5, 0}}},
		{Script: `a[2] = nil`, Input: map[string]interface{}{"a": []string{"a", "b"}}, Output: map[string]interface{}{"a": []string{"a", "b", ""}}},

		{Script: `a[2] = "a"`, Input: map[string]interface{}{"a": []int16{1, 2}}, RunError: fmt.Errorf("type string cannot be assigned to type int16 for slice index"), Output: map[string]interface{}{"a": []int16{1, 2}}},
		{Script: `a[2] = true`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunError: fmt.Errorf("type bool cannot be assigned to type int64 for slice index"), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[2] = "a"`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunError: fmt.Errorf("type string cannot be assigned to type int64 for slice index"), Output: map[string]interface{}{"a": []int64{1, 2}}},
		{Script: `a[2] = true`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunError: fmt.Errorf("type bool cannot be assigned to type float32 for slice index"), Output: map[string]interface{}{"a": []float32{1.1, 2.2}}},
		{Script: `a[2] = "a"`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunError: fmt.Errorf("type string cannot be assigned to type float64 for slice index"), Output: map[string]interface{}{"a": []float64{1.1, 2.2}}},
		{Script: `a[2] = true`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunError: fmt.Errorf("type bool cannot be assigned to type string for slice index"), Output: map[string]interface{}{"a": []string{"a", "b"}}},
		{Script: `a[2] = 1.1`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunError: fmt.Errorf("type float64 cannot be assigned to type string for slice index"), Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: [][]interface{}{}, Output: map[string]interface{}{"a": [][]interface{}{}}},
		{Script: `a[0] = []`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": [][]interface{}{{}}}},
		{Script: `a[0] = [nil]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{Script: `a[0] = [true]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{Script: `a[0] = [1]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{Script: `a[0] = [1.1]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},
		{Script: `a[0] = ["b"]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: []interface{}{"b"}, Output: map[string]interface{}{"a": [][]interface{}{{"b"}}}},

		{Script: `a[0] = [nil]; a[0][0]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": [][]interface{}{{nil}}}},
		{Script: `a[0] = [true]; a[0][0]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: true, Output: map[string]interface{}{"a": [][]interface{}{{true}}}},
		{Script: `a[0] = [1]; a[0][0]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": [][]interface{}{{int64(1)}}}},
		{Script: `a[0] = [1.1]; a[0][0]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": [][]interface{}{{float64(1.1)}}}},
		{Script: `a[0] = ["b"]; a[0][0]`, Input: map[string]interface{}{"a": [][]interface{}{}}, RunOutput: "b", Output: map[string]interface{}{"a": [][]interface{}{{"b"}}}},

		{Script: `a = make([]bool); a[0] = 1`, RunError: fmt.Errorf("type int64 cannot be assigned to type bool for slice index"), Output: map[string]interface{}{"a": []bool{}}},
		{Script: `a = make([]bool); a[0] = true; a[1] = false`, RunOutput: false, Output: map[string]interface{}{"a": []bool{true, false}}},

		{Script: `a = make([][]bool); a[0] = [true, 1]`, RunError: fmt.Errorf("type []interface {} cannot be assigned to type []bool for slice index"), Output: map[string]interface{}{"a": [][]bool{}}},
		{Script: `a = make([][]bool); a[0] = [true, false]`, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": [][]bool{{true, false}}}},

		{Script: `a = make([][][]bool); a[0] = [[true, 1]]`, RunError: fmt.Errorf("type []interface {} cannot be assigned to type [][]bool for slice index"), Output: map[string]interface{}{"a": [][][]bool{}}},
		{Script: `a = make([][][]bool); a[0] = [[true, false]]`, RunOutput: []interface{}{[]interface{}{true, false}}, Output: map[string]interface{}{"a": [][][]bool{{{true, false}}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeSlices(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `make([]foo)`, RunError: fmt.Errorf("undefined type 'foo'")},

		{Script: `make([]nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make([][]nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `make([][][]nilT)`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},

		{Script: `make([]bool, 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `make([]bool, 0, 1++)`, RunError: fmt.Errorf("invalid operation")},

		// spaces and/or newlines
		{Script: `[]`, RunOutput: []interface{}{}},
		{Script: `[ ]`, RunOutput: []interface{}{}},
		{Script: `[
]`, RunOutput: []interface{}{}},
		{Script: `[ 
 ]`, RunOutput: []interface{}{}},

		{Script: `make(slice2x)`, Types: map[string]interface{}{"slice2x": [][]interface{}{}}, RunOutput: [][]interface{}{}},

		{Script: `make([]bool)`, RunOutput: []bool{}},
		{Script: `make([]int32)`, RunOutput: []int32{}},
		{Script: `make([]int64)`, RunOutput: []int64{}},
		{Script: `make([]float32)`, RunOutput: []float32{}},
		{Script: `make([]float64)`, RunOutput: []float64{}},
		{Script: `make([]string)`, RunOutput: []string{}},

		{Script: `make([]bool, 0)`, RunOutput: []bool{}},
		{Script: `make([]int32, 0)`, RunOutput: []int32{}},
		{Script: `make([]int64, 0)`, RunOutput: []int64{}},
		{Script: `make([]float32, 0)`, RunOutput: []float32{}},
		{Script: `make([]float64, 0)`, RunOutput: []float64{}},
		{Script: `make([]string, 0)`, RunOutput: []string{}},

		{Script: `make([]bool, 2)`, RunOutput: []bool{false, false}},
		{Script: `make([]int32, 2)`, RunOutput: []int32{int32(0), int32(0)}},
		{Script: `make([]int64, 2)`, RunOutput: []int64{int64(0), int64(0)}},
		{Script: `make([]float32, 2)`, RunOutput: []float32{float32(0), float32(0)}},
		{Script: `make([]float64, 2)`, RunOutput: []float64{float64(0), float64(0)}},
		{Script: `make([]string, 2)`, RunOutput: []string{"", ""}},

		{Script: `make([]bool, 0, 2)`, RunOutput: []bool{}},
		{Script: `make([]int32, 0, 2)`, RunOutput: []int32{}},
		{Script: `make([]int64, 0, 2)`, RunOutput: []int64{}},
		{Script: `make([]float32, 0, 2)`, RunOutput: []float32{}},
		{Script: `make([]float64, 0, 2)`, RunOutput: []float64{}},
		{Script: `make([]string, 0, 2)`, RunOutput: []string{}},

		{Script: `make([]bool, 2, 2)`, RunOutput: []bool{false, false}},
		{Script: `make([]int32, 2, 2)`, RunOutput: []int32{int32(0), int32(0)}},
		{Script: `make([]int64, 2, 2)`, RunOutput: []int64{int64(0), int64(0)}},
		{Script: `make([]float32, 2, 2)`, RunOutput: []float32{float32(0), float32(0)}},
		{Script: `make([]float64, 2, 2)`, RunOutput: []float64{float64(0), float64(0)}},
		{Script: `make([]string, 2, 2)`, RunOutput: []string{"", ""}},

		{Script: `a = make([]bool, 0); a += true; a += false`, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a = make([]int32, 0); a += 1; a += 2`, RunOutput: []int32{int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{Script: `a = make([]int64, 0); a += 1; a += 2`, RunOutput: []int64{int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		{Script: `a = make([]float32, 0); a += 1.1; a += 2.2`, RunOutput: []float32{float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{Script: `a = make([]float64, 0); a += 1.1; a += 2.2`, RunOutput: []float64{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{Script: `a = make([]string, 0); a += "a"; a += "b"`, RunOutput: []string{"a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `a = make([]bool, 2); a[0] = true; a[1] = false`, RunOutput: false, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a = make([]int32, 2); a[0] = 1; a[1] = 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": []int32{int32(1), int32(2)}}},
		{Script: `a = make([]int64, 2); a[0] = 1; a[1] = 2`, RunOutput: int64(2), Output: map[string]interface{}{"a": []int64{int64(1), int64(2)}}},
		{Script: `a = make([]float32, 2); a[0] = 1.1; a[1] = 2.2`, RunOutput: float64(2.2), Output: map[string]interface{}{"a": []float32{float32(1.1), float32(2.2)}}},
		{Script: `a = make([]float64, 2); a[0] = 1.1; a[1] = 2.2`, RunOutput: float64(2.2), Output: map[string]interface{}{"a": []float64{float64(1.1), float64(2.2)}}},
		{Script: `a = make([]string, 2); a[0] = "a"; a[1] = "b"`, RunOutput: "b", Output: map[string]interface{}{"a": []string{"a", "b"}}},

		{Script: `make([]boolA)`, Types: map[string]interface{}{"boolA": []bool{}}, RunOutput: [][]bool{}},
		{Script: `make([]int32A)`, Types: map[string]interface{}{"int32A": []int32{}}, RunOutput: [][]int32{}},
		{Script: `make([]int64A)`, Types: map[string]interface{}{"int64A": []int64{}}, RunOutput: [][]int64{}},
		{Script: `make([]float32A)`, Types: map[string]interface{}{"float32A": []float32{}}, RunOutput: [][]float32{}},
		{Script: `make([]float64A)`, Types: map[string]interface{}{"float64A": []float64{}}, RunOutput: [][]float64{}},
		{Script: `make([]stringA)`, Types: map[string]interface{}{"stringA": []string{}}, RunOutput: [][]string{}},

		{Script: `make([]slice)`, Types: map[string]interface{}{"slice": []interface{}{}}, RunOutput: [][]interface{}{}},
		{Script: `a = make([]slice)`, Types: map[string]interface{}{"slice": []interface{}{}}, RunOutput: [][]interface{}{}, Output: map[string]interface{}{"a": [][]interface{}{}}},

		{Script: `make([][]bool)`, RunOutput: [][]bool{}},
		{Script: `make([][]int32)`, RunOutput: [][]int32{}},
		{Script: `make([][]int64)`, RunOutput: [][]int64{}},
		{Script: `make([][]float32)`, RunOutput: [][]float32{}},
		{Script: `make([][]float64)`, RunOutput: [][]float64{}},
		{Script: `make([][]string)`, RunOutput: [][]string{}},

		{Script: `make([][]bool)`, RunOutput: [][]bool{}},
		{Script: `make([][]int32)`, RunOutput: [][]int32{}},
		{Script: `make([][]int64)`, RunOutput: [][]int64{}},
		{Script: `make([][]float32)`, RunOutput: [][]float32{}},
		{Script: `make([][]float64)`, RunOutput: [][]float64{}},
		{Script: `make([][]string)`, RunOutput: [][]string{}},

		{Script: `make([][][]bool)`, RunOutput: [][][]bool{}},
		{Script: `make([][][]int32)`, RunOutput: [][][]int32{}},
		{Script: `make([][][]int64)`, RunOutput: [][][]int64{}},
		{Script: `make([][][]float32)`, RunOutput: [][][]float32{}},
		{Script: `make([][][]float64)`, RunOutput: [][][]float64{}},
		{Script: `make([][][]string)`, RunOutput: [][][]string{}},

		// slice type errors
		{Script: `[]nilT{"a"}`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `[]int64{1++}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `[]int64{"a"}`, RunError: fmt.Errorf("cannot use type string as type int64 as slice value")},

		// slice type
		{Script: `[]interface{nil}`, RunOutput: []interface{}{nil}},
		{Script: `[]bool{true, false}`, RunOutput: []bool{true, false}},
		{Script: `[]int32{1}`, RunOutput: []int32{int32(1)}},
		{Script: `[]int64{2}`, RunOutput: []int64{int64(2)}},
		{Script: `[]float32{3.5}`, RunOutput: []float32{float32(3.5)}},
		{Script: `[]float64{4.5}`, RunOutput: []float64{float64(4.5)}},
		{Script: `[]string{"a"}`, RunOutput: []string{"a"}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestSliceOfSlices(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a = [1, 2]; a[:]`, ParseError: fmt.Errorf("syntax error")},
		{Script: `(1++)[0:0]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = [1, 2]; a[1++:0]`, RunError: fmt.Errorf("invalid operation"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1, 2]; a[0:1++]`, RunError: fmt.Errorf("invalid operation"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1, 2]; a[:0]++`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1, 2]; a[:0]--`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},

		{Script: `a = [1, 2, 3]; a[nil:2]`, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:nil]`, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[-1:0]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:2]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:3]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:4]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[1:0]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:1]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:2]`, RunOutput: []interface{}{int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:3]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:4]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[2:1]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[2:2]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[2:3]`, RunOutput: []interface{}{int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[2:4]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[3:2]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[3:3]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[3:4]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[4:4]`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[-1:]`, RunError: fmt.Errorf("index out of range"), RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[2:]`, RunOutput: []interface{}{int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[3:]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[4:]`, RunError: fmt.Errorf("index out of range"), RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[:-1]`, RunError: fmt.Errorf("index out of range"), RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:2]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:3]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:4]`, RunError: fmt.Errorf("index out of range"), RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `b[1:2] = 4`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `a = [1, 2, 3]; a[1++:2] = 4`, RunError: fmt.Errorf("invalid operation"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:2++] = 4`, RunError: fmt.Errorf("invalid operation"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[nil:2] = 4`, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:nil] = 4`, RunError: fmt.Errorf("index must be a number"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[0:0] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:1] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:4] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:0] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:4] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[-1:] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[0:] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[1:] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[4:] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1, 2, 3]; a[:-1] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:0] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:1] = 4`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},
		{Script: `a = [1, 2, 3]; a[:4] = 4`, RunError: fmt.Errorf("index out of range"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		// slice assigned elem
		{Script: `a[0][0:0] = 4`, Input: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}}}, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}}}},

		{Script: `a = [{"b": "b"}, {"c": "c"}, {"d": "d"}]; a[0:2].a`, RunError: fmt.Errorf("type slice does not support member operation"), Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": "b"}, map[interface{}]interface{}{"c": "c"}, map[interface{}]interface{}{"d": "d"}}}},

		{Script: `a = [{"b": "b"}, {"c": "c"}, {"d": "d"}]; a[0:2]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": "b"}, map[interface{}]interface{}{"c": "c"}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": "b"}, map[interface{}]interface{}{"c": "c"}, map[interface{}]interface{}{"d": "d"}}}},
		{Script: `a = [{"b": "b"}, {"c": "c"}, {"d": "d"}]; a[0:2][0].b`, RunOutput: "b", Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": "b"}, map[interface{}]interface{}{"c": "c"}, map[interface{}]interface{}{"d": "d"}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:3]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][1:3]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][2:3]`, RunOutput: []interface{}{int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][3:3]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][1:]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][2:]`, RunOutput: []interface{}{int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][3:]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:2]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][0:3]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][:1]`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][:2]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[0][:3]`, RunOutput: []interface{}{int64(1), int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:3]`, RunOutput: []interface{}{int64(4), int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][1:3]`, RunOutput: []interface{}{int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][2:3]`, RunOutput: []interface{}{int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][3:3]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:]`, RunOutput: []interface{}{int64(4), int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][1:]`, RunOutput: []interface{}{int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][2:]`, RunOutput: []interface{}{int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][3:]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:1]`, RunOutput: []interface{}{int64(4)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:2]`, RunOutput: []interface{}{int64(4), int64(5)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][0:3]`, RunOutput: []interface{}{int64(4), int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][:0]`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][:1]`, RunOutput: []interface{}{int64(4)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][:2]`, RunOutput: []interface{}{int64(4), int64(5)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},
		{Script: `a = [[1, 2, 3], [4, 5, 6]]; a[1][:3]`, RunOutput: []interface{}{int64(4), int64(5), int64(6)}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2), int64(3)}, []interface{}{int64(4), int64(5), int64(6)}}}},

		{Script: `a = [["123"], ["456"]]; a[0][0][0:3]`, RunOutput: "123", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][1:3]`, RunOutput: "23", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][2:3]`, RunOutput: "3", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][3:3]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[0][0][0:]`, RunOutput: "123", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][1:]`, RunOutput: "23", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][2:]`, RunOutput: "3", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][3:]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[0][0][0:0]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][0:1]`, RunOutput: "1", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][0:2]`, RunOutput: "12", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][0:3]`, RunOutput: "123", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[0][0][:0]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][:1]`, RunOutput: "1", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][:2]`, RunOutput: "12", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[0][0][:3]`, RunOutput: "123", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[1][0][0:3]`, RunOutput: "456", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][1:3]`, RunOutput: "56", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][2:3]`, RunOutput: "6", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][3:3]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[1][0][0:]`, RunOutput: "456", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][1:]`, RunOutput: "56", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][2:]`, RunOutput: "6", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][3:]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[1][0][0:0]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][0:1]`, RunOutput: "4", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][0:2]`, RunOutput: "45", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][0:3]`, RunOutput: "456", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		{Script: `a = [["123"], ["456"]]; a[1][0][:0]`, RunOutput: "", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][:1]`, RunOutput: "4", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][:2]`, RunOutput: "45", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},
		{Script: `a = [["123"], ["456"]]; a[1][0][:3]`, RunOutput: "456", Output: map[string]interface{}{"a": []interface{}{[]interface{}{"123"}, []interface{}{"456"}}}},

		// cap errors
		{Script: `a = "a"; a[0:1:1]`, RunError: fmt.Errorf("type string does not support cap")},
		{Script: `a = [1,2,3,4]; a[1:3:-1]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:3:0]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:3:2]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:3:5]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[:2:-1]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[:2:0]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[:2:1]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[:2:5]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:3:"a"]`, RunError: fmt.Errorf("cap must be a number")},
		{Script: `a = [1,2,3,4]; a[1:3:1++]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = "a"; b = a[0:1:1]`, RunError: fmt.Errorf("type string does not support cap")},
		{Script: `a = [1,2,3,4]; b = a[1:3:-1]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; b = a[1:3:0]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; b = a[1:3:2]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; b = a[1:3:5]`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; b = a[1:3:"a"]`, RunError: fmt.Errorf("cap must be a number")},
		{Script: `a = [1,2,3,4]; b = a[1:3:1++]`, RunError: fmt.Errorf("invalid operation")},

		// cap assigned errors
		{Script: `a = [1,2,3,4]; a[1:1:-1] = 3`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:1:0] = 3`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:1:5] = 3`, RunError: fmt.Errorf("cap out of range")},
		{Script: `a = [1,2,3,4]; a[1:1:"a"] = 3`, RunError: fmt.Errorf("cap must be a number")},
		{Script: `a = [1,2,3,4]; a[1:3:1++] = 3`, RunError: fmt.Errorf("invalid operation")},

		// cap
		{Script: `a = [1,2,3,4]; a[1:3:3]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}}},
		{Script: `a = [1,2,3,4]; a[1:3:4]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}}},
		{Script: `a = [1,2,3,4]; a[:2:3]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}}},
		{Script: `a = [1,2,3,4]; a[:2:4]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}}},
		{Script: `a = [1,2,3,4]; b = a[1:3:3]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}, "b": []interface{}{int64(2), int64(3)}}},
		{Script: `a = [1,2,3,4]; b = a[1:3:4]`, RunOutput: []interface{}{int64(2), int64(3)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}, "b": []interface{}{int64(2), int64(3)}}},
		{Script: `a = [1,2,3,4]; b = a[:2:3]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1,2,3,4]; b = a[:2:4]`, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}, "b": []interface{}{int64(1), int64(2)}}},

		// cap assigned
		{Script: `a = [1,2,3,4]; a[1:1:1] = 3`, RunError: fmt.Errorf("slice cannot be assigned"), Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3), int64(4)}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestSliceAppendSlices(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a += 1`, Input: map[string]interface{}{"a": []bool{true}}, RunError: fmt.Errorf("invalid type conversion")},
		{Script: `a += 1.1`, Input: map[string]interface{}{"a": []bool{true}}, RunError: fmt.Errorf("invalid type conversion")},
		{Script: `a += "a"`, Input: map[string]interface{}{"a": []bool{true}}, RunError: fmt.Errorf("invalid type conversion")},
		{Script: `a += b`, Input: map[string]interface{}{"a": "a", "b": []bool{true}}, RunError: fmt.Errorf("invalid type conversion")},

		{Script: `a += b`, Input: map[string]interface{}{"a": []bool{true}, "b": []string{"b"}}, RunError: fmt.Errorf("invalid type conversion")},
		{Script: `b += a`, Input: map[string]interface{}{"a": []bool{true}, "b": []string{"b"}}, RunError: fmt.Errorf("invalid type conversion")},

		{Script: `a += nil`, Input: map[string]interface{}{"a": []bool{true}}, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []bool{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []bool{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []bool{true}}, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false}}},

		{Script: `b = [true]; b += a`, Input: map[string]interface{}{"a": []bool{}}, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": []bool{}, "b": []interface{}{true}}},
		{Script: `b = [true]; b += a`, Input: map[string]interface{}{"a": []bool{true}}, RunOutput: []interface{}{true, true}, Output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, true}}},
		{Script: `b = [true]; b += a`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: []interface{}{true, true, false}, Output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, true, false}}},

		{Script: `b = [true, false]; b += a`, Input: map[string]interface{}{"a": []bool{}}, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": []bool{}, "b": []interface{}{true, false}}},
		{Script: `b = [true, false]; b += a`, Input: map[string]interface{}{"a": []bool{true}}, RunOutput: []interface{}{true, false, true}, Output: map[string]interface{}{"a": []bool{true}, "b": []interface{}{true, false, true}}},
		{Script: `b = [true, false]; b += a`, Input: map[string]interface{}{"a": []bool{true, false}}, RunOutput: []interface{}{true, false, true, false}, Output: map[string]interface{}{"a": []bool{true, false}, "b": []interface{}{true, false, true, false}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int32(1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int32(1), int32(2)}}},

		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{int64(1), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int32(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{int64(1), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int32(1), int32(2)}}},

		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{int64(1), int64(2), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), int64(2), int32(1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{int64(1), int64(2), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), int64(2), int32(1), int32(2)}}},

		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{float64(1.1), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int32(1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{float64(1.1), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int32(1), int32(2)}}},

		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), float64(2.2), int32(1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int32(1), int32(2)}}},

		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{int64(1), float64(2.2)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{int64(1), float64(2.2), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{int64(1), float64(2.2), int32(1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{int64(1), float64(2.2), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{int64(1), float64(2.2), int32(1), int32(2)}}},

		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{}}, RunOutput: []interface{}{float64(1.1), int64(2)}, Output: map[string]interface{}{"a": []int32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{1}}, RunOutput: []interface{}{float64(1.1), int64(2), int32(1)}, Output: map[string]interface{}{"a": []int32{1}, "b": []interface{}{float64(1.1), int64(2), int32(1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int32{1, 2}}, RunOutput: []interface{}{float64(1.1), int64(2), int32(1), int32(2)}, Output: map[string]interface{}{"a": []int32{1, 2}, "b": []interface{}{float64(1.1), int64(2), int32(1), int32(2)}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2)}}},

		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{int64(1), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{int64(1), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(1), int64(2)}}},

		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{int64(1), int64(2), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), int64(2), int64(1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{int64(1), int64(2), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), int64(2), int64(1), int64(2)}}},

		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{float64(1.1), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{float64(1.1), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(1), int64(2)}}},

		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), float64(2.2), int64(1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), int64(1), int64(2)}}},

		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{int64(1), float64(2.2)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{int64(1), float64(2.2), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{int64(1), float64(2.2), int64(1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{int64(1), float64(2.2), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{int64(1), float64(2.2), int64(1), int64(2)}}},

		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{}}, RunOutput: []interface{}{float64(1.1), int64(2)}, Output: map[string]interface{}{"a": []int64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{1}}, RunOutput: []interface{}{float64(1.1), int64(2), int64(1)}, Output: map[string]interface{}{"a": []int64{1}, "b": []interface{}{float64(1.1), int64(2), int64(1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []int64{1, 2}}, RunOutput: []interface{}{float64(1.1), int64(2), int64(1), int64(2)}, Output: map[string]interface{}{"a": []int64{1, 2}, "b": []interface{}{float64(1.1), int64(2), int64(1), int64(2)}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float32(1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float32(1), float32(2)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float32(1.1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float32(1.1), float32(2.2)}}},

		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{int64(1), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float32(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{int64(1), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float32(1), float32(2)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{int64(1), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float32(1.1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{int64(1), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float32(1.1), float32(2.2)}}},

		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{int64(1), int64(2), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), int64(2), float32(1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{int64(1), int64(2), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), int64(2), float32(1), float32(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{int64(1), int64(2), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), int64(2), float32(1.1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float32(1.1), float32(2.2)}}},

		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{float64(1.1), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float32(1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{float64(1.1), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float32(1), float32(2)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{float64(1.1), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float32(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float32(1.1), float32(2.2)}}},

		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1), float32(2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float32(1.1), float32(2.2)}}},

		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{int64(1), float64(2.2)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{int64(1), float64(2.2), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{int64(1), float64(2.2), float32(1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{int64(1), float64(2.2), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{int64(1), float64(2.2), float32(1), float32(2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{int64(1), float64(2.2), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{int64(1), float64(2.2), float32(1.1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float32(1.1), float32(2.2)}}},

		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{}}, RunOutput: []interface{}{float64(1.1), int64(2)}, Output: map[string]interface{}{"a": []float32{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1}}, RunOutput: []interface{}{float64(1.1), int64(2), float32(1)}, Output: map[string]interface{}{"a": []float32{1}, "b": []interface{}{float64(1.1), int64(2), float32(1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1, 2}}, RunOutput: []interface{}{float64(1.1), int64(2), float32(1), float32(2)}, Output: map[string]interface{}{"a": []float32{1, 2}, "b": []interface{}{float64(1.1), int64(2), float32(1), float32(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1}}, RunOutput: []interface{}{float64(1.1), int64(2), float32(1.1)}, Output: map[string]interface{}{"a": []float32{1.1}, "b": []interface{}{float64(1.1), int64(2), float32(1.1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float32{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}, Output: map[string]interface{}{"a": []float32{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float32(1.1), float32(2.2)}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1), float64(2)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1)}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2)}}},

		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{int64(1), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{int64(1), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(1), float64(2)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{int64(1), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(1.1)}}},
		{Script: `b = [1]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{int64(1), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(1.1), float64(2.2)}}},

		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{int64(1), int64(2)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{int64(1), int64(2), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), int64(2), float64(1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{int64(1), int64(2), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), int64(2), float64(1), float64(2)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{int64(1), int64(2), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), int64(2), float64(1.1)}}},
		{Script: `b = [1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), int64(2), float64(1.1), float64(2.2)}}},

		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{float64(1.1), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{float64(1.1), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(1), float64(2)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{float64(1.1), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(1.1)}}},
		{Script: `b = [1.1]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(1.1), float64(2.2)}}},

		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), float64(2.2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1), float64(2)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1)}}},
		{Script: `b = [1.1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), float64(2.2), float64(1.1), float64(2.2)}}},

		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{int64(1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{int64(1), float64(2.2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{int64(1), float64(2.2), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{int64(1), float64(2.2), float64(1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{int64(1), float64(2.2), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{int64(1), float64(2.2), float64(1), float64(2)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{int64(1), float64(2.2), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{int64(1), float64(2.2), float64(1.1)}}},
		{Script: `b = [1, 2.2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{int64(1), float64(2.2), float64(1.1), float64(2.2)}}},

		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{}}, RunOutput: []interface{}{float64(1.1), int64(2)}, Output: map[string]interface{}{"a": []float64{}, "b": []interface{}{float64(1.1), int64(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1}}, RunOutput: []interface{}{float64(1.1), int64(2), float64(1)}, Output: map[string]interface{}{"a": []float64{1}, "b": []interface{}{float64(1.1), int64(2), float64(1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1, 2}}, RunOutput: []interface{}{float64(1.1), int64(2), float64(1), float64(2)}, Output: map[string]interface{}{"a": []float64{1, 2}, "b": []interface{}{float64(1.1), int64(2), float64(1), float64(2)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1}}, RunOutput: []interface{}{float64(1.1), int64(2), float64(1.1)}, Output: map[string]interface{}{"a": []float64{1.1}, "b": []interface{}{float64(1.1), int64(2), float64(1.1)}}},
		{Script: `b = [1.1, 2]; b += a`, Input: map[string]interface{}{"a": []float64{1.1, 2.2}}, RunOutput: []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}, Output: map[string]interface{}{"a": []float64{1.1, 2.2}, "b": []interface{}{float64(1.1), int64(2), float64(1.1), float64(2.2)}}},

		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": []string{}, "b": []interface{}{}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []string{"a"}}, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a"}}},
		{Script: `b = []; b += a`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: []interface{}{"a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b"}}},

		{Script: `b = ["a"]; b += a`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []string{}, "b": []interface{}{"a"}}},
		{Script: `b = ["a"]; b += a`, Input: map[string]interface{}{"a": []string{"a"}}, RunOutput: []interface{}{"a", "a"}, Output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "a"}}},
		{Script: `b = ["a"]; b += a`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: []interface{}{"a", "a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "a", "b"}}},

		{Script: `b = ["a", "b"]; b += a`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []interface{}{"a", "b"}, Output: map[string]interface{}{"a": []string{}, "b": []interface{}{"a", "b"}}},
		{Script: `b = ["a", "b"]; b += a`, Input: map[string]interface{}{"a": []string{"a"}}, RunOutput: []interface{}{"a", "b", "a"}, Output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{"a", "b", "a"}}},
		{Script: `b = ["a", "b"]; b += a`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: []interface{}{"a", "b", "a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{"a", "b", "a", "b"}}},

		{Script: `b = [1, "a"]; b += a`, Input: map[string]interface{}{"a": []string{}}, RunOutput: []interface{}{int64(1), "a"}, Output: map[string]interface{}{"a": []string{}, "b": []interface{}{int64(1), "a"}}},
		{Script: `b = [1, "a"]; b += a`, Input: map[string]interface{}{"a": []string{"a"}}, RunOutput: []interface{}{int64(1), "a", "a"}, Output: map[string]interface{}{"a": []string{"a"}, "b": []interface{}{int64(1), "a", "a"}}},
		{Script: `b = [1, "a"]; b += a`, Input: map[string]interface{}{"a": []string{"a", "b"}}, RunOutput: []interface{}{int64(1), "a", "a", "b"}, Output: map[string]interface{}{"a": []string{"a", "b"}, "b": []interface{}{int64(1), "a", "a", "b"}}},

		{Script: `a = []; a += [nil, nil]; a += [nil, nil]`, RunOutput: []interface{}{interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil)}, Output: map[string]interface{}{"a": []interface{}{interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil)}}},
		{Script: `a = []; a += [true, false]; a += [false, true]`, RunOutput: []interface{}{interface{}(true), interface{}(false), interface{}(false), interface{}(true)}, Output: map[string]interface{}{"a": []interface{}{interface{}(true), interface{}(false), interface{}(false), interface{}(true)}}},
		{Script: `a = []; a += [1, 2]; a += [3, 4]`, RunOutput: []interface{}{interface{}(int64(1)), interface{}(int64(2)), interface{}(int64(3)), interface{}(int64(4))}, Output: map[string]interface{}{"a": []interface{}{interface{}(int64(1)), interface{}(int64(2)), interface{}(int64(3)), interface{}(int64(4))}}},
		{Script: `a = []; a += [1.1, 2.2]; a += [3.3, 4.4]`, RunOutput: []interface{}{interface{}(float64(1.1)), interface{}(float64(2.2)), interface{}(float64(3.3)), interface{}(float64(4.4))}, Output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1)), interface{}(float64(2.2)), interface{}(float64(3.3)), interface{}(float64(4.4))}}},
		{Script: `a = []; a += ["a", "b"]; a += ["c", "d"]`, RunOutput: []interface{}{interface{}("a"), interface{}("b"), interface{}("c"), interface{}("d")}, Output: map[string]interface{}{"a": []interface{}{interface{}("a"), interface{}("b"), interface{}("c"), interface{}("d")}}},

		{Script: `a = []; a += [[nil, nil]]; a += [[nil, nil]]`, RunOutput: []interface{}{[]interface{}{nil, nil}, []interface{}{nil, nil}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{nil, nil}, []interface{}{nil, nil}}}},
		{Script: `a = []; a += [[true, false]]; a += [[false, true]]`, RunOutput: []interface{}{[]interface{}{true, false}, []interface{}{false, true}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{true, false}, []interface{}{false, true}}}},
		{Script: `a = []; a += [[1, 2]]; a += [[3, 4]]`, RunOutput: []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{Script: `a = []; a += [[1.1, 2.2]]; a += [[3.3, 4.4]]`, RunOutput: []interface{}{[]interface{}{float64(1.1), float64(2.2)}, []interface{}{float64(3.3), float64(4.4)}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{float64(1.1), float64(2.2)}, []interface{}{float64(3.3), float64(4.4)}}}},
		{Script: `a = []; a += [["a", "b"]]; a += [["c", "d"]]`, RunOutput: []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}}}},

		{Script: `a = make([]bool); a += 1`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{}}},
		{Script: `a = make([]bool); a += true; a += 1`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{true}}},
		{Script: `a = make([]bool); a += [1]`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{}}},
		{Script: `a = make([]bool); a += [true, 1]`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{}}},

		{Script: `a = make([]bool); a += true; a += false`, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a = make([]bool); a += [true]; a += [false]`, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},
		{Script: `a = make([]bool); a += [true, false]`, RunOutput: []bool{true, false}, Output: map[string]interface{}{"a": []bool{true, false}}},

		{Script: `a = make([]bool); a += true; a += b`, Input: map[string]interface{}{"b": int64(1)}, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{true}, "b": int64(1)}},
		{Script: `a = make([]bool); a += [true]; a += [b]`, Input: map[string]interface{}{"b": int64(1)}, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{true}, "b": int64(1)}},
		{Script: `a = make([]bool); a += [true, b]`, Input: map[string]interface{}{"b": int64(1)}, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{}, "b": int64(1)}},

		{Script: `a = make([]bool); a += b; a += b`, Input: map[string]interface{}{"b": true}, RunOutput: []bool{true, true}, Output: map[string]interface{}{"a": []bool{true, true}, "b": true}},
		{Script: `a = make([]bool); a += [b]; a += [b]`, Input: map[string]interface{}{"b": true}, RunOutput: []bool{true, true}, Output: map[string]interface{}{"a": []bool{true, true}, "b": true}},
		{Script: `a = make([]bool); a += [b, b]`, Input: map[string]interface{}{"b": true}, RunOutput: []bool{true, true}, Output: map[string]interface{}{"a": []bool{true, true}, "b": true}},

		{Script: `a = make([]bool); a += [true, false]; b = make([]int64); b += [1, 2]; a += b`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{true, false}, "b": []int64{int64(1), int64(2)}}},

		{Script: `a = []; b = []; b += true; b += false; a += b`, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": []interface{}{true, false}, "b": []interface{}{true, false}}},
		{Script: `a = []; b = make([]bool); b += true; b += false; a += b`, RunOutput: []interface{}{true, false}, Output: map[string]interface{}{"a": []interface{}{true, false}, "b": []bool{true, false}}},
		{Script: `a = []; b = []; b += [true]; b += [false]; a += [b]`, RunOutput: []interface{}{[]interface{}{true, false}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{true, false}}, "b": []interface{}{true, false}}},
		{Script: `a = []; b = make([]bool); b += [true]; b += [false]; a += [b]`, RunOutput: []interface{}{[]bool{true, false}}, Output: map[string]interface{}{"a": []interface{}{[]bool{true, false}}, "b": []bool{true, false}}},

		{Script: `a = [true, false]; b = [true, false]; a += b`, RunOutput: []interface{}{true, false, true, false}, Output: map[string]interface{}{"a": []interface{}{true, false, true, false}, "b": []interface{}{true, false}}},
		{Script: `a = make([]bool); a += [true, false]; b = make([]bool); b += [true, false]; a += b`, RunOutput: []bool{true, false, true, false}, Output: map[string]interface{}{"a": []bool{true, false, true, false}, "b": []bool{true, false}}},
		{Script: `a = make([]bool); a += [true, false]; b = [true, false]; a += b`, RunOutput: []bool{true, false, true, false}, Output: map[string]interface{}{"a": []bool{true, false, true, false}, "b": []interface{}{true, false}}},
		{Script: `a = [true, false]; b = make([]bool); b += [true, false]; a += b`, RunOutput: []interface{}{true, false, true, false}, Output: map[string]interface{}{"a": []interface{}{true, false, true, false}, "b": []bool{true, false}}},

		{Script: `a = make([][]bool); b = make([][]bool);  a += b`, RunOutput: [][]bool{}, Output: map[string]interface{}{"a": [][]bool{}, "b": [][]bool{}}},
		{Script: `a = make([][]bool); b = make([][]bool); b += [[]]; a += b`, RunOutput: [][]bool{{}}, Output: map[string]interface{}{"a": [][]bool{{}}, "b": [][]bool{{}}}},
		{Script: `a = make([][]bool); a += [[]]; b = make([][]bool); a += b`, RunOutput: [][]bool{{}}, Output: map[string]interface{}{"a": [][]bool{{}}, "b": [][]bool{}}},
		{Script: `a = make([][]bool); a += [[]]; b = make([][]bool); b += [[]]; a += b`, RunOutput: [][]bool{{}, {}}, Output: map[string]interface{}{"a": [][]bool{{}, {}}, "b": [][]bool{{}}}},

		{Script: `a = make([]bool); a += []; b = make([]bool); b += []; a += b`, RunOutput: []bool{}, Output: map[string]interface{}{"a": []bool{}, "b": []bool{}}},
		{Script: `a = make([]bool); a += [true]; b = make([]bool); b += []; a += b`, RunOutput: []bool{true}, Output: map[string]interface{}{"a": []bool{true}, "b": []bool{}}},
		{Script: `a = make([]bool); a += []; b = make([]bool); b += [true]; a += b`, RunOutput: []bool{true}, Output: map[string]interface{}{"a": []bool{true}, "b": []bool{true}}},
		{Script: `a = make([]bool); a += [true]; b = make([]bool); b += [true]; a += b`, RunOutput: []bool{true, true}, Output: map[string]interface{}{"a": []bool{true, true}, "b": []bool{true}}},

		{Script: `a = make([][]bool); a += [true, false];`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": [][]bool{}}},
		{Script: `a = make([][]bool); a += [[true, false]]; b = make([]bool); b += [true, false]; a += b`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": [][]bool{{true, false}}, "b": []bool{true, false}}},
		{Script: `a = make([][]bool); a += [[true, false]]; b = [[1.1]]; a += b`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": [][]bool{{true, false}}, "b": []interface{}{[]interface{}{1.1}}}},
		{Script: `a = make([]bool); a += [true, false]; b = make([][]bool); b += [[true, false]]; a += b`, RunError: fmt.Errorf("invalid type conversion"), Output: map[string]interface{}{"a": []bool{true, false}, "b": [][]bool{{true, false}}}},

		{Script: `a = make([][]interface); a += [[1, 2]]`, RunOutput: [][]interface{}{{int64(1), int64(2)}}, Output: map[string]interface{}{"a": [][]interface{}{{int64(1), int64(2)}}}},
		{Script: `a = make([][]interface); b = [1, 2]; a += [b]`, RunOutput: [][]interface{}{{int64(1), int64(2)}}, Output: map[string]interface{}{"a": [][]interface{}{{int64(1), int64(2)}}, "b": []interface{}{int64(1), int64(2)}}},
		{Script: `a = make([][]interface); a += [[1, 2],[3, 4]]`, RunOutput: [][]interface{}{{int64(1), int64(2)}, {int64(3), int64(4)}}, Output: map[string]interface{}{"a": [][]interface{}{{int64(1), int64(2)}, {int64(3), int64(4)}}}},
		{Script: `a = make([][]interface); b = [1, 2]; a += [b]; b = [3, 4]; a += [b]`, RunOutput: [][]interface{}{{int64(1), int64(2)}, {int64(3), int64(4)}}, Output: map[string]interface{}{"a": [][]interface{}{{int64(1), int64(2)}, {int64(3), int64(4)}}, "b": []interface{}{int64(3), int64(4)}}},

		{Script: `a = [["a", "b"], ["c", "d"]]; b = [["w", "x"], ["y", "z"]]; a += b`, RunOutput: []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}, []interface{}{"w", "x"}, []interface{}{"y", "z"}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}, []interface{}{"w", "x"}, []interface{}{"y", "z"}}, "b": []interface{}{[]interface{}{"w", "x"}, []interface{}{"y", "z"}}}},
		{Script: `a = make([][]string); a += [["a", "b"], ["c", "d"]]; b = make([][]string); b += [["w", "x"], ["y", "z"]]; a += b`, RunOutput: [][]string{{"a", "b"}, {"c", "d"}, {"w", "x"}, {"y", "z"}}, Output: map[string]interface{}{"a": [][]string{{"a", "b"}, {"c", "d"}, {"w", "x"}, {"y", "z"}}, "b": [][]string{{"w", "x"}, {"y", "z"}}}},
		{Script: `a = make([][]string); a += [["a", "b"], ["c", "d"]]; b = [["w", "x"], ["y", "z"]]; a += b`, RunOutput: [][]string{{"a", "b"}, {"c", "d"}, {"w", "x"}, {"y", "z"}}, Output: map[string]interface{}{"a": [][]string{{"a", "b"}, {"c", "d"}, {"w", "x"}, {"y", "z"}}, "b": []interface{}{[]interface{}{"w", "x"}, []interface{}{"y", "z"}}}},
		{Script: `a = [["a", "b"], ["c", "d"]]; b = make([][]string); b += [["w", "x"], ["y", "z"]]; a += b`, RunOutput: []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}, []string{"w", "x"}, []string{"y", "z"}}, Output: map[string]interface{}{"a": []interface{}{[]interface{}{"a", "b"}, []interface{}{"c", "d"}, []string{"w", "x"}, []string{"y", "z"}}, "b": [][]string{{"w", "x"}, {"y", "z"}}}},

		{Script: `a = make([][]int64); a += [[1, 2], [3, 4]]; b = make([][]int32); b += [[5, 6], [7, 8]]; a += b`, RunOutput: [][]int64{{int64(1), int64(2)}, {int64(3), int64(4)}, {int64(5), int64(6)}, {int64(7), int64(8)}}, Output: map[string]interface{}{"a": [][]int64{{int64(1), int64(2)}, {int64(3), int64(4)}, {int64(5), int64(6)}, {int64(7), int64(8)}}, "b": [][]int32{{int32(5), int32(6)}, {int32(7), int32(8)}}}},
		{Script: `a = make([][]int32); a += [[1, 2], [3, 4]]; b = make([][]int64); b += [[5, 6], [7, 8]]; a += b`, RunOutput: [][]int32{{int32(1), int32(2)}, {int32(3), int32(4)}, {int32(5), int32(6)}, {int32(7), int32(8)}}, Output: map[string]interface{}{"a": [][]int32{{int32(1), int32(2)}, {int32(3), int32(4)}, {int32(5), int32(6)}, {int32(7), int32(8)}}, "b": [][]int64{{int64(5), int64(6)}, {int64(7), int64(8)}}}},
		{Script: `a = make([][]int64); a += [[1, 2], [3, 4]]; b = make([][]float64); b += [[5, 6], [7, 8]]; a += b`, RunOutput: [][]int64{{int64(1), int64(2)}, {int64(3), int64(4)}, {int64(5), int64(6)}, {int64(7), int64(8)}}, Output: map[string]interface{}{"a": [][]int64{{int64(1), int64(2)}, {int64(3), int64(4)}, {int64(5), int64(6)}, {int64(7), int64(8)}}, "b": [][]float64{{float64(5), float64(6)}, {float64(7), float64(8)}}}},
		{Script: `a = make([][]float64); a += [[1, 2], [3, 4]]; b = make([][]int64); b += [[5, 6], [7, 8]]; a += b`, RunOutput: [][]float64{{float64(1), float64(2)}, {float64(3), float64(4)}, {float64(5), float64(6)}, {float64(7), float64(8)}}, Output: map[string]interface{}{"a": [][]float64{{float64(1), float64(2)}, {float64(3), float64(4)}, {float64(5), float64(6)}, {float64(7), float64(8)}}, "b": [][]int64{{int64(5), int64(6)}, {int64(7), int64(8)}}}},

		{Script: `a = make([][][]interface); a += [[[1, 2]]]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}}}}},
		{Script: `a = make([][][]interface); b = [[1, 2]]; a += [b]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}}}, "b": []interface{}{[]interface{}{int64(1), int64(2)}}}},
		{Script: `a = make([][][]interface); b = [1, 2]; a += [[b]]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}}}, "b": []interface{}{int64(1), int64(2)}}},

		{Script: `a = make([][][]interface); a += [[[1, 2],[3, 4]]]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}}},
		{Script: `a = make([][][]interface); b = [[1, 2],[3, 4]]; a += [b]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, "b": []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{Script: `a = make([][][]interface); b = [1, 2]; c = [b]; b = [3, 4]; c += [b]; a += [c]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, "b": []interface{}{int64(3), int64(4)}, "c": []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
		{Script: `a = make([][][]interface); b = [1, 2]; c = []; c += [b]; b = [3, 4]; c += [b]; a += [c]`, RunOutput: [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, Output: map[string]interface{}{"a": [][][]interface{}{{{int64(1), int64(2)}, {int64(3), int64(4)}}}, "b": []interface{}{int64(3), int64(4)}, "c": []interface{}{[]interface{}{int64(1), int64(2)}, []interface{}{int64(3), int64(4)}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		// TOFIX: this should have a ParseError
		{Script: `{ , }`, RunOutput: map[interface{}]interface{}{}},
		{Script: `{"a": }`, ParseError: fmt.Errorf("syntax error")},
		{Script: `{ :"a"}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `{ , "b":"b"}`, ParseError: fmt.Errorf("syntax error: unexpected ','"), RunOutput: map[interface{}]interface{}{"b": "b"}},
		{Script: `{"a":"a", , "c":"c"}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `{"b": 1++}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `{1++: 1}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = {}; a.b.c`, RunError: fmt.Errorf("type interface does not support member operation")},
		{Script: `a = {}; a.b.c = 1`, RunError: fmt.Errorf("type interface does not support member operation")},
		{Script: `a = {}; a[1++]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = {}; a[1++] = 1`, RunError: fmt.Errorf("invalid operation")},
		{Script: `b[1]`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `b[1] = 1`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `z.y.x = 1`, RunError: fmt.Errorf("undefined symbol 'z'")},

		{Script: `{}`, RunOutput: map[interface{}]interface{}{}},
		{Script: `{"b": nil}`, RunOutput: map[interface{}]interface{}{"b": nil}},
		{Script: `{"b": true}`, RunOutput: map[interface{}]interface{}{"b": true}},
		{Script: `{"b": 1}`, RunOutput: map[interface{}]interface{}{"b": int64(1)}},
		{Script: `{"b": 1.1}`, RunOutput: map[interface{}]interface{}{"b": float64(1.1)}},
		{Script: `{"b": "b"}`, RunOutput: map[interface{}]interface{}{"b": "b"}},

		{Script: `{1: nil}`, RunOutput: map[interface{}]interface{}{int64(1): nil}},
		{Script: `{1: true}`, RunOutput: map[interface{}]interface{}{int64(1): true}},
		{Script: `{1: 2}`, RunOutput: map[interface{}]interface{}{int64(1): int64(2)}},
		{Script: `{1: 2.2}`, RunOutput: map[interface{}]interface{}{int64(1): float64(2.2)}},
		{Script: `{1: "b"}`, RunOutput: map[interface{}]interface{}{int64(1): "b"}},

		{Script: `a = {}`, RunOutput: map[interface{}]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = map{}`, RunOutput: map[interface{}]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = map {}`, RunOutput: map[interface{}]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": nil}`, RunOutput: map[interface{}]interface{}{"b": nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": nil}}},
		{Script: `a = {"b": true}`, RunOutput: map[interface{}]interface{}{"b": true}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": true}}},
		{Script: `a = {"b": 1}`, RunOutput: map[interface{}]interface{}{"b": int64(1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1)}}},
		{Script: `a = {"b": 1.1}`, RunOutput: map[interface{}]interface{}{"b": float64(1.1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": float64(1.1)}}},
		{Script: `a = {"b": "b"}`, RunOutput: map[interface{}]interface{}{"b": "b"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		{Script: `a = {"b": {}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{}}}},
		{Script: `a = {"b": {"c": nil}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": nil}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": nil}}}},
		{Script: `a = {"b": {"c": true}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": true}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": true}}}},
		{Script: `a = {"b": {"c": 1}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": int64(1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": int64(1)}}}},
		{Script: `a = {"b": {"c": 1.1}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": float64(1.1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": float64(1.1)}}}},
		{Script: `a = {"b": {"c": "c"}}`, RunOutput: map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": "c"}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": "c"}}}},

		{Script: `a = {"b": {}}; a.b`, RunOutput: map[interface{}]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{}}}},
		{Script: `a = {"b": {"c": nil}}; a.b`, RunOutput: map[interface{}]interface{}{"c": nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": nil}}}},
		{Script: `a = {"b": {"c": true}}; a.b`, RunOutput: map[interface{}]interface{}{"c": true}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": true}}}},
		{Script: `a = {"b": {"c": 1}}; a.b`, RunOutput: map[interface{}]interface{}{"c": int64(1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": int64(1)}}}},
		{Script: `a = {"b": {"c": 1.1}}; a.b`, RunOutput: map[interface{}]interface{}{"c": float64(1.1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": float64(1.1)}}}},
		{Script: `a = {"b": {"c": "c"}}; a.b`, RunOutput: map[interface{}]interface{}{"c": "c"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": map[interface{}]interface{}{"c": "c"}}}},

		{Script: `a = {"b": []}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{}}}},
		{Script: `a = {"b": [nil]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{nil}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{nil}}}},
		{Script: `a = {"b": [true]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{true}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{true}}}},
		{Script: `a = {"b": [1]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{int64(1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{int64(1)}}}},
		{Script: `a = {"b": [1.1]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{float64(1.1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{float64(1.1)}}}},
		{Script: `a = {"b": ["c"]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{"c"}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"c"}}}},

		{Script: `a = {}; a.b`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": nil}; a.b`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": nil}}},
		{Script: `a = {"b": true}; a.b`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": true}}},
		{Script: `a = {"b": 1}; a.b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1)}}},
		{Script: `a = {"b": 1.1}; a.b`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": float64(1.1)}}},
		{Script: `a = {"b": "b"}; a.b`, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		{Script: `a = {}; a["b"]`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": nil}; a["b"]`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": nil}}},
		{Script: `a = {"b": true}; a["b"]`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": true}}},
		{Script: `a = {"b": 1}; a["b"]`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1)}}},
		{Script: `a = {"b": 1.1}; a["b"]`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": float64(1.1)}}},
		{Script: `a = {"b": "b"}; a["b"]`, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: map[string]interface{}{}, Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, RunOutput: map[string]interface{}{"b": nil}, Output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, RunOutput: map[string]interface{}{"b": true}, Output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, RunOutput: map[string]interface{}{"b": int32(1)}, Output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, RunOutput: map[string]interface{}{"b": int64(1)}, Output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, RunOutput: map[string]interface{}{"b": float32(1.1)}, Output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, RunOutput: map[string]interface{}{"b": float64(1.1)}, Output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{Script: `a`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, RunOutput: map[string]interface{}{"b": "b"}, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]bool{"a": false, "b": true}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, RunOutput: int32(2), Output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, RunOutput: int64(2), Output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, RunOutput: float32(2.2), Output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}},

		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, RunOutput: reflect.Value{}, Output: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": reflect.Value{}}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": nil}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": true}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": int32(1)}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": float32(1.1)}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{Script: `a[0:1]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, RunError: fmt.Errorf("type map does not support slice operation"), Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": nil}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int32(1)}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": int64(1)}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float32(1.1)}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": float64(1.1)}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "b"}},
		{Script: `a[c]`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": "c"}},

		{Script: `a.b = nil`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": nil}}},
		{Script: `a.b = true`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]interface{}{"b": true}}},
		{Script: `a.b = 1`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]interface{}{"b": int64(1)}}},
		{Script: `a.b = 1.1`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[string]interface{}{"b": float64(1.1)}}},
		{Script: `a.b = "b"`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a.b = true`, Input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		{Script: `a.b = 3`, Input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, RunOutput: int64(3), Output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{Script: `a.b = 3`, Input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, RunOutput: int64(3), Output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		{Script: `a.b = 3.3`, Input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{Script: `a.b = 3.3`, Input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{Script: `a.b = "c"`, Input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, RunOutput: "c", Output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{Script: `a["b"] = true`, Input: map[string]interface{}{"a": map[string]bool{"a": true, "b": false}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]bool{"a": true, "b": true}}},
		{Script: `a["b"] = 3`, Input: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(2)}}, RunOutput: int64(3), Output: map[string]interface{}{"a": map[string]int32{"a": int32(1), "b": int32(3)}}},
		{Script: `a["b"] = 3`, Input: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(2)}}, RunOutput: int64(3), Output: map[string]interface{}{"a": map[string]int64{"a": int64(1), "b": int64(3)}}},
		{Script: `a["b"] = 3.3`, Input: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(2.2)}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": map[string]float32{"a": float32(1.1), "b": float32(3.3)}}},
		{Script: `a["b"] = 3.3`, Input: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(2.2)}}, RunOutput: float64(3.3), Output: map[string]interface{}{"a": map[string]float64{"a": float64(1.1), "b": float64(3.3)}}},
		{Script: `a["b"] = "c"`, Input: map[string]interface{}{"a": map[string]string{"a": "a", "b": "b"}}, RunOutput: "c", Output: map[string]interface{}{"a": map[string]string{"a": "a", "b": "c"}}},

		{Script: `a[c] = "x"`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}, RunError: fmt.Errorf("index type bool cannot be used for map index type string"), RunOutput: nil, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": true}},
		{Script: `a[c] = "x"`, Input: map[string]interface{}{"a": map[bool]interface{}{true: "b"}, "c": true}, RunOutput: "x", Output: map[string]interface{}{"a": map[bool]interface{}{true: "x"}, "c": true}},

		// note if passed an uninitialized map there does not seem to be a way to update that map
		{Script: `a`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: testMapEmpty, Output: map[string]interface{}{"a": testMapEmpty}},
		{Script: `a.b`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: nil, Output: map[string]interface{}{"a": testMapEmpty}},
		{Script: `a.b = 1`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: nil, Output: map[string]interface{}{"a": testMapEmpty}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: nil, Output: map[string]interface{}{"a": testMapEmpty}},
		{Script: `a["b"] = 1`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": testMapEmpty}, RunOutput: nil, Output: map[string]interface{}{"a": testMapEmpty}},

		{Script: `a`, Input: map[string]interface{}{"a": testMap}, RunOutput: testMap, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.a`, Input: map[string]interface{}{"a": testMap}, RunOutput: nil, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.a = true`, Input: map[string]interface{}{"a": testMap}, RunOutput: true, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.a`, Input: map[string]interface{}{"a": testMap}, RunOutput: true, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.a = nil`, Input: map[string]interface{}{"a": testMap}, RunOutput: nil, Output: map[string]interface{}{"a": testMap}},

		{Script: `a`, Input: map[string]interface{}{"a": testMap}, RunOutput: testMap, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.b`, Input: map[string]interface{}{"a": testMap}, RunOutput: true, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.b = false`, Input: map[string]interface{}{"a": testMap}, RunOutput: false, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.b`, Input: map[string]interface{}{"a": testMap}, RunOutput: false, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.b = true`, Input: map[string]interface{}{"a": testMap}, RunOutput: true, Output: map[string]interface{}{"a": testMap}},

		{Script: `a`, Input: map[string]interface{}{"a": testMap}, RunOutput: testMap, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.c`, Input: map[string]interface{}{"a": testMap}, RunOutput: int64(1), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.c = 2`, Input: map[string]interface{}{"a": testMap}, RunOutput: int64(2), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.c`, Input: map[string]interface{}{"a": testMap}, RunOutput: int64(2), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.c = 1`, Input: map[string]interface{}{"a": testMap}, RunOutput: int64(1), Output: map[string]interface{}{"a": testMap}},

		{Script: `a`, Input: map[string]interface{}{"a": testMap}, RunOutput: testMap, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.d`, Input: map[string]interface{}{"a": testMap}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.d = 2.2`, Input: map[string]interface{}{"a": testMap}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.d`, Input: map[string]interface{}{"a": testMap}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": testMap}},
		{Script: `a.d = 1.1`, Input: map[string]interface{}{"a": testMap}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": testMap}},

		{Script: `a`, Input: map[string]interface{}{"a": testMap}, RunOutput: testMap, Output: map[string]interface{}{"a": testMap}},
		{Script: `a.e`, Input: map[string]interface{}{"a": testMap}, RunOutput: "e", Output: map[string]interface{}{"a": testMap}},
		{Script: `a.e = "x"`, Input: map[string]interface{}{"a": testMap}, RunOutput: "x", Output: map[string]interface{}{"a": testMap}},
		{Script: `a.e`, Input: map[string]interface{}{"a": testMap}, RunOutput: "x", Output: map[string]interface{}{"a": testMap}},
		{Script: `a.e = "e"`, Input: map[string]interface{}{"a": testMap}, RunOutput: "e", Output: map[string]interface{}{"a": testMap}},

		{Script: `a = {"b": 1, "c": nil}`, RunOutput: map[interface{}]interface{}{"b": int64(1), "c": nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.c`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.d`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},

		{Script: `a = {"b": 1, "c": nil}; a == nil`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a != nil`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.b == nil`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.b != nil`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.c == nil`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.c != nil`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.d == nil`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.d != nil`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},

		{Script: `a = {"b": 1, "c": nil}; a == 1`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a != 1`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.b == 1`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.b != 1`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.c == 1`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.c != 1`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.d == 1`, RunOutput: false, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},
		{Script: `a = {"b": 1, "c": nil}; a.d != 1`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": int64(1), "c": nil}}},

		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]bool{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]bool{"b": true}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}},
		{Script: `a["b"]`, Input: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}},

		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]bool{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]bool{"b": true}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}},
		{Script: `a.b`, Input: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}},

		{Script: `a = {}; a[true] = true`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{true: true}}},
		{Script: `a = {}; a[1] = 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int64(1): int64(1)}}},
		{Script: `a = {}; a[1.1] = 1.1`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float64(1.1): float64(1.1)}}},

		{Script: `a = {}; a[true] = true; a[true]`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{true: true}}},
		{Script: `a = {}; a[1] = 1; a[1]`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int64(1): int64(1)}}},
		{Script: `a = {}; a[1.1] = 1.1; a[1.1]`, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float64(1.1): float64(1.1)}}},

		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{nil: nil}}},
		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": int32(1)}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int32(1): int32(1)}}},
		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int64(1): int64(1)}}},
		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": float32(1.1)}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float32(1.1): float32(1.1)}}},
		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": float64(1.1)}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float64(1.1): float64(1.1)}}},
		{Script: `a = {}; a[b] = b`, Input: map[string]interface{}{"b": "b"}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": nil}, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{nil: nil}}},
		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": int32(1)}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int32(1): int32(1)}}},
		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{int64(1): int64(1)}}},
		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": float32(1.1)}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float32(1.1): float32(1.1)}}},
		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": float64(1.1)}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{float64(1.1): float64(1.1)}}},
		{Script: `a = {}; a[b] = b; a[b]`, Input: map[string]interface{}{"b": "b"}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		// test equal nil when not found
		{Script: `a = {"b":"b"}; if a.c == nil { return 1 }`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},
		{Script: `a = {"b":"b"}; if a["c"] == nil { return 1 }`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},

		// test map create with spacing and comma
		{Script: `{"b": "b",}`, RunOutput: map[interface{}]interface{}{"b": "b"}},
		{Script: `{"b": "b"
}`, RunOutput: map[interface{}]interface{}{"b": "b"}},
		{Script: `{"b": "b",
}`, RunOutput: map[interface{}]interface{}{"b": "b"}},
		{Script: `{"b": "b", "c": "c"}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b", "c": "c",}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b", "c": "c"
}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b", "c": "c",
}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b",
"c": "c"}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b",
"c": "c",}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b",
"c": "c"
}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
		{Script: `{"b": "b",
"c": "c",
}`, RunOutput: map[interface{}]interface{}{"b": "b", "c": "c"}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestExistenceOfKeyInMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a = {"b":"b"}; v, ok = a[1++]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = {"b":"b"}; b.c, ok = a["b"]`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `a = {"b":"b"}; v, b.c = a["b"]`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `a = {"b":"b"}; v, ok = a["a"]`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}, "v": nil, "ok": false}},
		{Script: `a = {"b":"b"}; v, ok = a["b"]`, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}, "v": "b", "ok": true}},
		{Script: `a = {"b":"b", "c":"c"}; v, ok = a["a"]`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b", "c": "c"}, "v": nil, "ok": false}},
		{Script: `a = {"b":"b", "c":"c"}; v, ok = a["b"]`, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b", "c": "c"}, "v": "b", "ok": true}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestDeleteMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `delete(1++, "b")`, RunError: fmt.Errorf("invalid operation")},
		{Script: `delete({}, 1++)`, RunError: fmt.Errorf("invalid operation")},
		{Script: `delete(nil, "b")`, RunError: fmt.Errorf("first argument to delete cannot be type interface")},
		{Script: `delete(1, "b")`, RunError: fmt.Errorf("first argument to delete cannot be type int64")},

		{Script: `delete(a, "")`, Input: map[string]interface{}{"a": testMapEmpty}, Output: map[string]interface{}{"a": testMapEmpty}},
		{Script: `delete(a, "")`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},
		{Script: `delete(a, "a")`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},
		{Script: `delete(a, "b")`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `delete(a, "a")`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b", "c": "c"}}, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b", "c": "c"}}},
		{Script: `delete(a, "b")`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b", "c": "c"}}, Output: map[string]interface{}{"a": map[string]interface{}{"c": "c"}}},

		{Script: `delete(a, 0)`, Input: map[string]interface{}{"a": map[int64]interface{}{1: 1}}, Output: map[string]interface{}{"a": map[int64]interface{}{1: 1}}},
		{Script: `delete(a, 1)`, Input: map[string]interface{}{"a": map[int64]interface{}{1: 1}}, Output: map[string]interface{}{"a": map[int64]interface{}{}}},
		{Script: `delete(a, 0)`, Input: map[string]interface{}{"a": map[int64]interface{}{1: 1, 2: 2}}, Output: map[string]interface{}{"a": map[int64]interface{}{1: 1, 2: 2}}},
		{Script: `delete(a, 1)`, Input: map[string]interface{}{"a": map[int64]interface{}{1: 1, 2: 2}}, Output: map[string]interface{}{"a": map[int64]interface{}{2: 2}}},

		{Script: `delete({}, "")`},
		{Script: `delete({}, 1)`},
		{Script: `delete({}, "a")`},
		{Script: `delete({"b":"b"}, "")`},
		{Script: `delete({"b":"b"}, 1)`},
		{Script: `delete({"b":"b"}, "a")`},
		{Script: `delete({"b":"b"}, "b")`},

		{Script: `a = {"b": "b"}; delete(a, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b"}}},
		{Script: `a = {"b": "b"}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": "b", "c":"c"}; delete(a, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": "b", "c": "c"}}},
		{Script: `a = {"b": "b", "c":"c"}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": "c"}}},

		{Script: `a = {"b": ["b"]}; delete(a, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}}}},
		{Script: `a = {"b": ["b"]}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": ["b"], "c": ["c"]}; delete(a, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}, "c": []interface{}{"c"}}}},
		{Script: `a = {"b": ["b"], "c": ["c"]}; delete(a, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": []interface{}{"c"}}}},

		{Script: `a = {"b": ["b"]}; b = &a; delete(*b, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}}}},
		{Script: `a = {"b": ["b"]}; b = &a; delete(*b, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"b": ["b"], "c": ["c"]}; b = &a; delete(*b, "a")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}, "c": []interface{}{"c"}}}},
		{Script: `a = {"b": ["b"], "c": ["c"]}; b = &a; delete(*b, "b")`, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": []interface{}{"c"}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `map[[]string]string {"a":"a"}`, RunError: fmt.Errorf("reflect.MapOf: invalid key type []string")},
	}
	runTests(t, tests, nil, &Options{Debug: false})

	tests = []Test{
		{Script: `make(mapStringBool)`, Types: map[string]interface{}{"mapStringBool": map[string]bool{}}, RunOutput: map[string]bool{}},
		{Script: `make(mapStringInt32)`, Types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, RunOutput: map[string]int32{}},
		{Script: `make(mapStringInt64)`, Types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, RunOutput: map[string]int64{}},
		{Script: `make(mapStringFloat32)`, Types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, RunOutput: map[string]float32{}},
		{Script: `make(mapStringFloat64)`, Types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, RunOutput: map[string]float64{}},
		{Script: `make(mapStringString)`, Types: map[string]interface{}{"mapStringString": map[string]string{}}, RunOutput: map[string]string{}},

		{Script: `a = make(mapStringBool)`, Types: map[string]interface{}{"mapStringBool": map[string]bool{}}, RunOutput: map[string]bool{}, Output: map[string]interface{}{"a": map[string]bool{}}},
		{Script: `a = make(mapStringInt32)`, Types: map[string]interface{}{"mapStringInt32": map[string]int32{}}, RunOutput: map[string]int32{}, Output: map[string]interface{}{"a": map[string]int32{}}},
		{Script: `a = make(mapStringInt64)`, Types: map[string]interface{}{"mapStringInt64": map[string]int64{}}, RunOutput: map[string]int64{}, Output: map[string]interface{}{"a": map[string]int64{}}},
		{Script: `a = make(mapStringFloat32)`, Types: map[string]interface{}{"mapStringFloat32": map[string]float32{}}, RunOutput: map[string]float32{}, Output: map[string]interface{}{"a": map[string]float32{}}},
		{Script: `a = make(mapStringFloat64)`, Types: map[string]interface{}{"mapStringFloat64": map[string]float64{}}, RunOutput: map[string]float64{}, Output: map[string]interface{}{"a": map[string]float64{}}},
		{Script: `a = make(mapStringString)`, Types: map[string]interface{}{"mapStringString": map[string]string{}}, RunOutput: map[string]string{}, Output: map[string]interface{}{"a": map[string]string{}}},

		{Script: `a = make(mapStringBool); a["b"] = true`, Types: map[string]interface{}{"mapStringBool": map[string]bool{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]bool{"b": true}}},
		{Script: `a = make(mapStringInt32); a["b"] = 1`, Types: map[string]interface{}{"mapStringInt32": map[string]int32{"b": int32(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]int32{"b": int32(1)}}},
		{Script: `a = make(mapStringInt64); a["b"] = 1`, Types: map[string]interface{}{"mapStringInt64": map[string]int64{"b": int64(1)}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[string]int64{"b": int64(1)}}},
		{Script: `a = make(mapStringFloat32); a["b"] = 1.1`, Types: map[string]interface{}{"mapStringFloat32": map[string]float32{"b": float32(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[string]float32{"b": float32(1.1)}}},
		{Script: `a = make(mapStringFloat64); a["b"] = 1.1`, Types: map[string]interface{}{"mapStringFloat64": map[string]float64{"b": float64(1.1)}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[string]float64{"b": float64(1.1)}}},
		{Script: `a = make(mapStringString); a["b"] = "b"`, Types: map[string]interface{}{"mapStringString": map[string]string{"b": "b"}}, RunOutput: "b", Output: map[string]interface{}{"a": map[string]string{"b": "b"}}},

		{Script: `a = make(mapStringBool); a.b = true`, Types: map[string]interface{}{"mapStringBool": map[string]bool{"b": true}}, RunOutput: true, Output: map[string]interface{}{"a": map[string]bool{"b": true}}},

		{Script: `a = make(mapInterfaceBool); a["b"] = true; a["b"]`, Types: map[string]interface{}{"mapInterfaceBool": map[interface{}]bool{}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]bool{"b": true}}},
		{Script: `a = make(mapInterfaceInt32); a["b"] = 1; a["b"]`, Types: map[string]interface{}{"mapInterfaceInt32": map[interface{}]int32{}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}},
		{Script: `a = make(mapInterfaceInt64); a["b"] = 1; a["b"]`, Types: map[string]interface{}{"mapInterfaceInt64": map[interface{}]int64{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}},
		{Script: `a = make(mapInterfaceFloat32); a["b"] = 1.1; a["b"]`, Types: map[string]interface{}{"mapInterfaceFloat32": map[interface{}]float32{}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}},
		{Script: `a = make(mapInterfaceFloat64); a["b"] = 1.1; a["b"]`, Types: map[string]interface{}{"mapInterfaceFloat64": map[interface{}]float64{}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}},
		{Script: `a = make(mapInterfaceString); a["b"] = "b"; a["b"]`, Types: map[string]interface{}{"mapInterfaceString": map[interface{}]string{}}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}},

		{Script: `a = make(mapInterfaceBool); a.b = true; a.b`, Types: map[string]interface{}{"mapInterfaceBool": map[interface{}]bool{}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]bool{"b": true}}},
		{Script: `a = make(mapInterfaceInt32); a.b = 1; a.b`, Types: map[string]interface{}{"mapInterfaceInt32": map[interface{}]int32{}}, RunOutput: int32(1), Output: map[string]interface{}{"a": map[interface{}]int32{"b": int32(1)}}},
		{Script: `a = make(mapInterfaceInt64); a.b = 1; a.b`, Types: map[string]interface{}{"mapInterfaceInt64": map[interface{}]int64{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]int64{"b": int64(1)}}},
		{Script: `a = make(mapInterfaceFloat32); a.b = 1.1; a.b`, Types: map[string]interface{}{"mapInterfaceFloat32": map[interface{}]float32{}}, RunOutput: float32(1.1), Output: map[string]interface{}{"a": map[interface{}]float32{"b": float32(1.1)}}},
		{Script: `a = make(mapInterfaceFloat64); a.b = 1.1; a.b`, Types: map[string]interface{}{"mapInterfaceFloat64": map[interface{}]float64{}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]float64{"b": float64(1.1)}}},
		{Script: `a = make(mapInterfaceString); a.b = "b"; a.b`, Types: map[string]interface{}{"mapInterfaceString": map[interface{}]string{}}, RunOutput: "b", Output: map[string]interface{}{"a": map[interface{}]string{"b": "b"}}},

		// map type errors
		{Script: `map[int64]string {"a":"a"}`, RunError: fmt.Errorf("cannot use type string as type int64 as map key")},
		{Script: `map[string]int64 {"a":"a"}`, RunError: fmt.Errorf("cannot use type string as type int64 as map value")},
		{Script: `map[nilT]interface {"a":"a"}`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `map[interface]nilT {"a":"a"}`, Types: map[string]interface{}{"nilT": nil}, RunError: fmt.Errorf("cannot make type nil")},
		{Script: `map[int64]int64 {1++:1}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `map[int64]int64 {1:1++}`, RunError: fmt.Errorf("invalid operation")},

		// map type
		{Script: `map[string]interface {"a":nil}`, RunOutput: map[string]interface{}{"a": nil}},
		{Script: `map[string]bool {"a":true}`, RunOutput: map[string]bool{"a": true}},
		{Script: `map[string]int32 {"a":1}`, RunOutput: map[string]int32{"a": 1}},
		{Script: `map[string]int64 {"a":2}`, RunOutput: map[string]int64{"a": 2}},
		{Script: `map[string]float32 {"a":3.5}`, RunOutput: map[string]float32{"a": 3.5}},
		{Script: `map[string]float64 {"a":4.5}`, RunOutput: map[string]float64{"a": 4.5}},
		{Script: `map[string]string {"a":"a"}`, RunOutput: map[string]string{"a": "a"}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestSlicesAndMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a = [{"b": nil}]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": interface{}(nil)}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(nil)}}}},
		{Script: `a = [{"b": true}]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": interface{}(true)}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(true)}}}},
		{Script: `a = [{"b": 1}]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": interface{}(int64(1))}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(int64(1))}}}},
		{Script: `a = [{"b": 1.1}]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": interface{}(float64(1.1))}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(float64(1.1))}}}},
		{Script: `a = [{"b": "b"}]`, RunOutput: []interface{}{map[interface{}]interface{}{"b": interface{}("b")}}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}("b")}}}},

		{Script: `a = [{"b": nil}]; a[0]`, RunOutput: map[interface{}]interface{}{"b": interface{}(nil)}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(nil)}}}},
		{Script: `a = [{"b": true}]; a[0]`, RunOutput: map[interface{}]interface{}{"b": interface{}(true)}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(true)}}}},
		{Script: `a = [{"b": 1}]; a[0]`, RunOutput: map[interface{}]interface{}{"b": interface{}(int64(1))}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(int64(1))}}}},
		{Script: `a = [{"b": 1.1}]; a[0]`, RunOutput: map[interface{}]interface{}{"b": interface{}(float64(1.1))}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}(float64(1.1))}}}},
		{Script: `a = [{"b": "b"}]; a[0]`, RunOutput: map[interface{}]interface{}{"b": interface{}("b")}, Output: map[string]interface{}{"a": []interface{}{map[interface{}]interface{}{"b": interface{}("b")}}}},

		{Script: `a = {"b": []}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{}}}},
		{Script: `a = {"b": [nil]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{nil}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{nil}}}},
		{Script: `a = {"b": [true]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{true}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{true}}}},
		{Script: `a = {"b": [1]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{int64(1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{int64(1)}}}},
		{Script: `a = {"b": [1.1]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{float64(1.1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{float64(1.1)}}}},
		{Script: `a = {"b": ["b"]}`, RunOutput: map[interface{}]interface{}{"b": []interface{}{"b"}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}}}},

		{Script: `a = {"b": []}; a.b`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{}}}},
		{Script: `a = {"b": [nil]}; a.b`, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{nil}}}},
		{Script: `a = {"b": [true]}; a.b`, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{true}}}},
		{Script: `a = {"b": [1]}; a.b`, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{int64(1)}}}},
		{Script: `a = {"b": [1.1]}; a.b`, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{float64(1.1)}}}},
		{Script: `a = {"b": ["b"]}; a.b`, RunOutput: []interface{}{"b"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{"b"}}}},

		{Script: `a.b = []`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {}}}},
		{Script: `a.b = [nil]`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{interface{}(nil)}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {interface{}(nil)}}}},
		{Script: `a.b = [true]`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {true}}}},
		{Script: `a.b = [1]`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {int64(1)}}}},
		{Script: `a.b = [1.1]`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {float64(1.1)}}}},
		{Script: `a.b = ["b"]`, Input: map[string]interface{}{"a": map[string][]interface{}{}}, RunOutput: []interface{}{"b"}, Output: map[string]interface{}{"a": map[string][]interface{}{"b": {"b"}}}},

		{Script: `b[0] = [nil]; a.b = b`, Input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, RunOutput: [][]interface{}{{interface{}(nil)}}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {{interface{}(nil)}}}, "b": [][]interface{}{{interface{}(nil)}}}},
		{Script: `b[0] = [true]; a.b = b`, Input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, RunOutput: [][]interface{}{{true}}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {{true}}}, "b": [][]interface{}{{true}}}},
		{Script: `b[0] = [1]; a.b = b`, Input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, RunOutput: [][]interface{}{{int64(1)}}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {{int64(1)}}}, "b": [][]interface{}{{int64(1)}}}},
		{Script: `b[0] = [1.1]; a.b = b`, Input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, RunOutput: [][]interface{}{{float64(1.1)}}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {{float64(1.1)}}}, "b": [][]interface{}{{float64(1.1)}}}},
		{Script: `b[0] = ["b"]; a.b = b`, Input: map[string]interface{}{"a": map[string][][]interface{}{}, "b": [][]interface{}{}}, RunOutput: [][]interface{}{{"b"}}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {{"b"}}}, "b": [][]interface{}{{"b"}}}},

		{Script: `a`, Input: map[string]interface{}{"a": map[string][][]interface{}{}}, RunOutput: map[string][][]interface{}{}, Output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{Script: `a.b = 1`, Input: map[string]interface{}{"a": map[string][][]interface{}{}}, RunError: fmt.Errorf("type int64 cannot be assigned to type [][]interface {} for map"), RunOutput: nil, Output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{Script: `a["b"] = 1`, Input: map[string]interface{}{"a": map[string][][]interface{}{}}, RunError: fmt.Errorf("type int64 cannot be assigned to type [][]interface {} for map"), RunOutput: nil, Output: map[string]interface{}{"a": map[string][][]interface{}{}}},

		{Script: `a = {}; a.b = []; a.b += 1; a.b[0] = {}; a.b[0].c = []; a.b[0].c += 2; a.b[0].c[0]`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{map[interface{}]interface{}{"c": []interface{}{int64(2)}}}}}},

		{Script: `a = {}; a.b = b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunError: fmt.Errorf("index out of range"), RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{}}, "b": [][]interface{}{}}},

		{Script: `a = {}; a.b = b; a.b[0] = []`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [nil]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{nil}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [nil]; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{{nil}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{nil}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [true]; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{{true}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{true}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1]; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{{int64(1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{int64(1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1.1]; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{{float64(1.1)}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{float64(1.1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = ["c"]; a.b`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{{"c"}}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{"c"}}}, "b": [][]interface{}{}}},

		{Script: `a = {}; a.b = b; a.b[0] = [nil]; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{nil}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{nil}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [true]; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{true}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{true}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1]; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{int64(1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{int64(1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1.1]; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{float64(1.1)}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{float64(1.1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = ["c"]; a.b[0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: []interface{}{"c"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{"c"}}}, "b": [][]interface{}{}}},

		{Script: `a = {}; a.b = b; a.b[0] = [nil]; a.b[0][0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{nil}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [true]; a.b[0][0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{true}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1]; a.b[0][0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{int64(1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1.1]; a.b[0][0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{float64(1.1)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = ["c"]; a.b[0][0]`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: "c", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{"c"}}}, "b": [][]interface{}{}}},

		{Script: `a = {}; a.b = b; a.b[0] = [nil]; a.b[0][1] = nil`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{nil, nil}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [true]; a.b[0][1] = true`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{true, true}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1]; a.b[0][1] = 2`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: int64(2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{int64(1), int64(2)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = [1.1]; a.b[0][1] = 2.2`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: float64(2.2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{float64(1.1), float64(2.2)}}}, "b": [][]interface{}{}}},
		{Script: `a = {}; a.b = b; a.b[0] = ["c"]; a.b[0][1] = "d"`, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: "d", Output: map[string]interface{}{"a": map[interface{}]interface{}{"b": [][]interface{}{{"c", "d"}}}, "b": [][]interface{}{}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeSlicesAndMaps(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `make([]aMap)`, Types: map[string]interface{}{"aMap": map[string]interface{}{}}, RunOutput: []map[string]interface{}{}},
		{Script: `make([][]aMap)`, Types: map[string]interface{}{"aMap": map[string]interface{}{}}, RunOutput: [][]map[string]interface{}{}},

		{Script: `make(mapSlice2x)`, Types: map[string]interface{}{"mapSlice2x": map[string][][]interface{}{}}, RunOutput: map[string][][]interface{}{}},
		{Script: `a = make(mapSlice2x)`, Types: map[string]interface{}{"mapSlice2x": map[string][][]interface{}{}}, RunOutput: map[string][][]interface{}{}, Output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{Script: `a = make(mapSlice2x); a`, Types: map[string]interface{}{"mapSlice2x": map[string][][]interface{}{}}, RunOutput: map[string][][]interface{}{}, Output: map[string]interface{}{"a": map[string][][]interface{}{}}},
		{Script: `a = make(mapSlice2x); a.b = b`, Types: map[string]interface{}{"mapSlice2x": map[string][][]interface{}{}}, Input: map[string]interface{}{"b": [][]interface{}{}}, RunOutput: [][]interface{}{}, Output: map[string]interface{}{"a": map[string][][]interface{}{"b": {}}, "b": [][]interface{}{}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeSlicesData(t *testing.T) {
	t.Parallel()

	stmts, err := parser.ParseSrc("make(slice)")
	if err != nil {
		t.Errorf("ParseSrc error - received %v - expected: %v", err, nil)
	}

	e := env.NewEnv()
	err = e.DefineType("slice", []string{})
	if err != nil {
		t.Errorf("DefineType error - received %v - expected: %v", err, nil)
	}

	value, err := Run(e, nil, stmts)
	if err != nil {
		t.Errorf("Run error - received %v - expected: %v", err, nil)
	}
	if !reflect.DeepEqual(value, []string{}) {
		t.Errorf("Run value - received %#v - expected: %#v", value, []string{})
	}

	a := value.([]string)
	if len(a) != 0 {
		t.Errorf("len value - received %#v - expected: %#v", len(a), 0)
	}
	a = append(a, "a")
	if a[0] != "a" {
		t.Errorf("Get value - received %#v - expected: %#v", a[0], "a")
	}
	if len(a) != 1 {
		t.Errorf("len value - received %#v - expected: %#v", len(a), 1)
	}

	stmts, err = parser.ParseSrc("make([]string)")
	if err != nil {
		t.Errorf("ParseSrc error - received %v - expected: %v", err, nil)
	}

	e = env.NewEnv()
	err = e.DefineType("string", "a")
	if err != nil {
		t.Errorf("DefineType error - received %v - expected: %v", err, nil)
	}

	value, err = Run(e, nil, stmts)
	if err != nil {
		t.Errorf("Run error - received %v - expected: %v", err, nil)
	}
	if !reflect.DeepEqual(value, []string{}) {
		t.Errorf("Run value - received %#v - expected: %#v", value, []string{})
	}

	b := value.([]string)
	if len(b) != 0 {
		t.Errorf("len value - received %#v - expected: %#v", len(b), 0)
	}
	b = append(b, "b")
	if b[0] != "b" {
		t.Errorf("Get value - received %#v - expected: %#v", b[0], "b")
	}
	if len(b) != 1 {
		t.Errorf("len value - received %#v - expected: %#v", len(b), 1)
	}
}

func TestMakeMapsData(t *testing.T) {
	t.Parallel()

	stmts, err := parser.ParseSrc("make(aMap)")
	if err != nil {
		t.Errorf("ParseSrc error - received %v - expected: %v", err, nil)
	}

	// test normal map
	e := env.NewEnv()
	err = e.DefineType("aMap", map[string]string{})
	if err != nil {
		t.Errorf("DefineType error - received %v - expected: %v", err, nil)
	}

	value, err := Run(e, nil, stmts)
	if err != nil {
		t.Errorf("Run error - received %v - expected: %v", err, nil)
	}
	if !reflect.DeepEqual(value, map[string]string{}) {
		t.Errorf("Run value - received %#v - expected: %#v", value, map[string]string{})
	}

	a := value.(map[string]string)
	a["a"] = "a"
	if a["a"] != "a" {
		t.Errorf("Get value - received %#v - expected: %#v", a["a"], "a")
	}

	// test url Values map
	e = env.NewEnv()
	err = e.DefineType("aMap", url.Values{})
	if err != nil {
		t.Errorf("DefineType error - received %v - expected: %v", err, nil)
	}

	value, err = Run(e, nil, stmts)
	if err != nil {
		t.Errorf("Run error - received %v - expected: %v", err, nil)
	}
	if !reflect.DeepEqual(value, url.Values{}) {
		t.Errorf("Run value - received %#v - expected: %#v", value, url.Values{})
	}

	b := value.(url.Values)
	b.Set("b", "b")
	if b.Get("b") != "b" {
		t.Errorf("Get value - received %#v - expected: %#v", b.Get("b"), "b")
	}
}

func TestStructs(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a["B"]`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunError: fmt.Errorf("type struct does not support index operation"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{Script: `a.C`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunError: fmt.Errorf("no member named 'C' for struct"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: nil, B: nil}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: nil, B: nil}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int32(1), B: int32(2)}},
			RunOutput: int32(2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int32(1), B: int32(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float32(1.1), B: float32(2.2)}},
			RunOutput: float32(2.2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float32(1.1), B: float32(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float64(1.1), B: float64(2.2)}},
			RunOutput: float64(2.2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float64(1.1), B: float64(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: "a", B: "b"}},
			RunOutput: "b",
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: "a", B: "b"}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{}},
			RunOutput: false,
			Output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{}},
			RunOutput: int32(0),
			Output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{}},
			RunOutput: int64(0),
			Output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{}},
			RunOutput: float32(0),
			Output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{}},
			RunOutput: float64(0),
			Output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A string
			B string
		}{}},
			RunOutput: "",
			Output: map[string]interface{}{"a": struct {
				A string
				B string
			}{}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{A: true, B: true}},
			RunOutput: true,
			Output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{A: true, B: true}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{A: int32(1), B: int32(2)}},
			RunOutput: int32(2),
			Output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{A: int32(1), B: int32(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(2),
			Output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{A: int64(1), B: int64(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{A: float32(1.1), B: float32(2.2)}},
			RunOutput: float32(2.2),
			Output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{A: float32(1.1), B: float32(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{A: float64(1.1), B: float64(2.2)}},
			RunOutput: float64(2.2),
			Output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{A: float64(1.1), B: float64(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A string
			B string
		}{A: "a", B: "b"}},
			RunOutput: "b",
			Output: map[string]interface{}{"a": struct {
				A string
				B string
			}{A: "a", B: "b"}}},

		{Script: `a.C = 3`, Input: map[string]interface{}{
			"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)},
		},
			RunError: fmt.Errorf("no member named 'C' for struct"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},

		{Script: `a.B = 3`, Input: map[string]interface{}{"a": &struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(3),
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)}}},

		{Script: `a.B = 3; a = *a`, Input: map[string]interface{}{"a": &struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)},
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)}}},

		// nil tests
		{Script: `a.A = nil; a.B = nil; a.C = nil;`, Input: map[string]interface{}{"a": &struct {
			A *bool
			B *int32
			C *int64
		}{A: new(bool), B: new(int32), C: new(int64)}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A *bool
				B *int32
				C *int64
			}{A: nil, B: nil, C: nil}}},
		{Script: `a.A = nil; a.B = nil; a.C = nil;`, Input: map[string]interface{}{"a": &struct {
			A *float32
			B *float64
			C *string
		}{A: new(float32), B: new(float64), C: new(string)}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A *float32
				B *float64
				C *string
			}{A: nil, B: nil, C: nil}}},
		{Script: `a.A = nil; a.B = nil; a.C = nil;`, Input: map[string]interface{}{"a": &struct {
			A interface{}
			B []interface{}
			C [][]interface{}
		}{A: "a", B: []interface{}{"b"}, C: [][]interface{}{{"c"}}}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B []interface{}
				C [][]interface{}
			}{A: nil, B: nil, C: nil}}},
		{Script: `a.A = nil; a.B = nil; a.C = nil; a.D = nil;`, Input: map[string]interface{}{"a": &struct {
			A map[string]string
			B map[string]interface{}
			C map[interface{}]string
			D map[interface{}]interface{}
		}{A: map[string]string{"a": "a"}, B: map[string]interface{}{"b": "b"}, C: map[interface{}]string{"c": "c"}, D: map[interface{}]interface{}{"d": "d"}}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A map[string]string
				B map[string]interface{}
				C map[interface{}]string
				D map[interface{}]interface{}
			}{A: nil, B: nil, C: nil, D: nil}}},
		{Script: `a.A.AA = nil;`, Input: map[string]interface{}{"a": &struct {
			A struct{ AA *int64 }
		}{A: struct{ AA *int64 }{AA: new(int64)}}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A struct{ AA *int64 }
			}{A: struct{ AA *int64 }{AA: nil}}}},
		{Script: `a.A = nil;`, Input: map[string]interface{}{"a": &struct {
			A *struct{ AA *int64 }
		}{A: &struct{ AA *int64 }{AA: new(int64)}}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": &struct {
				A *struct{ AA *int64 }
			}{A: nil}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestMakeStructs(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `a = make(struct1)`, Types: map[string]interface{}{"struct1": &testStruct1{}}, RunOutput: &testStruct1{}, Output: map[string]interface{}{"a": &testStruct1{}}},
		{Script: `a = make(struct2)`, Types: map[string]interface{}{"struct2": &testStruct2{}}, RunOutput: &testStruct2{}, Output: map[string]interface{}{"a": &testStruct2{}}},
		{Script: `make(struct1)`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: &struct {
				A interface{}
				B interface{}
			}{}},

		{Script: `a = make(struct1)`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: &struct {
				A interface{}
				B interface{}
			}{},
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{}}},

		{Script: `a = make(struct1); a.A = 3; a.B = 4`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(4),
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{A: interface{}(int64(3)), B: interface{}(int64(4))}}},

		{Script: `a = make(struct1); a = *a; a.A = 3; a.B = 4`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(4),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: interface{}(int64(3)), B: interface{}(int64(4))}}},

		{Script: `a = make(struct1); a.A = func () { return 1 }; a.A()`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(1)},
		{Script: `a = make(struct1); a.A = func () { return 1 }; a = *a; a.A()`, Types: map[string]interface{}{"struct1": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(1)},

		// make struct - new lines
		{Script: `make(struct { A int64, B float64 })`, RunOutput: struct {
			A int64
			B float64
		}{}},
		{Script: `make(struct {
A int64, B float64 })`, RunOutput: struct {
			A int64
			B float64
		}{}},
		{Script: `make(struct { A int64, 
B float64 })`, RunOutput: struct {
			A int64
			B float64
		}{}},
		{Script: `make(struct { A int64, B float64
})`, RunOutput: struct {
			A int64
			B float64
		}{}},
		{Script: `
make(struct {
	A int64,
	B float64
})`, RunOutput: struct {
			A int64
			B float64
		}{}},

		// make struct - with basic types
		{Script: `
make(struct {
	A bool,
	B int32,
	C int64,
	D float32,
	E float64,
	F string
})`, RunOutput: struct {
			A bool
			B int32
			C int64
			D float32
			E float64
			F string
		}{}},

		// make struct - with other types
		{Script: `
make(struct {
	A *int64,
	B []int64,
	C map[string]int64
})`, RunOutput: struct {
			A *int64
			B []int64
			C map[string]int64
		}{A: (*int64)(nil), B: []int64{}, C: map[string]int64{}}},

		// make struct within structs
		{Script: `
make(struct {
	A struct {
		AA int64,
		AB float64
	},
	B struct {
		BA []int64,
		BB map[string]int64
	}
})`, RunOutput: struct {
			A struct {
				AA int64
				AB float64
			}
			B struct {
				BA []int64
				BB map[string]int64
			}
		}{A: struct {
			AA int64
			AB float64
		}{AA: 0, AB: 0}, B: struct {
			BA []int64
			BB map[string]int64
		}{BA: []int64{}, BB: map[string]int64{}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}
