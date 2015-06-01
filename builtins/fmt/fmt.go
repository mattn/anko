// Package fmt implements json interface for anko script.
package fmt

import (
	pkg "fmt"
	"github.com/mattn/anko/vm"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Errorf", reflect.ValueOf(pkg.Errorf))
	m.Define("Fprint", reflect.ValueOf(pkg.Fprint))
	m.Define("Fprintf", reflect.ValueOf(pkg.Fprintf))
	m.Define("Fprintln", reflect.ValueOf(pkg.Fprintln))
	m.Define("Fscan", reflect.ValueOf(pkg.Fscan))
	m.Define("Fscanf", reflect.ValueOf(pkg.Fscanf))
	m.Define("Fscanln", reflect.ValueOf(pkg.Fscanln))
	m.Define("Print", reflect.ValueOf(pkg.Print))
	m.Define("Printf", reflect.ValueOf(pkg.Printf))
	m.Define("Println", reflect.ValueOf(pkg.Println))
	m.Define("Scan", reflect.ValueOf(pkg.Scan))
	m.Define("Scanf", reflect.ValueOf(pkg.Scanf))
	m.Define("Scanln", reflect.ValueOf(pkg.Scanln))
	m.Define("Sprint", reflect.ValueOf(pkg.Sprint))
	m.Define("Sprintf", reflect.ValueOf(pkg.Sprintf))
	m.Define("Sprintln", reflect.ValueOf(pkg.Sprintln))
	m.Define("Sscan", reflect.ValueOf(pkg.Sscan))
	m.Define("Sscanf", reflect.ValueOf(pkg.Sscanf))
	m.Define("Sscanln", reflect.ValueOf(pkg.Sscanln))
	return m
}
