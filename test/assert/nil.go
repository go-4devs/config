package assert

import "testing"

func Nil(t *testing.T, data any, msgAndArgs ...any) bool {
	t.Helper()

	if data != nil {
		t.Error(msgAndArgs...)

		return false
	}

	return true
}
