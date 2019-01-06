package db

import (
	"log"
)

// Storage ...
type Storage map[string][]uint8

// MemoryDB ...
type MemoryDB struct {
	storage Storage
}

// NewMemoryDB creates a MemoryDB database extending TransactionDB
// options argument (optional) used to specifies whether
// to use compression.
// - Enable compression when storing "blocks" since they compress well.
// - Disable compression when storing "state" Transactions Trie
// data structure serialised with Recursive Length Prefix (RLP) encoding
func NewMemoryDB(options *BaseOptions) *MemoryDB {
	return &MemoryDB{
		storage: Storage{},
	}
}

// Close ...
func (m *MemoryDB) Close() {
	m.Empty()
}

// Open ...
func (m *MemoryDB) Open() {
	m.Empty()
}

// Drop ...
func (m *MemoryDB) Drop() {
	m.Empty()
}

// Empty ...
func (m *MemoryDB) Empty() {
	m.storage = Storage{}
}

// Rename ...
func (m *MemoryDB) Rename(base, file string) {
	log.Println("rename is not implemented")
}

// Maintain ...
func (m *MemoryDB) Maintain(fn *ProgressCB) error {
	if fn != nil {
		f := *fn
		f(&ProgressValue{
			IsCompleted: true,
			Keys:        len(m.storage),
			Percent:     100,
		})
	}
	return nil
}

// Size ...
func (m *MemoryDB) Size() int {
	log.Println("size is not implemented")
	return 0
}

// Del ...
func (m *MemoryDB) Del(key []uint8) {
	delete(m.storage, string(key))
}

// Get ...
func (m *MemoryDB) Get(key []uint8) []uint8 {
	value, found := m.storage[string(key)]
	if found {
		return value
	}

	return nil
}

// Put ...
func (m *MemoryDB) Put(key []uint8, value []uint8) {
	m.storage[string(key)] = value
}
