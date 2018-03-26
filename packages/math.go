// +build go1.6

package packages

import (
	"math"
)

func init() {
	if _, ok := Packages["math"]; !ok {
		Packages["math"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["math"]; !ok {
		PackageTypes["math"] = make(map[string]interface{})
	}
	Packages["math"]["Abs"] = math.Abs
	Packages["math"]["Acos"] = math.Acos
	Packages["math"]["Acosh"] = math.Acosh
	Packages["math"]["Asin"] = math.Asin
	Packages["math"]["Asinh"] = math.Asinh
	Packages["math"]["Atan"] = math.Atan
	Packages["math"]["Atan2"] = math.Atan2
	Packages["math"]["Atanh"] = math.Atanh
	Packages["math"]["Cbrt"] = math.Cbrt
	Packages["math"]["Ceil"] = math.Ceil
	Packages["math"]["Copysign"] = math.Copysign
	Packages["math"]["Cos"] = math.Cos
	Packages["math"]["Cosh"] = math.Cosh
	Packages["math"]["Dim"] = math.Dim
	Packages["math"]["E"] = math.E
	Packages["math"]["Erf"] = math.Erf
	Packages["math"]["Erfc"] = math.Erfc
	Packages["math"]["Exp"] = math.Exp
	Packages["math"]["Exp2"] = math.Exp2
	Packages["math"]["Expm1"] = math.Expm1
	Packages["math"]["Float32bits"] = math.Float32bits
	Packages["math"]["Float32frombits"] = math.Float32frombits
	Packages["math"]["Float64bits"] = math.Float64bits
	Packages["math"]["Float64frombits"] = math.Float64frombits
	Packages["math"]["Floor"] = math.Floor
	Packages["math"]["Frexp"] = math.Frexp
	Packages["math"]["Gamma"] = math.Gamma
	Packages["math"]["Hypot"] = math.Hypot
	Packages["math"]["Ilogb"] = math.Ilogb
	Packages["math"]["Inf"] = math.Inf
	Packages["math"]["IsInf"] = math.IsInf
	Packages["math"]["IsNaN"] = math.IsNaN
	Packages["math"]["J0"] = math.J0
	Packages["math"]["J1"] = math.J1
	Packages["math"]["Jn"] = math.Jn
	Packages["math"]["Ldexp"] = math.Ldexp
	Packages["math"]["Lgamma"] = math.Lgamma
	Packages["math"]["Ln10"] = math.Ln10
	Packages["math"]["Ln2"] = math.Ln2
	Packages["math"]["Log"] = math.Log
	Packages["math"]["Log10"] = math.Log10
	Packages["math"]["Log10E"] = math.Log10E
	Packages["math"]["Log1p"] = math.Log1p
	Packages["math"]["Log2"] = math.Log2
	Packages["math"]["Log2E"] = math.Log2E
	Packages["math"]["Logb"] = math.Logb
	Packages["math"]["Max"] = math.Max
	Packages["math"]["MaxFloat32"] = math.MaxFloat32
	Packages["math"]["MaxFloat64"] = math.MaxFloat64
	Packages["math"]["MaxInt16"] = math.MaxInt16
	Packages["math"]["MaxInt32"] = math.MaxInt32
	Packages["math"]["MaxInt64"] = math.MaxInt64
	Packages["math"]["MaxInt8"] = math.MaxInt8
	Packages["math"]["MaxUint16"] = math.MaxUint16
	Packages["math"]["MaxUint32"] = math.MaxUint32
	Packages["math"]["MaxUint64"] = uint64(math.MaxUint64)
	Packages["math"]["MaxUint8"] = math.MaxUint8
	Packages["math"]["Min"] = math.Min
	Packages["math"]["MinInt16"] = math.MinInt16
	Packages["math"]["MinInt32"] = math.MinInt32
	Packages["math"]["MinInt64"] = math.MinInt64
	Packages["math"]["MinInt8"] = math.MinInt8
	Packages["math"]["Mod"] = math.Mod
	Packages["math"]["Modf"] = math.Modf
	Packages["math"]["NaN"] = math.NaN
	Packages["math"]["Nextafter"] = math.Nextafter
	Packages["math"]["Nextafter32"] = math.Nextafter32
	Packages["math"]["Phi"] = math.Phi
	Packages["math"]["Pi"] = math.Pi
	Packages["math"]["Pow"] = math.Pow
	Packages["math"]["Pow10"] = math.Pow10
	Packages["math"]["Remainder"] = math.Remainder
	Packages["math"]["Signbit"] = math.Signbit
	Packages["math"]["Sin"] = math.Sin
	Packages["math"]["Sincos"] = math.Sincos
	Packages["math"]["Sinh"] = math.Sinh
	Packages["math"]["SmallestNonzeroFloat32"] = math.SmallestNonzeroFloat32
	Packages["math"]["SmallestNonzeroFloat64"] = math.SmallestNonzeroFloat64
	Packages["math"]["Sqrt"] = math.Sqrt
	Packages["math"]["Sqrt2"] = math.Sqrt2
	Packages["math"]["SqrtE"] = math.SqrtE
	Packages["math"]["SqrtPhi"] = math.SqrtPhi
	Packages["math"]["SqrtPi"] = math.SqrtPi
	Packages["math"]["Tan"] = math.Tan
	Packages["math"]["Tanh"] = math.Tanh
	Packages["math"]["Trunc"] = math.Trunc
	Packages["math"]["Y0"] = math.Y0
	Packages["math"]["Y1"] = math.Y1
	Packages["math"]["Yn"] = math.Yn
}
