package clientdb

import (
	"github.com/opennetsys/godot/common/db"
	types "github.com/opennetsys/godot/types"
)

func createU8a(dbs db.BaseDB, fn types.StorageFunction) StorageMethodU8a {
	// TODO
	return StorageMethodU8a{}
}

func createBn(dbs db.BaseDB, fn types.StorageFunction, n int) StorageMethodBn {
	// TODO
	return StorageMethodBn{}
}

// NewBlockDB ...
func NewBlockDB(dbs db.BaseDB) *BlockDB {
	return &BlockDB{
		DB:         dbs,
		BestHash:   createU8a(dbs, KeyBestHash()),
		BestNumber: createBn(dbs, KeyBestNumber(), 64),
		BlockData:  createU8a(dbs, KeyBlockByHash()),
		Header:     createU8a(dbs, KeyHeaderByHash()),
		Hash:       createU8a(dbs, KeyHashByNumber()),
	}
}
