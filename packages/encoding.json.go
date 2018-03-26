// +build go1.6

package packages

import (
	"encoding/json"
	"reflect"
)

func init() {
	if _, ok := Packages["encoding/json"]; !ok {
		Packages["encoding/json"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["encoding/json"]; !ok {
		PackageTypes["encoding/json"] = make(map[string]interface{})
	}
	var decoder json.Decoder
	var delim json.Delim
	var encoder json.Encoder
	var invalidutf8error json.InvalidUTF8Error
	var invalidunmarshalerror json.InvalidUnmarshalError
	var marshaler json.Marshaler
	var marshalererror json.MarshalerError
	var number json.Number
	var rawmessage json.RawMessage
	var syntaxerror json.SyntaxError
	var token json.Token
	var unmarshalfielderror json.UnmarshalFieldError
	var unmarshaltypeerror json.UnmarshalTypeError
	var unmarshaler json.Unmarshaler
	var unsupportedtypeerror json.UnsupportedTypeError
	var unsupportedvalueerror json.UnsupportedValueError
	Packages["encoding/json"]["Compact"] = json.Compact
	Packages["encoding/json"]["HTMLEscape"] = json.HTMLEscape
	Packages["encoding/json"]["Indent"] = json.Indent
	Packages["encoding/json"]["Marshal"] = json.Marshal
	Packages["encoding/json"]["MarshalIndent"] = json.MarshalIndent
	Packages["encoding/json"]["NewDecoder"] = json.NewDecoder
	Packages["encoding/json"]["NewEncoder"] = json.NewEncoder
	Packages["encoding/json"]["Unmarshal"] = json.Unmarshal
	PackageTypes["encoding/json"]["Decoder"] = reflect.TypeOf(&decoder).Elem()
	PackageTypes["encoding/json"]["Delim"] = reflect.TypeOf(&delim).Elem()
	PackageTypes["encoding/json"]["Encoder"] = reflect.TypeOf(&encoder).Elem()
	PackageTypes["encoding/json"]["InvalidUTF8Error"] = reflect.TypeOf(&invalidutf8error).Elem()
	PackageTypes["encoding/json"]["InvalidUnmarshalError"] = reflect.TypeOf(&invalidunmarshalerror).Elem()
	PackageTypes["encoding/json"]["Marshaler"] = reflect.TypeOf(&marshaler).Elem()
	PackageTypes["encoding/json"]["MarshalerError"] = reflect.TypeOf(&marshalererror).Elem()
	PackageTypes["encoding/json"]["Number"] = reflect.TypeOf(&number).Elem()
	PackageTypes["encoding/json"]["RawMessage"] = reflect.TypeOf(&rawmessage).Elem()
	PackageTypes["encoding/json"]["SyntaxError"] = reflect.TypeOf(&syntaxerror).Elem()
	PackageTypes["encoding/json"]["Token"] = reflect.TypeOf(&token).Elem()
	PackageTypes["encoding/json"]["UnmarshalFieldError"] = reflect.TypeOf(&unmarshalfielderror).Elem()
	PackageTypes["encoding/json"]["UnmarshalTypeError"] = reflect.TypeOf(&unmarshaltypeerror).Elem()
	PackageTypes["encoding/json"]["Unmarshaler"] = reflect.TypeOf(&unmarshaler).Elem()
	PackageTypes["encoding/json"]["UnsupportedTypeError"] = reflect.TypeOf(&unsupportedtypeerror).Elem()
	PackageTypes["encoding/json"]["UnsupportedValueError"] = reflect.TypeOf(&unsupportedvalueerror).Elem()
}
