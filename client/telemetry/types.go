package telemetry

import (
	synctypes "github.com/opennetsys/golkadot/client/p2p/sync/types"
)

// Config ...
type Config struct {
	Name string
	URL  string
}

// InterfaceTelemetry ...
type InterfaceTelemetry interface {
	BlockImported()
	IntervalInfo(peers int, status synctypes.StatusEnum)
	Start()
}
