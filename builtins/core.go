// Package core implements core interface for anko script.
package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	env.Define("len", func(v interface{}) int64 {
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
	})

	env.Define("keys", func(v interface{}) []string {
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
	})

	env.Define("range", func(args ...int64) []int64 {
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
	})

	env.Define("toString", func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return string(b)
		}
		return fmt.Sprint(v)
	})

	env.Define("toInt", func(v interface{}) int64 {
		nt := reflect.TypeOf(1)
		rv := reflect.ValueOf(v)
		if !rv.Type().ConvertibleTo(nt) {
			return 0
		}
		return rv.Convert(nt).Int()
	})

	env.Define("toFloat", func(v interface{}) float64 {
		nt := reflect.TypeOf(1.0)
		rv := reflect.ValueOf(v)
		if !rv.Type().ConvertibleTo(nt) {
			return 0.0
		}
		return rv.Convert(nt).Float()
	})

	env.Define("toBool", func(v interface{}) bool {
		nt := reflect.TypeOf(true)
		rv := reflect.ValueOf(v)
		if !rv.Type().ConvertibleTo(nt) {
			return false
		}
		return rv.Convert(nt).Bool()
	})

	env.Define("toChar", func(s rune) string {
		return string(s)
	})

	env.Define("toRune", func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	})

	env.Define("toByteSlice", func(s string) []byte {
		return []byte(s)
	})

	env.Define("toRuneSlice", func(s string) []rune {
		return []rune(s)
	})

	env.Define("toBoolSlice", func(v []interface{}) []bool {
		var result []bool
		toSlice(v, &result)
		return result
	})

	env.Define("toFloatSlice", func(v []interface{}) []float64 {
		var result []float64
		toSlice(v, &result)
		return result
	})

	env.Define("toIntSlice", func(v []interface{}) []int64 {
		var result []int64
		toSlice(v, &result)
		return result
	})

	env.Define("toStringSlice", func(v []interface{}) []string {
		var result []string
		toSlice(v, &result)
		return result
	})

	env.Define("typeOf", func(v interface{}) string {
		return reflect.TypeOf(v).String()
	})

	env.Define("chanOf", func(t reflect.Type) reflect.Value {
		return reflect.MakeChan(t, 1)
	})

	env.Define("defined", func(s string) bool {
		_, err := env.Get(s)
		return err == nil
	})

	env.Define("load", func(s string) interface{} {
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
	})

	env.Define("panic", func(e interface{}) {
		os.Setenv("ANKO_DEBUG", "1")
		panic(e)
	})

	env.Define("print", fmt.Print)
	env.Define("println", fmt.Println)
	env.Define("printf", fmt.Printf)

	env.DefineType("int64", int64(0))
	env.DefineType("float64", float64(0.0))
	env.DefineType("bool", true)
	env.DefineType("string", "")
	return env
}

// toSlice takes in a "generic" slice and converts and copies
// it's elements into the typed slice pointed at by ptr.
// Note that this is a costly operation.
func toSlice(from []interface{}, ptr interface{}) {
	// Value of the pointer to the target
	obj := reflect.Indirect(reflect.ValueOf(ptr))
	// We can't just convert from interface{} to whatever the target is (diff memory layout),
	// so we need to create a New slice of the proper type and copy the values individually
	t := reflect.TypeOf(ptr).Elem()
	slice := reflect.MakeSlice(t, len(from), len(from))
	// Copying the data, val is an adressable Pointer of the actual target type
	val := reflect.Indirect(reflect.New(t.Elem()))
	for i := 0; i < len(from); i++ {
		v := reflect.ValueOf(from[i])
		val.Set(v)
		slice.Index(i).Set(v)
	}
	// Ok now assign our slice to the target pointer
	obj.Set(slice)
}
