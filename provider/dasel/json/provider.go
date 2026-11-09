package json

import (
	"fmt"

	"github.com/go-4devs/config/provider/dasel"
	"github.com/tomwright/dasel/v3/parsing/json"
)

const Name = "dasel:json"

// New returns a dasel provider for JSON bytes.
func New(data []byte) (dasel.Provider, error) {
	prov, err := dasel.New(data, json.JSON, dasel.WithName(Name))
	if err != nil {
		return prov, fmt.Errorf("dasel: %w", err)
	}

	return prov, nil
}
