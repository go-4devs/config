package arg_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/test"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	args := os.Args

	defer func() {
		os.Args = args
	}()

	os.Args = []string{
		"main.go",
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

	prov := arg.New()

	test.Run(t, prov, read)
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
	return []byte(fmt.Sprintf("%q", d)), nil
}
