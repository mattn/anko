// +build go1.7

package packages

import (
	"bytes"
)

func init() {
	if _, ok := Packages["bytes"]; !ok {
		Packages["bytes"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["bytes"]; !ok {
		PackageTypes["bytes"] = make(map[string]interface{})
	}
	Packages["bytes"]["ContainsAny"] = bytes.ContainsAny
	Packages["bytes"]["ContainsRune"] = bytes.ContainsRune
}
