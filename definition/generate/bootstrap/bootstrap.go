package bootstrap

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"gitoa.ru/go-4devs/config/definition/generate/pkg"
)

//go:embed *.tpl
var tpls embed.FS

type Boot struct {
	Config

	imp       *pkg.Imports
	Configure []string
	OutName   string
}

func (b Boot) Imports() []pkg.Import {
	return b.imp.Imports()
}

type Config interface {
	File() string
	Methods() []string
	SkipContext() bool
	Prefix() string
	Suffix() string
	FullPkg() string
	Pkg() string
}

func Bootstrap(ctx context.Context, cfg Config) (string, error) {
	fInfo, err := os.Stat(cfg.File())
	if err != nil {
		return "", fmt.Errorf("stat:%w", err)
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
			"gitoa.ru/go-4devs/config/definition",
			"gitoa.ru/go-4devs/config/definition/generate",
			"os",
			"fmt",
			"go/format",
			pkgPath,
		)

	data := Boot{
		imp:       imports,
		Configure: cfg.Methods(),
		OutName:   fInfo.Name()[0:len(fInfo.Name())-3] + "_config.go",
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
