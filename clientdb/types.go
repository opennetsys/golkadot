package clientdb

import (
	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triedb"
)

// DefaultPath ...
var DefaultPath = "~/.go-substrate"

// DefaultType ...
var DefaultType = "disk"

// DBPathPrefix ...
var DBPathPrefix = "database"

// DBConfigType ...
var DBConfigType = "disk" // other option is "memory"

// DBConfig ...
type DBConfig struct {
	Compact  bool
	IsTrieDb bool
	Path     string
	Snapshot bool
	Type     string // DBConfigType
}

// BlockDB ...
type BlockDB struct {
	DB         db.BaseDB
	BestHash   StorageMethodU8a
	BestNumber StorageMethodBn
	BlockData  StorageMethodU8a
	Hash       StorageMethodU8a
	Header     StorageMethodU8a
}

// StateDB ...
type StateDB struct {
	db   *triedb.TrieDB
	code StorageMethodU8a
}

// ChainDbs ...
type ChainDbs interface {
	Snapshot()
}

// StorageMethodU8a ...
// TODO
type StorageMethodU8a struct {
	//Del(params ...interface{})
	//Get(params ...interface{}) []uint8
	//Set(value []uint8, params ...interface{})
	//OnUpdate(callback func(value []uint8))
}

// StorageMethodBn ...
// TODO
type StorageMethodBn struct {
	//Del(params ...interface{})
	//Get(params ...interface{}) *big.Int
	//Set(value *big.Int, params ...interface{})
	//OnUpdate(callback func(value *big.Int))
}
