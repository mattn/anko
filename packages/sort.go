package packages

import (
	"sort"
)

func init() {
	Packages["sort"] = map[string]interface{}{
		"Float64s":          sort.Float64s,
		"Float64sAreSorted": sort.Float64sAreSorted,
		"Ints":              sort.Ints,
		"IntsAreSorted":     sort.IntsAreSorted,
		"IsSorted":          sort.IsSorted,
		"Search":            sort.Search,
		"SearchFloat64s":    sort.SearchFloat64s,
		"SearchInts":        sort.SearchInts,
		"SearchStrings":     sort.SearchStrings,
		"Sort":              sort.Sort,
		"Stable":            sort.Stable,
		"Strings":           sort.Strings,
		"StringsAreSorted":  sort.StringsAreSorted,
	}
	PackageTypes["sort"] = map[string]interface{}{
		"Float64Slice": sort.Float64Slice{},
		"IntSlice":     sort.IntSlice{},
		"StringSlice":  sort.StringSlice{},
	}
	sortGo18()
}
