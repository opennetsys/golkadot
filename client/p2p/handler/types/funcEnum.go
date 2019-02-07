package handler

import (
	"errors"
	"strings"
)

type funcEnum int

// note: do not change this order!
const (
	// Status ...
	Status funcEnum = iota
	// BlockRequest ...
	BlockRequest
	// BlockResponse ...
	BlockResponse
	// BlockAnnounce ...
	BlockAnnounce
	// Transactions ...
	Transactions
	// BFT ...
	BFT
	// Request ...
	//Request
	//// StateRequest ...
	//StateRequest
)

// ErrUnknownFunc is thrown when an unknown handler function is encountered.
var ErrUnknownFunc = errors.New("handler func: unknown")

// FuncEnum are the exported handler func enums
type FuncEnum interface {
	Type() funcEnum
	String() string
}

// Type returns the private handler func enum
func (f funcEnum) Type() funcEnum {
	return f
}

// AllFuncEnums returns all of the handler func enums
func AllFuncEnums() []FuncEnum {
	return []FuncEnum{
		BFT,
		BlockAnnounce,
		BlockRequest,
		BlockResponse,
		//Request,
		//StateRequest,
		Status,
		Transactions,
	}
}

// FuncEnumFromString parses a string to return the handler func enum
func FuncEnumFromString(s string) (FuncEnum, error) {
	switch strings.ToUpper(s) {
	case "BFT":
		return BFT, nil
	case "BLOCKANNOUNCE":
		return BlockAnnounce, nil
	case "BLOCKREQUEST":
		return BlockRequest, nil
	case "BLOCKRESPONSE":
		return BlockResponse, nil
	//case "REQUEST":
	//return Request, nil
	//case "STATEREQUEST":
	//return StateRequest, nil
	case "STATUS":
		return Status, nil
	case "TRANSACTIONS":
		return Transactions, nil
	default:
		return nil, ErrUnknownFunc
	}
}
