package validator

import (
	"fmt"
	"net/mail"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

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
		return NewError(ErrInvalid, text, "email")
	}

	return nil
}
