package http

import (
	"errors"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"net/http"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("http")
	m.Define("get", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		res, err := http.Get(args[0].String())
		if err != nil {
			return vm.NilValue, err
		}
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		return reflect.ValueOf(map[string]interface{} {
			"headers": res.Header,
			"content": b,
		}), nil
	}))
}
