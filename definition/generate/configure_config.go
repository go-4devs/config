package generate

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gitoa.ru/go-4devs/config"
)

func NewConfigure(ctx context.Context, prov config.Provider) ConfigureConfig {
	return ConfigureConfig{
		Provider: prov,
		ctx:      ctx,
		handle: func(err error) {
			if !errors.Is(err, config.ErrNotFound) {
				log.Print(err)
			}
		},
	}
}

type ConfigureConfig struct {
	config.Provider

	ctx    context.Context //nolint:containedctx
	handle func(err error)
}

func (i ConfigureConfig) ReadBuildTags() (string, error) {
	val, verr := i.Value(i.ctx, optionBuildTags)
	if verr != nil {
		return "", fmt.Errorf("get %v:%w", optionBuildTags, verr)
	}

	data, err := val.ParseString()
	if err != nil {
		return "", fmt.Errorf("parse %v:%w", optionBuildTags, err)
	}

	return data, nil
}

func (i ConfigureConfig) BuildTags() string {
	data, err := i.ReadBuildTags()
	if err != nil {
		i.handle(err)

		return ""
	}

	return data
}

func (i ConfigureConfig) ReadOutName() (string, error) {
	val, err := i.Value(i.ctx, optionOutName)
	if err != nil {
		return "", fmt.Errorf("get %v:%w", optionOutName, err)
	}

	data, derr := val.ParseString()
	if derr != nil {
		return "", fmt.Errorf("parse %v:%w", optionOutName, derr)
	}

	return data, nil
}

func (i ConfigureConfig) OutName() string {
	data, err := i.ReadOutName()
	if err != nil {
		i.handle(err)

		return ""
	}

	return data
}

func (i ConfigureConfig) ReadFile() (string, error) {
	val, err := i.Value(i.ctx, OptionFile)
	if err != nil {
		return "", fmt.Errorf("get %v:%w", OptionFile, err)
	}

	data, derr := val.ParseString()
	if derr != nil {
		return "", fmt.Errorf("parse %v:%w", OptionFile, derr)
	}

	return data, nil
}

func (i ConfigureConfig) File() string {
	data, err := i.ReadFile()
	if err != nil {
		i.handle(err)

		return ""
	}

	return data
}

func (i ConfigureConfig) ReadMethods() ([]string, error) {
	val, err := i.Value(i.ctx, optionMethod)
	if err != nil {
		return nil, fmt.Errorf("get %v:%w", optionMethod, err)
	}

	var data []string

	perr := val.Unmarshal(&data)
	if perr != nil {
		return nil, fmt.Errorf("unmarshal %v:%w", optionMethod, perr)
	}

	return data, nil
}

func (i ConfigureConfig) Methods() []string {
	data, err := i.ReadMethods()
	if err != nil {
		i.handle(err)

		return nil
	}

	return data
}

func (i ConfigureConfig) ReadSkipContext() (bool, error) {
	val, err := i.Value(i.ctx, optionSkipContext)
	if err != nil {
		return false, fmt.Errorf("get %v:%w", optionSkipContext, err)
	}

	data, derr := val.ParseBool()
	if derr != nil {
		return false, fmt.Errorf("parse %v:%w", optionSkipContext, derr)
	}

	return data, nil
}

func (i ConfigureConfig) SkipContext() bool {
	data, err := i.ReadSkipContext()
	if err != nil {
		i.handle(err)

		return false
	}

	return data
}

func (i ConfigureConfig) ReadPrefix() (string, error) {
	val, err := i.Value(i.ctx, optionPrefix)
	if err != nil {
		return "", fmt.Errorf("get %v: %w", optionPrefix, err)
	}

	data, derr := val.ParseString()
	if derr != nil {
		return "", fmt.Errorf("parse %v:%w", optionPrefix, derr)
	}

	return data, nil
}

func (i ConfigureConfig) Prefix() string {
	val, err := i.ReadPrefix()
	if err != nil {
		i.handle(err)

		return ""
	}

	return val
}

func (i ConfigureConfig) ReadSuffix() (string, error) {
	val, err := i.Value(i.ctx, optionSuffix)
	if err != nil {
		return "", fmt.Errorf("get %v:%w", optionSuffix, err)
	}

	data, derr := val.ParseString()
	if derr != nil {
		return "", fmt.Errorf("parse %v:%w", optionSuffix, derr)
	}

	return data, nil
}

func (i ConfigureConfig) Suffix() string {
	data, err := i.ReadSuffix()
	if err != nil {
		i.handle(err)

		return ""
	}

	return data
}

func (i ConfigureConfig) ReadFullPkg() (string, error) {
	val, err := i.Value(i.ctx, optionFullPkg)
	if err != nil {
		return "", fmt.Errorf("get %v:%w", optionFullPkg, err)
	}

	data, derr := val.ParseString()
	if derr != nil {
		return "", fmt.Errorf("parse %v:%w", optionFullPkg, derr)
	}

	return data, nil
}

func (i ConfigureConfig) FullPkg() string {
	data, err := i.ReadFullPkg()
	if err != nil {
		i.handle(err)

		return ""
	}

	return data
}
