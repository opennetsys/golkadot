package peers

import (
	"context"
	"crypto/rand"
	"testing"

	clienttypes "github.com/opennetsys/godot/client/types"
	mocktypes "github.com/opennetsys/godot/client/types/mock"

	"github.com/golang/mock/gomock"
	ic "github.com/libp2p/go-libp2p-crypto"
	libpeer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
)

func TestPeers(t *testing.T) {
	mn := mocknet.New(context.Background())
	mh, err := mn.GenPeer()
	if err != nil {
		t.Fatalf("err generating host\n%v", err)
	}

	ctrl := gomock.NewController(t)
	mockChains := mocktypes.NewMockInterfaceChains(ctrl)
	id := libpeer.ID("foobarbazforbarbaz")
	pi := pstore.PeerInfo{
		ID: id,
	}

	priv, pub, err := ic.GenerateKeyPairWithReader(ic.RSA, 256, rand.Reader)
	if err != nil {
		t.Fatalf("err generating key pair\n%v", err)
	}
	prsID, err := libpeer.IDFromPrivateKey(priv)
	if err != nil {
		t.Fatalf("err generating id from priv key\n%v", err)
	}
	prs, err := New(&clienttypes.ConfigClient{
		Peers: &clienttypes.ConfigPeers{
			Priv: priv,
			Pub:  pub,
			ID:   prsID,
		},
		P2P: &clienttypes.ConfigP2P{
			Context: context.Background(),
		},
	}, mockChains, mh)
	if err != nil {
		t.Fatalf("err building peers\n%v", err)
	}

	t.Run("it adds a peer", func(t *testing.T) {
		count, err := prs.CountAll()
		if err != nil {
			t.Errorf("err counting peers\n%v", err)
			return
		}
		if count != 0 {
			t.Errorf("expected 0, received %v", count)
			return
		}
		count, err = prs.Count()
		if err != nil {
			t.Errorf("err counting peers\n%v", err)
			return
		}
		if count != 0 {
			t.Errorf("expected 0, received %v", count)
			return
		}
		addedPeers, err := prs.KnownPeers()
		if err != nil {
			t.Errorf("err getting peers\n%v", err)
			return
		}
		if len(addedPeers) != 0 {
			t.Errorf("expected 0, received %v", len(addedPeers))
			return
		}

		kp, err := prs.Add(pi)
		if err != nil {
			t.Errorf("err adding peer\n%v", err)
			return
		}
		if kp == nil {
			t.Error("known peer is nil")
			return
		}

		count, err = prs.CountAll()
		if err != nil {
			t.Errorf("err counting peers\n%v", err)
			return
		}
		if count != 1 {
			t.Errorf("expected 1, received %v", count)
			return
		}

		kp1, err := prs.Get(pi)
		if err != nil {
			t.Errorf("err getting peer\n%v", err)
			return
		}
		if kp != kp1 {
			t.Errorf("expected %v, received %v", kp, kp1)
			return
		}
	})
	t.Run("it does not re-add a peer", func(t *testing.T) {
		kp, err := prs.Add(pi)
		if err != nil {
			t.Errorf("err adding peer\n%v", err)
			return
		}
		if kp == nil {
			t.Error("known peer is nil")
			return
		}
		kp, err = prs.Add(pi)
		if err != nil {
			t.Errorf("err adding peer\n%v", err)
			return
		}
		if kp == nil {
			t.Error("known peer is nil")
			return
		}

		count, err := prs.CountAll()
		if err != nil {
			t.Errorf("err counting peers\n%v", err)
			return
		}
		if count != 1 {
			t.Errorf("expected 1, received %v", count)
			return
		}
	})
	t.Run("it does not return a peer for info it does not have", func(t *testing.T) {
		_, err := prs.Add(pi)
		if err != nil {
			t.Errorf("err adding peer\n%v", err)
			return
		}

		if _, err = prs.Get(pstore.PeerInfo{}); err == nil || err != ErrNoSuchPeer {
			t.Error("expected err but got nil")
			return
		}
	})
	// TODO: test peer receiving a message...
}
