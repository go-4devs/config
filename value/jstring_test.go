package value_test

import (
	"math"
	"strconv"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/test/require"
	"gitoa.ru/go-4devs/config/value"
)

func TestJStringDuration(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		raw value.JString
		exp time.Duration
		err error
	}{
		"1m": {
			raw: value.JString("1m"),
			exp: time.Minute,
		},
		"number error": {
			raw: value.JString("100000000"),
			err: config.ErrInvalidValue,
		},
	}

	for name, data := range tests {
		require.Equal(t, data.exp, data.raw.Duration(), name)
		d, err := data.raw.ParseDuration()
		require.ErrorIsf(t, err, data.err, "%[1]s: expect:%#[2]v, got:%#[3]v", name, data.err, err)
		require.Equal(t, data.exp, d, name)
	}
}

func TestJStringInt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		raw value.JString
		exp int
		err error
	}{
		"empty": {
			raw: value.JString(strconv.Itoa(0)),
			exp: 0,
		},
		"42": {
			raw: value.JString("42"),
			exp: 42,
		},
		"err": {
			raw: value.JString("err"),
			exp: 0,
			err: config.ErrInvalidValue,
		},
		"float": {
			raw: value.JString("0.23"),
			exp: 0,
			err: config.ErrInvalidValue,
		},
		"maxInt": {
			raw: value.JString(strconv.Itoa(math.MaxInt)),
			exp: math.MaxInt,
		},
	}

	for name, data := range tests {
		require.Equal(t, data.exp, data.raw.Int(), name)
		res, err := data.raw.ParseInt()
		require.ErrorIs(t, err, data.err)
		require.Equal(t, data.exp, res, name)
	}
}
