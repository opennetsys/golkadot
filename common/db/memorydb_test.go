package db

import (
	"reflect"
	"testing"
)

func TestMemoryDB(t *testing.T) {
	memoryDb := NewMemoryDB(nil)
	baseDb := BaseDB(memoryDb)
	txDb := NewTransactionDB(&baseDb)

	// Open the memory database
	memoryDb.Open()

	// Declare key/value pair to allocate to store under a the key
	key := []uint8("key")
	value := []uint8("some value")

	// Store key/value pair in memory db
	memoryDb.Put(key, value)

	// Retrieve value for key from memory db
	if !reflect.DeepEqual(memoryDb.Get(key), value) {
		t.Fail()
	}

	// Delete key/value pair from memory db
	memoryDb.Del(key)

	if memoryDb.Get(key) != nil {
		t.Fail()
	}

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

	// Retrieve transaction value from memory db
	if !reflect.DeepEqual(memoryDb.Get(key), value) {
		t.Fail()
	}

	// Close the memory database
	memoryDb.Close()
}
