package arg

import (
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
)

func Default(in any) param.Option {
	return option.Default(value.New(in))
}

func Required(v param.Params) {
	option.Required(v)
}

func Slice(v param.Params) {
	option.Slice(v)
}

func String(name, description string, opts ...param.Option) option.Option {
	return option.String(name, description, append(opts, Argument)...)
}

func Bool(name, description string, opts ...param.Option) option.Option {
	return option.Bool(name, description, append(opts, Argument)...)
}

func Duration(name, description string, opts ...param.Option) option.Option {
	return option.Duration(name, description, append(opts, Argument)...)
}

func Float64(name, description string, opts ...param.Option) option.Option {
	return option.Float64(name, description, append(opts, Argument)...)
}

func Int(name, description string, opts ...param.Option) option.Option {
	return option.Int(name, description, append(opts, Argument)...)
}

func Int64(name, description string, opts ...param.Option) option.Option {
	return option.Int64(name, description, append(opts, Argument)...)
}

func Time(name, description string, opts ...param.Option) option.Option {
	return option.Time(name, description, append(opts, Argument)...)
}

func Uint(name, description string, opts ...param.Option) option.Option {
	return option.Uint(name, description, append(opts, Argument)...)
}

func Uint64(name, descriontion string, opts ...param.Option) option.Option {
	return option.Uint64(name, descriontion, append(opts, Argument)...)
}

func Err(err error, key ...string) option.Error {
	return option.Err(err, key)
}
