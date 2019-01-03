package triedb

import (
	"testing"

	"github.com/c3systems/go-substrate/common/db"
)

func TestImpl(t *testing.T) {
	// TODO: table tests

	memdb := db.NewMemoryDB(&db.BaseOptions{})
	impl := NewImpl(memdb, []uint8{0x1})

	_ = impl
}
