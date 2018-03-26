// +build go1.8

package packages

import (
	"net/http"
	"reflect"
)

func init() {
	if _, ok := Packages["net/http"]; !ok {
		Packages["net/http"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/http"]; !ok {
		PackageTypes["net/http"] = make(map[string]interface{})
	}
	var pushoptions http.PushOptions
	var pusher http.Pusher
	Packages["net/http"]["ErrAbortHandler"] = http.ErrAbortHandler
	Packages["net/http"]["ErrServerClosed"] = http.ErrServerClosed
	Packages["net/http"]["NoBody"] = http.NoBody
	Packages["net/http"]["TrailerPrefix"] = http.TrailerPrefix
	PackageTypes["net/http"]["PushOptions"] = reflect.TypeOf(&pushoptions).Elem()
	PackageTypes["net/http"]["Pusher"] = reflect.TypeOf(&pusher).Elem()
}
