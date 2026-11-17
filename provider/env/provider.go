package env

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/value"
)

const Name = "env"

var (
	_ config.Provider     = (*Provider)(nil)
	_ config.DumpProvider = (*Provider)(nil)
)

type Option func(*Provider)

// WithKeyFactory sets the function that joins a path into an environment variable name.
func WithKeyFactory(factory func(...string) string) Option {
	return func(p *Provider) { p.key = factory }
}

// WithName sets the provider name.
func WithName(name string) Option {
	return func(p *Provider) {
		p.name = name
	}
}

// WithPrefix sets the variable name prefix.
func WithPrefix(prefix string) Option {
	return func(p *Provider) {
		p.prefix = prefix
	}
}

// New returns a provider that reads environment variables for namespace and appName.
func New(namespace, appName string, opts ...Option) *Provider {
	provider := Provider{
		key: func(path ...string) string {
			return strings.ToUpper(strings.Join(path, "_"))
		},
		prefix: strings.ToUpper(namespace + "_" + appName + "_"),
		name:   Name,
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

// Key returns the environment variable name for the path.
func (p *Provider) Key(path ...string) string {
	return p.prefix + p.key(path...)
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Value returns the environment variable as raw text.
func (p *Provider) Value(_ context.Context, path ...string) (config.Value, error) {
	if val, ok := os.LookupEnv(p.Key(path...)); ok {
		return value.JString(val), nil
	}

	return nil, fmt.Errorf("%v:%w", p.Name(), config.ErrNotFound)
}

// DumpReference writes the variable names for opt.
func (p *Provider) DumpReference(_ context.Context, w io.Writer, opt config.Options) error {
	return DumpReference(w, opt, p.Key, EmptyPrefix)
}
