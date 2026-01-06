package arg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/provider/memory"
)

const (
	doubleDash           = `--`
	defaultLenLognOption = 2
	dash                 = `-`
	Name                 = "arg"
)

var (
	_ config.DumpProvider = (*Argv)(nil)
	_ config.BindProvider = (*Argv)(nil)
)

// Deprecated: use WithArgs.
func WithSkip(skip int) func(*Argv) {
	return func(ar *Argv) {
		res := 2

		switch {
		case skip > 0 && len(os.Args) > skip:
			res = skip
		case skip > 0:
			res = len(os.Args)
		case len(os.Args) == 1:
			res = 1
		case len(os.Args) > 1 && os.Args[1][0] == '-':
			res = 1
		}

		ar.args = os.Args[res:]
	}
}

func WithArgs(args []string) func(*Argv) {
	return func(a *Argv) {
		a.args = args
	}
}

func WithName(name string) func(*Argv) {
	return func(a *Argv) {
		a.name = name
	}
}

func New(opts ...func(*Argv)) *Argv {
	arg := &Argv{
		args: os.Args[1:],
		pos:  0,
		Map:  memory.Map{},
		name: Name,
	}

	for _, opt := range opts {
		opt(arg)
	}

	return arg
}

type Argv struct {
	memory.Map

	args []string
	pos  uint64
	name string
}

func (i *Argv) Value(ctx context.Context, key ...string) (config.Value, error) {
	if err := i.parse(); err != nil {
		return nil, fmt.Errorf("parse:%w", err)
	}

	data, err := i.Map.Value(ctx, key...)
	if err != nil {
		return nil, fmt.Errorf("map: %w", err)
	}

	return data, nil
}

func (i *Argv) Bind(ctx context.Context, def config.Variables) error {
	options := true

	for len(i.args) > 0 {
		var err error

		arg := i.args[0]
		i.args = i.args[1:]

		switch {
		case options && arg == doubleDash:
			options = false
		case options && len(arg) > 2 && arg[0:2] == doubleDash:
			err = i.parseLongOption(arg[2:], def)
		case options && arg[0:1] == "-":
			if len(arg) == 1 {
				return fmt.Errorf("%w: option name required given '-'", config.ErrInvalidName)
			}

			err = i.parseShortOption(arg[1:], def)
		default:
			err = i.parseArgument(arg, def)
		}

		if err != nil {
			return fmt.Errorf("arg bind:%w", err)
		}
	}

	if err := i.Map.Bind(ctx, def); err != nil {
		return fmt.Errorf("arg map:%w", err)
	}

	return nil
}

func (i *Argv) Name() string {
	return i.name
}

func (i *Argv) DumpReference(_ context.Context, w io.Writer, opt config.Options) error {
	return NewDump().Reference(w, opt)
}

func (i *Argv) parseLongOption(arg string, def config.Variables) error {
	var value *string

	name := arg

	if strings.Contains(arg, "=") {
		vals := strings.SplitN(arg, "=", defaultLenLognOption)
		name = vals[0]
		value = &vals[1]
	}

	opt, err := def.ByName(key.ByPath(name, dash)...)
	if err != nil {
		return Err(err, name)
	}

	return i.appendOption(value, opt)
}

func (i *Argv) appendOption(data *string, opt config.Variable) error {
	if i.HasOption(opt.Key()...) && !option.IsSlice(opt) {
		return fmt.Errorf("%w: got: array, expect: %T", config.ErrUnexpectedType, param.Type(opt))
	}

	var val string

	switch {
	case data != nil:
		val = *data
	case option.IsBool(opt):
		val = "true"
	case len(i.args) > 0 && len(i.args[0]) > 0 && i.args[0][0:1] != "-":
		val = i.args[0]
		i.args = i.args[1:]
	default:
		return Err(config.ErrRequired, opt.Key()...)
	}

	err := i.AppendOption(val, opt.Key()...)
	if err != nil {
		return Err(err, opt.Key()...)
	}

	return nil
}

func (i *Argv) parseShortOption(arg string, def config.Variables) error {
	name := arg

	var value string

	if len(name) > 1 {
		name, value = arg[0:1], strings.TrimSpace(arg[1:])
	}

	opt, err := def.ByParam(option.HasShort(name))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if option.IsBool(opt) && value != "" {
		err := i.parseShortOption(value, def)
		if err != nil {
			return err
		}

		value = ""
	}

	if value == "" {
		return i.appendOption(nil, opt)
	}

	return i.appendOption(&value, opt)
}

func (i *Argv) parseArgument(arg string, def config.Variables) error {
	opt, err := def.ByParam(PosArgument(i.pos))
	if err != nil {
		var maxArgs uint64
		if i.pos > 0 {
			maxArgs -= i.pos
		}

		if errors.Is(err, config.ErrNotFound) {
			return fmt.Errorf("argument[%s] by pos[%d] max[%d]: %w",
				arg,
				i.pos+1,
				maxArgs,
				config.ErrNotFound,
			)
		}

		return fmt.Errorf("find argiment by pos[%d] max[%d]: %w", i.pos+1, maxArgs, err)
	}

	i.pos++

	if err := i.AppendOption(arg, opt.Key()...); err != nil {
		return Err(err, opt.Key()...)
	}

	return nil
}

func (i *Argv) parse() error {
	if i.Len() > 0 {
		return nil
	}

	for _, arg := range i.args {
		name, value, err := i.parseOne(arg)
		if err != nil {
			return err
		}

		if name != "" {
			if err := i.AppendOption(value, name); err != nil {
				return fmt.Errorf("append %v: %w", name, err)
			}
		}
	}

	return nil
}

// parseOne return name, value, error.
func (i *Argv) parseOne(arg string) (string, string, error) {
	if arg[0] != '-' {
		return "", "", nil
	}

	numMinuses := 1

	if arg[1] == '-' {
		numMinuses++
	}

	name := strings.TrimSpace(arg[numMinuses:])
	if len(name) == 0 {
		return name, "", nil
	}

	if name[0] == '-' || name[0] == '=' {
		return "", "", fmt.Errorf("%w: bad flag syntax: %s", config.ErrInvalidValue, arg)
	}

	var val string

	for idx := 1; idx < len(name); idx++ {
		if name[idx] == '=' || name[idx] == ' ' {
			val = strings.TrimSpace(name[idx+1:])
			name = name[0:idx]

			break
		}
	}

	if val == "" && numMinuses == 1 && len(arg) > 2 {
		name, val = name[:1], name[1:]
	}

	return name, val, nil
}
