package main

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/vm"
	"reflect"
)

func setupBuiltins(env vm.Env) {
	env["println"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		for i, arg := range args {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(arg.Interface())
		}
		fmt.Println()
		return vm.NilValue, nil
	})

	env["len"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Array && args[0].Kind() != reflect.Slice {
			return vm.NilValue, errors.New("Argument should be array")
		}
		return reflect.ValueOf(args[0].Len()), nil
	})

	env["keys"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Map {
			return vm.NilValue, errors.New("Argument should be map")
		}
		keys := []string{}
		mk := args[0].MapKeys()
		for _, key := range mk {
			keys = append(keys, key.String())
		}
		return reflect.ValueOf(keys), nil
	})

	env["bytes"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		return reflect.ValueOf([]byte(args[0].String())), nil
	})

	env["runes"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		return reflect.ValueOf([]rune(args[0].String())), nil
	})

	env["string"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		return reflect.ValueOf(fmt.Sprint(args[0].Interface())), nil
	})

	env["char"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Int && args[0].Kind() != reflect.Int64 {
			return vm.NilValue, errors.New("Argument should be int")
		}
		return reflect.ValueOf(string(rune(args[0].Int()))), nil
	})

	env["rune"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		s := args[0].String()
		if len(s) == 0 {
			return reflect.ValueOf(0), nil
		}
		return reflect.ValueOf(s[0]), nil
	})
}
