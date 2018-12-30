package fileflatdb

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/c3systems/go-substrate/db"
)

var uintSize = 5
var keySize = 32
var keyTotalSize = keySize + uintSize + uintSize
var defaultEntryNum = 16 // nibbles, 256 for bytes (where serialize would be noop)
var defaultEntrySize = 1 + uintSize
var defaultBranchSize = defaultEntryNum * defaultEntrySize
var defaultFile = "store.db"
var lruBranchCount = 16384 // * 96 = bytes
var lruDataCount = 8192

// File ...
type File struct {
	Serializer
	fd       int
	fileSize int64
	path     string
	file     string
}

// NewFile ...
func NewFile(base, file string, options *db.BaseDBOptions) *File {

	var isCompressed bool
	if options != nil && options.IsCompressed {
		isCompressed = true
	}

	f := &File{
		fd:       -1,
		fileSize: 0,
		path:     fmt.Sprintf("%s/%s", base, file),
		file:     file,
	}

	f.IsCompressed = isCompressed

	if _, err := os.Stat(base); os.IsNotExist(err) {
		if err := os.MkdirAll(base, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	return f
}

// AssertOpen ...
func (f *File) AssertOpen(open bool) {
	var test bool
	if open {
		test = f.fd != -1
	} else {
		test = f.fd == -1
	}

	if !test {
		if open {
			log.Fatal("expected an open database")
		} else {
			log.Fatal("expected a closed database")
		}
	}
}

// Close ...
func (f *File) Close() {
	// close file descriptor
	if err := syscall.Close(int(f.fd)); err != nil {
		log.Fatal(err)
	}

	f.fd = -1
}

// Open ...
func (f *File) Open(filepath string, startEmpty bool) {
	_, err := os.Stat(filepath)
	isExisting := !os.IsNotExist(err)

	if !isExisting || startEmpty {
		if isExisting {
			os.Rename(filepath, fmt.Sprintf("%s.%d", filepath, time.Now().Unix()))
		}

		b := make([]byte, defaultBranchSize)
		ioutil.WriteFile(filepath, b, 0644)
	}

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	stat, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}

	f.fd = int(file.Fd())
	f.fileSize = stat.Size()
}
