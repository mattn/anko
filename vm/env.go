package vm

import (
	"fmt"
	"reflect"
)

type Env struct {
	env map[string]reflect.Value
	parent *Env
}

func NewEnv() *Env {
	return &Env {env: make(map[string]reflect.Value), parent: nil}
}

func (e *Env) New() *Env {
	return &Env {env: make(map[string]reflect.Value), parent: e}
}

func (e *Env) Get(k string) (reflect.Value, bool) {
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

func (e *Env) Define(k string, v reflect.Value) {
	e.env[k] = v
}
