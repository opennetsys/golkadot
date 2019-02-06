package peer

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math/big"

	"github.com/c3systems/go-substrate/client/p2p/defaults"
	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
	peertypes "github.com/c3systems/go-substrate/client/p2p/peer/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/common/stringutil"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"

	inet "github.com/libp2p/go-libp2p-net"
	pstore "github.com/libp2p/go-libp2p-peerstore"
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

	p := &Peer{
		chain:       chn,
		config:      cfg,
		connections: make(map[uint]inet.Conn),
		peerInfo:    pInfo,
		// TODO: base58?
		id:       string(pInfo.ID),
		shortID:  stringutil.Shorten(string(pInfo.ID), 6),
		handlers: make(map[peertypes.EventEnum]clienttypes.PeerEventCallback),
	}

	return p, nil
}

// AddConnection ...
func (p *Peer) AddConnection(conn inet.Conn, isWritable bool) (uint, error) {
	// TODO: check for chain nil, etc.
	connID := p.nextConnID
	p.nextConnID++

	if isWritable {
		p.connections[connID] = conn
		//go p.Receive(conn, connID)
		//pull(pushable, connection);

		bn, err := p.chain.GetBestBlocksNumber()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks number\n%v", err)
			return 0, err
		}
		bh, err := p.chain.GetBestBlocksHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain best blocks hash\n%v", err)
			return 0, err
		}
		gh, err := p.chain.GetGenesisHash()
		if err != nil {
			logger.Errorf("[peer] err getting chain genesis hash\n%v", err)
			return 0, err
		}

		ok, err := p.Send(&clienttypes.Status{
			Roles:       p.config.Roles,
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
	p.connections = nil
	p.handleEvent(peertypes.Disconnected, nil)

	return nil
}

// IsActive returns whether the peer is active or not
func (p *Peer) IsActive() (bool, error) {
	isWritable, err := p.IsWritable()
	if err != nil {
		logger.Errorf("[peer] err getting isWritable\n%v", err)
		return false, err
	}

	return len(p.BestHash) != 0 && isWritable, nil
}

// IsWritable returns whether the peer is writable or not
func (p *Peer) IsWritable() (bool, error) {
	pushables, err := p.pushables()
	if err != nil {
		return false, err
	}

	return len(pushables) != 0, nil
}

// GetNextID ...
func (p *Peer) GetNextID() uint {
	p.nextID++
	return p.nextID
}

// On defines the event handlers
func (p *Peer) On(event peertypes.EventEnum, cb clienttypes.PeerEventCallback) {
	p.handlers[event] = cb

	return
}

// Send is used to send the peer a message
func (p *Peer) Send(msg clienttypes.InterfaceMessage) (bool, error) {
	encoded, err := msg.Encode()
	if err != nil {
		logger.Errorf("[peer] %v message encoding send error\n%v", p.shortID, err)
		return false, err
	}

	lengthBuf := make([]byte, binary.MaxVarintLen64)
	// TODO: correct to be ignoring bytes written, here?
	_ = binary.PutVarint(lengthBuf, int64(len(encoded)))
	buffer := u8util.Concat(lengthBuf, encoded)

	logger.Infof("[peer] sending %v -> %v", p.shortID, u8util.ToHex(encoded, -1, true))

	pushables, err := p.pushables()
	if err != nil {
		return false, err
	}

	// TODO: use goroutines?
	var (
		ret    = true
		stream inet.Stream
		w      *bufio.Writer
	)
	for idx := range pushables {
		if pushables[idx] == nil {
			continue
		}

		// TODO: get stream or open a new one?
		stream, err = pushables[idx].NewStream()
		if err != nil {
			logger.Errorf("[peer] err getting new stream\n%v", err)
			ret = false
			continue
		}

		w = bufio.NewWriter(stream)
		if _, err = w.Write(buffer); err != nil {
			logger.Errorf("[peer] err writing message\n%v", err)
			if err = stream.Close(); err != nil {
				logger.Errorf("[peer] err closing stream\n%v", err)
			}
			ret = false
			continue
		}
		if err = w.Flush(); err != nil {
			logger.Errorf("[peer] err flushing message\n%v", err)
			if err = stream.Close(); err != nil {
				logger.Errorf("[peer] err closing stream\n%v", err)
			}
			ret = false
			continue
		}

		if err = stream.Close(); err != nil {
			logger.Errorf("[peer] err closing stream\n%v", err)
		}
	}

	return ret, nil
}

// Receive ...
func (p *Peer) Receive(stream inet.Stream) error {
	if stream == nil {
		return errors.New("nil stream")
	}
	// TODO: unhandled err...
	defer stream.Close()

	// TODO: limit the number of bytes that can be read?
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(stream)
	if err != nil && err != io.EOF {
		logger.Errorf("[peer] err reading from stream\n%v", err)
		return err
	}

	logger.Infof("[peer] read %v bytes", n)

	b := buf.Bytes()
	lengthBuf := make([]byte, binary.MaxVarintLen64)

	if len(b) <= len(lengthBuf) {
		logger.Errorf("[peer] have bytes len %d; need minimum %d", len(b), len(lengthBuf)+1)
		return errors.New("bytes length below minimum")
	}

	data := b[len(lengthBuf):]
	lengthBuf = b[:len(lengthBuf)]
	dataLen, err := binary.ReadVarint(bytes.NewBuffer(lengthBuf))
	if err != nil {
		logger.Errorf("[peer] err reading var int\n%v", err)
		return err
	}

	if dataLen != int64(len(data)) {
		logger.Errorf("[peer] expected data length %v, received %v", dataLen, len(data))
		return errors.New("incorrect bytes length")
	}

	msg, err := clienttypes.DecodeMessage(data)
	if err != nil && err != io.EOF {
		logger.Errorf("[peer] err decoding message\n%v", err)
		return err
	}

	p.handleEvent(peertypes.Message, msg)

	if msg.Kind() == handlertypes.Status {
		p.handleEvent(peertypes.Active, nil)
	}

	return nil
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
	if p.config == nil {
		return clienttypes.ConfigClient{}
	}

	return *p.config
}

// GetID ...
func (p *Peer) GetID() string {
	return p.id
}

// GetChain ...
func (p *Peer) GetChain() (clienttypes.InterfaceChains, error) {
	if p.chain == nil {
		return nil, errors.New("nil chain")
	}

	return p.chain, nil
}

// GetPeerInfo ...
func (p *Peer) GetPeerInfo() pstore.PeerInfo {
	return p.peerInfo
}

// GetShortID ...
func (p *Peer) GetShortID() string {
	return p.shortID
}

// GetBestNumber ...
func (p *Peer) GetBestNumber() *big.Int {
	return p.BestNumber
}

func (p *Peer) handleEvent(event peertypes.EventEnum, iface interface{}) {
	if event == nil {
		logger.Info("[peer] nil event")
		return
	}

	cb, ok := p.handlers[event]
	if !ok {
		logger.Infof("[peer] no event for %s", event.String())
		return
	}

	iface, err := cb(iface)
	logger.Infof("[peers] handled event %s\nresults:\n%v\n%v", event.String(), iface, err)
	return
}

func (p *Peer) clearConnection(connID uint) {
	// TODO: Dialer.ClosePeer?
	delete(p.connections, connID)

	isWriteable, err := p.IsWritable()
	if err != nil {
		logger.Errorf("[peer] err clearing connection\n%v", err)
		return
	}

	logger.Infof("[peer] clearConnection %v %v %v", connID, p.shortID, isWriteable)

	if !isWriteable {
		p.handleEvent(peertypes.Disconnected, nil)
	}
}

func (p *Peer) pushables() ([]inet.Conn, error) {
	var pushables []inet.Conn
	for k := range p.connections {
		pushables = append(pushables, p.connections[k])
	}

	return pushables, nil
}
