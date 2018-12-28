package db

import (
	"reflect"
	"testing"
)

func TestTransactionDB(t *testing.T) {
	memoryDB := NewMemoryDB(&BaseOptions{})
	baseDB := BaseDB(memoryDB)
	txdb := NewTransactionDB(&baseDB)

	t.Run("passed through values immediately when not in transaction", func(t *testing.T) {
		key := []uint8("test0")
		value := []uint8("value0")
		txdb.Put(key, value)

		if !reflect.DeepEqual(memoryDB.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("commits when transaction passes (result = true, put)", func(t *testing.T) {
		key := []uint8("test1")
		value := []uint8("value1")

		ok, err := txdb.Transaction(func() bool {
			txdb.Put(key, value)
			return true
		})
		if err != nil {
			t.Fail()
		}
		if !ok {
			t.Fail()
		}
		if !reflect.DeepEqual(memoryDB.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("commits when transaction passes (result = true, del)", func(t *testing.T) {
		key := []uint8("test0")

		ok, err := txdb.Transaction(func() bool {
			txdb.Del(key)
			return true
		})
		if err != nil {
			t.Fail()
		}
		if !ok {
			t.Fail()
		}
		if memoryDB.Get(key) != nil {
			t.Fail()
		}
	})

	t.Run("does not commit when transaction fails (result = false)", func(t *testing.T) {
		key := []uint8("test2")
		value := []uint8("value2")

		ok, err := txdb.Transaction(func() bool {
			txdb.Put(key, value)
			return false
		})
		if err != nil {
			t.Fail()
		}
		if ok {
			t.Fail()
		}

		if memoryDB.Get(key) != nil {
			t.Fail()
		}
	})

	t.Run("inside the transaction, the value is set", func(t *testing.T) {
		key := []uint8("test2")
		value := []uint8("value2")

		ok, err := txdb.Transaction(func() bool {
			txdb.Put(key, value)

			if !reflect.DeepEqual(txdb.Get(key), value) {
				t.Fail()
			}

			return false
		})
		if err != nil {
			t.Fail()
		}
		if ok {
			t.Fail()
		}
		if memoryDB.Get(key) != nil {
			t.Fail()
		}
	})
}
