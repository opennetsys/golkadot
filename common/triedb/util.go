package triedb

import (
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/triecodec"
)

// GetNodeType ...
func GetNodeType(node Node) int {
	if IsEmptyNode(node) {
		return NodeTypeEmpty
	} else if IsKvNode(node) {
		nodekv, ok := node.([][]uint8)
		if !ok {
			log.Fatal("not ok")
		}
		key := nodekv[0]
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
	case [][]uint8:
		if !IsEmptyNode(node) {
			if len(v) == 2 {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// IsBranchNode ...
func IsBranchNode(node Node) bool {
	switch v := node.(type) {
	case NodeBranch:
		if len(v) == 17 {
			return true
		}
		return false
	case [][]uint8:
		if !IsEmptyNode(node) {
			if len(v) == 17 {
				return true
			}
		}
		return false
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
	case NodeEmpty:
		return true
	case [][]uint8:
		if v == nil {
			return true
		}
		return false
	case []uint8:
		if v == nil {
			return true
		}
		return false
	default:
		return false
	}
}

// IsNull ...
func IsNull(node Node) bool {
	switch v := node.(type) {
	case NodeEmpty:
		return true
	case nil:
		return true
	case []uint8:
		if v == nil {
			return true
		}
		return false
	default:
		return false
	}
}

// IsSlice ...
func IsSlice(value interface{}) bool {
	switch value.(type) {
	case []uint8:
		return true
	case [][]uint8:
		return true
	default:
		return false
	}
}

// Size ...
func Size(value interface{}) int {
	switch v := value.(type) {
	case []uint8:
		return len(v)
	case []Node:
		return len(v)
	case []interface{}:
		return len(v)
	default:
		return 0
	}
}

// DecodeNode ...
func DecodeNode(encoded Node) Node {
	if IsNull(encoded) || Size(encoded) == 0 {
		return nil
	} else if IsSlice(encoded) {
		return encoded
	}

	encodedSlice, ok := encoded.([]uint8)
	if !ok {
		log.Fatal("not ok")
	}
	node := triecodec.Decode(encodedSlice)

	decodedSlice, ok := node.([][]uint8)
	if !ok {
		log.Fatal("not ok")
	}

	var nodes []Node
	for _, value := range decodedSlice {
		if value != nil && len(value) > 0 {
			nodes = append(nodes, NewNode(value))
		} else {
			nodes = append(nodes, NewNodeEmpty())
		}
	}

	return nodes
}

// EncodeNode ...
func EncodeNode(node Node) []uint8 {
	var i []interface{}
	i = append(i, node)
	return triecodec.Encode(i)
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
