package packages

import (
	"os"
	"testing"
)

func TestRegexp(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^simple$\"); re.MatchString(\"simple\")", runOutput: true},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^simple$\"); re.MatchString(\"no match\")", runOutput: false},

		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^simple$", "b": "simple"}, runOutput: true, output: map[string]interface{}{"a": "^simple$", "b": "simple"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^simple$", "b": "no match"}, runOutput: false, output: map[string]interface{}{"a": "^simple$", "b": "no match"}},

		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.String()", runOutput: "^a\\.\\d+\\.b$"},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.MatchString(\"a.1.b\")", runOutput: true},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.MatchString(\"a.22.b\")", runOutput: true},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.MatchString(\"a.333.b\")", runOutput: true},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.MatchString(\"no match\")", runOutput: false},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(\"^a\\\\.\\\\d+\\\\.b$\"); re.MatchString(\"a+1+b\")", runOutput: false},

		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.String()", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}, runOutput: "^a\\.\\d+\\.b$", output: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}, runOutput: true, output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}, runOutput: true, output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}, runOutput: true, output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}, runOutput: false, output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}},
		{script: "regexp = import(\"regexp\"); re = regexp.MustCompile(a); re.MatchString(b)", input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}, runOutput: false, output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}},
	}
	runTests(t, tests)
}
