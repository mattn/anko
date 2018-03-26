package main

import (
	"fmt"
	"strings"
)

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

func goVersionToBuildTagName(s string) string {
	s = goVersionToFileExtName(s)
	return fmt.Sprintf("%s.%s", s[:1], s[1:])
}

func goVersionToFileExtName(s string) string {
	return strings.Replace(strings.ToLower(s), "go", "", -1)
}

func genFileName(pkg, ext string) string {
	f := "packages/" + strings.Join(strings.Split(pkg, "/"), ".")
	if ext == "" {
		return f + ".go"
	} else {
		return f + "." + ext + ".go"
	}
}

func diffSlice(notTypes1, typeAssigns1 []pkgFuncType, typeVars1 []string, notTypes2, typeAssigns2 []pkgFuncType, typeVars2 []string) ([]pkgFuncType, []pkgFuncType, []string) {
	var diffPkgFuncTypeSlice = func(a, b []pkgFuncType) []pkgFuncType {
		var m = make(map[pkgFuncType]bool)
		for _, v := range b {
			m[v] = true
		}

		var d []pkgFuncType
		for _, v := range a {
			if !m[v] {
				d = append(d, v)
			}
		}
		return d
	}
	var diffStringSlice = func(a, b []string) []string {
		var m = make(map[string]bool)
		for _, v := range b {
			m[v] = true
		}

		var d []string
		for _, v := range a {
			if !m[v] {
				d = append(d, v)
			}
		}
		return d
	}

	notTypesDiff := diffPkgFuncTypeSlice(notTypes1, notTypes2)
	typeAssignsDiff := diffPkgFuncTypeSlice(typeAssigns1, typeAssigns2)
	typeVarsDiff := diffStringSlice(typeVars1, typeVars2)
	return notTypesDiff, typeAssignsDiff, typeVarsDiff
}
