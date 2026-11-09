package json

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/value"
	"github.com/tidwall/gjson"
)

const (
	Name      = "json"
	Separator = "."
)

var _ config.Provider = (*Provider)(nil)

// WithKey sets the function that joins a path into a JSON key.
func WithKey(fn func(...string) string) Option {
	return func(p *Provider) {
		p.key = fn
	}
}

// WithName sets the provider name.
func WithName(name string) Option {
	return func(p *Provider) {
		p.name = name
	}
}

// New returns a provider for the JSON bytes.
func New(json []byte, opts ...Option) *Provider {
	provider := Provider{
		key: func(s ...string) string {
			return strings.Join(s, Separator)
		},
		data: json,
		name: Name,
	}

	for _, opt := range opts {
		opt(&provider)
	}

	return &provider
}

// NewFile reads path and returns a JSON provider.
func NewFile(path string, opts ...Option) (*Provider, error) {
	file, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("%w: unable to read config file %#q: file not found or unreadable", err, path)
	}

	return New(file, opts...), nil
}

type Option func(*Provider)

type Provider struct {
	data []byte
	key  func(...string) string
	name string
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Value returns the JSON value at path.
func (p *Provider) Value(_ context.Context, path ...string) (config.Value, error) {
	key := p.key(path...)
	if val := gjson.GetBytes(p.data, key); val.Exists() {
		return value.JString(val.String()), nil
	}

	return nil, fmt.Errorf("%v:%w", p.Name(), config.ErrValueNotFound)
}
