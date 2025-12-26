package validator

import (
	"slices"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

func Enum(enum ...string) Validator {
	return func(_ param.Params, in config.Value) error {
		val := in.String()
		if slices.Contains(enum, val) {
			return nil
		}

		return NewError(ErrInvalid, val, enum)
	}
}
