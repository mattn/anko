// Package signal implements signal interface for anko script.
package signal

import (
	"github.com/mattn/anko/vm"
	pkg "os/signal"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("os/signal")

	//m.Define("Ignore", reflect.ValueOf(pkg.Ignore))
	m.Define("Notify", reflect.ValueOf(pkg.Notify))
	//m.Define("Reset", reflect.ValueOf(pkg.Reset))
	m.Define("Stop", reflect.ValueOf(pkg.Stop))
	return m
}
