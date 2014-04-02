// Package path implements path manipulation interface for anko script.
package filepath

import (
	"github.com/mattn/anko/vm"
	f "path/filepath"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("filepath")
	m.Define("Join", reflect.ValueOf(f.Join))
}
