package clientdb

import "github.com/opennetsys/golkadot/common/triedb"

// DefaultPath ...
var DefaultPath = "~/.golkadot"

// DefaultType ...
var DefaultType = "disk"

// DBPathPrefix ...
var DBPathPrefix = "database"

// DBConfigType ...
var DBConfigType = "disk" // other option is "memory"

// StateDB ...
type StateDB struct {
	DB   *triedb.TrieDB
	Code *StorageMethodU8a
}
