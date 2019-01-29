package peerstypes

import (
	"errors"
	"strings"
)

type eventEnum int

const (
	// Active ...
	Active eventEnum = iota
	// Connected ...
	Connected
	// Dialled ...
	Dialled
	// Disconnected ...
	Disconnected
	// Discovered ...
	Discovered
	// Message ...
	Message
	// Protocol ...
	Protocol
)

// ErrUnknownEvent is thrown when an unknown peers interface event is encountered.
var ErrUnknownEvent = errors.New("peers interface event: unknown")

// EventEnum are the exported peers events enums.
type EventEnum interface {
	Type() eventEnum
	String() string
}

// Type returns the private peers event enum.
func (e eventEnum) Type() eventEnum {
	return e
}

// AllEventEnums returns all of the peers interface event enums.
func AllEventEnums() []EventEnum {
	return []EventEnum{
		Active,
		Connected,
		Dialled,
		Disconnected,
		Discovered,
		Message,
		Protocol,
	}
}

// EventEnumFromString parses a string to return the peers interface event.
func EventEnumFromString(s string) (EventEnum, error) {
	switch strings.ToUpper(s) {
	case "ACTIVE":
		return Active, nil
	case "CONNECTED":
		return Connected, nil
	case "DIALLED":
		return Dialled, nil
	case "DISCONNECTED":
		return Disconnected, nil
	case "DISCOVERED":
		return Discovered, nil
	case "MESSAGE":
		return Message, nil
	case "PROTOCOL":
		return Protocol, nil
	default:
		return nil, ErrUnknownEvent
	}
}
