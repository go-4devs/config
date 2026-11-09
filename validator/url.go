package validator

import (
	"fmt"
	"net/url"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

// URL rejects text that is not an absolute URL.
// The option type must be string. The URL must include a scheme and a host.
func URL(fn param.Params, in config.Value) error {
	text, err := optionText("url", fn, in)
	if err != nil {
		return err
	}

	parsed, err := url.Parse(text)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return NewError(ErrInvalid, text, "url")
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
		return NewError(ErrInvalid, text, "uri")
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
