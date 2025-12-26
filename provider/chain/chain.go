package chain

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
)

const Name = "chain"

func New(c ...config.Provider) config.BindProvider {
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
