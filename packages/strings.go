// +build go1.6

package packages

import (
	"reflect"
	"strings"
)

func init() {
	if _, ok := Packages["strings"]; !ok {
		Packages["strings"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["strings"]; !ok {
		PackageTypes["strings"] = make(map[string]interface{})
	}
	var reader strings.Reader
	var replacer strings.Replacer
	Packages["strings"]["Compare"] = strings.Compare
	Packages["strings"]["Contains"] = strings.Contains
	Packages["strings"]["ContainsAny"] = strings.ContainsAny
	Packages["strings"]["ContainsRune"] = strings.ContainsRune
	Packages["strings"]["Count"] = strings.Count
	Packages["strings"]["EqualFold"] = strings.EqualFold
	Packages["strings"]["Fields"] = strings.Fields
	Packages["strings"]["FieldsFunc"] = strings.FieldsFunc
	Packages["strings"]["HasPrefix"] = strings.HasPrefix
	Packages["strings"]["HasSuffix"] = strings.HasSuffix
	Packages["strings"]["Index"] = strings.Index
	Packages["strings"]["IndexAny"] = strings.IndexAny
	Packages["strings"]["IndexByte"] = strings.IndexByte
	Packages["strings"]["IndexFunc"] = strings.IndexFunc
	Packages["strings"]["IndexRune"] = strings.IndexRune
	Packages["strings"]["Join"] = strings.Join
	Packages["strings"]["LastIndex"] = strings.LastIndex
	Packages["strings"]["LastIndexAny"] = strings.LastIndexAny
	Packages["strings"]["LastIndexByte"] = strings.LastIndexByte
	Packages["strings"]["LastIndexFunc"] = strings.LastIndexFunc
	Packages["strings"]["Map"] = strings.Map
	Packages["strings"]["NewReader"] = strings.NewReader
	Packages["strings"]["NewReplacer"] = strings.NewReplacer
	Packages["strings"]["Repeat"] = strings.Repeat
	Packages["strings"]["Replace"] = strings.Replace
	Packages["strings"]["Split"] = strings.Split
	Packages["strings"]["SplitAfter"] = strings.SplitAfter
	Packages["strings"]["SplitAfterN"] = strings.SplitAfterN
	Packages["strings"]["SplitN"] = strings.SplitN
	Packages["strings"]["Title"] = strings.Title
	Packages["strings"]["ToLower"] = strings.ToLower
	Packages["strings"]["ToLowerSpecial"] = strings.ToLowerSpecial
	Packages["strings"]["ToTitle"] = strings.ToTitle
	Packages["strings"]["ToTitleSpecial"] = strings.ToTitleSpecial
	Packages["strings"]["ToUpper"] = strings.ToUpper
	Packages["strings"]["ToUpperSpecial"] = strings.ToUpperSpecial
	Packages["strings"]["Trim"] = strings.Trim
	Packages["strings"]["TrimFunc"] = strings.TrimFunc
	Packages["strings"]["TrimLeft"] = strings.TrimLeft
	Packages["strings"]["TrimLeftFunc"] = strings.TrimLeftFunc
	Packages["strings"]["TrimPrefix"] = strings.TrimPrefix
	Packages["strings"]["TrimRight"] = strings.TrimRight
	Packages["strings"]["TrimRightFunc"] = strings.TrimRightFunc
	Packages["strings"]["TrimSpace"] = strings.TrimSpace
	Packages["strings"]["TrimSuffix"] = strings.TrimSuffix
	PackageTypes["strings"]["Reader"] = reflect.TypeOf(&reader).Elem()
	PackageTypes["strings"]["Replacer"] = reflect.TypeOf(&replacer).Elem()
}
