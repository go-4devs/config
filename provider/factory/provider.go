package factory

import (
	"context"

	"gitoa.ru/go-4devs/config"
)

var _ config.Factory = New("", nil)

type Create func(ctx context.Context, prov config.Provider) (config.Provider, error)

func New(name string, fn Create) Factory {
	return Factory{
		create: fn,
		name:   name,
	}
}

type Factory struct {
	create Create
	name   string
}

func (f Factory) Name() string {
	return f.name
}

func (f Factory) Create(ctx context.Context, prov config.Provider) (config.Provider, error) {
	return f.create(ctx, prov)
}
