// Package ioutil implements I/O interface for anko script.
package ioutil

import (
	"github.com/mattn/anko/vm"
	u "io/ioutil"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("ReadAll", reflect.ValueOf(u.ReadAll))
	m.Define("ReadDir", reflect.ValueOf(u.ReadDir))
	m.Define("ReadFile", reflect.ValueOf(u.ReadFile))
	m.Define("WriteFile", reflect.ValueOf(u.WriteFile))
	return m
}
