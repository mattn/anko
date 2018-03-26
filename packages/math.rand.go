// +build go1.6

package packages

import (
	"math/rand"
	"reflect"
)

func init() {
	if _, ok := Packages["math/rand"]; !ok {
		Packages["math/rand"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math/rand"]; !ok {
		PackageTypes["math/rand"] = make(map[string]interface{})
	}
	var rand_ rand.Rand
	var source rand.Source
	var zipf rand.Zipf
	Packages["math/rand"]["ExpFloat64"] = rand.ExpFloat64
	Packages["math/rand"]["Float32"] = rand.Float32
	Packages["math/rand"]["Float64"] = rand.Float64
	Packages["math/rand"]["Int"] = rand.Int
	Packages["math/rand"]["Int31"] = rand.Int31
	Packages["math/rand"]["Int31n"] = rand.Int31n
	Packages["math/rand"]["Int63"] = rand.Int63
	Packages["math/rand"]["Int63n"] = rand.Int63n
	Packages["math/rand"]["Intn"] = rand.Intn
	Packages["math/rand"]["New"] = rand.New
	Packages["math/rand"]["NewSource"] = rand.NewSource
	Packages["math/rand"]["NewZipf"] = rand.NewZipf
	Packages["math/rand"]["NormFloat64"] = rand.NormFloat64
	Packages["math/rand"]["Perm"] = rand.Perm
	Packages["math/rand"]["Read"] = rand.Read
	Packages["math/rand"]["Seed"] = rand.Seed
	Packages["math/rand"]["Uint32"] = rand.Uint32
	PackageTypes["math/rand"]["Rand"] = reflect.TypeOf(&rand_).Elem()
	PackageTypes["math/rand"]["Source"] = reflect.TypeOf(&source).Elem()
	PackageTypes["math/rand"]["Zipf"] = reflect.TypeOf(&zipf).Elem()
}
