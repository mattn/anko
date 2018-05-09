package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestBytes(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		// TOFIX: no member named 'WriteString' for struct
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b")`, RunOutput: []interface{}{1, nil}},
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b"); a.String()`, RunOutput: "b"},
	}
	testlib.RunTests(t, tests, &testlib.TestingOptions{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
