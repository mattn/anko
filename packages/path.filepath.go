// +build go1.6

package packages

import (
	"path/filepath"
	"reflect"
)

func init() {
	if _, ok := Packages["path/filepath"]; !ok {
		Packages["path/filepath"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["path/filepath"]; !ok {
		PackageTypes["path/filepath"] = make(map[string]interface{})
	}
	var walkfunc filepath.WalkFunc
	Packages["path/filepath"]["Abs"] = filepath.Abs
	Packages["path/filepath"]["Base"] = filepath.Base
	Packages["path/filepath"]["Clean"] = filepath.Clean
	Packages["path/filepath"]["Dir"] = filepath.Dir
	Packages["path/filepath"]["ErrBadPattern"] = filepath.ErrBadPattern
	Packages["path/filepath"]["EvalSymlinks"] = filepath.EvalSymlinks
	Packages["path/filepath"]["Ext"] = filepath.Ext
	Packages["path/filepath"]["FromSlash"] = filepath.FromSlash
	Packages["path/filepath"]["Glob"] = filepath.Glob
	Packages["path/filepath"]["HasPrefix"] = filepath.HasPrefix
	Packages["path/filepath"]["IsAbs"] = filepath.IsAbs
	Packages["path/filepath"]["Join"] = filepath.Join
	Packages["path/filepath"]["ListSeparator"] = filepath.ListSeparator
	Packages["path/filepath"]["Match"] = filepath.Match
	Packages["path/filepath"]["Rel"] = filepath.Rel
	Packages["path/filepath"]["Separator"] = filepath.Separator
	Packages["path/filepath"]["SkipDir"] = filepath.SkipDir
	Packages["path/filepath"]["Split"] = filepath.Split
	Packages["path/filepath"]["SplitList"] = filepath.SplitList
	Packages["path/filepath"]["ToSlash"] = filepath.ToSlash
	Packages["path/filepath"]["VolumeName"] = filepath.VolumeName
	Packages["path/filepath"]["Walk"] = filepath.Walk
	PackageTypes["path/filepath"]["WalkFunc"] = reflect.TypeOf(&walkfunc).Elem()
}
