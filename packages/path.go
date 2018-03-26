// +build go1.6

package packages

import (
	"path"
)

func init() {
	if _, ok := Packages["path"]; !ok {
		Packages["path"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["path"]; !ok {
		PackageTypes["path"] = make(map[string]interface{})
	}
	Packages["path"]["Base"] = path.Base
	Packages["path"]["Clean"] = path.Clean
	Packages["path"]["Dir"] = path.Dir
	Packages["path"]["ErrBadPattern"] = path.ErrBadPattern
	Packages["path"]["Ext"] = path.Ext
	Packages["path"]["IsAbs"] = path.IsAbs
	Packages["path"]["Join"] = path.Join
	Packages["path"]["Match"] = path.Match
	Packages["path"]["Split"] = path.Split
}
