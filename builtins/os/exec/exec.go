// Package exec implements os/exec interface for anko script.
package exec

import (
	"github.com/mattn/anko/vm"
	"reflect"
	e "os/exec"
)

func Import(env *vm.Env) {
	m := env.NewModule("exec")
	m.Define("ErrNotFound", reflect.ValueOf(e.ErrNotFound))
	m.Define("LookPath", reflect.ValueOf(e.LookPath))
	m.Define("Command", reflect.ValueOf(e.Command))
}
