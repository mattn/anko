// +build go1.8

package packages

import (
	"net"
	"reflect"
)

func init() {
	if _, ok := Packages["net"]; !ok {
		Packages["net"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net"]; !ok {
		PackageTypes["net"] = make(map[string]interface{})
	}
	var buffers net.Buffers
	var resolver net.Resolver
	Packages["net"]["DefaultResolver"] = net.DefaultResolver
	PackageTypes["net"]["Buffers"] = reflect.TypeOf(&buffers).Elem()
	PackageTypes["net"]["Resolver"] = reflect.TypeOf(&resolver).Elem()
}
