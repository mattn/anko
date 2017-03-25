// +build go1.8

package sort

import (
	"reflect"
	s "sort"

	"github.com/mattn/anko/vm"
)

func handleGo18(m *vm.Env) {
	m.Define("Slice", func(arr interface{}, less func(i, j int) reflect.Value) interface{} {
		s.Slice(arr, func(ii, jj int) bool {
			return less(ii, jj).Interface().(bool)
		})
		return arr
	})
}
