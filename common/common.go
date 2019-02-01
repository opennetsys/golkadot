package common

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"reflect"
	"strconv"

	"github.com/c3systems/go-substrate/common/bnutil"
	"github.com/c3systems/go-substrate/common/hexutil"
	"github.com/c3systems/go-substrate/common/u8compact"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/davecgh/go-spew/spew"
)

// ErrTypeUnsupported ...
var ErrTypeUnsupported = errors.New("type not supported")

// TypeToU8a ...
func TypeToU8a(value interface{}, isBare bool) []uint8 {
	switch v := value.(type) {
	case *big.Int:
		return bnutil.ToUint8Slice(v, 64, true, false)
	case int8:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 8, true, false)
	case int16:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 16, true, false)
	case int32:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 32, true, false)
	case int64:
		return bnutil.ToUint8Slice(big.NewInt(v), 64, true, false)
	case uint8:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 8, true, false)
	case uint16:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 16, true, false)
	case uint32:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 32, true, false)
	case uint64:
		return bnutil.ToUint8Slice(big.NewInt(int64(v)), 64, true, false)
	case []byte:
		return v[:]
	case *string:
		var encoded []byte
		if v != nil {
			encoded = []byte(*v)
		}

		if isBare {
			return encoded
		}

		return u8compact.AddLength(encoded, -1)
	case string:
		if isBare {
			return []byte(v)
		}

		return u8compact.AddLength([]byte(v), -1)
	}
	return nil
}

// TypeLen ...
func TypeLen(value interface{}) int {
	switch v := value.(type) {
	case *big.Int:
		return v.BitLen()
	case int8:
		return 8
	case int16:
		return 16
	case int32:
		return 32
	case int64:
		return 64
	case uint8:
		return 8
	case uint16:
		return 16
	case uint32:
		return 32
	case uint64:
		return 64
	case *string:
		if v != nil {
			return len(*v)
		}
	case string:
		return len(v)
	}

	return 0
}

// TypeBitLen ...
func TypeBitLen(value interface{}) int {
	switch v := value.(type) {
	case *big.Int:
		return v.BitLen()
	case int8:
		return 8
	case int16:
		return 16
	case int32:
		return 32
	case int64:
		return 64
	case uint8:
		return 8
	case uint16:
		return 16
	case uint32:
		return 32
	case uint64:
		return 64
	}

	return 0
}

// TypeEncodedLen ...
func TypeEncodedLen(value interface{}) int {
	switch v := value.(type) {
	case *big.Int:
		if v == nil {
			return 0
		}
		return v.BitLen()
	case int8:
		return 8
	case int16:
		return 16
	case int32:
		return 32
	case int64:
		return 64
	case uint8:
		return 8
	case uint16:
		return 16
	case uint32:
		return 32
	case uint64:
		return 64
	case *string:
		if v == nil {
			return 0
		}
		return TypeLen(v) + len(u8compact.CompactToUint8Slice(big.NewInt(int64(TypeLen(v))), -1))
	case string:
		return TypeLen(v) + len(u8compact.CompactToUint8Slice(big.NewInt(int64(TypeLen(v))), -1))

	}

	return 0
}

// TypeToString ...
func TypeToString(value interface{}) string {
	switch v := value.(type) {
	case *big.Int:
		return v.String()
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.Itoa(int(v))
	case uint8:
		return strconv.Itoa(int(v))
	case uint16:
		return strconv.Itoa(int(v))
	case uint32:
		return strconv.Itoa(int(v))
	case uint64:
		return strconv.Itoa(int(v))
	case *string:
		if v != nil {
			return *v
		}
		return ""
	case string:
		return v
	case []uint8:
		offset, length := u8compact.FromUint8Slice(v, 8)
		end := int(offset) + int(length.Uint64())
		if end > len(v) {
			end = len(v)
		}
		return string(v[offset:end])
	}

	return ""
}

// TypeToHex ...
func TypeToHex(value interface{}) string {
	spew.Dump()
	switch v := value.(type) {
	case *big.Int:
		return hexutil.HexFixLength(hex.EncodeToString(v.Bytes()), TypeBitLen(v), true)
	case int8:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 8, true)
	case int16:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 16, true)
	case int32:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 32, true)
	case int64:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 64, true)
	case uint8:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 8, true)
	case uint16:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 16, true)
	case uint32:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 32, true)
	case uint64:
		return hexutil.HexFixLength(hex.EncodeToString(TypeToBytes(v)), 61, true)
	case []uint8:
		return hexutil.HexFixLength(hex.EncodeToString(v), len(v), true)
	case *string:
		if hexutil.ValidHex(*v) {
			return u8util.ToHex(TypeToU8a(*v, false), -1, true)
		}

		return u8util.ToHex(TypeToU8a(*v, false), -1, true)
	case string:
		if hexutil.ValidHex(v) {
			return u8util.ToHex(TypeToU8a(v, false), -1, true)
		}

		return u8util.ToHex(TypeToU8a(v, false), -1, true)
	default:
		log.Fatal(ErrTypeUnsupported)
	}

	return ""
}

// TypeEquals ...
func TypeEquals(value interface{}, other interface{}) bool {
	return TypeToString(value) == TypeToString(other)
}

// TypeIsZeroValue ...
func TypeIsZeroValue(value interface{}) bool {
	switch v := value.(type) {
	case *big.Int:
		return v.Cmp(big.NewInt(0)) == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint8:
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	default:
		log.Fatal(ErrTypeUnsupported)
	}

	return false
}

// TypeToBytes ...
func TypeToBytes(value interface{}) []byte {
	switch v := value.(type) {
	case *big.Int:
		return v.Bytes()
	case int8:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:1]
	case int16:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:2]
	case int32:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:4]
	case int64:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:8]
	case uint8:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:1]
	case uint16:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:2]
	case uint32:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:4]
	case uint64:
		bs := make([]byte, 64)
		binary.LittleEndian.PutUint64(bs, uint64(v))
		return bs[:8]
	default:
		log.Fatal(ErrTypeUnsupported)
	}

	return nil
}

// TypeIsNil ...
func TypeIsNil(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}
