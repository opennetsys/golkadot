package peer

import (
	"encoding/binary"
	"errors"
	"log"
	"math/big"

	"github.com/c3systems/go-substrate/client/p2p/defaults"
	peertypes "github.com/c3systems/go-substrate/client/p2p/peer/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"

	inet "github.com/libp2p/go-libp2p-net"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	transport "github.com/libp2p/go-libp2p-transport"
)

// TODO ...
// note: ensure the struct implements the interface
var _ clienttypes.InterfacePeer = (*Peer)(nil)

// New ...
func New(cfg *clienttypes.ConfigClient, chn clienttypes.InterfaceChains, pInfo pstore.PeerInfo) (*Peer, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if chn == nil {
		return nil, ErrNoChain
	}

	m := make(map[libpeer.ID]*clienttypes.KnownPeer)
	svc := &Peer{
		Chain:  chn,
		Config: cfg,
		Map:    m,
	}

	return svc, nil
}

// AddConnection ...
func (p *Peer) AddConnection(conn inet.Conn, isWritable bool) (uint, error) {
	// TODO: check for chain nil, etc.
	// note: set first, and then increment?
	connID := p.NextConnID
	p.NextConnID++

	// TODO??? only make chan if isWritable?
	ch := make(chan interface{})
	p.Connections[int(connID)] = &clienttypes.Connection{
		Connection: conn,
		Pushable:   ch,
	}

	// TODO ...
	//this._receive(connection, connId);

	if isWritable {
		// TODO ...
		//pull(pushable, connection);
		bn, err := p.Chain.GetBestBlocksNumber()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks number\n%v", err)
			return 0, err
		}
		bh, err := p.Chain.GetBestBlocksHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks hash\n%v", err)
			return 0, err
		}
		gh, err := p.Chain.GetGenesisHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain genesis hash\n%v", err)
			return 0, err
		}

		ok, err := p.Send(&clienttypes.Status{
			Roles:       p.Config.Roles,
			BestNumber:  bn,
			BestHash:    bh,
			GenesisHash: gh,
			Version:     defaults.Defaults.ProtocolVersion,
		})
		if err != nil {
			logger.Errorf("[peer] err sending message\n%v", err)
			return 0, err
		}
		if !ok {
			logger.Print("[peer] message send returned not ok")
			return 0, errors.New("could not send message")
		}
	}

	return connID, nil
}

// Disconnect disconnects from the peer
func (p *Peer) Disconnect() error {
	p.Connections = nil
	// TODO ...
	//this.emit('disconnected');

	return nil
}

// IsActive returns whether the peer is active or not
func (p *Peer) IsActive() (bool, error) {
	pushables, err := p.pushables()
	if err != nil {
		logger.Errorf("[peer] err getting pushables\n%v", err)
		return false, err
	}

	return len(pushables) != 0, nil
}

// IsWritable returns whether the peer is writable or not
func (p *Peer) IsWritable() (bool, error) {
	pushables, err := p.pushables()
	if err != nil {
		return false, err
	}

	return len(pushables) != 0, nil
}

// GetNextID TODO
func (p *Peer) GetNextID() (uint, error) {
	// note: increment first and then return?
	p.NextID++
	return p.NextID, nil
}

// On defines the event handlers
func (p *Peer) On(event peertypes.EventEnum, cb clienttypes.PeerEventCallback) {
	// TODO
	return
}

// Send is used to send the peer a message
func (p *Peer) Send(msg clienttypes.InterfaceMessage) (bool, error) {
	encoded, err := msg.Encode()
	if err != nil {
		logger.Errorf("[peer] %v message encoding send error\n%v", p.ShortID, err)
		return false, err
	}

	lengthBuf := make([]byte, binary.MaxVarintLen64)
	// TODO: correct to be ignoring bytes written, here?
	_ = binary.PutUvarint(lengthBuf, uint64(len(encoded)))
	buffer, err := u8util.ToBuffer(u8util.Concat(lengthBuf, encoded))
	if err != nil {
		return false, err
	}

	logger.Infof("[peer] sending %v -> %v", p.ShortID, u8util.ToHex(encoded, -1, true))

	pushables, err := p.pushables()
	if err != nil {
		return false, err
	}

	for idx := range pushables {
		pushables[idx] <- buffer
	}

	return true, nil
}

// SetBest sets a new block
func (p *Peer) SetBest(blockNumber *big.Int, hash []byte) error {
	// TODO: check for nils?
	p.BestHash = hash
	p.BestNumber = blockNumber

	return nil
}

// Cfg ...
func (p *Peer) Cfg() clienttypes.ConfigClient {
	if p.Config == nil {
		return clienttypes.ConfigClient{}
	}

	return *p.Config
}

func (p *Peer) clearConnection(connID int) {
	delete(p.Connections, connID)

	isWriteable, err := p.IsWritable()
	// TODO: fix!
	log.Println(err)

	logger.Infof("[peer] clearConnection %v %v %v", connID, p.ShortID, isWriteable)

	// TODO ...
	//if (!this.isWritable()) {
	//this.emit('disconnected');
	//}
}

func (p *Peer) pushables() ([]chan<- interface{}, error) {
	// TODO...
	return nil, nil
}

// TODO ...
func (p *Peer) receive(conn transport.Conn, connID int) (bool, error) {
	return false, nil
}

// GetID ...
func (p *Peer) GetID() string {
	return p.ID
}
