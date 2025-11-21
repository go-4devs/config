package definition

import (
	"gitoa.ru/go-4devs/config"
)

func New() *Definition {
	return &Definition{
		options: nil,
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
