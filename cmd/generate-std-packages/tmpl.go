package main

import (
	"bytes"
	"strings"
	"text/template"
)

var (
	MainTmpl = `// +build go{{.go_version}}

package packages

import (
	{{.import_packages}}
)

{{.extra}}

func init() {
	if _, ok := Packages["{{.name}}"]; !ok {
		Packages["{{.name}}"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["{{.name}}"]; !ok {
		PackageTypes["{{.name}}"] = make(map[string]interface{})
	}

	{{.package_no_struct_types}}

	{{range $index, $element := .package_funcs}}
		Packages["{{.PkgName}}"]["{{.FuncTypeName}}"] = {{.Expr}}
	{{end}}

	{{range $index, $element := .package_types}}
		PackageTypes["{{.PkgName}}"]["{{.FuncTypeName}}"] = {{.Expr}}
	{{end}}
}
`

	DiffTmpl = `// +build go{{.go_version}}

package packages

import (
	{{.import_packages}}
)

func init() {
	if _, ok := Packages["{{.name}}"]; !ok {
		Packages["{{.name}}"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["{{.name}}"]; !ok {
		PackageTypes["{{.name}}"] = make(map[string]interface{})
	}

	{{.package_no_struct_types}}

	{{range $index, $element := .package_funcs}}
		Packages["{{.PkgName}}"]["{{.FuncTypeName}}"] = {{.Expr}}
	{{end}}

	{{range $index, $element := .package_types}}
		PackageTypes["{{.PkgName}}"]["{{.FuncTypeName}}"] = {{.Expr}}
	{{end}}
}
`

	PackagesTmpl = `package packages

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
`
)

func hasPkgFuncTypeField(v interface{}, name string) bool {
	m, ok := v.(map[string]interface{})
	if !ok {
		return false
	}

	n, ok := m[name]
	if !ok {
		return false
	}

	p, ok := n.([]pkgFuncType)

	return ok && len(p) > 0
}

func parseTmpl(tmpl string, data map[string]interface{}) ([]byte, error) {
	for k, v := range data {
		if v == "" {
			tmpl = strings.Replace(tmpl, "{{."+k+"}}", "", -1)
		}
	}

	parsedTmpl, err := template.New("tmpl").Funcs(template.FuncMap{"hasPkgFuncTypeField": hasPkgFuncTypeField}).Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	s := strings.Split(result.String(), "\n")
	var s2 []string
	for i := 0; i < len(s); i++ {
		if s3 := strings.TrimSpace(s[i]); s3 == "" {
			continue
		}

		s2 = append(s2, s[i])
		if strings.HasPrefix(s[i], "// +build") {
			s2 = append(s2, "\n")
		}
	}

	return []byte(strings.Join(s2, "\n")), nil
}
