package arg_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/test"
	"gitoa.ru/go-4devs/config/test/require"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	args := []string{
		"--listen=8080",
		"--config=config.hcl",
		"--url=http://4devs.io",
		"--url=https://4devs.io",
		"--timeout=1m",
		"--timeout=1h",
		"--start-at=2010-01-02T15:04:05Z",
		"--end-after=2009-01-02T15:04:05Z",
		"--end-after=2008-01-02T15:04:05+03:00",
	}
	read := []test.Read{
		test.NewRead(8080, "listen"),
		test.NewRead("config.hcl", "config"),
		test.NewRead(test.Time("2010-01-02T15:04:05Z"), "start-at"),
		test.NewReadUnmarshal(&[]string{"http://4devs.io", "https://4devs.io"}, &[]string{}, "url"),
		test.NewReadUnmarshal(&[]Duration{{time.Minute}, {time.Hour}}, &[]Duration{}, "timeout"),
		test.NewReadUnmarshal(&[]time.Time{
			test.Time("2009-01-02T15:04:05Z"),
			test.Time("2008-01-02T15:04:05+03:00"),
		}, &[]time.Time{}, "end-after"),
	}

	prov := arg.New(arg.WithArgs(args))

	test.Run(t, prov, read)
}

func TestProviderBind(t *testing.T) {
	t.Parallel()

	args := []string{
		"-p 8080",
		"--config=config.hcl",
		"-u http://4devs.io",
		"--url=https://4devs.io",
		"-t 1m",
		"--timeout=1h",
		"--start-at=2010-01-02T15:04:05Z",
		"--end-after=2009-01-02T15:04:05Z",
		"--end-after=2008-01-02T15:04:05+03:00",
	}

	read := []test.Read{
		test.NewRead(8080, "listen"),
		test.NewRead(test.Time("2010-01-02T15:04:05Z"), "start-at"),
		test.NewReadUnmarshal(&[]string{"http://4devs.io", "https://4devs.io"}, &[]string{}, "url"),
		test.NewReadUnmarshal(&[]Duration{{time.Minute}, {time.Hour}}, &[]Duration{}, "timeout"),
		test.NewReadUnmarshal(&[]time.Time{
			test.Time("2009-01-02T15:04:05Z"),
			test.Time("2008-01-02T15:04:05+03:00"),
		}, &[]time.Time{}, "end", "after"),
	}

	ctx := context.Background()
	prov := arg.New(arg.WithArgs(args))
	require.NoError(t, prov.Bind(ctx, testVariables(t)))

	test.Run(t, prov, read)
}

func testVariables(t *testing.T) config.Variables {
	t.Helper()

	def := definition.New()
	def.Add(
		option.Int("listen", "listen", option.Short('p')),
		option.String("config", "config"),
		option.String("url", "url", option.Short('u'), option.Slice),
		option.Duration("timeout", "timeout", option.Short('t'), option.Slice),
		option.Time("start-at", "start at"),
		group.New("end", "end at",
			option.Time("after", "after", option.Slice),
			proto.New("service", "service after",
				option.Time("after", "after"),
			),
		),
	)

	return config.NewVars(def.Options()...)
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(in []byte) error {
	o, err := time.ParseDuration(strings.Trim(string(in), `"`))
	if err != nil {
		return fmt.Errorf("parse:%w", err)
	}

	d.Duration = o

	return nil
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "%q", d), nil
}
