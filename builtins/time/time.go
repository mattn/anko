// Package time implements time interface for anko script.
package math

import (
	"github.com/mattn/anko/vm"
	"reflect"
	t "time"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("After", reflect.ValueOf(t.After))
	m.Define("Sleep", reflect.ValueOf(t.Sleep))
	m.Define("Tick", reflect.ValueOf(t.Tick))
	m.Define("Since", reflect.ValueOf(t.Since))
	m.Define("FixedZone", reflect.ValueOf(t.FixedZone))
	m.Define("LoadLocation", reflect.ValueOf(t.LoadLocation))
	m.Define("NewTicker", reflect.ValueOf(t.NewTicker))
	m.Define("Date", reflect.ValueOf(t.Date))
	m.Define("Now", reflect.ValueOf(t.Now))
	m.Define("Parse", reflect.ValueOf(t.Parse))
	m.Define("ParseDuration", reflect.ValueOf(t.ParseDuration))
	m.Define("ParseInLocation", reflect.ValueOf(t.ParseInLocation))
	m.Define("Unix", reflect.ValueOf(t.Unix))
	m.Define("AfterFunc", reflect.ValueOf(t.AfterFunc))
	m.Define("NewTimer", reflect.ValueOf(t.NewTimer))
	return m
}
