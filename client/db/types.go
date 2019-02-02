package clientdb

import "github.com/c3systems/go-substrate/common/triedb"

// DefaultPath ...
var DefaultPath = "~/.go-substrate"

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
