package triedb

import (
	"testing"

	"github.com/c3systems/go-substrate/db"
)

func TestImpl(t *testing.T) {
	// TODO: table tests

	memDB := db.NewMemoryDB(&db.BaseOptions{})
	txDB := TxDB(memDB)
	impl := NewImpl(&txDB, []uint8{0x1})

	_ = impl
}
