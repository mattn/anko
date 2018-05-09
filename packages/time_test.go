package packages

import (
	"os"
	"testing"
)

func TestTime(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `time = import("time"); a = make(time.Time); a.IsZero()`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: true},
	}
	RunTests(t, tests, &TestingOptions{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
