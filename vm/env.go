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
	e.Define(n, reflect.ValueOf(m))
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
func (e *Env) Get(k string) (reflect.Value, error) {
	tok := strings.Split(k, ".")
	v := NilValue

	for e != nil {
		v = reflect.ValueOf(e)
		found := true
		for _, n := range tok {
			if vv, ok := v.Interface().(*Env); ok {
				ve, ok := vv.env[n]
				if !ok {
					found = false
					break
				}
				v = ve
			} else {
				if v.Kind() == reflect.Interface {
					v = v.Elem()
				}
				if v.Kind() == reflect.Slice {
					v = v.Index(0)
				}

				if !v.IsValid() {
					found = false
					break
				}

				m := v.MethodByName(n)
				if m.IsValid() {
					v = m
				} else {
					if v.Kind() == reflect.Ptr {
						v = v.Elem()
					}
					if v.Kind() == reflect.Struct {
						v = v.FieldByName(n)
					} else if v.Kind() == reflect.Map {
						v = v.MapIndex(reflect.ValueOf(n))
					}
				}
			}
		}

		if found {
			break
		}
		e = e.parent
		if e == nil {
			return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
		}
	}

	return v, nil
}

// Set modify the value which specified as symbol. If it can't be found in
// whole. This function return error
func (e *Env) Set(k string, nv reflect.Value) error {
	tok := strings.Split(k, ".")

	if len(tok) == 1 {
		for {
			if e.parent == nil {
				if _, ok := e.env[k]; !ok {
					return fmt.Errorf("Unknown symbol '%s'", k)
				}
				e.env[k] = nv
				return nil
			}
			if _, ok := e.env[k]; ok {
				e.env[k] = nv
				return nil
			}
			e = e.parent
		}
	}

	v := reflect.ValueOf(e)
	for e != nil {
		found := true
		for i, n := range tok {
			if vv, vok := v.Interface().(*Env); vok {
				v, _ = vv.env[n]
				if !v.IsValid() {
					return fmt.Errorf("nil reference for '%s'", strings.Join(tok[:i], "."))
				}
			} else {
				if v.Kind() == reflect.Interface {
					v = v.Elem()
				}
				if v.Kind() == reflect.Slice {
					v = v.Index(0)
				}

				if !v.IsValid() {
					return fmt.Errorf("nil reference for '%s'", strings.Join(tok[:i], "."))
				}

				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					v = v.FieldByName(n)
				} else if v.Kind() == reflect.Map {
					v = v.MapIndex(reflect.ValueOf(n))
				}
			}
		}

		if found {
			if v.CanSet() {
				v.Set(nv)
			} else {
				return fmt.Errorf("Invalid assign operation '%s'", k)
			}
			break
		}
		e = e.parent
		if e == nil {
			return fmt.Errorf("Undefined symbol '%s'", k)
		}
	}
	return nil
}

// DefineGlobal defines global symbol.
func (e *Env) DefineGlobal(k string, v reflect.Value) error {
	global := e
	for global.parent != nil {
		global = global.parent
	}
	return global.Set(k, v)
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
