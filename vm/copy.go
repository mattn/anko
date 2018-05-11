package vm

import "reflect"

// Copy the state of the virtual machine environment
func (env *Env) Copy() *Env {
	b := false
	copy := Env {
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		interrupt: &b,
	}
	for name, value := range env.env {
		copy.env[name] = value
	}
	for name, typ := range env.typ {
		copy.typ[name] = typ
	}
	return &copy
}
