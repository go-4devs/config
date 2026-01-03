package key_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/test/require"
)

func TestWild(t *testing.T) {
	require.True(t, key.IsWild(key.Wild("test")))
	require.True(t, !key.IsWild("test"))
	require.True(t, key.IsWild("test", key.Wild("test"), "key"))
}
