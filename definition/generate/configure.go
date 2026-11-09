package generate

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/generate/view"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
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

// WithPrefix sets the generated struct name prefix.
func WithPrefix(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionPrefix)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionPrefix, err)
		}

		return nil
	}
}

// WithSuffix sets the generated struct name suffix.
func WithSuffix(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionSuffix)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionSuffix, err)
		}

		return nil
	}
}

// WithSkipContext sets whether generated methods omit context.
func WithSkipContext(in bool) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(strconv.FormatBool(in), optionSkipContext)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionSkipContext, err)
		}

		return nil
	}
}

// WithBuildTags sets the build tags for the generator.
func WithBuildTags(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionBuildTags)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionBuildTags, err)
		}

		return nil
	}
}

// WithOutName sets the generated file name.
func WithOutName(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionOutName)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionOutName, err)
		}

		return nil
	}
}

// WithMethods sets the configure functions to generate.
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

// WithFullPkg sets the full package path of the generated file.
func WithFullPkg(in string) func(*memory.Map) error {
	return func(m *memory.Map) error {
		err := m.AppendOption(in, optionFullPkg)
		if err != nil {
			return fmt.Errorf("append %v:%w", optionFullPkg, err)
		}

		return nil
	}
}

// NewMemoryProvider returns the generator settings, including the source file.
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

// Configure declares the generator options.
func Configure(_ context.Context, def config.Definition) error {
	def.Add(
		option.String(OptionFile, "the configure file", option.Required),
		option.String(optionPrefix, "the struct name prefix"),
		option.String(optionSuffix, "the struct name suffix", option.Default("Config")),
		option.Bool(optionSkipContext, "whether generated methods skip context"),
		option.String(optionBuildTags, "the build tags"),
		option.String(optionOutName, "the output file name"),
		option.New(optionMethod, "the configure method names", []string{}, view.WithFuncName("methods")),
		option.New(optionLeaveTemp, "the temporary files to keep, such as bootstrap and config", []string{}, view.WithFuncName("LeaveTemps")),
		option.String(optionFullPkg, "the full package path"),
	)

	return nil
}
