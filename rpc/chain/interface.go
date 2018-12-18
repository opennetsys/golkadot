package chain

import (
	"github.com/c3systems/go-substrate/p2p/pubsub"
	"github.com/c3systems/go-substrate/runtime"
)

// ServiceInterface defines the methods implemented by the chain service
type ServiceInterface interface {
	// GetHeader returns the header of a relay chain block.
	GetHeader(blockHash *string) (*runtime.Header, error)
	// GetBlock returns the header and body of a relay chain block.
	GetBlock(blockHash *string) (runtime.SignedBlock, error)
	// GetBlockHash returns the hash of the n-th block in the canon chain.
	//
	// By default returns latest block hash.
	GetBlockHash(blockHash *string) (*string, error)
	// GetFinalizedHead returns the hash of the last finalised block in the canon chain.
	GetFinalizedHead() (string, error)
	// SubscribeNewHead creates a new head subscription.
	SubscribeNewHead(subscriber pubsub.Subscriber) error
	// UnsubscribeNewHead unsubscribes from new head subscription.
	UnsubscribeNewHead(id pubsub.SubscriptionID) (bool, error)
	// SubscribeFinalizedHeads returns a new head subscription
	SubscribeFinalizedHeads(subscriber pubsub.Subscriber) error
	// UnsubscribeFinalizedHeads from new head subscription.
	UnsubscribeFinalizedHeads(id pubsub.SubscriptionID) (bool, error)
}
