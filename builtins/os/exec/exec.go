// Package exec implements os/exec interface for anko script.
package exec

import (
	"github.com/mattn/anko/vm"
	e "os/exec"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("exec")
	m.Define("ErrNotFound", reflect.ValueOf(e.ErrNotFound))
	m.Define("LookPath", reflect.ValueOf(e.LookPath))
	m.Define("Command", reflect.ValueOf(e.Command))
	return m
}
