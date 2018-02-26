package packages

import (
	"bytes"
)

func init() {
	Packages["bytes"] = map[string]interface{}{
		"Compare":        bytes.Compare,
		"Contains":       bytes.Contains,
		"ContainsRune":   bytes.ContainsRune,
		"Count":          bytes.Count,
		"Equal":          bytes.Equal,
		"EqualFold":      bytes.EqualFold,
		"Fields":         bytes.Fields,
		"FieldsFunc":     bytes.FieldsFunc,
		"HasPrefix":      bytes.HasPrefix,
		"HasSuffix":      bytes.HasSuffix,
		"Index":          bytes.Index,
		"IndexAny":       bytes.IndexAny,
		"IndexByte":      bytes.IndexByte,
		"IndexFunc":      bytes.IndexFunc,
		"IndexRune":      bytes.IndexRune,
		"Join":           bytes.Join,
		"LastIndex":      bytes.LastIndex,
		"LastIndexAny":   bytes.LastIndexAny,
		"LastIndexByte":  bytes.LastIndexByte,
		"LastIndexFunc":  bytes.LastIndexFunc,
		"Map":            bytes.Map,
		"Repeat":         bytes.Repeat,
		"Replace":        bytes.Replace,
		"Runes":          bytes.Runes,
		"Split":          bytes.Split,
		"SplitAfter":     bytes.SplitAfter,
		"SplitAfterN":    bytes.SplitAfterN,
		"SplitN":         bytes.SplitN,
		"Title":          bytes.Title,
		"ToLower":        bytes.ToLower,
		"ToLowerSpecial": bytes.ToLowerSpecial,
		"ToTitle":        bytes.ToTitle,
		"ToTitleSpecial": bytes.ToTitleSpecial,
		"ToUpper":        bytes.ToUpper,
		"ToUpperSpecial": bytes.ToUpperSpecial,
		"Trim":           bytes.Trim,
		"TrimFunc":       bytes.TrimFunc,
		"TrimLeft":       bytes.TrimLeft,
		"TrimLeftFunc":   bytes.TrimLeftFunc,
		"TrimPrefix":     bytes.TrimPrefix,
		"TrimRight":      bytes.TrimRight,
		"TrimRightFunc":  bytes.TrimRightFunc,
		"TrimSpace":      bytes.TrimSpace,
		"TrimSuffix":     bytes.TrimSuffix,
	}
	PackageTypes["bytes"] = map[string]interface{}{
		"Buffer": bytes.Buffer{},
		"Reader": bytes.Reader{},
	}
}
