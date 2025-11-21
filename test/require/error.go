package require

import (
	"testing"
)

func NoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()

	if err != nil {
		t.Error(msgAndArgs...)
		t.FailNow()
	}
}

func NoErrorf(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err != nil {
		t.Errorf(msg, args...)
		t.FailNow()
	}
}
