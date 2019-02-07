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

// Kind ...
func (t *Transactions) Kind() handlertypes.FuncEnum {
	return handlertypes.Transactions
}

// Encode serializes the message into a bytes array
// TODO: fix...
func (t *Transactions) Encode() ([]byte, error) {
	jsn, err := t.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msgBytes, err := codec.Encode(jsn)
	if err != nil {
		return nil, err
	}

	return u8util.Concat(bnutil.ToUint8Slice(big.NewInt(int64(handlertypes.Transactions)), 8, true, false), msgBytes), nil
}

// Decode deserializes a bytes array into a message
// TODO: fix...
func (t *Transactions) Decode(data []byte) error {
	if data == nil || len(data) == 0 {
		return errors.New("nil data")
	}

	bn := bnutil.ToBN(data[0], true)
	if bn == nil {
		return errors.New("nil kind")
	}

	if bn.Int64() != int64(handlertypes.Transactions) {
		logger.Errorf("[transactions] expected Transactions, but received %v", bn.Int64())
		return errors.New("wrong kind")
	}

	return t.UnmarshalJSON(data[1:])
}

// MarshalJSON returns json
func (t *Transactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Message)
}

// UnmarshalJSON converts json to a message
func (t *Transactions) UnmarshalJSON(data []byte) error {
	msg := new(TransactionsMessage)
	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	t.Message = msg
	return nil
}

// GetHeader ...
func (t *Transactions) GetHeader() *Header {
	return nil
}
