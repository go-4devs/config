package memory

import (
	"context"
	"errors"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/validator"
)

var _ config.BindProvider = (*Map)(nil)

// Valid returns a provider that checks values during Bind.
func Valid(in config.Provider) Validate {
	return Validate{
		Provider: in,
		name:     "",
	}
}

type Validate struct {
	config.Provider

	name string
}

// Bind reads each variable and validates it. A missing value is skipped.
func (a Validate) Bind(ctx context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		val, err := a.Value(ctx, opt.Key()...)
		if errors.Is(err, config.ErrNotFound) {
			continue
		}

		if err != nil {
			return option.Err(err, opt.Key()) //nolint:wrapcheck
		}

		if err = validator.Validate(opt.Key(), opt, val); err != nil {
			return option.Err(err, opt.Key()) //nolint:wrapcheck
		}
	}

	return nil
}

// Name returns the provider name.
func (a Validate) Name() string {
	if a.name != "" {
		return a.name
	}

	return a.Provider.Name()
}
