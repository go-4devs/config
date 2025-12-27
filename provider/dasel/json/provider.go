package json //nolint:revive

import (
	"fmt"

	"github.com/tomwright/dasel/v3/parsing"
	"github.com/tomwright/dasel/v3/parsing/json"
	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/dasel"
)

const Name = "dasel:json"

func New(data []byte) (dasel.Provider, error) {
	readOption := parsing.DefaultReaderOptions()

	reader, err := json.JSON.NewReader(readOption)
	if err != nil {
		return dasel.Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, err)
	}

	val, verr := reader.Read(data)
	if verr != nil {
		return dasel.Provider{}, fmt.Errorf("%w:%w", config.ErrInitFactory, verr)
	}

	return dasel.New(val, dasel.WithName(Name)), nil
}
