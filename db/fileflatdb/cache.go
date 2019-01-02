package fileflatdb

import (
	"log"
	"os"
)

// LruMap ...
type LruMap map[int64][]byte

// Cache ...
type Cache struct {
	file      *File
	lruBranch LruMap
	lruData   LruMap
}

// NewCache ...
func NewCache(file *File) *Cache {
	return &Cache{
		file:      file,
		lruBranch: make(LruMap, lruBranchCount),
		lruData:   make(LruMap, lruDataCount),
	}
}

// CacheBranch ...
func (c *Cache) CacheBranch(branchAt int64, branch []byte) {
	c.lruBranch[branchAt] = branch
}

// CacheData ...
func (c *Cache) CacheData(dataAt int64, data []byte) {
	c.lruData[dataAt] = data
}

// GetCachedBranch ...
func (c *Cache) GetCachedBranch(branchAt int64) []byte {
	branch, found := c.lruBranch[branchAt]
	if !found {
		branch = make([]byte, branchSize)
		fd := os.NewFile(uintptr(c.file.fd), "temp")
		_, err := fd.ReadAt(branch, branchAt)
		if err != nil {
			log.Fatal(err)
		}

		c.CacheBranch(branchAt, branch)
	}

	return branch
}

// GetCachedData ...
func (c *Cache) GetCachedData(dataAt int64, length int64) []byte {
	data, found := c.lruData[dataAt]
	if !found {
		data = make([]byte, length)
		fd := os.NewFile(uintptr(c.file.fd), "temp")
		_, err := fd.ReadAt(data, dataAt)
		if err != nil {
			log.Fatal(err)
		}

		c.CacheData(dataAt, data)
	}

	return data
}
