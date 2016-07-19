// +build appengine

// Package url implements url interface for anko script.
package url

import (
	"github.com/leonelquinteros/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	panic("can't import 'url'")
	return nil
}
