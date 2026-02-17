// Code generated gitoa.ru/go-4devs/config DO NOT EDIT.
package generate

import (
	"context"
	"fmt"
	"gitoa.ru/go-4devs/config"
)

func WithConfigureConfigHandle(fn func(context.Context, error)) func(*ConfigureConfig) {
	return func(ci *ConfigureConfig) {
		ci.handle = fn
	}
}

func NewConfigureConfig(ctx context.Context, prov config.Provider, opts ...func(*ConfigureConfig)) ConfigureConfig {
	i := ConfigureConfig{
		Provider: prov,
		handle: func(_ context.Context, err error) {
			fmt.Printf("ConfigureConfig:%v", err)
		},
		ctx: ctx,
	}

	for _, opt := range opts {
		opt(&i)
	}

	return i
}

type ConfigureConfig struct {
	config.Provider
	handle func(context.Context, error)
	ctx    context.Context
}

// readFile set file.
func (i ConfigureConfig) readFile(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "file")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"file"}, err)

	}

	return val.ParseString()

}

// ReadFile set file.
func (i ConfigureConfig) ReadFile() (string, error) {
	return i.readFile(i.ctx)
}

// File set file.
func (i ConfigureConfig) File() string {
	val, err := i.readFile(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readPrefix struct prefix.
func (i ConfigureConfig) readPrefix(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "prefix")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"prefix"}, err)

	}

	return val.ParseString()

}

// ReadPrefix struct prefix.
func (i ConfigureConfig) ReadPrefix() (string, error) {
	return i.readPrefix(i.ctx)
}

// Prefix struct prefix.
func (i ConfigureConfig) Prefix() string {
	val, err := i.readPrefix(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readSuffix struct suffix.
func (i ConfigureConfig) readSuffix(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "suffix")
	if err != nil {
		i.handle(ctx, err)

		return "Config", nil
	}

	return val.ParseString()

}

// ReadSuffix struct suffix.
func (i ConfigureConfig) ReadSuffix() (string, error) {
	return i.readSuffix(i.ctx)
}

// Suffix struct suffix.
func (i ConfigureConfig) Suffix() string {
	val, err := i.readSuffix(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readSkipContext skip contect to method.
func (i ConfigureConfig) readSkipContext(ctx context.Context) (v bool, e error) {
	val, err := i.Value(ctx, "skip-context")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"skip-context"}, err)

	}

	return val.ParseBool()

}

// ReadSkipContext skip contect to method.
func (i ConfigureConfig) ReadSkipContext() (bool, error) {
	return i.readSkipContext(i.ctx)
}

// SkipContext skip contect to method.
func (i ConfigureConfig) SkipContext() bool {
	val, err := i.readSkipContext(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readBuildTags add build tags.
func (i ConfigureConfig) readBuildTags(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "build-tags")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"build-tags"}, err)

	}

	return val.ParseString()

}

// ReadBuildTags add build tags.
func (i ConfigureConfig) ReadBuildTags() (string, error) {
	return i.readBuildTags(i.ctx)
}

// BuildTags add build tags.
func (i ConfigureConfig) BuildTags() string {
	val, err := i.readBuildTags(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readOutName set out name.
func (i ConfigureConfig) readOutName(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "out-name")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"out-name"}, err)

	}

	return val.ParseString()

}

// ReadOutName set out name.
func (i ConfigureConfig) ReadOutName() (string, error) {
	return i.readOutName(i.ctx)
}

// OutName set out name.
func (i ConfigureConfig) OutName() string {
	val, err := i.readOutName(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readMethods set method.
func (i ConfigureConfig) readMethods(ctx context.Context) (v []string, e error) {
	val, err := i.Value(ctx, "method")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"method"}, err)

	}

	return v, val.Unmarshal(&v)
}

// ReadMethods set method.
func (i ConfigureConfig) ReadMethods() ([]string, error) {
	return i.readMethods(i.ctx)
}

// Methods set method.
func (i ConfigureConfig) Methods() []string {
	val, err := i.readMethods(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readLeaveTemps leave temp files example:[bootstrap,config].
func (i ConfigureConfig) readLeaveTemps(ctx context.Context) (v []string, e error) {
	val, err := i.Value(ctx, "leave-temp")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"leave-temp"}, err)

	}

	return v, val.Unmarshal(&v)
}

// ReadLeaveTemps leave temp files example:[bootstrap,config].
func (i ConfigureConfig) ReadLeaveTemps() ([]string, error) {
	return i.readLeaveTemps(i.ctx)
}

// LeaveTemps leave temp files example:[bootstrap,config].
func (i ConfigureConfig) LeaveTemps() []string {
	val, err := i.readLeaveTemps(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readFullPkg set full pkg.
func (i ConfigureConfig) readFullPkg(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "full-pkg")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"full-pkg"}, err)

	}

	return val.ParseString()

}

// ReadFullPkg set full pkg.
func (i ConfigureConfig) ReadFullPkg() (string, error) {
	return i.readFullPkg(i.ctx)
}

// FullPkg set full pkg.
func (i ConfigureConfig) FullPkg() string {
	val, err := i.readFullPkg(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}
