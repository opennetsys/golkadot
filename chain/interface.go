package chain

import (
	"math/big"

	synctypes "github.com/c3systems/go-substrate/p2p/sync/types"
)

// InterfaceChain describes the methods of the chain service
type InterfaceChain interface {
	// note: required from p2p.peer.AddConnection
	GetBestBlocksNumber() (*big.Int, error)
	GetBestBlocksHash() ([]byte, error)
	GetGenesisHash() ([]byte, error)
	// note: required by sync.processBlock
	ImportBlock(block synctypes.StateBlock) (bool, error)
	// note required by sync.QueuBlocks
	GetBlockDataByHash(hash []byte) (synctypes.StateBlock, error)
}
