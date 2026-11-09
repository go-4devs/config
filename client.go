package config

import (
	"context"
	"errors"
	"fmt"
)

var _ Providers = (*Client)(nil)

// Must returns a client for the providers and panics if New fails.
func Must(providers ...any) *Client {
	client, err := New(providers...)
	if err != nil {
		panic(err)
	}

	return client
}

// New returns a client that reads the providers in order.
// A provider may be a Provider or a Factory. A Factory is created on first use.
func New(providers ...any) (*Client, error) {
	client := &Client{
		providers: make([]Provider, len(providers)),
		name:      make(map[string]int),
		chain:     make([]Providers, 0, len(providers)),
	}

	for idx, prov := range providers {
		var name string

		switch current := prov.(type) {
		case Provider:
			client.providers[idx] = current
			name = current.Name()
		case Factory:
			client.providers[idx] = WrapFactory(current, client)
			name = current.Name()
		default:
			return nil, fmt.Errorf("provier[%d]: %w %T", idx, ErrUnknowType, prov)
		}

		client.name[name] = idx
		if current, ok := prov.(Providers); ok {
			client.chain = append(client.chain, current)
		}
	}

	return client, nil
}

type Client struct {
	providers []Provider
	name      map[string]int
	chain     []Providers
}

// Name returns the client name.
func (c *Client) Name() string {
	return "client"
}

// Value returns the value at path.
// The next provider is used when the current one returns ErrNotFound or ErrInitFactory.
func (c *Client) Value(ctx context.Context, path ...string) (Value, error) {
	var (
		value Value
		err   error
	)

	for _, provider := range c.providers {
		value, err = provider.Value(ctx, path...)
		if err == nil || (!errors.Is(err, ErrNotFound) && !errors.Is(err, ErrInitFactory)) {
			break
		}
	}

	if err != nil {
		return value, fmt.Errorf("client failed get value: %w", err)
	}

	return value, nil
}

// Watch calls Watch on every provider that implements WatchProvider.
// ErrNotFound and ErrInitFactory from a provider are skipped.
func (c *Client) Watch(ctx context.Context, callback WatchCallback, path ...string) error {
	for idx, prov := range c.providers {
		provider, ok := prov.(WatchProvider)
		if !ok {
			continue
		}

		err := provider.Watch(ctx, callback, path...)
		if err != nil {
			if errors.Is(err, ErrNotFound) || errors.Is(err, ErrInitFactory) {
				continue
			}

			return fmt.Errorf("client: failed watch by provider[%d]: %w", idx, err)
		}
	}

	return nil
}

// Bind calls Bind on every provider that implements BindProvider.
func (c *Client) Bind(ctx context.Context, data Variables) error {
	for idx, prov := range c.providers {
		provider, ok := prov.(BindProvider)
		if !ok {
			continue
		}

		if err := provider.Bind(ctx, data); err != nil {
			return fmt.Errorf("bind[%d] %v:%w", idx, provider.Name(), err)
		}
	}

	return nil
}

// ByName returns the provider with that name.
func (c *Client) ByName(name string) (Provider, error) {
	if idx, ok := c.name[name]; ok {
		return c.providers[idx], nil
	}

	for _, prov := range c.chain {
		if cprov, err := prov.ByName(name); err == nil {
			return cprov, nil
		}
	}

	return nil, fmt.Errorf("provider[%v]:%w", c.Name(), ErrNotFound)
}

// Names returns the provider names, including names from nested provider lists.
func (c *Client) Names() []string {
	names := make([]string, 0, len(c.providers))
	for name := range c.name {
		names = append(names, name)
	}

	for _, prov := range c.chain {
		names = append(names, prov.Names()...)
	}

	return names
}
