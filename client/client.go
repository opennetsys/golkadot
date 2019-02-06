package client

import (
	"fmt"
	"math/big"
	"time"

	clienttypes "github.com/opennetsys/go-substrate/client/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client/src/index.ts

// TODO: these are placeholders. need to implement in their respective package

// InformantDelay ...
var InformantDelay = 10000

// Client ...
type Client struct {
	Chain       clienttypes.InterfaceChains
	InformantID interface{}
	P2P         clienttypes.InterfaceP2P
	RPC         clienttypes.InterfaceRPC
	Telemetry   clienttypes.InterfaceTelemetry
	PrevBest    big.Int
	PrevTime    int64
}

// NewClient ..
func NewClient() *Client {
	return &Client{
		PrevTime: time.Now().Unix(),
	}
}

// Start ...
func (c *Client) Start(config *clienttypes.ConfigClient) {
	// TODO: implement
	/*
		c.chain = clientchain.NewChain(config)
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

	if c.P2P == nil {
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
	if c.InformantID != nil {
		// TODO: implement
		// clearInterval(c.informantID)
	}

	c.InformantID = nil
}

// RunInformant ...
func (c *Client) RunInformant() {
	if c.Chain == nil || c.P2P == nil || c.RPC == nil {
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

	if c.Telemetry != nil {
		// TODO: implement
		// c.telemetry.intervalInfo(numPeers, status)
	}
}

// ID ...
func ID() string {
	// TODO: dynamic
	version := "0.0.1"
	name := "go-substrate"
	var stability string
	isDevelopment := true
	if isDevelopment {
		stability = "development"
	} else {
		stability = "release"
	}

	return fmt.Sprintf("%s/%s-%s", name, version, stability)
}
