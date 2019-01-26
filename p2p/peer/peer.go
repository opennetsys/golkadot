package peer

import (
	"math/big"

	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/message/status"
	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"

	pstore "github.com/libp2p/go-libp2p-peerstore"
	transport "github.com/libp2p/go-libp2p-transport"
)

// TODO ...
// note: ensure the struct implements the interface
var _ peertypes.InterfacePeer = (*Peer)(nil)

// New ...
func New(cfg *peertypes.Config, chn chain.Interface, pInfo pstore.PeerInfo) (*Peer, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if chn == nil {
		return nil, ErrNoChain
	}

	m := make(map[pstore.PeerInfo]*KnownPeer)
	svc := &Peer{
		Chain:  chn,
		Config: cfg,
		Map:    m,
	}

	return svc, nil
}

// AddConnection ...
func (p *Peer) AddConnection(conn transport.Conn, isWritable bool) (uint, error) {
	// TODO: check for chain nil, etc.
	// note: set first, and then increment?
	connID := s.nextConnId
	p.NextConnID++

	// TODO??? only make chan if isWritable?
	ch := make(chan interface{})
	p.Connections[connID] = &Connection{
		Connection: conn,
		Pushable:   ch,
	}

	// TODO ...
	//this._receive(connection, connId);

	if isWritable {
		// TODO ...
		//pull(pushable, connection);
		bn, err := s.Chain.GetBestBlocksNumber()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks number\n%v", err)
			return 0, err
		}
		bh, err := s.Chain.GetBestBlocksHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks hash\n%v", err)
			return 0, err
		}
		gh, err := s.Chain.GetGenesisHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain genesis hash\n%v", err)
			return 0, err
		}

		p.Send(&status.Status{
			Roles:       p.Config.Roles,
			BestNumber:  bn,
			BestHash:    bh,
			GenesisHash: gh,
			Version:     defaults.PROTOCOL_VERSION,
		})
	}

	return connID
}

// Disconnect disconnects from the peer
func (p *Peer) Disconnect() error {
	s.Connections = nil
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
	return p.BestHash != nil && len(p.BestHash) != 0 && p.isWritable(), nil
}

// GetNextID TODO
func (p *Peer) GetNextID() (uint, error) {
	// note: increment first and then return?
	p.NextID++
	return p.NextID, nil
}

// On defines the event handlers
func (p *Peer) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO
	return nil, nil
}

// Send is used to send the peer a message
func (p *Peer) Send(msg message.Interface) (bool, error) {
	encoded, err := msg.Encode()
	if err != nil {
		logger.Errorf("[peer] %v message encoding send error\n%v", p.ShortID, err)
		return false, err
	}

	length := varint.encode(len(encoded))
	// note: is []byte(length) correct?
	buffer := u8util.ToBuffer(u8util.Concat([]byte(length), encoded))

	logger.Infof("[peer] sending %v -> %v", p.ShortID, u8util.ToHex(encoded))

	pushables := s.pushables()
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
func (p *Peer) Cfg() Config {
	if p.Config == nil {
		return Config{}
	}

	return *p.Config
}

func (p *Peer) clearConnection(connID int) {
	delete(p.Connections, connId)

	logger.Infof("[peer] clearConnection %v %v %v", connIsDd, p.ShortID, p.isWritable())

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
