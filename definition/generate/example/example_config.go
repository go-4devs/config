// Code generated gitoa.ru/go-4devs/config DO NOT EDIT.
package example

import (
	"context"
	"fmt"
	"gitoa.ru/go-4devs/config"
)

func WithExampleConfigHandle(fn func(context.Context, error)) func(*ExampleConfig) {
	return func(ci *ExampleConfig) {
		ci.handle = fn
	}
}

func NewExampleConfig(ctx context.Context, prov config.Provider, opts ...func(*ExampleConfig)) ExampleConfig {
	i := ExampleConfig{
		Provider: prov,
		handle: func(_ context.Context, err error) {
			fmt.Printf("ExampleConfig:%v", err)
		},
		ctx: ctx,
	}

	for _, opt := range opts {
		opt(&i)
	}

	return i
}

type ExampleConfig struct {
	config.Provider
	handle func(context.Context, error)
	ctx    context.Context
}

// readTest test string.
func (i ExampleConfig) readTest(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "test")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"test"}, err)

	}

	return val.ParseString()

}

// ReadTest test string.
func (i ExampleConfig) ReadTest() (string, error) {
	return i.readTest(i.ctx)
}

// Test test string.
func (i ExampleConfig) Test() string {
	val, err := i.readTest(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

type UserConfig struct {
	ExampleConfig
}

// User configure user.
func (i ExampleConfig) User() UserConfig {
	return UserConfig{i}
}

type LogConfig struct {
	ExampleConfig
}

// Log configure logger.
func (i ExampleConfig) Log() LogConfig {
	return LogConfig{i}
}

// readLevel log level.
func (i LogConfig) readLevel(ctx context.Context) (v Level, e error) {
	val, err := i.Value(ctx, "log", "level")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", "level"}, err)

	}

	pval, perr := val.ParseString()
	if perr != nil {
		return v, fmt.Errorf("parse [%v]:%w", []string{"log", "level"}, perr)
	}

	return v, v.UnmarshalText([]byte(pval))
}

// ReadLevel log level.
func (i LogConfig) ReadLevel(ctx context.Context) (Level, error) {
	return i.readLevel(ctx)
}

// Level log level.
func (i LogConfig) Level(ctx context.Context) Level {
	val, err := i.readLevel(ctx)
	if err != nil {
		i.handle(ctx, err)
	}

	return val
}

type LogServiceConfig struct {
	LogConfig
	service string
}

// Service servise logger.
func (i LogConfig) Service(key string) LogServiceConfig {
	return LogServiceConfig{i, key}
}

// readLevel log level.
func (i LogServiceConfig) readLevel(ctx context.Context) (v Level, e error) {
	val, err := i.Value(ctx, "log", i.service, "level")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"log", i.service, "level"}, err)

	}

	pval, perr := val.ParseString()
	if perr != nil {
		return v, fmt.Errorf("parse [%v]:%w", []string{"log", i.service, "level"}, perr)
	}

	return v, v.UnmarshalText([]byte(pval))
}

// ReadLevel log level.
func (i LogServiceConfig) ReadLevel(ctx context.Context) (Level, error) {
	return i.readLevel(ctx)
}

// Level log level.
func (i LogServiceConfig) Level(ctx context.Context) Level {
	val, err := i.readLevel(ctx)
	if err != nil {
		i.handle(ctx, err)
	}

	return val
}
