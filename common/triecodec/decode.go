package triecodec

import (
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/u8compact"
)

// Decode ...
func Decode(input []uint8) interface{} {
	header := NewNodeHeader(input)
	nodeType := header.NodeType()

	var isNull bool // todo

	if nodeType == NODE_TYPE_NULL || isNull {
		return input
	} else if nodeType == NODE_TYPE_BRANCH {
		return decodeBranch(header, input)
	} else if nodeType == NODE_TYPE_EXT || nodeType == NODE_TYPE_LEAF {
		return decodeKv(header, input)
	}

	log.Fatal(ErrUnreachableCode)

	return nil
}

func decodeBranch(header *NodeHeader, input []uint8) []interface{} {
	offset := header.EncodedLength()
	branch, ok := header.value.(*BranchHeader)
	if !ok {
		log.Fatal(ErrTypeAssertion)
	}
	bitmap := int(input[offset]) + (int(input[offset+1]) * 256)
	var value []uint8

	offset += 2

	if branch.Value() == true {
		length, bytes := u8compact.StripLength(input[offset:], 32)
		value = bytes
		offset += length
	}

	cursor := 1

	emptyBranch := make([]interface{}, 16)
	emptyBranch = append(emptyBranch, value)

	for index := 0; index < len(emptyBranch); index++ {
		value := emptyBranch[index]
		var result interface{}
		result = value

		if (index < 16) && (bitmap&cursor) != 0 {
			length, bytes := u8compact.StripLength(input[offset:], 32)

			if len(bytes) == 32 {
				result = bytes
			} else {
				result = Decode(bytes)
			}

			offset += length
		}

		cursor = cursor << 1

		emptyBranch[index] = result
	}

	return emptyBranch
}

func decodeKv(header *NodeHeader, input []uint8) []interface{} {
	offset := header.EncodedLength()
	nibbleCount := header.Value()
	nibbleLength := int(math.Floor(float64(nibbleCount+1) / float64(2)))
	nibbleData := input[offset : offset+nibbleLength]

	// for odd, ignore the first nibble, data starts at offset 1
	nibbles := ToNibbles(nibbleData)[(nibbleCount % 2):]

	offset += len(nibbleData)
	_, value := u8compact.StripLength(input[offset:], 32)
	if header.NodeType() == NODE_TYPE_LEAF {
		return []interface{}{
			EncodeNibbles(AddNibblesTerminator(nibbles)),
			value,
		}
	}

	return []interface{}{
		EncodeNibbles(nibbles),
		value,
	}
}
