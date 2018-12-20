package p2p

import (
	"errors"
	"strings"
)

type p2pInterfaceEvent int

const (
	// Started ...
	Started p2pInterfaceEvent = iota
	// Stopped ...
	Stopped
)

// ErrUnknownP2pInterfaceEvent is thrown when an unknown p2p interface event is encountered.
var ErrUnknownP2pInterfaceEvent = errors.New("p2p interface event: unknown")

// P2pInterfaceEventEnum are the available p2p interface events.
type P2pInterfaceEventEnum interface {
	Type() p2pInterfaceEvent
	String() string
}

// every base must fullfill the supported interface.
func (p p2pInterfaceEvent) Type() p2pInterfaceEvent {
	return p
}

// AllP2pInterfaceEventEnums returns all of the p2p interface event enums.
func AllP2pInterfaceEventEnums() []P2pInterfaceEventEnum {
	return []P2pInterfaceEventEnum{
		Started,
		Stopped,
	}
}

// P2pInterfaceEventEnumFromString parses a string to return the p2p interface event.
func P2pInterfaceEventEnumFromString(s string) (P2pInterfaceEventEnum, error) {
	switch strings.ToUpper(s) {
	case "STARTED":
		return Started, nil
	case "STOPPED":
		return Stopped, nil
	default:
		return nil, ErrUnknownP2pInterfaceEvent
	}
}
