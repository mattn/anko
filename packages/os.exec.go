// +build go1.6

package packages

import (
	"os/exec"
	"reflect"
)

func init() {
	if _, ok := Packages["os/exec"]; !ok {
		Packages["os/exec"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["os/exec"]; !ok {
		PackageTypes["os/exec"] = make(map[string]interface{})
	}
	var cmd exec.Cmd
	var error_ exec.Error
	var exiterror exec.ExitError
	Packages["os/exec"]["Command"] = exec.Command
	Packages["os/exec"]["ErrNotFound"] = exec.ErrNotFound
	Packages["os/exec"]["LookPath"] = exec.LookPath
	PackageTypes["os/exec"]["Cmd"] = reflect.TypeOf(&cmd).Elem()
	PackageTypes["os/exec"]["Error"] = reflect.TypeOf(&error_).Elem()
	PackageTypes["os/exec"]["ExitError"] = reflect.TypeOf(&exiterror).Elem()
}
