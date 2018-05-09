package packages

import (
	"os"
	"testing"
)

func TestBytes(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		// TOFIX: no member named 'WriteString' for struct
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b")`, RunOutput: []interface{}{1, nil}},
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b"); a.String()`, RunOutput: "b"},
	}
	RunTests(t, tests, &TestingOptions{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
