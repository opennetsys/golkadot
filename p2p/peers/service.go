package peers

import (
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/peer"

	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
)

// Ensure the struct implements the interface
var _ Interface = (*Service)(nil)

// New ...
func New(cfg *Config) (*Service, error) {
	// TODO ...
	if cfg == nil {
		return nil, ErrNoConfig
	}

	ps := pstoremem.NewPeerstore()
	if err := ps.AddPrivKey(cfg.ID, cfg.Priv); err != nil {
		logger.Errorf("[peers] err adding private keey to peer store\n%v", err)
		return nil, err
	}
	if err := ps.AddPubKey(cfg.ID, cfg.Pub); err != nil {
		logger.Errorf("[peers] err adding public key to peer store\n%v", err)
		return nil, err
	}

	pMap := make(map[pstore.PeerInfo]*peer.KnownPeer)

	return &Service{
		Store: ps,
		Peers: mMap,
	}, nil
}

func (s *Service) Add(pi pstore.PeerInfo) (*peer.KnownPeer, error) {
	if s.Store == nil {
		return nil, ErrNoStore
	}
	if s.Peers == nil {
		return nil, ErrNoPeerMap
	}

	// note: connect as well???
	s.Store.AddAddrs(pi.ID, pi.Addrs, pstore.PermanentAddrTTL)

	// TODO...
	p, err := peer.New(nil)
	if err != nil {
		logger.Errorf("[peers] err building new peer\n%v", err)
		return nil, err
	}

	p.Peers[pi] = p

	return p, nil
}

// Count returns the number of connected peers
func (s *Service) Count() (int, error) {
	if s.Store == nil {
		return nil, ErrNoStore
	}

	return s.Store.PeersWithAddrs().Len(), nil
}

// Get returns a peer
func (s *Service) Get(pi pstore.PeerInfo) (*peer.KnownPeer, error) {
	if s.Peers == nil {
		return nil, ErrNoPeerMap
	}

	p, ok := s.Peers[pi]
	if !ok {
		return nil, ErrNoSuchPeer
	}

	return p, nil
}

// Log TODO
func (s *Service) Log(event EventEnum, p *peer.KnownPeer) error {
	if event == nil {
		return ErrNilEvent
	}

	// TODO: log pinfo? or peer id?
	logger.Infof("[peers] peer event: %s, from peer: %s", event.String(), p)

	return nil
}

// On handles peers events
func (s *Service) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO ...
	return nil, nil
}

// Peers returns the peers
func (s *Service) Peers() ([]*peer.KnownPeer, error) {
	if s.Peers == nil {
		return nil, ErrNoPeerMap
	}

	var knownPeers []*peer.KnownPeer
	for _, v := range s.Peers {
		knownPeers = append(knownPeers, v)
	}

	return knownPeers, nil
}

// TODO: implement _onConnect, _onDisconnect, _onDiscovery
