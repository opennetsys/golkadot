package runtime

import "github.com/opennetsys/godot/common/triedb"

// Runtime ...
type Runtime struct {
}

// NewRuntime ...
func NewRuntime(stateDB *triedb.TrieDB) *Runtime {
	return &Runtime{}
}
