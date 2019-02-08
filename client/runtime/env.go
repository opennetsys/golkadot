package runtime

import "github.com/opennetsys/golkadot/common/triedb"

// Env ...
type Env struct {
	DB   *triedb.TrieDB
	Heap *Heap
}

// NewEnv ...
func NewEnv(db *triedb.TrieDB) *Env {
	return &Env{
		DB:   db,
		Heap: NewHeap(),
	}
}
