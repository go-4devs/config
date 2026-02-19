package memory

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

func WithParamProcess(process config.ProcessFunc) func(*Param) {
	return func(p *Param) {
		p.process = process
	}
}

func NewParam(
	name string,
	resolve func(param.Params) (any, bool),
	opts ...func(*Param),
) *Param {
	param := Param{
		name: name,
		process: func(_ context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
			return in, nil
		},
		data:  NewMap(name),
		param: resolve,
	}

	for _, opt := range opts {
		opt(&param)
	}

	return &param
}

type Param struct {
	process config.ProcessFunc
	name    string
	data    *Map
	param   func(p param.Params) (any, bool)
}

func (p *Param) Value(ctx context.Context, key ...string) (config.Value, error) {
	v, err := p.data.Value(ctx, key...)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	pv, perr := p.process.Process(ctx, v)
	if perr != nil {
		return nil, fmt.Errorf("%w", perr)
	}

	return pv, nil
}

func (p *Param) Bind(_ context.Context, def config.Variables) error {
	for _, opt := range def.Variables() {
		if data, ok := p.param(opt); ok {
			p.data.SetOption(data, opt.Key()...)
		}
	}

	return nil
}

func (p *Param) Name() string {
	return p.name
}
