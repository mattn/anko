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

		{script: `1 in [1]`, runOutput: true},
		{script: `1 in [2]`, runOutput: false},
		{script: `1 in [3, 2, 1]`, runOutput: true},
		{script: `1.0 in [1]`, runOutput: false},
		{script: `1 in ["1"]`, runOutput: false},

		{script: `true in ["true"]`, runOutput: false},
		{script: `true in [true]`, runOutput: true},
		{script: `true in [true, false]`, runOutput: true},
		{script: `false in [false, true]`, runOutput: true},

		{script: `"a" in ["b", "a", "c"][1:]`, runOutput: true},
		{script: `"a" in ["b", "a", "c"][:1]`, runOutput: false},
		{script: `"a" in ["b", "a", "c"][1:2]`, runOutput: true},
		{script: `l = ["b", "a", "c"];"a" in l[1:]`, runOutput: true},
		{script: `l = ["b", "a", "c"];"a" in l[:1]`, runOutput: false},
		{script: `l = ["b", "a", "c"];"a" in l[1:2]`, runOutput: true},

		// for i in list && item in list
		{script: `list_of_list = [["a"]];for l in list_of_list { return "a" in l }`, runOutput: true},

		// not slice or array
		// todo: support `"a" in "aaa"` ?
		{script: `"a" in "aaa"`, runError: fmt.Errorf("second argument must be slice or array; but get string")},
		{script: `1 in 12345`, runError: fmt.Errorf("type int64 does not support slice operation")},

		// a in item in list
		{script: `"a" in 5 in [1, 2, 3]`, runError: fmt.Errorf("type bool does not support slice operation")},
	}
	runTests(t, tests)
}
