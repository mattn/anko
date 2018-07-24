package corelib

import (
	"reflect"

	"github.com/mattn/anko/ast"
)

// Env is an interface representing vm.Env.
type Env interface {
	DefineType(string, interface{}) error
	Define(string, interface{}) error
	Get(string) (interface{}, error)
	Run([]ast.Stmt) (interface{}, error)
}

// NewEnv is function provided for tests
var NewEnv func() Env

// ValueEqual return true if v1 and v2 is same value. If passed function, does
// extra checks otherwise just doing reflect.DeepEqual
func ValueEqual(v1 interface{}, v2 interface{}) bool {
	v1RV := reflect.ValueOf(v1)
	switch v1RV.Kind() {
	case reflect.Func:
		// This is best effort to check if functions match, but it could be wrong
		v2RV := reflect.ValueOf(v2)
		if !v1RV.IsValid() || !v2RV.IsValid() {
			if v1RV.IsValid() != !v2RV.IsValid() {
				return false
			}
			return true
		} else if v1RV.Kind() != v2RV.Kind() {
			return false
		} else if v1RV.Type() != v2RV.Type() {
			return false
		} else if v1RV.Pointer() != v2RV.Pointer() {
			// From reflect: If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.
			return false
		}
		return true
	}
	return reflect.DeepEqual(v1, v2)
}
