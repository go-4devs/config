package env

import (
	"fmt"
	"io"
	"slices"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/key"
	"github.com/go-4devs/config/param"
)

// DumpReference writes env names for def.
// prefix reads key segments for an option and returns nil when the option has none.
func DumpReference(
	w io.Writer,
	def *definition.Definition,
	envKey func(...string) string,
	prefix func(param.Params) ([]string, bool),
) error {
	return reference{envKey: envKey, prefix: prefix, head: 0}.writeOptions(w, def)
}

// EmptyPrefix is the default reader: the option has no prefix segments.
func EmptyPrefix(param.Params) ([]string, bool) {
	return nil, false
}

type reference struct {
	envKey func(...string) string
	prefix func(param.Params) ([]string, bool)
	head   int
}

func (r reference) writeOptions(w io.Writer, opt config.Options, keys ...string) error {
	for idx, option := range opt.Options() {
		if err := r.writeOption(w, option, keys...); err != nil {
			return fmt.Errorf("option[%d]:%w", idx, err)
		}
	}

	return nil
}

func (r reference) writeOption(w io.Writer, opt config.Option, keys ...string) error {
	if next, ok := r.prefix(opt); ok {
		keys = slices.Concat(next, keys[r.head:])
		r.head = len(next)
	}

	if desc := param.Description(opt); desc != "" {
		if _, derr := fmt.Fprintf(w, "# %v.\n", desc); derr != nil {
			return fmt.Errorf("write description:%w", derr)
		}
	}

	var err error

	switch one := opt.(type) {
	case config.Group:
		err = r.writeOptions(w, one, append(keys, one.Name())...)
	case config.Options:
		err = r.writeOptions(w, one, keys...)
	default:
		def, dok := param.Default(opt)

		mark := ""
		if !dok || key.IsWild(keys...) {
			mark = "#"
		}

		if !dok {
			def = ""
		}

		_, err = fmt.Fprintf(w, "%s%s=%v\n", mark, r.envKey(append(keys, one.Name())...), def)
	}

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
