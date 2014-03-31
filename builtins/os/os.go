package os

import (
	"errors"
	"github.com/mattn/anko/vm"
	"os"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("os")
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
		return reflect.ValueOf(os.Getenv(args[0].String())), nil
	}))
}
