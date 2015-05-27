// +build !appengine

// Package url implements url interface for anko script.
package url

import (
	"github.com/mattn/anko/vm"
	u "net/url"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Parse", reflect.ValueOf(u.Parse))
	return m
}
