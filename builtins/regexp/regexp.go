// Package regexp implements regexp interface for anko script.
package sort

import (
	"github.com/mattn/anko/vm"
	"reflect"
	r "regexp"
)

func Import(env *vm.Env) {
	m := env.NewModule("regexp")
	m.Define("Match", reflect.ValueOf(r.Match))
	m.Define("MatchReader", reflect.ValueOf(r.MatchReader))
	m.Define("MatchString", reflect.ValueOf(r.MatchString))
	m.Define("QuoteMeta", reflect.ValueOf(r.QuoteMeta))
	m.Define("Compile", reflect.ValueOf(r.Compile))
	m.Define("CompilePOSIX", reflect.ValueOf(r.CompilePOSIX))
	m.Define("MustCompile", reflect.ValueOf(r.MustCompile))
	m.Define("MustCompilePOSIX", reflect.ValueOf(r.MustCompilePOSIX))
}
