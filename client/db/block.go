package db

import (
	"github.com/c3systems/go-substrate/common/db"
)

func createU8a(dbs *db.BaseDB, fn func()) StorageMethodU8a {
	// TODO
	return StorageMethodU8a{}
}

func createBn(dbs *db.BaseDB, fn func(), n int) StorageMethodBn {
	// TODO
	return StorageMethodBn{}
}

// NewBlockDB ...
func NewBlockDB(dbs *db.BaseDB) *BlockDB {
	return &BlockDB{
		DB:         *dbs,
		BestHash:   createU8a(dbs, KeyBestHash()),
		BestNumber: createBn(dbs, KeyBestNumber(), 64),
		BlockData:  createU8a(dbs, KeyBlockByHash()),
		Header:     createU8a(dbs, KeyHeaderByHash()),
		Hash:       createU8a(dbs, KeyHashByNumber()),
	}
}
