package triecodec

import (
	"log"
	"math"

	"github.com/opennetsys/go-substrate/common/crypto"
	"github.com/opennetsys/go-substrate/common/u8compact"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// Type ...
const Type = "Substrate"

// BranchValueIndex ...
const BranchValueIndex = 16

// Encode ...
func Encode(input []interface{}) []uint8 {
	debugLog("triecodec; Encode, input", input)
	if len(input) == 1 {
		v, ok := input[0].([]uint8)
		if ok {
			debugLog("triecodec; Encode, returning already encoded input", v)
			return v
		}
	}

	header := NewNodeHeader(input)
	nodeType := header.NodeType()

	debugLog("triecodec; Encode, new node header value", header.Value())
	debugLog("triecodec; Encode, new node header type", nodeType)

	if nodeType == NODE_TYPE_NULL {
		debugLog("triecodec; Encode, is null")
		result := header.ToUint8Slice()
		debugLog("triecodec; Encode, is null, result", result)
		return result
	} else if nodeType == NODE_TYPE_BRANCH {
		debugLog("triecodec; Encode, is branch")
		result := encodeBranch(header, input)
		debugLog("triecodec; Encode, is branch result", result)
		return result
	} else if nodeType == NODE_TYPE_EXT || nodeType == NODE_TYPE_LEAF {
		debugLog("triecodec; Encode, is kv")
		result := encodeKv(header, input)
		debugLog("triecodec; Encode, is kv, result", result)
		return result
	}

	log.Fatal(ErrUnreachableCode)

	return nil
}

func encodeBranch(header *NodeHeader, input []interface{}) []uint8 {
	debugLog("triecodec; encodeBranch, header value", header.Value())
	debugLog("triecodec; encodeBranch, header type", header.NodeType())
	debugLog("triecodec; encodeBranch, input", input)

	valuesU8a := []uint8{}
	var bitmap int64
	var cursor int64 = 1

	for index := 0; index < len(input); index++ {

		value := input[index]
		var isNull bool
		switch v := value.(type) {
		case *Null:
			isNull = true
		case nil:
			isNull = true
		case []uint8:
			isNull = v == nil
		}
		isLt := (index < BranchValueIndex)
		debugLog("triecodec; encodeBranch, reduce, isLt? isNull?", isLt, isNull)
		if isLt && !isNull {
			bitmap = bitmap | cursor
			ev := encodeValue(value)

			debugLog("triecodec; encodeBranch, reduce, bitmap", bitmap)
			debugLog("triecodec; encodeBranch, reduce, valuesU8a", valuesU8a)
			debugLog("triecodec; encodeBranch, reduce, encoded value", ev)

			valuesU8a = u8util.Concat(valuesU8a, ev)
		}

		cursor = cursor << 1
	}

	debugLog("triecodec; encodeBranch, end cursor", cursor)
	debugLog("triecodec; encodeBranch, end bitmap", bitmap)

	h := header.ToUint8Slice()
	hk := []uint8{uint8(bitmap % 256), uint8(math.Floor(float64(bitmap) / float64(256)))}
	ev := encodeValue(input[BranchValueIndex])

	debugLog("triecodec; encodeBranch, header uint8 slice", h)
	debugLog("triecodec; encodeBranch, header-k uint8 slice", hk)
	debugLog("triecodec; encodeBranch, encoded value", ev)
	debugLog("triecodec; encodeBranch, values uint8 slice", valuesU8a)

	return u8util.Concat(
		h,
		hk,
		ev,
		valuesU8a,
	)
}

func encodeKv(header *NodeHeader, input []interface{}) []uint8 {
	key, ok := input[0].([]uint8)
	if !ok {
		log.Fatal(ErrTypeAssertion)
	}

	value := input[1]

	h := header.ToUint8Slice()
	k := encodeKey(key)
	v := encodeValue(value)

	debugLog("triecodec; encodeKv, header", h)
	debugLog("triecodec; encodeKv, encoded key", k)
	debugLog("triecodec; encodeKv, encoded value", v)

	result := u8util.Concat(h, k, v)
	debugLog("triecodec; encodeKv, concatenated result", result)

	return result
}

// in the case of odd nibbles, the first byte is encoded as a single
// byte from the nibble, with the remainder of the nibbles is converted
// as nomral nibble combined bytes
func encodeKey(input []uint8) []uint8 {
	nibbles := ExtractKey(input)

	if len(nibbles)%2 == 1 {
		return u8util.Concat(
			[]uint8{uint8(nibbles[0])},
			FromNibbles(nibbles[1:]),
		)
	}

	return FromNibbles(nibbles)
}

func encodeValue(input interface{}) []uint8 {
	var isNull bool
	switch v := input.(type) {
	case nil:
		isNull = true
	case *Null:
		isNull = true
	case []uint8:
		isNull = v == nil
	case *crypto.Blake2b256Hash:
		isNull = v == nil
	case *crypto.Blake2b512Hash:
		isNull = v == nil
	case *crypto.Hash:
		isNull = v == nil
	case []interface{}:
		isNull = v == nil
	case interface{}:
		isNull = v == nil
	default:
		debugLog("triecodec; encodeValue, default:", input)
		debugDump(v)
	}

	debugLog("triecodec; encodeValue, input", input)
	debugLog("triecodec; encodeValue, isNull", isNull)

	if isNull {
		return []uint8{}
	}

	var i []interface{}
	switch v := input.(type) {
	case [][]uint8:
		for _, arr := range v {
			i = append(i, arr)
		}
	case []interface{}:
		for _, arr := range v {
			i = append(i, arr)
		}
	default:
		i = append(i, input)
	}

	debugLog("triecodec; encodeValue, call Encode", i)

	encoded := Encode(i)
	return u8compact.AddLength(encoded, 32)
}
