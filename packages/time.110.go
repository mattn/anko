// +build go1.10

package packages

import (
	"time"
)

func init() {
	if _, ok := Packages["time"]; !ok {
		Packages["time"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["time"]; !ok {
		PackageTypes["time"] = make(map[string]interface{})
	}
	Packages["time"]["LoadLocationFromTZData"] = time.LoadLocationFromTZData
}
