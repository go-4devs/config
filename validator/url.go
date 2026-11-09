package validator

import (
	"fmt"
	"net/url"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
)

var (
	// ErrURL is returned when the text is not an absolute URL.
	ErrURL = fmt.Errorf("%w: url", ErrInvalid)
	// ErrURI is returned when the text is not a URI.
	ErrURI = fmt.Errorf("%w: uri", ErrInvalid)
)

// URLError is returned when Got is not an absolute URL.
type URLError struct {
	Got string
}

// Error returns the received text.
func (e URLError) Error() string {
	return fmt.Sprintf("%s: want scheme://host, given %v", ErrURL, e.Got)
}

// Unwrap returns ErrURL.
func (e URLError) Unwrap() error {
	return ErrURL
}

// URIError is returned when Got is not a URI.
type URIError struct {
	Got string
}

// Error returns the received text.
func (e URIError) Error() string {
	return fmt.Sprintf("%s: want /path or scheme:value, given %v", ErrURI, e.Got)
}

// Unwrap returns ErrURI.
func (e URIError) Unwrap() error {
	return ErrURI
}

// URL rejects text that is not an absolute URL.
// The option type must be string. The URL must include a scheme and a host.
func URL(fn param.Params, in config.Value) error {
	text, err := optionText("url", fn, in)
	if err != nil {
		return err
	}

	parsed, err := url.Parse(text)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return URLError{Got: text}
	}

	return nil
}

// URI rejects text that is not a URI.
// The option type must be string. An absolute path or a scheme with a value is enough.
func URI(fn param.Params, in config.Value) error {
	text, err := optionText("uri", fn, in)
	if err != nil {
		return err
	}

	parsed, err := url.ParseRequestURI(text)
	if err != nil || parsed.Host == "" && parsed.Path == "" && parsed.Opaque == "" {
		return URIError{Got: text}
	}

	return nil
}

func optionText(kind string, fn param.Params, in config.Value) (string, error) {
	if _, ok := param.Type(fn).(string); !ok {
		return "", ErrInvalid
	}

	text, err := in.ParseString()
	if err != nil {
		return "", fmt.Errorf("%s:%w", kind, err)
	}

	return text, nil
}
