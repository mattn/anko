// +build go1.7

package packages

import (
	"os/exec"
)

func init() {
	if _, ok := Packages["os/exec"]; !ok {
		Packages["os/exec"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["os/exec"]; !ok {
		PackageTypes["os/exec"] = make(map[string]interface{})
	}
	Packages["os/exec"]["CommandContext"] = exec.CommandContext
}
