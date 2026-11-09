package config

import (
	"context"
	"io"

	"github.com/go-4devs/config/param"
)

type Provider interface {
	// Value returns the value at path.
	Value(ctx context.Context, path ...string) (Value, error)
	// Name returns the provider name.
	Name() string
}

type WatchCallback func(ctx context.Context, oldVar, newVar Value) error

type WatchProvider interface {
	// Watch reports changes under path to callback.
	// The callback stops the watch by returning ErrStopWatch.
	Watch(ctx context.Context, callback WatchCallback, path ...string) error
}

type Factory interface {
	// Name returns the factory name.
	Name() string
	// Create builds the provider. prov is the client that owns the factory.
	Create(ctx context.Context, prov Provider) (Provider, error)
}

type Option interface {
	// Name returns the option name.
	Name() string
	// Param returns the value stored under key.
	Param(key any) (any, bool)
}

type Group interface {
	Option
	Options
}

type Options interface {
	// Options returns the child options.
	Options() []Option
	param.Params
}

type BindProvider interface {
	Provider

	// Bind writes provider values into data.
	Bind(ctx context.Context, data Variables) error
}

type DumpProvider interface {
	Provider

	// DumpReference writes a description of opts to w.
	DumpReference(ctx context.Context, w io.Writer, opts Options) error
}

type Providers interface {
	Provider

	// ByName returns the nested provider with that name.
	ByName(name string) (Provider, error)
	// Names returns the nested provider names.
	Names() []string
}

type Variables interface {
	// ByName returns the variable at path.
	ByName(name ...string) (Variable, error)
	// ByParam returns the single variable accepted by filter.
	ByParam(filter param.Has) (Variable, error)
	// Variables returns every variable.
	Variables() []Variable
}

type Definition interface {
	// Add appends options to the definition.
	Add(opts ...Option)
}

type ProcessFunc func(ctx context.Context, in Value, opts ...param.Option) (Value, error)

// Process calls the function.
func (o ProcessFunc) Process(ctx context.Context, in Value, opts ...param.Option) (Value, error) {
	return o(ctx, in, opts...)
}

type Processor interface {
	// Process converts in into another value.
	Process(ctx context.Context, in Value, opts ...param.Option) (Value, error)
}
