package validator

import (
	"fmt"
	"regexp"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

// Regex returns a validator that accepts text matching pattern.
// The option type must be string. The match must cover the whole text.
// An invalid pattern returns ErrInvalid.
func Regex(pattern string) Validator {
	re, compErr := regexp.Compile(pattern)

	return func(fn param.Params, in config.Value) error {
		if compErr != nil {
			return fmt.Errorf("%w: %w", ErrInvalid, compErr)
		}

		text, err := optionText("regex", fn, in)
		if err != nil {
			return err
		}

		span := re.FindStringIndex(text)
		if span == nil || span[0] != 0 || span[1] != len(text) {
			return NewError(ErrInvalid, text, pattern)
		}

		return nil
	}
}
