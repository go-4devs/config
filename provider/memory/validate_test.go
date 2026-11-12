package memory_test

import (
	"context"
	"errors"
	"strings"
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

func (brokenProvider) Bind(context.Context, config.Variables) error {
	return nil
}

type bindFail struct{}

func (bindFail) Name() string { return "bind-fail" }

func (bindFail) Value(context.Context, ...string) (config.Value, error) {
	return nil, config.ErrNotFound
}

func (bindFail) Bind(context.Context, config.Variables) error {
	return errBroken
}

func TestValidBindRejectsBlank(t *testing.T) {
	t.Parallel()

	store := memory.NewMap("mem")
	store.SetOption("", "host")

	err := memory.BindValid(store).Bind(context.Background(), config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	))
	require.ErrorIs(t, err, validator.ErrNotBlank)

	if strings.Count(err.Error(), "[host]") != 1 {
		t.Fatalf("key: %s", err)
	}

	store.SetOption("localhost", "host")

	require.NoError(t, memory.BindValid(store).Bind(context.Background(), config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	)))
}

func TestValidBindReadError(t *testing.T) {
	t.Parallel()

	vars := config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	)

	require.NoError(t, memory.BindValid(memory.NewMap("mem")).Bind(context.Background(), vars))
	require.ErrorIs(t, memory.BindValid(brokenProvider{}).Bind(context.Background(), vars), errBroken)
}

func TestValidBindRunsInnerBind(t *testing.T) {
	t.Parallel()

	blank := config.NewVars(
		option.String("host", "host", option.Default(""), validator.Valid(validator.NotBlank)),
	)
	require.ErrorIs(t, memory.BindValid(&memory.Default{}).Bind(context.Background(), blank), validator.ErrNotBlank)

	vars := config.NewVars(
		option.String("host", "host", validator.Valid(validator.NotBlank)),
	)
	require.ErrorIs(t, memory.BindValid(bindFail{}).Bind(context.Background(), vars), errBroken)
}
