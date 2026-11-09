package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

// NewImports returns an import set for the generated package.
func NewImports(pkg string) *Packages {
	imp := Packages{
		data: make(map[string]string),
		pkg:  pkg,
	}

	return &imp
}

type Packages struct {
	data map[string]string
	pkg  string
}

// Imports returns the registered imports.
func (i *Packages) Imports() []Import {
	imports := make([]Import, 0, len(i.data))
	for name, alias := range i.data {
		imports = append(imports, Import{
			Package: name,
			Alias:   alias,
		})
	}

	return imports
}

// Short returns the name used for fullType in generated code.
func (i *Packages) Short(fullType string) (string, error) {
	idx := strings.LastIndexByte(fullType, '.')
	if idx == -1 {
		return "", fmt.Errorf("%w: expect package.Type", ErrWrongFormat)
	}

	if alias, ok := i.data[fullType[:idx]]; ok {
		return alias + fullType[idx:], nil
	}

	return "", fmt.Errorf("%w alias for pkg %v", ErrNotFound, fullType[:idx])
}

// AddType registers the package of fullType and returns the name used in generated code.
func (i *Packages) AddType(fullType string) (string, error) {
	idx := strings.LastIndexByte(fullType, '.')
	if idx == -1 {
		return "", fmt.Errorf("%w: expect pckage.Type got %v", ErrWrongFormat, fullType)
	}

	imp := i.Add(fullType[:idx])

	if imp.Alias == "" {
		return fullType[idx+1:], nil
	}

	return imp.Alias + fullType[idx:], nil
}

// Adds registers the packages.
func (i *Packages) Adds(pkgs ...string) *Packages {
	for _, pkg := range pkgs {
		i.Add(pkg)
	}

	return i
}

// Add registers pkg. The generated package itself is imported with an empty alias.
func (i *Packages) Add(pkg string) Import {
	if pkg == i.pkg {
		return Import{
			Alias:   "",
			Package: pkg,
		}
	}

	alias := pkg

	if _, name, ok := strings.CutLast(pkg, "/"); ok {
		alias = AliasName(name)
	}

	if al, ok := i.data[pkg]; ok {
		return Import{Package: pkg, Alias: al}
	}

	for _, al := range i.data {
		if al == alias {
			alias += strconv.Itoa(len(i.data))
		}
	}

	i.data[pkg] = alias

	return Import{
		Alias:   alias,
		Package: pkg,
	}
}

// Pkg returns the last element of the generated package path.
func (i *Packages) Pkg() string {
	return Pkg(i.pkg)
}

type Import struct {
	Alias   string
	Package string
}

// Pkg returns the last element of the import path.
func (i Import) Pkg() string {
	return Pkg(i.Package)
}

// String returns the import clause.
func (i Import) String() string {
	if i.Alias == i.Pkg() {
		return fmt.Sprintf("%q", i.Package)
	}

	return fmt.Sprintf("%s %q", i.Alias, i.Package)
}

// Pkg returns the last element of the import path.
func Pkg(fullPkg string) string {
	if _, name, ok := strings.CutLast(fullPkg, "/"); ok {
		return name
	}

	return fullPkg
}
