package runtime

import (
	"reflect"
	"testing"

	"github.com/opennetsys/go-substrate/common/db"
)

func TestStorage(t *testing.T) {
	t.Run("data", func(t *testing.T) {
		memdb := db.NewMemoryDB(&db.BaseOptions{})
		basedb := db.BaseDB(memdb)
		txdbt := db.NewTransactionDB(&basedb)
		txdb := db.TXDB(txdbt)

		s := NewData(NewHeap(), txdb)
		buffer := []byte{0xff, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f}
		s.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		s.Heap.GrowMemory(1)
		s.Heap.Set(0, buffer)

		dbvalue := []uint8{0x1, 0x2, 0x3, 0x4, 0x5}
		//s.DB.Put([]uint8{255}, dbvalue)
		s.DB.Put([]uint8{83, 97, 121}, dbvalue)
		s.DB.Put([]uint8{}, dbvalue)

		t.Run("retrieves the correct value from storage", func(t *testing.T) {
			// TODO
			t.Skip()
			length := s.GetStorageInto(1, 3, 3, 3, 0)
			if length != U32Max {
				t.Fail()
			}

			if !reflect.DeepEqual(s.Heap.Get(1, 3), []uint8{0x53, 0x61, 0x79}) {
				t.Fail()
			}
		})

		t.Run("retrieves the full value when length >= available", func(t *testing.T) {
			// TODO
			t.Skip()
			length := s.GetStorageInto(1, 0, 3, 10, 0)
			if length != 5 {
				t.Fail()
			}
		})

		t.Run("retrieves a partial value when length < available", func(t *testing.T) {
			// TODO
			t.Skip()
			length := s.GetStorageInto(1, 0, 3, 3, 0)
			if length != 3 {
				t.Fail()
			}
		})

		t.Run("retrieves a partial value with offset", func(t *testing.T) {
			// TODO
			t.Skip()
			length := s.GetStorageInto(1, 0, 3, 13, 2)
			if length != 3 {
				t.Fail()
			}
		})

		t.Run("retrieves zero value when not available", func(t *testing.T) {
			// TODO
			t.Skip()
			length := s.GetStorageInto(0, 3, 3, 5, 0)
			if length != 5 {
				t.Fail()
			}
		})
	})

	t.Run("storage", func(t *testing.T) {
		t.Skip()
		s := NewStorage(NewHeap(), nil)
		buffer := []byte{0xff, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f}
		s.Data.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		s.Data.Heap.GrowMemory(1)
		s.Data.Heap.Set(0, buffer)

		t.Run("retrieves the correct value from storage", func(t *testing.T) {
			// TODO
			t.Skip()
			//s.GetStorageInto(1, 3, 3, 3, 0)
		})
	})
}
