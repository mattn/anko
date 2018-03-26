// +build go1.8

package packages

import (
	"sort"
)

func init() {
	if _, ok := Packages["sort"]; !ok {
		Packages["sort"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["sort"]; !ok {
		PackageTypes["sort"] = make(map[string]interface{})
	}
	Packages["sort"]["Slice"] = sort.Slice
	Packages["sort"]["SliceIsSorted"] = sort.SliceIsSorted
	Packages["sort"]["SliceStable"] = sort.SliceStable
	PackageTypes["sort"]["SortFuncsStruct"] = &SortFuncsStruct{}
}
