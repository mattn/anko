// +build appengine

// Package net implements net interface for anko script.
package net

import (
	"github.com/leonelquinteros/anko/vm"
)

func Import(env *vm.Env) *vm.Env {
	panic("can't import 'net'")
	return nil
}
