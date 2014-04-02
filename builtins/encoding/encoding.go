// Package encoding implements encoding interface for anko script.
package io

import (
	"github.com/mattn/anko/builtins/encoding/json"
	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) {
	m := env.NewModule("encoding")
	json.Import(m)
}
