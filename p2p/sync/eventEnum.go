package sync

import (
	"errors"
	"strings"
)

type eventEnum int

const (
	// Imported TODO
	Imported eventEnum = iota
)

// ErrUnknownEvent is thrown when an unknown sync event is encountered.
var ErrUnknownEvent = errors.New("sync event: unknown")

// EventEnum are the exported sync events
type EventEnum interface {
	Type() eventEnum
	String() string
}

// Type returns the non-exported enum
func (e eventEnum) Type() eventEnum {
	return e
}

// AllEventEnums returns all of the sync events
func AllEventEnums() []EventEnum {
	return []EventEnum{
		Imported,
	}
}

// EventEnumFromString parses a string to return the sync event
func EventEnumFromString(s string) (EventEnum, error) {
	switch strings.ToUpper(s) {
	case "IMPORTED":
		return Imported, nil
	default:
		return nil, ErrUnknownEvent
	}
}
