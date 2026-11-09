package etcd //nolint:testpackage // notify is unexported and needs no etcd server

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-4devs/config"
	pb "go.etcd.io/etcd/api/v3/mvccpb"
	client "go.etcd.io/etcd/client/v3"
)

type stopWatch struct {
	client.KV
	client.Watcher

	stopped chan struct{}
	prefix  bool
}

func (s *stopWatch) Watch(ctx context.Context, key string, opts ...client.OpOption) client.WatchChan {
	s.prefix = client.IsOptsWithPrefix(opts)

	ch := make(chan client.WatchResponse, 1)
	ch <- client.WatchResponse{
		Header:          nil,
		Events:          []*client.Event{{Kv: &pb.KeyValue{Key: []byte(key), Value: []byte("ok")}}},
		CompactRevision: 0,
		Canceled:        false,
		Created:         false,
		CancelReason:    "",
	}

	go func() {
		<-ctx.Done()
		close(ch)
		close(s.stopped)
	}()

	return ch
}

func TestWatchStopsOnErrStopWatch(t *testing.T) {
	t.Parallel()

	stub := &stopWatch{
		KV:      nil,
		Watcher: nil,
		stopped: make(chan struct{}),
		prefix:  false,
	}
	prov := New("fdevs", "config", stub)

	err := prov.Watch(context.Background(), func(context.Context, config.Value, config.Value) error {
		return config.ErrStopWatch
	}, "port")
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-stub.stopped:
	case <-time.After(time.Second):
		t.Fatal("watch did not stop")
	}

	if stub.prefix {
		t.Fatal("prefix watch")
	}
}

func TestNotifySkipsOtherKey(t *testing.T) {
	t.Parallel()

	const key = "fdevs/config/port"

	p := New("fdevs", "config", nil)
	calls := 0

	err := p.notify(context.Background(), key, []*client.Event{{
		Kv: &pb.KeyValue{Key: []byte(key + "_extra"), Value: []byte(`"no"`)},
	}}, func(context.Context, config.Value, config.Value) error {
		calls++

		return nil
	})
	if !errors.Is(err, config.ErrValueNotFound) || calls != 0 {
		t.Fatalf("calls %d err %v", calls, err)
	}
}

func TestNotifyFindsKeyAndStops(t *testing.T) {
	t.Parallel()

	const key = "fdevs/config/port"

	p := New("fdevs", "config", nil)

	var got string

	err := p.notify(context.Background(), key, []*client.Event{
		{Kv: nil},
		{Kv: &pb.KeyValue{Key: []byte(key), Value: []byte("ok")}},
	}, func(_ context.Context, oldVar, newVar config.Value) error {
		if oldVar != nil {
			t.Fatal("old value")
		}

		got = newVar.String()

		return config.ErrStopWatch
	})
	if !errors.Is(err, config.ErrStopWatch) || got != "ok" {
		t.Fatalf("got %q err %v", got, err)
	}
}
