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

// TODO

// Kind ...
func (s *Status) Kind() handlertypes.FuncEnum {
	return handlertypes.Status
}

// Encode serializes the message into a bytes array
// TODO: fix...
func (s *Status) Encode() ([]byte, error) {
	jsn, err := s.MarshalJSON()
	if err != nil {
		return nil, err
	}

	msgBytes, err := codec.Encode(jsn)
	if err != nil {
		return nil, err
	}

	return u8util.Concat(bnutil.ToUint8Slice(big.NewInt(int64(handlertypes.Status)), 8, true, false), msgBytes), nil
}

// Decode deserializes a bytes array into a message
// TODO: fix...
func (s *Status) Decode(data []byte) error {
	if data == nil || len(data) == 0 {
		return errors.New("nil data")
	}

	bn := bnutil.ToBN(data[0], true)
	if bn == nil {
		return errors.New("nil kind")
	}

	if bn.Int64() != int64(handlertypes.Status) {
		logger.Errorf("[bft] expected Status, but received %v", bn.Int64())
		return errors.New("wrong kind")
	}

	return s.UnmarshalJSON(data[1:])
}

// MarshalJSON returns json
func (s *Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Message)
}

// UnmarshalJSON converts json to a message
func (s *Status) UnmarshalJSON(data []byte) error {
	msg := new(StatusMessage)
	if err := json.Unmarshal(data, msg); err != nil {
		return err
	}

	s.Message = msg
	return nil
}

// GetHeader ...
func (s *Status) GetHeader() *Header {
	return nil
}
