// Package json implements json interface to manipulate JSON for anko script.
package os

import (
	"errors"
	"github.com/mattn/anko/vm"
	o "os"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("os")

	m.Define("Stdin", reflect.ValueOf(o.Stdin))
	m.Define("Stdout", reflect.ValueOf(o.Stdout))
	m.Define("Stderr", reflect.ValueOf(o.Stderr))

	m.Define("Getenv", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		return reflect.ValueOf(o.Getenv(args[0].String())), nil
	}))
}
