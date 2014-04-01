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
	m.Define("NewClient", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(&h.Client{}), nil
	}))

	m.Define("NewServeMux", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		return reflect.ValueOf(h.NewServeMux()), nil
	}))

	m.Define("Handle", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 2 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 2 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		handle, ok := args[1].Interface().(h.HandlerFunc)
		if !ok {
			return vm.NilValue, errors.New("Argument should be http.Handler")
		}
		h.Handle(args[0].String(),handle)
		return vm.NilValue, nil
	}))

	m.Define("HandleFunc", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 2 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 2 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		if args[1].Kind() != reflect.Func {
			return vm.NilValue, errors.New("Argument should be function")
		}
		f := args[1].Interface().(vm.Func)
		h.HandleFunc(args[0].String(), func(w h.ResponseWriter, r *h.Request) {
			f(reflect.ValueOf(w), reflect.ValueOf(r))
		})
		return vm.NilValue, nil
	}))

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

	m.Define("ListenAndServe", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 2 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 2 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		err := h.ListenAndServe(args[0].String(), nil)
		return vm.NilValue, err
	}))
}
