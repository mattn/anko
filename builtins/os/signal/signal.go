// Package signal implements signal interface for anko script.
package signal

import (
	"github.com/mattn/anko/vm"
	pkg "os/signal"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("os/signal")

	m.DefineType("Ignore", reflect.TypeOf(pkg.Ignore))
	m.DefineType("Notify", reflect.TypeOf(pkg.Notify))
	m.DefineType("Reset", reflect.TypeOf(pkg.Reset))
	m.DefineType("Stop", reflect.TypeOf(pkg.Stop))
	return m
}
