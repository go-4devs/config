package key

import (
	"strings"
)

const (
	prefixByPath = "byPath"
)

func newMap() *Map {
	return &Map{
		idx:      0,
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
	for name := range m.children {
		if after, ok := strings.CutPrefix(path, name); ok {
			data := m.children[name]
			if len(after) == 0 {
				return data, true
			}

			after, ok = strings.CutPrefix(after, sep)
			if !ok {
				return data, false
			}

			if data.wild == nil {
				return data.byPath(after, sep)
			}

			if idx := strings.Index(after, sep); idx != -1 {
				return data.wild.byPath(after[idx+1:], sep)
			}

			return data, false
		}
	}

	return m, false
}

func (m *Map) find(path []string) (*Map, bool) {
	name := path[0]

	last := len(path) == 1
	if !last {
		path = path[1:]
	}

	if m.wild != nil {
		return m.wild.find(path)
	}

	data, ok := m.children[name]
	if !ok {
		return data, false
	}

	if last {
		return data, data.children == nil
	}

	return data.find(path)
}
