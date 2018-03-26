// +build go1.6

package packages

import (
	"io/ioutil"
)

func init() {
	if _, ok := Packages["io/ioutil"]; !ok {
		Packages["io/ioutil"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["io/ioutil"]; !ok {
		PackageTypes["io/ioutil"] = make(map[string]interface{})
	}
	Packages["io/ioutil"]["Discard"] = ioutil.Discard
	Packages["io/ioutil"]["NopCloser"] = ioutil.NopCloser
	Packages["io/ioutil"]["ReadAll"] = ioutil.ReadAll
	Packages["io/ioutil"]["ReadDir"] = ioutil.ReadDir
	Packages["io/ioutil"]["ReadFile"] = ioutil.ReadFile
	Packages["io/ioutil"]["TempDir"] = ioutil.TempDir
	Packages["io/ioutil"]["TempFile"] = ioutil.TempFile
	Packages["io/ioutil"]["WriteFile"] = ioutil.WriteFile
}
