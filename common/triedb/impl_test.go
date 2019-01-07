package triedb

import (
	"testing"

	"github.com/c3systems/go-substrate/common/db"
)

func TestImpl(t *testing.T) {
	// TODO: table tests

	memdb := db.NewMemoryDB(&db.BaseOptions{})
	basedb := db.BaseDB(memdb)
	txdbt := db.NewTransactionDB(&basedb)
	txdb := db.TXDB(txdbt)
	rootHash := []uint8{0x1}
	codec := NewRLPCodec()
	impl := NewImpl(txdb, rootHash, codec)

	_ = impl
}
