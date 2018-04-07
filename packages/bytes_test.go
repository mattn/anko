package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestBytes(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
		// TOFIX: no member named 'WriteString' for struct
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b")`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: []interface{}{1, nil}},
		// {Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("b"); a.String()`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: "b"},
		// Comment to trigger build
	}
	vm.RunTests(t, tests)
}
