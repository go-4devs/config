package env

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

const Name = "env"

var _ config.Provider = (*Provider)(nil)

type Option func(*Provider)

func WithKeyFactory(factory func(...string) string) Option {
	return func(p *Provider) { p.key = factory }
}

func New(namespace, appName string, opts ...Option) *Provider {
	provider := Provider{
		key: func(path ...string) string {
			return strings.ToUpper(strings.Join(path, "_"))
		},
		prefix: strings.ToUpper(namespace + "_" + appName + "_"),
		name:   "",
	}

	for _, opt := range opts {
		opt(&provider)
	}

	return &provider
}

type Provider struct {
	key    func(...string) string
	name   string
	prefix string
}

func (p *Provider) Key(path ...string) string {
	return p.prefix + p.key(path...)
}

func (p *Provider) Name() string {
	return p.name
}

func (p *Provider) Value(_ context.Context, path ...string) (config.Value, error) {
	if val, ok := os.LookupEnv(p.Key(path...)); ok {
		return value.JString(val), nil
	}

	return nil, fmt.Errorf("%v:%w", p.Name(), config.ErrNotFound)
}

func (p *Provider) DumpReference(_ context.Context, w io.Writer, opt config.Options) error {
	return p.writeOptions(w, opt)
}

func (p *Provider) writeOptions(w io.Writer, opt config.Options, key ...string) error {
	for idx, option := range opt.Options() {
		if err := p.writeOption(w, option, key...); err != nil {
			return fmt.Errorf("option[%d]:%w", idx, err)
		}
	}

	return nil
}

func (p *Provider) writeOption(w io.Writer, opt config.Option, keys ...string) error {
	if desc := param.Description(opt); desc != "" {
		if _, derr := fmt.Fprintf(w, "# %v.\n", desc); derr != nil {
			return fmt.Errorf("write description:%w", derr)
		}
	}

	var err error

	switch one := opt.(type) {
	case config.Group:
		err = p.writeOptions(w, one, append(keys, one.Name())...)
	case config.Options:
		err = p.writeOptions(w, one, keys...)
	default:
		def, dok := param.Default(opt)

		prefix := ""
		if !dok || key.IsWild(keys...) {
			prefix = "#"
		}

		if !dok {
			def = ""
		}

		_, err = fmt.Fprintf(w, "%s%s=%v\n", prefix, p.Key(append(keys, one.Name())...), def)
	}

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
