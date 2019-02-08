package runtime

import (
	"fmt"
	"math"

	"github.com/opennetsys/godot/common/db"
	"github.com/opennetsys/godot/common/triehash"
)

// Storage ...
type Storage struct {
	Data *Data
	Trie *Trie
}

// NewStorage ...
func NewStorage(heap *Heap, dbs db.TXDB) *Storage {
	return &Storage{
		Data: NewData(heap, dbs),
		Trie: NewTrie(heap, nil),
	}
}

// U32Max ...
var U32Max = int64(math.Pow(float64(2), float64(32))) - 1

// Data ...
type Data struct {
	Heap *Heap
	DB   db.TXDB
}

// NewData ...
func NewData(heap *Heap, dbs db.TXDB) *Data {
	return &Data{
		Heap: heap,
		DB:   dbs,
	}
}

// ClearPrefix ...
func (d *Data) ClearPrefix(prefixPtr Pointer, prefixLength int64) {
	fmt.Println("clear_prefix has not been implemented, only stubbed")
}

// ClearStorage ...
func (d *Data) ClearStorage(keyPtr Pointer, keyLength int64) {
	key := d.Heap.Get(keyPtr, keyLength)
	d.DB.Del(key)
}

// ExistsStorage  ...
func (d *Data) ExistsStorage(keyPtr Pointer, keyLength int64) int {
	key := d.Heap.Get(keyPtr, keyLength)
	data := get(d.DB, key, 0, U32Max)
	hasEntry := 0
	if data != nil && len(data) > 0 {
		hasEntry = 1
	}

	return hasEntry
}

// GetAllocatedStorage ...
func (d *Data) GetAllocatedStorage(keyPtr Pointer, keyLength int64, lenPtr Pointer) Pointer {
	key := d.Heap.Get(keyPtr, keyLength)
	data := get(d.DB, key, 0, U32Max)

	length := U32Max
	if data != nil {
		length = int64(len(data))
	}

	d.Heap.SetU32(lenPtr, uint32(length))

	if data == nil {
		return 0
	}

	return d.Heap.Set(d.Heap.Allocate(length), data)
}

// GetStorageInto ...
func (d *Data) GetStorageInto(keyPtr Pointer, keyLength int64, dataPtr Pointer, dataLength int64, offset int64) int64 {
	key := d.Heap.Get(keyPtr, keyLength)
	data := get(d.DB, key, offset, dataLength)

	if data == nil {
		return U32Max
	}

	d.Heap.Set(dataPtr, data)

	return int64(len(data))
}

// SetStorage ...
func (d *Data) SetStorage(keyPtr Pointer, keyLength int64, dataPtr Pointer, dataLength int64) {
	key := d.Heap.Get(keyPtr, keyLength)
	data := d.Heap.Get(dataPtr, dataLength)

	d.DB.Put(key, data)
}

func get(dbs db.TXDB, key []uint8, offset int64, maxLength int64) []uint8 {
	data := dbs.Get(key)

	if data == nil {
		return nil
	}

	dataLength := int64(math.Min(float64(maxLength), float64(len(data))-float64(offset)))

	return data[offset : offset+dataLength]
}

// Trie ...
type Trie struct {
	Heap *Heap
	DB   db.TXDB
}

// NewTrie ...
func NewTrie(heap *Heap, dbs db.TXDB) *Trie {
	return &Trie{
		Heap: heap,
		DB:   dbs,
	}
}

// Blake2b256EnumeratedTrieRoot ...
func (t *Trie) Blake2b256EnumeratedTrieRoot(valuesPtr Pointer, lenPtr Pointer, count int64, resultPtr Pointer) {
	pairs := make([][]uint8, count)

	for index := range pairs {
		length := t.Heap.GetU32(Pointer(int(lenPtr) + (index * 4)))
		data := t.Heap.Get(valuesPtr, int64(length))

		valuesPtr += Pointer(length)

		pairs[index] = data
	}

	root := triehash.TrieRootOrdered(pairs)[:]

	t.Heap.Set(resultPtr, root)
}

// StorageChangesRoot ...
func (t *Trie) StorageChangesRoot(parentHashData Pointer, parentHashLen int64, parentNumHi int64, parentNumLo int64, result Pointer) int64 {
	fmt.Println("storage_changes_root has not been implemented, only stubbed")
	return 0
}

// StorageRoot ...
func (t *Trie) StorageRoot(resultPtr Pointer) {
	root := t.DB.GetRoot()[:]

	t.Heap.Set(resultPtr, root)
}
