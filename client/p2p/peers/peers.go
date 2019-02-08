package peers

import (
	"errors"
	"time"

	"github.com/opennetsys/golkadot/client/p2p/defaults"
	"github.com/opennetsys/golkadot/client/p2p/peer"
	peertypes "github.com/opennetsys/golkadot/client/p2p/peer/types"
	peerstypes "github.com/opennetsys/golkadot/client/p2p/peers/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/logger"

	libp2pHost "github.com/libp2p/go-libp2p-host"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
	"github.com/libp2p/go-libp2p/p2p/discovery"
)

// Ensure the struct implements the interface
var _ clienttypes.InterfacePeers = (*Peers)(nil)

// New ...
func New(cfg *clienttypes.ConfigClient, chn clienttypes.InterfaceChains, host libp2pHost.Host) (*Peers, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if cfg.Peers == nil {
		return nil, errors.New("nil peers config")
	}
	if cfg.P2P == nil {
		return nil, errors.New("nil p2p config")
	}
	if chn == nil {
		return nil, errors.New("nil chain")
	}
	if host == nil {
		return nil, errors.New("nil host")
	}

	ps := pstoremem.NewPeerstore()
	if err := ps.AddPrivKey(cfg.Peers.ID, cfg.Peers.Priv); err != nil {
		logger.Errorf("[peers] err adding private keey to peer store\n%v", err)
		return nil, err
	}
	if err := ps.AddPubKey(cfg.Peers.ID, cfg.Peers.Pub); err != nil {
		logger.Errorf("[peers] err adding public key to peer store\n%v", err)
		return nil, err
	}

	p := &Peers{
		Store:         ps,
		KnownPeersMap: make(map[libpeer.ID]*clienttypes.KnownPeer),
		cfg:           cfg,
		handlers:      make(map[peerstypes.EventEnum]clienttypes.PeersEventCallback),
		chain:         chn,
		host:          host,
	}
	if err := p.onDiscovery(); err != nil {
		return nil, err
	}
	if err := p.onConnectAndDisconnect(); err != nil {
		return nil, err
	}

	return p, nil
}

// Add ...
func (p *Peers) Add(pi pstore.PeerInfo) (*clienttypes.KnownPeer, error) {
	if p.Store == nil {
		return nil, ErrNoStore
	}
	if p.KnownPeersMap == nil {
		return nil, ErrNoPeerMap
	}

	if kp, ok := p.KnownPeersMap[pi.ID]; ok {
		return kp, nil
	}

	// note: connect as well???
	p.Store.AddAddrs(pi.ID, pi.Addrs, pstore.PermanentAddrTTL)

	// TODO...
	pr, err := peer.New(p.cfg, p.chain, pi)
	if err != nil {
		logger.Errorf("[peers] err building new peer\n%v", err)
		return nil, err
	}

	kp := &clienttypes.KnownPeer{
		Peer: pr,
		// TODO: true?
		IsConnected: false,
	}
	p.KnownPeersMap[pi.ID] = kp

	kp.Peer.On(peertypes.Active, func(iface interface{}) (interface{}, error) {
		//kp, ok := p.KnownPeersMap[pi.ID]
		//if !ok {
		//return nil, errors.New("peer not known")
		//}
		//if kp == nil {
		//return nil, errors.New("nil peer")
		//}

		kp.IsConnected = true
		if err := p.Log(peerstypes.Active, kp); err != nil {
			logger.Errorf("[peers] logging err\n%v", err)
			return nil, err
		}

		return nil, nil
	})

	kp.Peer.On(peertypes.Disconnected, func(iface interface{}) (interface{}, error) {
		//kp, ok := p.KnownPeersMap[pi.ID]
		//if !ok {
		//return nil, errors.New("peer not known")
		//}
		//if kp == nil {
		//return nil, errors.New("nil peer")
		//}

		if !kp.IsConnected {
			return nil, nil
		}

		kp.IsConnected = false
		if err := p.Log(peerstypes.Disconnected, kp); err != nil {
			logger.Errorf("[peers] logging err\n%v", err)
			return nil, err
		}

		return nil, nil
	})

	kp.Peer.On(peertypes.Message, func(iface interface{}) (interface{}, error) {
		if iface == nil {
			return nil, errors.New("nil message")
		}
		msg, ok := iface.(clienttypes.InterfaceMessage)
		if !ok {
			logger.Errorf("[peers] want InterfaceMessagel have %T", iface)
			return nil, errors.New("iface not message")
		}

		p.handleEvent(peerstypes.Message, &clienttypes.OnMessage{
			Peer:    kp.Peer,
			Message: msg,
		})

		return nil, nil
	})

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

// CountAll returns the number of known peers
func (p *Peers) CountAll() (uint, error) {
	if p.KnownPeersMap == nil {
		return 0, errors.New("no known peers map")
	}

	pCount := len(p.KnownPeersMap)
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

// GetFromID ...
func (p *Peers) GetFromID(id libpeer.ID) (*clienttypes.KnownPeer, error) {
	if p.KnownPeersMap == nil {
		return nil, ErrNoPeerMap
	}

	pr, ok := p.KnownPeersMap[id]
	if !ok {
		return nil, ErrNoSuchPeer
	}

	return pr, nil

}

// Log ...
func (p *Peers) Log(event peerstypes.EventEnum, iface interface{}) error {
	if event == nil {
		return ErrNilEvent
	}

	// TODO: log pinfo? or peer id?
	logger.Infof("[peers] peer event: %s, interface: %v", event.String(), iface)

	p.handleEvent(event, iface)

	return nil
}

// On handles peers events
func (p *Peers) On(event peerstypes.EventEnum, cb clienttypes.PeersEventCallback) {
	p.handlers[event] = cb

	return
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

func (p *Peers) handleEvent(event peerstypes.EventEnum, iface interface{}) {
	if event == nil {
		logger.Info("[peers] nil event")
		return
	}

	cb, ok := p.handlers[event]
	if !ok {
		logger.Infof("[peers] no event for %s", event.String())
		return
	}

	iface, err := cb(iface)
	logger.Infof("[peers] handled event %s\nresults:\n%v\n%v", event.String(), iface, err)
	return
}

func (p *Peers) onDiscovery() error {
	if p.cfg == nil {
		return errors.New("nil config")
	}
	if p.cfg.P2P == nil {
		return errors.New("nil p2p config")
	}

	// TODO: we already have a discovery service in P2P, is this duplicating? How else to notify on peer found?
	discoverySvc, err := discovery.NewMdnsService(p.cfg.P2P.Context, p.host, time.Second, defaults.Defaults.Name)
	if err != nil {
		logger.Errorf("[peers] err starting discover service\n%v", err)
		return err
	}

	discoverySvc.RegisterNotifee(&DiscoveryNotifee{p})

	return nil
}

func (p *Peers) onConnectAndDisconnect() error {
	if p.host == nil {
		return errors.New("nil host")
	}

	n := &Notifiee{p}
	// TODO: nil check?
	p.host.Network().Notify(n)

	return nil
}
