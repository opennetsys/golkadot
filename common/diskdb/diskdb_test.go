package db

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/opennetsys/golkadot/common/db"
)

func TestDiskDB(t *testing.T) {
	setUp()
	defer cleanUp()

	// Creat diskDb instance that wraps LruDb with backing and Lru cache
	diskDb := NewDiskDB(getLocation(), "store.db", nil)

	// Creat txDb instance that uses diskDb
	baseDb := db.BaseDB(diskDb)
	txDb := db.NewTransactionDB(&baseDb)

	// Open the disk db backing database. Clears the Lru cache
	diskDb.Open()

	// Declare key/value pair to allocate to store under a the key
	key := []uint8("key")
	value := []uint8("some value")

	// Store key/value pair in disk db backing and also in Lru cache
	diskDb.Put(key, value)

	// Retrieve value for key from disk db. Returns cached value if key
	// in cache, otherwise returns backing value for key and stores
	// this latest retrieved key/value pair in Lru cache
	if !reflect.DeepEqual(diskDb.Get(key), value) {
		t.Fail()
	}

	// Delete key/value pair from disk db backing and set the key to null in Lru cache
	// diskDb.Del(key) // Del not implemented

	// Transaction to Store key/value pair in transaction db
	isTxSuccess, err := txDb.Transaction(func() bool {
		txDb.Put(key, value)

		// Boolean to indicate whether transaction was successful or not
		return true
	})
	if err != nil {
		t.Error(err)
	}

	if isTxSuccess != true {
		t.Fail()
	}

	// Retrieve transaction value from diskDb db
	if !reflect.DeepEqual(diskDb.Get(key), value) {
		t.Fail()
	}

	// Close the diskDb database and clear Lru cache
	diskDb.Close()
}

func setUp() {
	testpath := getLocation()
	if _, err := os.Stat(testpath); os.IsNotExist(err) {
		err := os.MkdirAll(testpath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func cleanUp() {
	testpath := getLocation()
	if _, err := os.Stat(testpath); !os.IsNotExist(err) {
		err := os.RemoveAll(testpath)
		if err != nil {
			panic(err)
		}
	}
}

func getLocation() string {
	dirpath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/test", dirpath)
}
