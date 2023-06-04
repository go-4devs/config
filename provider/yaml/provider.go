package yaml

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
	"gopkg.in/yaml.v3"
)

var _ config.Provider = (*Provider)(nil)

func keyFactory(_ context.Context, key config.Key) []string {
	return strings.Split(key.Name, "/")
}

func NewFile(name string, opts ...Option) (*Provider, error) {
	in, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("yaml_file: read error: %w", err)
	}

	return New(in, opts...)
}

func New(yml []byte, opts ...Option) (*Provider, error) {
	var data yaml.Node
	if err := yaml.Unmarshal(yml, &data); err != nil {
		return nil, fmt.Errorf("yaml: unmarshal err: %w", err)
	}

	return create(opts...).With(&data), nil
}

func create(opts ...Option) *Provider {
	prov := Provider{
		key: keyFactory,
	}

	for _, opt := range opts {
		opt(&prov)
	}

	return &prov
}

type Option func(*Provider)

type Provider struct {
	data node
	key  func(context.Context, config.Key) []string
}

func (p *Provider) Name() string {
	return "yaml"
}

func (p *Provider) Read(ctx context.Context, key config.Key) (config.Variable, error) {
	k := p.key(ctx, key)

	return p.data.read(p.Name(), k)
}

func (p *Provider) With(data *yaml.Node) *Provider {
	return &Provider{
		key:  p.key,
		data: node{Node: data},
	}
}

type node struct {
	*yaml.Node
}

func (n *node) read(name string, keys []string) (config.Variable, error) {
	val, err := getData(n.Node.Content[0].Content, keys)
	if err != nil {
		if errors.Is(err, config.ErrVariableNotFound) {
			return config.Variable{}, fmt.Errorf("%w: %s", config.ErrVariableNotFound, name)
		}

		return config.Variable{}, fmt.Errorf("%w: %s", err, name)
	}

	return config.Variable{
		Name:     strings.Join(keys, "."),
		Provider: name,
		Value:    value.Decode(val),
	}, nil
}

func getData(node []*yaml.Node, keys []string) (func(interface{}) error, error) {
	for idx := len(node) - 1; idx > 0; idx -= 2 {
		if node[idx-1].Value == keys[0] {
			if len(keys) > 1 {
				return getData(node[idx].Content, keys[1:])
			}

			return node[idx].Decode, nil
		}
	}

	return nil, config.ErrVariableNotFound
}
