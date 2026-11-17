package bootstrap

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-4devs/config/definition/generate/pkg"
)

//go:embed *.tpl
var tpls embed.FS

type configure struct {
	Name       string
	StructName string
}

type Boot struct {
	Config
	*pkg.Packages

	Configure []configure
}

// Pkg returns the last element of the package path.
func (b Boot) Pkg() string {
	return pkg.Pkg(b.FullPkg())
}

type Config interface {
	File() string
	Methods() []string
	SkipContext() bool
	Prefix() string
	Suffix() string
	FullPkg() string
}

// Bootstrap writes a temporary program that calls the configure functions.
func Bootstrap(ctx context.Context, cfg Config) (string, error) {
	fInfo, err := os.Stat(cfg.File())
	if err != nil {
		return "", fmt.Errorf("stat[%v]:%w", cfg.File(), err)
	}

	pkgPath, err := pkg.ByPath(ctx, cfg.File(), fInfo.IsDir())
	if err != nil {
		return "", fmt.Errorf("pkg by path:%w", err)
	}

	tmpFile, err := os.CreateTemp(filepath.Dir(fInfo.Name()), "config-bootstrap")
	if err != nil {
		return "", fmt.Errorf("create tmp file:%w", err)
	}

	tpl, err := template.ParseFS(tpls, "bootstrap.go.tpl")
	if err != nil {
		return "", fmt.Errorf("parse template:%w", err)
	}

	imports := pkg.NewImports("main").
		Adds(
			"context",
			"github.com/go-4devs/config/definition",
			"github.com/go-4devs/config",
			"github.com/go-4devs/config/definition/generate/view",
			"github.com/go-4devs/config/param",
			"github.com/go-4devs/config/definition/generate",
			"os",
			"io",
			"fmt",
			pkgPath,
		)

	names, err := structNames(cfg.File())
	if err != nil {
		return "", fmt.Errorf("struct names:%w", err)
	}

	methods := cfg.Methods()
	conf := make([]configure, len(methods))

	for idx, name := range methods {
		conf[idx] = configure{
			Name:       name,
			StructName: names[name],
		}
	}

	data := Boot{
		Packages:  imports,
		Configure: conf,
		Config:    cfg,
	}

	if err := tpl.Execute(tmpFile, data); err != nil {
		return "", fmt.Errorf("execute:%w", err)
	}

	src := tmpFile.Name()
	if err := tmpFile.Close(); err != nil {
		return src, fmt.Errorf("close file:%w", err)
	}

	dest := src + ".go"

	if err := os.Rename(src, dest); err != nil {
		return dest, fmt.Errorf("rename idt:%w", err)
	}

	return dest, nil
}

const structNameDirective = "fconfig:structname:"

var errInvalidStructName = errors.New("invalid struct name")

// structNames reads //fconfig:structname: Name from function doc comments.
func structNames(file string) (map[string]string, error) {
	pfile, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parse:%w", err)
	}

	names := make(map[string]string)

	var invalid error

	ast.Inspect(pfile, func(node ast.Node) bool {
		fn, ok := node.(*ast.FuncDecl)
		if !ok || fn.Recv != nil || fn.Name == nil {
			return true
		}

		name, ok := structName(fn.Doc)
		if !ok {
			return true
		}

		if name == "" || !token.IsIdentifier(name) || token.IsKeyword(name) {
			invalid = fmt.Errorf("%s: %w: %q", fn.Name.Name, errInvalidStructName, name)

			return false
		}

		names[fn.Name.Name] = name

		return true
	})

	if invalid != nil {
		return nil, invalid
	}

	return names, nil
}

func structName(doc *ast.CommentGroup) (string, bool) {
	if doc == nil {
		return "", false
	}

	for _, comment := range doc.List {
		text := strings.TrimSpace(comment.Text)
		text = strings.TrimSpace(strings.TrimPrefix(text, "//"))

		name, ok := strings.CutPrefix(text, structNameDirective)
		if !ok {
			continue
		}

		return strings.TrimSpace(name), true
	}

	return "", false
}
