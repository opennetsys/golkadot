package peer

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"testing"
	"time"

	peertypes "github.com/opennetsys/golkadot/client/p2p/peer/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	mocktypes "github.com/opennetsys/golkadot/client/types/mock"
	"github.com/opennetsys/golkadot/common/u8util"

	"github.com/golang/mock/gomock"
	inet "github.com/libp2p/go-libp2p-net"
	libpeer "github.com/libp2p/go-libp2p-peer"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	protocol "github.com/libp2p/go-libp2p-protocol"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
)

type mockStream struct {
	//io.Reader
	//io.Writer

	//io.Closer
	*io.PipeReader
	*io.PipeWriter
}

func (m mockStream) Close() error {
	return nil
}

func (m mockStream) Reset() error {
	return nil
}

func (m mockStream) SetReadDeadline(t time.Time) error {
	return nil
}

func (m mockStream) SetDeadline(t time.Time) error {
	return nil
}

func (m mockStream) SetWriteDeadline(t time.Time) error {
	return nil
}

func (m mockStream) Protocol() protocol.ID {
	return protocol.ID("foo")
}

func (m mockStream) SetProtocol(p protocol.ID) {
	return
}

func (m mockStream) Stat() inet.Stat {
	return inet.Stat{}
}

func (m mockStream) Conn() inet.Conn {
	// note: is this the easiest way to get a mocked conn?
	mn := mocknet.New(context.Background())
	c, _ := mn.ConnectPeers(peer.ID("foo"), peer.ID("bar"))
	return c
}

func TestSetBest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockChains := mocktypes.NewMockInterfaceChains(ctrl)
	pr, err := New(&clienttypes.ConfigClient{}, mockChains, pstore.PeerInfo{})
	if err != nil {
		t.Fatalf("err generating peer\n%v", err)
	}

	for i, tt := range []struct {
		blockNumber *big.Int
		hash        []byte
	}{
		{
			big.NewInt(32),
			[]byte("foo"),
		},
		{
			big.NewInt(0),
			[]byte("bar"),
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if err = pr.SetBest(tt.blockNumber, tt.hash); err != nil {
				t.Error(err)
				return
			}

			if tt.blockNumber.String() != pr.BestNumber.String() || !bytes.Equal(tt.hash, pr.BestHash) {
				t.Errorf("expected block number %s, hash %s; received block number %s, hash %s", tt.blockNumber.String(), string(tt.hash), pr.BestNumber.String(), string(pr.BestHash))
			}
		})
	}
}

func TestReceive(t *testing.T) {
	received := clienttypes.BFT{}
	msg := clienttypes.BFT{
		Message: map[string]interface{}{
			"foo": 1,
		},
	}
	encoded, err := msg.Encode()
	if err != nil {
		t.Fatalf("err encoding message\n%v", err)
	}

	lengthBuf := make([]byte, binary.MaxVarintLen64)
	_ = binary.PutVarint(lengthBuf, int64(len(encoded)))
	buffer := u8util.Concat(lengthBuf, encoded)

	ctrl := gomock.NewController(t)
	mockChains := mocktypes.NewMockInterfaceChains(ctrl)
	pr, err := New(&clienttypes.ConfigClient{}, mockChains, pstore.PeerInfo{})
	if err != nil {
		t.Fatalf("err generating peer\n%v", err)
	}
	pr.On(peertypes.Message, func(iface interface{}) (interface{}, error) {
		t.Logf("message\n%v", iface)
		tmp, ok := iface.(*clienttypes.BFT)
		if !ok || tmp == nil {
			return nil, errors.New("wrong type")
		}

		received = *tmp
		return nil, nil
	})

	for i, tt := range []struct {
		input []byte
		toErr bool
	}{
		{
			buffer,
			false,
		},
		{
			[]byte("foo"),
			true,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			r, w := io.Pipe()
			if w == nil || r == nil {
				t.Errorf("received nil pipe %v %v", w, r)
				return
			}
			stream := &mockStream{r, w}
			go func() {
				defer w.Close()
				time.Sleep(1 * time.Second)
				if n, err := w.Write(tt.input); err != nil || n == 0 {
					t.Errorf("err writing to %d bytes stream\n%v", n, err)
					return
				}
			}()
			err = pr.Receive(stream)
			if tt.toErr {
				if err == nil {
					t.Error("expected error, received nil")
					return
				}
			} else {
				if err != nil {
					t.Errorf("err receiving stream\n%v", err)
					return
				}

				if len(msg.Message) != len(received.Message) {
					t.Errorf("expected len %d, received len %d", len(msg.Message), len(received.Message))
					return
				}

				for k, v := range msg.Message {
					val, ok := received.Message[k]
					if !ok {
						t.Errorf("no received value for key %v", k)
						return
					}

					// TODO: fix codec. We passed an int and got a float64.
					if v.(int) != int(val.(float64)) {
						t.Errorf("expected %v, received %v", v, val)
					}
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockChains := mocktypes.NewMockInterfaceChains(ctrl)
	id := libpeer.ID("foobarbazforbarbaz")
	pi := pstore.PeerInfo{
		ID: id,
	}
	pr, err := New(&clienttypes.ConfigClient{}, mockChains, pi)
	if err != nil {
		t.Fatalf("err generating peer\n%v", err)
	}

	t.Run("it stores the id and short id", func(t *testing.T) {
		if pr.GetID() != string(id) {
			t.Errorf("wrong id; expected %s, received %s", string(id), pr.GetID())
		}
		if pr.GetShortID() != "foobar..barbaz" {
			t.Errorf("wrong short id; expected %s, received %s", "foobar..barbaz", pr.GetShortID())
		}
	})
	t.Run("it stores the pinfo", func(t *testing.T) {
		pinfo := pr.GetPeerInfo()
		if !reflect.DeepEqual(pi, pinfo) {
			t.Errorf("wrong peer info; expected %v, received %v", pi, pinfo)
		}
	})
}

func TestGetNextID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockChains := mocktypes.NewMockInterfaceChains(ctrl)
	id := libpeer.ID("foobarbazforbarbaz")
	pi := pstore.PeerInfo{
		ID: id,
	}
	pr, err := New(&clienttypes.ConfigClient{}, mockChains, pi)
	if err != nil {
		t.Fatalf("err generating peer\n%v", err)
	}

	t.Run("it increments by 1", func(t *testing.T) {
		nextID := pr.GetNextID()
		if nextID != 1 {
			t.Fatalf("expected 1, received %v", nextID)
		}
	})

	t.Run("it increments by 1, again", func(t *testing.T) {
		nextID := pr.GetNextID()
		if nextID != 2 {
			t.Fatalf("expected 2, received %v", nextID)
		}
	})
}
