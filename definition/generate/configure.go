package generate

import (
	"context"
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/generate/view"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/memory"
)

//go:generate go run ../../cmd/config/main.go config:generate --skip-context
const (
	OptionFile        = "file"
	optionPrefix      = "prefix"
	optionSuffix      = "suffix"
	optionSkipContext = "skip-context"
	optionBuildTags   = "build-tags"
	optionOutName     = "out-name"
	optionMethod      = "method"
	optionFullPkg     = "full-pkg"
	optionLeaveTemp   = "leave-temp"
)

func WithPrefix(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionPrefix)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionPrefix, err)
		}

		return nil
	}
}

func WithSuffix(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionSuffix)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionSuffix, err)
		}

		return nil
	}
}

func WithSkipContext(in bool) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(strconv.FormatBool(in), optionSkipContext)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionSkipContext, err)
		}

		return nil
	}
}

func WithBuildTags(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionBuildTags)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionBuildTags, err)
		}

		return nil
	}
}

func WithOutName(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionOutName)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionOutName, err)
		}

		return nil
	}
}

func WithMethods(methods ...string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		for _, method := range methods {
			err := m.AppendOption(method, optionMethod)
			if err != nil {
				return fmt.Errorf("append %v:%w", optionMethod, err)
			}
		}

		return nil
	}
}

func WithFullPkg(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionFullPkg)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionFullPkg, err)
		}

		return nil
	}
}

func NewMemoryProvider(file string, opts ...func(*memory.Map) error) (*memory.Map, error) {
	cfg := new(memory.Map)

	err := cfg.AppendOption(file, OptionFile)
	if err != nil {
		return nil, fmt.Errorf("append %v:%w", OptionFile, err)
	}

	for idx, opt := range opts {
		err = opt(cfg)
		if err != nil {
			return nil, fmt.Errorf("opt[%d]:%w", idx, err)
		}
	}

	return cfg, nil
}

func Configure(_ context.Context, def config.Definition) error {
	def.Add(
		option.String(OptionFile, "set file", option.Required),
		option.String(optionPrefix, "struct prefix"),
		option.String(optionSuffix, "struct suffix", option.Default("Config")),
		option.Bool(optionSkipContext, "skip contect to method"),
		option.String(optionBuildTags, "add build tags"),
		option.String(optionOutName, "set out name"),
		option.New(optionMethod, "set method", []string{}, view.WithFuncName("methods")),
		option.New(optionLeaveTemp, "leave temp files example:[bootstrap,config]", []string{}, view.WithFuncName("LeaveTemps")),
		option.String(optionFullPkg, "set full pkg"),
	)

	return nil
}
