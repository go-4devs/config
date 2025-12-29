package group

import (
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

var _ config.Group = New("", "")

func New(name, desc string, opts ...config.Option) *Group {
	group := Group{
		name:   name,
		opts:   opts,
		Params: param.New(option.Description(desc)),
	}

	return &group
}

type Group struct {
	param.Params

	name string
	opts []config.Option
}

func (g *Group) Name() string {
	return g.name
}

func (g *Group) Options() []config.Option {
	return g.opts
}

func (g *Group) Add(opts ...config.Option) {
	g.opts = append(g.opts, opts...)
}

func (g *Group) With(opts ...param.Option) *Group {
	group := New(g.Name(), "")
	group.Params = param.Chain(g.Params, param.New(opts...))

	return group
}
