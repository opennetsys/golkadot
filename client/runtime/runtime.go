package runtime

import "github.com/c3systems/go-substrate/common/triedb"

// Runtime ...
type Runtime struct {
}

// NewRuntime ...
func NewRuntime(stateDB *triedb.TrieDB) *Runtime {
	return &Runtime{}
}
