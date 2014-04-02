// Package json implements json interface for anko script.
package json

import (
	"encoding/json"
	"github.com/mattn/anko/vm"
	"reflect"
)

func Import(env *vm.Env) {
	m := env.NewModule("json")
	m.Define("Marshal", reflect.ValueOf(json.Marshal))
	m.Define("Unmarshal", reflect.ValueOf(json.Unmarshal))
}
