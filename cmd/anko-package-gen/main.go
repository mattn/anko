package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func pkgName(f string) string {
	file, err := parser.ParseFile(token.NewFileSet(), f, nil, parser.PackageClauseOnly)
	if err != nil || file == nil {
		return ""
	}
	return file.Name.Name
}

func isGoFile(dir os.FileInfo) bool {
	return !dir.IsDir() &&
		!strings.HasPrefix(dir.Name(), ".") && // ignore .files
		filepath.Ext(dir.Name()) == ".go"
}

func isPkgFile(dir os.FileInfo) bool {
	return isGoFile(dir) && !strings.HasSuffix(dir.Name(), "_test.go") // ignore test files
}

func parseDir(p string) (map[string]*ast.Package, error) {
	isGoFile := func(d os.FileInfo) bool {
		return !d.IsDir() && !strings.HasSuffix(d.Name(), "_test.go") && !strings.HasPrefix(d.Name(), "example_")
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), p, isGoFile, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return pkgs, nil
}

func main() {
	pkg := "flag"
	if len(os.Args) == 2 {
		pkg = os.Args[1]
	}
	b, err := exec.Command("go", "env", "GOROOT").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	paths := []string{filepath.Join(strings.TrimSpace(string(b)), "src")}
	b, err = exec.Command("go", "env", "GOPATH").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range strings.Split(strings.TrimSpace(string(b)), string(filepath.ListSeparator)) {
		paths = append(paths, filepath.Join(p, "src"))
	}
	for _, p := range paths {
		pp := filepath.Join(p, pkg)
		pkgs, err := parseDir(pp)
		if err != nil || len(pkgs) == 0 {
			continue
		}
		names := map[string]bool{}
		pn := pkg
		for _, pp := range pkgs {
			pn = pp.Name
			for _, f := range pp.Files {
				for _, d := range f.Decls {
					switch decl := d.(type) {
					case *ast.GenDecl:
						for _, spec := range decl.Specs {
							if vspec, ok := spec.(*ast.ValueSpec); ok {
								for _, n := range vspec.Names {
									c := n.Name[0]
									if c < 'A' || c > 'Z' {
										continue
									}
									names[n.Name] = true
								}
							}
						}
					case *ast.FuncDecl:
						if decl.Recv != nil {
							continue
						}
						c := decl.Name.Name[0]
						if c < 'A' || c > 'Z' {
							continue
						}
						names[decl.Name.Name] = true
					}
				}
			}
		}
		keys := []string{}
		for k := range names {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		fmt.Printf(`// Package %s implements %s interface for anko script.
package %s

import (
	"%s"
)

func init() {
	Packages["%s"] = map[string]interface{}{
`, pn, pkg, pn, pkg, pn)
		for _, k := range keys {
			fmt.Printf(`	"%s": %s.%s,`+"\n", k, pn, k)
		}
		fmt.Println(`	}
}`)
	}
}
