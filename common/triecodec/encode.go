package triecodec

import (
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/u8compact"
	"github.com/c3systems/go-substrate/common/u8util"
)

// Type ...
const Type = "Substrate"

// BranchValueIndex ...
const BranchValueIndex = 16

// Encode ...
func Encode(input []interface{}) []uint8 {
	if len(input) == 1 {
		v, ok := input[0].([]uint8)
		if ok {
			return v
		}
	}

	header := NewNodeHeader(input)
	nodeType := header.NodeType()

	if nodeType == NODE_TYPE_NULL {
		return header.ToUint8Slice()
	} else if nodeType == NODE_TYPE_BRANCH {
		return encodeBranch(header, input)
	} else if nodeType == NODE_TYPE_EXT || nodeType == NODE_TYPE_LEAF {
		return encodeKv(header, input)
	}

	log.Fatal(ErrUnreachableCode)

	return nil
}

func encodeBranch(header *NodeHeader, input []interface{}) []uint8 {
	valuesU8a := []uint8{}
	bitmap := 0

	cursor := 1
	for index := 0; index < len(input); index++ {
		value := input[index]
		_, ok := value.(*Null)
		if (index < BranchValueIndex) && !ok {
			bitmap = bitmap | cursor
			valuesU8a = u8util.Concat(
				valuesU8a,
				encodeValue(value),
			)
		}

		cursor = cursor << 1
	}

	return u8util.Concat(
		header.ToUint8Slice(),
		[]uint8{uint8(bitmap % 256), uint8(math.Floor(float64(bitmap) / float64(256)))},
		encodeValue(input[BranchValueIndex]),
		valuesU8a,
	)
}

func encodeKv(header *NodeHeader, input []interface{}) []uint8 {
	key, ok := input[0].([]uint8)
	if !ok {
		log.Fatal(ErrCastingType)
	}

	value := input[1]
	return u8util.Concat(
		header.ToUint8Slice(),
		encodeKey(key),
		encodeValue(value),
	)
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
	_, ok := input.(*Null)
	if ok {
		isNull = true
	}

	if input == nil || isNull {
		return []uint8{}
	}

	var i []interface{}
	v, ok := input.([][]uint8)
	if ok {
		for _, arr := range v {
			i = append(i, arr)
		}
	} else {
		i = append(i, input)
	}

	encoded := Encode(i)
	return u8compact.AddLength(encoded, 32)
}
