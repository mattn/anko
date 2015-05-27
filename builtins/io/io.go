// Package io implements io interface for anko script.
package io

import (
	"github.com/mattn/anko/vm"
	pkg "io"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Copy", reflect.ValueOf(pkg.Copy))
	m.Define("CopyN", reflect.ValueOf(pkg.CopyN))
	m.Define("EOF", reflect.ValueOf(pkg.EOF))
	m.Define("ErrClosedPipe", reflect.ValueOf(pkg.ErrClosedPipe))
	m.Define("ErrNoProgress", reflect.ValueOf(pkg.ErrNoProgress))
	m.Define("ErrShortBuffer", reflect.ValueOf(pkg.ErrShortBuffer))
	m.Define("ErrShortWrite", reflect.ValueOf(pkg.ErrShortWrite))
	m.Define("ErrUnexpectedEOF", reflect.ValueOf(pkg.ErrUnexpectedEOF))
	m.Define("LimitReader", reflect.ValueOf(pkg.LimitReader))
	m.Define("MultiReader", reflect.ValueOf(pkg.MultiReader))
	m.Define("MultiWriter", reflect.ValueOf(pkg.MultiWriter))
	m.Define("NewSectionReader", reflect.ValueOf(pkg.NewSectionReader))
	m.Define("Pipe", reflect.ValueOf(pkg.Pipe))
	m.Define("ReadAtLeast", reflect.ValueOf(pkg.ReadAtLeast))
	m.Define("ReadFull", reflect.ValueOf(pkg.ReadFull))
	m.Define("TeeReader", reflect.ValueOf(pkg.TeeReader))
	m.Define("WriteString", reflect.ValueOf(pkg.WriteString))
	return m
}
