package state

import (
	"github.com/c3systems/go-substrate/client/p2p/pubsub"
	"github.com/c3systems/go-substrate/common/storage"
)

// Service implements the service interface
type Service interface{}

// CallArgs is passed to the Call method
type CallArgs struct {
	Method    string
	Data      []byte
	BlockHash *string
}

// StorageArgs is passed to the various storage methods
type StorageArgs struct {
	Key       *storage.Key
	BlockHash *string
}

// QueryStorageArgs is passed to the QueryStorage method
type QueryStorageArgs struct {
	Keys          []*storage.Key
	FromBlockHash string
	ToBlockHash   *string
}

// SubscribeStorageArgs is passed to the SubscribeRuntimeVersion method
type SubscribeStorageArgs struct {
	Subscriber pubsub.Subscriber
	Keys       []*storage.Key
}
