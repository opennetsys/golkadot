package triecodec

// FromNibbles ...
func FromNibbles(input []uint8) []uint8 {
	result := make([]uint8, len(input)/2)

	for index := 0; index < len(result); index++ {
		nibIndex := index * 2

		x := (input[nibIndex] << 4) + input[nibIndex+1]

		result[index] = x
	}

	return result
}
