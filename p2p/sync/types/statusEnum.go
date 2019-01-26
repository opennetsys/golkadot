package synctypes

import (
	"errors"
	"strings"
)

type statusEnum int

const (
	// Idle TODO
	Idle statusEnum = iota
	// SyncStatus TODO
	// note: we cannot use sync bc the collision with the sync struct
	SyncStatus
)

// ErrUnknownStatus is thrown when an unknown sync status is encountered.
var ErrUnknownStatus = errors.New("sync status: unknown")

// StatusEnum are the exported sync statuses of the node
type StatusEnum interface {
	Type() statusEnum
	String() string
}

// Type returns the internal, package level sync status enum
func (s statusEnum) Type() statusEnum {
	return s
}

// AllStatusEnums returns all of the sync statuses
func AllStatusEnums() []StatusEnum {
	return []StatusEnum{
		Idle,
		SyncStatus,
	}
}

// StatusEnumFromString parses a string to return the sync status
func StatusEnumFromString(s string) (StatusEnum, error) {
	switch strings.ToUpper(s) {
	case "IDLE":
		return Idle, nil
	case "SYNC":
		return SyncStatus, nil
	default:
		return nil, ErrUnknownStatus
	}
}

// String ...
// note: we have to implement this ourselves and not with the stringer package bc of the collision
func (s statusEnum) String() string {
	switch s {
	case Idle:
		return "idle"
	default:
		return "sync"
	}
}
