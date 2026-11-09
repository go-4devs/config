package config_test

import (
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
)

func TestNewRejectsDuplicateName(t *testing.T) {
	t.Parallel()

	_, err := config.New(memory.NewMap("mem"), memory.NewMap("mem"))
	require.ErrorIs(t, err, config.ErrDuplicateName)
}
