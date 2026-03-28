package vm

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	_ "github.com/mattn/anko/packages"
)

func TestImportPackages(t *testing.T) {
	type Descriptor struct {
		ID   int
		Name string
		Desc string
	}
	var hello = func() {
		fmt.Println("Msg: Hello, from 'custom' Package")
	}
	var sum = func(args ...float64) float64 {
		s := 0.0
		for _, v := range args {
			s += v
		}
		return s
	}

	var pkg = map[string]map[string]reflect.Value{
		"custom": {
			"hello":  reflect.ValueOf(hello),
			"sum":    reflect.ValueOf(sum),
			"printf": reflect.ValueOf(fmt.Printf),
		},
	}
	var typ = map[string]map[string]reflect.Type{
		"custom": {
			"Descriptor": reflect.TypeOf(Descriptor{}),
		},
	}

	// Tests inputs
	tNoPkg := Test{Script: `a = 1; b = 2; c = a+b;`, RunOutput: int64(3)}
	tStdFmt := Test{Script: `var fmt = import("fmt"); fmt.Println("Msg: Hello from 'stdlib.fmt'"); x=0`, RunOutput: int64(0)}
	tCust := Test{Script: `var custom = import("custom"); custom.hello(); s = custom.sum(1, 2, 3, 11); custom.printf("custom: sum = %f\n", s); out=s`, RunOutput: 17.0}
	tCustFmt := Test{Script: `var custom, fmt = import("custom"), import("fmt"); custom.hello(); s = custom.sum(1, 2, 3, 11); fmt.Printf("stdlib.fmt: sum = %f\n", s); out=s`, RunOutput: 17.0}
	tCustTimeErr := Test{Script: `var custom, time = import("custom"), import("time"); custom.hello(); now = time.Now(); custom.printf("Now %v\n", now); out=0`,
		RunError: errors.New("package not found: time")}

	// --- importer not specified (backward compatibility) ---
	// should not error
	runTests(t, []Test{tNoPkg, tStdFmt}, nil, nil)

	// option without specifying package importer
	runTests(t, []Test{tNoPkg, tStdFmt}, nil, &Options{Debug: true})

	// --- explicitly specify importer --- //
	// import only custom package
	optCustom := Options{
		PkgImporter: NewPackagesImporter(pkg, typ),
	}
	runTests(t, []Test{tNoPkg, tCust}, nil, &optCustom)

	// import custom package and ALL stdlib
	optCustomStdAll := Options{
		PkgImporter: NewPackagesWithStdImporter(pkg, typ),
	}
	runTests(t, []Test{tNoPkg, tStdFmt, tCust, tCustFmt}, nil, &optCustomStdAll)

	// import custom package and `fmt` from stdlib
	optCustomStdFmt := Options{
		PkgImporter: NewPackagesWithStdImporter(pkg, typ, "fmt"),
	}
	runTests(t, []Test{tNoPkg, tStdFmt, tCust, tCustFmt, tCustTimeErr}, nil, &optCustomStdFmt)

}
