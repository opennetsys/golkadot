package runtime

import (
	"reflect"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("malloc", func(t *testing.T) {
		t.Run("allocates space", func(t *testing.T) {
			mem := newMemForMallocTest()
			mem.Malloc(50)

			if mem.Heap.Used().Allocated != 50 {
				t.Fail()
			}
		})
		t.Run("deallocates space", func(t *testing.T) {
			mem := newMemForMallocTest()
			mem.Free(mem.Malloc(666))

			if mem.Heap.Used().Deallocated != 666 {
				t.Fail()
			}
		})
	})

	t.Run("memcmp", func(t *testing.T) {
		t.Run("returns 0 when arrays are equal", func(t *testing.T) {
			mem := newMemForMemcmpTest()
			if mem.Memcmp(0, 6, 4) != 0 {
				t.Fail()
			}
		})
		t.Run("returns -1 when first is lt", func(t *testing.T) {
			mem := newMemForMemcmpTest()
			if mem.Memcmp(0, 7, 4) != -1 {
				t.Fail()
			}
		})
		t.Run("returns 1 when first is gt", func(t *testing.T) {
			mem := newMemForMemcmpTest()
			if mem.Memcmp(1, 6, 4) != 1 {
				t.Fail()
			}
		})
	})

	t.Run("memcpy", func(t *testing.T) {
		t.Run("copies the src to dst", func(t *testing.T) {
			mem := newMemForMemcpyTest()
			mem.Memcpy(4, 2, 3)

			if !reflect.DeepEqual(mem.Heap.Get(0, 8), []uint8{1, 2, 3, 4, 3, 4, 5, 8}) {
				t.Fail()
			}
		})
	})

	t.Run("memmove", func(t *testing.T) {
		t.Run("copies the src to dst", func(t *testing.T) {
			mem := newMemForMemmoveTest()
			mem.Memmove(0, 2, 2)

			if !reflect.DeepEqual(mem.Heap.Get(0, 5), []uint8{3, 4, 3, 4, 5}) {
				t.Fail()
			}
		})
	})

	t.Run("memset", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			mem := newMemForMemsetTest()
			mem.Memset(0, 2, 3)

			if !reflect.DeepEqual(mem.Heap.Get(0, 5), []uint8{2, 2, 2, 0, 0}) {
				t.Fail()
			}
		})
	})
}

func newMemForMallocTest() *Memory {
	rt := NewEnv(nil)
	rt.Heap.SetWASMMemory(&WasmMemory{Buffer: make([]uint8, 1024*1024)}, -1)

	return NewMemory(rt)
}

func newMemForMemcmpTest() *Memory {
	rt := NewEnv(nil)
	rt.Heap.SetWASMMemory(&WasmMemory{Buffer: []uint8{0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5}}, -1)

	return NewMemory(rt)
}

func newMemForMemcpyTest() *Memory {
	rt := NewEnv(nil)
	rt.Heap.SetWASMMemory(&WasmMemory{Buffer: []uint8{1, 2, 3, 4, 5, 6, 7, 8}}, -1)

	return NewMemory(rt)
}

func newMemForMemmoveTest() *Memory {
	rt := NewEnv(nil)
	rt.Heap.SetWASMMemory(&WasmMemory{Buffer: []uint8{1, 2, 3, 4, 5}}, -1)

	return NewMemory(rt)
}

func newMemForMemsetTest() *Memory {
	rt := NewEnv(nil)
	rt.Heap.SetWASMMemory(&WasmMemory{Buffer: make([]uint8, 5)}, -1)

	return NewMemory(rt)
}
