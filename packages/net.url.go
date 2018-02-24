// +build !appengine

package packages

import (
	"net/url"
)

func init() {
	Packages["net/url"] = map[string]interface{}{
		"PathEscape":      url.PathEscape,
		"PathUnescape":    url.PathUnescape,
		"QueryEscape":     url.QueryEscape,
		"QueryUnescape":   url.QueryUnescape,
		"Parse":           url.Parse,
		"ParseRequestURI": url.ParseRequestURI,
		"User":            url.User,
		"UserPassword":    url.UserPassword,
		"ParseQuery":      url.ParseQuery,
	}
	PackageTypes["net/url"] = map[string]interface{}{
		"Error":  url.Error{},
		"URL":    url.URL{},
		"Values": url.Values{},
	}
}
