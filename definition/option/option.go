package option

import (
	"reflect"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

var _ config.Option = New("", "", nil)

// New returns an option named name, described by desc, with value type vtype.
func New(name, desc string, vtype any, opts ...param.Option) Option {
	opts = append(opts, param.WithDescription(desc), WithType(vtype))

	if vtype != nil && isSlice(vtype) {
		opts = append(opts, Slice)
	}

	res := Option{
		name:   name,
		Params: param.New(opts...),
	}

	return res
}

func isSlice(vtype any) bool {
	rtype, ok := vtype.(reflect.Type)
	if !ok {
		rtype = reflect.TypeOf(vtype)
	}

	return rtype.Kind() == reflect.Slice
}

type Option struct {
	param.Params

	name string
}

// Name returns the option name.
func (o Option) Name() string {
	return o.name
}

// String returns a string option.
func String(name, description string, opts ...param.Option) Option {
	return New(name, description, "", opts...)
}

// Bool returns a bool option.
func Bool(name, description string, opts ...param.Option) Option {
	return New(name, description, false, opts...)
}

// Duration returns a time.Duration option.
func Duration(name, description string, opts ...param.Option) Option {
	return New(name, description, time.Duration(0), opts...)
}

// Float64 returns a float64 option.
func Float64(name, description string, opts ...param.Option) Option {
	return New(name, description, float64(0), opts...)
}

// Int returns an int option.
func Int(name, description string, opts ...param.Option) Option {
	return New(name, description, int(0), opts...)
}

// Int64 returns an int64 option.
func Int64(name, description string, opts ...param.Option) Option {
	return New(name, description, int64(0), opts...)
}

// Time returns a time.Time option.
func Time(name, description string, opts ...param.Option) Option {
	return New(name, description, time.Time{}, opts...)
}

// Uint returns a uint option.
func Uint(name, description string, opts ...param.Option) Option {
	return New(name, description, uint(0), opts...)
}

// Uint64 returns a uint64 option.
func Uint64(name, descriontion string, opts ...param.Option) Option {
	return New(name, descriontion, uint64(0), opts...)
}
