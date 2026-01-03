package env_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/provider/env"
	"gitoa.ru/go-4devs/config/test"
	"gitoa.ru/go-4devs/config/test/require"
)

func TestProvider(t *testing.T) {
	t.Setenv("FDEVS_CONFIG_DSN", test.DSN)
	t.Setenv("FDEVS_CONFIG_PORT", "8080")

	provider := env.New("fdevs", "config")

	read := []test.Read{
		test.NewRead(test.DSN, "dsn"),
		test.NewRead(8080, "port"),
	}
	test.Run(t, provider, read)
}

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

func testOptions(t *testing.T) config.Options {
	t.Helper()

	return group.New("test", "test",
		group.New("log", "configure log",
			option.String("level", "level", option.Default("info")),
			proto.New("service", "configure log service",
				option.String("level", "level"),
			),
		),
	)
}
