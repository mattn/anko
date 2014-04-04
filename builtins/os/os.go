// Package os implements os related interface for anko script.
package os

import (
	"github.com/mattn/anko/builtins/os/exec"
	"github.com/mattn/anko/vm"
	o "os"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("os")
	m.Define("Stdin", reflect.ValueOf(o.Stdin))
	m.Define("Stdout", reflect.ValueOf(o.Stdout))
	m.Define("Stderr", reflect.ValueOf(o.Stderr))
	m.Define("Getenv", reflect.ValueOf(o.Getenv))
	m.Define("Setenv", reflect.ValueOf(o.Setenv))
	m.Define("Exit", reflect.ValueOf(o.Exit))

	exec.Import(m)
}
