package memory

import (
	"context"
	"fmt"
	"io"

	"gitoa.ru/go-4devs/config"
)

var (
	_ config.BindProvider = Wrap(nil)
	_ config.DumpProvider = Wrap(nil)
	_ config.Providers    = Wrap(nil)
)

// Wrap returns a provider that forwards Bind, DumpReference, ByName and Names.
func Wrap(prov config.Provider) WrapProvider {
	return WrapProvider{
		Provider: prov,
	}
}

type WrapProvider struct {
	config.Provider
}

// Bind calls Bind when the inner provider implements it.
func (w WrapProvider) Bind(ctx context.Context, data config.Variables) error {
	if prov, ok := w.Provider.(config.BindProvider); ok {
		if err := prov.Bind(ctx, data); err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

// DumpReference calls DumpReference when the inner provider implements it.
func (w WrapProvider) DumpReference(ctx context.Context, in io.Writer, opts config.Options) error {
	if prov, ok := w.Provider.(config.DumpProvider); ok {
		err := prov.DumpReference(ctx, in, opts)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil
}

// ByName returns a nested provider when the inner provider has one.
func (w WrapProvider) ByName(name string) (config.Provider, error) {
	if prov, ok := w.Provider.(config.Providers); ok {
		out, err := prov.ByName(name)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return out, nil
	}

	return nil, fmt.Errorf("%w", config.ErrNotFound)
}

// Names returns nested provider names when the inner provider has them.
func (w WrapProvider) Names() []string {
	if prov, ok := w.Provider.(config.Providers); ok {
		return prov.Names()
	}

	return nil
}
