package chain

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
)

const Name = "chain"

type Providers interface {
	config.BindProvider
	config.Providers
}

func New(c ...config.Provider) Providers {
	return chain(c)
}

type chain []config.Provider

func (c chain) Value(ctx context.Context, name ...string) (config.Value, error) {
	for _, in := range c {
		if val, err := in.Value(ctx, name...); err == nil {
			return val, nil
		}
	}

	return nil, fmt.Errorf("%w", config.ErrNotFound)
}

func (c chain) Bind(ctx context.Context, def config.Variables) error {
	for _, input := range c {
		if prov, ok := input.(config.BindProvider); ok {
			if err := prov.Bind(ctx, def); err != nil {
				return fmt.Errorf("%T:%w", input, err)
			}
		}
	}

	return nil
}

func (c chain) Name() string {
	return Name
}

func (c chain) Provider(name string) (config.Provider, error) {
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

		if in, err := cprov.Provider(name); err == nil {
			return in, nil
		}
	}

	return nil, fmt.Errorf("prov[%v]:%w", c.Name(), config.ErrNotFound)
}

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
