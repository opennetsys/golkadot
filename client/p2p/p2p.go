package p2p

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"io"
	"time"

	"github.com/opennetsys/go-substrate/client/p2p/defaults"
	"github.com/opennetsys/go-substrate/client/p2p/handler"
	"github.com/opennetsys/go-substrate/client/p2p/peers"
	peerstypes "github.com/opennetsys/go-substrate/client/p2p/peers/types"
	"github.com/opennetsys/go-substrate/client/p2p/sync"
	p2ptypes "github.com/opennetsys/go-substrate/client/p2p/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/logger"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	inet "github.com/libp2p/go-libp2p-net"
	libpeer "github.com/libp2p/go-libp2p-peer"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	pstoremem "github.com/libp2p/go-libp2p-peerstore/pstoremem"
	protocol "github.com/libp2p/go-libp2p-protocol"
	swarm "github.com/libp2p/go-libp2p-swarm"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	bhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	tcp "github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// Ensure the struct implements the interface
var _ clienttypes.InterfaceP2P = (*P2P)(nil)

// New builds a new p2p service
func New(ctx context.Context, cancel context.CancelFunc, ch chan interface{}, cfg *clienttypes.ConfigClient, c clienttypes.InterfaceChains) (*P2P, error) {
	// 1. check inputs
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if cfg.P2P == nil {
		return nil, errors.New("nil p2p config")
	}
	if c == nil {
		return nil, ErrNoChainService
	}

	// 1. parse the public key for the pid
	pid, err := libpeer.IDFromPublicKey(cfg.P2P.Pub)
	if err != nil {
		logger.Errorf("[p2p] err generating peer id from public key\n%v", err)
		return nil, err
	}

	// 2. build the peerstore and save the keys
	// TODO: pass in pstore?
	// TODO: this is being built in peers, too!
	ps := pstoremem.NewPeerstore()
	if err := ps.AddPrivKey(pid, cfg.P2P.Priv); err != nil {
		logger.Errorf("[p2p] err adding private keey to peer store\n%v", err)
		return nil, err
	}
	if err := ps.AddPubKey(pid, cfg.P2P.Pub); err != nil {
		logger.Errorf("[p2p] err adding public key to peer store\n%v", err)
		return nil, err
	}

	// 3. build the swarm network
	// TODO ...
	swarmNet := swarm.NewSwarm(cfg.P2P.Context, pid, ps, nil)
	tcpTransport := tcp.NewTCPTransport(genUpgrader(swarmNet))
	if err := swarmNet.AddTransport(tcpTransport); err != nil {
		logger.Errorf("[p2p] err adding transport to swarmnet\n%v", err)
		return nil, err
	}
	bNode := bhost.New(swarmNet)

	// 4. build the dht
	// TODO ...
	dhtSvc, err := dht.New(cfg.P2P.Context, bNode)
	if err != nil {
		logger.Errorf("[p2p] err building dht service\n%v", err)
		return nil, err
	}
	if err := dhtSvc.Bootstrap(cfg.P2P.Context); err != nil {
		logger.Errorf("[p2p] err bootstrapping dht\n%v", err)
		return nil, err
	}

	// 5. build the host
	newNode := rhost.Wrap(bNode, dhtSvc)
	for i, addr := range newNode.Addrs() {
		logger.Infof("[p2p] %d: %s/ipfs/%s\n", i, addr, newNode.ID().Pretty())
	}

	// 6. build the discovery service
	// TODO ...
	// note: https://libp2p.io/implementations/#discovery
	// note: use https://github.com/libp2p/go-libp2p/blob/master/p2p/discovery/mdns.go rather than whyrusleeping
	discoverySvc, err := discovery.NewMdnsService(cfg.P2P.Context, newNode, time.Second, defaults.Defaults.Name)
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

	snc, err := sync.New(ctx, cfg, c)
	if err != nil {
		logger.Errorf("[p2p] err creating syncer\n%v", err)
		return nil, err
	}

	prs, err := peers.New(cfg, c, newNode)
	if err != nil {
		logger.Errorf("[p2p] err creating new peers\n%v", err)
		return nil, err
	}

	p := &P2P{
		state: &clienttypes.State{
			Chain:  c,
			Config: cfg,
			// TODO ...
			SyncState: &clienttypes.SyncState{},
			Peers:     prs,
			Host:      newNode,
		},
		cfg:       cfg,
		ctx:       ctx,
		ch:        ch,
		sync:      snc,
		cancel:    cancel,
		dialQueue: make(map[string]*clienttypes.QueuedPeer),
		handlers:  make(map[p2ptypes.EventEnum]clienttypes.EventCallback),
	}

	nb := &inet.NotifyBundle{
		ConnectedF:    p.onConn,
		OpenedStreamF: p.onStream,
	}
	newNode.Network().Notify(nb)

	return p, nil
}

// IsStarted ...
func (p *P2P) IsStarted() bool {
	// TODO: best practice for determining if server is started?
	if p.state == nil || p.state.Host == nil || len(p.state.Host.Addrs()) == 0 {
		return false
	}

	return true
}

// GetNumPeers ...
func (p *P2P) GetNumPeers() (uint, error) {
	if p.state == nil {
		return 0, ErrUninitializedService
	}
	if p.state.Peers == nil {
		return 0, ErrNoPeersService
	}

	return p.state.Peers.Count()
}

// On handles messages
func (p *P2P) On(event p2ptypes.EventEnum, cb clienttypes.EventCallback) {
	p.handlers[event] = cb

	return
}

// Start starts the p2p service
func (p *P2P) Start() error {
	if p.cfg == nil {
		return errors.New("nil config")
	}
	if p.state == nil {
		return errors.New("nil state")
	}
	if p.sync == nil {
		return errors.New("err nil sync")
	}

	listen, err := ma.NewMultiaddr(p.cfg.P2P.Address)
	if err != nil {
		logger.Errorf("[p2p] err parsing address\n%v", err)
		return err
	}
	if err := p.state.Host.Network().Listen(listen); err != nil {
		logger.Errorf("[p2p] err adding swarm listen addr %v to swarmnet\n%v", listen, err)
		return err
	}

	if err = p.handleProtocol(); err != nil {
		return err
	}
	if err = p.handlePing(); err != nil {
		return err
	}
	go p.dialPeers(nil)

	p.handleEvent(p2ptypes.Started)

	return nil
}

// Stop stops the p2p service
func (p *P2P) Stop() error {
	if p.state == nil {
		return ErrUninitializedService
	}
	if p.state.Host == nil {
		return ErrNoHost
	}

	if p.cancel != nil {
		p.cancel()
	}

	p.handleEvent(p2ptypes.Stopped)

	return p.state.Host.Close()
}

// Cfg returns the cfg
func (p *P2P) Cfg() clienttypes.ConfigClient {
	// TODO: return err?
	if p.cfg == nil {
		return clienttypes.ConfigClient{}
	}

	return *p.cfg
}

// GetSyncer ...
func (p *P2P) GetSyncer() (clienttypes.InterfaceSync, error) {
	if p.sync == nil {
		return nil, errors.New("nil sync")
	}

	return p.sync, nil
}

func (p *P2P) onConn(network inet.Network, conn inet.Conn) {
	logger.Infof("[p2p] peer did connect\nid %v peerAddr %v", conn.RemotePeer().Pretty(), conn.RemoteMultiaddr())

	p.addAddr(conn)
}

func (p *P2P) onStream(network inet.Network, stream inet.Stream) {
	if network == nil || stream == nil {
		logger.Errorf("[p2p] network or stream is nil")
		return
	}

	switch stream.Protocol() {
	case protocol.ID(defaults.Defaults.ProtocolPing), protocol.ID(defaults.Defaults.ProtocolDot):
		{
			logger.Info("[p2p] new stream was opened")
			return
		}

	default:
		{
			if p.state == nil {
				logger.Error("[p2p] nil state")
				return
			}
			if p.state.Peers == nil {
				logger.Error("[p2p] nil peers")
				return
			}

			kp, err := p.state.Peers.GetFromID(stream.Conn().RemotePeer())
			if err != nil {
				logger.Errorf("[p2p] err getting known peer from id\n%v", err)
				return
			}
			if kp == nil || kp.Peer == nil {
				logger.Error("[p2p] known peer is nil")
				return
			}

			go func() {
				if err = kp.Peer.Receive(stream); err != nil {
					logger.Errorf("[p2p] err receiving stream\n%v", err)
					return
				}

				logger.Infof("[p2p] message received from %s", kp.Peer.GetShortID())
			}()
		}
	}
}

func (p *P2P) addAddr(conn inet.Conn) {
	if p.state == nil || p.state.Host == nil {
		logger.Warnf("[p2p] no state or host, cannot add peer %s", conn.RemoteMultiaddr())
		return
	}

	for _, peer := range p.state.Host.Peerstore().Peers() {
		if conn.RemotePeer() == peer {
			// note: we already have info on this peer
			logger.Warnf("[p2p] already have peer in our peerstore")
			return
		}
	}

	// note: we don't have this peer's info
	p.state.Host.Peerstore().AddAddr(conn.RemotePeer(), conn.RemoteMultiaddr(), peerstore.PermanentAddrTTL)
	logger.Infof("[p2p] added %s to our peerstore", conn.RemoteMultiaddr())

	if _, err := p.state.Host.Network().DialPeer(context.Background(), conn.RemotePeer()); err != nil {
		logger.Warnf("[p2p] err connecting to a peer\n%v", err)

		return
	}

	logger.Infof("[p2p] connected to %s", conn.RemoteMultiaddr())
}

func (p *P2P) handleProtocol() error {
	if p.state == nil {
		return errors.New("nil state")
	}
	if p.state.Host == nil {
		return errors.New("nil host")
	}

	// TODO: is this the correct method?
	p.state.Host.SetStreamHandler(protocol.ID(defaults.Defaults.ProtocolDot), p.protocolHandler)

	return nil
}

func (p *P2P) protocolHandler(stream inet.Stream) {
	defer stream.Close()
	if p.state == nil {
		logger.Error("nil state")
		return
	}
	if p.state.Peers == nil {
		logger.Error("nil peers")
		return
	}

	if stream == nil {
		logger.Error("[p2p] got nil stream")
		return
	}

	multiAddr := stream.Conn().RemoteMultiaddr()

	pinfo, err := peerstore.InfoFromP2pAddr(multiAddr)
	if err != nil {
		logger.Errorf("[p2p] err getting info from multiaddr %v\n%v", multiAddr, err)
		return
	}
	if pinfo == nil {
		logger.Error("nil pinfo")
		return
	}

	pr, err := p.state.Peers.Add(*pinfo)
	if err != nil {
		logger.Errorf("[p2p] err adding peer\n%v", err)
		return
	}
	if pr == nil || pr.Peer == nil {
		logger.Error("[p2p] known peer is null")
		return
	}

	// TODO: check if is connected?
	ok, err := pr.Peer.IsWritable()
	if err != nil {
		logger.Errorf("[p2p] err checking if peer is writable\n%v", err)
		return
	}
	if !ok {
		// TODO: is this right? It runs for ever. Pass context?
		// Just add to the queue?
		//go p.dialPeers(pr.Peer)
		p.dialQueue[pr.Peer.GetID()] = &clienttypes.QueuedPeer{
			Peer:     pr.Peer,
			NextDial: time.Now(),
		}
	}

	return
}

func (p *P2P) dialPeers(pr clienttypes.InterfacePeer) {
	if !p.IsStarted() {
		logger.Error("p2p host not started")
		return
	}

	if pr != nil {
		if _, ok := p.dialQueue[pr.GetID()]; !ok {
			p.dialQueue[pr.GetID()] = &clienttypes.QueuedPeer{
				Peer:     pr,
				NextDial: time.Now(),
			}
		}
	}

	var (
		now time.Time
		k   string
	)
	for {
		select {
		case <-p.ctx.Done():
			{
				logger.Info("[p2p] context canceled. Stopping dialPeers.")
				return
			}
		case <-time.After(time.Duration(defaults.Defaults.DialInterval)):
			{
				now = time.Now()
				var (
					err    error
					active bool
				)
				// TODO: use go routine?
				for k = range p.dialQueue {
					if p.dialQueue[k] == nil || p.dialQueue[k].NextDial.After(now) || p.dialQueue[k].Peer == nil {
						continue
					}
					active, err = p.dialQueue[k].Peer.IsActive()
					if err != nil || active {
						continue
					}

					p.dialQueue[k].NextDial = p.dialQueue[k].NextDial.Add(time.Duration(defaults.Defaults.DialBackoff))
					if err = p.dialPeer(p.dialQueue[k].Peer); err != nil {
						// TODO: nil check
						logger.Errorf("[p2p] err dialing peer id %s\n%v", p.dialQueue[k].Peer.GetID(), err)
					}
				}
			}
		}
	}
}

func (p *P2P) dialPeer(pr clienttypes.InterfacePeer) error {
	if pr == nil {
		return errors.New("cannot dial nil peer")
	}
	if !p.IsStarted() {
		return errors.New("p2p host not started")
	}

	var (
		conn inet.Conn
		err  error
	)
	// note: check for nil?
	conns := p.state.Host.Network().ConnsToPeer(pr.Cfg().Peer.ID)
	if conns == nil || len(conns) == 0 {
		logger.Infof("[p2p] dialing peer with id %s", pr.GetID())
		conn, err = p.state.Host.Network().DialPeer(context.Background(), pr.Cfg().Peer.ID)
		if err != nil {
			return err
		}
	} else {
		conn = conns[0]
	}

	_, err = pr.AddConnection(conn, true)

	return err
}

func (p *P2P) handlePing() error {
	if p.state == nil {
		return errors.New("nil state")
	}
	if p.state.Host == nil {
		return errors.New("nil host")
	}

	// TODO: is this the correct method?
	p.state.Host.SetStreamHandler(protocol.ID(defaults.Defaults.ProtocolPing), p.pingHandler)
	return nil
}

func (p *P2P) pingHandler(stream inet.Stream) {
	defer stream.Close()

	if p.state == nil {
		logger.Error("nil state")
		return
	}
	if p.state.Peers == nil {
		logger.Error("nil peers")
		return
	}

	if stream == nil {
		logger.Error("[p2p] got nil stream")
		return
	}

	// TODO: use read writer?
	//rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	// TODO: use io.Pipe()?
	r := bufio.NewReader(stream)
	w := bufio.NewWriter(stream)

	if _, err := io.Copy(w, r); err != nil {
		logger.Errorf("[p2p] err handling ping\n%v", err)
	}
}

func (p *P2P) pingPeer(pr clienttypes.InterfacePeer) error {
	var err error
	for {
		select {
		case <-p.ctx.Done():
			{
				logger.Info("[p2p] context canceled. Stopping pingPeer.")
				return nil
			}
		case <-time.After(time.Duration(defaults.Defaults.PingInterval)):
			{
				if err = p.sendPingToPeer(pr); err != nil {
					logger.Errorf("[p2p] err sending ping to peer with ID %v\n%v", pr.Cfg().Peer.ID, err)
				}
			}
		}
	}
}

func (p *P2P) sendPingToPeer(pr clienttypes.InterfacePeer) error {
	if pr == nil {
		return errors.New("cannot ping nil peer")
	}
	if p.state == nil {
		return errors.New("nil state")
	}
	if p.state.Host == nil {
		return errors.New("nil host")
	}

	// TODO: nil check
	stream, err := p.state.Host.NewStream(context.Background(), pr.Cfg().Peer.ID, protocol.ID(defaults.Defaults.ProtocolPing))
	if err != nil {
		// TODO: disconnect from peer?
		return err
	}
	// TODO: defer peer disconnect?
	defer stream.Close()

	if err = stream.SetDeadline(time.Now().Add(time.Duration(defaults.Defaults.PingTimeout))); err != nil {
		return err
	}

	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return err
	}

	// TODO: use read writer?
	// TODO: use io.Pipe?
	//rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	r := bufio.NewReader(stream)
	w := bufio.NewWriter(stream)

	if _, err = w.Write(b); err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}

	var (
		b2 []byte
		c  byte
		nb int
	)
	// TODO: use inet.AwaitEOF(stream)?
	for {
		c, err = r.ReadByte()
		if err == nil {
			b2[nb] = c
			nb++
			continue
		}
		if err == io.EOF {
			break
		}
		if nb >= 32 {
			break
		}

		return err
	}

	if !bytes.Equal(b, b2) {
		logger.Errorf("[p2p] ping err from %v\nexpected %v\nreceived %v", pr.Cfg().Peer.ID, b, b2)
		return errors.New("pong does not match ping")
	}

	return nil
}

func (p *P2P) handlePeerMessage(msg *clienttypes.OnMessage) error {
	if msg == nil {
		return errors.New("nil msg")
	}
	if msg.Message == nil {
		return errors.New("nil message")
	}

	h, err := handler.FromEnum(msg.Message.Kind())
	if err != nil {
		return err
	}
	if h == nil {
		return errors.New("nil handler")
	}
	if h.Type() != msg.Message.Kind() {
		return errors.New("wrong handler")
	}

	return h.Func(p, msg.Peer, msg.Message)
}

func (p *P2P) requestAny() {
	var (
		knownPeers []*clienttypes.KnownPeer
		err        error
		idx        int
	)
	for {
		select {
		case <-p.ctx.Done():
			{
				logger.Info("[p2p] context canceled. Stopping requestAny.")
				return
			}
		case <-time.After(time.Duration(defaults.Defaults.RequestInterval)):
			knownPeers, err = p.state.Peers.KnownPeers()
			if err != nil {
				logger.Errorf("[p2p] err requesting known peers\n%v", err)
				continue
			}

			for idx = range knownPeers {
				// TODO: check peer is connected?
				if err = p.sync.RequestBlocks(knownPeers[idx].Peer); err != nil {
					logger.Errorf("[p2p] err requesting block from peer\n%v", err)
					continue
				}
			}
		}
	}
}

func (p *P2P) onPeerDiscovery() {
	if p.state == nil {
		logger.Error("[p2p] err nil state")
		return
	}
	if p.state.Peers == nil {
		logger.Error("[p2p] err nil peers")
		return
	}

	p.state.Peers.On(peerstypes.Discovered, func(iface interface{}) (interface{}, error) {
		if iface == nil {
			return nil, errors.New("cannot add nil peer")
		}
		pr, ok := iface.(clienttypes.InterfacePeer)
		if !ok {
			return nil, errors.New("iface not peer type")
		}
		if pr == nil {
			return nil, errors.New("nil peer")
		}

		p.dialQueue[pr.GetID()] = &clienttypes.QueuedPeer{
			Peer:     pr,
			NextDial: time.Now(),
		}

		return nil, nil
	})
}

func (p *P2P) onPeerMessage() {
	if p.state == nil {
		logger.Error("[p2p] err nil state")
		return
	}
	if p.state.Peers == nil {
		logger.Error("[p2p] err nil peers")
		return
	}

	p.state.Peers.On(peerstypes.Message, func(iface interface{}) (interface{}, error) {
		if iface == nil {
			return nil, errors.New("cannot add nil peer")
		}
		msg, ok := iface.(*clienttypes.OnMessage)
		if !ok {
			return nil, errors.New("iface not message type")
		}
		if msg == nil {
			return nil, errors.New("nil msg")
		}

		return nil, p.handlePeerMessage(msg)
	})
}

func (p *P2P) handleEvent(event p2ptypes.EventEnum) {
	if event == nil {
		logger.Info("nil event")
		return
	}

	cb, ok := p.handlers[event]
	if !ok {
		logger.Infof("[p2p] no event for %s", event.String())
		return
	}

	iface, err := cb()
	logger.Infof("[p2p] handled event %s\nresults:\n%v\n%v", event.String(), iface, err)
	return
}

//// TODO ...
//func (p *P2P) announceBlock(hash *crypto.Blake2b256Hash, header []byte, body []byte) {
//return
//}
