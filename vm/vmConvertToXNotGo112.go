// +build !go1.12

package vm

import (
	"context"
	"fmt"
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
	// Before Go 1.12 the only way to do this is to get all the keys.
	// Note this is costly for large maps.
	mapKeys := rv.MapKeys()
	for i := 0; i < len(mapKeys); i++ {
		newKey, err := convertReflectValueToType(mapKeys[i], rtKey)
		if err != nil {
			return rv, err
		}
		value := rv.MapIndex(mapKeys[i])
		value, err = convertReflectValueToType(value, rtElem)
		if err != nil {
			return rv, err
		}
		newMap.SetMapIndex(newKey, value)
	}

	return newMap, nil
}
