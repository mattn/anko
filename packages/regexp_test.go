package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestRegexp(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^simple$"); re.MatchString("simple")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^simple$"); re.MatchString("no match")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: false},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^simple$", "b": "simple"}, RunOutput: true, Output: map[string]interface{}{"a": "^simple$", "b": "simple"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^simple$", "b": "no match"}, RunOutput: false, Output: map[string]interface{}{"a": "^simple$", "b": "no match"}},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.String()`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: "^a\\.\\d+\\.b$"},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.1.b")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.22.b")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.333.b")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("no match")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: false},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a+1+b")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: false},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.String()`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}, RunOutput: "^a\\.\\d+\\.b$", Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}, RunOutput: false, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}, RunOutput: false, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}},
	}
	vm.RunTests(t, tests)
}
