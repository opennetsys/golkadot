package peer

import (
	"math/big"

	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/client"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/message"
	"github.com/c3systems/go-substrate/p2p/message/status"

	pstore "github.com/libp2p/go-libp2p-peerstore"
	transport "github.com/libp2p/go-libp2p-transport"
)

// TODO ...

// New ...
func New(cfg *client.Config, chn chain.Interface, pInfo pstore.PeerInfo) (*Service, error) {
	if cfg == nil {
		return nil, ErrNoConfig
	}
	if chn == nil {
		return nil, ErrNoChain
	}

	m := make(map[pstore.PeerInfo]*KnownPeer)
	svc := &Service{
		Chain:  chn,
		Config: cfg,
		Map:    m,
	}

	return svc, nil
}

// AddConnection ...
func (s *Service) AddConnection(conn transport.Conn, isWritable bool) (uint, error) {
	// TODO: check for chain nil, etc.
	// note: set first, and then increment?
	connID := s.nextConnId
	s.NextConnID++

	// TODO??? only make chan if isWritable?
	ch := make(chan interface{})
	s.Connections[connID] = &Connection{
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

		s.Send(&status.Status{
			Roles:       s.Config.Roles,
			BestNumber:  bn,
			BestHash:    bh,
			GenesisHash: gh,
			Version:     defaults.PROTOCOL_VERSION,
		})
	}

	return connID
}

// Disconnect disconnects from the peer
func (s *Service) Disconnect() error {
	s.Connections = nil
	// TODO ...
	//this.emit('disconnected');

	return nil
}

// IsActive returns whether the peer is active or not
func (s *Service) IsActive() (bool, error) {
	pushables, err := s.pushables()
	if err != nil {
		logger.Errorf("[peer] err getting pushables\n%v", err)
		return false, err
	}

	return len(pushables) != 0, nil
}

// IsWritable returns whether the peer is writable or not
func (s *Service) IsWritable() (bool, error) {
	return s.BestHash != nil && len(s.BestHash) != 0 && s.isWritable(), nil
}

// GetNextID TODO
func (s *Service) GetNextID() (uint, error) {
	// note: increment first and then return?
	s.NextID++
	return s.NextID, nil
}

// On defines the event handlers
func (s *Service) On(event EventEnum, cb EventCallback) (interface{}, error) {
	// TODO
	return nil, nil
}

// Send is used to send the peer a message
func (s *Service) Send(msg message.Interface) (bool, error) {
	encoded, err := msg.Encode()
	if err != nil {
		logger.Errorf("[peer] %v message encoding send error\n%v", s.ShortID, err)
		return false, err
	}

	length := varint.encode(len(encoded))
	// note: is []byte(length) correct?
	buffer := u8util.ToBuffer(u8util.Concat([]byte(length), encoded))

	logger.Infof("[peer] sending %v -> %v", s.ShortID, u8util.ToHex(encoded))

	pushables := s.pushables()
	for idx := range pushables {
		pushables[idx] <- buffer
	}

	return true, nil
}

// SetBest sets a new block
func (s *Service) SetBest(blockNumber *big.Int, hash []byte) error {
	// TODO: check for nils?
	s.BestHash = hash
	s.BestNumber = blockNumber

	return nil
}

func (s *Service) clearConnection(connID int) {
	delete(s.Connections, connId)

	logger.Infof("[peer] clearConnection %v %v %v", connIsDd, s.ShortID, s.isWritable())

	// TODO ...
	//if (!this.isWritable()) {
	//this.emit('disconnected');
	//}
}

func (s *Service) pushables() ([]chan<- interface{}, error) {
	// TODO...
	return nil, nil
}

// TODO ...
func (s *Service) receive(conn transport.Conn, connID int) (bool, error) {
	return false, nil
}
