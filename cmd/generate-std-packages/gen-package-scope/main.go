package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/importer"
	"go/types"
	"os"
	"reflect"
	"strings"
)

var packageName string

func init() {
	flag.StringVar(&packageName, "package", "", "package name")
	flag.Parse()
}

type pkgFuncType struct {
	PkgName      string `json:"pkg_name"`
	FuncTypeName string `json:"func_type_name"`
	Expr         string `json:"expr"`
}

type output struct {
	NotTypes    []pkgFuncType `json:"not_types"`
	TypeAssigns []pkgFuncType `json:"type_assigns"`
	TypeVars    []string      `json:"type_vars"`
}

func getpackageTailName(packageImportName string) string {
	n := strings.Split(packageImportName, "/")
	return n[len(n)-1]
}

func getConvertedName(pkgImportName, name string) string {
	n := getpackageTailName(pkgImportName) + "." + name

	return n
}

var keyword = map[string]bool{
	"interface": true,
	"func":      true,
	"time":      true,
	"error":     true,
	"url":       true,
	"rand":      true,
	"flag":      true,
	"regexp":    true,
	"int":       true,
}

func dealPackageScope(pkgImportName string) ([]pkgFuncType, []pkgFuncType, []string, error) {
	pkg, err := importer.Default().Import(pkgImportName)
	if err != nil {
		return nil, nil, nil, err
	}

	var notTypes []pkgFuncType
	var typeAssigns []pkgFuncType
	var typeVars []string
	var scope = pkg.Scope()

	for _, name := range scope.Names() {

		sc := scope.Lookup(name)
		if sc.Exported() {
			switch reflect.TypeOf(sc) {
			case reflect.TypeOf(new(types.TypeName)):
				varName := strings.ToLower(name)
				if keyword[varName] {
					varName = varName + "_"
				}
				typeVars = append(typeVars, fmt.Sprintf(`        var %s %s.%s`, varName, getpackageTailName(pkgImportName), name))
				typeAssigns = append(typeAssigns, pkgFuncType{pkgImportName, name, fmt.Sprintf(`reflect.TypeOf(&%s).Elem()`, varName)})
			default:
				notTypes = append(notTypes, pkgFuncType{pkgImportName, name, getConvertedName(pkgImportName, name)})
			}
		}
	}

	return notTypes, typeAssigns, typeVars, nil
}

func main() {
	if packageName == "" {
		fmt.Fprint(os.Stderr, "package name is empty")
		return
	}
	notTypes, typeAssigns, typeVars, err := dealPackageScope(packageName)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	b, err := json.MarshalIndent(output{notTypes, typeAssigns, typeVars}, "", "  ")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fmt.Printf("%s", b)
}
