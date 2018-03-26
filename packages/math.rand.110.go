// +build go1.10

package packages

import (
	"math/rand"
)

func init() {
	if _, ok := Packages["math/rand"]; !ok {
		Packages["math/rand"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math/rand"]; !ok {
		PackageTypes["math/rand"] = make(map[string]interface{})
	}
	Packages["math/rand"]["Shuffle"] = rand.Shuffle
}
