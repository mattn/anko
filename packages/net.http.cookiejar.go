// +build go1.6

package packages

import (
	"net/http/cookiejar"
	"reflect"
)

func init() {
	if _, ok := Packages["net/http/cookiejar"]; !ok {
		Packages["net/http/cookiejar"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/http/cookiejar"]; !ok {
		PackageTypes["net/http/cookiejar"] = make(map[string]interface{})
	}
	var jar cookiejar.Jar
	var options cookiejar.Options
	var publicsuffixlist cookiejar.PublicSuffixList
	Packages["net/http/cookiejar"]["New"] = cookiejar.New
	PackageTypes["net/http/cookiejar"]["Jar"] = reflect.TypeOf(&jar).Elem()
	PackageTypes["net/http/cookiejar"]["Options"] = reflect.TypeOf(&options).Elem()
	PackageTypes["net/http/cookiejar"]["PublicSuffixList"] = reflect.TypeOf(&publicsuffixlist).Elem()
}
