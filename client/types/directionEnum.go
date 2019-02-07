package clienttypes

import (
	"errors"
	"strings"
)

type directionEnum int

const (
	// Ascending ...
	Ascending directionEnum = iota
	// Descending ...
	Descending
)

// ErrUnknownDirection ...
var ErrUnknownDirection = errors.New("direction: unknown")

// DirectionEnum ...
type DirectionEnum interface {
	Type() directionEnum
	String() string
}

// Type ...
func (d directionEnum) Type() directionEnum {
	return d
}

// AllDirectionEnums returns all of the direction enums
func AllDirectionEnums() []DirectionEnum {
	return []DirectionEnum{
		Ascending,
		Descending,
	}
}

// DirectionEnumFromString ...
func DirectionEnumFromString(s string) (DirectionEnum, error) {
	switch strings.ToUpper(s) {
	case "ASCENDING":
		return Ascending, nil
	case "DESCENDING":
		return Descending, nil
	default:
		return nil, ErrUnknownDirection
	}
}
