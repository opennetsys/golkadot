package db

import (
	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/fileflatdb"
)

// DiskDB ...
type DiskDB struct {
	db.TransactionDB
}

// NewDiskDB creates DiskDB database using LruDB for caching with FileFlatDB and extending TransactionDB.
func NewDiskDB(base, name string, options *db.BaseDBOptions) *DiskDB {
	flatdb := fileflatdb.NewFileFlatDB(base, name)
	basedb := db.BaseDB(flatdb)
	lrudb := db.NewLruDB(basedb, -1)
	backingdb := db.BaseDB(lrudb)
	diskdb := &DiskDB{}
	diskdb.Backing = backingdb
	return diskdb
}
