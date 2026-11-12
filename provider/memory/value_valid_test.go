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

func TestValueValidRejectsBlank(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	store := memory.NewMap("mem")
	store.SetOption("", "host")
	store.SetOption("app", "name")

	prov := memory.ValueValid(store)
	require.NoError(t, prov.Bind(ctx, config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
		option.String("name", "name"),
	)))

	_, err := prov.Value(ctx, "host")
	require.ErrorIs(t, err, validator.ErrNotBlank)

	val, err := prov.Value(ctx, "name")
	require.NoError(t, err)
	require.Equal(t, "app", val.String())

	_, err = prov.Value(ctx, "missing")
	require.ErrorIs(t, err, config.ErrNotFound)
}
