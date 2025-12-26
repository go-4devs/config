package definition

import (
	"gitoa.ru/go-4devs/config"
)

func New(opts ...config.Option) *Definition {
	return &Definition{
		options: opts,
	}
}

type Definition struct {
	options []config.Option
}

func (d *Definition) Add(opts ...config.Option) {
	d.options = append(d.options, opts...)
}

func (d *Definition) Options() []config.Option {
	return d.options
}
