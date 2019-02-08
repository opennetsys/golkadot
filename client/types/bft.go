package clienttypes

import (
	"encoding/json"
	"errors"
	"math/big"

	handlertypes "github.com/opennetsys/godot/client/p2p/handler/types"
	"github.com/opennetsys/godot/common/bnutil"
	"github.com/opennetsys/godot/common/codec"
	"github.com/opennetsys/godot/common/u8util"
	"github.com/opennetsys/godot/logger"
)

// Kind ...
func (b *BFT) Kind() handlertypes.FuncEnum {
	return handlertypes.BFT
}

// Encode serializes the message into a bytes array
// TODO: fix...
func (b *BFT) Encode() ([]byte, error) {
	jsn, err := b.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msgBytes, err := codec.Encode(jsn)
	if err != nil {
		return nil, err
	}

	return u8util.Concat(bnutil.ToUint8Slice(big.NewInt(int64(handlertypes.BFT)), 8, true, false), msgBytes), nil
}

// Decode deserializes a bytes array into a message
// TODO: fix...
func (b *BFT) Decode(data []byte) error {
	if data == nil || len(data) == 0 {
		return errors.New("nil data")
	}

	bn := bnutil.ToBN(data[0], true)
	if bn == nil {
		return errors.New("nil kind")
	}

	if bn.Int64() != int64(handlertypes.BFT) {
		logger.Errorf("[bft] expected BFT, but received %v", bn.Int64())
		return errors.New("wrong kind")
	}

	return b.UnmarshalJSON(data[1:])
}

// MarshalJSON returns json
func (b *BFT) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Message)
}

// UnmarshalJSON converts json to a message
func (b *BFT) UnmarshalJSON(data []byte) error {
	var msg map[string]interface{}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	b.Message = msg
	return nil
}

// GetHeader ...
func (b *BFT) GetHeader() *Header {
	return nil
}
