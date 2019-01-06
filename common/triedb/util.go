package triedb

import (
	"fmt"
	"log"
	"math"

	"github.com/c3systems/go-substrate/common/triecodec"
	"github.com/ethereum/go-ethereum/rlp"
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
func DecodeNode(encoded Node) Node {
	fmt.Println("Debug: DecodedNode, encoded input", encoded)
	if IsNull(encoded) || Size(encoded) == 0 {
		fmt.Println("Debug: DecodedNode, is null, returning nil")
		return nil
	}

	// replaces above isSlice func
	// TODO: refactor
	encodedSlice := NewUint8FromNode(encoded)
	if len(encodedSlice) == 0 {
		fmt.Println("Debug: DecodeNode is array, returning", encoded)
		return encoded
	}

	fmt.Println("Debug: DecodeNode, encoded arg to rlp", encoded)

	var decoded []interface{}
	err := rlp.DecodeBytes(encodedSlice, &decoded)
	if err != nil {
		log.Println("Debug: DecodedNode, DecodeBytes err", err)
		return encoded
	}

	fmt.Println("Debug: DecodeNode, decoded bytes from rlp", decoded)

	var nodes []Node
	for _, s := range decoded {
		nodes = append(nodes, NewNode(s))
	}

	fmt.Println("Debug: DecodeNode, decoded bytes to Node type", nodes)

	return nodes
}

// EncodeNode ...
func EncodeNode(node Node) []uint8 {
	fmt.Println("Debug: EncodeNode, node input", node)
	var i []interface{}

	// TODO: refactor this
	slice, ok := node.([]Node)
	if ok {
		for _, s := range slice {
			if s == nil {
				// NOTE: empty string is required for nil values
				i = append(i, "")
			} else {
				slice, ok := s.([]Node)
				if ok {
					var n []Node
					for _, s := range slice {
						if s == nil {
							// NOTE: empty string is required for nil values
							n = append(n, "")
						} else {
							n = append(n, s)
						}
					}
					i = append(i, n)
				} else {
					i = append(i, s)
				}
			}
		}
	} else {
		slice, ok := node.([][][]uint8)
		if ok {
			for _, s := range slice {
				if s == nil {
					// NOTE: empty string is required for nil values
					i = append(i, "")
				} else {
					i = append(i, s)
				}
			}
		} else {
			slice, ok := node.([][]uint8)
			if ok {
				for _, s := range slice {
					if s == nil {
						// NOTE: empty string is required for nil values
						i = append(i, "")
					} else {
						i = append(i, s)
					}
				}
			} else {
				i = append(i, slice)
			}
		}
	}

	fmt.Println("Debug: EncodeNode, decoded arg to rlp", i)

	encoded, err := rlp.EncodeToBytes(&i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Debug: EncodedNode, encoded bytes from rlp", encoded)

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
