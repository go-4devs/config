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

// NewMap returns an empty in-memory provider.
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

// Len returns the number of stored values.
func (m *Map) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return len(m.vals)
}

// Value returns the value stored under key.
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

// Bind does nothing.
func (m *Map) Bind(_ context.Context, _ config.Variables) error {
	return nil
}

// Name returns the provider name.
func (m *Map) Name() string {
	if m.name != "" {
		return m.name
	}

	return NameMap
}

// HasOption reports whether key is stored.
func (m *Map) HasOption(key ...string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.idx.Index(key)

	return ok
}

// SetOption stores val under key.
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

// AppendOption appends val to the string list stored under key.
func (m *Map) AppendOption(val string, keys ...string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	id, ok := m.idx.Index(keys)
	if !ok {
		m.idx.Add(len(m.vals), keys)
		m.vals = append(m.vals, value.New(value.Strings{val}))

		return nil
	}

	old, tok := m.vals[id].(value.Strings)
	if !tok {
		return fmt.Errorf("%w:%T", config.ErrWrongType, m.vals[id])
	}

	m.vals[id] = value.New(old.Append(val))

	return nil
}
