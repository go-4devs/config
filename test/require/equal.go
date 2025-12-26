package require

import (
	"testing"

	"gitoa.ru/go-4devs/config/test/assert"
)

func Equal(t *testing.T, expected any, actual any, msgAndArgs ...any) {
	t.Helper()

	if assert.Equal(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

func Equalf(t *testing.T, expected any, actual any, msg string, args ...any) {
	t.Helper()

	if assert.Equalf(t, expected, actual, msg, args...) {
		return
	}

	t.FailNow()
}
