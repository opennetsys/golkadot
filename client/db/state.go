package clientdb

import (
	storagetypes "github.com/opennetsys/godot/client/storage/types"
	"github.com/opennetsys/godot/common/db"
	"github.com/opennetsys/godot/common/triedb"
)

func createStateDB(dbs *triedb.TrieDB) *StateDB {
	storageMethodU8a := createU8a(db.BaseDB(dbs), storagetypes.Substrate.Code)
	return &StateDB{
		DB:   dbs,
		Code: &storageMethodU8a,
	}
}
