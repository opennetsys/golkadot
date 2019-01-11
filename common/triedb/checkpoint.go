package triedb

import (
	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/c3systems/go-substrate/common/triehash"
)

// Checkpoint ...
type Checkpoint struct {
	rootHash []uint8
	txRoot   []uint8
}

// NewCheckpoint ...
func NewCheckpoint(rootHash []uint8) *Checkpoint {
	if rootHash == nil {
		rootHash = triecodec.Hashing(triehash.TrieRoot(nil))
	}

	return &Checkpoint{
		rootHash: rootHash,
		txRoot:   rootHash,
	}
}

// CreateCheckpoint ...
func (c *Checkpoint) CreateCheckpoint() []uint8 {
	c.txRoot = c.rootHash

	return c.txRoot
}

// CommitCheckpoint ...
func (c *Checkpoint) CommitCheckpoint() []uint8 {
	return c.rootHash
}

// RevertCheckpoint ...
func (c *Checkpoint) RevertCheckpoint() []uint8 {
	c.rootHash = c.txRoot

	return c.rootHash
}

// RootHash ...
func (c *Checkpoint) RootHash() []uint8 {
	return c.rootHash
}

// TxRoot ...
func (c *Checkpoint) TxRoot() []uint8 {
	return c.txRoot
}
