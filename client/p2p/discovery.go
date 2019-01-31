package p2p

import (
	"context"

	"github.com/c3systems/go-substrate/logger"

	host "github.com/libp2p/go-libp2p-host"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

// TODO ...

// DiscoveryNotifee ...
// TODO: make this private?
type DiscoveryNotifee struct {
	h host.Host
}

// HandlePeerFound ...
func (n *DiscoveryNotifee) HandlePeerFound(pi peerstore.PeerInfo) {
	n.h.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
	if err := n.h.Connect(context.Background(), pi); err != nil {
		logger.Errorf("[node] found peer %s\nerr connecting %v", pi.Addrs, err)

		return
	}

	logger.Infof("[node] found peer %s\nadded to peerstore and connected", pi.Addrs)
}
