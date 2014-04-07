// Package encoding implements encoding interface for anko script.
package io

import (
	pkg_json "github.com/mattn/anko/builtins/encoding/json"
	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) {
	m := env.NewModule("encoding")
	pkg_json.Import(m)
}
