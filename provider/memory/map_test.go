package memory_test

import (
	"context"
	"sync"
	"testing"

	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
)

func TestAppendOptionConcurrent(t *testing.T) {
	t.Parallel()

	store := memory.NewMap("mem")

	const n = 32

	var wg sync.WaitGroup

	for range n {
		wg.Go(func() {
			if err := store.AppendOption("a", "host"); err != nil {
				t.Error(err)
			}
		})
	}

	wg.Wait()

	got, err := store.Value(context.Background(), "host")
	require.NoError(t, err)

	var items []string

	require.NoError(t, got.Unmarshal(&items))
	require.Equal(t, n, len(items))
}
