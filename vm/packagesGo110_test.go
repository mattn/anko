// +build go1.10

package vm

import (
	"testing"

	_ "github.com/mattn/anko/packages"
)

func TestPackagesStringsGo110(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{Script: `strings = import("strings"); a = make(strings.Builder); _, err = a.WriteString("a"); if err != nil { return err.Error() }; _, err = a.WriteString("b"); if err != nil { return err.Error() }; _, err = a.WriteString("c"); if err != nil { return err.Error() }; a.String()`, RunOutput: "abc"},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}
