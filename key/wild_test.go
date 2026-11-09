package key_test

import (
	"testing"

	"github.com/go-4devs/config/key"
	"github.com/go-4devs/config/test/require"
)

func TestWild(t *testing.T) {
	t.Parallel()

	require.True(t, key.IsWild(key.Wild("test")))
	require.True(t, !key.IsWild("test"))
	require.True(t, key.IsWild("test", key.Wild("test"), "key"))
}
