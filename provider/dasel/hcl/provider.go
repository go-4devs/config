package hcl

import (
	"fmt"

	"github.com/tomwright/dasel/v3/parsing"
	"github.com/tomwright/dasel/v3/parsing/hcl"
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/dasel"
)

const Name = "dasel:hcl"

func New(data []byte) (dasel.Provider, error) {
	readOption := parsing.DefaultReaderOptions()

	reader, err := hcl.HCL.NewReader(readOption)
	if err != nil {
		return dasel.Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, err)
	}

	val, verr := reader.Read(data)
	if verr != nil {
		return dasel.Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, verr)
	}

	return dasel.New(val, dasel.WithName(Name)), nil
}
