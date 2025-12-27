package json //nolint:revive

import (
	"github.com/tomwright/dasel/v3/parsing/json"
	"gitoa.ru/go-4devs/config/provider/dasel"
)

const Name = "dasel:json"

//nolint:wrapcheck
func New(data []byte) (dasel.Provider, error) {
	return dasel.New(data, json.JSON, dasel.WithName(Name))
}
