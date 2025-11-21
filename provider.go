package config

import "context"

type Provider interface {
	Value(ctx context.Context, path ...string) (Value, error)
}

type NamedProvider interface {
	Name() string
	Provider
}

type WatchCallback func(ctx context.Context, oldVar, newVar Value) error

type WatchProvider interface {
	Watch(ctx context.Context, callback WatchCallback, path ...string) error
}

type Factory func(ctx context.Context, cfg Provider) (Provider, error)

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

type Definition interface {
	Add(opts ...Option)
}

type BindProvider interface {
	Bind(ctx context.Context, def Definition)
}
