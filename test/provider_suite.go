package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/test/require"
)

const (
	DSN       = "pgsql://user@pass:127.0.0.1:5432"
	Namespace = "fdevs"
	AppName   = "config"
)

func Run(t *testing.T, provider config.Provider, read []Read) {
	t.Helper()

	ctx := context.Background()

	for idx, read := range read {
		t.Run(fmt.Sprintf("%v:%v", idx, read.Key), func(t *testing.T) {
			val, err := provider.Value(ctx, read.Key...)
			read.Error(t, err)
			read.Assert(t, val)
		})
	}
}

type Read struct {
	Key    []string
	Assert func(t *testing.T, v config.Value)
	Error  func(t *testing.T, ex error)
}

type Config struct {
	Duration time.Duration
	Enabled  bool
}

func NewReadConfig(key ...string) Read {
	ex := &Config{
		Duration: 21 * time.Minute,
		Enabled:  true,
	}

	return NewReadUnmarshal(ex, &Config{}, key...)
}

func NewReadUnmarshal(expected, target any, key ...string) Read {
	return Read{
		Key: key,
		Assert: func(t *testing.T, v config.Value) {
			t.Helper()
			require.NoErrorf(t, v.Unmarshal(target), "unmarshal")
			require.Equal(t, expected, target, "unmarshal")
		},
		Error: exError(key...),
	}
}

func exError(path ...string) func(t *testing.T, err error) {
	return func(t *testing.T, err error) {
		t.Helper()
		require.NoError(t, err, path)
	}
}

func Time(value string) time.Time {
	t, _ := time.Parse(time.RFC3339, value)

	return t
}

// NewRead test data.
func NewRead(expected any, key ...string) Read {
	return Read{
		Key:   key,
		Error: exError(key...),
		Assert: func(t *testing.T, v config.Value) {
			t.Helper()

			var (
				val   any
				err   error
				short any
			)

			switch expected.(type) {
			case bool:
				val, err = v.ParseBool()
				short = v.Bool()
			case int:
				val, err = v.ParseInt()
				short = v.Int()
			case int64:
				val, err = v.ParseInt64()
				short = v.Int64()
			case uint:
				val, err = v.ParseUint()
				short = v.Uint()
			case uint64:
				val, err = v.ParseUint64()
				short = v.Uint64()
			case string:
				val, err = v.ParseString()
				short = v.String()
			case float64:
				val, err = v.ParseFloat64()
				short = v.Float64()
			case time.Duration:
				val, err = v.ParseDuration()
				short = v.Duration()
			case time.Time:
				val, err = v.ParseTime()
				short = v.Time()
			default:
				require.Fail(t, "unexpected type:%+T", expected)
			}

			require.Equalf(t, val, short, "%q!=%q, type:%T", val, short, expected)
			require.NoErrorf(t, err, "err:%v type:%T", err, expected)
			require.Equalf(t, expected, val, "%q!=%q type:%T", expected, val, expected)
		},
	}
}

func NewErrorIs(exErr error, path ...string) Read {
	return Read{
		Key:    path,
		Assert: func(*testing.T, config.Value) {},
		Error: func(t *testing.T, err error) {
			t.Helper()
			require.ErrorIs(t, err, exErr)
		},
	}
}
