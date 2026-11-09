package etcd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/value"
	pb "go.etcd.io/etcd/api/v3/mvccpb"
	client "go.etcd.io/etcd/client/v3"
)

const (
	Name      = "etcd"
	Separator = "/"
)

var (
	_ config.Provider      = (*Provider)(nil)
	_ config.WatchProvider = (*Provider)(nil)
)

type Client interface {
	client.KV
	client.Watcher
}

// WithName sets the provider name.
func WithName(name string) func(*Provider) {
	return func(p *Provider) {
		p.name = name
	}
}

// WithLog sets the function used for watch errors.
func WithLog(fn func(context.Context, string, ...any)) func(*Provider) {
	return func(p *Provider) {
		p.log = fn
	}
}

// WithPrefix sets the key prefix.
func WithPrefix(prefix string) func(*Provider) {
	return func(p *Provider) {
		p.prefix = prefix
	}
}

// WithKey sets the function that joins a path into an etcd key.
func WithKey(fn func(...string) string) func(*Provider) {
	return func(p *Provider) {
		p.key = fn
	}
}

// New returns an etcd provider for namespace and appName.
func New(namespace, appName string, client Client, opts ...func(*Provider)) *Provider {
	prov := Provider{
		client: client,
		key: func(s ...string) string {
			return strings.Join(s, Separator)
		},
		name:   Name,
		prefix: namespace + Separator + appName,
		log: func(_ context.Context, format string, args ...any) {
			log.Printf(format, args...)
		},
	}

	for _, opt := range opts {
		opt(&prov)
	}

	return &prov
}

type Provider struct {
	client Client
	key    func(...string) string
	name   string
	prefix string
	log    func(context.Context, string, ...any)
}

// Name returns the provider name.
func (p *Provider) Name() string {
	return p.name
}

// Key returns the full etcd key for the path.
func (p *Provider) Key(s []string) string {
	return p.prefix + Separator + p.key(s...)
}

// Value returns the raw string stored at the key.
func (p *Provider) Value(ctx context.Context, path ...string) (config.Value, error) {
	name := p.Key(path)

	resp, err := p.client.Get(ctx, name, client.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("%w: key:%s, prov:%s", err, name, p.Name())
	}

	val, err := p.resolve(name, resp.Kvs)
	if err != nil {
		return nil, fmt.Errorf("%w: key:%s, prov:%s", err, name, p.Name())
	}

	return val, nil
}

// Watch reports changes of the key. The callback is not called for other keys under the prefix. ErrStopWatch stops the watch.
func (p *Provider) Watch(ctx context.Context, callback config.WatchCallback, path ...string) error {
	go func(ctx context.Context, key string, callback config.WatchCallback) {
		watch := p.client.Watch(ctx, key, client.WithPrevKV(), client.WithPrefix())
		for w := range watch {
			if errors.Is(p.notify(ctx, key, w.Events, callback), config.ErrStopWatch) {
				return
			}
		}
	}(ctx, p.Key(path), callback)

	return nil
}

func (p *Provider) notify(ctx context.Context, key string, events []*client.Event, callback config.WatchCallback) error {
	kvs, olds := p.getEventKvs(events)

	newVar, err := p.resolve(key, kvs)
	if err != nil {
		return err
	}

	oldVar, err := p.resolve(key, olds)
	if err != nil {
		oldVar = nil
	}

	if err = callback(ctx, oldVar, newVar); err != nil {
		if errors.Is(err, config.ErrStopWatch) {
			return err
		}

		p.log(ctx, "watch callback[%v] %v:%v", p.Name(), key, err)
	}

	return nil
}

func (p *Provider) getEventKvs(events []*client.Event) ([]*pb.KeyValue, []*pb.KeyValue) {
	kvs := make([]*pb.KeyValue, 0, len(events))
	old := make([]*pb.KeyValue, 0, len(events))

	for i := range events {
		kvs = append(kvs, events[i].GetKv())
		old = append(old, events[i].GetPrevKv())
	}

	return kvs, old
}

func (p *Provider) resolve(key string, kvs []*pb.KeyValue) (config.Value, error) {
	for _, kv := range kvs {
		if kv == nil || string(kv.GetKey()) != key {
			continue
		}

		return value.JString(kv.GetValue()), nil
	}

	return nil, fmt.Errorf("%w: name %s", config.ErrValueNotFound, key)
}
