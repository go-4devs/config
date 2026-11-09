package json

import (
	"github.com/go-4devs/config/provider/dasel"
	"github.com/tomwright/dasel/v3/parsing/json"
)

const Name = "dasel:json"

//nolint:wrapcheck
func New(data []byte) (dasel.Provider, error) {
	return dasel.New(data, json.JSON, dasel.WithName(Name))
}
