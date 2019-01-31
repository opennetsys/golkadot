package address

import (
	"errors"
	"strings"
)

type prefixEnum int

const (
	// Zero ...
	Zero prefixEnum = 0
	// One ...
	One prefixEnum = 1
	// Three ..
	Three prefixEnum = 3
	// FortyTwo ...
	FortyTwo prefixEnum = 42
	// FortyThree ...
	FortyThree prefixEnum = 43
	// SixtyEight ...
	SixtyEight prefixEnum = 68
	// SixtyNine ...
	SixtyNine prefixEnum = 69
)

// ErrUnknownPrefix ...
var ErrUnknownPrefix = errors.New("prefix: unknown")

// PrefixEnum ...
type PrefixEnum interface {
	Type() prefixEnum
	String() string
}

// Type ...
func (p prefixEnum) Type() prefixEnum {
	return p
}

// AllPrefixEnums returns all of the prefix enums
func AllPrefixEnums() []PrefixEnum {
	return []PrefixEnum{
		Zero,
		One,
		Three,
		FortyTwo,
		FortyThree,
		SixtyEight,
		SixtyNine,
	}
}

// PrefixEnumFromString ...
func PrefixEnumFromString(s string) (PrefixEnum, error) {
	switch strings.ToUpper(s) {
	case "ZERO":
		return Zero, nil
	case "ONE":
		return One, nil
	case "THREE":
		return Three, nil
	case "FORTYTWO":
		return FortyTwo, nil
	case "FORTYTHREE":
		return FortyThree, nil
	case "SIXYEIGHT":
		return SixtyEight, nil
	case "SIXTYNINE":
		return SixtyNine, nil
	default:
		return nil, ErrUnknownPrefix
	}
}
