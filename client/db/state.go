package clientdb

import (
	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triedb"
	"github.com/c3systems/go-substrate/storagetypes"
)

func createStateDB(dbs *triedb.TrieDB) *StateDB {
	storageMethodU8a := createU8a(db.BaseDB(dbs), storagetypes.Substrate.Code)
	return &StateDB{
		DB:   dbs,
		Code: &storageMethodU8a,
	}
}
