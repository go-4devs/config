package validator

import (
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/param"
)

func Chain(v ...Validator) Validator {
	return func(vr param.Params, in config.Value) error {
		for _, valid := range v {
			err := valid(vr, in)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

const paramValid = "param.valid"

type Validator func(param.Params, config.Value) error

func Valid(in ...Validator) param.Option {
	return func(v param.Params) param.Params {
		return param.With(v, paramValid, in)
	}
}

func Validate(key []string, fn param.Params, in config.Value) error {
	params, ok := fn.Param(paramValid)
	if !ok {
		return nil
	}

	valids, _ := params.([]Validator)

	for _, valid := range valids {
		err := valid(fn, in)
		if err != nil {
			return fmt.Errorf("%s:%w", key, err)
		}
	}

	return nil
}
