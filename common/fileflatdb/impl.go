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
	cache *Cache
}

// NewImpl ...
func NewImpl(cache *Cache) *Impl {
	return &Impl{
		cache: cache,
	}
}

// GetKeyValue ...
func (i *Impl) GetKeyValue(keyAt int64) []byte {
	return i.cache.GetCachedData(keyAt, int64(keyTotalSize))
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
	prevKey := i.cache.file.serializer.SerializeKey(keyValue[0:keySize])
	matchIndex := keyIndex

	for matchIndex < keySize {
		if matchIndex >= len(prevKey.Nibbles) || matchIndex >= len(key.Nibbles) {
			break
		}
		if prevKey.Nibbles[matchIndex] != key.Nibbles[matchIndex] {
			break
		}

		matchIndex++
	}

	if matchIndex != keySize {
		if doCreate {
			return i.WriteNewBranch(branch, int64(branchAt), int64(entryIndex), key, int64(keyAt.Uint64()), prevKey, uint64(matchIndex), int64(matchIndex-keyIndex-1))
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
	var entryIndex int
	if len(key.Nibbles) > 0 {
		entryIndex = int(key.Nibbles[keyIndex]) * entrySize
	}
	branch := i.cache.GetCachedBranch(branchAt)
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
	value := i.cache.GetCachedData(valueInfo.ValueAt, valueInfo.ValueLength)
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

	writeUIntBE(keyValue, int64(len(value)), int64(keySize), int64(uintSize))
	writeUIntBE(keyValue, int64(valueAt), int64(keySize)+int64(uintSize), int64(uintSize))
	file := os.NewFile(i.fd(), "fileflatdb")
	_, err := file.WriteAt(keyValue[keySize:keySize+2*uintSize], int64(keyAt)+int64(keySize))
	if err != nil {
		log.Fatal(err)
	}

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
func (i *Impl) WriteNewBranch(branch []byte, branchAt int64, entryIndex int64, key *NibbleBuffer, prevAt int64, prevKey *NibbleBuffer, matchIndex uint64, depth int64) *Key {

	newKey := i.WriteNewKey(key)

	if matchIndex >= uint64(len(key.Nibbles)) && len(key.Nibbles) > 0 {
		matchIndex = uint64(len(key.Nibbles) - 1)
	}

	var keyIndex int
	if len(key.Nibbles) > 0 {
		keyIndex = int(key.Nibbles[matchIndex]) * entrySize
	}

	var prevIndex int
	if len(prevKey.Nibbles) > 0 {
		prevIndex = int(prevKey.Nibbles[matchIndex]) * entrySize
	}
	var buffers [][]byte
	newBranchAt := i.cache.file.fileSize
	newBranch := make([]byte, branchSize)

	newBranch[keyIndex] = byte(SlotLeaf)
	writeUIntBE(newBranch, int64(newKey.KeyAt), int64(keyIndex)+1, int64(uintSize))

	newBranch[prevIndex] = byte(SlotLeaf)
	writeUIntBE(newBranch, int64(prevAt), int64(prevIndex)+1, int64(uintSize))

	buffers = append(buffers, newBranch)

	var offset int64 = 1
	for depth > 0 {
		branchIndex := int64(key.Nibbles[int64(matchIndex)-offset]) * int64(entrySize)

		newBranch = make([]byte, branchSize)
		newBranch[branchIndex] = byte(SlotBranch)

		writeUIntBE(newBranch, int64(newBranchAt), int64(branchIndex)+1, int64(uintSize))

		buffers = append(buffers, newBranch)
		newBranchAt += int64(branchSize)

		depth--
		offset++
	}

	i.WriteNewBuffers(buffers)

	branch[entryIndex] = byte(SlotBranch)

	writeUIntBE(branch, int64(newBranchAt), int64(entryIndex)+1, int64(uintSize))

	file := os.NewFile(i.fd(), "fileflatdb")
	_, err := file.WriteAt(branch[entryIndex:entryIndex+int64(entrySize)], int64(branchAt)+int64(entryIndex))
	if err != nil {
		log.Fatal(err)
	}

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
	writeUIntBE(branch, int64(newKey.KeyAt), int64(entryIndex)+1, int64(uintSize))
	file := os.NewFile(i.fd(), "fileflatdb")
	_, err := file.WriteAt(branch[entryIndex:entryIndex+int64(entrySize)], int64(branchAt)+int64(entryIndex))
	if err != nil {
		log.Fatal(err)
	}

	return &Key{
		Key:      key,
		KeyAt:    newKey.KeyAt,
		KeyValue: newKey.KeyValue,
	}
}

// WriteUpdatedBuffer  ...
func (i *Impl) WriteUpdatedBuffer(buffer []byte, bufferAt int64) int64 {

	file := os.NewFile(i.fd(), "fileflatdb")
	_, err := file.WriteAt(buffer, bufferAt)
	if err != nil {
		log.Fatal(err)
	}

	i.cache.CacheData(bufferAt, buffer)
	return bufferAt
}

// WriteNewBuffer  ...
func (i *Impl) WriteNewBuffer(buffer []byte, withCache bool) int64 {
	startAt := i.cache.file.fileSize

	file := os.NewFile(i.fd(), "fileflatdb")
	_, err := file.WriteAt(buffer, startAt)
	if err != nil {
		log.Fatalf("error writing new buffer to file: %v \n", err)
	}

	if withCache {
		i.cache.CacheData(startAt, buffer)
	}

	i.cache.file.fileSize += int64(len(buffer))

	return startAt
}

// WriteNewBuffers  ...
func (i *Impl) WriteNewBuffers(buffers [][]byte) int64 {
	bufferAt := i.cache.file.fileSize

	for _, buffer := range buffers {
		i.cache.CacheData(bufferAt, buffer)
		bufferAt += int64(len(buffer))
	}

	var concenatedBuffers []byte
	for _, buffer := range buffers {
		concenatedBuffers = append(concenatedBuffers, buffer...)
	}

	return i.WriteNewBuffer(concenatedBuffers, false)
}

// fd ...
func (i *Impl) fd() uintptr {
	filepath := i.cache.file.path

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Fatalf("[fileflatdb] file %q does not exist", filepath)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatalf("[fileflatdb] error opening file: %v", err)
	}
	return file.Fd()
}
