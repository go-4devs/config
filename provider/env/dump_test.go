package env_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/group"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/definition/proto"
	"github.com/go-4devs/config/provider/env"
	"github.com/go-4devs/config/test/require"
)

func TestProvider_DumpReference(t *testing.T) {
	t.Parallel()

	const expect = `# configure log.
# level.
FDEVS_CONFIG_LOG_LEVEL=info
# configure log service.
# level.
#FDEVS_CONFIG_LOG_{SERVICE}_LEVEL=
`

	ctx := context.Background()
	prov := env.New("fdevs", "config")
	buf := bytes.NewBuffer(nil)

	require.NoError(t, prov.DumpReference(ctx, buf, testOptions(t)))
	require.Equal(t, buf.String(), expect)
}

func TestProvider_DumpReference_Prefix(t *testing.T) {
	t.Parallel()

	const expect = `# configure log.
# level.
FDEVS_CONFIG_EXTRA_LOG_LEVEL=info
# port.
#FDEVS_CONFIG_NODE_LOG_PORT=
# configure log service.
# level.
#FDEVS_CONFIG_EXTRA_LOG_{SERVICE}_LEVEL=
`

	ctx := context.Background()
	prov := env.NewHelper(env.New("fdevs", "config"))

	require.NoError(t, prov.Bind(ctx, config.NewVars(
		option.String("port", "port").With(env.Prefix("other")),
	)))

	buf := bytes.NewBuffer(nil)
	require.NoError(t, prov.DumpReference(ctx, buf, prefixOptions(t)))
	require.Equal(t, expect, buf.String())

	const each = `# port.
#FDEVS_CONFIG_EXTRA_PORT=
# host.
#FDEVS_CONFIG_OTHER_HOST=
`

	own := bytes.NewBuffer(nil)
	require.NoError(t, prov.DumpReference(ctx, own, definition.New(
		option.Int("port", "port").With(env.Prefix("extra")),
		option.Int("host", "host").With(env.Prefix("other")),
	)))
	require.Equal(t, each, own.String())
}

func TestProvider_DumpReference_PrefixSibling(t *testing.T) {
	t.Parallel()

	const expect = `# port.
#FDEVS_CONFIG_EXTRA_PORT=
# dsn.
#FDEVS_CONFIG_DSN=
`

	buf := bytes.NewBuffer(nil)
	prov := env.NewHelper(env.New("fdevs", "config"))

	require.NoError(t, prov.DumpReference(context.Background(), buf, definition.New(
		option.Int("port", "port").With(env.Prefix("extra")),
		option.String("dsn", "dsn"),
	)))
	require.Equal(t, expect, buf.String())
}

func TestProvider_DumpReference_PrefixMany(t *testing.T) {
	t.Parallel()

	const expect = `# gate.
#FDEVS_CONFIG_EXTRA_NODE_GATE=
# db.
# port.
#FDEVS_CONFIG_EXTRA_NODE_DB_PORT=
`

	buf := bytes.NewBuffer(nil)
	prov := env.NewHelper(env.New("fdevs", "config"))

	require.NoError(t, prov.DumpReference(context.Background(), buf, definition.New(
		option.Int("gate", "gate").With(env.Prefix("extra", "node")),
		group.New("db", "db", option.Int("port", "port")).With(env.Prefix("extra", "node")),
	)))
	require.Equal(t, expect, buf.String())
}

func prefixOptions(t *testing.T) *definition.Definition {
	t.Helper()

	return definition.New(
		group.New("log", "configure log",
			option.String("level", "level", option.Default("info")),
			option.String("port", "port").With(env.Prefix("node")),
			proto.New("service", "configure log service",
				option.String("level", "level"),
			),
		).With(env.Prefix("extra")),
	)
}

func testOptions(t *testing.T) *definition.Definition {
	t.Helper()

	return definition.New(
		group.New("log", "configure log",
			option.String("level", "level", option.Default("info")),
			proto.New("service", "configure log service",
				option.String("level", "level"),
			),
		),
	)
}
