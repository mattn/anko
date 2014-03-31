package http

import (
	"errors"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	h "net/http"
	"reflect"
)

type Client struct {
	c *h.Client
}

func (c *Client) Get(args ...reflect.Value) (reflect.Value, error) {
	if len(args) < 1 {
		return vm.NilValue, errors.New("Missing arguments")
	}
	if len(args) > 1 {
		return vm.NilValue, errors.New("Too many arguments")
	}
	if args[0].Kind() != reflect.String {
		return vm.NilValue, errors.New("Argument should be string")
	}
	res, err := h.Get(args[0].String())
	return reflect.ValueOf(res), err
}

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
		res, err := h.Get(args[0].String())
		if err != nil {
			return vm.NilValue, err
		}
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		return reflect.ValueOf(map[string]interface{} {
			"headers": map[string][]string(res.Header),
			"content": b,
		}), nil
	}))

	m.Define("NewClient", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(&Client{}), nil
	}))
}
