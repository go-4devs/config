package etcd //nolint:testpackage // notify is unexported and needs no etcd server

import (
	"context"
	"errors"
	"testing"

	"gitoa.ru/go-4devs/config"
	pb "go.etcd.io/etcd/api/v3/mvccpb"
	client "go.etcd.io/etcd/client/v3"
)

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
