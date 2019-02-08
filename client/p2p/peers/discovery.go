package peers

import (
	peerstypes "github.com/opennetsys/golkadot/client/p2p/peers/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/logger"

	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

// TODO ...

// DiscoveryNotifee ...
type DiscoveryNotifee struct {
	peers clienttypes.InterfacePeers
}

// HandlePeerFound ...
func (d *DiscoveryNotifee) HandlePeerFound(pi peerstore.PeerInfo) {
	_, err := d.peers.Get(pi)
	if err == ErrNoSuchPeer {
		kp, err := d.peers.Add(pi)
		if err != nil {
			logger.Errorf("[peers] err adding peer with info %v\n%v", pi, err)
			return
		}

		if err = d.peers.Log(peerstypes.Discovered, kp); err != nil {
			logger.Errorf("[peers] err logging discovered\n%v", err)
		}
		return
	}
	if err != nil {
		logger.Errorf("[peers] err getting known peer from pi %v\n%v", pi, err)
		return
	}
}
