package memory

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

const NameDefault = "default"

var _ config.BindProvider = (*Default)(nil)

type Default struct {
	data Map
	name string
}

func (a *Default) Value(ctx context.Context, key ...string) (config.Value, error) {
	if v, err := a.data.Value(ctx, key...); err == nil {
		return v, nil
	}

	return nil, fmt.Errorf("%w", config.ErrNotFound)
}

func (a *Default) Bind(_ context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		if data, ok := param.Default(opt); ok {
			a.data.SetOption(data, opt.Key()...)
		}
	}

	return nil
}

func (a *Default) Name() string {
	if a.name != "" {
		return a.name
	}

	return NameDefault
}
