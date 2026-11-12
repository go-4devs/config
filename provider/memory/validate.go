package memory

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
)

var _ config.BindProvider = (*Validate)(nil)

// BindValid returns a provider that checks values during Bind.
func BindValid(in config.BindProvider) Validate {
	return Validate{
		BindProvider: in,
		name:         "",
	}
}

type Validate struct {
	config.BindProvider

	name string
}

// Bind binds the inner provider, then checks each variable. A missing value is skipped.
func (a Validate) Bind(ctx context.Context, def config.Variables) error {
	if err := a.BindProvider.Bind(ctx, def); err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, opt := range def.Variables() {
		val, err := a.Value(ctx, opt.Key()...)
		if errors.Is(err, config.ErrNotFound) {
			continue
		}

		if err != nil {
			return fmt.Errorf("%w", option.Err(err, opt.Key()))
		}

		if err = validator.Validate(opt, val); err != nil {
			return fmt.Errorf("%w", option.Err(err, opt.Key()))
		}
	}

	return nil
}

// Name returns the provider name.
func (a Validate) Name() string {
	if a.name != "" {
		return a.name
	}

	return a.BindProvider.Name()
}
