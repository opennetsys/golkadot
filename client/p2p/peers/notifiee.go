package peers

import (
	peerstypes "github.com/opennetsys/godot/client/p2p/peers/types"
	clienttypes "github.com/opennetsys/godot/client/types"
	"github.com/opennetsys/godot/logger"

	inet "github.com/libp2p/go-libp2p-net"
	ma "github.com/multiformats/go-multiaddr"
)

// Notifiee ...
// TODO: make private?
type Notifiee struct {
	peers clienttypes.InterfacePeers
}

// Listen ...
func (n *Notifiee) Listen(net inet.Network, addr ma.Multiaddr) {}

// ListenClose ...
func (n *Notifiee) ListenClose(net inet.Network, addr ma.Multiaddr) {}

// Connected ...
func (n *Notifiee) Connected(net inet.Network, conn inet.Conn) {
	if net == nil || conn == nil {
		logger.Info("[peers] net and conn required")
		return
	}

	pID := conn.RemotePeer()
	ps := net.Peerstore()
	if ps == nil {
		logger.Error("[peers] nil pstore")
		return
	}

	pi := ps.PeerInfo(pID)

	kp, err := n.peers.Get(pi)
	if err != nil {
		logger.Errorf("[peers] err getting known peer from pi %v\n%v", pi, err)
		return
	}
	if kp == nil {
		logger.Info("[peers] nil known peer")
		return
	}

	if !kp.IsConnected {
		kp.IsConnected = true
		if err = n.peers.Log(peerstypes.Connected, kp.Peer); err != nil {
			logger.Errorf("[peers] err logging connected\n%v", err)
		}
	}
}

// Disconnected ...
func (n *Notifiee) Disconnected(net inet.Network, conn inet.Conn) {
	if net == nil || conn == nil {
		logger.Info("[peers] net and conn required")
		return
	}

	pID := conn.RemotePeer()
	ps := net.Peerstore()
	if ps == nil {
		logger.Error("[peers] nil pstore")
		return
	}

	pi := ps.PeerInfo(pID)

	kp, err := n.peers.Get(pi)
	if err != nil {
		logger.Errorf("[peers] err getting known peer with pi %v\n%v", pi, err)
		return
	}

	if !kp.IsConnected {
		return
	}

	kp.IsConnected = false

	if err = n.peers.Log(peerstypes.Disconnected, kp.Peer); err != nil {
		logger.Errorf("[peers] err logging disconected\n%v", err)
	}
}

// OpenedStream ...
func (n *Notifiee) OpenedStream(net inet.Network, stream inet.Stream) {

}

// ClosedStream ...
func (n *Notifiee) ClosedStream(net inet.Network, stream inet.Stream) {

}
