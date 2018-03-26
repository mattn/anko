// +build go1.6

package packages

import (
	"reflect"
	"regexp"
)

func init() {
	if _, ok := Packages["regexp"]; !ok {
		Packages["regexp"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["regexp"]; !ok {
		PackageTypes["regexp"] = make(map[string]interface{})
	}
	var regexp_ regexp.Regexp
	Packages["regexp"]["Compile"] = regexp.Compile
	Packages["regexp"]["CompilePOSIX"] = regexp.CompilePOSIX
	Packages["regexp"]["Match"] = regexp.Match
	Packages["regexp"]["MatchReader"] = regexp.MatchReader
	Packages["regexp"]["MatchString"] = regexp.MatchString
	Packages["regexp"]["MustCompile"] = regexp.MustCompile
	Packages["regexp"]["MustCompilePOSIX"] = regexp.MustCompilePOSIX
	Packages["regexp"]["QuoteMeta"] = regexp.QuoteMeta
	PackageTypes["regexp"]["Regexp"] = reflect.TypeOf(&regexp_).Elem()
}
