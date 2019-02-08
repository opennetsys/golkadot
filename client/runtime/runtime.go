package runtime

import "github.com/opennetsys/golkadot/common/triedb"

// Runtime ...
type Runtime struct {
}

// NewRuntime ...
func NewRuntime(stateDB *triedb.TrieDB) *Runtime {
	return &Runtime{}
}
