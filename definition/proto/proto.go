package proto

import (
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
)

var (
	_ config.Group = New("", "")
)

func New(name string, desc string, opts ...config.Option) Proto {
	return Proto{
		name:   name,
		opts:   opts,
		Params: param.New(option.Description(desc)),
	}
}

type Proto struct {
	param.Params

	opts []config.Option
	name string
}

func (p Proto) Options() []config.Option {
	return p.opts
}

func (p Proto) Name() string {
	return "{" + p.name + "}"
}
