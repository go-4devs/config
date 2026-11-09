package validator

import (
	"errors"
	"fmt"
)

// ErrInvalid is the base validation error. Validator sentinels wrap it.
var ErrInvalid = errors.New("invalid value")

// ErrNotBlank is returned when the value is empty.
var ErrNotBlank = fmt.Errorf("%w: not blank", ErrInvalid)
