package clienttypes

import (
	"errors"

	handlertypes "github.com/opennetsys/golkadot/client/p2p/handler/types"
	"github.com/opennetsys/golkadot/logger"
)

// DecodeMessage ...
func DecodeMessage(data []byte) (InterfaceMessage, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("data is nil")
	}

	switch int(data[0]) {
	case int(handlertypes.Status):
		{
			ret := &Status{}
			err := ret.Decode(data)
			return ret, err

		}
	case int(handlertypes.BlockRequest):
		{
			ret := &BlockRequest{}
			err := ret.Decode(data)
			return ret, err

		}
	case int(handlertypes.BlockResponse):
		{
			ret := &BlockResponse{}
			err := ret.Decode(data)
			return ret, err

		}
	case int(handlertypes.BlockAnnounce):
		{
			ret := &BlockAnnounce{}
			err := ret.Decode(data)
			return ret, err

		}
	case int(handlertypes.Transactions):
		{
			ret := &Transactions{}
			err := ret.Decode(data)
			return ret, err

		}
	case int(handlertypes.BFT):
		{
			ret := &BFT{}
			err := ret.Decode(data)
			return ret, err
		}
	default:
		{
			logger.Errorf("[client types] received unknown message with id %d", int(data[0]))
			return nil, errors.New("unknown message")
		}
	}
}
