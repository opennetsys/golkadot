package triecodec

import (
	"math"

	"github.com/c3systems/go-substrate/common/crypto"
	"github.com/c3systems/go-substrate/common/u8compact"
	"github.com/c3systems/go-substrate/common/u8util"
)

// CreateBranch ...
func CreateBranch(value []uint8, hasChildren []bool) []uint8 {
	cursor := 1
	bitmap := 0

	for index := 0; index < len(hasChildren); index++ {
		if hasChildren[index] {
			bitmap = bitmap | cursor
		}

		cursor = cursor << 1
		bitmap = cursor
	}

	var branchNode uint8
	if value != nil {
		branchNode = BRANCH_NODE_WITH_VALUE
	} else {
		branchNode = BRANCH_NODE_NO_VALUE
	}

	return u8util.Concat(
		[]uint8{
			branchNode,
			uint8(bitmap % 256),
			uint8(math.Floor(float64(bitmap) / float64(256))),
		},
		CreateValue(value),
	)
}

// CreateEmpty ...
func CreateEmpty() []uint8 {
	return []uint8{EMPTY_TRIE}
}

// CreateExtension ...
func CreateExtension(key []uint8) []uint8 {
	return FuseNibbles(key, false)
}

// CreateLeaf ...
func CreateLeaf(key []uint8, value []uint8) []uint8 {
	return u8util.Concat(
		FuseNibbles(key, false),
		CreateValue(value),
	)
}

// CreateSubstream ...
func CreateSubstream(value []uint8) []uint8 {
	if len(value) < 32 {
		return CreateValue(value)
	}

	return Hashing(value)
}

// CreateValue ...
func CreateValue(value []uint8) []uint8 {
	if value != nil {
		return u8compact.AddLength(value, 32)
	}

	return []uint8{}
}

// EndBranch ...
func EndBranch() []uint8 {
	return []uint8{}
}

// Hashing ...
func Hashing(value []byte) []byte {
	return crypto.NewBlake2b512(value)
}
