package config_test

import (
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/test/require"
)

func TestNestedGroupKeysStayDistinct(t *testing.T) {
	t.Parallel()

	vars := config.NewVars(
		group.New("a", "a",
			group.New("b", "b",
				group.New("c", "c",
					option.String("x", "x"),
					option.String("y", "y"),
				),
			),
		),
	)

	_, err := vars.ByName("a", "b", "c", "x")
	require.NoError(t, err)

	_, err = vars.ByName("a", "b", "c", "y")
	require.NoError(t, err)
}
