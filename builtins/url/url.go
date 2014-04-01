// Package url implements url interface for anko script.
package url

import (
	"errors"
	"github.com/mattn/anko/vm"
	u "net/url"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("url")
	m.Define("Parse", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		ui, err := u.Parse(args[0].String())
		if err != nil {
			return vm.NilValue, err
		}
		return reflect.ValueOf(ui), nil
	}))
}
