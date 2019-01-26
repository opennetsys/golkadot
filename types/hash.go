package types

// Hash ...
// TODO: impl u8fixed methods
type Hash [256]uint8

// NewHash ...
func NewHash(input []byte) *Hash {
	h := new(Hash)
	copy(h[:], input)
	return h
}

// Value ...
func (h Hash) Value() [256]uint8 {
	return h
}
