package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestBytes(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("a"); if err != nil { return err }; n`, RunOutput: 1},
		{Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("a"); if err != nil { return err }; a.String()`, RunOutput: "a"},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
