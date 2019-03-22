package vm

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/mattn/anko/ast"
	"github.com/mattn/anko/parser"
)

// EnvResolver provides an interface for extrenal values and types
type EnvResolver interface {
	Get(string) (reflect.Value, error)
	Type(string) (reflect.Type, error)
}

// Env provides interface to run VM. This mean function scope and blocked-scope.
// If stack goes to blocked-scope, it will make new Env.
type Env struct {
	name     string
	env      map[string]reflect.Value
	typ      map[string]reflect.Type
	parent   *Env
	external EnvResolver
	sync.RWMutex
}

var basicTypes = map[string]reflect.Type{
	"interface": reflect.ValueOf([]interface{}{int64(1)}).Index(0).Type(),
	"bool":      reflect.TypeOf(true),
	"string":    reflect.TypeOf("a"),
	"int":       reflect.TypeOf(int(1)),
	"int32":     reflect.TypeOf(int32(1)),
	"int64":     reflect.TypeOf(int64(1)),
	"uint":      reflect.TypeOf(uint(1)),
	"uint32":    reflect.TypeOf(uint32(1)),
	"uint64":    reflect.TypeOf(uint64(1)),
	"byte":      reflect.TypeOf(byte(1)),
	"rune":      reflect.TypeOf('a'),
	"float32":   reflect.TypeOf(float32(1)),
	"float64":   reflect.TypeOf(float64(1)),
}

// NewEnv creates new global scope.
func NewEnv() *Env {
	return &Env{
		env:    make(map[string]reflect.Value),
		typ:    make(map[string]reflect.Type),
		parent: nil,
	}
}

// NewEnv creates new child scope.
func (e *Env) NewEnv() *Env {
	return &Env{
		env:    make(map[string]reflect.Value),
		typ:    make(map[string]reflect.Type),
		parent: e,
		name:   e.name,
	}
}

// NewPackage creates a new env with a name
func NewPackage(n string) *Env {
	return &Env{
		env:    make(map[string]reflect.Value),
		typ:    make(map[string]reflect.Type),
		parent: nil,
		name:   n,
	}
}

// NewPackage creates a new env with a name under the parent env
func (e *Env) NewPackage(n string) *Env {
	return &Env{
		env:    make(map[string]reflect.Value),
		typ:    make(map[string]reflect.Type),
		parent: e,
		name:   n,
	}
}

// AddPackage creates a new env with a name that has methods and types in it. Created under the parent env
func (e *Env) AddPackage(name string, methods map[string]interface{}, types map[string]interface{}) (*Env, error) {
	if strings.Contains(name, ".") {
		return nil, fmt.Errorf("unknown symbol '%s'", name)
	}
	var err error
	pack := e.NewPackage(name)

	for methodName, methodValue := range methods {
		err = pack.Define(methodName, methodValue)
		if err != nil {
			return pack, err
		}
	}
	for typeName, typeValue := range types {
		err = pack.DefineType(typeName, typeValue)
		if err != nil {
			return pack, err
		}
	}

	// can ignore error from Define because of check at the start of the function
	e.Define(name, pack)
	return pack, nil
}

// SetExternal sets an external resolver
func (e *Env) SetExternal(res EnvResolver) {
	e.external = res
}

// NewModule creates new module.
func (e *Env) NewModule(n string) *Env {
	m := &Env{
		env:    make(map[string]reflect.Value),
		typ:    make(map[string]reflect.Type),
		parent: e,
		name:   n,
	}
	e.Define(n, m)
	return m
}

// SetName sets a name of the scope. This means that the scope is module.
func (e *Env) SetName(n string) {
	e.Lock()
	e.name = n
	e.Unlock()
}

// GetName returns module name.
func (e *Env) GetName() string {
	e.RLock()
	defer e.RUnlock()

	return e.name
}

// Addr returns pointer value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Addr(k string) (reflect.Value, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.env[k]; ok {
		if v.CanAddr() {
			return v.Addr(), nil
		}
		return nilValue, fmt.Errorf("unaddressable")
	}
	if e.external != nil {
		v, err := e.external.Get(k)
		if err == nil {
			if v.CanAddr() {
				return v.Addr(), nil
			}
			return nilValue, fmt.Errorf("unaddressable")
		}
	}
	if e.parent == nil {
		return nilValue, fmt.Errorf("undefined symbol '%s'", k)
	}
	return e.parent.Addr(k)
}

// Type returns type which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Type(k string) (reflect.Type, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.typ[k]; ok {
		return v, nil
	}
	if e.external != nil {
		v, err := e.external.Type(k)
		if err == nil {
			return v, nil
		}
	}
	if e.parent == nil {
		if v, ok := basicTypes[k]; ok {
			return v, nil
		}
		return nilType, fmt.Errorf("undefined type '%s'", k)
	}
	return e.parent.Type(k)
}

// Get returns value which specified symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Get(k string) (interface{}, error) {
	rv, err := e.get(k)
	return rv.Interface(), err
}

func (e *Env) get(k string) (reflect.Value, error) {
	e.RLock()
	defer e.RUnlock()

	if v, ok := e.env[k]; ok {
		return v, nil
	}
	if e.external != nil {
		v, err := e.external.Get(k)
		if err == nil {
			return v, nil
		}
	}
	if e.parent == nil {
		return nilValue, fmt.Errorf("undefined symbol '%s'", k)
	}
	return e.parent.get(k)
}

// Set modifies value which specified as symbol. It goes to upper scope until
// found or returns error.
func (e *Env) Set(k string, v interface{}) error {
	if v == nil {
		return e.setValue(k, nilValue)
	}
	return e.setValue(k, reflect.ValueOf(v))
}

func (e *Env) setValue(k string, v reflect.Value) error {
	e.RLock()
	_, ok := e.env[k]
	e.RUnlock()
	if ok {
		e.Lock()
		e.env[k] = v
		e.Unlock()
		return nil
	}
	if e.parent == nil {
		return fmt.Errorf("unknown symbol '%s'", k)
	}
	return e.parent.setValue(k, v)
}

// DefineGlobal defines symbol in global scope.
func (e *Env) DefineGlobal(k string, v interface{}) error {
	for e.parent != nil {
		e = e.parent
	}
	return e.Define(k, v)
}

// defineGlobalValue defines symbol in global scope.
func (e *Env) defineGlobalValue(k string, v reflect.Value) error {
	for e.parent != nil {
		e = e.parent
	}
	return e.defineValue(k, v)
}

// Define defines symbol in current scope.
func (e *Env) Define(k string, v interface{}) error {
	if v == nil {
		return e.defineValue(k, nilValue)
	}
	return e.defineValue(k, reflect.ValueOf(v))
}

// defineValue defines symbol in current scope.
func (e *Env) defineValue(k string, v reflect.Value) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("unknown symbol '%s'", k)
	}

	e.Lock()
	e.env[k] = v
	e.Unlock()

	return nil
}

// Delete deletes symbol in current scope.
func (e *Env) Delete(k string) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}

	e.Lock()
	delete(e.env, k)
	e.Unlock()

	return nil
}

// DeleteGlobal deletes the first matching symbol found in current or parent scope.
func (e *Env) DeleteGlobal(k string) error {
	if e.parent == nil {
		return e.Delete(k)
	}

	e.RLock()
	_, ok := e.env[k]
	e.RUnlock()

	if ok {
		return e.Delete(k)
	}

	return e.parent.DeleteGlobal(k)
}

// DefineGlobalType defines type in global scope.
func (e *Env) DefineGlobalType(k string, t interface{}) error {
	for e.parent != nil {
		e = e.parent
	}
	return e.DefineType(k, t)
}

// DefineGlobalReflectType defines type in global scope.
func (e *Env) DefineGlobalReflectType(k string, t reflect.Type) error {
	for e.parent != nil {
		e = e.parent
	}
	return e.DefineReflectType(k, t)
}

// DefineType defines type in current scope.
func (e *Env) DefineType(k string, t interface{}) error {
	var typ reflect.Type
	if t == nil {
		typ = nilType
	} else {
		var ok bool
		typ, ok = t.(reflect.Type)
		if !ok {
			typ = reflect.TypeOf(t)
		}
	}

	return e.DefineReflectType(k, typ)
}

// DefineReflectType defines type in current scope.
func (e *Env) DefineReflectType(k string, t reflect.Type) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("unknown symbol '%s'", k)
	}

	e.Lock()
	e.typ[k] = t
	e.Unlock()

	return nil
}

// String return the name of current scope.
func (e *Env) String() string {
	e.RLock()
	defer e.RUnlock()

	return e.name
}

// Dump show symbol values in the scope.
func (e *Env) Dump() {
	e.RLock()
	fmt.Printf("Name: %v\n", e.name)
	fmt.Printf("Has parent: %v\n", e.parent != nil)
	for k, v := range e.env {
		fmt.Printf("%v = %#v\n", k, v)
	}
	e.RUnlock()
}

// Execute parses and runs source in current scope.
func (e *Env) Execute(src string) (interface{}, error) {
	return e.ExecuteContext(context.Background(), src)
}

// ExecuteContext parses and runs source in current scope.
func (e *Env) ExecuteContext(ctx context.Context, src string) (interface{}, error) {
	stmt, err := parser.ParseSrc(src)
	if err != nil {
		return nilValue, err
	}
	return RunContext(ctx, stmt, e)
}

// Run runs statement in current scope.
func (e *Env) Run(stmt ast.Stmt) (interface{}, error) {
	return e.RunContext(context.Background(), stmt)
}

// RunContext runs statement in current scope.
func (e *Env) RunContext(ctx context.Context, stmt ast.Stmt) (interface{}, error) {
	return RunContext(ctx, stmt, e)
}

// Copy the state of the virtual machine environment
func (e *Env) Copy() *Env {
	e.Lock()
	defer e.Unlock()
	copy := Env{
		name:     e.name,
		env:      make(map[string]reflect.Value),
		typ:      make(map[string]reflect.Type),
		parent:   e.parent,
		external: e.external,
	}
	for name, value := range e.env {
		copy.env[name] = value
	}
	for name, typ := range e.typ {
		copy.typ[name] = typ
	}
	return &copy
}

// DeepCopy copy recursively the state of the virtual machine environment
func (e *Env) DeepCopy() *Env {
	copy := e.Copy()
	if copy.parent != nil {
		copy.parent = copy.parent.DeepCopy()
	}
	return copy
}
