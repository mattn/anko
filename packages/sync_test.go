package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestSync(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `sync = import("sync"); once = make(sync.Once); a = []; func add() { a += "a" }; once.Do(add); once.Do(add); a`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []interface{}{"a"}}},
		{Script: `sync = import("sync"); waitGroup = make(sync.WaitGroup); waitGroup.Add(2);  func done() { waitGroup.Done() }; go done(); go done(); waitGroup.Wait(); "a"`, RunOutput: "a"},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
