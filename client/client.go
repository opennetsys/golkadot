package client

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	clientchain "github.com/opennetsys/golkadot/client/chain"
	p2p "github.com/opennetsys/golkadot/client/p2p"
	clienttypes "github.com/opennetsys/golkadot/client/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client/src/index.ts

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
	var err error
	c.Chain, err = clientchain.NewChain(config)
	if err != nil {
		log.Fatal(err)
	}
	c.P2P, err = p2p.NewP2P(context.Background(), nil, nil, config, c.Chain)
	if err != nil {
		log.Fatal(err)
	}
	//c.RPC = NewRPC(config, c.Chain)
	//c.Telemetry = NewTelemetry(config, c.Chain)

	c.P2P.Start()
	//c.RPC.Start()
	//c.Telemetry.Start()

	c.StartInformant()
}

// Stop ...
func (c *Client) Stop() {
	c.StopInformant()
}

// StartInformant ...
func (c *Client) StartInformant() {
	// TODO: implement
	//c.InformantID = setInterval(c.RunInformant, InformationDelay)

	if c.P2P == nil {
		return
	}

	// TODO: implement
	/*
		c.P2P.Sync.on("imported", func () {
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
		// c.Telemetry.intervalInfo(numPeers, status)
	}
}

// ID ...
func ID() string {
	// TODO: dynamic
	version := "0.0.1"
	name := "golkadot"
	var stability string
	isDevelopment := true
	if isDevelopment {
		stability = "development"
	} else {
		stability = "release"
	}

	return fmt.Sprintf("%s/%s-%s", name, version, stability)
}
