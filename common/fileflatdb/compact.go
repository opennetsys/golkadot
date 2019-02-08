package fileflatdb

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"syscall"
	"time"

	"github.com/opennetsys/golkadot/common/db"
)

// Compact ...
type Compact struct {
	fd   int64
	file string
}

// NewCompact ...
func NewCompact(file string) *Compact {
	return &Compact{
		fd:   -1,
		file: file,
	}
}

// Maintain ...
func (c *Compact) Maintain(fn *db.ProgressCB) {
	if c.fd != -1 {
		log.Fatal("database cannot be open for compacting")
	}

	start := time.Now().Unix()
	newFile := fmt.Sprintf("%s.compacted", c.file)
	newFd := c.Open(newFile, true)
	oldFd := c.Open(c.file, false)
	keys := c.Compact(*fn, newFd, oldFd)

	closeFd(oldFd)
	closeFd(newFd)

	newStat, err := os.Stat(newFile)
	if err != nil {
		log.Fatal(err)
	}
	oldStat, err := os.Stat(c.file)
	if err != nil {
		log.Fatal(err)
	}
	percentage := 100 * (newStat.Size() / oldStat.Size())
	sizeMB := newStat.Size() / (1024 * 1024)
	elapsed := time.Now().Unix() - start

	err = os.Remove(c.file)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(newFile, c.file)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("compacted in %d, %dk keys, %dMB (%d%%)", elapsed, keys/1e3, sizeMB, percentage)
}

// Open ...
func (c *Compact) Open(file string, startEmpty bool) int {
	_, err := os.Stat(file)
	isExisting := !os.IsNotExist(err)
	if !isExisting || startEmpty {
		data := make([]byte, branchSize)
		err := ioutil.WriteFile(file, data, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return int(f.Fd())
}

// doCompact ...
func (c *Compact) doCompact(keys *int, percent int, fn db.ProgressCB, newFd, oldFd int, newAt int, oldAt int, depth int) {
	increment := (100 / float64(entryNum)) / math.Pow(float64(entryNum), float64(depth))

	for index := 0; index < entryNum; index++ {
		entry := c.CompactReadEntry(oldFd, oldAt, index)
		dataAt := new(big.Int)
		dataAt.SetBytes(entry[1 : 1+uintSize])
		entryType := entry[0]

		if int(entryType) == SlotEmpty {
			percent += int(increment)
		} else if int(entryType) == SlotLeaf {
			key, value := c.CompactReadKey(oldFd, int64(dataAt.Uint64()))
			keyAt := c.CompactWriteKey(newFd, key, value)

			c.CompactUpdateLink(newFd, newAt, index, keyAt, SlotLeaf)

			newKeys := *keys + 1
			keys = &newKeys
			percent += int(increment)
		} else if int(entryType) == SlotBranch {
			headerAt := c.CompactWriteHeader(newFd, newAt, index)

			c.doCompact(keys, percent, fn, newFd, oldFd, int(headerAt), int(dataAt.Uint64()), depth+1)
		} else {
			log.Fatalf("Unknown entry type %d", entryType)
		}

		var isCompleted bool
		if depth == 0 && index == entryNum-1 {
			isCompleted = true
		}

		if fn != nil {
			fn(&db.ProgressValue{
				IsCompleted: isCompleted,
				Keys:        *keys,
				Percent:     percent,
			})
		}
	}
}

// Compact ...
func (c *Compact) Compact(fn db.ProgressCB, newFd, oldFd int) int {
	var keys int
	var percent int

	c.doCompact(&keys, percent, fn, newFd, oldFd, 0, 0, 0)

	return keys
}

// CompactReadEntry ...
func (c *Compact) CompactReadEntry(fd int, at int, index int) []byte {
	entry := make([]byte, entrySize)
	entryAt := at + (index * entrySize)

	file := os.NewFile(uintptr(fd), "temp")
	_, err := file.ReadAt(entry, int64(entryAt))
	if err != nil {
		log.Fatal(err)
	}

	return entry
}

// CompactReadKey ...
func (c *Compact) CompactReadKey(fd int, at int64) ([]byte, []byte) {
	key := make([]byte, keyTotalSize)
	file := os.NewFile(uintptr(fd), "temp")
	_, err := file.ReadAt(key, at)
	if err != nil {
		log.Fatal(err)
	}

	valueLength := new(big.Int)
	valueLength.SetBytes(key[keySize : keySize+uintSize])

	valueAt := new(big.Int)
	valueAt.SetBytes(key[(keySize + uintSize) : (keySize+uintSize)+uintSize])

	value := make([]byte, valueLength.Uint64())
	_, err = file.ReadAt(value, int64(valueAt.Uint64()))
	if err != nil {
		log.Fatal(err)
	}

	return key, value
}

// CompactWriteKey ...
func (c *Compact) CompactWriteKey(fd int, key, value []byte) int64 {
	file := os.NewFile(uintptr(fd), "temp")
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	valueAt := stat.Size()
	keyAt := valueAt + int64(len(value))

	writeUIntBE(key, int64(valueAt), int64(keySize)+int64(uintSize), int64(uintSize))

	_, err = file.WriteAt(value, valueAt)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteAt(key, keyAt)
	if err != nil {
		log.Fatal(err)
	}

	return keyAt
}

// CompactUpdateLink ...
func (c *Compact) CompactUpdateLink(fd int, at int, index int, pointer int64, kind int) {
	entry := make([]byte, entrySize)
	entryAt := at + (index * entrySize)

	entry[0] = byte(kind)
	writeUIntBE(entry, int64(pointer), int64(1), int64(uintSize))

	file := os.NewFile(uintptr(fd), "temp")
	_, err := file.WriteAt(entry, int64(entryAt))
	if err != nil {
		log.Fatal(err)
	}
}

// CompactWriteHeader ...
func (c *Compact) CompactWriteHeader(fd int, at int, index int) int64 {
	file := os.NewFile(uintptr(fd), "temp")
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	headerAt := stat.Size()

	header := make([]byte, branchSize)
	_, err = file.WriteAt(header, headerAt)
	if err != nil {
		log.Fatal(err)
	}

	c.CompactUpdateLink(fd, at, index, headerAt, SlotBranch)

	return headerAt
}

func closeFd(fd int) {
	// close file descriptor
	if err := syscall.Close(fd); err != nil {
		log.Fatal(err)
	}
}
