package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestStrings(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
		{Script: `strings = import("strings"); a = " one two "; b = strings.TrimSpace(a)`, RunOutput: "one two", Output: map[string]interface{}{"a": " one two ", "b": "one two"}},
		{Script: `strings = import("strings"); a = "a b c d"; b = strings.Split(a, " ")`, RunOutput: []string{"a", "b", "c", "d"}, Output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c", "d"}}},
		{Script: `strings = import("strings"); a = "a b c d"; b = strings.SplitN(a, " ", 3)`, RunOutput: []string{"a", "b", "c d"}, Output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c d"}}},
	}
	vm.RunTests(t, tests, &vm.TestingOptions{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
