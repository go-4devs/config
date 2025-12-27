package dasel

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	dasel "github.com/tomwright/dasel/v3"
	"github.com/tomwright/dasel/v3/model"
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
)

var _ config.Provider = New(nil)

const Name = "dasel"

type Option func(*Provider)

func WithName(in string) Option {
	return func(p *Provider) {
		p.name = in
	}
}

func New(data *model.Value, opts ...Option) Provider {
	prov := Provider{
		data: data,
		key: func(path ...string) string {
			return strings.Join(path, ".")
		},
		name: Name,
	}
	for _, opt := range opts {
		opt(&prov)
	}

	return prov
}

type Provider struct {
	data *model.Value
	key  func(path ...string) string
	name string
}

func (p Provider) Value(ctx context.Context, path ...string) (config.Value, error) {
	selector := p.key(path...)

	data, cnt, err := dasel.Query(ctx, p.data, selector)
	if err != nil {
		return nil, fmt.Errorf("query: %w:%w", config.ErrInvalidValue, err)
	}

	if cnt > 1 {
		return nil, fmt.Errorf("count: %v:%w", cnt, config.ErrToManyArgs)
	}

	if cnt == 0 {
		return value.EmptyValue(), nil
	}

	val, verr := data[0].GoValue()
	if verr != nil {
		return nil, fmt.Errorf("go value: %w:%w", config.ErrInvalidValue, verr)
	}

	res, merr := json.Marshal(val)
	if merr != nil {
		return nil, fmt.Errorf("marshal: %w:%w", config.ErrInvalidValue, merr)
	}

	return value.JBytes(res), nil
}

func (p Provider) Name() string {
	return p.name
}
