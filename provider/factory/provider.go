package factory

import (
	"context"

	"github.com/go-4devs/config"
)

var _ config.Factory = New("", nil)

type Create func(ctx context.Context, prov config.Provider) (config.Provider, error)

// New returns a factory named name that calls fn.
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

// Name returns the factory name.
func (f Factory) Name() string {
	return f.name
}

// Create calls the factory function.
func (f Factory) Create(ctx context.Context, prov config.Provider) (config.Provider, error) {
	return f.create(ctx, prov)
}
