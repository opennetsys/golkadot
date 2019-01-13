package crypto

import (
	"crypto/sha256"
	"math"

	"github.com/pierrec/xxHash/xxHash64"
	"golang.org/x/crypto/blake2b"
)

// Hash ...
type Hash [sha256.Size]uint8

// Blake2b256Hash ...
type Blake2b256Hash [blake2b.Size256]uint8

// Blake2b512Hash ...
type Blake2b512Hash [blake2b.Size]uint8

// NewSHA256 ...
func NewSHA256(data []byte) *Hash {
	var hash Hash
	hash = sha256.Sum256(data)
	return &hash
}

// NewBlake2b256 ...
func NewBlake2b256(data []byte) *Blake2b256Hash {
	var hash Blake2b256Hash
	hash = blake2b.Sum256(data)
	return &hash
}

// NewBlake2b512 ...
func NewBlake2b512(data []byte) *Blake2b512Hash {
	var hash Blake2b512Hash
	hash = blake2b.Sum512(data)
	return &hash
}

// NewXXHash64 ...
func NewXXHash64(data []byte) [8]byte {
	var hash [8]byte
	copy(hash[:], newXXHash(data, 64))
	return hash
}

// NewXXHash128 ...
func NewXXHash128(data []byte) [16]byte {
	var hash [16]byte
	copy(hash[:], newXXHash(data, 128))
	return hash
}

// NewXXHash256 ...
func NewXXHash256(data []byte) [32]byte {
	var hash [32]byte
	copy(hash[:], newXXHash(data, 256))
	return hash
}

func newXXHash(data []byte, bitLength uint) []byte {
	byteLength := int64(math.Ceil(float64(bitLength) / float64(8)))
	iterations := int64(math.Ceil(float64(bitLength) / float64(64)))
	var hash = make([]byte, byteLength)

	for seed := int64(0); seed < iterations; seed++ {
		digest := xxHash64.New(uint64(seed))
		digest.Write(data)
		copy(hash[seed*8:], digest.Sum(nil))
	}

	return hash
}
