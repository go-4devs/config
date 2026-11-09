package ini

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
	"gopkg.in/ini.v1"
)

const (
	Name      = "ini"
	Separator = "."
)

var _ config.Provider = (*Provider)(nil)

// WithName sets the provider name.
func WithName(name string) func(*Provider) {
	return func(p *Provider) {
		p.name = name
	}
}

// WithResolve sets how a path becomes an INI section and key.
func WithResolve(fn func([]string) (string, string)) func(*Provider) {
	return func(p *Provider) {
		p.resolve = fn
	}
}

// New returns a provider for the INI file.
func New(data *ini.File, opts ...func(*Provider)) *Provider {
	prov := &Provider{
		data: data,
		resolve: func(path []string) (string, string) {
			if len(path) == 1 {
				return "", path[0]
			}

			return strings.Join(path[:len(path)-1], Separator), strings.ToUpper(path[len(path)-1])
		},
		name: Name,
	}

	for _, opt := range opts {
		opt(prov)
	}

	return prov
}

type Provider struct {
	data    *ini.File
	resolve func(path []string) (string, string)
	name    string
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Value returns the INI value at path.
func (p *Provider) Value(_ context.Context, path ...string) (config.Value, error) {
	section, name := p.resolve(path)

	iniSection, err := p.data.GetSection(section)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", config.ErrValueNotFound, p.Name(), err)
	}

	iniKey, err := iniSection.GetKey(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", config.ErrValueNotFound, p.Name(), err)
	}

	return value.JString(iniKey.String()), nil
}
