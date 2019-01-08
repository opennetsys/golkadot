package client

import (
	"math/big"
	"time"
)

// InformantDelay ...
var InformantDelay = 10000

// TODO: these are placeholders. need to implement

// InterfaceChain ...
type InterfaceChain interface{}

// InterfaceP2P ...
type InterfaceP2P interface{}

// InterfaceRPC ...
type InterfaceRPC interface{}

// InterfaceTelemetry ...
type InterfaceTelemetry interface{}

// ChainName ...
type ChainName struct{}

// DBConfig ...
type DBConfig struct{}

// DevConfig ...
type DevConfig struct {
	genBlocks bool
}

// P2PConfig ...
type P2PConfig struct{}

// RPCConfig ...
type RPCConfig struct{}

// RolesConfig ...
type RolesConfig struct{}

// TelemetryConfig ...
type TelemetryConfig struct{}

// WasmConfig ...
type WasmConfig struct{}

// Client ...
type Client struct {
	chain       InterfaceChain
	informantID interface{}
	p2p         InterfaceP2P
	rpc         InterfaceRPC
	telemetry   InterfaceTelemetry
	prevBest    big.Int
	prevTime    int64
}

// Config ...
type Config struct {
	chain     ChainName
	db        DBConfig
	dev       DevConfig
	p2p       P2PConfig
	rpc       RPCConfig
	roles     []string
	telemetry TelemetryConfig
	wasm      WasmConfig
}

// NewClient ..
func NewClient() *Client {
	return &Client{
		prevTime: time.Now().Unix(),
	}
}

// Start ...
func (c *Client) Start(config *Config) {
	// TODO: implement
	/*
		c.chain = NewChain(config)
		c.p2p = NewP2P(config, c.chain)
		c.rpc = NewRPC(config, c.chain)
		c.telemetry = NewTelemetry(config, c.chain)

		c.p2p.Start()
		c.rpc.Start()
		c.telemetry.Start()
	*/

	c.StartInformant()
}

// Stop ...
func (c *Client) Stop() {
	c.StopInformant()
}

// StartInformant ...
func (c *Client) StartInformant() {
	// TODO: implement
	//c.informantID = setInterval(c.runInformant, InformationDelay)

	if c.p2p == nil {
		return
	}

	// TODO: implement
	/*
		c.p2p.sync.on("imported", func () {
			if c.telemetry != nil {
				c.telemetry.BlockImported()
			}
		})
	*/
}

// StopInformant ...
func (c *Client) StopInformant() {
	if c.informantID != nil {
		// TODO: implement
		// clearInterval(c.informantID)
	}

	c.informantID = nil
}

// RunInformant ...
func (c *Client) RunInformant() {
	if c.chain == nil || c.p2p == nil || c.rpc == nil {
		c.StopInformant()
		return
	}

	// TODO: implement
	/*
			now := time.Now().Unix()
			elapsed := now - c.prevTime
			numPeers := c.p2p.GetPeersCount()
			bestHash := c.chain.Blocks.BestHash.Get()
			bestNumber := c.chain.Blocks.BestNumber.Get()
			status := c.p2p.Sync.Status()
			isSync := status == "Sync"
			hasBlocks := c.prevBest && c.prevBest.lt(bestNumber)
			var numBlocks big.Int
			if hasBlocks && c.prevBest.Cmp(big.NewInt(0)) > 0 {
				numBlocks = new(big.int).Sub(bestNumber, c.prevBest)
			} else {
				numBlocks = big.NewInt(1)
			}

			var newSpeedStr string
			if isSync {
				newSpeedStr = fmt.Sprintf("%dms/block", elapsed/numBlocks.Uint64())
			}
			var newBlocksStr string
			if hasBlocks && c.prevBest.Cmp(big.NewInt(0)) > 0 {
				newBlocksStr = fmt.Sprintf(", +%sblocks%s", numBlocks.String(), newSpeedStr)
			}
			var targetStr string
			if isSync {
				targetStr = fmt.Sprintf(", target #", c.p2p.Sync.BestSeen.String())
			}

			fmt.Printf("%s (%d peers)%s, current #%s, %s%s", status, numPeers, targetStr, bestNumber.String(), u8util.ToHex(bestHash, 48, true), newBlocksStr)

		c.prevBest = bestNumber
		c.prevTime = now
	*/

	if c.telemetry != nil {
		// TODO: implement
		// c.telemetry.intervalInfo(numPeers, status)
	}
}
