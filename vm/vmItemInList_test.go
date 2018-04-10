package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestItemInList(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `"a" in ["a"]`, RunOutput: true},
		{Script: `"a" in ["b"]`, RunOutput: false},
		{Script: `"a" in ["c", "b", "a"]`, RunOutput: true},
		{Script: `"a" in ["a", "b", 1]`, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"a"}}, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"b"}}, RunOutput: false},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"c", "b", "a"}}, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"a", "b", 1}}, RunOutput: true},

		{Script: `1 in [1]`, RunOutput: true},
		{Script: `1 in [2]`, RunOutput: false},
		{Script: `1 in [3, 2, 1]`, RunOutput: true},
		{Script: `1 in ["1"]`, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `"1" in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{2}}, RunOutput: false},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{3, 2, 1}}, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{"1"}}, RunOutput: true},

		{Script: `0.9 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1.0 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1.1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in [0.9]`, Input: map[string]interface{}{"l": []interface{}{0.9}}, RunOutput: false},
		{Script: `1 in [1.0]`, Input: map[string]interface{}{"l": []interface{}{1.0}}, RunOutput: true},
		{Script: `1 in [1.1]`, Input: map[string]interface{}{"l": []interface{}{1.1}}, RunOutput: false},
		{Script: `0.9 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1.0 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1.1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in [0.9]`, Input: map[string]interface{}{"l": []interface{}{0.9}}, RunOutput: false},
		{Script: `1 in [1.0]`, Input: map[string]interface{}{"l": []interface{}{1.0}}, RunOutput: true},
		{Script: `1 in [1.1]`, Input: map[string]interface{}{"l": []interface{}{1.1}}, RunOutput: false},

		{Script: `true in ["true"]`, RunOutput: true},
		{Script: `true in [true]`, RunOutput: true},
		{Script: `true in [true, false]`, RunOutput: true},
		{Script: `false in [false, true]`, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{"true"}}, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{true}}, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{true, false}}, RunOutput: true},
		{Script: `false in l`, Input: map[string]interface{}{"l": []interface{}{false, true}}, RunOutput: true},

		{Script: `"a" in ["b", "a", "c"][1:]`, RunOutput: true},
		{Script: `"a" in ["b", "a", "c"][:1]`, RunOutput: false},
		{Script: `"a" in ["b", "a", "c"][1:2]`, RunOutput: true},
		{Script: `l = ["b", "a", "c"];"a" in l[1:]`, RunOutput: true},
		{Script: `l = ["b", "a", "c"];"a" in l[:1]`, RunOutput: false},
		{Script: `l = ["b", "a", "c"];"a" in l[1:2]`, RunOutput: true},
		{Script: `"a" in l[1:]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: true},
		{Script: `"a" in l[:1]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: false},
		{Script: `"a" in l[1:2]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: true},

		// for i in list && item in list
		{Script: `list_of_list = [["a"]];for l in list_of_list { return "a" in l }`, RunOutput: true},
		{Script: `for l in list_of_list { return "a" in l }`, Input: map[string]interface{}{"list_of_list": []interface{}{[]interface{}{"a"}}}, RunOutput: true},

		// not slice or array
		// todo: support `"a" in "aaa"` ?
		{Script: `"a" in "aaa"`, RunError: fmt.Errorf("second argument must be slice or array; but have string")},
		{Script: `1 in 12345`, RunError: fmt.Errorf("type int64 does not support slice operation")},

		// a in item in list
		{Script: `"a" in 5 in [1, 2, 3]`, RunError: fmt.Errorf("type bool does not support slice operation")},

		// applying a in b in several part of expresstion/statement
		{Script: `switch 1 in [1] {case true: return true;default: return false}`, RunOutput: true},
		{Script: `switch 1 in [2,3] {case true: return true;default: return false}`, RunOutput: false},
		{Script: `switch true {case 1 in [1]: return true;default: return false}`, RunOutput: true},
		{Script: `switch false {case 1 in [1]: return true;default: return false}`, RunOutput: false},
		{Script: `if 1 in [1] {return true} else {return false}`, RunOutput: true},
		{Script: `if 1 in [2,3] {return true} else {return false}`, RunOutput: false},
		{Script: `for i in [1,2,3] { i++ }`},
		{Script: `a=1; a=a in [1]`, RunOutput: true},
		{Script: `a=1; a=a in [2,3]`, RunOutput: false},
		{Script: `1 in [1] && true`, RunOutput: true},
		{Script: `1 in [1] && false`, RunOutput: false},
		{Script: `1 in [1] || true`, RunOutput: true},
		{Script: `1 in [1] || false`, RunOutput: true},
		{Script: `1 in [2,3] && true`, RunOutput: false},
		{Script: `1 in [2,3] && false`, RunOutput: false},
		{Script: `1 in [2,3] || true`, RunOutput: true},
		{Script: `1 in [2,3] || false`, RunOutput: false},
		{Script: `1++ in [1, 2, 3]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `3++ in [1, 2, 3]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `1 in 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a=1;a++ in [1, 2, 3]`, RunOutput: true},
		{Script: `a=3;a++ in [1, 2, 3]`, RunOutput: false},
		{Script: `switch 1 in l {case true: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `switch 1 in l {case true: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `switch true {case 1 in l: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `switch false {case 1 in l: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `if 1 in l {return true} else {return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `if 1 in l {return true} else {return false}`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `for i in l { i++ }`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}},
		{Script: `a=1; a=a in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `a=1; a=a in l`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l && true`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l && false`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in l || true`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l || false`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l && true`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l && false`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l || true`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: true},
		{Script: `1 in l || false`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunError: fmt.Errorf("invalid operation")},
		{Script: `3++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunError: fmt.Errorf("invalid operation")},
		{Script: `a=1;a++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunOutput: true},
		{Script: `a=3;a++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunOutput: false},

		// multidimensional slice
		{Script: `1 in [1]`, RunOutput: true},
		{Script: `1 in [[1]]`, RunOutput: false},
		{Script: `1 in [[[1]]]`, RunOutput: false},
		{Script: `1 in [[1],[[1]],1]`, RunOutput: true},
		{Script: `[1] in [1]`, RunOutput: false},
		{Script: `[1] in [[1]]`, RunOutput: true},
		{Script: `[1] in [[[1]]]`, RunOutput: false},
		{Script: `[[1]] in [1]`, RunOutput: false},
		{Script: `[[1]] in [[1]]`, RunOutput: false},
		{Script: `[[1]] in [[[1]]]`, RunOutput: true},
		{Script: `1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: false},
		{Script: `1 in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: false},
		{Script: `1 in [[1],[[1]],1]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}, []interface{}{[]interface{}{1}}, 1}}, RunOutput: true},
		{Script: `[1] in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `[1] in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: true},
		{Script: `[1] in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: false},
		{Script: `[[1]] in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `[[1]] in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: false},
		{Script: `[[1]] in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: true},
	}
	RunTests(t, tests, nil)
}
