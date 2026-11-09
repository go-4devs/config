package arg

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

// Default stores the default value.
func Default(in any) param.Option {
	return option.Default(value.New(in))
}

// Required marks the option as required.
func Required(v param.Params) param.Params {
	return option.Required(v)
}

// Slice marks the option as a slice.
func Slice(v param.Params) param.Params {
	return option.Slice(v)
}

// String returns a string option and marks it as a positional argument.
func String(name, description string, opts ...param.Option) option.Option {
	return option.String(name, description, append(opts, Argument)...)
}

// Bool returns a bool option and marks it as a positional argument.
func Bool(name, description string, opts ...param.Option) option.Option {
	return option.Bool(name, description, append(opts, Argument)...)
}

// Duration returns a duration option and marks it as a positional argument.
func Duration(name, description string, opts ...param.Option) option.Option {
	return option.Duration(name, description, append(opts, Argument)...)
}

// Float64 returns a float64 option and marks it as a positional argument.
func Float64(name, description string, opts ...param.Option) option.Option {
	return option.Float64(name, description, append(opts, Argument)...)
}

// Int returns an int option and marks it as a positional argument.
func Int(name, description string, opts ...param.Option) option.Option {
	return option.Int(name, description, append(opts, Argument)...)
}

// Int64 returns an int64 option and marks it as a positional argument.
func Int64(name, description string, opts ...param.Option) option.Option {
	return option.Int64(name, description, append(opts, Argument)...)
}

// Time returns a time option and marks it as a positional argument.
func Time(name, description string, opts ...param.Option) option.Option {
	return option.Time(name, description, append(opts, Argument)...)
}

// Uint returns a uint option and marks it as a positional argument.
func Uint(name, description string, opts ...param.Option) option.Option {
	return option.Uint(name, description, append(opts, Argument)...)
}

// Uint64 returns a uint64 option and marks it as a positional argument.
func Uint64(name, descriontion string, opts ...param.Option) option.Option {
	return option.Uint64(name, descriontion, append(opts, Argument)...)
}

// Err returns an option error tagged with key.
func Err(err error, key ...string) option.Error {
	return option.Err(err, key)
}
