package builtins

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"reflect"
)

func SetupBuiltins(env *vm.Env) {
	env.Define("println", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		for i, arg := range args {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(arg.Interface())
		}
		fmt.Println()
		return vm.NilValue, nil
	}))

	env.Define("len", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("keys", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("bytes", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("runes", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("string", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		return reflect.ValueOf(fmt.Sprint(args[0].Interface())), nil
	}))

	env.Define("char", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("rune", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
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
	}))

	env.Define("load", vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.String {
			return vm.NilValue, errors.New("Argument should be string")
		}
		body, err := ioutil.ReadFile(args[0].String())
		if err != nil {
			return vm.NilValue, err
		}
		scanner := new(parser.Scanner)
		scanner.Init(string(body))
		stmts, err := parser.Parse(scanner)
		if err != nil {
			return vm.NilValue, err
		}
		return vm.RunStmts(stmts, env)
	}))
}
