package peers

import (
	"github.com/c3systems/go-substrate/client/p2p/peer"
	peerstypes "github.com/c3systems/go-substrate/client/p2p/peers/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/logger"

	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
)

// Ensure the struct implements the interface
var _ clienttypes.InterfacePeers = (*Peers)(nil)

// New ...
func New(cfg *clienttypes.ConfigPeers) (*Peers, error) {
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

	pMap := make(map[libpeer.ID]*clienttypes.KnownPeer)

	return &Peers{
		Store:         ps,
		KnownPeersMap: pMap,
	}, nil
}

// Add ...
func (p *Peers) Add(pi pstore.PeerInfo) (*clienttypes.KnownPeer, error) {
	if p.Store == nil {
		return nil, ErrNoStore
	}
	if p.KnownPeersMap == nil {
		return nil, ErrNoPeerMap
	}

	// note: connect as well???
	p.Store.AddAddrs(pi.ID, pi.Addrs, pstore.PermanentAddrTTL)

	// TODO...
	cfg := &clienttypes.ConfigPeer{}
	pr, err := peer.New(cfg, nil, pstore.PeerInfo{})
	if err != nil {
		logger.Errorf("[peers] err building new peer\n%v", err)
		return nil, err
	}

	kp := &clienttypes.KnownPeer{
		Peer: pr,
		// TODO: true?
		IsConnected: true,
	}
	p.KnownPeersMap[pi.ID] = kp

	return kp, nil
}

// Count returns the number of connected peers
func (p *Peers) Count() (uint, error) {
	if p.Store == nil {
		return 0, ErrNoStore
	}

	pCount := p.Store.PeersWithAddrs().Len()
	return uint(pCount), nil
}

// Get returns a peer
func (p *Peers) Get(pi pstore.PeerInfo) (*clienttypes.KnownPeer, error) {
	if p.KnownPeersMap == nil {
		return nil, ErrNoPeerMap
	}

	pr, ok := p.KnownPeersMap[pi.ID]
	if !ok {
		return nil, ErrNoSuchPeer
	}

	return pr, nil
}

// Log TODO
func (p *Peers) Log(event peerstypes.EventEnum, kp *clienttypes.KnownPeer) error {
	if event == nil {
		return ErrNilEvent
	}

	// TODO: log pinfo? or peer id?
	logger.Infof("[peers] peer event: %s, from peer: %v", event.String(), p)

	return nil
}

// On handles peers events
func (p *Peers) On(event peerstypes.EventEnum, cb peerstypes.EventCallback) (interface{}, error) {
	// TODO ...
	return nil, nil
}

// KnownPeers returns the peers
func (p *Peers) KnownPeers() ([]*clienttypes.KnownPeer, error) {
	if p.KnownPeersMap == nil {
		return nil, ErrNoPeerMap
	}

	var knownPeers []*clienttypes.KnownPeer
	for _, v := range p.KnownPeersMap {
		knownPeers = append(knownPeers, v)
	}

	return knownPeers, nil
}

// TODO: implement _onConnect, _onDisconnect, _onDiscovery
