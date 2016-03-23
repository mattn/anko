package vm

import (
	"fmt"
	"reflect"
	"strings"
)

// Env provides interface to run VM. This mean function scope and blocked-scope.
// If stack goes to blocked-scope, it will make new Env.
type Env struct {
	name      string
	env       map[string]reflect.Value
	typ       map[string]reflect.Type
	parent    *Env
	interrupt *bool
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	b := false

	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		interrupt: &b,
	}
}

// NewEnv creates new child scope.
func (e *Env) NewEnv() *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    e,
		name:      e.name,
		interrupt: e.interrupt,
	}
}

func NewPackage(n string) *Env {
	b := false

	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    nil,
		name:      n,
		interrupt: &b,
	}
}

func (e *Env) NewPackage(n string) *Env {
	return &Env{
		env:       make(map[string]reflect.Value),
		typ:       make(map[string]reflect.Type),
		parent:    e,
		name:      n,
		interrupt: e.interrupt,
	}
}

// Destroy deletes current scope.
func (e *Env) Destroy() {
	if e.parent == nil {
		return
	}
	for k, v := range e.parent.env {
		if v.IsValid() && v.Interface() == e {
			delete(e.parent.env, k)
		}
	}
	e.parent = nil
	e.env = nil
}

// NewModule creates new module scope as global.
func (e *Env) NewModule(n string) *Env {
	m := &Env{env: make(map[string]reflect.Value), parent: e, name: n}
	e.Define(n, reflect.ValueOf(m))
	return m
}

// SetName sets a name of the scope. This means that the scope is module.
func (e *Env) SetName(n string) {
	e.name = n
}

// GetName returns module name.
func (e *Env) GetName() string {
	return e.name
}

// Addr returns pointer value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Addr(k string) (reflect.Value, error) {
	for {
		if e.parent == nil {
			v, ok := e.env[k]
			if !ok {
				return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
			}
			return v, nil
		}
		if v, ok := e.env[k]; ok {
			return v.Addr(), nil
		}
		e = e.parent
	}
	return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
}

// Type returns type which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Type(k string) (reflect.Type, error) {
	for {
		if e.parent == nil {
			v, ok := e.typ[k]
			if !ok {
				return NilType, fmt.Errorf("Undefined type '%s'", k)
			}
			return v, nil
		}
		if v, ok := e.typ[k]; ok {
			return v, nil
		}
		e = e.parent
	}
	return NilType, fmt.Errorf("Undefined type '%s'", k)
}

// Get returns value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Get(k string) (reflect.Value, error) {
	for {
		if e.parent == nil {
			v, ok := e.env[k]
			if !ok {
				return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
			}
			return v, nil
		}
		if v, ok := e.env[k]; ok {
			return v, nil
		}
		e = e.parent
	}
	return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
}

// Set modifies value which specified as symbol. It goes to upper scope until
// found or returns error.
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
	return nil
}

// DefineGlobal defines symbol in global scope.
func (e *Env) DefineGlobal(k string, v reflect.Value) error {
	global := e
	for global.parent != nil {
		global = global.parent
	}
	return global.Define(k, v)
}

// DefineType defines type which specifis symbol in global scope.
func (e *Env) DefineType(k string, v reflect.Type) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	global := e
	name := []string{}
	for {
		if global.parent == nil {
			break
		}
		if global.name != "" {
			name = append(name, global.name)
		}
		global = global.parent
	}
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}

	if len(name) > 0 {
		global.typ[strings.Join(name, ".")+"."+k] = v
	} else {
		global.typ[k] = v
	}
	return nil
}

// Define defines symbol in current scope.
func (e *Env) Define(k string, v reflect.Value) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	e.env[k] = v
	return nil
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

// Execute parses and runs source in current scope.
func (e *Env) Execute(src string) (reflect.Value, error) {
	stmts, err := parser.ParseSrc(src)
	if err != nil {
		return NilValue, err
	}
	return Run(stmts, e)
}
