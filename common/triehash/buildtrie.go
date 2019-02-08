package triehash

import (
	"math"

	"github.com/opennetsys/golkadot/common/triecodec"
	"github.com/opennetsys/golkadot/common/u8util"
)

// BuildTrie ...
// FIXME: This is problematic, the stream implementation is Substrate-only as is the branch here
func BuildTrie(input [][][]uint8, cursor int) []uint8 {
	if len(input) == 0 {
		return triecodec.CreateEmpty()
	}

	firstKey := input[0][0]
	firstValue := input[0][1]

	if len(input) == 1 {
		return triecodec.CreateLeaf(firstKey[cursor:], firstValue)
	}

	sharedNibbleCount := 0
	for index := 0; index < len(input); index++ {
		key := input[index][0]
		if index == 0 {
			sharedNibbleCount = len(key)
		}

		sharedNibbleCount = int(math.Min(float64(sharedNibbleCount), float64(triecodec.SharedPrefixLength(key, firstKey))))
	}

	if sharedNibbleCount > cursor {
		return u8util.Concat(
			triecodec.CreateExtension(firstKey[cursor:sharedNibbleCount]),
			triecodec.CreateSubstream(
				BuildTrie(input, sharedNibbleCount),
			),
		)
	}

	var value interface{}
	if len(firstKey) == cursor {
		value = firstValue
	} else {
		value = triecodec.NewNull()
	}

	var sharedNibbleCounts [16]uint8

	var start uint8
	if len(firstKey) == cursor {
		start = 1
	}

	begin := start

	for index := 0; index < 16; index++ {
		var x uint8
		for j := 0; j < len(input); j++ {
			if uint8(j) >= begin {
				k := input[j][0]
				if k[cursor] == uint8(index) {
					x++
				}
			}
		}

		sharedNibbleCounts[index] = x
		begin += sharedNibbleCounts[index]
	}

	begin = start

	var hasChildren []bool
	for _, val := range sharedNibbleCounts {
		hasChildren = append(hasChildren, val > 0)
	}

	var stream [][]uint8
	for _, count := range sharedNibbleCounts {
		var result []uint8
		if count > 0 {
			result = triecodec.CreateSubstream(
				BuildTrie(input[begin:begin+count], cursor+1),
			)

			begin += count
		}

		stream = append(stream, result)
	}

	var val []uint8
	val, ok := value.([]uint8)
	if !ok {
		val = nil
	}

	branch := triecodec.CreateBranch(
		val,
		hasChildren,
	)

	var col [][]uint8
	col = append(col, branch)
	col = append(col, stream...)
	col = append(col, triecodec.EndBranch())

	return u8util.Concat(col...)
}
