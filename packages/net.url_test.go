// +build !appengine

package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestURL(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	env := vm.NewEnv()
	DefineImport(env)
	value, err := env.Execute(`
url = import("net/url")
v1 = make("url.Values")
v1.Set("a", "a")
if v1.Get("a") != "a" {
	return "value a not set"
} 
v2 = make("url.Values") 
v2.Set("b", "b")
if v2.Get("b") != "b" {
	return "value b not set"
} 
v2.Get("a")
`)
	if err != nil {
		t.Errorf("Execute error - received: %v expected: %v", err, nil)
	}
	if value != "" {
		t.Errorf("Execute value - received: %#v expected: %#v", value, "")
	}
}
