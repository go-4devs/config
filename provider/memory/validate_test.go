package memory_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
)

func TestValidBindRejectsBlank(t *testing.T) {
	t.Parallel()

	store := memory.NewMap("mem")
	store.SetOption("", "host")

	err := memory.Valid(store).Bind(context.Background(), config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	))
	require.ErrorIs(t, err, validator.ErrNotBlank)

	store.SetOption("localhost", "host")

	require.NoError(t, memory.Valid(store).Bind(context.Background(), config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	)))
}
