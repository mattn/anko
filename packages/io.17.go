// +build go1.7

package packages

import (
	"io"
)

func init() {
	if _, ok := Packages["io"]; !ok {
		Packages["io"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["io"]; !ok {
		PackageTypes["io"] = make(map[string]interface{})
	}
	Packages["io"]["SeekCurrent"] = io.SeekCurrent
	Packages["io"]["SeekEnd"] = io.SeekEnd
	Packages["io"]["SeekStart"] = io.SeekStart
}
