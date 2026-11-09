package yaml

import (
	"context"
	"fmt"
	"os"

	"github.com/go-4devs/config"
	"gopkg.in/yaml.v3"
)

const NameWatch = "yaml_watch"

// NewWatch returns a provider that reads the YAML file on every Value call.
func NewWatch(name string, opts ...Option) *Watch {
	wath := Watch{
		file: name,
		prov: create(opts...),
		name: NameWatch,
	}

	return &wath
}

type Watch struct {
	file string
	prov *Provider
	name string
}

// Name returns the provider name.
func (p *Watch) Name() string {
	return p.name
}

// Value reads the file and returns the value at path.
func (p *Watch) Value(ctx context.Context, path ...string) (config.Value, error) {
	in, err := os.ReadFile(p.file)
	if err != nil {
		return nil, fmt.Errorf("yaml_file: read error: %w", err)
	}

	var yNode yaml.Node
	if err = yaml.Unmarshal(in, &yNode); err != nil {
		return nil, fmt.Errorf("yaml_file: unmarshal error: %w", err)
	}

	return p.prov.With(&yNode).Value(ctx, path...)
}
