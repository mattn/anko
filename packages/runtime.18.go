// +build go1.8

package packages

import (
	"runtime"
)

func init() {
	if _, ok := Packages["runtime"]; !ok {
		Packages["runtime"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["runtime"]; !ok {
		PackageTypes["runtime"] = make(map[string]interface{})
	}
	Packages["runtime"]["MutexProfile"] = runtime.MutexProfile
	Packages["runtime"]["SetMutexProfileFraction"] = runtime.SetMutexProfileFraction
}
