package vm

import (
	"fmt"
	"reflect"
	"strings"
)

// Env provides interface to run VM. This mean function scope and blocked-scope.
// If stack go to blocked-scope, it will make new Env.
type Env struct {
	name   string
	env    map[string]reflect.Value
	parent *Env
}

// NewEnv create new global scope.
func NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: nil}
}

// NewEnv create new child scope.
func (e *Env) NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: e, name: e.name}
}

// NewEnv create new module scope as global.
func (e *Env) NewModule(n string) *Env {
	m := &Env{env: make(map[string]reflect.Value), parent: e, name: n}
	e.DefineGlobal(n, reflect.ValueOf(m))
	return m
}

// SetName make a name of the scope. This mean that the scope is module.
func (e *Env) SetName(n string) {
	e.name = n
}

// GetName return module name.
func (e *Env) GetName() string {
	return e.name
}

// Get return value which specified symbol. It go to upper scope until found.
func (e *Env) Get(k string) (reflect.Value, bool) {
	ns := strings.Split(k, "::")
	if len(ns) > 1 {
		global := e
		for global.parent != nil {
			global = global.parent
		}
		for i, n := range ns {
			v, ok := global.env[n]
			if !ok {
				return NilValue, false
			}
			if i == len(ns)-1 {
				return v, ok
			}
			if vv, ok := v.Interface().(*Env); ok {
				global = vv
			} else {
				return NilValue, false
			}
		}
		return NilValue, false
	}

	for {
		if e.parent == nil {
			v, ok := e.env[k]
			return v, ok
		}
		if v, ok := e.env[k]; ok {
			return v, ok
		}
		e = e.parent
	}
	return NilValue, false
}

// Set modify the value which specified as symbol. If it can't be found in
// whole. This function return error
func (e *Env) Set(k string, v reflect.Value) error {
	for {
		if e.parent == nil {
			if _, ok := e.env[k]; !ok {
				return fmt.Errorf("Unknown symbol '%s'", k)
			}
			e.env[k] = v
			return nil
		}
		if _, ok := e.env[k]; ok {
			e.env[k] = v
			return nil
		}
		e = e.parent
	}
}

// DefineGlobal defines global symbol.
func (e *Env) DefineGlobal(k string, v reflect.Value) {
	global := e
	for global.parent != nil {
		global = global.parent
	}
	global.env[k] = v
}

// Define defines symbol in current scope.
func (e *Env) Define(k string, v reflect.Value) {
	e.env[k] = v
}

// String return the name of current scope.
func (e *Env) String() string {
	return e.name
}

// Dump show symbol values in the scope.
func (e *Env) Dump() {
	for k, v := range e.env {
		fmt.Printf("%v = %v\n", k, v)
	}
}
