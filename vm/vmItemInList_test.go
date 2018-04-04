package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestItemInList(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `"a" in ["a"]`, runOutput: true},
		{script: `"a" in ["b"]`, runOutput: false},
		{script: `"a" in ["c", "b", "a"]`, runOutput: true},
		{script: `"a" in ["a", "b", 1]`, runOutput: true},
		{script: `"a" in l`, input: map[string]interface{}{"l": []interface{}{"a"}}, runOutput: true},
		{script: `"a" in l`, input: map[string]interface{}{"l": []interface{}{"b"}}, runOutput: false},
		{script: `"a" in l`, input: map[string]interface{}{"l": []interface{}{"c", "b", "a"}}, runOutput: true},
		{script: `"a" in l`, input: map[string]interface{}{"l": []interface{}{"a", "b", 1}}, runOutput: true},

		{script: `1 in [1]`, runOutput: true},
		{script: `1 in [2]`, runOutput: false},
		{script: `1 in [3, 2, 1]`, runOutput: true},
		{script: `1 in ["1"]`, runOutput: true},
		{script: `1 in l`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `"1" in l`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1 in l`, input: map[string]interface{}{"l": []interface{}{2}}, runOutput: false},
		{script: `1 in l`, input: map[string]interface{}{"l": []interface{}{3, 2, 1}}, runOutput: true},
		{script: `1 in l`, input: map[string]interface{}{"l": []interface{}{"1"}}, runOutput: true},

		{script: `0.9 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `1.0 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1.1 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `1 in [0.9]`, input: map[string]interface{}{"l": []interface{}{0.9}}, runOutput: false},
		{script: `1 in [1.0]`, input: map[string]interface{}{"l": []interface{}{1.0}}, runOutput: true},
		{script: `1 in [1.1]`, input: map[string]interface{}{"l": []interface{}{1.1}}, runOutput: false},
		{script: `0.9 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `1.0 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1.1 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `1 in [0.9]`, input: map[string]interface{}{"l": []interface{}{0.9}}, runOutput: false},
		{script: `1 in [1.0]`, input: map[string]interface{}{"l": []interface{}{1.0}}, runOutput: true},
		{script: `1 in [1.1]`, input: map[string]interface{}{"l": []interface{}{1.1}}, runOutput: false},

		{script: `true in ["true"]`, runOutput: true},
		{script: `true in [true]`, runOutput: true},
		{script: `true in [true, false]`, runOutput: true},
		{script: `false in [false, true]`, runOutput: true},
		{script: `true in l`, input: map[string]interface{}{"l": []interface{}{"true"}}, runOutput: true},
		{script: `true in l`, input: map[string]interface{}{"l": []interface{}{true}}, runOutput: true},
		{script: `true in l`, input: map[string]interface{}{"l": []interface{}{true, false}}, runOutput: true},
		{script: `false in l`, input: map[string]interface{}{"l": []interface{}{false, true}}, runOutput: true},

		{script: `"a" in ["b", "a", "c"][1:]`, runOutput: true},
		{script: `"a" in ["b", "a", "c"][:1]`, runOutput: false},
		{script: `"a" in ["b", "a", "c"][1:2]`, runOutput: true},
		{script: `l = ["b", "a", "c"];"a" in l[1:]`, runOutput: true},
		{script: `l = ["b", "a", "c"];"a" in l[:1]`, runOutput: false},
		{script: `l = ["b", "a", "c"];"a" in l[1:2]`, runOutput: true},
		{script: `"a" in l[1:]`, input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, runOutput: true},
		{script: `"a" in l[:1]`, input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, runOutput: false},
		{script: `"a" in l[1:2]`, input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, runOutput: true},

		// for i in list && item in list
		{script: `list_of_list = [["a"]];for l in list_of_list { return "a" in l }`, runOutput: true},
		{script: `for l in list_of_list { return "a" in l }`, input: map[string]interface{}{"list_of_list": []interface{}{[]interface{}{"a"}}}, runOutput: true},

		// not slice or array
		// todo: support `"a" in "aaa"` ?
		{script: `"a" in "aaa"`, runError: fmt.Errorf("second argument must be slice or array; but have string")},
		{script: `1 in 12345`, runError: fmt.Errorf("type int64 does not support slice operation")},

		// a in item in list
		{script: `"a" in 5 in [1, 2, 3]`, runError: fmt.Errorf("type bool does not support slice operation")},

		// applying a in b in several part of expresstion/statement
		{script: `switch 1 in [1] {case true: return true;default: return false}`, runOutput: true},
		{script: `switch 1 in [2,3] {case true: return true;default: return false}`, runOutput: false},
		{script: `switch true {case 1 in [1]: return true;default: return false}`, runOutput: true},
		{script: `switch false {case 1 in [1]: return true;default: return false}`, runOutput: false},
		{script: `if 1 in [1] {return true} else {return false}`, runOutput: true},
		{script: `if 1 in [2,3] {return true} else {return false}`, runOutput: false},
		{script: `for i in [1,2,3] { i++ }`},
		{script: `a=1; a=a in [1]`, runOutput: true},
		{script: `a=1; a=a in [2,3]`, runOutput: false},
		{script: `1 in [1] && true`, runOutput: true},
		{script: `1 in [1] && false`, runOutput: false},
		{script: `1 in [1] || true`, runOutput: true},
		{script: `1 in [1] || false`, runOutput: true},
		{script: `1 in [2,3] && true`, runOutput: false},
		{script: `1 in [2,3] && false`, runOutput: false},
		{script: `1 in [2,3] || true`, runOutput: true},
		{script: `1 in [2,3] || false`, runOutput: false},
		{script: `1++ in [1, 2, 3]`, runError: fmt.Errorf("invalid operation")},
		{script: `3++ in [1, 2, 3]`, runError: fmt.Errorf("invalid operation")},
		{script: `1 in 1++`, runError: fmt.Errorf("invalid operation")},
		{script: `a=1;a++ in [1, 2, 3]`, runOutput: true},
		{script: `a=3;a++ in [1, 2, 3]`, runOutput: false},
		{script: `switch 1 in l {case true: return true;default: return false}`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `switch 1 in l {case true: return true;default: return false}`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `switch true {case 1 in l: return true;default: return false}`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `switch false {case 1 in l: return true;default: return false}`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `if 1 in l {return true} else {return false}`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `if 1 in l {return true} else {return false}`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `for i in l { i++ }`, input: map[string]interface{}{"l": []interface{}{1, 2, 3}}},
		{script: `a=1; a=a in l`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `a=1; a=a in l`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `1 in l && true`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1 in l && false`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `1 in l || true`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1 in l || false`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1 in l && true`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `1 in l && false`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `1 in l || true`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: true},
		{script: `1 in l || false`, input: map[string]interface{}{"l": []interface{}{2, 3}}, runOutput: false},
		{script: `1++ in l`, input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, runError: fmt.Errorf("invalid operation")},
		{script: `3++ in l`, input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, runError: fmt.Errorf("invalid operation")},
		{script: `a=1;a++ in l`, input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, runOutput: true},
		{script: `a=3;a++ in l`, input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, runOutput: false},

		// multidimensional slice
		{script: `1 in [1]`, runOutput: true},
		{script: `1 in [[1]]`, runOutput: false},
		{script: `1 in [[[1]]]`, runOutput: false},
		{script: `1 in [[1],[[1]],1]`, runOutput: true},
		{script: `[1] in [1]`, runOutput: false},
		{script: `[1] in [[1]]`, runOutput: true},
		{script: `[1] in [[[1]]]`, runOutput: false},
		{script: `[[1]] in [1]`, runOutput: false},
		{script: `[[1]] in [[1]]`, runOutput: false},
		{script: `[[1]] in [[[1]]]`, runOutput: true},
		{script: `1 in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: true},
		{script: `1 in [[1]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, runOutput: false},
		{script: `1 in [[[1]]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, runOutput: false},
		{script: `1 in [[1],[[1]],1]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{1}, []interface{}{[]interface{}{1}}, 1}}, runOutput: true},
		{script: `[1] in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `[1] in [[1]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, runOutput: true},
		{script: `[1] in [[[1]]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, runOutput: false},
		{script: `[[1]] in [1]`, input: map[string]interface{}{"l": []interface{}{1}}, runOutput: false},
		{script: `[[1]] in [[1]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, runOutput: false},
		{script: `[[1]] in [[[1]]]`, input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, runOutput: true},
	}
	runTests(t, tests)
}
