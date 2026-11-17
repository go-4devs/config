package memory

import (
	"context"
	"fmt"
	"slices"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/key"
	"github.com/go-4devs/config/param"
)

type pkey int

const keyPrefix pkey = 1

var (
	_ config.Provider     = (*PrefixProvider)(nil)
	_ config.BindProvider = (*PrefixProvider)(nil)
	_ config.DumpProvider = (*PrefixProvider)(nil)
)

// Prefix stores key segments inserted before the option path.
// The option path itself stays the same.
func Prefix(keys ...string) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, keyPrefix, keys)
	}
}

// ParamPrefixFn reads key segments from option parameters.
// A nil function uses ParamPrefix.
type ParamPrefixFn func(param.Params) ([]string, bool)

// ParamPrefix returns the key segments stored by Prefix.
func ParamPrefix(in param.Params) ([]string, bool) {
	val, ok := in.Param(keyPrefix)
	if !ok {
		return nil, false
	}

	keys, ok := val.([]string)

	return keys, ok
}

// NewPrefix returns a provider that prepends each variable's Prefix segments to its key.
func NewPrefix(parent config.Provider, opts ...func(*PrefixProvider)) *PrefixProvider {
	provider := &PrefixProvider{
		WrapProvider: Wrap(parent),
		paramPrefix:  ParamPrefix,
		idx:          key.Map{},
		keys:         nil,
	}

	for _, opt := range opts {
		opt(provider)
	}

	return provider
}

// WithParamPrefix sets the function Bind uses to read key segments.
func WithParamPrefix(fn ParamPrefixFn) func(*PrefixProvider) {
	return func(p *PrefixProvider) {
		p.paramPrefix = fn
	}
}

type PrefixProvider struct {
	WrapProvider

	paramPrefix ParamPrefixFn
	idx         key.Map
	keys        [][]string
}

// Bind remembers the Prefix of each variable. A later Bind replaces the previous map.
func (p *PrefixProvider) Bind(ctx context.Context, def config.Variables) error {
	p.idx = key.Map{}
	p.keys = nil

	for _, variable := range def.Variables() {
		segs, ok := p.paramPrefix(variable)
		if !ok {
			continue
		}

		p.idx.Add(len(p.keys), variable.Key())
		p.keys = append(p.keys, slices.Clone(segs))
	}

	if err := p.WrapProvider.Bind(ctx, def); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Value returns the parent value. Prefix segments stored for path are inserted before it.
func (p *PrefixProvider) Value(ctx context.Context, path ...string) (config.Value, error) {
	if idx, ok := p.idx.Index(path); ok {
		path = slices.Concat(p.keys[idx], path)
	}

	val, err := p.Provider.Value(ctx, path...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return val, nil
}
