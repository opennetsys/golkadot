package common

import (
	"math/big"
	"reflect"
	"testing"
)

func TestTypeToU8aj(t *testing.T) {
	t.Run("int64 to []uint8", func(t *testing.T) {
		if !reflect.DeepEqual(TypeToU8a(int64(1234567), false), []uint8{135, 214, 18, 0, 0, 0, 0, 0}) {
			t.Fail()
		}
	})
	t.Run("int32 to []uint8", func(t *testing.T) {
		if !reflect.DeepEqual(TypeToU8a(int32(1234567), false), []uint8{135, 214, 18, 0}) {
			t.Fail()
		}
	})

	t.Run("int8 to []uint8", func(t *testing.T) {
		if !reflect.DeepEqual(TypeToU8a(int8(9), false), []uint8{9}) {
			t.Fail()
		}
	})
}

func TestTypeBitLen(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		if TypeBitLen(uint8(9)) != 8 {
			t.Fail()
		}
	})

	t.Run("int64", func(t *testing.T) {
		if TypeBitLen(int64(1234567)) != 64 {
			t.Fail()
		}
	})

	t.Run("big int", func(t *testing.T) {
		if TypeBitLen(big.NewInt(int64(1234567))) != 21 {
			t.Fail()
		}
	})
}

func TestTypeToStrign(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		if TypeToString(uint8(9)) != "9" {
			t.Fail()
		}
	})

	t.Run("int64", func(t *testing.T) {
		if TypeToString(int64(1234567)) != "1234567" {
			t.Fail()
		}
	})

	t.Run("big num", func(t *testing.T) {
		if TypeToString(big.NewInt(int64(1234567))) != "1234567" {
			t.Fail()
		}
	})
}

func TestTypeIsZeroValue(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		if TypeIsZeroValue(uint8(9)) {
			t.Fail()
		}
		if !TypeIsZeroValue(uint8(0)) {
			t.Fail()
		}
	})

	t.Run("int64", func(t *testing.T) {
		if TypeIsZeroValue(int64(1234567)) {
			t.Fail()
		}
		if !TypeIsZeroValue(int64(0)) {
			t.Fail()
		}
	})

	t.Run("big num", func(t *testing.T) {
		if TypeIsZeroValue(big.NewInt(int64(1234567))) {
			t.Fail()
		}
		if !TypeIsZeroValue(big.NewInt(0)) {
			t.Fail()
		}
	})
}

func TestTypeToBytes(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		if !reflect.DeepEqual(TypeToBytes(uint8(9)), []uint8{9}) {
			t.Fail()
		}
	})

	t.Run("uint64", func(t *testing.T) {
		if !reflect.DeepEqual(TypeToBytes(uint64(1234567)), []uint8{135, 214, 18, 0, 0, 0, 0, 0}) {
			t.Fail()
		}
	})
}

func TestTypeToHex(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		if TypeToHex(uint8(9)) != "0x09" {
			t.Fail()
		}
	})

	t.Run("int16", func(t *testing.T) {
		if TypeToHex(int16(18)) != "0x1200" {
			t.Fail()
		}
	})

	t.Run("uint64", func(t *testing.T) {
		if TypeToHex(uint64(1234567)) != "0x87d6120000000000" {
			t.Fail()
		}
	})
}
