package peers

import (
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/peer"

	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
)

// Ensure the struct implements the interface
var _ InterfacePeers = (*Peers)(nil)

// New ...
func New(cfg *Config) (*Peers, error) {
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

	return &Peers{
		Store:      ps,
		KnownPeers: mMap,
	}, nil
}

func (p *Peers) Add(pi pstore.PeerInfo) (*peer.KnownPeer, error) {
	if p.Store == nil {
		return nil, ErrNoStore
	}
	if p.KnownPeers == nil {
		return nil, ErrNoPeerMap
	}

	// note: connect as well???
	p.Store.AddAddrs(pi.ID, pi.Addrs, pstore.PermanentAddrTTL)

	// TODO...
	pr, err := peer.New(nil)
	if err != nil {
		logger.Errorf("[peers] err building new peer\n%v", err)
		return nil, err
	}

	pr.KnownPeers[pi] = pr

	return pr, nil
}

// Count returns the number of connected peers
func (p *Peers) Count() (int, error) {
	if p.Store == nil {
		return nil, ErrNoStore
	}

	return p.Store.PeersWithAddrs().Len(), nil
}

// Get returns a peer
func (p *Peers) Get(pi pstore.PeerInfo) (*peer.KnownPeer, error) {
	if p.KnownPeers == nil {
		return nil, ErrNoPeerMap
	}

	pr, ok := s.KnownPeers[pi]
	if !ok {
		return nil, ErrNoSuchPeer
	}

	return pr, nil
}

// Log TODO
func (p *Peers) Log(event EventEnum, p *peer.KnownPeer) error {
	if event == nil {
		return ErrNilEvent
	}

	// TODO: log pinfo? or peer id?
	logger.Infof("[peers] peer event: %s, from peer: %s", event.String(), p)

	return nil
}

// On handles peers events
func (p *Peers) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO ...
	return nil, nil
}

// Peers returns the peers
func (p *Peers) KnownPeers() ([]*peer.KnownPeer, error) {
	if p.KnownPeers == nil {
		return nil, ErrNoPeerMap
	}

	var knownPeers []*peer.KnownPeer
	for _, v := range p.KnownPeers {
		knownPeers = append(knownPeers, v)
	}

	return knownPeers, nil
}

// TODO: implement _onConnect, _onDisconnect, _onDiscovery
