// +build !appengine

// Package url implements url interface for anko script.
package url

import (
	"github.com/mattn/anko/vm"
	u "net/url"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("url")
	m.Define("Parse", reflect.ValueOf(u.Parse))
}
