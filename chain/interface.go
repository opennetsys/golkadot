package chain

// Interface describes the methods of the chain service
type Interface interface {
	// note: required from p2p.peer.AddConnection
	GetBestBlocksNumber() (*math.Big, error)
	GetBestBlocksHash() ([]byte, error)
	GetGenesisHash() ([]byte, error)
}
