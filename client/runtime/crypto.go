package runtime

import (
	"github.com/opennetsys/golkadot/common/crypto"
)

// Crypto ...
type Crypto struct {
	Heap *Heap
}

// NewCrypto ...
func NewCrypto(heap *Heap) *Crypto {
	return &Crypto{
		Heap: heap,
	}
}

// Blake2b256 ...
func (c *Crypto) Blake2b256(dataPtr Pointer, dataLength int64, outPtr Pointer) {
	data := c.Heap.Get(dataPtr, dataLength)
	hash := crypto.NewBlake2b256(data)

	c.Heap.Set(outPtr, hash[:])
}

// ED25519Verify ...
func (c *Crypto) ED25519Verify(msgPtr Pointer, msgLength int64, sigPtr Pointer, pubkeyPtr Pointer) int64 {
	var pubKey [32]byte
	copy(pubKey[:], c.Heap.Get(pubkeyPtr, 32))
	if crypto.NaclVerify(
		c.Heap.Get(msgPtr, msgLength),
		c.Heap.Get(sigPtr, 64),
		pubKey,
	) {
		return 0
	}

	return 5
}

// Twox128 ...
func (c *Crypto) Twox128(dataPtr Pointer, dataLength int64, outPtr Pointer) {
	c.twox(128, dataPtr, dataLength, outPtr)
}

// Twox256 ...
func (c *Crypto) Twox256(dataPtr Pointer, dataLength int64, outPtr Pointer) {
	c.twox(256, dataPtr, dataLength, outPtr)
}

// twox ...
func (c *Crypto) twox(bitLength int64, dataPtr Pointer, dataLength int64, outPtr Pointer) {
	data := c.Heap.Get(dataPtr, dataLength)
	hash := crypto.NewXXHash(data, bitLength)

	c.Heap.Set(outPtr, hash)
}
