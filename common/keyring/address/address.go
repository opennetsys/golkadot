package address

import (
	"errors"

	"github.com/c3systems/go-substrate/common/crypto"
	"github.com/c3systems/go-substrate/common/u8util"

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
