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

func ByPath(name, sep string) []string {
	return []string{prefixByPath, name, sep}
}

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

func (m *Map) Add(idx int, path []string) {
	m.add(path).idx = idx
}

func (m *Map) add(path []string) *Map {
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

	for name := range m.children {
		if after, ok := strings.CutPrefix(path, name); ok {
			data := m.children[name]
			if len(after) == 0 || len(after) == len(sep) {
				return data, data.isValid()
			}

			return data.byPath(after[len(sep):], sep)
		}
	}

	if m.wild == nil {
		return m, m.isValid()
	}

	if idx := strings.Index(path, sep); idx != -1 {
		return m.wild.byPath(path[idx+1:], sep)
	}

	return m, m.isValid()
}

func (m *Map) find(path []string) (*Map, bool) {
	name := path[0]

	last := len(path) == 1
	if !last {
		path = path[1:]
	}

	data, ok := m.children[name]
	if !ok && m.wild != nil {
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
