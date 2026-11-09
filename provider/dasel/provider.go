package dasel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/value"
	dasel "github.com/tomwright/dasel/v3"
	"github.com/tomwright/dasel/v3/model"
	"github.com/tomwright/dasel/v3/parsing"
)

var _ config.Provider = (*Provider)(nil)

const Name = "dasel"

type Option func(*Provider)

// WithName sets the provider name.
func WithName(in string) Option {
	return func(p *Provider) {
		p.name = in
	}
}

// New returns a provider for the document. A read error is ErrInitFactory.
func New(in []byte, format parsing.Format, opts ...Option) (Provider, error) {
	reader, err := format.NewReader(parsing.DefaultReaderOptions())
	if err != nil {
		return Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, err)
	}

	data, verr := reader.Read(in)
	if verr != nil {
		return Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, verr)
	}

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

	return prov, nil
}

type Provider struct {
	data *model.Value
	key  func(path ...string) string
	name string
}

// Value returns the single value selected by the dotted path.
// A missing key is ErrValueNotFound.
func (p Provider) Value(ctx context.Context, path ...string) (config.Value, error) {
	selector := p.key(path...)

	data, cnt, err := dasel.Query(ctx, p.data, selector)

	missing := model.MapKeyNotFound{Key: ""}
	if errors.As(err, &missing) {
		return nil, fmt.Errorf("%s:%w", p.Name(), config.ErrValueNotFound)
	}

	if err != nil {
		return nil, fmt.Errorf("query: %w:%w", config.ErrInvalidValue, err)
	}

	if cnt > 1 {
		return nil, fmt.Errorf("count: %v:%w", cnt, config.ErrToManyArgs)
	}

	if cnt == 0 {
		return nil, fmt.Errorf("%s:%w", p.Name(), config.ErrValueNotFound)
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

// Name returns the provider name.
func (p Provider) Name() string {
	return p.name
}
