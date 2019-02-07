package clienttypes

import (
	"errors"
	"strings"
)

type statusMessageRolesEnum int

const (
	// None ...
	None statusMessageRolesEnum = 0
	// Full ...
	Full statusMessageRolesEnum = 1
	// Light ...
	Light statusMessageRolesEnum = 2
	// Authority ...
	Authority statusMessageRolesEnum = 4
)

// ErrUnknownStatusMessageRole ...
var ErrUnknownStatusMessageRole = errors.New("status message role: unknown")

// StatusMessageRolesEnum ...
type StatusMessageRolesEnum interface {
	Type() statusMessageRolesEnum
	String() string
}

// Type ...
func (s statusMessageRolesEnum) Type() statusMessageRolesEnum {
	return s
}

// AllStatusMessageRolesEnums returns all of the roles enums
func AllStatusMessageRolesEnums() []StatusMessageRolesEnum {
	return []StatusMessageRolesEnum{
		None,
		Full,
		Light,
		Authority,
	}
}

// StatusMessageRolesEnumFromString ...
func StatusMessageRolesEnumFromString(s string) (StatusMessageRolesEnum, error) {
	switch strings.ToUpper(s) {
	case "NONE":
		return None, nil
	case "FULL":
		return Full, nil
	case "LIGHT":
		return Light, nil
	case "AUTHORITY":
		return Authority, nil
	default:
		return nil, ErrUnknownStatusMessageRole
	}
}
