package require

import "testing"

func Fail(t *testing.T, msg string, args ...any) {
	t.Helper()
	t.Errorf(msg, args...)
	t.FailNow()
}
