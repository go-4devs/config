package proto

import (
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/param"
)

var _ config.Group = New("", "")

// New returns a group whose name is a {name} wildcard.
func New(name string, desc string, opts ...config.Option) Proto {
	return Proto{
		name:   key.Wild(name),
		opts:   opts,
		Params: param.New(param.WithDescription(desc)),
	}
}

type Proto struct {
	param.Params

	opts []config.Option
	name string
}

// Options returns the child options.
func (p Proto) Options() []config.Option {
	return p.opts
}

// Name returns the wildcard name.
func (p Proto) Name() string {
	return p.name
}
