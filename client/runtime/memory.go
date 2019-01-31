package runtime

import "github.com/c3systems/go-substrate/common/triedb"

// Memory ...
type Memory struct {
	Heap *Heap
	DB   *triedb.TrieDB
}

// NewMemory ...
func NewMemory(env *Env) *Memory {
	return &Memory{
		Heap: env.Heap,
		DB:   env.DB,
	}
}

// Free ...
func (m *Memory) Free(ptr Pointer) {
	m.Heap.Deallocate(ptr)
}

// Malloc ...
func (m *Memory) Malloc(size int64) Pointer {
	return m.Heap.Allocate(size)
}

// Memcpy ...
func (m *Memory) Memcpy(dst Pointer, src Pointer, num int64) Pointer {
	return memcpy(m.Heap, dst, src, num)
}

// Memcmp ...
func (m *Memory) Memcmp(s1 Pointer, s2 Pointer, length int64) int64 {
	return memcmp(m.Heap, s1, s2, length)
}

// Memmove ...
func (m *Memory) Memmove(dst Pointer, src Pointer, num int64) Pointer {
	return memmove(m.Heap, dst, src, num)
}

// Memset ...
func (m *Memory) Memset(dst Pointer, val uint8, num int64) Pointer {
	return memset(m.Heap, dst, val, num)
}

func memcmp(heap *Heap, s1 Pointer, s2 Pointer, length int64) int64 {
	v1 := heap.Get(s1, length)
	v2 := heap.Get(s2, length)

	for index := int64(0); index < length; index++ {
		if v1[index] > v2[index] {
			return 1
		} else if v1[index] < v2[index] {
			return -1
		}
	}

	return 0
}

func memcpy(heap *Heap, dst, src Pointer, num int64) Pointer {
	return heap.Set(dst, heap.Get(src, num))
}

func memmove(heap *Heap, dst, src Pointer, num int64) Pointer {
	heap.Set(dst, heap.Dup(src, num))

	return dst
}

func memset(heap *Heap, dst Pointer, value uint8, length int64) Pointer {
	heap.Fill(dst, value, length)

	return dst
}
