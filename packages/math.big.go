// +build go1.6

package packages

import (
	"math/big"
	"reflect"
)

func init() {
	if _, ok := Packages["math/big"]; !ok {
		Packages["math/big"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math/big"]; !ok {
		PackageTypes["math/big"] = make(map[string]interface{})
	}
	var accuracy big.Accuracy
	var errnan big.ErrNaN
	var float big.Float
	var int_ big.Int
	var rat big.Rat
	var roundingmode big.RoundingMode
	var word big.Word
	Packages["math/big"]["Above"] = big.Above
	Packages["math/big"]["AwayFromZero"] = big.AwayFromZero
	Packages["math/big"]["Below"] = big.Below
	Packages["math/big"]["Exact"] = big.Exact
	Packages["math/big"]["Jacobi"] = big.Jacobi
	Packages["math/big"]["MaxBase"] = big.MaxBase
	Packages["math/big"]["MaxExp"] = big.MaxExp
	Packages["math/big"]["MaxPrec"] = big.MaxPrec
	Packages["math/big"]["MinExp"] = big.MinExp
	Packages["math/big"]["NewFloat"] = big.NewFloat
	Packages["math/big"]["NewInt"] = big.NewInt
	Packages["math/big"]["NewRat"] = big.NewRat
	Packages["math/big"]["ParseFloat"] = big.ParseFloat
	Packages["math/big"]["ToNearestAway"] = big.ToNearestAway
	Packages["math/big"]["ToNearestEven"] = big.ToNearestEven
	Packages["math/big"]["ToNegativeInf"] = big.ToNegativeInf
	Packages["math/big"]["ToPositiveInf"] = big.ToPositiveInf
	Packages["math/big"]["ToZero"] = big.ToZero
	PackageTypes["math/big"]["Accuracy"] = reflect.TypeOf(&accuracy).Elem()
	PackageTypes["math/big"]["ErrNaN"] = reflect.TypeOf(&errnan).Elem()
	PackageTypes["math/big"]["Float"] = reflect.TypeOf(&float).Elem()
	PackageTypes["math/big"]["Int"] = reflect.TypeOf(&int_).Elem()
	PackageTypes["math/big"]["Rat"] = reflect.TypeOf(&rat).Elem()
	PackageTypes["math/big"]["RoundingMode"] = reflect.TypeOf(&roundingmode).Elem()
	PackageTypes["math/big"]["Word"] = reflect.TypeOf(&word).Elem()
}
