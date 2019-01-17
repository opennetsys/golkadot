package address

import (
	"errors"
	"strings"
)

type prefixEnum int

const (
	Zero       prefixEnum = 0
	One        prefixEnum = 1
	Three      prefixEnum = 3
	FortyTwo   prefixEnum = 42
	FortyThree prefixEnum = 43
	SixtyEight prefixEnum = 68
	SixtyNine  prefixEnum = 69
)

// ErrUnknownPrefix ...
var ErrUnknownPrefix = errors.New("prefix: unknown")

type PrefixEnum interface {
	Type() prefixEnum
	String() string
}

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
