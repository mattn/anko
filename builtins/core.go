// Package core implements core interface for anko script.
package core

import (
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"os"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	env.Define("len", reflect.ValueOf(func(v interface{}) int64 {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.String {
			return int64(len([]byte(rv.String())))
		}
		if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
			panic("Argument #1 should be array")
		}
		return int64(rv.Len())
	}))

	env.Define("keys", reflect.ValueOf(func(v interface{}) []string {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Map {
			panic("Argument #1 should be map")
		}
		keys := []string{}
		mk := rv.MapKeys()
		for _, key := range mk {
			keys = append(keys, key.String())
		}
		return keys
	}))

	env.Define("range", reflect.ValueOf(func(args ...int64) []int64 {
		if len(args) < 1 {
			panic("Missing arguments")
		}
		if len(args) > 2 {
			panic("Too many arguments")
		}
		var min, max int64
		if len(args) == 1 {
			min = 0
			max = args[0] - 1
		} else {
			min = args[0]
			max = args[1]
		}
		arr := []int64{}
		for i := min; i <= max; i++ {
			arr = append(arr, i)
		}
		return arr
	}))

	env.Define("toBytes", reflect.ValueOf(func(s string) []byte {
		return []byte(s)
	}))

	env.Define("toRunes", reflect.ValueOf(func(s string) []rune {
		return []rune(s)
	}))

	env.Define("toString", reflect.ValueOf(func(v interface{}) string {
		return fmt.Sprint(v)
	}))

	env.Define("toInt", reflect.ValueOf(func(v interface{}) int64 {
		nt := reflect.TypeOf(1)
		rv := reflect.ValueOf(v)
		if rv.Type().ConvertibleTo(nt) {
			return 0
		}
		return rv.Convert(nt).Int()
	}))

	env.Define("toFloat", reflect.ValueOf(func(v interface{}) float64 {
		nt := reflect.TypeOf(1.0)
		rv := reflect.ValueOf(v)
		if rv.Type().ConvertibleTo(nt) {
			return 0.0
		}
		return rv.Convert(nt).Float()
	}))

	env.Define("toBool", reflect.ValueOf(func(v interface{}) bool {
		nt := reflect.TypeOf(true)
		rv := reflect.ValueOf(v)
		if rv.Type().ConvertibleTo(nt) {
			return false
		}
		return rv.Convert(nt).Bool()
	}))

	env.Define("toChar", reflect.ValueOf(func(s rune) string {
		return string(s)
	}))

	env.Define("toRune", reflect.ValueOf(func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	}))

	env.Define("string", reflect.ValueOf(func(b []byte) string {
		return string(b)
	}))

	env.Define("typeof", reflect.ValueOf(func(v interface{}) string {
		return reflect.TypeOf(v).String()
	}))

	env.Define("defined", reflect.ValueOf(func(s string) bool {
		_, err := env.Get(s)
		return err == nil
	}))

	env.Define("load", reflect.ValueOf(func(s string) interface{} {
		body, err := ioutil.ReadFile(s)
		if err != nil {
			panic(err)
		}
		scanner := new(parser.Scanner)
		scanner.Init(string(body))
		stmts, err := parser.Parse(scanner)
		if err != nil {
			if pe, ok := err.(*parser.Error); ok {
				pe.Filename = s
				panic(pe)
			}
			panic(err)
		}
		rv, err := vm.Run(stmts, env)
		if err != nil {
			panic(err)
		}
		if rv.IsValid() && rv.CanInterface() {
			return rv.Interface()
		}
		return nil
	}))

	env.Define("panic", reflect.ValueOf(func(e interface{}) {
		os.Setenv("ANKO_DEBUG", "1")
		panic(e)
	}))

	env.Define("print", reflect.ValueOf(fmt.Print))
	env.Define("println", reflect.ValueOf(fmt.Println))
	env.Define("printf", reflect.ValueOf(fmt.Printf))

	return env
}
