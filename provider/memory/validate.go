package memory

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
)

var _ config.BindProvider = (*Map)(nil)

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

func (a Validate) Bind(ctx context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		val, err := a.Value(ctx, opt.Key()...)
		if err != nil {
			continue
		}

		if err = validator.Validate(opt.Key(), opt, val); err != nil {
			return option.Err(err, opt.Key()) //nolint:wrapcheck
		}
	}

	return nil
}

func (a Validate) Name() string {
	if a.name != "" {
		return a.name
	}

	return a.Provider.Name()
}
