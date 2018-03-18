package packages

import (
	"math/rand"
)

func init() {
	Packages["math/rand"] = map[string]interface{}{
		"ExpFloat64":  rand.ExpFloat64,
		"Float32":     rand.Float32,
		"Float64":     rand.Float64,
		"Int":         rand.Int,
		"Int31":       rand.Int31,
		"Int31n":      rand.Int31n,
		"Int63":       rand.Int63,
		"Int63n":      rand.Int63n,
		"Intn":        rand.Intn,
		"NormFloat64": rand.NormFloat64,
		"Perm":        rand.Perm,
		"Seed":        rand.Seed,
		"Uint32":      rand.Uint32,
	}
}
