package json

import (
	"encoding/json"
	"errors"
	"github.com/mattn/anko/vm"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("json")
	m.Define("Marshal", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		b, err := json.Marshal(args[0].Interface())
		if err != nil {
			return vm.NilValue, err
		}
		return reflect.ValueOf(string(b)), nil
	}))

	m.Define("Unmarshal", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		var v interface{}
		err := json.Unmarshal([]byte(args[0].String()), &v)
		return reflect.ValueOf(v), err
	}))
}
