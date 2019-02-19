package p2p

import (
	"context"

	"github.com/opennetsys/golkadot/logger"

	host "github.com/libp2p/go-libp2p-host"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

// TODO ...

// DiscoveryNotifee ...
// TODO: make this private?
type DiscoveryNotifee struct {
	host host.Host
}

// HandlePeerFound ...
func (d *DiscoveryNotifee) HandlePeerFound(pi peerstore.PeerInfo) {
	d.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
	if err := d.host.Connect(context.Background(), pi); err != nil {
		logger.Errorf("[node] found peer %v but failed to connect: %v", pi.Addrs, err)

		return
	}

	logger.Infof("[node] found peer %s\nadded to peerstore and connected", pi.Addrs)
}
