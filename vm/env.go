package vm

import (
	"fmt"
	"reflect"
	"strings"
)

type Env struct {
	name   string
	env    map[string]reflect.Value
	parent *Env
}

func NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: nil}
}

func (e *Env) NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: e, name: e.name}
}

func (e *Env) NewModule(n string) *Env {
	m := &Env{env: make(map[string]reflect.Value), parent: e, name: n}
	e.DefineGlobal(n, reflect.ValueOf(m))
	return m
}

func (e *Env) SetName(n string) {
	e.name = n
}

func (e *Env) GetName() string {
	return e.name
}

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
			if i == len(ns) - 1 {
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

func (e *Env) DefineGlobal(k string, v reflect.Value) {
	global := e
	for global.parent != nil {
		global = global.parent
	}
	global.env[k] = v
}

func (e *Env) Define(k string, v reflect.Value) {
	e.env[k] = v
}

func (e *Env) String() string {
	return e.name
}
