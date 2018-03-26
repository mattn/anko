// +build go1.10

package packages

import (
	"reflect"
	"strings"
)

func init() {
	if _, ok := Packages["strings"]; !ok {
		Packages["strings"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["strings"]; !ok {
		PackageTypes["strings"] = make(map[string]interface{})
	}
	var builder strings.Builder
	PackageTypes["strings"]["Builder"] = reflect.TypeOf(&builder).Elem()
}
