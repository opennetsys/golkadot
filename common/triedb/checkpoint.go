package triedb

// Checkpoint ...
type Checkpoint struct {
	rootHash []uint8
	txRoot   []uint8
}

// NewCheckpoint ...
func NewCheckpoint(rootHash []uint8) *Checkpoint {
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
