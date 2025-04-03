package vm

import (
	"reflect"
	"sync"

	"github.com/mattn/anko/env"
)

// Importer is an interface to handle `import` statement
type Importer interface {
	// Import package with designated `name`
	Import(name string) (map[string]reflect.Value, bool)

	// ImportTypes imports type definition for designated `name`
	ImportTypes(name string) (map[string]reflect.Type, bool)

	// Each iterate over available packages
	Each(func(string, map[string]reflect.Value))

	// EachTypes iterate over available package types
	EachTypes(func(string, map[string]reflect.Type))

	// AppendMap merge packages and package types into this Importer.
	// If `pkgs` is not specified, ALL available packages and types will be merged.
	// Otherwise, only designated `pkgs` are merged into this importer.
	AppendMap(pkg map[string]map[string]reflect.Value, pkgTypes map[string]map[string]reflect.Type, pkgs ...string) Importer

	// Append merge packages and package types into this Importer.
	// If `pkgs` is not specified, ALL available packages and types will be merged.
	// Otherwise, only designated `pkgs` are merged into this importer.
	Append(i Importer, pkgs ...string) Importer

	// Remove package and types specified in `pkgs`
	Remove(pkgs ...string) Importer
}

type pkgImporter struct {
	sync.RWMutex
	pack      map[string]map[string]reflect.Value
	packTypes map[string]map[string]reflect.Type
}

// NewPackagesImporter return import handler from map of packages and types
func NewPackagesImporter(packages map[string]map[string]reflect.Value, packageTypes map[string]map[string]reflect.Type) Importer {
	imp := pkgImporter{}
	imp.copyMap(packages, packageTypes)
	return &imp
}

// NewStdPackagesImporter create package importer for packages defined in `anko/packages`.
// If `pkgs` is not set, ALL packages will be available for import.
// Otherwise, if `pkgs` are specified, only designated packages will be available for import.
func NewStdPackagesImporter(pkgs ...string) Importer {
	imp := &pkgImporter{}
	return imp.AppendMap(env.Packages, env.PackageTypes, pkgs...)

}

// NewPackagesWithStdImporter combines NewPackagesImporter and NewStdPackagesImporter.
func NewPackagesWithStdImporter(packages map[string]map[string]reflect.Value, packageTypes map[string]map[string]reflect.Type, pkgs ...string) Importer {
	imp := NewPackagesImporter(packages, packageTypes)
	return imp.AppendMap(env.Packages, env.PackageTypes, pkgs...)
}

func (i *pkgImporter) Import(name string) (map[string]reflect.Value, bool) {
	i.RLock()
	defer i.RUnlock()

	if i.pack == nil {
		return nil, false
	}
	v, ok := i.pack[name]
	return v, ok
}
func (i *pkgImporter) ImportTypes(name string) (map[string]reflect.Type, bool) {
	i.RLock()
	defer i.RUnlock()

	if i.packTypes == nil {
		return nil, false
	}
	t, ok := i.packTypes[name]
	return t, ok
}
func (i *pkgImporter) Each(fn func(string, map[string]reflect.Value)) {
	i.RLock()
	for name, pkg := range i.pack {
		fn(name, pkg)
	}
	i.RUnlock()
}
func (i *pkgImporter) EachTypes(fn func(string, map[string]reflect.Type)) {
	i.RLock()
	for name, typ := range i.packTypes {
		fn(name, typ)
	}
	i.RUnlock()
}
func (i *pkgImporter) AppendMap(pkg map[string]map[string]reflect.Value, pkgTypes map[string]map[string]reflect.Type, pkgs ...string) Importer {
	i.Lock()
	i.makeMap()
	for n, p := range pkg {
		if i.contains(pkgs, n) {
			i.pack[n] = p
		}
	}
	for n, t := range pkgTypes {
		if i.contains(pkgs, n) {
			i.packTypes[n] = t
		}
	}
	i.Unlock()

	return i
}

func (i *pkgImporter) Append(other Importer, pkgs ...string) Importer {
	i.Lock()
	i.makeMap()
	other.Each(func(name string, pkg map[string]reflect.Value) {
		if i.contains(pkgs, name) {
			i.pack[name] = pkg
		}
	})
	other.EachTypes(func(name string, typ map[string]reflect.Type) {
		if i.contains(pkgs, name) {
			i.packTypes[name] = typ
		}
	})
	i.Unlock()

	return i
}
func (i *pkgImporter) Remove(pkgs ...string) Importer {
	i.Lock()
	for _, name := range pkgs {
		delete(i.pack, name)
		delete(i.packTypes, name)
	}
	i.Unlock()

	return i
}

func (i *pkgImporter) contains(pkgs []string, name string) bool {
	if len(pkgs) == 0 {
		return true
	}
	for _, item := range pkgs {
		if item == name {
			return true
		}
	}
	return false
}

func (i *pkgImporter) copyMap(pkg map[string]map[string]reflect.Value, pkgTypes map[string]map[string]reflect.Type) {
	i.makeMap()
	for n, p := range pkg {
		i.pack[n] = p
	}
	for n, t := range pkgTypes {
		i.packTypes[n] = t
	}
}

func (i *pkgImporter) makeMap() {
	if i.pack == nil {
		i.pack = make(map[string]map[string]reflect.Value)
	}
	if i.packTypes == nil {
		i.packTypes = make(map[string]map[string]reflect.Type)
	}
}

// Import get packages from PkgImporter OR env.Packages.
func (o *Options) Import(name string) (map[string]reflect.Value, bool) {
	if o.PkgImporter != nil {
		return o.PkgImporter.Import(name)
	}
	v, ok := env.Packages[name]
	return v, ok
}

// ImportTypes get types from PkgImporter OR env.PackageTypes
func (o *Options) ImportTypes(name string) (map[string]reflect.Type, bool) {
	if o.PkgImporter != nil {
		return o.PkgImporter.ImportTypes(name)
	}
	t, ok := env.PackageTypes[name]
	return t, ok
}
