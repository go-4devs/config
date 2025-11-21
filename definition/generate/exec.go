package generate

import (
	"context"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"

	"gitoa.ru/go-4devs/config/definition/generate/bootstrap"
)

func NewGConfig(fname string, opts ...Option) (Config, error) {
	opts = append([]Option{
		WithFile(fname),
	}, opts...)

	return NewConfig(opts...), nil
}

type GConfig interface {
	BuildTags() string
	OutName() string
	bootstrap.Config
}

func Generate(ctx context.Context, cfg GConfig) error {
	path, err := bootstrap.Bootstrap(ctx, cfg)
	defer os.Remove(path)

	if err != nil {
		return fmt.Errorf("build bootstrap:%w", err)
	}

	tmpFile, err := os.Create(cfg.File() + ".tmp")
	if err != nil {
		return fmt.Errorf("create tmp file:%w", err)
	}

	defer os.Remove(tmpFile.Name()) // will not remove after rename

	execArgs := []string{"run"}
	if len(cfg.BuildTags()) > 0 {
		execArgs = append(execArgs, "-tags", cfg.BuildTags())
	}

	execArgs = append(execArgs, filepath.Base(path))
	cmd := exec.CommandContext(ctx, "go", execArgs...)

	cmd.Stdout = tmpFile
	cmd.Stderr = os.Stderr

	cmd.Dir = filepath.Dir(path)
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("start cmd:%w", err)
	}

	tmpFile.Close()

	// format file and write to out path
	in, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	out, err := format.Source(in)
	if err != nil {
		return fmt.Errorf("format source:%w", err)
	}

	err = os.WriteFile(cfg.OutName(), out, 0644) //nolint:gosec,mnd
	if err != nil {
		return fmt.Errorf("write file:%w", err)
	}

	return nil
}
