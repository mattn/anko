// Package os implements os related interface for anko script.
package os

import (
	"github.com/mattn/anko/builtins/net/http"
	"github.com/mattn/anko/builtins/net/url"
	"github.com/mattn/anko/vm"
	n "net"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("net")
	http.Import(m)
	url.Import(m)
	m.Define("Dial", reflect.ValueOf(n.Dial))
}
