package key_test

import (
	"testing"

	"gitoa.ru/go-4devs/config/key"
)

func TestMap_ByPath(t *testing.T) {
	t.Parallel()

	const (
		expID int = 1
		newID int = 42
	)

	data := key.Map{}
	data.Add(expID, []string{"test", "data", "three"})
	data.Add(expID, []string{"test", "other"})
	data.Add(newID, []string{"new", "{data}", "test"})

	idx, ok := data.Index(key.ByPath("test-other", "-"))
	if !ok {
		t.Error("key not found")
	}

	if idx != expID {
		t.Errorf("idx exp:%v got:%v", expID, idx)
	}

	if nidx, nok := data.Index(key.ByPath("new-service-test", "-")); !nok && nidx != newID {
		t.Errorf("idx exp:%v got:%v", newID, nidx)
	}
}

func TestMap_Add(t *testing.T) {
	t.Parallel()

	const (
		expID int = 1
		newID int = 42
	)

	data := key.Map{}
	data.Add(expID, []string{"test", "data"})
	data.Add(expID, []string{"test", "other"})
	data.Add(newID, []string{"new"})

	idx, ok := data.Index([]string{"test", "data"})
	if !ok {
		t.Error("key not found")
	}

	if idx != expID {
		t.Errorf("idx exp:%v got:%v", expID, idx)
	}

	if nidx, nok := data.Index([]string{"new"}); !nok && nidx != newID {
		t.Errorf("idx exp:%v got:%v", newID, nidx)
	}
}

func TestMap_Wild(t *testing.T) {
	t.Parallel()

	const (
		expID int = 1
		newID int = 42
	)

	data := key.Map{}
	data.Add(expID, []string{"test", "{data}", "id"})
	data.Add(newID, []string{"new", "data"})

	idx, ok := data.Index([]string{"test", "data", "id"})
	if !ok {
		t.Error("key not found")
	}

	if idx != expID {
		t.Errorf("idx exp:%v got:%v", expID, idx)
	}
}
