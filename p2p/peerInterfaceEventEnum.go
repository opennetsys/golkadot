package p2p

import (
	"errors"
	"strings"
)

type peerInterfaceEvent int

const (
	// ActivePeer ...
	ActivePeer peerInterfaceEvent = iota
	// MessagePeer ...
	MessagePeer
	// DisconnectedPeer ...
	DisconnectedPeer
)

// ErrUnknownPeerInterfaceEvent is thrown when an unknown peer interface event is encountered.
var ErrUnknownPeerInterfaceEvent = errors.New("peer interface event: unknown")

// PeerInterfaceEventEnum are the available peer interface events
type PeerInterfaceEventEnum interface {
	Type() peerInterfaceEvent
	String() string
}

// Type returns the enum
func (p peerInterfaceEvent) Type() peerInterfaceEvent {
	return p
}

// AllPeerInterfaceEventEnums returns all of the peer interface event enums
func AllPeerInterfaceEventEnums() []PeerInterfaceEventEnum {
	return []PeerInterfaceEventEnum{
		ActivePeer,
		MessagePeer,
		DisconnectedPeer,
	}
}

// PeerInterfaceEventEnumFromString parses a string to return the peer interface event.
func PeerInterfaceEventEnumFromString(s string) (PeerInterfaceEventEnum, error) {
	switch strings.ToUpper(s) {
	case "ACTIVEPEER":
		return ActivePeer, nil
	case "MESSAGEPEER":
		return MessagePeer, nil
	case "DISCONNECTEDPEER":
		return DisconnectedPeer, nil
	default:
		return nil, ErrUnknownPeerInterfaceEvent
	}
}
