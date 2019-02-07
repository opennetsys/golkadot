package clienttypes

import (
	"encoding/json"
	"errors"
	"math/big"

	handlertypes "github.com/c3systems/go-substrate/client/p2p/handler/types"
	"github.com/c3systems/go-substrate/common/bnutil"
	"github.com/c3systems/go-substrate/common/codec"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"
)

// TODO...

// Kind ...
func (b *BlockRequest) Kind() handlertypes.FuncEnum {
	return handlertypes.BlockRequest
}

// Encode serializes the message into a bytes array
func (b *BlockRequest) Encode() ([]byte, error) {
	jsn, err := b.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msgBytes, err := codec.Encode(jsn)
	if err != nil {
		return nil, err
	}

	return u8util.Concat(bnutil.ToUint8Slice(big.NewInt(int64(handlertypes.BlockRequest)), 8, true, false), msgBytes), nil
}

// Decode deserializes a bytes array into a message
func (b *BlockRequest) Decode(data []byte) error {
	if data == nil || len(data) == 0 {
		return errors.New("nil data")
	}

	bn := bnutil.ToBN(data[0], true)
	if bn == nil {
		return errors.New("nil kind")
	}

	if bn.Int64() != int64(handlertypes.BlockRequest) {
		logger.Errorf("[block request] expected Block Request, but received %v", bn.Int64())
		return errors.New("wrong kind")
	}

	return b.UnmarshalJSON(data[1:])
}

// MarshalJSON returns json
func (b *BlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Message)
}

// UnmarshalJSON converts json to a message
func (b *BlockRequest) UnmarshalJSON(data []byte) error {
	msg := new(BlockRequestMessage)
	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	b.Message = msg
	return nil
}

// GetHeader ...
func (b *BlockRequest) GetHeader() *Header {
	return nil
}
