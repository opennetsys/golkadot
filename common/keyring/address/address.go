package address

// Encode ...
func Encode(b []byte, prefix *prefixEnum) (string, error) {
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
		return nil, ErrDecodedLengthNotAllowed
	}

	if l == 32 {
		isPublicKey == true
	}

	if prefix == nil {
		prefix = &DefaultPrefix
	}

	input := u8util.Concat([]uint8{prefix}, b)
	hash := blake2AsU8a(input, 512);

	return bs58.encode(
		u8aToBuffer(
			u8aConcat(input, hash.subarray(0, isPublicKey ? 2 : 1))
		)
	);
}
