package hcl

import (
	"fmt"

	"github.com/go-4devs/config/provider/dasel"
	"github.com/tomwright/dasel/v3/parsing/hcl"
)

const Name = "dasel:hcl"

// New returns a dasel provider for HCL bytes.
func New(data []byte) (dasel.Provider, error) {
	prov, err := dasel.New(data, hcl.HCL, dasel.WithName(Name))
	if err != nil {
		return prov, fmt.Errorf("dasel: %w", err)
	}

	return prov, nil
}
