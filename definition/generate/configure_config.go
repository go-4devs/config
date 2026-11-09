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

// readFile reads the configure file.
func (i ConfigureConfig) readFile(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "file")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"file"}, err)

	}

	return val.ParseString()

}

// ReadFile reads the configure file.
func (i ConfigureConfig) ReadFile() (string, error) {
	return i.readFile(i.ctx)
}

// File returns the configure file.
func (i ConfigureConfig) File() string {
	val, err := i.readFile(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readPrefix reads the struct name prefix.
func (i ConfigureConfig) readPrefix(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "prefix")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"prefix"}, err)

	}

	return val.ParseString()

}

// ReadPrefix reads the struct name prefix.
func (i ConfigureConfig) ReadPrefix() (string, error) {
	return i.readPrefix(i.ctx)
}

// Prefix returns the struct name prefix.
func (i ConfigureConfig) Prefix() string {
	val, err := i.readPrefix(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readSuffix reads the struct name suffix.
func (i ConfigureConfig) readSuffix(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "suffix")
	if err != nil {
		i.handle(ctx, err)

		return "Config", nil
	}

	return val.ParseString()

}

// ReadSuffix reads the struct name suffix.
func (i ConfigureConfig) ReadSuffix() (string, error) {
	return i.readSuffix(i.ctx)
}

// Suffix returns the struct name suffix.
func (i ConfigureConfig) Suffix() string {
	val, err := i.readSuffix(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readSkipContext reads whether generated methods skip context.
func (i ConfigureConfig) readSkipContext(ctx context.Context) (v bool, e error) {
	val, err := i.Value(ctx, "skip-context")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"skip-context"}, err)

	}

	return val.ParseBool()

}

// ReadSkipContext reads whether generated methods skip context.
func (i ConfigureConfig) ReadSkipContext() (bool, error) {
	return i.readSkipContext(i.ctx)
}

// SkipContext returns whether generated methods skip context.
func (i ConfigureConfig) SkipContext() bool {
	val, err := i.readSkipContext(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readBuildTags reads the build tags.
func (i ConfigureConfig) readBuildTags(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "build-tags")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"build-tags"}, err)

	}

	return val.ParseString()

}

// ReadBuildTags reads the build tags.
func (i ConfigureConfig) ReadBuildTags() (string, error) {
	return i.readBuildTags(i.ctx)
}

// BuildTags returns the build tags.
func (i ConfigureConfig) BuildTags() string {
	val, err := i.readBuildTags(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readOutName reads the output file name.
func (i ConfigureConfig) readOutName(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "out-name")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"out-name"}, err)

	}

	return val.ParseString()

}

// ReadOutName reads the output file name.
func (i ConfigureConfig) ReadOutName() (string, error) {
	return i.readOutName(i.ctx)
}

// OutName returns the output file name.
func (i ConfigureConfig) OutName() string {
	val, err := i.readOutName(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readMethods reads the configure method names.
func (i ConfigureConfig) readMethods(ctx context.Context) (v []string, e error) {
	val, err := i.Value(ctx, "method")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"method"}, err)

	}

	return v, val.Unmarshal(&v)
}

// ReadMethods reads the configure method names.
func (i ConfigureConfig) ReadMethods() ([]string, error) {
	return i.readMethods(i.ctx)
}

// Methods returns the configure method names.
func (i ConfigureConfig) Methods() []string {
	val, err := i.readMethods(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readLeaveTemps reads the temporary files to keep, such as bootstrap and config.
func (i ConfigureConfig) readLeaveTemps(ctx context.Context) (v []string, e error) {
	val, err := i.Value(ctx, "leave-temp")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"leave-temp"}, err)

	}

	return v, val.Unmarshal(&v)
}

// ReadLeaveTemps reads the temporary files to keep, such as bootstrap and config.
func (i ConfigureConfig) ReadLeaveTemps() ([]string, error) {
	return i.readLeaveTemps(i.ctx)
}

// LeaveTemps returns the temporary files to keep, such as bootstrap and config.
func (i ConfigureConfig) LeaveTemps() []string {
	val, err := i.readLeaveTemps(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}

// readFullPkg reads the full package path.
func (i ConfigureConfig) readFullPkg(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "full-pkg")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"full-pkg"}, err)

	}

	return val.ParseString()

}

// ReadFullPkg reads the full package path.
func (i ConfigureConfig) ReadFullPkg() (string, error) {
	return i.readFullPkg(i.ctx)
}

// FullPkg returns the full package path.
func (i ConfigureConfig) FullPkg() string {
	val, err := i.readFullPkg(i.ctx)
	if err != nil {
		i.handle(i.ctx, err)
	}

	return val
}
