package triedb

import (
	"github.com/c3systems/go-substrate/common/crypto"
	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/c3systems/go-substrate/common/triehash"
)

// Checkpoint ...
type Checkpoint struct {
	rootHash *crypto.Blake2b256Hash
	txRoot   *crypto.Blake2b256Hash
}

// NewCheckpoint ...
func NewCheckpoint(rootHash *crypto.Blake2b256Hash) *Checkpoint {
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
func (c *Checkpoint) CreateCheckpoint() *crypto.Blake2b256Hash {
	c.txRoot = c.rootHash

	return c.txRoot
}

// CommitCheckpoint ...
func (c *Checkpoint) CommitCheckpoint() *crypto.Blake2b256Hash {
	return c.rootHash
}

// RevertCheckpoint ...
func (c *Checkpoint) RevertCheckpoint() *crypto.Blake2b256Hash {
	c.rootHash = c.txRoot

	return c.rootHash
}

// RootHash ...
func (c *Checkpoint) RootHash() *crypto.Blake2b256Hash {
	return c.rootHash
}

// TxRoot ...
func (c *Checkpoint) TxRoot() *crypto.Blake2b256Hash {
	return c.txRoot
}
