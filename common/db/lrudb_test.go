package db

import (
	"reflect"
	"testing"
)

func TestLruDB(t *testing.T) {
	memoryDB := NewMemoryDB(&BaseOptions{})
	baseDB := BaseDB(memoryDB)
	lrudb := NewLruDB(baseDB, -1)

	t.Run("retrieves an item from the backing when not available (caching it)", func(t *testing.T) {
		key := []uint8("test1")
		value := []uint8("value1")
		memoryDB.Put(key, value)

		if !reflect.DeepEqual(lrudb.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("replaces an item (caching and backing)", func(t *testing.T) {
		key := []uint8("test1")
		value := []uint8("test")
		lrudb.Put(key, value)

		if !reflect.DeepEqual(lrudb.Get(key), value) {
			t.Fail()
		}
		if !reflect.DeepEqual(memoryDB.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("retrieves item from LRU when available", func(t *testing.T) {
		key := []uint8("test1")
		value := []uint8("test")
		memoryDB.Del(key)

		if !reflect.DeepEqual(lrudb.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("puts item both in LRU and backing", func(t *testing.T) {
		key := []uint8("test0")
		value := []uint8("value0")
		lrudb.Put(key, value)

		if !reflect.DeepEqual(memoryDB.Get(key), value) {
			t.Fail()
		}
	})

	t.Run("deletes an item from both backing and db", func(t *testing.T) {
		key := []uint8("test0")
		lrudb.Del(key)

		if memoryDB.Get(key) != nil {
			t.Fail()
		}
		if lrudb.Get(key) != nil {
			t.Fail()
		}
	})
}
