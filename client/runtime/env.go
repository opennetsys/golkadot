package runtime

import "github.com/opennetsys/godot/common/triedb"

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
