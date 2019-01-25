package types

import (
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	t.Run("provides a to bn interface", func(t *testing.T) {
		if NewInt(-1234, -1).BN().Int64() != -1234 {
			t.Fail()
		}
	})

	t.Run("provides a to number interface", func(t *testing.T) {
		if NewInt(-1234, -1).Int64() != -1234 {
			t.Fail()
		}
	})

	t.Run("converts to LE from the provided value", func(t *testing.T) {
		if !reflect.DeepEqual(NewInt(1234567, -1).ToU8a(), []byte{135, 214, 18, 0, 0, 0, 0, 0}) {
			t.Fail()
		}
	})

	t.Run("converts to hex/string", func(t *testing.T) {
		u := NewIntFromHex("0x12", 16)
		if u.Hex() != "0x0012" {
			t.Fail()
		}
		if u.String() != "18" {
			t.Fail()
		}
	})

	t.Run("converts to equivalents", func(t *testing.T) {
		u := NewIntFromString("123", -1)
		if u.Int64() != 123 {
			t.Fail()
		}
	})

	t.Run("equals", func(t *testing.T) {
		u := NewInt(12345, -1)
		if !u.Equals(NewInt(12345, -1)) {
			t.Fail()
		}

		if !u.Equals(NewIntFromHex("0x3039", -1)) {
			t.Fail()
		}
	})

	t.Run("bit length", func(t *testing.T) {
		if NewInt(12345, -1).BitLen() != 64 {
			t.Fail()
		}
	})
}
