package types

import (
	"reflect"
	"testing"
)

func TestU8Fixed(t *testing.T) {
	t.Run("can decode from uint8 slice", func(t *testing.T) {
		if NewU8Fixed([]byte{1, 2, 3, 4, 5}, -1).String() != "0x0102030405" {
			t.Fail()
		}
	})

	t.Run("can decode from hex string", func(t *testing.T) {
		if NewU8FixedFromHex("0x0102030405", -1).String() != "0x0102030405" {
			t.Fail()
		}
	})

	t.Run("test if empty value", func(t *testing.T) {
		if !NewU8Fixed([]uint8{0, 0, 0}, -1).IsEmpty() {
			t.Fail()
		}

		if NewU8Fixed([]uint8{0, 1, 0}, -1).IsEmpty() {
			t.Fail()
		}
	})

	t.Run("contains the length of the elements", func(t *testing.T) {
		if NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Len() != 5 {
			t.Fail()
		}
	})

	t.Run("correctly encodes length", func(t *testing.T) {
		if NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).EncodedLen() != 5 {
			t.Fail()
		}
	})

	t.Run("implements subarray correctly", func(t *testing.T) {
		if !reflect.DeepEqual(NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Sub(1, 3), []uint8{2, 3}) {
			t.Fail()
		}
	})

	t.Run("compares against other U8a", func(t *testing.T) {
		if !NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Equals(NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1)) {
			t.Fail()
		}
	})

	t.Run("compares against other U8a (non-length)", func(t *testing.T) {
		if NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Equals(NewU8Fixed([]uint8{1, 2, 3, 4}, -1)) {
			t.Fail()
		}
	})

	t.Run("compares against other U8a (mismatch)", func(t *testing.T) {
		if NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Equals(NewU8Fixed([]uint8{1, 2, 3, 4, 5, 6}, -1)) {
			t.Fail()
		}
	})

	t.Run("compares against hex inputs", func(t *testing.T) {
		if !NewU8Fixed([]uint8{1, 2, 3, 4, 5}, -1).Equals(NewU8FixedFromHex("0x0102030405", -1)) {
			t.Fail()
		}
	})
}
