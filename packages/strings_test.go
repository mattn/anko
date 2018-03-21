package packages

import (
	"os"
	"testing"
)

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `strings = import("strings"); a = " one two "; b = strings.TrimSpace(a)`, runOutput: "one two", output: map[string]interface{}{"a": " one two ", "b": "one two"}},
		{script: `strings = import("strings"); a = "a b c d"; b = strings.Split(a, " ")`, runOutput: []string{"a", "b", "c", "d"}, output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c", "d"}}},
		{script: `strings = import("strings"); a = "a b c d"; b = strings.SplitN(a, " ", 3)`, runOutput: []string{"a", "b", "c d"}, output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c d"}}},
	}
	runTests(t, tests)
}
