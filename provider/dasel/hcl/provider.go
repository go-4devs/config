package hcl

import (
	"github.com/tomwright/dasel/v3/parsing/hcl"
	"gitoa.ru/go-4devs/config/provider/dasel"
)

const Name = "dasel:hcl"

//nolint:wrapcheck
func New(data []byte) (dasel.Provider, error) {
	return dasel.New(data, hcl.HCL, dasel.WithName(Name))
}
