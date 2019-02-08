package triecodec

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/opennetsys/godot/common/crypto"
)

// Null ...
type Null struct {
}

// BranchHeader ...
type BranchHeader struct {
	value int // bool
}

// Value ....
func (b *BranchHeader) Value() bool {
	return b.value == 1
}

// ExtensionHeader ...
type ExtensionHeader struct {
	value int
}

// LeafHeader ...
type LeafHeader struct {
	value int
}

// NodeHeader ...
type NodeHeader struct {
	nodeType int
	value    interface{}
}

// NewNull ...
func NewNull() *Null {
	return &Null{}
}

// NewBranchHeader ...
func NewBranchHeader(value bool) *BranchHeader {
	var i int
	if value {
		i = 1
	}
	return &BranchHeader{value: i}
}

// NewExtensionHeader ...
func NewExtensionHeader(value int) *ExtensionHeader {
	return &ExtensionHeader{value}
}

// NewLeafHeader ...
func NewLeafHeader(value int) *LeafHeader {
	return &LeafHeader{value}
}

// NewNodeHeader ...
func NewNodeHeader(input interface{}) *NodeHeader {
	nodeType, value := DecodeNodeHeader(input)

	header := &NodeHeader{
		nodeType: nodeType,
		value:    value,
	}

	return header
}

// NodeType ...
func (n *NodeHeader) NodeType() int {
	return n.nodeType
}

// Value ...
func (n *NodeHeader) Value() int {
	switch v := n.value.(type) {
	case *Null:
		return 0
	case *ExtensionHeader:
		return v.value
	case *BranchHeader:
		return v.value
	case *LeafHeader:
		return v.value
	}

	return 0
}

// EncodedLength ...
func (n *NodeHeader) EncodedLength() int {
	if n.nodeType == NODE_TYPE_NULL || n.nodeType == NODE_TYPE_BRANCH {
		return 1
	} else if n.nodeType == NODE_TYPE_EXT {
		header, ok := n.value.(*ExtensionHeader)
		if !ok {
			log.Fatal(ErrTypeAssertion, 1)
		}

		nibbleCount := header.value

		if nibbleCount < EXTENSION_NODE_THRESHOLD {
			return 1
		}

		return 2
	} else if n.nodeType == NODE_TYPE_LEAF {
		header, ok := n.value.(*LeafHeader)
		if !ok {
			log.Fatal(ErrTypeAssertion, 2)
		}

		nibbleCount := header.value

		if nibbleCount < LEAF_NODE_THRESHOLD {
			return 1
		}
		return 2
	}

	log.Fatal(ErrUnreachableCode, 1)

	return 0
}

// ToUint8Slice ...
func (n *NodeHeader) ToUint8Slice() []uint8 {
	if n.nodeType == NODE_TYPE_NULL {
		return []uint8{uint8(EMPTY_TRIE)}
	} else if n.nodeType == NODE_TYPE_BRANCH {
		header, ok := n.value.(*BranchHeader)
		if !ok {
			log.Fatal(ErrTypeAssertion, 3)
		}

		if header.value == 1 {
			return []uint8{uint8(BRANCH_NODE_WITH_VALUE)}
		}
		return []uint8{uint8(BRANCH_NODE_NO_VALUE)}
	} else if n.nodeType == NODE_TYPE_EXT {
		header, ok := n.value.(*ExtensionHeader)
		if !ok {
			log.Fatal(ErrTypeAssertion, 4)
		}

		nibbleCount := header.value
		if nibbleCount < EXTENSION_NODE_THRESHOLD {
			return []uint8{uint8(EXTENSION_NODE_OFFSET + nibbleCount)}
		}

		return []uint8{uint8(EXTENSION_NODE_BIG), uint8(nibbleCount - EXTENSION_NODE_THRESHOLD)}
	} else if n.nodeType == NODE_TYPE_LEAF {
		header, ok := n.value.(*LeafHeader)
		if !ok {
			log.Fatal(ErrTypeAssertion, 10)
		}

		nibbleCount := header.value
		if nibbleCount < LEAF_NODE_THRESHOLD {
			return []uint8{uint8(LEAF_NODE_OFFSET + nibbleCount)}
		}

		return []uint8{uint8(LEAF_NODE_BIG), uint8(nibbleCount - LEAF_NODE_THRESHOLD)}
	}

	log.Fatal(ErrUnreachableCode, 2)

	return nil
}

// DecodeNodeHeader ...
func DecodeNodeHeader(input interface{}) (int, interface{}) {
	debugLog("triecodec, DecodeNodeHeader, input", input)

	switch v := input.(type) {
	case []uint8:
		debugLog("triecodec, DecodeNodeHeader, is []uint8")
		return DecodeNodeHeaderUint8Slice(v)
	case []interface{}:
		if len(v) == 1 {
			arr, ok := v[0].([]uint8)
			if ok {
				debugLog("triecodec, DecodeNodeHeader, 0 is []uint8")
				return DecodeNodeHeaderUint8Slice(arr)
			}
		}

		debugLog("triecodec, DecodeNodeHeader, slices")
		return DecodeNodeHeaderUint8Slices(v)
	}

	log.Fatal(ErrUnreachableCode, 3)

	return 0, nil
}

// DecodeNodeHeaderUint8Slice ...
func DecodeNodeHeaderUint8Slice(input []uint8) (int, interface{}) {
	firstByte := EMPTY_TRIE
	if len(input) > 0 {
		firstByte = int(input[0])
	}

	if firstByte == EMPTY_TRIE {
		return NODE_TYPE_NULL, NewNull()
	} else if firstByte == BRANCH_NODE_NO_VALUE {
		return NODE_TYPE_BRANCH, NewBranchHeader(false)
	} else if firstByte == BRANCH_NODE_WITH_VALUE {
		return NODE_TYPE_BRANCH, NewBranchHeader(true)
	} else if firstByte >= EXTENSION_NODE_OFFSET && firstByte <= EXTENSION_NODE_SMALL_MAX {
		return NODE_TYPE_EXT, NewExtensionHeader(firstByte - EXTENSION_NODE_OFFSET)
	} else if firstByte == EXTENSION_NODE_BIG {
		return NODE_TYPE_EXT, NewExtensionHeader(int(input[1] + EXTENSION_NODE_THRESHOLD))
	} else if firstByte >= LEAF_NODE_OFFSET && firstByte <= LEAF_NODE_SMALL_MAX {
		return NODE_TYPE_LEAF, NewLeafHeader(firstByte - LEAF_NODE_OFFSET)
	} else if firstByte == LEAF_NODE_BIG {
		return NODE_TYPE_LEAF, NewLeafHeader(int(input[1] + LEAF_NODE_THRESHOLD))
	}

	log.Fatal(ErrUnreachableCode, 4)

	return 0, nil
}

// DecodeNodeHeaderUint8Slices ...
func DecodeNodeHeaderUint8Slices(input []interface{}) (int, interface{}) {
	debugLog("triecodec, DecodeNodeHeaderUint8Slices")
	var isNull bool
	if len(input) == 1 {
		switch input[0].(type) {
		case *Null:
			isNull = true
		case nil:
			isNull = true
		}
	}

	debugLog("triecodec, DecodeNodeHeaderUint8Slices size", len(input))
	if len(input) == 0 || isNull {
		return NODE_TYPE_NULL, NewNull()
	} else if len(input) == 2 {
		value, ok := input[0].([]uint8)
		if !ok {
			spew.Dump(input[0])
			log.Fatal(ErrTypeAssertion, 6)
		}
		nibbles := DecodeNibbles(value)
		isTerminated := IsNibblesTerminated(nibbles)
		if isTerminated {
			return NODE_TYPE_LEAF, NewLeafHeader(len(nibbles) - 1)
		}

		return NODE_TYPE_EXT, NewExtensionHeader(len(nibbles))
	} else if len(input) == 17 {
		var value bool
		switch v := input[16].(type) {
		case nil:
			value = false
		case []uint8:
			value = v != nil
		case *crypto.Blake2b256Hash:
			value = v != nil
		case *crypto.Blake2b512Hash:
			value = v != nil
		case *crypto.Hash:
			value = v != nil
		case *Null:
			value = false
		case *BranchHeader:
			if v.value == 1 {
				value = true
			}
		default:
			value = true
		}

		debugLog("triecodec, DecodeNodeHeaderUint8Slices returning new branch header", NODE_TYPE_BRANCH, value)
		return NODE_TYPE_BRANCH, NewBranchHeader(value)
	}

	log.Fatal(ErrUnreachableCode, 5)

	return 0, nil
}
