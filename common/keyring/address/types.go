package address

import "errors"

var (
	// ErrDecodedLengthNotAllowed ...
	ErrDecodedLengthNotAllowed = errors.New("decoded length not allowed")
	// ErrInvalidChecksum ...
	ErrInvalidChecksum = errors.New("Invalid decoded address checksum")
)
