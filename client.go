package config

import (
	"context"
	"errors"
	"fmt"
)

func Must(providers ...any) *Client {
	client, err := New(providers...)
	if err != nil {
		panic(err)
	}

	return client
}

func New(providers ...any) (*Client, error) {
	client := &Client{
		providers: make([]Provider, len(providers)),
	}

	for idx, prov := range providers {
		switch current := prov.(type) {
		case Provider:
			client.providers[idx] = current
		case Factory:
			client.providers[idx] = WrapFactory(current, client)
		default:
			return nil, fmt.Errorf("provier[%d]: %w %T", idx, ErrUnknowType, prov)
		}
	}

	return client, nil
}

type Client struct {
	providers []Provider
}

func (c *Client) Name() string {
	return "client"
}

// Value get value by name.
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
