package pkg

import (
	"bytes"
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var cache = sync.Map{}

func ByPath(ctx context.Context, fname string, isDir bool) (string, error) {
	if !filepath.IsAbs(fname) {
		pwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}

		fname = filepath.Join(pwd, fname)
	}

	goModPath, _ := goModPath(ctx, fname, isDir)
	if strings.Contains(goModPath, "go.mod") {
		pkgPath, err := getPkgPathFromGoMod(fname, isDir, goModPath)
		if err != nil {
			return "", err
		}

		return pkgPath, nil
	}

	return getPkgPathFromGOPATH(fname, isDir)
}

// empty if no go.mod, GO111MODULE=off or go without go modules support.
func goModPath(ctx context.Context, fname string, isDir bool) (string, error) {
	root := fname
	if !isDir {
		root = filepath.Dir(fname)
	}

	var modPath string

	loadModPath, ok := cache.Load(root)
	if ok {
		modPath, _ = loadModPath.(string)

		return modPath, nil
	}

	defer func() {
		cache.Store(root, modPath)
	}()

	cmd := exec.CommandContext(ctx, "go", "env", "GOMOD")
	cmd.Dir = root

	stdout, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	modPath = string(bytes.TrimSpace(stdout))

	return modPath, nil
}

func getPkgPathFromGoMod(fname string, isDir bool, goModPath string) (string, error) {
	modulePath := getModulePath(goModPath)
	if modulePath == "" {
		return "", fmt.Errorf("c%w module path from %s", ErrNotFound, goModPath)
	}

	rel := path.Join(modulePath, filePathToPackagePath(strings.TrimPrefix(fname, filepath.Dir(goModPath))))

	if !isDir {
		return path.Dir(rel), nil
	}

	return path.Clean(rel), nil
}

func getModulePath(goModPath string) string {
	var pkgPath string

	cacheOkgPath, ok := cache.Load(goModPath)
	if ok {
		pkgPath, _ = cacheOkgPath.(string)

		return pkgPath
	}

	defer func() {
		cache.Store(goModPath, pkgPath)
	}()

	data, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}

	pkgPath = modulePath(data)

	return pkgPath
}

func getPkgPathFromGOPATH(fname string, isDir bool) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	for _, p := range strings.Split(gopath, string(filepath.ListSeparator)) {
		prefix := filepath.Join(p, "src") + string(filepath.Separator)

		rel, err := filepath.Rel(prefix, fname)
		if err == nil && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			if !isDir {
				return path.Dir(filePathToPackagePath(rel)), nil
			}

			return path.Clean(filePathToPackagePath(rel)), nil
		}
	}

	return "", fmt.Errorf("%w: file '%v' is not in GOPATH '%v'", ErrNotFound, fname, gopath)
}

func filePathToPackagePath(path string) string {
	return filepath.ToSlash(path)
}

var (
	slashSlash = []byte("//")
	moduleStr  = []byte("module")
)

// modulePath returns the module path from the gomod file text.
// If it cannot find a module path, it returns an empty string.
// It is tolerant of unrelated problems in the go.mod file.
func modulePath(mod []byte) string {
	for len(mod) > 0 {
		line := mod

		mod = nil
		if i := bytes.IndexByte(line, '\n'); i >= 0 {
			line, mod = line[:i], line[i+1:]
		}

		if i := bytes.Index(line, slashSlash); i >= 0 {
			line = line[:i]
		}

		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, moduleStr) {
			continue
		}

		line = line[len(moduleStr):]
		n := len(line)

		line = bytes.TrimSpace(line)
		if len(line) == n || len(line) == 0 {
			continue
		}

		if line[0] == '"' || line[0] == '`' {
			p, err := strconv.Unquote(string(line))
			if err != nil {
				return "" // malformed quoted string or multiline module path
			}

			return p
		}

		return string(line)
	}

	return "" // missing module path
}
