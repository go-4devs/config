package watcher

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-4devs/config"
)

var (
	_ config.Provider      = (*Provider)(nil)
	_ config.WatchProvider = (*Provider)(nil)
)

const defaultInterval = 10 * time.Minute

// New returns a provider that polls the inner provider every 10 minutes.
func New(provider config.Provider, opts ...Option) *Provider {
	prov := &Provider{
		Provider: provider,
		duration: defaultInterval,
		logger:   slog.ErrorContext,
	}

	for _, opt := range opts {
		opt(prov)
	}

	return prov
}

// WithDuration sets the poll interval. A duration that is not positive is replaced with 10 minutes.
func WithDuration(duration time.Duration) Option {
	if duration <= 0 {
		duration = defaultInterval
	}

	return func(p *Provider) {
		p.duration = duration
	}
}

// WithLogger sets the function used when a poll fails.
func WithLogger(l func(context.Context, string, ...any)) Option {
	return func(p *Provider) {
		p.logger = l
	}
}

type Option func(*Provider)

type Provider struct {
	config.Provider

	duration time.Duration
	logger   func(context.Context, string, ...any)
}

// Watch polls the value and calls callback when it changes. ErrStopWatch stops the poll.
func (p *Provider) Watch(ctx context.Context, callback config.WatchCallback, key ...string) error {
	old, err := p.Value(ctx, key...)
	if err != nil {
		return fmt.Errorf("failed watch variable: %w", err)
	}

	go func(oldVar config.Value) {
		ticker := time.NewTicker(p.duration)

		defer func() {
			ticker.Stop()
		}()

		for {
			select {
			case <-ticker.C:
				newVar, err := p.Value(ctx, key...)
				if err != nil {
					p.logger(ctx, "get value%v:%v", key, err.Error())
				} else if !newVar.IsEquals(oldVar) {
					if err := callback(ctx, oldVar, newVar); err != nil {
						if errors.Is(err, config.ErrStopWatch) {
							return
						}

						p.logger(ctx, "callback %v:%v", key, err)
					}

					oldVar = newVar
				}
			case <-ctx.Done():
				return
			}
		}
	}(old)

	return nil
}
