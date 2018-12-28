package db

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// FilePath ...
type FilePath struct {
	Directory string
	File      string
}

// FileTreeDB ...
type FileTreeDB struct {
	location string
}

// NewFileTreeDBDB ...
func NewFileTreeDBDB(location string) *FileTreeDB {
	return &FileTreeDB{
		location: location,
	}
}

// Close ...
func (f *FileTreeDB) Close() {
	// NOTE: noop
}

// Open ...
func (f *FileTreeDB) Open() {
	// NOTE: noop
}

// Drop ...
func (f *FileTreeDB) Drop() {
	log.Println("drop is not implemented")
}

// Empty ...
func (f *FileTreeDB) Empty() {
	log.Println("empty is not implemented")
}

// Rename ...
func (f *FileTreeDB) Rename(base, file string) {
	log.Println("rename is not implemented")
}

// Maintain ...
func (f *FileTreeDB) Maintain(fn *ProgressCB) error {
	if fn != nil {
		f := *fn
		f(&ProgressValue{
			IsCompleted: true,
			Keys:        0,
			Percent:     100,
		})
	}
	return nil
}

// Size ...
func (f *FileTreeDB) Size() int {
	log.Println("size is not implemented")
	return 0
}

// Del ...
func (f *FileTreeDB) Del(key []uint8) {
	filepath := f.getFilePath(key)

	if _, err := os.Stat(filepath.File); !os.IsNotExist(err) {
		if os.Remove(filepath.File); err != nil {
			log.Fatal(err)
		}
	}
}

// Get ...
func (f *FileTreeDB) Get(key []uint8) []uint8 {
	filepath := f.getFilePath(key)

	if _, err := os.Stat(filepath.File); os.IsNotExist(err) {
		return nil
	}

	b, err := ioutil.ReadFile(filepath.File)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// Put ...
func (f *FileTreeDB) Put(key, value []uint8) {
	filepath := f.getFilePath(key)

	if _, err := os.Stat(filepath.Directory); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Directory, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if err := ioutil.WriteFile(filepath.File, value, 0644); err != nil {
		log.Fatal(err)
	}
}

// getFilePath ...
func (f *FileTreeDB) getFilePath(key []uint8) *FilePath {
	dirDepth := 1

	// NOTE: We want to limit the number of entries in any specific directory. Split the
	// key into parts and use this to construct the path and the actual filename. We want
	// to limit the entries per directory, but at the same time minimize the number of
	// directories we need to create (when non-existent as well as the size overhead)
	re := regexp.MustCompile(`.{1,6}`)
	parts := re.FindAllString(hex.EncodeToString(key), -1)
	directory := fmt.Sprintf("%s/%s", f.location, strings.Join(parts[0:dirDepth], "/"))
	file := fmt.Sprintf("%s/%s", directory, strings.Join(parts[dirDepth:], ""))

	return &FilePath{
		Directory: directory,
		File:      file,
	}
}
