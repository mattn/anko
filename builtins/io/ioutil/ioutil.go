// Package ioutil implements I/O interface for anko script.
package ioutil

import (
	"github.com/mattn/anko/vm"
	u "io/ioutil"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("ioutil")
	m.Define("ReadAll", reflect.ValueOf(u.ReadAll))
}
