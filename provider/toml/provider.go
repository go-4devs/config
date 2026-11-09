package toml

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/value"
	toml "github.com/pelletier/go-toml/v2"
)

const (
	Name      = "toml"
	Separator = "."
)

var _ config.Provider = (*Provider)(nil)

type Option func(*Provider)

// WithName sets the provider name.
func WithName(in string) Option {
	return func(p *Provider) {
		p.name = in
	}
}

// New returns a provider for the TOML bytes.
func New(in []byte, opts ...Option) (*Provider, error) {
	var data Data
	if err := toml.Unmarshal(in, &data); err != nil {
		return nil, fmt.Errorf("toml failed load data: %w", err)
	}

	prov := &Provider{
		data: data,
		name: Name,
	}

	for _, opt := range opts {
		opt(prov)
	}

	return prov, nil
}

type Provider struct {
	data Data
	name string
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Value returns the value at path as JSON bytes.
func (p *Provider) Value(_ context.Context, path ...string) (config.Value, error) {
	val, err := p.data.Value(path...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	data, merr := json.Marshal(val)
	if merr != nil {
		return nil, fmt.Errorf("toml:%w", merr)
	}

	return value.JBytes(data), nil
}

type Data map[string]any

// Value returns the nested map value at path.
func (d Data) Value(path ...string) (any, error) {
	if len(path) == 1 {
		val, ok := d[path[0]]
		if !ok {
			return "", config.ErrValueNotFound
		}

		return val, nil
	}

	key, path := path[0], path[1:]

	val, ok := d[key]
	if !ok {
		return nil, config.ErrValueNotFound
	}

	data, _ := val.(map[string]any)

	return Data(data).Value(path...)
}
