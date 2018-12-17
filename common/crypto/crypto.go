package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/blake2b"
)

// NewSHA256 ...
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// NewBlake2b256 ...
func NewBlake2b256(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}

// NewBlake2b512 ...
func NewBlake2b512(data []byte) []byte {
	hash := blake2b.Sum512(data)
	return hash[:]
}
