package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (i *Packages) Adds(pkgs ...string) *Packages {
	for _, pkg := range pkgs {
		i.Add(pkg)
	}

	return i
}

func (i *Packages) Add(pkg string) Import {
	if pkg == i.pkg {
		return Import{
			Alias:   "",
			Package: pkg,
		}
	}

	alias := pkg

	if idx := strings.LastIndexByte(pkg, '/'); idx != -1 {
		alias = AliasName(pkg[idx+1:])
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

func (i *Packages) Pkg() string {
	return Pkg(i.pkg)
}

type Import struct {
	Alias   string
	Package string
}

func (i Import) Pkg() string {
	return Pkg(i.Package)
}

func (i Import) String() string {
	if i.Alias == i.Pkg() {
		return fmt.Sprintf("%q", i.Package)
	}

	return fmt.Sprintf("%s %q", i.Alias, i.Package)
}

func Pkg(fullPkg string) string {
	if idx := strings.LastIndex(fullPkg, "/"); idx != -1 {
		return fullPkg[idx+1:]
	}

	return fullPkg
}
