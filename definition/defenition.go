package definition

import (
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

var _ config.Options = (*Definition)(nil)

// New returns a definition that holds opts.
func New(opts ...config.Option) *Definition {
	return &Definition{
		options: opts,
		Params:  param.New(),
	}
}

type Definition struct {
	param.Params

	options []config.Option
}

// Add appends options.
func (d *Definition) Add(opts ...config.Option) {
	d.options = append(d.options, opts...)
}

// Options returns the options.
func (d *Definition) Options() []config.Option {
	return d.options
}

// With returns a copy whose parameters are chained with params.
func (d *Definition) With(params param.Params) *Definition {
	def := &Definition{
		options: make([]config.Option, len(d.options)),
		Params:  param.Chain(params, d.Params),
	}

	copy(def.options, d.options)

	return def
}
