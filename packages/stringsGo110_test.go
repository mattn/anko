package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestStringsGo110(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `strings = import("strings"); a = make(strings.Builder); _, err = a.WriteString("a"); if err != nil { return err.Error() }; _, err = a.WriteString("b"); if err != nil { return err.Error() }; _, err = a.WriteString("c"); if err != nil { return err.Error() }; a.String()`, RunOutput: "abc"},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
