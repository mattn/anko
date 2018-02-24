package packages

import (
	"fmt"

	"github.com/mattn/anko/vm"
)

var Packages = make(map[string]map[string]interface{}, 16)
var PackageTypes = make(map[string]map[string]interface{}, 4)

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
