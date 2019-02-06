package handler

import (
	"errors"

	handlertypes "github.com/opennetsys/go-substrate/client/p2p/handler/types"
)

// FromEnum ...
func FromEnum(typ handlertypes.FuncEnum) (InterfaceHandler, error) {
	switch typ {
	case handlertypes.BFT:
		{
			return &BFTHandler{}, nil

		}
	case handlertypes.BlockAnnounce:
		{
			return &BlockAnnounceHandler{}, nil

		}
	case handlertypes.BlockRequest:
		{
			return &BlockRequestHandler{}, nil

		}
	case handlertypes.BlockResponse:
		{
			return &BlockResponseHandler{}, nil

		}
	case handlertypes.Status:
		{
			return &StatusHandler{}, nil

		}
	case handlertypes.Transactions:
		{
			return &TransactionsHandler{}, nil
		}
	default:
		{
			return nil, errors.New("no such handler function")
		}
	}
}
