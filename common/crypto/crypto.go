package crypto

import (
	"crypto/sha256"

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

// NewBlake2b256Sig ...
func NewBlake2b256Sig(key, data []byte) ([]byte, error) {
	hash, err := blake2b.New256(key)
	if err != nil {
		return nil, err
	}

	hash.Write(data)
	return hash.Sum(nil), nil
}

// NewBlake2b512Sig ...
func NewBlake2b512Sig(key, data []byte) ([]byte, error) {
	hash, err := blake2b.New512(key)
	if err != nil {
		return nil, err
	}

	hash.Write(data)
	return hash.Sum(nil), nil
}
