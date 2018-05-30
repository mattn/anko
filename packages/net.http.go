// +build !appengine

package packages

import (
	"net/http"
)

func init() {
	Packages["net/http"] = map[string]interface{}{
		"DefaultClient":     http.DefaultClient,
		"DefaultServeMux":   http.DefaultServeMux,
		"DefaultTransport":  http.DefaultTransport,
		"Handle":            http.Handle,
		"HandleFunc":        http.HandleFunc,
		"ListenAndServe":    http.ListenAndServe,
		"ListenAndServeTLS": http.ListenAndServeTLS,
		"NewRequest":        http.NewRequest,
		"NewServeMux":       http.NewServeMux,
		"Serve":             http.Serve,
		"SetCookie":         http.SetCookie,
	}
	PackageTypes["net/http"] = map[string]interface{}{
		"Client":   http.Client{},
		"Cookie":   http.Cookie{},
		"Request":  http.Request{},
		"Response": http.Response{},
	}
}
