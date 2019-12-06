// Package core implements core interface for anko script.
package core

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

// Import defines core language builtins - keys, range, println,  etc.
func Import(e *env.Env) *env.Env {
	e.Define("keys", func(v interface{}) []interface{} {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		mapKeysValue := rv.MapKeys()
		mapKeys := make([]interface{}, len(mapKeysValue))
		for i := 0; i < len(mapKeysValue); i++ {
			mapKeys[i] = mapKeysValue[i].Interface()
		}
		return mapKeys
	})

	e.Define("range", func(args ...int64) []int64 {
		var start, stop int64
		var step int64 = 1

		switch len(args) {
		case 0:
			panic("range expected at least 1 argument, got 0")
		case 1:
			stop = args[0]
		case 2:
			start = args[0]
			stop = args[1]
		case 3:
			start = args[0]
			stop = args[1]
			step = args[2]
			if step == 0 {
				panic("range argument 3 must not be zero")
			}
		default:
			panic(fmt.Sprintf("range expected at most 3 arguments, got %d", len(args)))
		}

		arr := []int64{}
		for i := start; (step > 0 && i < stop) || (step < 0 && i > stop); i += step {
			arr = append(arr, i)
		}
		return arr
	})

	e.Define("typeOf", func(v interface{}) string {
		return reflect.TypeOf(v).String()
	})

	e.Define("kindOf", func(v interface{}) string {
		typeOf := reflect.TypeOf(v)
		if typeOf == nil {
			return "nil"
		}
		return typeOf.Kind().String()
	})

	e.Define("defined", func(s string) bool {
		_, err := e.Get(s)
		return err == nil
	})

	e.Define("load", func(s string) interface{} {
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
		rv, err := vm.Run(e, nil, stmts)
		if err != nil {
			panic(err)
		}
		return rv
	})

	e.Define("print", fmt.Print)
	e.Define("println", fmt.Println)
	e.Define("printf", fmt.Printf)

	ImportToX(e)

	return e
}
