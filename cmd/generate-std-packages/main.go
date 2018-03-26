package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync"
)

type GoBin struct {
	Go16  string
	Go17  string
	Go18  string
	Go19  string
	Go110 string
}

var goBin GoBin
var goGoroutine bool

func (r GoBin) RangeName() []string {
	// for order
	return []string{"Go16", "Go17", "Go18", "Go19", "Go110"}
}

func init() {
	flag.StringVar(&goBin.Go16, "go16", "", "where is go 1.6 bin path")
	flag.StringVar(&goBin.Go17, "go17", "", "where is go 1.7 bin path")
	flag.StringVar(&goBin.Go18, "go18", "", "where is go 1.8 bin path")
	flag.StringVar(&goBin.Go19, "go19", "", "where is go 1.9 bin path")
	flag.StringVar(&goBin.Go110, "go110", "", "where is go 1.10 bin path")
	flag.BoolVar(&goGoroutine, "p", false, "if open goroutine to generate code")
	flag.Parse()
}

func genPackageScope(pkgImportName, goBin string) (*output, error) {
	var stdout = new(bytes.Buffer)
	var stderr = new(bytes.Buffer)

	cmd := exec.Command(goBin, "run", "./cmd/generate-std-packages/gen-package-scope/main.go", "-package", pkgImportName)
	cmd.Env = append(os.Environ(), "GOROOT="+strings.TrimSuffix(goBin, "/bin/go"))
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("err %s, stderr %s", err, stderr.String())
	} else if s := stderr.String(); s != "" {
		return nil, fmt.Errorf(s)
	}

	var resp = output{}
	err = json.Unmarshal(stdout.Bytes(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func dealExtra(pkgImportName string) (string, []pkgFuncType, []pkgFuncType) {
	var extraDetail string
	var notTypes []pkgFuncType
	var typeAssigns []pkgFuncType

	if extra, ok := extraList[pkgImportName]; ok {
		extraDetail = extra.detail
		if extra.packageImport != nil {
			for _, v := range extra.packageImport {
				notTypes = append(notTypes, pkgFuncType{pkgImportName, v.k, v.v})
			}
		}
		if extra.packageTypesImport != nil {
			for _, v := range extra.packageTypesImport {
				typeAssigns = append(typeAssigns, pkgFuncType{pkgImportName, v.k, v.v})
			}
		}
	}

	return extraDetail, notTypes, typeAssigns
}

func renderPkgFile(pkgImportName, pkgExtraTmpl, goVersion string, notTypes, typeAssigns []pkgFuncType, typeVars []string) ([]byte, error) {
	extraDetail, notTypes2, typeAssigns2 := dealExtra(pkgImportName)

	notTypes = convert(pkgImportName, append(notTypes, notTypes2...))
	typeAssigns = append(typeAssigns, typeAssigns2...)

	var importPackages = []string{fmt.Sprintf(`"%s"`, pkgImportName)}
	if len(typeVars) > 0 {
		importPackages = append(importPackages, `"reflect"`)
	}

	var importMultiVersionCall string

	return parseTmpl(pkgExtraTmpl, map[string]interface{}{
		"go_version":                goVersion,
		"name":                      pkgImportName,
		"package_funcs":             notTypes,
		"package_types":             typeAssigns,
		"package_no_struct_types":   strings.Join(typeVars, "\n"),
		"import_packages":           strings.Join(importPackages, "\n"),
		"extra":                     extraDetail,
		"import_multi_version_call": importMultiVersionCall,
	})
}

func genPkgImportCodeAndSave(wg *sync.WaitGroup, pkgImportName string) {
	defer wg.Done()

	binNames := goBin.RangeName()

	var allOutput = make(map[string]*output)

	goBinV := reflect.ValueOf(goBin)
	for _, name := range binNames {
		resp, err := genPackageScope(pkgImportName, goBinV.FieldByName(name).String())
		if err != nil {
			panic(err)
		}
		allOutput[name] = resp
	}

	// first, now is 17
	firstResp := allOutput[binNames[0]]
	notTypes, typeAssigns, typeVars := firstResp.NotTypes, firstResp.TypeAssigns, firstResp.TypeVars
	result, err := renderPkgFile(pkgImportName, MainTmpl, goVersionToBuildTagName(binNames[0]), notTypes, typeAssigns, typeVars)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(genFileName(pkgImportName, ""), result, 0644); err != nil {
		panic(err)
	}

	for i := 0; i < len(binNames)-1; i++ {
		preName := binNames[i]
		sufName := binNames[i+1]
		notTypes1, typeAssigns1, typeVars1 := allOutput[preName].NotTypes, allOutput[preName].TypeAssigns, allOutput[preName].TypeVars
		notTypes2, typeAssigns2, typeVars2 := allOutput[sufName].NotTypes, allOutput[sufName].TypeAssigns, allOutput[sufName].TypeVars

		notTypes, typeAssigns, typeVars := diffSlice(notTypes2, typeAssigns2, typeVars2, notTypes1, typeAssigns1, typeVars1)
		if err != nil {
			panic(err)
		}
		if len(notTypes) != 0 || len(typeAssigns) != 0 || len(typeVars) != 0 {
			result, err := renderPkgFile(pkgImportName, DiffTmpl, goVersionToBuildTagName(sufName), notTypes, typeAssigns, typeVars)
			if err != nil {
				panic(err)
			}
			if err := ioutil.WriteFile(genFileName(pkgImportName, goVersionToFileExtName(sufName)), result, 0644); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	goBinV := reflect.ValueOf(goBin)
	goBinT := reflect.TypeOf(goBin)
	for i := 0; i < goBinV.NumField(); i++ {
		if goBinV.Field(i).String() == "" {
			panic(fmt.Sprintf("please input %s bin path", goBinT.Field(i).Name))
		}
	}

	wg := new(sync.WaitGroup)
	for _, pkgImportName := range []string{
		"encoding/json",
		"log",
		"fmt",
		"flag",
		"time",
		"math/big",
		"io/ioutil",
		"runtime",
		"os",
		"io",
		"net",
		"net/http",
		"os/exec",
		"sort",
		"os/signal",
		"path/filepath",
		"regexp",
		"net/url",
		"strconv",
		"strings",
		"path",
		"math/rand",
		"bytes",
		"net/http/cookiejar",
		"math",
		"errors",
	} {
		wg.Add(1)
		if goGoroutine {
			go genPkgImportCodeAndSave(wg, pkgImportName)
		} else {
			fmt.Printf("run %s\n", pkgImportName)

			genPkgImportCodeAndSave(wg, pkgImportName)
		}

		if err := ioutil.WriteFile(genFileName("packages", ""), []byte(PackagesTmpl), 0644); err != nil {
			panic(err)
		}

		wg.Wait()
	}
}
