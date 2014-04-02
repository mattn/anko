// Package io implements I/O interface for anko script.
package io

import (
	"github.com/mattn/anko/builtins/io/ioutil"
	"github.com/mattn/anko/vm"
)

func Import(env *vm.Env) {
	m := env.NewModule("io")

	ioutil.Import(m)
}
