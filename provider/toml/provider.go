package toml

import (
	"context"
	"encoding/json"
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
)

const (
	Name      = "toml"
	Separator = "."
)

var _ config.Provider = (*Provider)(nil)

type Option func(*Provider)

func WithName(in string) Option {
	return func(p *Provider) {
		p.name = in
	}
}

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

func (p *Provider) Name() string {
	return p.name
}

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
