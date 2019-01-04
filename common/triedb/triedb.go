package triedb

import (
	"log"
	"time"

	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triecodec"
)

// Trie ...
type Trie struct {
	impl *Impl
}

// NewTrie ...
func NewTrie(db db.TXDB, rootHash []byte) *Trie {
	impl := NewImpl(db, rootHash)
	return &Trie{
		impl: impl,
	}
}

// Transaction ...
func (t *Trie) Transaction(fn func() bool) bool {
	t.impl.checkpoint.CreateCheckpoint()

	result := t.impl.db.Transaction(fn)

	if result {
		t.impl.checkpoint.CommitCheckpoint()
	} else {
		t.impl.checkpoint.RevertCheckpoint()
	}

	return result
}

// Open ...
func (t *Trie) Open() {
	t.impl.db.Open()
}

// Close ...
func (t *Trie) Close() {
	t.impl.db.Close()
}

// Empty ...
func (t *Trie) Empty() {
	t.impl.db.Empty()
}

// Drop ...
func (t *Trie) Drop() {
	t.impl.db.Drop()
}

// Maintain ...
func (t *Trie) Maintain(fn *db.ProgressCB) {
	t.impl.db.Maintain(fn)
}

// Rename ...
func (t *Trie) Rename(base, file string) {
	t.impl.db.Rename(base, file)
}

// Size ...
func (t *Trie) Size() int {
	return t.impl.db.Size()
}

// Delete ...
func (t *Trie) Delete(key []uint8) {
	node := t.impl.Del(
		t.impl.GetNode(t.impl.checkpoint.rootHash),
		triecodec.ToNibbles(key),
	)

	t.impl.SetRootNode(node)
}

// Get ...
func (t *Trie) Get(key []uint8) []uint8 {
	value := t.impl.Get(
		t.impl.GetNode(t.impl.checkpoint.rootHash),
		triecodec.ToNibbles(key),
	)

	return value.([]uint8)
}

// Put ...
func (t *Trie) Put(key, value []uint8) {
	node := t.impl.Put(
		t.impl.GetNode(t.impl.checkpoint.rootHash),
		triecodec.ToNibbles(key),
		value,
	)

	t.impl.SetRootNode(node)
}

// GetRoot ...
func (t *Trie) GetRoot() []uint8 {
	rootnode := t.GetNode(nil)

	if IsNull(rootnode) {
		return []uint8{}
	}

	return t.impl.checkpoint.rootHash
}

// GetNode ...
func (t *Trie) GetNode(hash []uint8) Node {
	if hash == nil {
		hash = t.impl.checkpoint.rootHash
	}

	return t.impl.GetNode(hash)
}

// SetRoot ...
func (t *Trie) SetRoot(rootHash []uint8) {
	t.impl.checkpoint.rootHash = rootHash
}

// Snapshot ...
func (t *Trie) Snapshot(dest Trie, fn db.ProgressCB) int {
	start := time.Now().Unix()

	keys := t.impl.Snapshot(dest, fn, t.impl.checkpoint.rootHash, 0, 0, 0)
	elapsed := time.Now().Unix() - start

	dest.SetRoot(t.impl.checkpoint.rootHash)

	newSize := dest.impl.db.Size()
	percentage := 100 * (newSize / t.impl.db.Size())
	sizeMB := newSize / (1024 * 1024)

	log.Printf("snapshot created in %d, %dk keys, %dMB (%d%%)", elapsed, keys/1e3, sizeMB, percentage)

	fn(&db.ProgressValue{
		IsCompleted: true,
		Keys:        keys,
		Percent:     100,
	})

	return keys
}
