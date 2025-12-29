package generate

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"slices"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate/pkg"
)

const NameSuffix = "_config.go"

func Parse(ctx context.Context, name string) (Parser, error) {
	var parse Parser

	parse.file = name

	stats, err := os.Stat(name)
	if err != nil {
		return parse, fmt.Errorf("stats:%w", err)
	}

	parse.fullPkg, err = pkg.ByPath(ctx, name, stats.IsDir())
	if err != nil {
		return parse, fmt.Errorf("get pkg:%w", err)
	}

	parse.methods, err = NewParseMethods(
		name,
		[]reflect.Type{
			reflect.TypeFor[context.Context](),
			reflect.TypeFor[config.Definition](),
		},
		[]reflect.Type{reflect.TypeFor[error]()},
	)
	if err != nil {
		return parse, fmt.Errorf("parse methods:%w", err)
	}

	return parse, nil
}

func NewParseMethods(file string, params []reflect.Type, results []reflect.Type) ([]string, error) {
	pfile, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse:%w", err)
	}

	resultAlias := importAlias(pfile, results)
	paramsAlias := importAlias(pfile, params)

	var methods []string

	ast.Inspect(pfile, func(anode ast.Node) bool {
		if fn, ok := anode.(*ast.FuncDecl); ok &&
			fn.Recv == nil &&
			fn.Type != nil &&
			(fn.Type.Params != nil && len(params) == len(fn.Type.Params.List) || len(params) == 0 && fn.Type.Params == nil) &&
			(fn.Type.Results != nil && len(results) == len(fn.Type.Results.List) || len(results) == 0 && fn.Type.Results == nil) {
			if hasFields(fn.Type.Params, params, paramsAlias) && hasFields(fn.Type.Results, results, resultAlias) {
				methods = append(methods, fn.Name.String())
			}
		}

		return true
	})

	return methods, nil
}

func importAlias(file *ast.File, params []reflect.Type) map[int][]string {
	paramsAlias := make(map[int][]string, len(params))
	for idx := range params {
		name := params[idx].Name()
		if pkgPath := params[idx].PkgPath(); pkgPath != "" {
			name = pkg.Pkg(pkgPath)
		}

		paramsAlias[idx] = append(paramsAlias[idx], name)
	}

	ast.Inspect(file, func(anode ast.Node) bool {
		if exp, ok := anode.(*ast.ImportSpec); ok {
			pathName := strings.Trim(exp.Path.Value, "\"")
			pname := pkg.Pkg(pathName)

			if exp.Name != nil {
				pname = exp.Name.String()
			}

			for idx, param := range params {
				if pathName == param.PkgPath() {
					paramsAlias[idx] = append(paramsAlias[idx], pname)
				}
			}
		}

		return true
	})

	return paramsAlias
}

func hasFields(fields *ast.FieldList, params []reflect.Type, alias map[int][]string) bool {
	for idx, one := range fields.List {
		iparam := params[idx]
		if ident, iok := one.Type.(*ast.Ident); iok && iparam.String() == ident.String() {
			return true
		}

		selector, sok := one.Type.(*ast.SelectorExpr)
		if !sok {
			return false
		}

		if iparam.Name() != selector.Sel.String() {
			return false
		}

		salias, saok := selector.X.(*ast.Ident)
		if iparam.PkgPath() != "" && saok && !slices.Contains(alias[idx], salias.String()) {
			return false
		}
	}

	return true
}

type Parser struct {
	file    string
	fullPkg string
	methods []string
}

func (p Parser) OutName() string {
	return strings.ReplaceAll(p.file, ".go", NameSuffix)
}

func (p Parser) FullPkg() string {
	return p.fullPkg
}

func (p Parser) Methods() []string {
	return p.methods
}
