package address

import (
	"errors"

	"github.com/opennetsys/go-substrate/common/crypto"
	"github.com/opennetsys/go-substrate/common/hexutil"
	"github.com/opennetsys/go-substrate/common/u8util"

	"github.com/mr-tron/base58/base58"
)

// Encode ...
func Encode(b []byte, prefix PrefixEnum) (string, error) {
	var (
		allowed, isPublicKey bool
	)

	l := len(b)
	for idx := range DefaultAllowedDecodedLengths {
		if l == DefaultAllowedDecodedLengths[idx] {
			allowed = true
			break
		}
	}
	if !allowed {
		return "", ErrDecodedLengthNotAllowed
	}

	if l == 32 {
		isPublicKey = true
	}

	if prefix == nil {
		prefix = DefaultPrefix
	}

	input := u8util.Concat([]uint8{uint8(prefix.Type())}, b)
	hash := crypto.NewBlake2b512(input)
	if hash == nil {
		return "", errors.New("nil blake hash")
	}

	ending := 1
	if isPublicKey {
		ending = 2
	}
	if len(hash) < ending {
		return "", errors.New("invalid hash length")
	}

	return base58.Encode(append(input, hash[0:ending]...)), nil
}

// Decode ...
func Decode(encoded string, prefix PrefixEnum) ([]byte, error) {
	var (
		allowed, isPublicKey bool
	)

	if hexutil.ValidHex(encoded) {
		return hexutil.ToUint8Slice(encoded, -1)
	}

	decoded, err := base58.Decode(encoded)
	if err != nil {
		return nil, err
	}

	l := len(decoded)
	for idx := range DefaultAllowedEncodedLengths {
		if l == DefaultAllowedEncodedLengths[idx] {
			allowed = true
			break
		}
	}
	if !allowed {
		return nil, ErrDecodedLengthNotAllowed
	}

	// TODO Unless it is an "use everywhere" prefix, throw an error
	// if (decoded[0] !== prefix) {
	//   console.log(`WARN: Expected ${prefix}, found ${decoded[0]}`);
	// }

	if l == 35 {
		isPublicKey = true
	}

	// non-publicKeys has 1 byte checksums, else default to 2
	ending := l - 1
	if isPublicKey {
		ending = l - 2
	}
	if len(decoded) < ending {
		return nil, errors.New("invalid decoded length")
	}

	// calculate the hash and do the checksum byte checks
	hash := crypto.NewBlake2b512(decoded[0:ending])
	if hash == nil {
		return nil, errors.New("nil blake hash")
	}

	// run checks
	switch isPublicKey {
	case true:
		{
			if l < 2 || decoded[l-2] != hash[0] || decoded[l-1] != hash[1] {
				return nil, ErrInvalidChecksum
			}
		}

	default:
		{
			if l < 1 || decoded[l-1] != hash[0] {
				return nil, ErrInvalidChecksum
			}
		}
	}

	return decoded[1:ending], nil
}
