package hcl

import (
	"github.com/go-4devs/config/provider/dasel"
	"github.com/tomwright/dasel/v3/parsing/hcl"
)

const Name = "dasel:hcl"

//nolint:wrapcheck
func New(data []byte) (dasel.Provider, error) {
	return dasel.New(data, hcl.HCL, dasel.WithName(Name))
}
