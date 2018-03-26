// +build go1.6

package packages

import (
	"net/url"
	"reflect"
)

func init() {
	if _, ok := Packages["net/url"]; !ok {
		Packages["net/url"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/url"]; !ok {
		PackageTypes["net/url"] = make(map[string]interface{})
	}
	var error_ url.Error
	var escapeerror url.EscapeError
	var invalidhosterror url.InvalidHostError
	var url_ url.URL
	var userinfo url.Userinfo
	var values url.Values
	Packages["net/url"]["Parse"] = url.Parse
	Packages["net/url"]["ParseQuery"] = url.ParseQuery
	Packages["net/url"]["ParseRequestURI"] = url.ParseRequestURI
	Packages["net/url"]["QueryEscape"] = url.QueryEscape
	Packages["net/url"]["QueryUnescape"] = url.QueryUnescape
	Packages["net/url"]["User"] = url.User
	Packages["net/url"]["UserPassword"] = url.UserPassword
	PackageTypes["net/url"]["Error"] = reflect.TypeOf(&error_).Elem()
	PackageTypes["net/url"]["EscapeError"] = reflect.TypeOf(&escapeerror).Elem()
	PackageTypes["net/url"]["InvalidHostError"] = reflect.TypeOf(&invalidhosterror).Elem()
	PackageTypes["net/url"]["URL"] = reflect.TypeOf(&url_).Elem()
	PackageTypes["net/url"]["Userinfo"] = reflect.TypeOf(&userinfo).Elem()
	PackageTypes["net/url"]["Values"] = reflect.TypeOf(&values).Elem()
}
