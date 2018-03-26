// +build go1.10

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
	Packages["math"]["Erfcinv"] = math.Erfcinv
	Packages["math"]["Erfinv"] = math.Erfinv
	Packages["math"]["Round"] = math.Round
	Packages["math"]["RoundToEven"] = math.RoundToEven
}
