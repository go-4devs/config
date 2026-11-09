package validator

import (
	"fmt"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

const (
	maxHostnameLen = 253
	maxLabelLen    = 63
)

// ErrHostname is returned when the text is not a hostname.
var ErrHostname = fmt.Errorf("%w: hostname", ErrInvalid)

// HostnameError is returned when Got is not a hostname.
type HostnameError struct {
	Got string
}

// Error returns the received text.
func (e HostnameError) Error() string {
	return fmt.Sprintf("%s: want name or name.domain, given %v", ErrHostname, e.Got)
}

// Unwrap returns ErrHostname.
func (e HostnameError) Unwrap() error {
	return ErrHostname
}

// Hostname rejects text that is not a hostname.
// The option type must be string. Each label is 1-63 ASCII letters, digits, or hyphens
// and does not start or end with a hyphen. The whole name is at most 253 characters.
func Hostname(fn param.Params, in config.Value) error {
	text, err := optionText("hostname", fn, in)
	if err != nil {
		return err
	}

	if !hostnameOK(text) {
		return HostnameError{Got: text}
	}

	return nil
}

func hostnameOK(host string) bool {
	if host == "" || len(host) > maxHostnameLen {
		return false
	}

	for label := range strings.SplitSeq(host, ".") {
		if !labelOK(label) {
			return false
		}
	}

	return true
}

func labelOK(label string) bool {
	size := len(label)
	if size == 0 || size > maxLabelLen || label[0] == '-' || label[size-1] == '-' {
		return false
	}

	for idx := range size {
		if !hostChar(label[idx]) {
			return false
		}
	}

	return true
}

func hostChar(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '-'
}
