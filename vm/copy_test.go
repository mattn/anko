package vm

import (
	"testing"
)

func TestCopy(t *testing.T) {
	env := NewEnv()
	env.Define("a", "a")
	copy := env.Copy()
	if v, e := copy.Get("a"); e != nil || v != "a" {
		t.Errorf("Copy doesn't retain original values")
	}
	copy.Set("a", "b")
	if v, e := env.Get("a"); e != nil || v != "a" {
		t.Errorf("Original was modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "b" {
		t.Errorf("Copy kept the old value")
	}
	env.Set("a", "c")
	if v, e := env.Get("a"); e != nil || v != "c" {
		t.Errorf("Original was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "b" {
		t.Errorf("Copy was modified")
	}
}
