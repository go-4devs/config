package memory

import (
	"context"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

const NameDefault = "default"

var _ config.BindProvider = (*Default)(nil)

type Default struct {
	data Map
	name string
}

// Value returns the value stored under key.
func (a *Default) Value(ctx context.Context, key ...string) (config.Value, error) {
	if v, err := a.data.Value(ctx, key...); err == nil {
		return v, nil
	}

	return nil, fmt.Errorf("%w", config.ErrNotFound)
}

// Bind stores each variable default under its key.
func (a *Default) Bind(_ context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		if data, ok := param.Default(opt); ok {
			a.data.SetOption(data, opt.Key()...)
		}
	}

	return nil
}

// Name returns the provider name.
func (a *Default) Name() string {
	if a.name != "" {
		return a.name
	}

	return NameDefault
}
