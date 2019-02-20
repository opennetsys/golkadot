package p2p

import (
	"context"

	host "github.com/libp2p/go-libp2p-host"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/opennetsys/golkadot/client/p2p/peers"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/logger"
)

// TODO ...

type discoveryNotifee struct {
	host    host.Host
	peers   clienttypes.InterfacePeers
	context context.Context
}

// HandlePeerFound ...
func (d *discoveryNotifee) HandlePeerFound(pi peerstore.PeerInfo) {
	kp, err := d.peers.Get(pi)
	if err != nil {
		if err == peers.ErrNoSuchPeer {
			logger.Infof("[p2p] found new peer %s; adding and connecting", pi.ID)
			if err = d.connect(pi); err != nil {
				logger.Errorf("[p2p] err connecting to peer %s\n%v", pi.ID, err)
				return
			}
		}

		logger.Errorf("[p2p] err getting peer %s from peerstore\n%v", pi.ID, err)
		return
	}

	if kp == nil {
		logger.Warnf("[p2p] nil peer for %s; re-adding", pi.ID)
		if err = d.connect(pi); err != nil {
			logger.Errorf("[p2p] err connecting to peer %s\n%v", pi.ID, err)
			return
		}

		return
	}

	// TODO ...
	//if !kp.IsConnected {
	//logger.Warnf("[p2p] peer %s not connected; re-connecting", pi.ID)
	//if err = d.connect(pi); err != nil {
	//logger.Errorf("[p2p] err connecting to peer %s\n%v", pi.ID, err)
	//return
	//}

	//return
	//}
}

func (d *discoveryNotifee) connect(pi peerstore.PeerInfo) error {
	d.host.Peerstore().AddAddrs(pi.ID, pi.Addrs, peerstore.PermanentAddrTTL)
	if err := d.host.Connect(d.context, pi); err != nil {
		logger.Errorf("[p2p] found peer %v but failed to connect: %v", pi.Addrs, err)

		return err
	}

	logger.Infof("[p2p] found peer %s\nadded to peerstore and connected", pi.Addrs)
	return nil
}
