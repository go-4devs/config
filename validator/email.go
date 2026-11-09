package validator

import (
	"fmt"
	"net/mail"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

// ErrEmail is returned when the text is not an email address.
var ErrEmail = fmt.Errorf("%w: email", ErrInvalid)

// EmailError is returned when Got is not an email address.
type EmailError struct {
	Got string
}

// Error returns the received text.
func (e EmailError) Error() string {
	return fmt.Sprintf("%s: want user@domain, given %v", ErrEmail, e.Got)
}

// Unwrap returns ErrEmail.
func (e EmailError) Unwrap() error {
	return ErrEmail
}

// Email rejects text that is not an email address.
// The option type must be string. The whole text must be the address.
func Email(fn param.Params, in config.Value) error {
	if _, ok := param.Type(fn).(string); !ok {
		return ErrInvalid
	}

	text, err := in.ParseString()
	if err != nil {
		return fmt.Errorf("email:%w", err)
	}

	addr, err := mail.ParseAddress(text)
	if err != nil || addr.Address != text {
		return EmailError{Got: text}
	}

	return nil
}
