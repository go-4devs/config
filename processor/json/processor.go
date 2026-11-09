package json

import (
	"context"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/value"
)

// JSON reads a string value as JSON text. A value that is not a string is returned unchanged.
func JSON(_ context.Context, in config.Value, _ ...param.Option) (config.Value, error) {
	data, ok := jsonText(in)
	if !ok {
		return in, nil
	}

	return value.JString(data), nil
}

func jsonText(in config.Value) (string, bool) {
	text, err := in.ParseString()
	if err != nil {
		return "", false
	}

	return text, true
}
