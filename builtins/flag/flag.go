// Package flag implements flag interface for anko script.
package flag

import (
	pkg "flag"
	"github.com/mattn/anko/vm"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Arg", reflect.ValueOf(pkg.Arg))
	m.Define("Args", reflect.ValueOf(pkg.Args))
	m.Define("Bool", reflect.ValueOf(pkg.Bool))
	m.Define("BoolVar", reflect.ValueOf(pkg.BoolVar))
	m.Define("CommandLine", reflect.ValueOf(pkg.CommandLine))
	m.Define("ContinueOnError", reflect.ValueOf(pkg.ContinueOnError))
	m.Define("Duration", reflect.ValueOf(pkg.Duration))
	m.Define("DurationVar", reflect.ValueOf(pkg.DurationVar))
	m.Define("ErrHelp", reflect.ValueOf(pkg.ErrHelp))
	m.Define("ExitOnError", reflect.ValueOf(pkg.ExitOnError))
	m.Define("Float64", reflect.ValueOf(pkg.Float64))
	m.Define("Float64Var", reflect.ValueOf(pkg.Float64Var))
	m.Define("Int", reflect.ValueOf(pkg.Int))
	m.Define("Int64", reflect.ValueOf(pkg.Int64))
	m.Define("Int64Var", reflect.ValueOf(pkg.Int64Var))
	m.Define("IntVar", reflect.ValueOf(pkg.IntVar))
	m.Define("Lookup", reflect.ValueOf(pkg.Lookup))
	m.Define("NArg", reflect.ValueOf(pkg.NArg))
	m.Define("NFlag", reflect.ValueOf(pkg.NFlag))
	m.Define("NewFlagSet", reflect.ValueOf(pkg.NewFlagSet))
	m.Define("PanicOnError", reflect.ValueOf(pkg.PanicOnError))
	m.Define("Parse", reflect.ValueOf(pkg.Parse))
	m.Define("Parsed", reflect.ValueOf(pkg.Parsed))
	m.Define("PrintDefaults", reflect.ValueOf(pkg.PrintDefaults))
	m.Define("Set", reflect.ValueOf(pkg.Set))
	m.Define("String", reflect.ValueOf(pkg.String))
	m.Define("StringVar", reflect.ValueOf(pkg.StringVar))
	m.Define("Uint", reflect.ValueOf(pkg.Uint))
	m.Define("Uint64", reflect.ValueOf(pkg.Uint64))
	m.Define("Uint64Var", reflect.ValueOf(pkg.Uint64Var))
	m.Define("UintVar", reflect.ValueOf(pkg.UintVar))
	m.Define("Usage", reflect.ValueOf(pkg.Usage))
	m.Define("Var", reflect.ValueOf(pkg.Var))
	m.Define("Visit", reflect.ValueOf(pkg.Visit))
	m.Define("VisitAll", reflect.ValueOf(pkg.VisitAll))
	return m
}
