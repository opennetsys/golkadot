package block

import (
	"errors"
	"strings"
)

type digestEnum int

const (
	// AuthoritiesChange ...
	AuthoritiesChange digestEnum = iota
	// ChangesTrieRoot ...
	ChangesTrieRoot
	// Other ...
	Other
	// Seal ...
	Seal
)

// ErrUnknownDigest ...
var ErrUnknownDigest = errors.New("digest: unknown")

type DigestEnum interface {
	Type() digestEnum
	String() string
}

func (d digestEnum) Type() digestEnum {
	return d
}

// AllDigestEnums returns all of the digest enums
func AllDigestEnums() []DigestEnum {
	return []DigestEnum{
		AuthoritiesChange,
		ChangesTrieRoot,
		Other,
		Seal,
	}
}

// DigestEnumFromString ...
func DigestEnumFromString(s string) (DigestEnum, error) {
	switch strings.ToUpper(s) {
	case "AUTHORITIESCHANGE":
		return AuthoritiesChange, nil
	case "CHANGESTRIEROOT":
		return ChangesTrieRoot, nil
	case "OTHER":
		return Other, nil
	case "SEAL":
		return Seal, nil
	default:
		return nil, ErrUnknownDigest
	}
}
