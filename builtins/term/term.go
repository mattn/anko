// Package term implements terminal interface for anko script.
package term

import (
	"errors"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/anko/vm"
	"reflect"
)

var ntoc = map[string]ct.Color{
	"none":    ct.None,
	"black":   ct.Black,
	"red":     ct.Red,
	"green":   ct.Green,
	"yellow":  ct.Yellow,
	"blue":    ct.Blue,
	"mazenta": ct.Magenta,
	"cyan":    ct.Cyan,
	"white":   ct.White,
}

func colorOf(name string) ct.Color {
	if c, ok := ntoc[name]; ok {
		return c
	}
	return ct.None
}

func Import(env *vm.Env) {
	m := env.NewModule("term")

	m.Define("ChangeColor", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) != 2 && len(args) != 4 {
			return vm.NilValue, errors.New("Arguments should be 2 or 4")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument #1 should be string")
		}
		if args[1].Kind() != reflect.Bool {
			return vm.NilValue, errors.New("Argument #2 should be bool")
		}
		if len(args) == 4 {
			if args[2].Kind() != reflect.String {
				return vm.NilValue, errors.New("Argument #3 should be string")
			}
			if args[3].Kind() != reflect.Bool {
				return vm.NilValue, errors.New("Argument #4 should be bool")
			}
			ct.ChangeColor(colorOf(args[0].String()), args[1].Bool(), colorOf(args[2].String()), args[3].Bool())
		} else {
			ct.ChangeColor(colorOf(args[0].String()), args[1].Bool(), ct.None, false)
		}
		return vm.NilValue, nil
	}))

	m.Define("ResetColor", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) > 0 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		ct.ResetColor()
		return vm.NilValue, nil
	}))
}
