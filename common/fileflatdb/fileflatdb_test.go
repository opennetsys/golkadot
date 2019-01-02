package fileflatdb

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFileFlatDB(t *testing.T) {
	setUp()
	defer cleanUp()

	store := NewFileFlatDB(getLocation(), "store.db")
	store.Open()

	keyA := []byte{0x10, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	keyB := []byte{0x20, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	keyC := []byte{0x10, 2, 6, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	keyD := []byte{0x10, 2, 6, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	keyE := []byte{0x10, 2, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	keyF := []byte{0x50, 2, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	valA := []byte{0x42, 1, 0x69}
	valB := []byte{0x42, 1, 2, 0x69}
	valC := []byte{0x42, 1, 2, 3, 0x69}
	valD := []byte{0x42, 1, 2, 3, 4, 0x69}
	valE := []byte{0x42, 1, 2, 3, 4, 5, 0x69}
	valF := []byte{0x42, 1, 2, 3, 4, 5, 6, 0x69}

	t.Run("writes an entry", func(t *testing.T) {
		store.Put(keyA, valA)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
	})

	t.Run("writes an entry (additional)", func(t *testing.T) {
		store.Put(keyB, valB)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyB), valB) {
			t.Fail()
		}
	})

	t.Run("writes an entry (expanding the tree)", func(t *testing.T) {
		store.Put(keyC, valC)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyB), valB) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyC), valC) {
			t.Fail()
		}
	})

	t.Run("writes an entry (expanding the tree, again)", func(t *testing.T) {
		store.Put(keyD, valD)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyB), valB) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyC), valC) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyD), valD) {
			t.Fail()
		}
	})

	t.Run("writes an entry (expanding the tree, yet again)", func(t *testing.T) {
		store.Put(keyE, valE)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyB), valB) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyC), valC) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyD), valD) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyE), valE) {
			t.Fail()
		}
	})

	t.Run("writes an entry, expanding the top-level", func(t *testing.T) {
		store.Put(keyF, valF)
		if !reflect.DeepEqual(store.Get(keyA), valA) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyB), valB) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyC), valC) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyD), valD) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyE), valE) {
			t.Fail()
		}
		if !reflect.DeepEqual(store.Get(keyF), valF) {
			t.Fail()
		}
	})

	t.Run("overrides with smaller values", func(t *testing.T) {
		store.Put(keyF, valA)
		if !reflect.DeepEqual(store.Get(keyF), valA) {
			t.Fail()
		}
	})

	t.Run("overrides with larger values", func(t *testing.T) {
		store.Put(keyA, valF)
		if !reflect.DeepEqual(store.Get(keyA), valF) {
			t.Fail()
		}
	})

	store.Close()
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
