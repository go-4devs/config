package config

import (
	"context"

	"gitoa.ru/go-4devs/config/param"
)

type Provider interface {
	Value(ctx context.Context, path ...string) (Value, error)
	Name() string
}

type WatchCallback func(ctx context.Context, oldVar, newVar Value) error

type WatchProvider interface {
	Watch(ctx context.Context, callback WatchCallback, path ...string) error
}

type Factory interface {
	Name() string
	Create(ctx context.Context, prov Provider) (Provider, error)
}

type Option interface {
	Name() string
	Param(key any) (any, bool)
}

type Group interface {
	Option
	Options
}

type Options interface {
	Options() []Option
}

type BindProvider interface {
	Provider

	Bind(ctx context.Context, data Variables) error
}

type Variables interface {
	ByName(name ...string) (Variable, error)
	ByParam(filter param.Has) (Variable, error)
	Variables() []Variable
}

type Definition interface {
	Add(opts ...Option)
}

type ProcessFunc func(ctx context.Context, in Value, opts ...param.Option) (Value, error)

func (o ProcessFunc) Process(ctx context.Context, in Value, opts ...param.Option) (Value, error) {
	return o(ctx, in, opts...)
}

type Processor interface {
	Process(ctx context.Context, in Value, opts ...param.Option) (Value, error)
}
