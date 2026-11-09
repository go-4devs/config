package validator_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/test/require"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
)

func TestFileExists(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "note.txt")
	require.NoError(t, os.WriteFile(path, []byte("x"), 0o600))

	opt := option.String("path", "path")

	require.NoError(t, validator.FileExists(opt, value.New(path)))
	require.NoError(t, validator.FileExists(opt, value.JString(path)))
	require.ErrorIs(t, validator.FileExists(opt, value.New(filepath.Join(dir, "missing.txt"))), validator.ErrFileExists)
	require.ErrorIs(t, validator.FileExists(opt, value.New(dir)), validator.ErrFileExists)
	require.ErrorIs(t, validator.FileExists(opt, value.New("")), validator.ErrFileExists)
	require.ErrorIs(t, validator.FileExists(option.Int("port", "port"), value.New(1)), validator.ErrInvalid)
}
