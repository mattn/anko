package packages

import (
	"os"
	"testing"
)

func TestBytes(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		// TOFIX
		// {script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b")`, runOutput: []interface{}{1, nil}},
		// {script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b"); a.String()`, runOutput: "b",
	}
	runTests(t, tests)
}
