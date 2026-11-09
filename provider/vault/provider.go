package vault

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/value"
	"github.com/hashicorp/vault/api"
)

const (
	Name      = "vault"
	Separator = "/"
	Prefix    = "secret/data/"
	ValueName = "value"
)

var _ config.Provider = (*Provider)(nil)

type SecretOption func(*Provider)

// WithSecretResolve sets how a path becomes a secret path and a field.
func WithSecretResolve(f func(key []string) (string, string)) SecretOption {
	return func(s *Provider) { s.resolve = f }
}

// WithName sets the provider name.
func WithName(name string) SecretOption {
	return func(p *Provider) {
		p.name = name
	}
}

// WithPrefix sets the secret path prefix.
func WithPrefix(prefix string) SecretOption {
	return func(p *Provider) {
		p.prefix = prefix
	}
}

// New returns a Vault provider for namespace and appName.
func New(namespace, appName string, client *api.Client, opts ...SecretOption) *Provider {
	prov := Provider{
		client: client,
		resolve: func(key []string) (string, string) {
			keysLen := len(key)
			if keysLen == 1 {
				return "", key[0]
			}

			return strings.Join(key[:keysLen-1], Separator), key[keysLen-1]
		},
		name:   Name,
		prefix: Prefix + namespace + Separator + appName,
	}

	for _, opt := range opts {
		opt(&prov)
	}

	return &prov
}

type Provider struct {
	client  *api.Client
	resolve func(key []string) (string, string)
	name    string
	prefix  string
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Key returns the secret path and the field name.
func (p *Provider) Key(in []string) (string, string) {
	path, val := p.resolve(in)
	if path == "" {
		return p.prefix, val
	}

	return p.prefix + Separator + path, val
}

// Value returns the secret field.
func (p *Provider) Value(_ context.Context, key ...string) (config.Value, error) {
	path, field := p.Key(key)

	secret, err := p.read(path, field)
	if err != nil {
		return nil, fmt.Errorf("%w: path:%s, field:%s, provider:%s", err, path, field, p.Name())
	}

	if secret == nil || len(secret.Data) == 0 {
		return nil, fmt.Errorf("%w: path:%s, field:%s, provider:%s", config.ErrValueNotFound, path, field, p.Name())
	}

	data, ok := secret.Data["data"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%w: path:%s, field:%s, provider:%s", config.ErrValueNotFound, path, field, p.Name())
	}

	if val, ok := data[field]; ok {
		return value.JString(fmt.Sprint(val)), nil
	}

	if val, ok := data[ValueName]; ok {
		return value.JString(fmt.Sprint(val)), nil
	}

	md, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", config.ErrInvalidValue, err)
	}

	return value.JBytes(md), nil
}

func (p *Provider) read(path, key string) (*api.Secret, error) {
	secret, err := p.client.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("read[%s:%s]:%w", path, key, err)
	}

	if secret == nil && key != ValueName {
		return p.read(path+Separator+key, ValueName)
	}

	return secret, nil
}
