package memory

import (
	"context"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/key"
	"github.com/go-4devs/config/validator"
)

var _ config.BindProvider = (*ValueValidate)(nil)

// ValueValid returns a provider that checks a value when it is read.
// Bind records the variables. A value read before Bind is not checked.
func ValueValid(in config.Provider) *ValueValidate {
	return &ValueValidate{
		WrapProvider: Wrap(in),
		idx:          key.Map{},
		vars:         nil,
	}
}

type ValueValidate struct {
	WrapProvider

	idx  key.Map
	vars []config.Variable
}

// Bind records the variables, then binds the inner provider.
func (a *ValueValidate) Bind(ctx context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		a.idx.Add(len(a.vars), opt.Key())
		a.vars = append(a.vars, opt)
	}

	if err := a.WrapProvider.Bind(ctx, def); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Value returns the inner value. A variable recorded by Bind is checked.
func (a *ValueValidate) Value(ctx context.Context, path ...string) (config.Value, error) {
	val, err := a.Provider.Value(ctx, path...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	id, ok := a.idx.Index(path)
	if !ok {
		return val, nil
	}

	if err = validator.Validate(a.vars[id], val); err != nil {
		return nil, fmt.Errorf("%w", option.Err(err, path))
	}

	return val, nil
}
