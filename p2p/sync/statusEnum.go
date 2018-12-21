package sync

import (
	"errors"
	"strings"
)

type statusEnum int

const (
	// Idle ...
	Idle statusEnum = iota
	// Sync ...
	Sync
)

// ErrUnknownStatus is thrown when an unknown sync status is encountered.
var ErrUnknownStatus = errors.New("sync status: unknown")

// StatusEnum are the available sync statuses of the node
type StatusEnum interface {
	Type() statusEnum
	String() string
}

// every base must fullfill the supported interface
func (s statusEnum) Type() statusEnum {
	return s
}

// AllStatusEnums returns all of the sync statuses
func AllStatusEnums() []StatusEnum {
	return []StatusEnum{
		Idle,
		Sync,
	}
}

// StatusEnumFromString parses a string to return the sync status
func StatusEnumFromString(s string) (StatusEnum, error) {
	switch strings.ToUpper(s) {
	case "IDLE":
		return Idle, nil
	case "SYNC":
		return Sync, nil
	default:
		return nil, ErrUnknownStatus
	}
}
