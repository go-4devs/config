package env

import (
	"context"
	"os"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/value"
)

var _ config.Provider = (*Provider)(nil)

type Option func(*Provider)

func WithKeyFactory(factory config.KeyFactory) Option {
	return func(p *Provider) { p.key = factory }
}

func New(opts ...Option) *Provider {
	provider := Provider{
		key: func(ctx context.Context, k config.Key) string {
			return strings.ToUpper(key.NsAppName("_")(ctx, k))
		},
	}

	for _, opt := range opts {
		opt(&provider)
	}

	return &provider
}

type Provider struct {
	key config.KeyFactory
}

func (p *Provider) Name() string {
	return "env"
}

func (p *Provider) IsSupport(ctx context.Context, key config.Key) bool {
	return p.key(ctx, key) != ""
}

func (p *Provider) Read(ctx context.Context, key config.Key) (config.Variable, error) {
	name := p.key(ctx, key)
	if val, ok := os.LookupEnv(name); ok {
		return config.Variable{
			Name:     name,
			Provider: p.Name(),
			Value:    value.JString(val),
		}, nil
	}

	return config.Variable{}, config.ErrVariableNotFound
}
