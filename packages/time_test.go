package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestTime(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
		{Script: `time = import("time"); a = make(time.Time); a.IsZero()`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
	}
	vm.RunTests(t, tests)
}
