package env_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/env"
	"github.com/go-4devs/config/test"
	"github.com/go-4devs/config/test/require"
)

func TestProvider_Prefix(t *testing.T) {
	t.Setenv("FDEVS_CONFIG_EXTRA_PORT", "8080")
	t.Setenv("FDEVS_CONFIG_OTHER_HOST", "6060")
	t.Setenv("FDEVS_CONFIG_EXTRA_NODE_GATE", "7070")
	t.Setenv("FDEVS_CONFIG_EXTRA_DB_PORT", "5050")
	t.Setenv("FDEVS_CONFIG_DSN", "plain")
	t.Setenv("FDEVS_CONFIG_PORT", "9090")
	t.Setenv("FDEVS_CONFIG_HOST", "1111")

	ctx := context.Background()
	provider := env.NewPrefixProvider(env.New("fdevs", "config"))

	require.NoError(t, provider.Bind(ctx, config.NewVars(
		option.Int("port", "port").With(env.Prefix("extra")),
		option.Int("host", "host").With(env.Prefix("other")),
		option.String("dsn", "dsn"),
		option.Int("gate", "gate").With(env.Prefix("extra", "node")),
		group.New("db", "db", option.Int("port", "port")).With(env.Prefix("extra")),
	)))
	test.Run(t, provider, []test.Read{
		test.NewRead(8080, "port"),
		test.NewRead(6060, "host"),
		test.NewRead("plain", "dsn"),
		test.NewRead(7070, "gate"),
		test.NewRead(5050, "db", "port"),
	})

	require.NoError(t, provider.Bind(ctx, config.NewVars(
		option.Int("port", "port"),
		option.Int("host", "host"),
	)))
	test.Run(t, provider, []test.Read{
		test.NewRead(9090, "port"),
		test.NewRead(1111, "host"),
	})
}
