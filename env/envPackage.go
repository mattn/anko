package env

import "reflect"

// Package returns the methods for the package defined by name.
// If no package is found by this name in the environment,
// the globally defined package with this name is returned.
func (e *Env) Package(name string) map[string]reflect.Value {
	e.rwMutex.RLock()
	pkg, ok := e.packages[name]
	e.rwMutex.RUnlock()

	if ok {
		return pkg
	}

	if e.parent == nil {
		pkg, ok := Packages[name]
		if ok {
			return pkg
		}
	}

	return nil
}

// DefinePackage defines methods for the package name.
func (e *Env) DefinePackage(name string, values map[string]reflect.Value) {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	if e.packages == nil {
		e.packages = map[string]map[string]reflect.Value{}
	}
	e.packages[name] = make(map[string]reflect.Value, len(values))
	for k, v := range values {
		e.packages[name][k] = v
	}
}

// PackageTypes returns the types for the package defined by name.
// If no package is found by this name in the environment,
// the globally defined package with this name is returned
func (e *Env) PackageTypes(name string) map[string]reflect.Type {
	e.rwMutex.RLock()
	pkg, ok := e.packageTypes[name]
	e.rwMutex.RUnlock()

	if ok {
		return pkg
	}

	if e.parent == nil {
		pkg, ok := PackageTypes[name]
		if ok {
			return pkg
		}
	}

	return nil
}

// DefinePackageTypes defines types for the package name.
func (e *Env) DefinePackageTypes(name string, types map[string]reflect.Type) {
	e.rwMutex.RLock()
	defer e.rwMutex.RUnlock()

	if e.packageTypes == nil {
		e.packageTypes = map[string]map[string]reflect.Type{}
	}

	e.packageTypes[name] = make(map[string]reflect.Type, len(types))
	for k, v := range types {
		e.packageTypes[name][k] = v
	}
}
