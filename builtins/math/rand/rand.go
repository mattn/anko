// Package rand implements math/rand interface for anko script.
package rand

import (
	"github.com/mattn/anko/vm"
	t "math/rand"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("ExpFloat64", reflect.ValueOf(t.ExpFloat64))
	m.Define("Float32", reflect.ValueOf(t.Float32))
	m.Define("Float64", reflect.ValueOf(t.Float64))
	m.Define("Int", reflect.ValueOf(t.Int))
	m.Define("Int31", reflect.ValueOf(t.Int31))
	m.Define("Int31n", reflect.ValueOf(t.Int31n))
	m.Define("Int63", reflect.ValueOf(t.Int63))
	m.Define("Int63n", reflect.ValueOf(t.Int63n))
	m.Define("Intn", reflect.ValueOf(t.Intn))
	m.Define("NormFloat64", reflect.ValueOf(t.NormFloat64))
	m.Define("Perm", reflect.ValueOf(t.Perm))
	m.Define("Read", reflect.ValueOf(t.Read))
	m.Define("Seed", reflect.ValueOf(t.Seed))
	m.Define("Uint32", reflect.ValueOf(t.Uint32))
	return m
}
