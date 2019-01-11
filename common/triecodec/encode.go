package triecodec

import (
	"fmt"
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/u8compact"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/davecgh/go-spew/spew"
)

// Type ...
const Type = "Substrate"

// BranchValueIndex ...
const BranchValueIndex = 16

// Encode ...
func Encode(input []interface{}) []uint8 {
	fmt.Println("Debug: triecodec; Encode, input", input)
	if len(input) == 1 {
		v, ok := input[0].([]uint8)
		if ok {
			fmt.Println("Debug: triecodec; Encode, returning already encoded input", v)
			return v
		}
	}

	header := NewNodeHeader(input)
	nodeType := header.NodeType()

	fmt.Println("Debug: triecodec; Encode, new node header value", header.Value())
	fmt.Println("Debug: triecodec; Encode, new node header type", nodeType)

	if nodeType == NODE_TYPE_NULL {
		fmt.Println("Debug: triecodec; Encode, is null")
		result := header.ToUint8Slice()
		fmt.Println("Debug: triecodec; Encode, is null, result", result)
		return result
	} else if nodeType == NODE_TYPE_BRANCH {
		fmt.Println("Debug: triecodec; Encode, is branch")
		result := encodeBranch(header, input)
		fmt.Println("Debug: triecodec; Encode, is branch result", result)
		return result
	} else if nodeType == NODE_TYPE_EXT || nodeType == NODE_TYPE_LEAF {
		fmt.Println("Debug: triecodec; Encode, is kv")
		result := encodeKv(header, input)
		fmt.Println("Debug: triecodec; Encode, is kv, result", result)
		return result
	}

	log.Fatal(ErrUnreachableCode)

	return nil
}

func encodeBranch(header *NodeHeader, input []interface{}) []uint8 {
	fmt.Println("Debug: triecodec; encodeBranch, header value", header.Value())
	fmt.Println("Debug: triecodec; encodeBranch, header type", header.NodeType())
	fmt.Println("Debug: triecodec; encodeBranch, input", input)

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
		fmt.Println("Debug: triecodec; encodeBranch, reduce, isLt? isNull?", isLt, isNull)
		if isLt && !isNull {
			bitmap = bitmap | cursor
			ev := encodeValue(value)

			fmt.Println("Debug: triecodec; encodeBranch, reduce, bitmap", bitmap)
			fmt.Println("Debug: triecodec; encodeBranch, reduce, valuesU8a", valuesU8a)
			fmt.Println("Debug: triecodec; encodeBranch, reduce, encoded value", ev)

			valuesU8a = u8util.Concat(valuesU8a, ev)
		}

		cursor = cursor << 1
	}

	fmt.Println("Debug: triecodec; encodeBranch, end cursor", cursor)
	fmt.Println("Debug: triecodec; encodeBranch, end bitmap", bitmap)

	h := header.ToUint8Slice()
	hk := []uint8{uint8(bitmap % 256), uint8(math.Floor(float64(bitmap) / float64(256)))}
	ev := encodeValue(input[BranchValueIndex])

	fmt.Println("Debug: triecodec; encodeBranch, header uint8 slice", h)
	fmt.Println("Debug: triecodec; encodeBranch, header-k uint8 slice", hk)
	fmt.Println("Debug: triecodec; encodeBranch, encoded value", ev)
	fmt.Println("Debug: triecodec; encodeBranch, values uint8 slice", valuesU8a)

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

	fmt.Println("Debug: triecodec; encodeKv, header", h)
	fmt.Println("Debug: triecodec; encodeKv, encoded key", k)
	fmt.Println("Debug: triecodec; encodeKv, encoded value", v)

	result := u8util.Concat(h, k, v)
	fmt.Println("Debug: triecodec; encodeKv, concatenated result", result)

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
	case []interface{}:
		isNull = v == nil
	case interface{}:
		isNull = v == nil
	default:
		spew.Dump(v)
	}

	fmt.Println("Debug: triecodec; encodeValue, input", input)
	fmt.Println("Debug: triecodec; encodeValue, isNull", isNull)

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

	fmt.Println("Debug: triecodec; encodeValue, call Encode", i)

	encoded := Encode(i)
	return u8compact.AddLength(encoded, 32)
}
