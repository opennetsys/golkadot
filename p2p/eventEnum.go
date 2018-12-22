package p2p

import (
	"errors"
	"strings"
)

type eventEnum int

const (
	// Started ...
	Started eventEnum = iota
	// Stopped ...
	Stopped
)

// ErrUnknownEvent is thrown when an unknown p2p interface event is encountered.
var ErrUnknownEvent = errors.New("p2p interface event: unknown")

// EventEnum are the available p2p interface events.
type EventEnum interface {
	Type() eventEnum
	String() string
}

// Type is used to return the enum.
func (e eventEnum) Type() eventEnum {
	return e
}

// AllEventEnums returns all of the p2p interface event enums.
func AllEventEnums() []EventEnum {
	return []EventEnum{
		Started,
		Stopped,
	}
}

// EventEnumFromString parses a string to return the p2p interface event.
func EventEnumFromString(s string) (EventEnum, error) {
	switch strings.ToUpper(s) {
	case "STARTED":
		return Started, nil
	case "STOPPED":
		return Stopped, nil
	default:
		return nil, ErrUnknownEvent
	}
}
