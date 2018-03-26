// +build go1.10

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
	Packages["os"]["ErrNoDeadline"] = os.ErrNoDeadline
	Packages["os"]["IsTimeout"] = os.IsTimeout
}
