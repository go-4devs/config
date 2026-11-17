package memory_test

import (
	"context"
	"strings"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/value"
)

func TestParamPrefixDefault(t *testing.T) {
	t.Parallel()

	keys, ok := memory.ParamPrefix(param.New(memory.Prefix("extra", "node")))
	require.True(t, ok)
	require.Equal(t, []string{"extra", "node"}, keys)
}

func TestParamPrefixRead(t *testing.T) {
	t.Parallel()

	parent := &pathProvider{}
	prov := memory.NewPrefix(parent, memory.WithParamPrefix(func(param.Params) ([]string, bool) {
		return []string{"extra"}, true
	}))

	require.NoError(t, prov.Bind(context.Background(), config.NewVars(
		option.String("port", "port"),
	)))

	_, err := prov.Value(context.Background(), "port")
	require.NoError(t, err)
	require.Equal(t, []string{"extra", "port"}, parent.got)
}

type pathProvider struct {
	got []string
}

func (p *pathProvider) Name() string { return "path" }

func (p *pathProvider) Value(_ context.Context, path ...string) (config.Value, error) {
	p.got = path

	return value.String(strings.Join(path, "_")), nil
}
