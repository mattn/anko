// +build go1.8

package packages

import (
	"os"
)

func init() {
	if _, ok := Packages["os"]; !ok {
		Packages["os"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["os"]; !ok {
		PackageTypes["os"] = make(map[string]interface{})
	}
	Packages["os"]["ErrClosed"] = os.ErrClosed
	Packages["os"]["Executable"] = os.Executable
}
