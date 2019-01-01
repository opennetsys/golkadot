package fileflatdb

import (
	"log"
	"math/big"
	"os"
)

// Key ...
type Key struct {
	Key      *NibbleBuffer
	KeyAt    int64
	KeyValue []byte
}

// Value ...
type Value struct {
	Value   []byte
	ValueAt int64
}

// ValueInfo ...
type ValueInfo struct {
	ValueLength int64
	ValueAt     int64
}

// Impl ...
type Impl struct {
	Cache
}

// NewImpl ...
func NewImpl() *Impl {
	return &Impl{}
}

// GetKeyValue ...
func (i *Impl) GetKeyValue(keyAt int64) []byte {
	return i.GetCachedData(keyAt, int64(keyTotalSize))
}

// RetrieveBranch ...
func (i *Impl) RetrieveBranch(doCreate bool, branch []byte, entryIndex int, keyIndex int, key *NibbleBuffer) *Key {
	nextBranchAt := new(big.Int)
	nextBranchAt.SetBytes(branch[entryIndex+1 : entryIndex+1+uintSize])

	return i.FindKey(key, doCreate, int64(keyIndex+1), int64(nextBranchAt.Uint64()))
}

// RetrieveEmpty ...
func (i *Impl) RetrieveEmpty(doCreate bool, branch []byte, branchAt int64, entryIndex int64, key *NibbleBuffer) *Key {
	if doCreate {
		return i.WriteNewLeaf(branch, branchAt, entryIndex, key)
	}

	return nil
}

// RetrieveLeaf ...
func (i *Impl) RetrieveLeaf(doCreate bool, branch []byte, branchAt int, entryIndex int, keyIndex int, key *NibbleBuffer) *Key {
	keyAt := new(big.Int)
	keyAt.SetBytes(branch[entryIndex+1 : entryIndex+1+uintSize])
	keyValue := i.GetKeyValue(int64(keyAt.Uint64()))
	prevKey := i.SerializeKey(keyValue[0:keySize])
	matchIndex := keyIndex

	for matchIndex < keySize {
		if prevKey.Nibbles[matchIndex] != key.Nibbles[matchIndex] {
			break
		}

		matchIndex++
	}

	if matchIndex != keySize {
		if doCreate {
			return i.WriteNewBranch(branch, int64(branchAt), int64(entryIndex), key, int64(keyAt.Uint64()), prevKey, int64(matchIndex), int64(matchIndex-keyIndex-1))
		}

		return nil
	}

	return &Key{
		Key:      key,
		KeyAt:    int64(keyAt.Uint64()),
		KeyValue: keyValue,
	}
}

// FindKey ...
func (i *Impl) FindKey(key *NibbleBuffer, doCreate bool, keyIndex, branchAt int64) *Key {
	entryIndex := int(key.Nibbles[keyIndex]) * entrySize
	branch := i.GetCachedBranch(branchAt)
	entryType := branch[entryIndex]
	switch int(entryType) {
	case SlotBranch:
		return i.RetrieveBranch(doCreate, branch, entryIndex, int(keyIndex), key)
	case SlotEmpty:
		return i.RetrieveEmpty(doCreate, branch, branchAt, int64(entryIndex), key)
	case SlotLeaf:
		return i.RetrieveLeaf(doCreate, branch, int(branchAt), entryIndex, int(keyIndex), key)
	default:
		log.Fatalf("Unhandled entry type %v\n", entryType)
	}
	return nil
}

// ExtractValueInfo ...
func (i *Impl) ExtractValueInfo(keyValue []byte) *ValueInfo {
	valueLength := new(big.Int)
	valueLength.SetBytes(keyValue[keySize : keySize+uintSize])

	valueAt := new(big.Int)
	valueAt.SetBytes(keyValue[keySize+uintSize : keySize+uintSize+uintSize])

	return &ValueInfo{
		ValueLength: int64(valueLength.Uint64()),
		ValueAt:     int64(valueAt.Uint64()),
	}
}

// ReadValue ...
func (i *Impl) ReadValue(keyValue []byte) *Value {
	valueInfo := i.ExtractValueInfo(keyValue)
	value := i.GetCachedData(valueInfo.ValueAt, valueInfo.ValueLength)
	return &Value{
		Value:   value,
		ValueAt: valueInfo.ValueAt,
	}
}

// WriteValue ...
func (i *Impl) WriteValue(keyAt int, keyValue []byte, value []byte) *Value {
	current := i.ExtractValueInfo(keyValue)
	var valueAt int64
	if int64(len(value)) > current.ValueLength {
		valueAt = i.WriteNewBuffer(value, true)
	} else {
		valueAt = int64(i.WriteUpdatedBuffer(value, current.ValueAt))
	}

	keyValue[len(keyValue)-(keySize+uintSize)-1] = byte(len(value))
	keyValue[len(keyValue)-(keySize+uintSize+uintSize)-1] = byte(valueAt)

	file := os.NewFile(uintptr(i.fd), "temp")
	kv := keyValue[keySize : keySize+2*uintSize]
	file.WriteAt(kv, int64(keyAt)+int64(keySize))

	return &Value{
		Value:   value,
		ValueAt: valueAt,
	}
}

// WriteNewKey ...
func (i *Impl) WriteNewKey(key *NibbleBuffer) *Key {
	keyValue := make([]byte, keyTotalSize)
	for i, b := range key.Buffer {
		keyValue[i] = b
	}

	keyAt := i.WriteNewBuffer(keyValue, true)

	return &Key{
		Key:      key,
		KeyAt:    keyAt,
		KeyValue: keyValue,
	}
}

// WriteNewBranch ...
func (i *Impl) WriteNewBranch(branch []byte, branchAt int64, entryIndex int64, key *NibbleBuffer, prevAt int64, prevKey *NibbleBuffer, matchIndex int64, depth int64) *Key {

	newKey := i.WriteNewKey(key)
	keyIndex := int(key.Nibbles[matchIndex]) * entrySize
	prevIndex := int(prevKey.Nibbles[matchIndex]) * entrySize
	var buffers [][]byte
	newBranchAt := i.fileSize
	newBranch := make([]byte, branchSize)

	newBranch[keyIndex] = byte(SlotLeaf)
	newBranch[len(newBranch)-keyIndex+1+uintSize] = byte(newKey.KeyAt)
	newBranch[prevIndex] = byte(SlotLeaf)
	newBranch[len(newBranch)-prevIndex+1+uintSize] = byte(prevAt)
	buffers = append(buffers, newBranch)

	var offset int64 = 1
	for depth > 0 {
		branchIndex := int64(key.Nibbles[matchIndex-offset]) * int64(entrySize)

		newBranch = make([]byte, branchSize)
		newBranch[branchIndex] = byte(SlotBranch)
		newBranch[int64(len(newBranch))-branchIndex+1+int64(uintSize)] = byte(newBranchAt)
		buffers = append(buffers, newBranch)
		newBranchAt += int64(branchSize)

		depth--
		offset++
	}

	i.WriteNewBuffers(buffers)

	branch[entryIndex] = byte(SlotBranch)
	branch[int64(len(branch))-entryIndex+1+int64(uintSize)] = byte(newBranchAt)

	file := os.NewFile(uintptr(i.fd), "temp")
	file.WriteAt(branch[entryIndex:entryIndex+int64(entrySize)], int64(branchAt)+int64(entryIndex))

	return &Key{
		Key:      key,
		KeyAt:    newKey.KeyAt,
		KeyValue: newKey.KeyValue,
	}
}

// WriteNewLeaf ...
func (i *Impl) WriteNewLeaf(branch []byte, branchAt int64, entryIndex int64, key *NibbleBuffer) *Key {
	newKey := i.WriteNewKey(key)
	branch[entryIndex] = byte(SlotLeaf)
	branch[entryIndex+1] = byte(newKey.KeyAt)

	file := os.NewFile(uintptr(i.fd), "temp")
	b := branch[entryIndex : entryIndex+int64(entrySize)]
	file.WriteAt(b, int64(newKey.KeyAt)+int64(keySize))

	return &Key{
		Key:      key,
		KeyAt:    newKey.KeyAt,
		KeyValue: newKey.KeyValue,
	}
}

// WriteUpdatedBuffer  ...
func (i *Impl) WriteUpdatedBuffer(buffer []byte, bufferAt int64) int64 {

	file := os.NewFile(uintptr(i.fd), "temp")
	file.WriteAt(buffer, bufferAt)

	i.CacheData(bufferAt, buffer)
	return bufferAt
}

// WriteNewBuffer  ...
func (i *Impl) WriteNewBuffer(buffer []byte, withCache bool) int64 {
	startAt := i.fileSize

	file := os.NewFile(uintptr(i.fd), "temp")
	file.WriteAt(buffer, startAt)

	if withCache {
		i.CacheData(startAt, buffer)
	}

	i.fileSize += int64(len(buffer))

	return startAt
}

// WriteNewBuffers  ...
func (i *Impl) WriteNewBuffers(buffers [][]byte) int64 {
	bufferAt := i.fileSize

	for _, buffer := range buffers {
		i.CacheData(bufferAt, buffer)
		bufferAt += int64(len(buffer))
	}

	var concenatedBuffers []byte
	for _, buffer := range buffers {
		concenatedBuffers = append(concenatedBuffers, buffer...)
	}

	return i.WriteNewBuffer(concenatedBuffers, false)
}
