package triedb

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"

	"github.com/opennetsys/golkadot/common/triecodec"
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
func IsKvNode(node interface{}) bool {
	val := reflect.ValueOf(node)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len() == 2
	}

	return false
}

// IsBranchNode ...
func IsBranchNode(node interface{}) bool {
	val := reflect.ValueOf(node)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len() == 17
	}

	return false
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
func IsEmptyNode(node interface{}) bool {
	if node == nil {
		return true
	}

	return false
}

// IsNull ...
func IsNull(node interface{}) bool {
	if node == nil {
		return true
	}

	return false
}

// Size ...
func Size(value interface{}) int {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len()
	}

	return 0
}

// DecodeNode ...
func DecodeNode(encoded Node, codec InterfaceCodec) Node {
	debugLog("DecodedNode, encoded input", encoded)
	if IsNull(encoded) || Size(encoded) == 0 {
		debugLog("DecodedNode, is null, returning nil")
		return nil
	}

	// replaces above isSlice func
	// TODO: refactor
	if IsMultiSlice(encoded) {
		debugLog("DecodeNode is 'array', returning", encoded)
		return encoded
	}

	encodedSlice := NewUint8FromNode(encoded)

	debugLog("DecodeNode: encodedslice", encodedSlice)

	debugLog("DecodeNode is not 'array'")
	debugLog("DecodeNode, encoded arg to codec decoder", encodedSlice)

	var decoded []interface{}
	if err := codec.Decode(encodedSlice, &decoded); err != nil {
		debugLog("DecodedNode, DecodeBytes err", err)
		return encoded
	}

	debugLog("DecodeNode, decoded bytes from codec decoder", decoded)

	var nodes []Node
	for _, s := range decoded {
		nodes = append(nodes, NewNode(s))
	}

	debugLog("DecodeNode, decoded bytes to Node type", nodes)

	return nodes
}

// EncodeNode ...
func EncodeNode(node Node, codec InterfaceCodec) []uint8 {
	debugLog("EncodeNode, node input", node)

	encoded, err := codec.Encode(node)
	if err != nil {
		log.Fatal(err)
	}

	debugLog("EncodeNode, encoded bytes from codec encoder", encoded)

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
	case []interface{}:
		return len(v) > 1
	}

	return false
}

var debugEnabled bool

func debugLog(args ...interface{}) {
	if debugEnabled {
		args = append([]interface{}{"Debug: "}, args...)
		fmt.Println(args...)
	}
}

func init() {
	debugEnabled = os.Getenv("DEBUG") != ""
}
