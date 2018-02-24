// +build go1.8

package packages

import (
	"sort"
)

func sortGo18() {
	Packages["sort"]["Slice"] = sort.Slice
	Packages["sort"]["SliceIsSorted"] = sort.SliceIsSorted
	Packages["sort"]["SliceStable"] = sort.SliceStable
}
