package p2p

import (
	"errors"
	"strings"
)

type syncStatus int

const (
	// Idle ...
	Idle syncStatus = iota
	// Sync ...
	Sync
)

// ErrUnknownSyncStatus is thrown when an unknown sync status is encountered.
var ErrUnknownSyncStatus = errors.New("sync status: unknown")

// SyncStatusEnum are the available sync statuses of the node
type SyncStatusEnum interface {
	Type() syncStatus
	String() string
}

// every base must fullfill the supported interface
func (s syncStatus) Type() syncStatus {
	return s
}

// AllSyncStatusEnums returns all of the sync statuses
func AllSyncStatusEnums() []SyncStatusEnum {
	return []SyncStatusEnum{
		Idle,
		Sync,
	}
}

// SyncStatusEnumFromString parses a string to return the sync status
func SyncStatusEnumFromString(s string) (SyncStatusEnum, error) {
	switch strings.ToUpper(s) {
	case "IDLE":
		return Idle, nil
	case "SYNC":
		return Sync, nil
	default:
		return nil, ErrUnknownSyncStatus
	}
}
