package runtime

import "github.com/c3systems/go-substrate/common/triedb"

// Env ...
type Env struct {
	db   *triedb.TrieDB
	heap *Heap
}

// NewEnv ...
func NewEnv(db *triedb.TrieDB) *Env {
	return &Env{
		db:   db,
		heap: NewHeap(),
	}
}
