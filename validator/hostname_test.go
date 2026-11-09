package validator_test

import (
	"strings"
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestHostname(t *testing.T) {
	t.Parallel()

	opt := option.String("host", "host")

	require.NoError(t, validator.Hostname(opt, value.New("example.com")))
	require.NoError(t, validator.Hostname(opt, value.New("localhost")))
	require.NoError(t, validator.Hostname(opt, value.New("Example.COM")))
	require.NoError(t, validator.Hostname(opt, value.JString("192.168.0.1")))
	require.NoError(t, validator.Hostname(opt, value.New(strings.Repeat("a", 63)+".com")))
	require.ErrorIs(t, validator.Hostname(opt, value.New("")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("example.com.")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("example.com:80")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("-bad.com")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("bad-.com")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("a..b")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New(strings.Repeat("a", 64))), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(opt, value.New("https://example.com")), validator.ErrHostname)
	require.ErrorIs(t, validator.Hostname(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
