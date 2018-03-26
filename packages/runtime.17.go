// +build go1.7

package packages

import (
	"reflect"
	"runtime"
)

func init() {
	if _, ok := Packages["runtime"]; !ok {
		Packages["runtime"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["runtime"]; !ok {
		PackageTypes["runtime"] = make(map[string]interface{})
	}
	var frame runtime.Frame
	var frames runtime.Frames
	Packages["runtime"]["CallersFrames"] = runtime.CallersFrames
	Packages["runtime"]["KeepAlive"] = runtime.KeepAlive
	Packages["runtime"]["SetCgoTraceback"] = runtime.SetCgoTraceback
	PackageTypes["runtime"]["Frame"] = reflect.TypeOf(&frame).Elem()
	PackageTypes["runtime"]["Frames"] = reflect.TypeOf(&frames).Elem()
}
