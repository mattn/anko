// +build go1.6

package packages

import (
	"reflect"
	"sort"
)

// SortFuncsStruct provides functions to be used with Sort
type SortFuncsStruct struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (s SortFuncsStruct) Len() int           { return s.LenFunc() }
func (s SortFuncsStruct) Less(i, j int) bool { return s.LessFunc(i, j) }
func (s SortFuncsStruct) Swap(i, j int)      { s.SwapFunc(i, j) }
func init() {
	if _, ok := Packages["sort"]; !ok {
		Packages["sort"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["sort"]; !ok {
		PackageTypes["sort"] = make(map[string]interface{})
	}
	var float64slice sort.Float64Slice
	var intslice sort.IntSlice
	var interface_ sort.Interface
	var stringslice sort.StringSlice
	Packages["sort"]["Float64s"] = sort.Float64s
	Packages["sort"]["Float64sAreSorted"] = sort.Float64sAreSorted
	Packages["sort"]["Ints"] = sort.Ints
	Packages["sort"]["IntsAreSorted"] = sort.IntsAreSorted
	Packages["sort"]["IsSorted"] = sort.IsSorted
	Packages["sort"]["Reverse"] = sort.Reverse
	Packages["sort"]["Search"] = sort.Search
	Packages["sort"]["SearchFloat64s"] = sort.SearchFloat64s
	Packages["sort"]["SearchInts"] = sort.SearchInts
	Packages["sort"]["SearchStrings"] = sort.SearchStrings
	Packages["sort"]["Sort"] = sort.Sort
	Packages["sort"]["Stable"] = sort.Stable
	Packages["sort"]["Strings"] = sort.Strings
	Packages["sort"]["StringsAreSorted"] = sort.StringsAreSorted
	PackageTypes["sort"]["Float64Slice"] = reflect.TypeOf(&float64slice).Elem()
	PackageTypes["sort"]["IntSlice"] = reflect.TypeOf(&intslice).Elem()
	PackageTypes["sort"]["Interface"] = reflect.TypeOf(&interface_).Elem()
	PackageTypes["sort"]["StringSlice"] = reflect.TypeOf(&stringslice).Elem()
	PackageTypes["sort"]["SortFuncsStruct"] = &SortFuncsStruct{}
}
