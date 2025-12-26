package require

import (
	"errors"
	"testing"
)

func NoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()

	if err != nil {
		t.Errorf("no error got:%v", err)
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

func ErrorIs(t *testing.T, err, ex error, msgAndArgs ...any) {
	t.Helper()

	if !errors.Is(err, ex) {
		t.Errorf("expect:%#v got:%#v", ex, err)
		t.Error(msgAndArgs...)
		t.FailNow()
	}
}

func ErrorIsf(t *testing.T, err, ex error, msg string, args ...any) {
	t.Helper()

	if !errors.Is(ex, err) {
		t.Errorf(msg, args...)
		t.FailNow()
	}
}
