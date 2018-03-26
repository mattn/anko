// +build go1.6

package packages

import (
	"reflect"
	"strconv"
)

func init() {
	if _, ok := Packages["strconv"]; !ok {
		Packages["strconv"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["strconv"]; !ok {
		PackageTypes["strconv"] = make(map[string]interface{})
	}
	var numerror strconv.NumError
	Packages["strconv"]["AppendBool"] = strconv.AppendBool
	Packages["strconv"]["AppendFloat"] = strconv.AppendFloat
	Packages["strconv"]["AppendInt"] = strconv.AppendInt
	Packages["strconv"]["AppendQuote"] = strconv.AppendQuote
	Packages["strconv"]["AppendQuoteRune"] = strconv.AppendQuoteRune
	Packages["strconv"]["AppendQuoteRuneToASCII"] = strconv.AppendQuoteRuneToASCII
	Packages["strconv"]["AppendQuoteRuneToGraphic"] = strconv.AppendQuoteRuneToGraphic
	Packages["strconv"]["AppendQuoteToASCII"] = strconv.AppendQuoteToASCII
	Packages["strconv"]["AppendQuoteToGraphic"] = strconv.AppendQuoteToGraphic
	Packages["strconv"]["AppendUint"] = strconv.AppendUint
	Packages["strconv"]["Atoi"] = strconv.Atoi
	Packages["strconv"]["CanBackquote"] = strconv.CanBackquote
	Packages["strconv"]["ErrRange"] = strconv.ErrRange
	Packages["strconv"]["ErrSyntax"] = strconv.ErrSyntax
	Packages["strconv"]["FormatBool"] = strconv.FormatBool
	Packages["strconv"]["FormatFloat"] = strconv.FormatFloat
	Packages["strconv"]["FormatInt"] = strconv.FormatInt
	Packages["strconv"]["FormatUint"] = strconv.FormatUint
	Packages["strconv"]["IntSize"] = strconv.IntSize
	Packages["strconv"]["IsGraphic"] = strconv.IsGraphic
	Packages["strconv"]["IsPrint"] = strconv.IsPrint
	Packages["strconv"]["Itoa"] = strconv.Itoa
	Packages["strconv"]["ParseBool"] = strconv.ParseBool
	Packages["strconv"]["ParseFloat"] = strconv.ParseFloat
	Packages["strconv"]["ParseInt"] = strconv.ParseInt
	Packages["strconv"]["ParseUint"] = strconv.ParseUint
	Packages["strconv"]["Quote"] = strconv.Quote
	Packages["strconv"]["QuoteRune"] = strconv.QuoteRune
	Packages["strconv"]["QuoteRuneToASCII"] = strconv.QuoteRuneToASCII
	Packages["strconv"]["QuoteRuneToGraphic"] = strconv.QuoteRuneToGraphic
	Packages["strconv"]["QuoteToASCII"] = strconv.QuoteToASCII
	Packages["strconv"]["QuoteToGraphic"] = strconv.QuoteToGraphic
	Packages["strconv"]["Unquote"] = strconv.Unquote
	Packages["strconv"]["UnquoteChar"] = strconv.UnquoteChar
	PackageTypes["strconv"]["NumError"] = reflect.TypeOf(&numerror).Elem()
}
