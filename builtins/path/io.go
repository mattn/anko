// Package path implements path manipulation interface for anko script.
package path

import (
	"github.com/mattn/anko/builtins/path/filepath"
	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) {
	m := env.NewModule("path")

	filepath.Import(m)
}
