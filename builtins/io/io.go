package io

import (
	"errors"
	"github.com/mattn/anko/vm"
	i "io"
	u "io/ioutil"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("io")
	m.Define("ReadAll", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		r, ok := args[0].Interface().(i.Reader)
		if !ok {
			return vm.NilValue, errors.New("Argument should be Reader")
		}
		b, err := u.ReadAll(r)
		if !ok {
			return vm.NilValue, err
		}
		return reflect.ValueOf(b), nil
	}))
}
