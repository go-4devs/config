package key_test

import (
	"testing"

	"github.com/go-4devs/config/key"
)

func TestMap_ByPath(t *testing.T) {
	t.Parallel()

	const (
		expID int = 1
		othID int = 0
		newID int = 42
		grpID int = 27
	)

	data := key.Map{}
	data.Add(expID, []string{"test", "data", "three"})
	data.Add(othID, []string{"test", "other"})
	data.Add(newID, []string{"new", "{data}", "test"})
	data.Add(grpID, []string{"new", "group"})

	idx, ok := data.Index(key.ByPath("test-other", "-"))
	if !ok {
		t.Error("key not found")
	}

	if idx != othID {
		t.Errorf("idx exp:%v got:%v", othID, idx)
	}

	if nidx, nok := data.Index(key.ByPath("new-service-test", "-")); !nok || nidx != newID {
		t.Errorf("idx exp:%v got:%v ok:%v", newID, nidx, nok)
	}

	if gidx, gok := data.Index(key.ByPath("new-group", "-")); !gok || gidx != grpID {
		t.Errorf("idx exp:%v got:%v ok:%v", grpID, gidx, gok)
	}

	const (
		portID = 9
		vID    = 8
		verbID = 7
	)

	data.Add(portID, []string{"port"})
	data.Add(vID, []string{"v"})
	data.Add(verbID, []string{"verbose"})

	if pidx, pok := data.Index(key.ByPath("portX", "-")); pok {
		t.Errorf("portX matched %v", pidx)
	}

	if vidx, vok := data.Index(key.ByPath("verbose", "-")); !vok || vidx != verbID {
		t.Errorf("idx exp:%v got:%v ok:%v", verbID, vidx, vok)
	}

	if sidx, sok := data.Index(key.ByPath("v", "-")); !sok || sidx != vID {
		t.Errorf("idx exp:%v got:%v ok:%v", vID, sidx, sok)
	}
}

func TestMap_TrailingWild(t *testing.T) {
	t.Parallel()

	const id = 1

	data := key.Map{}
	data.Add(id, []string{"user", "{id}"})

	idx, ok := data.Index([]string{"user", "5"})
	if !ok || idx != id {
		t.Fatalf("idx exp:%v got:%v ok:%v", id, idx, ok)
	}

	idx, ok = data.Index(key.ByPath("user-5", "-"))
	if !ok || idx != id {
		t.Fatalf("path idx exp:%v got:%v ok:%v", id, idx, ok)
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

	if nidx, nok := data.Index([]string{"new"}); !nok || nidx != newID {
		t.Errorf("idx exp:%v got:%v ok:%v", newID, nidx, nok)
	}
}

func TestMap_WildKeepsSiblingFields(t *testing.T) {
	t.Parallel()

	data := key.Map{}
	data.Add(1, []string{"end", "{service}", "after"})
	data.Add(2, []string{"end", "{service}", "host"})

	idx, ok := data.Index([]string{"end", "api", "after"})
	if !ok || idx != 1 {
		t.Fatalf("after idx %v ok %v", idx, ok)
	}

	idx, ok = data.Index([]string{"end", "api", "host"})
	if !ok || idx != 2 {
		t.Fatalf("host idx %v ok %v", idx, ok)
	}
}

func TestMap_Wild(t *testing.T) {
	t.Parallel()

	const (
		expID int = 1
		grpID int = 27
		newID int = 42
	)

	data := key.Map{}
	data.Add(expID, []string{"test", "{data}", "id"})
	data.Add(grpID, []string{"test", "group"})
	data.Add(newID, []string{"new", "data"})

	idx, ok := data.Index([]string{"test", "data", "id"})
	if !ok {
		t.Error("key not found")
	}

	if idx != expID {
		t.Errorf("idx exp:%v got:%v", expID, idx)
	}

	gidx, gok := data.Index([]string{"test", "group"})
	if !gok {
		t.Error("key not found")
	}

	if gidx != grpID {
		t.Errorf("idx exp:%v got:%v", grpID, idx)
	}
}
