package generate

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate/pkg"
	"gitoa.ru/go-4devs/config/definition/generate/render"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/param"
)

type Option func(*Config)

func WithSkipContext(c *Config) {
	view.WithSkipContext(c.Params)
}

func WithPrefix(name string) Option {
	return func(c *Config) {
		c.prefix = name
	}
}

func WithSuffix(name string) Option {
	return func(c *Config) {
		c.suffix = name
	}
}

// WithMethods set methosd.
//
// generate.WithMethods(runtime.FuncForPC(reflect.ValueOf(configure).Pointer()).Name()).
func WithMethods(in ...string) Option {
	return func(c *Config) {
		c.methods = in
	}
}

func WithOutName(in string) Option {
	return func(c *Config) {
		c.outName = in
	}
}

func WithFile(in string) Option {
	return func(c *Config) {
		c.file = in
	}
}

func WithFullPkg(in string) Option {
	return func(c *Config) {
		c.fullPkg = in
	}
}

func NewConfig(opts ...Option) Config {
	var cfg Config

	cfg.Params = param.New()
	cfg.prefix = "Input"

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

type Config struct {
	param.Params

	methods   []string
	prefix    string
	suffix    string
	fullPkg   string
	pkg       string
	file      string
	buildTags string
	outName   string
}

func (c Config) BuildTags() string {
	return c.buildTags
}

func (c Config) Pkg() string {
	if c.pkg == "" {
		if idx := strings.LastIndex(c.fullPkg, "/"); idx != -1 {
			c.pkg = c.fullPkg[idx+1:]
		}
	}

	return c.pkg
}

func (c Config) FullPkg() string {
	return c.fullPkg
}

func (c Config) SkipContext() bool {
	return view.IsSkipContext(c.Params)
}

func (c Config) Methods() []string {
	return c.methods
}

func (c Config) Prefix() string {
	return c.prefix
}

func (c Config) Suffix() string {
	return c.suffix
}

func (c Config) File() string {
	return c.file
}

func (c Config) OutName() string {
	return c.outName
}

//go:embed tpl/*
var tpls embed.FS

var initTpl = template.Must(template.New("tpls").ParseFS(tpls, "tpl/*.tpl")).Lookup("init.go.tpl")

func Run(_ context.Context, cfg Config, w io.Writer, inputs ...Input) error {
	data := Data{
		Config: cfg,
		imp:    pkg.NewImports(cfg.FullPkg()).Adds("fmt", "context", "gitoa.ru/go-4devs/config"),
	}

	var buff bytes.Buffer

	for _, in := range inputs {
		vi := view.NewViews(in.Method(), data.Params, in.Options(), view.WithKeys())

		if err := render.Render(&buff, vi, data); err != nil {
			return fmt.Errorf("render:%w", err)
		}
	}

	if err := initTpl.Execute(w, data); err != nil {
		return fmt.Errorf("render base:%w", err)
	}

	if _, err := io.Copy(w, &buff); err != nil {
		return fmt.Errorf("copy:%w", err)
	}

	return nil
}

type Data struct {
	Config

	imp *pkg.Imports
}

func (f Data) Imports() []pkg.Import {
	return f.imp.Imports()
}

func (f Data) StructName(name string) string {
	return f.Prefix() + FuncName(name) + f.Suffix()
}

func (f Data) FuncName(in string) string {
	return FuncName(in)
}

func (f Data) AddType(pkg string) (string, error) {
	short, err := f.imp.AddType(pkg)
	if err != nil {
		return "", fmt.Errorf("data: %w", err)
	}

	return short, nil
}

func FuncName(name string) string {
	data := strings.Builder{}
	toUp := true

	for _, char := range name {
		isLeter := unicode.IsLetter(char)
		isAllowed := isLeter || unicode.IsDigit(char)

		switch {
		case isAllowed && !toUp:
			data.WriteRune(char)
		case !isAllowed:
			toUp = true
		case toUp:
			data.WriteString(strings.ToUpper(string(char)))

			toUp = false
		}
	}

	return data.String()
}

func NewInput(method string, options config.Options) Input {
	return Input{
		method:  method,
		options: options,
	}
}

type Input struct {
	options config.Options
	method  string
}

func (i Input) Method() string {
	return i.method
}

func (i Input) Options() config.Options {
	return i.options
}
