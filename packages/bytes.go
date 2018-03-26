// +build go1.6

package packages

import (
	"bytes"
	"reflect"
)

func init() {
	if _, ok := Packages["bytes"]; !ok {
		Packages["bytes"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["bytes"]; !ok {
		PackageTypes["bytes"] = make(map[string]interface{})
	}
	var buffer bytes.Buffer
	var reader bytes.Reader
	Packages["bytes"]["Compare"] = bytes.Compare
	Packages["bytes"]["Contains"] = bytes.Contains
	Packages["bytes"]["Count"] = bytes.Count
	Packages["bytes"]["Equal"] = bytes.Equal
	Packages["bytes"]["EqualFold"] = bytes.EqualFold
	Packages["bytes"]["ErrTooLarge"] = bytes.ErrTooLarge
	Packages["bytes"]["Fields"] = bytes.Fields
	Packages["bytes"]["FieldsFunc"] = bytes.FieldsFunc
	Packages["bytes"]["HasPrefix"] = bytes.HasPrefix
	Packages["bytes"]["HasSuffix"] = bytes.HasSuffix
	Packages["bytes"]["Index"] = bytes.Index
	Packages["bytes"]["IndexAny"] = bytes.IndexAny
	Packages["bytes"]["IndexByte"] = bytes.IndexByte
	Packages["bytes"]["IndexFunc"] = bytes.IndexFunc
	Packages["bytes"]["IndexRune"] = bytes.IndexRune
	Packages["bytes"]["Join"] = bytes.Join
	Packages["bytes"]["LastIndex"] = bytes.LastIndex
	Packages["bytes"]["LastIndexAny"] = bytes.LastIndexAny
	Packages["bytes"]["LastIndexByte"] = bytes.LastIndexByte
	Packages["bytes"]["LastIndexFunc"] = bytes.LastIndexFunc
	Packages["bytes"]["Map"] = bytes.Map
	Packages["bytes"]["MinRead"] = bytes.MinRead
	Packages["bytes"]["NewBuffer"] = bytes.NewBuffer
	Packages["bytes"]["NewBufferString"] = bytes.NewBufferString
	Packages["bytes"]["NewReader"] = bytes.NewReader
	Packages["bytes"]["Repeat"] = bytes.Repeat
	Packages["bytes"]["Replace"] = bytes.Replace
	Packages["bytes"]["Runes"] = bytes.Runes
	Packages["bytes"]["Split"] = bytes.Split
	Packages["bytes"]["SplitAfter"] = bytes.SplitAfter
	Packages["bytes"]["SplitAfterN"] = bytes.SplitAfterN
	Packages["bytes"]["SplitN"] = bytes.SplitN
	Packages["bytes"]["Title"] = bytes.Title
	Packages["bytes"]["ToLower"] = bytes.ToLower
	Packages["bytes"]["ToLowerSpecial"] = bytes.ToLowerSpecial
	Packages["bytes"]["ToTitle"] = bytes.ToTitle
	Packages["bytes"]["ToTitleSpecial"] = bytes.ToTitleSpecial
	Packages["bytes"]["ToUpper"] = bytes.ToUpper
	Packages["bytes"]["ToUpperSpecial"] = bytes.ToUpperSpecial
	Packages["bytes"]["Trim"] = bytes.Trim
	Packages["bytes"]["TrimFunc"] = bytes.TrimFunc
	Packages["bytes"]["TrimLeft"] = bytes.TrimLeft
	Packages["bytes"]["TrimLeftFunc"] = bytes.TrimLeftFunc
	Packages["bytes"]["TrimPrefix"] = bytes.TrimPrefix
	Packages["bytes"]["TrimRight"] = bytes.TrimRight
	Packages["bytes"]["TrimRightFunc"] = bytes.TrimRightFunc
	Packages["bytes"]["TrimSpace"] = bytes.TrimSpace
	Packages["bytes"]["TrimSuffix"] = bytes.TrimSuffix
	PackageTypes["bytes"]["Buffer"] = reflect.TypeOf(&buffer).Elem()
	PackageTypes["bytes"]["Reader"] = reflect.TypeOf(&reader).Elem()
}
