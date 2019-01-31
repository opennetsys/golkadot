package pair

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type encodingContentEnum int

const (
	// PKCS8 ...
	PKCS8 encodingContentEnum = iota
)

// ErrUnknownEncodingContent ...
var ErrUnknownEncodingContent = errors.New("encoding content: unknown")

// EncodingContentEnum ...
type EncodingContentEnum interface {
	Type() encodingContentEnum
	String() string
}

func (e encodingContentEnum) Type() encodingContentEnum {
	return e
}

// AllEncodingContentEnums returns all of the encoding content enums
func AllEncodingContentEnums() []EncodingContentEnum {
	return []EncodingContentEnum{
		PKCS8,
	}
}

// EncodingContentEnumFromString ...
func EncodingContentEnumFromString(s string) (EncodingContentEnum, error) {
	switch strings.ToUpper(s) {
	case "PKCS8":
		return PKCS8, nil
	default:
		return nil, ErrUnknownEncodingContent
	}
}

// MarshalJSON ...
func (e encodingContentEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, e.String())), nil
}

// Unmarshal ...
func (e *encodingContentEnum) UnmarshalJSON(data []byte) error {
	ECE, err := EncodingContentEnumFromString(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		log.Println("err getting content from string")
		return err
	}

	typ := ECE.Type()
	// TODO: nil check?
	*e = typ

	return nil
}
