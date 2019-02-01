package clienttypes

import (
	"errors"

	"github.com/c3systems/go-substrate/logger"
)

// DecodeMessage ...
func DecodeMessage(data []byte) (InterfaceMessage, error) {
	if data == nil || len(data) == 0 {
		return nil, errors.New("data is nil")
	}

	switch data[0] {
	case 0:
		{
			ret := &BFT{}
			err := ret.Decode(data)
			return ret, err
		}
	case 1:
		{
			ret := &BlockAnnounce{}
			err := ret.Decode(data)
			return ret, err

		}
	case 2:
		{
			ret := &BlockRequest{}
			err := ret.Decode(data)
			return ret, err

		}
	case 3:
		{
			ret := &BlockResponse{}
			err := ret.Decode(data)
			return ret, err

		}
	case 4:
		{
			ret := &Status{}
			err := ret.Decode(data)
			return ret, err

		}
	case 5:
		{
			ret := &Transactions{}
			err := ret.Decode(data)
			return ret, err

		}
	default:
		{
			logger.Errorf("[client types] received message with id %d", int(data[0]))
			return nil, errors.New("unknown message")
		}
	}
}
