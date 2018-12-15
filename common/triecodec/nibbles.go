package triecodec

var (
	// HPFlag2 ...
	HPFlag2 uint8 = 2
	// HPFlag0 ...
	HPFlag0 uint8
	// NibbleTerminator ...
	NibbleTerminator uint8 = 16
	// NeedsTerminator ...
	NeedsTerminator = []uint8{HPFlag2, HPFlag2 + 1}
	// IsOddLength ...
	IsOddLength = []uint8{HPFlag0 + 1, HPFlag2 + 1}
)

// ExtractKey ...
func ExtractKey(key []uint8) []uint8 {
	return RemoveNibblesTerminator(DecodeNibbles(key))
}

// ExtractNodeKey ...
func ExtractNodeKey(node [][]uint8) []uint8 {
	if len(node) < 1 {
		return []uint8{}
	}
	key := node[0]
	return RemoveNibblesTerminator(DecodeNibbles(key))
}

// DecodeNibbles ....
func DecodeNibbles(value []uint8) []uint8 {
	nibblesWithFlag := ToNibbles(value)
	flag := nibblesWithFlag[0]

	var rawNibbles []uint8
	if Includes(IsOddLength, flag) {
		rawNibbles = nibblesWithFlag[1:]
	} else {
		rawNibbles = nibblesWithFlag[2:]
	}

	if Includes(NeedsTerminator, flag) {
		return AddNibblesTerminator(rawNibbles)
	}

	return rawNibbles
}

// EncodeNibbles ...
func EncodeNibbles(nibbles []uint8) []uint8 {
	var flag uint8

	if IsNibblesTerminated(nibbles) {
		flag = HPFlag2
	} else {
		flag = HPFlag0
	}

	rawNibbles := RemoveNibblesTerminator(nibbles)

	var prefix []uint8
	if len(rawNibbles)%2 != 0 {
		prefix = []uint8{flag + 1}
	} else {
		prefix = []uint8{flag, 0}
	}

	prefixed := make([]uint8, len(prefix)+len(rawNibbles))

	for i, x := range prefix {
		prefixed[i] = x
	}

	j := 0
	for i := len(prefix); i < len(rawNibbles); i++ {
		prefixed[i] = rawNibbles[j]
		j++
	}

	return FromNibbles(prefixed)
}

// AddNibblesTerminator ...
func AddNibblesTerminator(nibbles []uint8) []uint8 {
	if IsNibblesTerminated(nibbles) {
		return nibbles
	}

	terminated := make([]uint8, len(nibbles)+1)

	for i := 0; i < len(terminated); i++ {
		terminated[i] = nibbles[i]
	}

	terminated[len(nibbles)] = NibbleTerminator

	return terminated
}

// RemoveNibblesTerminator ...
func RemoveNibblesTerminator(nibbles []uint8) []uint8 {
	if IsNibblesTerminated(nibbles) {
		return nibbles[0 : len(nibbles)-1]
	}

	return nibbles
}

// IsNibblesTerminated ...
func IsNibblesTerminated(nibbles []uint8) bool {
	return nibbles[len(nibbles)-1] == NibbleTerminator
}

// Includes ...
func Includes(haystack []uint8, needle uint8) bool {
	for _, x := range haystack {
		if x == needle {
			return true
		}
	}

	return false
}
