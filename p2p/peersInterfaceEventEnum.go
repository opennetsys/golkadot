package p2p

import (
	"errors"
	"strings"
)

type peersInterfaceEvent int

const (
	// ActivePeers ...
	ActivePeers peersInterfaceEvent = iota
	// ConnectedPeers ...
	ConnectedPeers
	// DialledPeers ...
	DialledPeers
	// DisconnectedPeers ...
	DisconnectedPeers
	// DiscoveredPeers ...
	DiscoveredPeers
	// MessagePeers ...
	MessagePeers
	// ProtocolPeers ...
	ProtocolPeers
)

// ErrUnknownPeersInterfaceEvent is thrown when an unknown peers interface event is encountered.
var ErrUnknownPeersInterfaceEvent = errors.New("peers interface event: unknown")

// PeersInterfaceEventEnum are the available peers interface events.
type PeersInterfaceEventEnum interface {
	Type() peersInterfaceEvent
	String() string
}

// every base must fullfill the supported interface.
func (p peersInterfaceEvent) Type() peersInterfaceEvent {
	return p
}

// AllPeersInterfaceEventEnums returns all of the peers interface event enums.
func AllPeersInterfaceEventEnums() []PeersInterfaceEventEnum {
	return []PeersInterfaceEventEnum{
		ActivePeers,
		ConnectedPeers,
		DialledPeers,
		DisconnectedPeers,
		DiscoveredPeers,
		MessagePeers,
		ProtocolPeers,
	}
}

// PeersInterfaceEventEnumFromString parses a string to return the peers interface event.
func PeersInterfaceEventEnumFromString(s string) (PeersInterfaceEventEnum, error) {
	switch strings.ToUpper(s) {
	case "ACTIVEPEERS":
		return ActivePeers, nil
	case "CONNECTEDPEERS":
		return ConnectedPeers, nil
	case "DIALLEDPEERS":
		return DialledPeers, nil
	case "DISCONNECTEDPEERS":
		return DisconnectedPeers, nil
	case "DISCOVEREDPEERS":
		return DiscoveredPeers, nil
	case "MESSAGEPEERS":
		return MessagePeers, nil
	case "PROTOCOLPEERS":
		return ProtocolPeers, nil
	default:
		return nil, ErrUnknownPeersInterfaceEvent
	}
}
