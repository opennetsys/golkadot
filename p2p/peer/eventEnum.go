package peer

import (
	"errors"
	"strings"
)

type eventEnum int

const (
	// Active ...
	Active eventEnum = iota
	// Message ...
	Message
	// Disconnected ...
	Disconnected
)

// ErrUnknownEvent is thrown when an unknown peer interface event is encountered.
var ErrUnknownEvent = errors.New("peer interface event: unknown")

// EventEnum are the exported peer event enums.
type EventEnum interface {
	Type() eventEnum
	String() string
}

// Type returns the private enum
func (e eventEnum) Type() eventEnum {
	return e
}

// AllEventEnums returns all of the peer interface event enums
func AllEventEnums() []EventEnum {
	return []EventEnum{
		Active,
		Message,
		Disconnected,
	}
}

// EventEnumFromString parses a string to return the peer interface event.
func EventEnumFromString(s string) (EventEnum, error) {
	switch strings.ToUpper(s) {
	case "ACTIVE":
		return Active, nil
	case "MESSAGE":
		return Message, nil
	case "DISCONNECTED":
		return Disconnected, nil
	default:
		return nil, ErrUnknownEvent
	}
}
