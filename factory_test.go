package config_test

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/factory"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/test/require"
)

var (
	errFactoryDown = errors.New("down")
	errHostValue   = errors.New("host")
)

func TestWrapFactoryRetriesAfterFailure(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32

	mem := memory.NewMap("mem")
	mem.SetOption("ok", "host")

	client, err := config.New(factory.New("late", func(context.Context, config.Provider) (config.Provider, error) {
		if calls.Add(1) == 1 {
			return nil, errFactoryDown
		}

		return mem, nil
	}))
	require.NoError(t, err)

	_, err = client.Value(context.Background(), "host")
	require.ErrorIs(t, err, errFactoryDown)

	got, err := client.Value(context.Background(), "host")
	require.NoError(t, err)
	require.Equal(t, "ok", got.String())
}

func TestWrapFactoryWaitsForInit(t *testing.T) {
	t.Parallel()

	mem := memory.NewMap("mem")
	mem.SetOption("ok", "host")

	entered := make(chan struct{})
	release := make(chan struct{})

	client, err := config.New(factory.New("late", func(context.Context, config.Provider) (config.Provider, error) {
		close(entered)
		<-release

		return mem, nil
	}))
	require.NoError(t, err)

	ctx := context.Background()

	const readers = 5

	errc := make(chan error, readers)

	var wg sync.WaitGroup

	wg.Go(func() {
		errc <- readHost(ctx, client)
	})

	<-entered

	for range readers - 1 {
		wg.Go(func() {
			errc <- readHost(ctx, client)
		})
	}

	close(release)
	wg.Wait()
	close(errc)

	for err := range errc {
		require.NoError(t, err)
	}
}

func readHost(ctx context.Context, client *config.Client) error {
	got, err := client.Value(ctx, "host")
	if err != nil {
		return err
	}

	if got.String() != "ok" {
		return errHostValue
	}

	return nil
}
