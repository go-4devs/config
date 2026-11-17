package env

import (
	"context"
	"io"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/provider/memory"
)

var (
	_ config.Provider     = (*Helper)(nil)
	_ config.BindProvider = (*Helper)(nil)
)

// NewHelper wraps provider so each variable's Prefix is part of its env name.
func NewHelper(provider *Provider) *Helper {
	return &Helper{
		PrefixProvider: memory.NewPrefix(provider, memory.WithParamPrefix(ParamPrefix)),
		provider:       provider,
	}
}

// Helper reads env values with per-variable prefix segments.
type Helper struct {
	*memory.PrefixProvider

	provider *Provider
}

// DumpReference writes variable names. Prefix segments come from each option.
func (h *Helper) DumpReference(_ context.Context, w io.Writer, def *definition.Definition) error {
	return DumpReference(w, def, h.provider.Key, ParamPrefix)
}
