package triecodec

import (
	"fmt"
	"math"

	"github.com/opennetsys/go-substrate/common/assert"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// FromNibbles converts the uint8 slice input from nibbles. Calculate and return a uint8 slice that's made from a list of nibbles.
func FromNibbles(input []uint8) []uint8 {
	result := make([]uint8, len(input)/2)

	for index := 0; index < len(result); index++ {
		nibIndex := index * 2
		result[index] = (input[nibIndex] << 4) + input[nibIndex+1]
	}

	return result
}

// ToNibbles converts the uint8 slice input to nibbles. Calculate and return a list of nibbles that makes up the input.
func ToNibbles(input []uint8) []uint8 {
	result := make([]uint8, len(input)*2)

	for index := 0; index < len(input); index++ {
		b := input[index]
		result[index*2] = b >> 4
		result[(index*2)+1] = b & 0xF
	}

	return result
}

// SharedPrefixLength returns the shared prefix length. Calculates the minimum distance that both uint8 slices share.
func SharedPrefixLength(first []uint8, second []uint8) int {
	length := int(math.Min(float64(len(first)), float64(len(second))))

	for index := 0; index < length; index++ {
		if first[index] != second[index] {
			return index
		}
	}

	return length
}

var (
	// ExtensionNodeBig ...
	ExtensionNodeBig = 253
	// ExtensionNodeOffset ...
	ExtensionNodeOffset = 128
	// LeafNodeBig ...
	LeafNodeBig = 127
	// LeafNodeOffset ...
	LeafNodeOffset = 1
	// MaxFuseLength ...
	MaxFuseLength = 255 + 126
)

// FuseNibbles ...
func FuseNibbles(nibbles []uint8, isLeaf bool) []uint8 {
	assert.Assert(len(nibbles) < MaxFuseLength, fmt.Sprintf("Input to fuseNibbles too large, found %v >= %v", len(nibbles), MaxFuseLength))

	var firstByteSmall int
	var bigThreshold int

	if isLeaf {
		firstByteSmall = LeafNodeOffset
		bigThreshold = LeafNodeBig - LeafNodeOffset
	} else {
		firstByteSmall = ExtensionNodeOffset
		bigThreshold = ExtensionNodeBig - ExtensionNodeOffset
	}

	result := [][]uint8{}

	result = append(result, []uint8{uint8(firstByteSmall + int(math.Min(float64(len(nibbles)), float64(bigThreshold))))})

	oddFlag := len(nibbles) % 2

	if len(nibbles) >= bigThreshold {
		result = append(result, []uint8{uint8(len(nibbles) - bigThreshold)})
	}

	if oddFlag == 1 {
		result = append(result, nibbles[0:1])
	}

	result = append(result, FromNibbles(nibbles[oddFlag:]))

	return u8util.Concat(result...)
}
