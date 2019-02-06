package runtime

import "github.com/opennetsys/go-substrate/common/triedb"

// Runtime ...
type Runtime struct {
}

// NewRuntime ...
func NewRuntime(stateDB *triedb.TrieDB) *Runtime {
	return &Runtime{}
}
