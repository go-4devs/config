package handler

import (
	"context"
	"fmt"
	"log"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/param"
)

type pkey uint8

const (
	processorKey pkey = iota + 1
)

func Process(fn config.Processor) param.Option {
	return func(p param.Params) param.Params {
		return param.With(p, processorKey, fn)
	}
}

func getProcess(in param.Params) (config.Processor, bool) {
	p, ok := in.Param(processorKey)
	if !ok {
		return nil, false
	}

	data, tok := p.(config.Processor)

	return data, tok
}

func Processor(parent config.Provider) *ProcessHandler {
	handler := &ProcessHandler{
		Provider: parent,
		name:     "process:" + parent.Name(),
		idx:      key.Map{},
		process:  nil,
	}

	return handler
}

type ProcessHandler struct {
	config.Provider

	idx     key.Map
	process []config.Processor
	name    string
}

func (p *ProcessHandler) Name() string {
	return p.name
}

func (p *ProcessHandler) Bind(ctx context.Context, vars config.Variables) error {
	for _, one := range vars.Variables() {
		process, ok := getProcess(one)
		if !ok {
			continue
		}

		p.idx.Add(len(p.process), one.Key())
		p.process = append(p.process, process)
	}

	log.Print(p.idx.Index([]string{"group", "int"}))

	if bind, bok := p.Provider.(config.BindProvider); bok {
		berr := bind.Bind(ctx, vars)
		if berr != nil {
			return fmt.Errorf("%w", berr)
		}
	}

	return nil
}

func (p *ProcessHandler) Value(ctx context.Context, key ...string) (config.Value, error) {
	pval, perr := p.Provider.Value(ctx, key...)
	if perr != nil {
		return nil, fmt.Errorf("%w", perr)
	}

	idx, iok := p.idx.Index(key)
	if !iok {
		return pval, nil
	}

	prov := p.process[idx]

	val, err := prov.Process(ctx, pval)
	if err != nil {
		return nil, fmt.Errorf("process[%v]:%w", p.Name(), err)
	}

	return val, nil
}
