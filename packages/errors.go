// +build go1.6

package packages

import (
	"errors"
)

func init() {
	if _, ok := Packages["errors"]; !ok {
		Packages["errors"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["errors"]; !ok {
		PackageTypes["errors"] = make(map[string]interface{})
	}
	Packages["errors"]["New"] = errors.New
}
