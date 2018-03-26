// +build go1.9

package packages

import (
	"encoding/json"
)

func init() {
	if _, ok := Packages["encoding/json"]; !ok {
		Packages["encoding/json"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["encoding/json"]; !ok {
		PackageTypes["encoding/json"] = make(map[string]interface{})
	}
	Packages["encoding/json"]["Valid"] = json.Valid
}
