package validator

import (
	"fmt"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

// ErrFileExists is returned when the path is not an existing file.
var ErrFileExists = fmt.Errorf("%w: file exists", ErrInvalid)

// FileExistsError is returned when Got is not an existing file.
type FileExistsError struct {
	Got string
	Err error
}

// Error returns the path and the reason it is not a file.
func (e FileExistsError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: want existing file, given %v: %v", ErrFileExists, e.Got, e.Err)
	}

	return fmt.Sprintf("%s: want existing file, given %v", ErrFileExists, e.Got)
}

// Unwrap returns ErrFileExists. A stat failure also wraps that error.
func (e FileExistsError) Unwrap() error {
	if e.Err == nil {
		return ErrFileExists
	}

	return fmt.Errorf("%w: %w", ErrFileExists, e.Err)
}

// FileExists rejects a path that is not an existing file.
// The option type must be string. A directory is rejected.
func FileExists(fn param.Params, in config.Value) error {
	text, err := optionText("file exists", fn, in)
	if err != nil {
		return err
	}

	info, err := os.Stat(text)
	if err != nil {
		return FileExistsError{Got: text, Err: err}
	}

	if info.IsDir() {
		return FileExistsError{Got: text, Err: nil}
	}

	return nil
}
