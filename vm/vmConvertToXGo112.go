// +build go1.12

package vm

import (
	"reflect"
)

// convertMap trys to covert the reflect.Value map to the map reflect.Type
func convertMap(rv reflect.Value, rt reflect.Type) (reflect.Value, error) {
	rtKey := rt.Key()
	rtElem := rt.Elem()

	// create new map
	// note creating slice as work around to create map
	// just doing MakeMap can give incorrect type for defined types
	newMap := reflect.MakeSlice(reflect.SliceOf(rt), 0, 1)
	newMap = reflect.Append(newMap, reflect.MakeMap(reflect.MapOf(rtKey, rtElem))).Index(0)

	// copy keys to new map
	// For Go 1.12 and after can use MapRange
	mapIter := rv.MapRange()
	var value reflect.Value
	for mapIter.Next() {
		newKey, err := convertReflectValueToType(mapIter.Key(), rtKey)
		if err != nil {
			return rv, err
		}
		value, err = convertReflectValueToType(mapIter.Value(), rtElem)
		if err != nil {
			return rv, err
		}
		newMap.SetMapIndex(newKey, value)
	}

	return newMap, nil
}
