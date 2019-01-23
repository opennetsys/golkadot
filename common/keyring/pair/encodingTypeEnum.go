package pair

import (
	"errors"
	"strings"
)

type encodingTypeEnum int

const (
	// XSalsa20_Poly1305 ...
	XSalsa20_Poly1305 encodingTypeEnum = iota
	// None ...
	None
)

// ErrUnknownEncodingType ...
var ErrUnknownEncodingType = errors.New("encoding type: unknown")

type EncodingTypeEnum interface {
	Type() encodingTypeEnum
	String() string
}

func (e encodingTypeEnum) Type() encodingTypeEnum {
	return p
}

// AllEncodingTypeEnums returns all of the encoding type enums
func AllEncodingTypeEnums() []EncodingTypeEnum {
	return []EncodingTypeEnum{
		XSalsa20_Poly1305,
	}
}

// EncodingTypeEnumFromString ...
func EncodingTypeEnumFromString(s string) (EncodingTypeEnum, error) {
	switch strings.ToUpper(s) {
	case "XSALSA20_POLY1305":
		return XSalsa20_Poly1305, nil
	default:
		return nil, ErrUnknownEncodingType
	}
}
