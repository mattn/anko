// +build go1.8

package packages

import (
	"math/rand"
	"reflect"
)

func init() {
	if _, ok := Packages["math/rand"]; !ok {
		Packages["math/rand"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math/rand"]; !ok {
		PackageTypes["math/rand"] = make(map[string]interface{})
	}
	var source64 rand.Source64
	Packages["math/rand"]["Uint64"] = rand.Uint64
	PackageTypes["math/rand"]["Source64"] = reflect.TypeOf(&source64).Elem()
}
