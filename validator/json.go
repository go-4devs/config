package validator

import (
	"encoding/json"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

// ErrJSON is returned when the text is not valid JSON.
var ErrJSON = fmt.Errorf("%w: json", ErrInvalid)

// JSONError is returned when Got is not JSON text.
type JSONError struct {
	Got string
}

// Error returns the received text.
func (e JSONError) Error() string {
	return fmt.Sprintf("%s: want json text, given %v", ErrJSON, e.Got)
}

// Unwrap returns ErrJSON.
func (e JSONError) Unwrap() error {
	return ErrJSON
}

// JSON rejects text that is not valid JSON.
// The option type must be string.
func JSON(fn param.Params, in config.Value) error {
	if _, ok := param.Type(fn).(string); !ok {
		return ErrInvalid
	}

	text, err := in.ParseString()
	if err != nil {
		return fmt.Errorf("json:%w", err)
	}

	if json.Valid([]byte(text)) {
		return nil
	}

	return JSONError{Got: text}
}
