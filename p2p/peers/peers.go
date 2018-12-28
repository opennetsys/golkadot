package peers

import (
	"github.com/c3systems/go-substrate/p2p/peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// New ...
func New(cfg *Config) (*Service, error) {
	// TODO
	return nil, nil
}

// Add ...
func (s *Service) Add(peerInfo pstore.PeerInfo) (peer.Interface, error) {
	// TODO
	return nil, nil
}

// Count ...
func (s *Service) Count() uint {
	// TODO
	return 0
}

// Get ...
func (s *Service) Get(peerInfo pstore.PeerInfo) (*peer.KnownPeer, error) {
	// TODO
	return nil, nil
}

// Log ...
func (s *Service) Log(event EventEnum, p peer.Interface) error {
	// TODO
	return nil
}

// On ...
func (s *Service) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO
	return nil, nil
}

// Peers ...
func (s *Service) Peers() []peer.Interface {
	// TODO
	return nil
}
