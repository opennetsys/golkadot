package runtime

import (
	"reflect"
	"testing"
)

func TestCrypto(t *testing.T) {
	t.Run("blake2b256", func(t *testing.T) {
		c := NewCrypto(NewHeap())
		buffer := []byte{0x01, 0x61, 0x62, 0x63, 0x64, 0x01}
		c.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		c.Heap.GrowMemory(1)
		c.Heap.Set(0, buffer)

		t.Run("stores the retrieved value", func(t *testing.T) {
			c.Blake2b256(1, 3, 6)

			if !reflect.DeepEqual(c.Heap.Get(0, 32+6), []uint8{1, 97, 98, 99, 100, 1, 189, 221, 129, 60, 99, 66, 57, 114, 49, 113, 239, 63, 238, 152, 87, 155, 148, 150, 78, 59, 177, 203, 62, 66, 114, 98, 200, 192, 104, 213, 35, 25}) {
				t.Fail()
			}
		})
	})

	t.Run("ed25519Verify", func(t *testing.T) {
		c := NewCrypto(NewHeap())
		buffer := []byte{
			0x61, 0x62, 0x63, 0x64,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			// publicKey, offset 4 + 6 = 10
			180, 114, 93, 155, 165, 255, 217, 82, 16, 250, 209, 11, 193, 10, 88, 218, 190, 190, 41, 193, 236, 252, 1, 152, 216, 214, 0, 41, 45, 138, 13, 53,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			// signature, offset 32 + 8 = 50
			209, 234, 164, 44, 182, 218, 103, 16, 205, 238, 97, 222, 123, 112, 2, 240, 24, 192, 26, 134, 11, 170, 167, 153, 141, 108, 187, 171, 241, 125, 226, 179, 244, 232, 131, 61, 44, 68, 87, 41, 141, 131, 88, 36, 175, 173, 57, 29, 12, 112, 26, 200, 247, 89, 14, 64, 224, 188, 211, 198, 233, 119, 158, 6,
		}
		c.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		c.Heap.GrowMemory(1)
		c.Heap.Set(0, buffer)

		t.Run("verifies correct signatures", func(t *testing.T) {
			if c.ED25519Verify(0, 4, 50, 10) != 0 {
				t.Fail()
			}
		})

		t.Run("fails correct signatures", func(t *testing.T) {
			if c.ED25519Verify(0, 3, 50, 10) == 0 {
				t.Fail()
			}
		})
	})

	t.Run("twox128", func(t *testing.T) {
		c := NewCrypto(NewHeap())
		buffer := []byte{0x01, 0x61, 0x62, 0x63, 0x64, 0x01}
		c.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		c.Heap.GrowMemory(1)
		c.Heap.Set(0, buffer)

		t.Run("stores the retrieved value", func(t *testing.T) {
			c.Twox128(1, 3, 16)

			if !reflect.DeepEqual(c.Heap.Get(0, 16+16), []uint8{1, 97, 98, 99, 100, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 153, 9, 119, 173, 245, 44, 188, 68, 8, 137, 50, 153, 129, 202, 169, 190}) {
				t.Fail()
			}
		})
	})

	t.Run("twox256", func(t *testing.T) {
		c := NewCrypto(NewHeap())
		buffer := []byte{0x01, 0x61, 0x62, 0x63, 0x64, 0x01}
		c.Heap.SetWASMMemory(&WasmMemory{buffer}, -1)
		c.Heap.GrowMemory(1)
		c.Heap.Set(0, buffer)

		t.Run("stores the retrieved value", func(t *testing.T) {
			c.Twox256(1, 3, 16)

			if !reflect.DeepEqual(c.Heap.Get(0, 16+32), []uint8{1, 97, 98, 99, 100, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 153, 9, 119, 173, 245, 44, 188, 68, 8, 137, 50, 153, 129, 202, 169, 190, 247, 218, 87, 112, 178, 184, 160, 83, 3, 183, 93, 149, 54, 13, 214, 43}) {
				t.Fail()
			}
		})
	})
}
