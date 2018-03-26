package main

func convert(pkgImportName string, notTypes []pkgFuncType) []pkgFuncType {
	switch pkgImportName {
	case "math":
		for k := range notTypes {
			if notTypes[k].FuncTypeName == "MaxUint64" {
				notTypes[k].Expr = "uint64(math.MaxUint64)"
			}
		}
	}

	return notTypes
}
