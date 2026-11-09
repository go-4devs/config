package config_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
)

func TestWatchRequiresWatcher(t *testing.T) {
	t.Parallel()

	client, err := config.New(memory.NewMap("mem"))
	require.NoError(t, err)

	err = client.Watch(context.Background(), nil, "host")
	require.ErrorIs(t, err, config.ErrNotFound)
}

func TestNewRejectsDuplicateName(t *testing.T) {
	t.Parallel()

	_, err := config.New(memory.NewMap("mem"), memory.NewMap("mem"))
	require.ErrorIs(t, err, config.ErrDuplicateName)
}
