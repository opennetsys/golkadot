package clientdb

import (
	storagetypes "github.com/opennetsys/golkadot/client/storage/types"
	"github.com/opennetsys/golkadot/common/db"
	"github.com/opennetsys/golkadot/common/triedb"
)

func createStateDB(dbs *triedb.TrieDB) *StateDB {
	storageMethodU8a := createU8a(db.BaseDB(dbs), storagetypes.Substrate.Code)
	return &StateDB{
		DB:   dbs,
		Code: &storageMethodU8a,
	}
}
