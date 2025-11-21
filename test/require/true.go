package require

import (
	"testing"
)

func Truef(t *testing.T, value bool, msg string, args ...any) {
	t.Helper()

	if !value {
		t.Errorf(msg, args...)
		t.FailNow()
	}
}
