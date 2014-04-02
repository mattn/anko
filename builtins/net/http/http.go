// Package http implements http interface for anko script.
package http

import (
	"errors"
	"github.com/mattn/anko/vm"
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
	m.Define("DefaultClient", reflect.ValueOf(h.DefaultClient))
	m.Define("NewServeMux", reflect.ValueOf(h.NewServeMux))
	m.Define("Handle", reflect.ValueOf(h.Handle))
	m.Define("HandleFunc", reflect.ValueOf(h.HandleFunc))
	m.Define("ListenAndServe", reflect.ValueOf(h.ListenAndServe))

	m.Define("toHandleFunc", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Func {
			return vm.NilValue, errors.New("Argument should be function")
		}
		f := args[0].Interface().(vm.Func)
		return reflect.ValueOf(func(w h.ResponseWriter, r *h.Request) {
			f(reflect.ValueOf(w), reflect.ValueOf(r))
		}), nil
	}))
}
