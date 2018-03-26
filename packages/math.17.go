// +build go1.7

package packages

import (
	"math"
)

func init() {
	if _, ok := Packages["math"]; !ok {
		Packages["math"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math"]; !ok {
		PackageTypes["math"] = make(map[string]interface{})
	}
	Packages["math"]["MaxUint64"] = uint64(math.MaxUint64)
}
