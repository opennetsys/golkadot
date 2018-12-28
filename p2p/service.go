package p2p

import (
	"context"
	"time"

	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/peers"
	p2psync "github.com/c3systems/go-substrate/p2p/sync"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	net "github.com/libp2p/go-libp2p-net"
	libpeer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	pstoremem "github.com/libp2p/go-libp2p-peerstore/pstoremem"
	swarm "github.com/libp2p/go-libp2p-swarm"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	bhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	tcp "github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// Ensure the struct implements the interface
var _ Interface = (*Service)(nil)

// New builds a new p2p service
func New(cfg *Config, c chain.Interface) (*Service, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if c == nil {
		return nil, ErrNoChainService
	}

	pid, err := libpeer.IDFromPublicKey(cfg.Pub)
	if err != nil {
		logger.Errorf("[p2p] err generating peer id from public key\n%v", err)
		return nil, err
	}

	listen, err := ma.NewMultiaddr(cfg.Address)
	if err != nil {
		logger.Errorf("[p2p] err parsing address\n%v", err)
		return nil, err
	}

	// TODO: pass in pstore?
	ps := pstoremem.NewPeerstore()
	if err := ps.AddPrivKey(pid, cfg.Priv); err != nil {
		logger.Errorf("[p2p] err adding private keey to peer store\n%v", err)
		return nil, err
	}
	if err := ps.AddPubKey(pid, cfg.Pub); err != nil {
		logger.Errorf("[p2p] err adding public key to peer store\n%v", err)
		return nil, err
	}

	// TODO ...
	swarmNet := swarm.NewSwarm(cfg.Context, pid, ps, nil)
	tcpTransport := tcp.NewTCPTransport(genUpgrader(swarmNet))
	if err := swarmNet.AddTransport(tcpTransport); err != nil {
		logger.Errorf("[p2p] err adding transport to swarmnet\n%v", err)
		return nil, err
	}
	if err := swarmNet.AddListenAddr(listen); err != nil {
		logger.Errorf("[p2p] err adding swarm listen addr %v to swarmnet\n%v", listen, err)
		return nil, err
	}
	bNode := bhost.New(swarmNet)

	// TODO ...
	dhtSvc, err := dht.New(cfg.Context, bNode)
	if err != nil {
		logger.Errorf("[p2p] err building dht service\n%v", err)
		return nil, err
	}
	if err := dhtSvc.Bootstrap(cfg.Context); err != nil {
		logger.Errorf("[p2p] err bootstrapping dht\n%v", err)
		return nil, err
	}

	newNode := rhost.Wrap(bNode, dhtSvc)
	for i, addr := range newNode.Addrs() {
		logger.Infof("[p2p] %d: %s/ipfs/%s\n", i, addr, newNode.ID().Pretty())
	}

	// TODO ...
	// note: https://libp2p.io/implementations/#discovery
	// note: use https://github.com/libp2p/go-libp2p/blob/master/p2p/discovery/mdns.go rather than whyrusleeping
	discoverySvc, err := discovery.NewMdnsService(cfg.Context, newNode, time.Second, Name)
	if err != nil {
		logger.Errorf("[p2p] err starting discover service\n%v", err)
		return nil, err
	}
	discoverySvc.RegisterNotifee(&DiscoveryNotifee{newNode})

	// TODO: pubsub chan
	//pubsub, err := floodsub.NewFloodSub(ctx, newNode)
	//if err != nil {
	//return nil, fmt.Errorf("err building new pubsub service\n%v", err)
	//}

	// TODO ...
	//if cfg.Peer != "" {
	//addr, err := ipfsaddr.ParseString(cfg.Peer)
	//if err != nil {
	//return nil, fmt.Errorf("err parsing node uri flag: %s\n%v", cfg.URI, err)
	//}

	//pinfo, err := peerstore.InfoFromP2pAddr(addr.Multiaddr())
	//if err != nil {
	//return nil, fmt.Errorf("err getting info from peerstore\n%v", err)
	//}

	//log.Println("[node] FULL", addr.String())
	//log.Println("[node] PIN INFO", pinfo)

	//if err := newNode.Connect(ctx, *pinfo); err != nil {
	//return nil, fmt.Errorf("[node] bootstrapping a peer failed\n%v", err)
	//}

	//newNode.Peerstore().AddAddrs(pinfo.ID, pinfo.Addrs, peerstore.PermanentAddrTTL)
	//}

	s := &Service{
		state: &State{
			Chain:  c,
			Config: cfg,
			// TODO ...
			SyncState: &p2psync.State{},
		},
	}
	nb := &net.NotifyBundle{
		ConnectedF: s.onConn,
	}
	newNode.Network().Notify(nb)
	s.state.Host = newNode

	prs, err := peers.New(&peers.Config{})
	if err != nil {
		logger.Errorf("[p2p] err creating new peers\n%v", err)
	}
	s.state.Peers = prs

	return s, nil
}

// IsStarted ...
func (s *Service) IsStarted() bool {
	// TODO: best practice for determining if server is started?
	if s.state == nil || s.state.Host == nil || len(s.state.Host.Addrs()) == 0 {
		return false
	}

	return true
}

// GetNumPeers ...
func (s *Service) GetNumPeers() (uint, error) {
	if s.state == nil {
		return 0, ErrUninitializedService
	}
	if s.state.Peers == nil {
		return 0, ErrNoPeersService
	}

	return s.state.Peers.Count(), nil
}

// On handles messages
func (s *Service) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO
	return nil, nil
}

// Start starts the p2p service
func (s *Service) Start(ctx context.Context, ch chan interface{}) error {
	//if err := n.listenForEvents(); err != nil {
	//return nil, fmt.Errorf("error starting listener\n%v", err)
	//}

	// TODO
	return nil
}

// Stop stops the p2p service
func (s *Service) Stop(ch chan interface{}) error {
	// TODO
	return nil
}

func (s *Service) onConn(network net.Network, conn net.Conn) {
	logger.Infof("[p2p] peer did connect\nid %v peerAddr %v", conn.RemotePeer().Pretty(), conn.RemoteMultiaddr())

	s.addAddr(conn)
}

func (s *Service) addAddr(conn net.Conn) {
	if s.state == nil || s.state.Host == nil {
		logger.Warnf("[p2p] no state or host, cannot add peer %s", conn.RemoteMultiaddr())
		return
	}

	for _, peer := range s.state.Host.Peerstore().Peers() {
		if conn.RemotePeer() == peer {
			// note: we already have info on this peer
			logger.Warnf("[p2p] already have peer in our peerstore")
			return
		}
	}

	// note: we don't have this peer's info
	s.state.Host.Peerstore().AddAddr(conn.RemotePeer(), conn.RemoteMultiaddr(), peerstore.PermanentAddrTTL)
	logger.Infof("[p2p] added %s to our peerstore", conn.RemoteMultiaddr())

	if _, err := s.state.Host.Network().DialPeer(context.Background(), conn.RemotePeer()); err != nil {
		logger.Warnf("[p2p] err connecting to a peer\n%v", err)

		return
	}

	logger.Infof("[p2p] connected to %s", conn.RemoteMultiaddr())
}
