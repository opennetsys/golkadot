package triedb

import (
	"github.com/opennetsys/golkadot/common/triecodec"
	"github.com/opennetsys/golkadot/common/triehash"
)

// Checkpoint ...
type Checkpoint struct {
	rootHash []byte
	txRoot   []byte
}

// NewCheckpoint ...
func NewCheckpoint(rootHash []byte) *Checkpoint {
	if rootHash == nil {
		tmpHash := triehash.TrieRoot(nil)
		rootHash = triecodec.Hashing(tmpHash[:])
	}

	return &Checkpoint{
		rootHash: rootHash,
		txRoot:   rootHash,
	}
}

// CreateCheckpoint ...
func (c *Checkpoint) CreateCheckpoint() []byte {
	c.txRoot = c.rootHash

	return c.txRoot
}

// CommitCheckpoint ...
func (c *Checkpoint) CommitCheckpoint() []byte {
	return c.rootHash
}

// RevertCheckpoint ...
func (c *Checkpoint) RevertCheckpoint() []byte {
	c.rootHash = c.txRoot

	return c.rootHash
}

// RootHash ...
func (c *Checkpoint) RootHash() []byte {
	return c.rootHash
}

// TxRoot ...
func (c *Checkpoint) TxRoot() []byte {
	return c.txRoot
}
