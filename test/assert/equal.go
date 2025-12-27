package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected any, actual any, msgAndArgs ...any) bool {
	t.Helper()

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	t.Errorf("not equal expect:%v actual:%v", expected, actual)
	t.Error(msgAndArgs...)

	return false
}

func Equalf(t *testing.T, expected any, actual any, msg string, args ...any) bool {
	t.Helper()

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	t.Errorf("not equal expect:%#v acctual: %#v", expected, actual)
	t.Errorf(msg, args...)

	return false
}
