package option

import (
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

var _ config.Option = New("", "", nil)

func New(name, desc string, vtype any, opts ...param.Option) Option {
	opts = append(opts, Description(desc), WithType(vtype))
	res := Option{
		name:   name,
		Params: param.New(opts...),
	}

	return res
}

type Option struct {
	param.Params

	name string
}

func (o Option) Name() string {
	return o.name
}

func String(name, description string, opts ...param.Option) Option {
	return New(name, description, "", opts...)
}

func Bool(name, description string, opts ...param.Option) Option {
	return New(name, description, false, opts...)
}

func Duration(name, description string, opts ...param.Option) Option {
	return New(name, description, time.Duration(0), opts...)
}

func Float64(name, description string, opts ...param.Option) Option {
	return New(name, description, float64(0), opts...)
}

func Int(name, description string, opts ...param.Option) Option {
	return New(name, description, int(0), opts...)
}

func Int64(name, description string, opts ...param.Option) Option {
	return New(name, description, int64(0), opts...)
}

func Time(name, description string, opts ...param.Option) Option {
	return New(name, description, time.Time{}, opts...)
}

func Uint(name, description string, opts ...param.Option) Option {
	return New(name, description, uint(0), opts...)
}

func Uint64(name, descriontion string, opts ...param.Option) Option {
	return New(name, descriontion, uint64(0), opts...)
}
