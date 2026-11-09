package validator

import (
	"fmt"
	"regexp"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

// ErrRegex is returned when the text does not match the pattern or the pattern is invalid.
var ErrRegex = fmt.Errorf("%w: regex", ErrInvalid)

// RegexError is returned when Got does not match Pattern, or Pattern does not compile.
type RegexError struct {
	Pattern string
	Got     string
	Err     error
}

// Error returns the pattern and the received text.
func (e RegexError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: want valid regular expression, given %v: %v", ErrRegex, e.Pattern, e.Err)
	}

	return fmt.Sprintf("%s: want %v, given %v", ErrRegex, e.Pattern, e.Got)
}

// Unwrap returns ErrRegex. A compile failure also wraps that error.
func (e RegexError) Unwrap() error {
	if e.Err == nil {
		return ErrRegex
	}

	return fmt.Errorf("%w: %w", ErrRegex, e.Err)
}

// Regex returns a validator that accepts text matching pattern.
// The option type must be string. The match must cover the whole text.
// An invalid pattern returns ErrRegex.
func Regex(pattern string) Validator {
	re, compErr := regexp.Compile(pattern)

	return func(fn param.Params, in config.Value) error {
		if compErr != nil {
			return RegexError{Pattern: pattern, Got: "", Err: compErr}
		}

		text, err := optionText("regex", fn, in)
		if err != nil {
			return err
		}

		span := re.FindStringIndex(text)
		if span == nil || span[0] != 0 || span[1] != len(text) {
			return RegexError{Pattern: pattern, Got: text, Err: nil}
		}

		return nil
	}
}
