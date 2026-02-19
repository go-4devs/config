package memory

import (
	"context"
	"fmt"
	"sync"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/key"
	"gitoa.ru/go-4devs/config/value"
)

const NameMap = "map"

var _ config.BindProvider = (*Map)(nil)

func NewMap(name string) *Map {
	return &Map{
		name: name,
		vals: make([]config.Value, 0),
		idx:  key.Map{},
		mu:   sync.Mutex{},
	}
}

type Map struct {
	mu   sync.Mutex
	vals []config.Value
	idx  key.Map
	name string
}

func (m *Map) Len() int {
	return len(m.vals)
}

func (m *Map) Value(_ context.Context, key ...string) (config.Value, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idx, ok := m.idx.Index(key)
	if !ok {
		return nil, fmt.Errorf("%w", config.ErrNotFound)
	}

	val := m.vals[idx]

	return val, nil
}

func (m *Map) Bind(_ context.Context, _ config.Variables) error {
	return nil
}

func (m *Map) Name() string {
	if m.name != "" {
		return m.name
	}

	return NameMap
}

func (m *Map) HasOption(key ...string) bool {
	_, ok := m.idx.Index(key)

	return ok
}

func (m *Map) SetOption(val any, key ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if id, ok := m.idx.Index(key); ok {
		m.vals[id] = value.New(val)

		return
	}

	m.idx.Add(len(m.vals), key)
	m.vals = append(m.vals, value.New(val))
}

func (m *Map) AppendOption(val string, keys ...string) error {
	id, ok := m.idx.Index(keys)
	if !ok {
		data := value.Strings{val}

		m.SetOption(data, keys...)

		return nil
	}

	old, tok := m.vals[id].(value.Strings)
	if !tok {
		return fmt.Errorf("%w:%T", config.ErrWrongType, m.vals[id])
	}

	m.SetOption(old.Append(val), keys...)

	return nil
}
