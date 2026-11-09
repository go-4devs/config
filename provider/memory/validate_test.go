package memory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
)

var errBroken = errors.New("broken")

type brokenProvider struct{}

func (brokenProvider) Name() string { return "broken" }

func (brokenProvider) Value(context.Context, ...string) (config.Value, error) {
	return nil, errBroken
}

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

func TestValidBindReadError(t *testing.T) {
	t.Parallel()

	vars := config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	)

	require.NoError(t, memory.Valid(memory.NewMap("mem")).Bind(context.Background(), vars))
	require.ErrorIs(t, memory.Valid(brokenProvider{}).Bind(context.Background(), vars), errBroken)
}
