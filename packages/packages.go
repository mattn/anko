package packages

import (
	"fmt"

	"github.com/mattn/anko/vm"
)

var (
	// Packages is a where all the packages are stored so they can be imported when wanted
	Packages = make(map[string]map[string]interface{}, 16)
	// PackageTypes is a where all the package types are stored so they can be imported when wanted
	PackageTypes = make(map[string]map[string]interface{}, 4)
)

func init() {
	notAppEngine()
}

// DefineImport defines the vm import command that will import packages and package types when wanted
func DefineImport(e *vm.Env) {
	e.Define("import", func(source string) *vm.Env {
		methods, ok := Packages[source]
		if !ok {
			panic(fmt.Sprintf("package '%s' not found", source))
		}
		var err error
		pack := e.NewPackage(source)
		for methodName, methodValue := range methods {
			err = pack.Define(methodName, methodValue)
			if err != nil {
				panic(fmt.Sprintf("import Define error: %v", err))
			}
		}
		types, ok := PackageTypes[source]
		if ok {
			for typeName, typeValue := range types {
				err = pack.DefineType(typeName, typeValue)
				if err != nil {
					panic(fmt.Sprintf("import DefineType error: %v", err))
				}
			}
		}
		return pack
	})
}
