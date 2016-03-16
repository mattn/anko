// Package path implements path manipulation interface for anko script.
package filepath

import (
	"github.com/mattn/anko/vm"
	f "path/filepath"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("filepath")
	m.Define("Join", reflect.ValueOf(f.Join))
	m.Define("Clean", reflect.ValueOf(f.Join))
	m.Define("Abs", reflect.ValueOf(f.Abs))
	m.Define("Base", reflect.ValueOf(f.Base))
	m.Define("Clean", reflect.ValueOf(f.Clean))
	m.Define("Dir", reflect.ValueOf(f.Dir))
	m.Define("EvalSymlinks", reflect.ValueOf(f.EvalSymlinks))
	m.Define("Ext", reflect.ValueOf(f.Ext))
	m.Define("FromSlash", reflect.ValueOf(f.FromSlash))
	m.Define("Glob", reflect.ValueOf(f.Glob))
	m.Define("HasPrefix", reflect.ValueOf(f.HasPrefix))
	m.Define("IsAbs", reflect.ValueOf(f.IsAbs))
	m.Define("Join", reflect.ValueOf(f.Join))
	m.Define("Match", reflect.ValueOf(f.Match))
	m.Define("Rel", reflect.ValueOf(f.Rel))
	m.Define("Split", reflect.ValueOf(f.Split))
	m.Define("SplitList", reflect.ValueOf(f.SplitList))
	m.Define("ToSlash", reflect.ValueOf(f.ToSlash))
	m.Define("VolumeName", reflect.ValueOf(f.VolumeName))
	return m
}
