// +build go1.9

package packages

import (
	"net/http"
)

func init() {
	if _, ok := Packages["net/http"]; !ok {
		Packages["net/http"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/http"]; !ok {
		PackageTypes["net/http"] = make(map[string]interface{})
	}
	Packages["net/http"]["ServeTLS"] = http.ServeTLS
}
