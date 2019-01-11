package runtime

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		heap := NewHeap()
		buffer := []uint8{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0}
		heap.SetWASMMemory(&WasmMemory{buffer}, -1)

		t.Run("uses set to set a part of the buffer", func(t *testing.T) {
			test := []uint8{0x6, 0x7, 0x8}

			if !reflect.DeepEqual(heap.Get(heap.Set(3, test), 3), test) {
				t.Fail()
			}
		})

		t.Run("allows set of LE u32 values", func(t *testing.T) {
			if !reflect.DeepEqual(heap.Get(heap.SetU32(Pointer(4), intToBytes(0x12345, 4)), 4), []uint8{0x45, 0x23, 0x01, 0}) {
				t.Fail()
			}
		})
	})

	t.Run("get", func(t *testing.T) {
		heap := NewHeap()
		buffer := []uint8{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0}
		heap.SetWASMMemory(&WasmMemory{buffer}, -1)

		t.Run("uses get to return data", func(t *testing.T) {
			if !reflect.DeepEqual(heap.Get(1, 3), []uint8{0x2, 0x3, 0x4}) {
				t.Fail()
			}
		})

		t.Run("allows retrieval of LE u32 values", func(t *testing.T) {
			if !reflect.DeepEqual(heap.GetU32(heap.SetU32(4, intToBytes(0x12345, 4))), intToBytes(0x12345, 4)) {
				t.Fail()
			}
		})
	})

	t.Run("fill", func(t *testing.T) {
		heap := NewHeap()
		buffer := []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
		heap.SetWASMMemory(&WasmMemory{buffer}, -1)

		t.Run("uses fill to set values", func(t *testing.T) {
			if !reflect.DeepEqual(heap.Fill(2, 5, 4), []uint8{0x0, 0x0, 0x5, 0x5, 0x5, 0x5, 0x0, 0x0}) {
				t.Fail()
			}
		})
	})

	t.Run("dup", func(t *testing.T) {
		heap := NewHeap()
		buffer := []uint8{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0}
		heap.SetWASMMemory(&WasmMemory{buffer}, -1)

		t.Run("uses dup to return a section", func(t *testing.T) {
			if !reflect.DeepEqual(heap.Dup(0, 5), []uint8{0x1, 0x2, 0x3, 0x4, 0x0}) {
				t.Fail()
			}
		})
	})

	t.Run("allocate", func(t *testing.T) {
		t.Run("returns 0 when size is 0", func(t *testing.T) {
			heap := newHeapForAllocateTest()
			if heap.Allocate(0) != 0 {
				t.Fail()
			}
		})

		t.Run("returns 0 when requested is > available", func(t *testing.T) {
			heap := newHeapForAllocateTest()
			if heap.Allocate(1024) != 0 {
				t.Fail()
			}
		})

		t.Run("returns a pointer as allocated", func(t *testing.T) {
			heap := newHeapForAllocateTest()
			if heap.Allocate(100) != 100 {
				t.Fail()
			}
		})

		t.Run("adds the allocated map to the alloc heap section", func(t *testing.T) {
			heap := newHeapForAllocateTest()
			heap.Allocate(100)

			if !reflect.DeepEqual(heap.memory.Allocated, MemoryBuffer{100: 100}) {
				t.Fail()
			}
		})

		t.Run("updates the internal offset for next allocation", func(t *testing.T) {
			heap := newHeapForAllocateTest()
			heap.Allocate(20)

			if heap.Allocate(50) != 120 {
				t.Fail()
			}
		})

		t.Run("re-allocates previous de-allocated space", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForAllocateTest()
			a := heap.Allocate(166)
			fmt.Println(a)
			if a != 3 {
				t.Fail()
			}
		})
	})

	t.Run("dealloc", func(t *testing.T) {
		t.Run("expect error when no allocation found", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForDeallocateTest()
			_, err := heap.Deallocate(456)
			if err == nil {
				t.Fail()
			}
		})

		t.Run("removes the allocation from the allocation table", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForDeallocateTest()

			if !reflect.DeepEqual(heap.memory.Allocated, MemoryBuffer{}) {
				t.Fail()
			}
		})

		t.Run("adds the allocation from the deallocated table", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForDeallocateTest()

			if !reflect.DeepEqual(heap.memory.Deallocated, MemoryBuffer{123: 456}) {
				t.Fail()
			}
		})
	})

	t.Run("freealloc", func(t *testing.T) {
		t.Run("returns 0 when matching size is not found", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForFreeAllocTest()
			if heap.FreeAlloc(501) != 0 {
				t.Fail()
			}
		})

		t.Run("returns the smallest matching slot (exact)", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForFreeAllocTest()
			if heap.FreeAlloc(120) != 3 {
				t.Fail()
			}
		})

		t.Run("returns the smallest matching slot (lesser)", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForFreeAllocTest()
			if heap.FreeAlloc(100) != 3 {
				t.Fail()
			}
		})

		t.Run("removes the previous deallocated slot", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForFreeAllocTest()
			heap.FreeAlloc(100)

			_, found := heap.memory.Deallocated[int64(3)]
			if found {
				t.Fail()
			}
		})

		t.Run("adds the allocated slot", func(t *testing.T) {
			// TODO
			t.Skip()
			heap := newHeapForFreeAllocTest()
			heap.FreeAlloc(100)

			if heap.memory.Allocated[int64(3)] != 100 {
				t.Fail()
			}
		})
	})

}

func intToBytes(n uint32, s int) []byte {
	b := make([]byte, s)
	binary.LittleEndian.PutUint32(b, n)
	return b
}

func newHeapForAllocateTest() *Heap {
	heap := NewHeap()
	heap.memory = &Memory{
		Deallocated: MemoryBuffer{
			0: 3,
			3: 166,
		},
		Allocated: make(MemoryBuffer),
		End:       110,
		Offset:    100,
		Size:      250,
	}

	return heap
}

func newHeapForDeallocateTest() *Heap {
	heap := NewHeap()
	heap.memory = &Memory{
		Deallocated: make(MemoryBuffer),
		Allocated: MemoryBuffer{
			123: 456,
		},
	}

	return heap
}

func newHeapForFreeAllocTest() *Heap {
	heap := NewHeap()
	heap.memory = &Memory{
		// NOTE: these don't make much sense as a layout, but it allows for sorting inside the actual findContaining function to find the first & smallest
		Deallocated: MemoryBuffer{
			0:   200,
			3:   120,
			4:   120,
			122: 500,
		},
		Allocated: make(MemoryBuffer),
	}

	return heap
}
