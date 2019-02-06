package clientdb

import (
	storagetypes "github.com/opennetsys/go-substrate/client/storage/types"
	"github.com/opennetsys/go-substrate/common/db"
	"github.com/opennetsys/go-substrate/common/triedb"
)

func createStateDB(dbs *triedb.TrieDB) *StateDB {
	storageMethodU8a := createU8a(db.BaseDB(dbs), storagetypes.Substrate.Code)
	return &StateDB{
		DB:   dbs,
		Code: &storageMethodU8a,
	}
}
