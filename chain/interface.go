package chain

import "github.com/c3systems/go-substrate/p2p/sync"

// Interface describes the methods of the chain service
type Interface interface {
	// note: required from p2p.peer.AddConnection
	GetBestBlocksNumber() (*math.Big, error)
	GetBestBlocksHash() ([]byte, error)
	GetGenesisHash() ([]byte, error)
	// note: required by sync.processBlock
	ImportBlock(block sync.StateBlock) (bool, error)
	// note required by sync.QueuBlocks
	GetBlockDataByHash(hash []byte) (sync.StateBlock, error)
}
