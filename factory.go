package config

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

func WrapFactory(fn Factory, prov Provider) *WrapProvider {
	return &WrapProvider{
		factory: func(ctx context.Context) (Provider, error) {
			return fn(ctx, prov)
		},
		mu:       sync.Mutex{},
		done:     0,
		provider: nil,
	}
}

type WrapProvider struct {
	mu       sync.Mutex
	done     uint32
	provider Provider
	factory  func(ctx context.Context) (Provider, error)
}

func (p *WrapProvider) Watch(ctx context.Context, callback WatchCallback, path ...string) error {
	if err := p.init(ctx); err != nil {
		return fmt.Errorf("init read:%w", err)
	}

	watch, ok := p.provider.(WatchProvider)
	if !ok {
		return nil
	}

	if err := watch.Watch(ctx, callback, path...); err != nil {
		return fmt.Errorf("factory provider: %w", err)
	}

	return nil
}
func (p *WrapProvider) Value(ctx context.Context, path ...string) (Value, error) {
	if err := p.init(ctx); err != nil {
		return nil, fmt.Errorf("init read:%w", err)
	}

	variable, err := p.provider.Value(ctx, path...)
	if err != nil {
		return nil, fmt.Errorf("factory provider: %w", err)
	}

	return variable, nil
}

func (p *WrapProvider) Name() string {
	if err := p.init(context.Background()); err != nil {
		return fmt.Sprintf("%T", p.provider)
	}

	return p.provider.Name()
}

func (p *WrapProvider) Bind(ctx context.Context, data Variables) error {
	if err := p.init(ctx); err != nil {
		return fmt.Errorf("init bind: %w", err)
	}

	prov, ok := p.provider.(BindProvider)
	if !ok {
		return nil
	}

	if perr := prov.Bind(ctx, data); perr != nil {
		return fmt.Errorf("init bind provider: %w", perr)
	}

	return nil
}

func (p *WrapProvider) init(ctx context.Context) error {
	if atomic.LoadUint32(&p.done) == 0 {
		if !p.mu.TryLock() {
			return fmt.Errorf("%w", ErrInitFactory)
		}
		defer atomic.StoreUint32(&p.done, 1)
		defer p.mu.Unlock()

		var err error
		if p.provider, err = p.factory(ctx); err != nil {
			return fmt.Errorf("init provider factory:%w", err)
		}
	}

	return nil
}
