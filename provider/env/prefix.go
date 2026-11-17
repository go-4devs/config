package env

import (
	"context"
	"io"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/provider/memory"
)

var (
	_ config.Provider     = (*PrefixProvider)(nil)
	_ config.BindProvider = (*PrefixProvider)(nil)
)

// NewPrefixProvider wraps provider so each variable's Prefix is part of its env name.
func NewPrefixProvider(provider *Provider) *PrefixProvider {
	return &PrefixProvider{
		PrefixProvider: memory.NewPrefix(provider, memory.WithParamPrefix(ParamPrefix)),
		provider:       provider,
	}
}

// PrefixProvider reads env values with per-variable prefix segments.
type PrefixProvider struct {
	*memory.PrefixProvider

	provider *Provider
}

// DumpReference writes variable names. Prefix segments come from each option.
func (h *PrefixProvider) DumpReference(_ context.Context, w io.Writer, def *definition.Definition) error {
	return DumpReference(w, def, h.provider.Key, ParamPrefix)
}
