package validator_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
)

func TestURL(t *testing.T) {
	t.Parallel()

	opt := option.String("url", "url")

	require.NoError(t, validator.URL(opt, value.New("https://example.com/path")))
	require.NoError(t, validator.URL(opt, value.JString("http://example.com")))
	require.ErrorIs(t, validator.URL(opt, value.New("/path")), validator.ErrURL)
	require.ErrorIs(t, validator.URL(opt, value.New("example.com")), validator.ErrURL)
	require.ErrorIs(t, validator.URL(opt, value.New("mailto:user@example.com")), validator.ErrURL)
	require.ErrorIs(t, validator.URL(opt, value.New("")), validator.ErrURL)
	require.ErrorIs(t, validator.URL(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}

func TestURI(t *testing.T) {
	t.Parallel()

	opt := option.String("uri", "uri")

	require.NoError(t, validator.URI(opt, value.New("https://example.com/path")))
	require.NoError(t, validator.URI(opt, value.New("/path")))
	require.NoError(t, validator.URI(opt, value.JString("mailto:user@example.com")))
	require.ErrorIs(t, validator.URI(opt, value.New("example.com")), validator.ErrURI)
	require.ErrorIs(t, validator.URI(opt, value.New("http://")), validator.ErrURI)
	require.ErrorIs(t, validator.URI(opt, value.New("")), validator.ErrURI)
	require.ErrorIs(t, validator.URI(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
