// +build !appengine

package os

import (
	"github.com/leonelquinteros/anko/vm"
	pkg "os"
	"reflect"
)

func handleAppEngine(m *vm.Env) {
	m.Define("Getppid", reflect.ValueOf(pkg.Getppid))
}
