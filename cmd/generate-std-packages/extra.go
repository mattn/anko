package main

type kv struct {
	k string
	v string
}
type extra struct {
	detail             string
	packageImport      []kv
	packageTypesImport []kv
}

var extraList = map[string]extra{
	"sort": {
		detail: `// SortFuncsStruct provides functions to be used with Sort
type SortFuncsStruct struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (s SortFuncsStruct) Len() int           { return s.LenFunc() }
func (s SortFuncsStruct) Less(i, j int) bool { return s.LessFunc(i, j) }
func (s SortFuncsStruct) Swap(i, j int)      { s.SwapFunc(i, j) }
`,
		packageTypesImport: []kv{{"SortFuncsStruct", "&SortFuncsStruct{}"}},
	},
}
