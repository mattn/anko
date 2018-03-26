// +build go1.8

package packages

import (
	"net/url"
)

func init() {
	if _, ok := Packages["net/url"]; !ok {
		Packages["net/url"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/url"]; !ok {
		PackageTypes["net/url"] = make(map[string]interface{})
	}
	Packages["net/url"]["PathEscape"] = url.PathEscape
	Packages["net/url"]["PathUnescape"] = url.PathUnescape
}
