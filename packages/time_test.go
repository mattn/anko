package packages

import (
	"os"
	"testing"
)

func TestTime(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `time = import("time"); a = make(time.Time); a.IsZero()`, runOutput: true},
	}
	runTests(t, tests)
}
