// Package flag implements flag interface for anko script.
package flag

import (
	"github.com/mattn/anko/vm"
	f "flag"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("flag")
	m.Define("Usage", reflect.ValueOf(f.Usage))
	m.Define("Parse", reflect.ValueOf(f.Parse))
	m.Define("Parsed", reflect.ValueOf(f.Parsed))
	m.Define("PrintDefaults", reflect.ValueOf(f.PrintDefaults))
	m.Define("Arg", reflect.ValueOf(f.Arg))
	m.Define("Args", reflect.ValueOf(f.Args))
	m.Define("NArg", reflect.ValueOf(f.NArg))
	m.Define("NFlag", reflect.ValueOf(f.NFlag))
	m.Define("Bool", reflect.ValueOf(f.Bool))
	m.Define("Duration", reflect.ValueOf(f.Duration))
	m.Define("Int", reflect.ValueOf(f.Int))
	m.Define("Int64", reflect.ValueOf(f.Int64))
	m.Define("Uint64", reflect.ValueOf(f.Uint64))
	m.Define("Float64", reflect.ValueOf(f.Float64))
	m.Define("String", reflect.ValueOf(f.String))
}
