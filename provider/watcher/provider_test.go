package watcher_test

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/watcher"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

var _ config.Provider = (*provider)(nil)

type provider struct {
	cnt int32
}

func (p *provider) Name() string {
	return "test"
}

func (p *provider) Value(context.Context, ...string) (config.Value, error) {
	p.cnt++

	return value.JString(strconv.Itoa(int(p.cnt))), nil
}

func TestWatcher(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer func() {
		cancel()
	}()

	prov := &provider{}

	w := watcher.New(time.Second/3, prov)
	wg := sync.WaitGroup{}
	wg.Add(2)

	var cnt int32

	err := w.Watch(
		ctx,
		func(_ context.Context, _, _ config.Value) error {
			atomic.AddInt32(&cnt, 1)
			wg.Done()

			if atomic.LoadInt32(&cnt) == 2 {
				return config.ErrStopWatch
			}

			return nil
		},
		"tmpname",
	)

	wg.Wait()

	require.NoError(t, err)
	require.Equal(t, int32(2), cnt)
}
