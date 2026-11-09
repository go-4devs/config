package group

import (
	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

var _ config.Group = New("", "")

// New returns a group named name with description desc and child opts.
func New(name, desc string, opts ...config.Option) *Group {
	group := Group{
		name:   name,
		opts:   opts,
		Params: param.New(param.WithDescription(desc)),
	}

	return &group
}

type Group struct {
	param.Params

	name string
	opts []config.Option
}

// Name returns the group name.
func (g *Group) Name() string {
	return g.name
}

// Options returns the child options.
func (g *Group) Options() []config.Option {
	return g.opts
}

// Add appends child options.
func (g *Group) Add(opts ...config.Option) {
	g.opts = append(g.opts, opts...)
}

// With returns a copy of the group with extra parameters.
func (g *Group) With(opts ...param.Option) *Group {
	group := New(g.Name(), "")
	group.Params = param.Chain(g.Params, param.New(opts...))
	group.opts = make([]config.Option, len(g.opts))
	copy(group.opts, g.opts)

	return group
}
