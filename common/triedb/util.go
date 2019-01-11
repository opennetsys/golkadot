package triedb

import (
	"fmt"
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/davecgh/go-spew/spew"
)

// GetNodeType ...
func GetNodeType(node Node) int {
	if IsEmptyNode(node) {
		return NodeTypeEmpty
	} else if IsKvNode(node) {
		key := NewFirstUint8ListFromNode(node)
		nibbles := triecodec.DecodeNibbles(key)

		if triecodec.IsNibblesTerminated(nibbles) {
			return NodeTypeLeaf
		}

		return NodeTypeExtension
	} else if IsBranchNode(node) {
		return NodeTypeBranch
	}

	log.Fatal("Unable to determine node type")
	return -1
}

// IsKvNode ...
func IsKvNode(node Node) bool {
	switch v := node.(type) {
	case []Node:
		if IsEmptyNode(node) {
			return false
		}
		return len(v) == 2
	case [][]uint8:
		if IsEmptyNode(node) {
			return false
		}
		return len(v) == 2
	case []interface{}:
		return len(v) == 2
	default:
		return false
	}
}

// IsBranchNode ...
func IsBranchNode(node Node) bool {
	switch v := node.(type) {
	case [][]uint8:
		if IsEmptyNode(v) {
			return false
		}
		return len(v) == 17
	case []Node:
		if IsEmptyNode(v) {
			return false
		}
		return len(v) == 17
	case []interface{}:
		if IsEmptyNode(v) {
			return false
		}
		return len(v) == 17
	default:
		return false
	}
}

// IsExtensionNode ...
func IsExtensionNode(node Node) bool {
	return GetNodeType(node) == NodeTypeExtension
}

// IsLeafNode ...
func IsLeafNode(node Node) bool {
	return GetNodeType(node) == NodeTypeLeaf
}

// IsEmptyNode ...
func IsEmptyNode(node Node) bool {
	switch v := node.(type) {
	case nil:
		return true
	case Node:
		switch u := v.(type) {
		case []Node:
			return u == nil || len(u) == 0
		}
		return v == nil
	case [][]uint8:
		return v == nil
	case []uint8:
		return v == nil
	case []Node:
		return v == nil
	default:
		return false
	}
}

// IsNull ...
func IsNull(node Node) bool {
	switch v := node.(type) {
	case nil:
		return true
	case []uint8:
		return v == nil
	case [][]uint8:
		return v == nil
	case []Node:
		return v == nil
	default:
		return false
	}
}

// Size ...
func Size(value interface{}) int {
	switch v := value.(type) {
	case nil:
		return 0
	case []uint8:
		return len(v)
	case []Node:
		return len(v)
	case Node:
		return 1
	case []interface{}:
		return len(v)
	default:
		return 0
	}
}

// DecodeNode ...
func DecodeNode(encoded Node, codec InterfaceCodec) Node {
	fmt.Println("Debug: DecodedNode, encoded input", encoded)
	if IsNull(encoded) || Size(encoded) == 0 {
		fmt.Println("Debug: DecodedNode, is null, returning nil")
		return nil
	}

	// replaces above isSlice func
	// TODO: refactor
	if IsMultiSlice(encoded) {
		fmt.Println("Debug: DecodeNode is 'array', returning", encoded)
		return encoded
	}

	encodedSlice := NewUint8FromNode(encoded)

	fmt.Println(encodedSlice)
	spew.Dump(encodedSlice)

	fmt.Println("Debug: DecodeNode is not 'array'")
	fmt.Println("Debug: DecodeNode, encoded arg to codec decoder", encodedSlice)

	var decoded []interface{}
	if err := codec.Decode(encodedSlice, &decoded); err != nil {
		log.Println("Debug: DecodedNode, DecodeBytes err", err)
		return encoded
	}

	fmt.Println("Debug: DecodeNode, decoded bytes from codec decoder", decoded)

	var nodes []Node
	for _, s := range decoded {
		nodes = append(nodes, NewNode(s))
	}

	fmt.Println("Debug: DecodeNode, decoded bytes to Node type", nodes)

	return nodes
}

// EncodeNode ...
func EncodeNode(node Node, codec InterfaceCodec) []uint8 {
	fmt.Println("Debug: EncodeNode, node input", node)

	encoded, err := codec.Encode(node)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Debug: EncodeNode, encoded bytes from codec encoder", encoded)

	return encoded
}

// KeyEquals ...
func KeyEquals(key []uint8, test []uint8) bool {
	if IsNull(key) && IsNull(test) {
		return true
	} else if IsNull(key) || IsNull(test) || len(key) != len(test) {
		return false
	}

	return KeyStartsWith(key, test)
}

// KeyStartsWith ...
func KeyStartsWith(key []uint8, partial []uint8) bool {
	if IsNull(key) && IsNull(partial) {
		return true
	} else if IsNull(key) || IsNull(partial) || len(key) < len(partial) {
		return false
	}

	for index := 0; index < len(partial); index++ {
		if key[index] != partial[index] {
			return false
		}
	}

	return true
}

// ComputeExtensionKey ...
func ComputeExtensionKey(nibbles []uint8) EncodedPath {
	return NewEncodedPath(triecodec.EncodeNibbles(nibbles))
}

// ComputeLeafKey ...
func ComputeLeafKey(nibbles []uint8) EncodedPath {
	return NewEncodedPath(
		triecodec.EncodeNibbles(
			triecodec.AddNibblesTerminator(nibbles),
		),
	)
}

// GetCommonPrefixLength ...
func GetCommonPrefixLength(left []uint8, right []uint8) int {
	for index := 0; index < len(left) && index < len(right); index++ {
		if left[index] != right[index] {
			return index
		}
	}

	return int(math.Min(float64(len(left)), float64(len(right))))
}

// ConsumeCommonPrefix ...
func ConsumeCommonPrefix(left []uint8, right []uint8) [][]uint8 {
	length := GetCommonPrefixLength(left, right)

	return [][]uint8{
		left[0:length],
		left[length:],
		right[length:],
	}
}

// IsMultiSlice ...
func IsMultiSlice(value interface{}) bool {
	switch v := value.(type) {
	case []Node:
		return len(v) > 1
	case [][]uint8:
		return len(v) > 1
	case Node:
		switch u := v.(type) {
		case []uint8:
			return false
		case [][]uint8:
			return len(u) > 1
		case []interface{}:
			return len(u) > 1
		default:
			return false
		}
	case []interface{}:
		return len(v) > 1
	default:
		return false
	}
}
