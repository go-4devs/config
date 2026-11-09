package validator

import (
	"encoding/json"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

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

	return NewError(ErrInvalid, text, "json")
}
