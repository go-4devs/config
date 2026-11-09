package etcd_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/etcd"
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
	prov := etcd.New("fdevs", "config", stub)

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
