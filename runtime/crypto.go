package runtime

import (
	"github.com/c3systems/go-substrate/common/crypto"
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

/*
oexport default function crypto ({ l, heap }: RuntimeEnv): RuntimeInterface$Crypto {
  const twox = (bitLength: number, dataPtr: Pointer, dataLen: number, outPtr: Pointer): void => {
    const data = heap.get(dataPtr, dataLen);
    const hash = xxhashAsU8a(data, bitLength);

    l.debug(() => [`twox_${bitLength}`, [dataPtr, dataLen, outPtr], '<-', u8aDisplay(data), '->', u8aToHex(hash)]);

    heap.set(outPtr, hash);
  };

  return {
    blake2_256: (dataPtr: Pointer, dataLen: number, outPtr: Pointer): void =>
      instrument('blake2_256', (): void => {
        const data = heap.get(dataPtr, dataLen);
        const hash = blake2AsU8a(data, 256);

        l.debug(() => ['blake2_256', [dataPtr, dataLen, outPtr], '<-', u8aToHex(data), '->', u8aToHex(hash)]);

        heap.set(outPtr, hash);
      }),
    ed25519_verify: (msgPtr: Pointer, msgLen: number, sigPtr: Pointer, pubkeyPtr: Pointer): number =>
      instrument('ed25519_verify', (): number => {
        l.debug(() => ['ed25519_verify', [msgPtr, msgLen, sigPtr, pubkeyPtr]]);

        return naclVerify(
          heap.get(msgPtr, msgLen),
          heap.get(sigPtr, 64),
          heap.get(pubkeyPtr, 32)
        ) ? 0 : 5;
      }),
    twox_128: (dataPtr: Pointer, dataLen: number, outPtr: Pointer): void =>
      instrument('twox_128', (): void =>
        twox(128, dataPtr, dataLen, outPtr)
      ),
    twox_256: (dataPtr: Pointer, dataLen: number, outPtr: Pointer): void =>
      instrument('twox_128', (): void =>
        twox(256, dataPtr, dataLen, outPtr)
      )
  };
}
*/
