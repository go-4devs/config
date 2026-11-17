package memory

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
)

// ErrConflict is returned when two options resolve to the same key after prefixes.
var ErrConflict = errors.New("prefix conflict")

// CheckFn is called with the definition before the conflict check.
type CheckFn func(context.Context, *definition.Definition)

// CheckPrefix builds a definition through fn and reports prefix conflicts using prefix.
func CheckPrefix(fn CheckFn, prefix ParamPrefixFn) error {
	def := definition.New()
	fn(context.Background(), def)

	seen := map[string]string{}

	for _, variable := range config.NewVars(def.Options()...).Variables() {
		segs, _ := prefix(variable)
		full := strings.Join(slices.Concat(segs, variable.Key()), "/")
		path := strings.Join(variable.Key(), "/")

		if prev, ok := seen[full]; ok {
			return fmt.Errorf("%w: %s and %s -> %s", ErrConflict, prev, path, full)
		}

		seen[full] = path
	}

	return nil
}
