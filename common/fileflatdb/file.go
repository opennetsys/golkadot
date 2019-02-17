package fileflatdb

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	db "github.com/opennetsys/golkadot/common/db"
	dirutil "github.com/opennetsys/golkadot/common/dirutil"
)

var uintSize = 5
var keySize = 32
var keyTotalSize = keySize + uintSize + uintSize
var entryNum = 16 // nibbles, 256 for bytes (where serialize would be noop)
var entrySize = 1 + uintSize
var branchSize = entryNum * entrySize
var defaultFile = "store.db"
var lruBranchCount = 16384 // * 96 = bytes
var lruDataCount = 8192

// File ...
type File struct {
	serializer *Serializer
	fd         uintptr
	fileSize   int64
	path       string
	file       string
}

// NewFile ...
func NewFile(base, file string, options *db.BaseDBOptions) *File {
	if file == "" {
		file = defaultFile
	}

	var isCompressed bool
	if options != nil && options.IsCompressed {
		isCompressed = true
	}

	filepath := dirutil.NormalizePath(fmt.Sprintf("%s/%s", base, file))

	f := &File{
		serializer: NewSerializer(),
		fd:         0,
		fileSize:   0,
		path:       filepath,
		file:       file,
	}

	f.serializer.IsCompressed = isCompressed

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
		test = f.fd != 0
	} else {
		test = f.fd == 0
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

	f.fd = 0
}

// Open ...
func (f *File) Open(filepath string, startEmpty bool) {
	filepath = dirutil.NormalizePath(filepath)

	_, err := os.Stat(filepath)
	isExisting := !os.IsNotExist(err)

	if !isExisting || startEmpty {
		if isExisting {
			os.Rename(filepath, fmt.Sprintf("%s.%d", filepath, time.Now().Unix()))
		}

		b := make([]byte, branchSize)

		paths := strings.Split(filepath, "/")
		folderPath := strings.Join(paths[:len(paths)-1], "/")
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		ioutil.WriteFile(filepath, b, 0644)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	stat, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}

	f.fd = file.Fd()
	f.fileSize = stat.Size()
}
