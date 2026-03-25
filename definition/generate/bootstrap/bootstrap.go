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
	*pkg.Packages

	Configure []string
}

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
			"gitoa.ru/go-4devs/config/definition",
			"gitoa.ru/go-4devs/config",
			"gitoa.ru/go-4devs/config/definition/generate/view",
			"gitoa.ru/go-4devs/config/param",
			"gitoa.ru/go-4devs/config/definition/generate",
			"os",
			"io",
			"fmt",
			pkgPath,
		)

	data := Boot{
		Packages:  imports,
		Configure: cfg.Methods(),
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
