
// Package strings implements strings interface for anko script.
package sort

import (
	"github.com/mattn/anko/vm"
	"reflect"
	s "strings"
)

func Import(env *vm.Env) {
	m := env.NewModule("strings")
	m.Define("Count", reflect.ValueOf(s.Count))
	m.Define("Contains", reflect.ValueOf(s.Contains))
	m.Define("ContainsAny", reflect.ValueOf(s.ContainsAny))
	m.Define("ContainsRune", reflect.ValueOf(s.ContainsRune))
	m.Define("Index", reflect.ValueOf(s.Index))
	m.Define("LastIndex", reflect.ValueOf(s.LastIndex))
	m.Define("IndexRune", reflect.ValueOf(s.IndexRune))
	m.Define("IndexAny", reflect.ValueOf(s.IndexAny))
	m.Define("LastIndexAny", reflect.ValueOf(s.LastIndexAny))
	m.Define("SplitN", reflect.ValueOf(s.SplitN))
	m.Define("SplitAfterN", reflect.ValueOf(s.SplitAfterN))
	m.Define("Split", reflect.ValueOf(s.Split))
	m.Define("SplitAfter", reflect.ValueOf(s.SplitAfter))
	m.Define("Fields", reflect.ValueOf(s.Fields))
	m.Define("FieldsFunc", reflect.ValueOf(s.FieldsFunc))
	m.Define("Join", reflect.ValueOf(s.Join))
	m.Define("HasPrefix", reflect.ValueOf(s.HasPrefix))
	m.Define("HasSuffix", reflect.ValueOf(s.HasSuffix))
	m.Define("Map", reflect.ValueOf(s.Map))
	m.Define("Repeat", reflect.ValueOf(s.Repeat))
	m.Define("ToUpper", reflect.ValueOf(s.ToUpper))
	m.Define("ToLower", reflect.ValueOf(s.ToLower))
	m.Define("ToTitle", reflect.ValueOf(s.ToTitle))
	m.Define("ToUpperSpecial", reflect.ValueOf(s.ToUpperSpecial))
	m.Define("ToLowerSpecial", reflect.ValueOf(s.ToLowerSpecial))
	m.Define("ToTitleSpecial", reflect.ValueOf(s.ToTitleSpecial))
	m.Define("Title", reflect.ValueOf(s.Title))
	m.Define("TrimLeftFunc", reflect.ValueOf(s.TrimLeftFunc))
	m.Define("TrimRightFunc", reflect.ValueOf(s.TrimRightFunc))
	m.Define("TrimFunc", reflect.ValueOf(s.TrimFunc))
	m.Define("IndexFunc", reflect.ValueOf(s.IndexFunc))
	m.Define("LastIndexFunc", reflect.ValueOf(s.LastIndexFunc))
	m.Define("Trim", reflect.ValueOf(s.Trim))
	m.Define("TrimLeft", reflect.ValueOf(s.TrimLeft))
	m.Define("TrimRight", reflect.ValueOf(s.TrimRight))
	m.Define("TrimSpace", reflect.ValueOf(s.TrimSpace))
	m.Define("TrimPrefix", reflect.ValueOf(s.TrimPrefix))
	m.Define("TrimSuffix", reflect.ValueOf(s.TrimSuffix))
	m.Define("Replace", reflect.ValueOf(s.Replace))
	m.Define("EqualFold", reflect.ValueOf(s.EqualFold))
}
