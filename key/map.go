package key

import (
	"strings"
)

const (
	prefixByPath = "byPath"
	wrongIDx     = -1
)

func newMap() *Map {
	return &Map{
		idx:      wrongIDx,
		wild:     nil,
		children: nil,
	}
}

type Map struct {
	idx      int
	wild     *Map
	children map[string]*Map
}

// ByPath returns the lookup key for a name split by sep.
func ByPath(name, sep string) []string {
	return []string{prefixByPath, name, sep}
}

// Index returns the stored index for path. A byPath key matches the longest name.
func (m *Map) Index(path []string) (int, bool) {
	if data, ok := m.find(path); ok {
		return data.idx, true
	}

	if len(path) == 3 && path[0] == prefixByPath {
		data, ok := m.byPath(path[1], path[2])

		return data.idx, ok
	}

	return 0, false
}

// Add stores idx at path. A trailing {name} segment is a wildcard.
func (m *Map) Add(idx int, path []string) {
	m.add(path).idx = idx
}

func (m *Map) add(path []string) *Map {
	if len(path) == 0 {
		return m
	}

	name, path := path[0], path[1:]

	if IsWild(name) {
		m.wild = newMap()

		return m.wild.add(path)
	}

	if m.children == nil {
		m.children = map[string]*Map{}
	}

	if _, ok := m.children[name]; !ok {
		m.children[name] = newMap()
	}

	if len(path) > 0 {
		return m.children[name].add(path)
	}

	return m.children[name]
}

func (m *Map) byPath(path, sep string) (*Map, bool) {
	if len(path) == 0 {
		return m, m.isValid()
	}

	var (
		best    *Map
		bestLen int
	)

	for name, data := range m.children {
		after, ok := strings.CutPrefix(path, name)
		if !ok || after != "" && (sep == "" || !strings.HasPrefix(after, sep)) {
			continue
		}

		next, nok := data, data.isValid()
		if after != "" {
			next, nok = data.byPath(after[len(sep):], sep)
		}

		if !nok || len(name) <= bestLen {
			continue
		}

		best, bestLen = next, len(name)
	}

	if best != nil {
		return best, true
	}

	if m.wild == nil || sep == "" {
		return m, false
	}

	if _, rest, ok := strings.Cut(path, sep); ok {
		return m.wild.byPath(rest, sep)
	}

	return m.wild, m.wild.isValid()
}

func (m *Map) find(path []string) (*Map, bool) {
	name := path[0]

	last := len(path) == 1
	if !last {
		path = path[1:]
	}

	data, ok := m.children[name]
	if !ok && m.wild != nil {
		if last {
			return m.wild, m.wild.isValid()
		}

		return m.wild.find(path)
	}

	if !ok {
		return data, false
	}

	if last {
		return data, data.isValid()
	}

	return data.find(path)
}

func (m *Map) isValid() bool {
	return m.idx != wrongIDx
}
