package chain

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-4devs/config"
)

const Name = "chain"

type Providers interface {
	config.BindProvider
	config.Providers
}

// New returns a provider that reads c in order.
func New(c ...config.Provider) Providers {
	return chain(c)
}

type chain []config.Provider

// Value returns the first value. The next provider is used only when the error is ErrNotFound.
func (c chain) Value(ctx context.Context, name ...string) (config.Value, error) {
	var err error

	for _, in := range c {
		var val config.Value

		val, err = in.Value(ctx, name...)
		if err == nil {
			return val, nil
		}

		if !errors.Is(err, config.ErrNotFound) {
			break
		}
	}

	if err == nil {
		err = config.ErrNotFound
	}

	return nil, fmt.Errorf("%w", err)
}

// Bind calls Bind on every provider that implements BindProvider.
func (c chain) Bind(ctx context.Context, def config.Variables) error {
	for _, input := range c {
		if prov, ok := input.(config.BindProvider); ok {
			if err := prov.Bind(ctx, def); err != nil {
				return fmt.Errorf("prov[%s]:%w", input.Name(), err)
			}
		}
	}

	return nil
}

// Name returns "chain".
func (c chain) Name() string {
	return Name
}

// ByName returns this chain or a nested provider with that name.
func (c chain) ByName(name string) (config.Provider, error) {
	if c.Name() == name {
		return c, nil
	}

	for _, prov := range c {
		if prov.Name() == name {
			return prov, nil
		}

		cprov, ok := prov.(config.Providers)
		if !ok {
			continue
		}

		if in, err := cprov.ByName(name); err == nil {
			return in, nil
		}
	}

	return nil, fmt.Errorf("prov[%v]:%w", c.Name(), config.ErrNotFound)
}

// Names returns the provider names, including nested provider lists.
func (c chain) Names() []string {
	names := make([]string, 0, len(c))
	for _, prov := range c {
		names = append(names, prov.Name())

		if cprov, ok := prov.(config.Providers); ok {
			names = append(names, cprov.Names()...)
		}
	}

	return names
}
