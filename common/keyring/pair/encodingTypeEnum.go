package pair

import (
	"errors"
	"fmt"
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

// EncodingTypeEnum ...
type EncodingTypeEnum interface {
	Type() encodingTypeEnum
	String() string
}

func (e encodingTypeEnum) Type() encodingTypeEnum {
	return e
}

// AllEncodingTypeEnums returns all of the encoding type enums
func AllEncodingTypeEnums() []EncodingTypeEnum {
	return []EncodingTypeEnum{
		XSalsa20_Poly1305,
		None,
	}
}

// EncodingTypeEnumFromString ...
func EncodingTypeEnumFromString(s string) (EncodingTypeEnum, error) {
	switch strings.ToUpper(s) {
	case "XSALSA20-POLY1305":
		return XSalsa20_Poly1305, nil
	case "NONE":
		return None, nil
	default:
		return nil, ErrUnknownEncodingType
	}
}

// String ...
// note: we do not use the built-in 'stringer' tool because we need a hyphen rather than an underscore for xsalsa.
func (e encodingTypeEnum) String() string {
	switch e {
	case XSalsa20_Poly1305:
		return "xsalsa20-poly1305"
	case None:
		return "none"
	default:
		return ""
	}
}

// MarshalJSON ...
func (e encodingTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, e.String())), nil
}

// UnmarshalJSON ...
func (e *encodingTypeEnum) UnmarshalJSON(data []byte) error {
	ETE, err := EncodingTypeEnumFromString(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	typ := ETE.Type()
	// TODO: nil check?
	*e = typ

	return nil
}
