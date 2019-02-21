package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"

	ic "github.com/libp2p/go-libp2p-crypto"
	libpeer "github.com/libp2p/go-libp2p-peer"
	"github.com/opennetsys/golkadot/client"
	clientchain "github.com/opennetsys/golkadot/client/chain"
	clientdbtypes "github.com/opennetsys/golkadot/client/db/types"
	"github.com/opennetsys/golkadot/client/p2p"
	p2psync "github.com/opennetsys/golkadot/client/p2p/sync"
	"github.com/opennetsys/golkadot/client/rpc"
	"github.com/opennetsys/golkadot/client/telemetry"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/client/wasm"
	"github.com/opennetsys/golkadot/common/db"
	"github.com/opennetsys/golkadot/logger"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func setup() {
	var dbCompact bool
	var dbSnapshot bool
	var dbPath string
	var dbType string

	var p2pAddress string
	var p2pNodes []string
	var p2pPort uint
	var p2pMaxPeers uint
	var p2pNoBootnodes bool

	var rpcPath string
	var rpcTypes []string
	var rpcPort uint

	var telemetryName string
	var telemetryURL string

	var wasmHeapSize uint

	var chain string
	var role string
	var clientID string

	rootCmd = &cobra.Command{
		Use:   "start",
		Short: "Start node",
		Long:  `Starts the Polkadot node`,
		RunE: func(cmd *cobra.Command, args []string) error {
			priv, pub, err := ic.GenerateKeyPairWithReader(ic.RSA, 2048, rand.Reader)
			if err != nil {
				return err
			}

			pi, err := libpeer.IDFromPrivateKey(priv)
			if err != nil {
				return err
			}

			// TODO: implement
			config := &clienttypes.ConfigClient{
				DB: &clientdbtypes.Config{
					Compact:  dbCompact,
					IsTrieDB: true,
					Path:     dbPath,
					Snapshot: dbSnapshot,
					Type:     dbType,
				},
				Chain: "dev",
				Peer: &clienttypes.ConfigPeer{
					BestHash:   nil,
					BestNumber: nil,
					ID:         pi,
					PeerInfo:   nil,
					ShortID:    "",
				},
				Peers: &clienttypes.ConfigPeers{
					Priv: priv,
					Pub:  pub,
					ID:   pi,
				},
				RPC: &clienttypes.ConfigRPC{
					Host:          nil,
					SystemService: nil,
					StateService:  nil,
					ChainService:  nil,
					AuthorService: nil,
					ID:            nil,
				},
				Telemetry: &clienttypes.TelemetryConfig{},
				Wasm:      &clienttypes.WasmConfig{},
			}

			cl := client.NewClient()
			cl.Chain, err = clientchain.NewChain(config)
			if err != nil {
				return err
			}

			syncer, err := p2psync.New(context.Background(), cl.Chain)
			if err != nil {
				return err
			}

			config.P2P = &clienttypes.ConfigP2P{
				Address:     p2pAddress,
				ClientID:    clientID,
				MaxPeers:    p2pMaxPeers,
				Nodes:       p2pNodes,
				NoBootNodes: p2pNoBootnodes,
				Port:        p2pPort,
				Syncer:      syncer,
				Priv:        priv,
				Pub:         pub,
				Context:     context.Background(),
			}

			cl.Start(config)

			return nil
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&dbCompact, "db-compact", "", false, "Compact existing databases")
	rootCmd.PersistentFlags().BoolVarP(&dbSnapshot, "db-snapshot", "", false, "Create trie snapshot, drop & restore database from this")
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db-path", "", db.DefaultPath, "Sets the path for all storage operations")
	rootCmd.PersistentFlags().StringVarP(&dbType, "db-type", "", db.DefaultType, "The type of database storage to use")

	rootCmd.PersistentFlags().StringVarP(&p2pAddress, "p2p-address", "", p2p.DefaultAddress, "The interface to bind to (p2p-port > 0)")
	rootCmd.PersistentFlags().StringArrayVarP(&p2pNodes, "p2p-nodes", "", nil, "Reserved nodes to make initial connections to")
	rootCmd.PersistentFlags().UintVarP(&p2pMaxPeers, "p2p-max-peers", "", p2p.DefaultMaxPeers, "The maximum allowed peers")
	rootCmd.PersistentFlags().UintVarP(&p2pPort, "p2p-port", "", p2p.DefaultPort, "Sets the peer-to-peer port, 0 for non-listening mode")
	rootCmd.PersistentFlags().BoolVarP(&p2pNoBootnodes, "p2p-no-bootnodes", "", false, "When specified, do not make connections to chain-specific bootnodes")

	rootCmd.PersistentFlags().StringVarP(&rpcPath, "rpc-path", "", rpc.DefaultPath, "Sets the endpoint for RPC POST requests")
	rootCmd.PersistentFlags().StringArrayVarP(&rpcTypes, "rpc-types", "", nil, "Sets the available RPC protocol types")
	rootCmd.PersistentFlags().UintVarP(&rpcPort, "rpc-port", "", p2p.DefaultPort, "Sets the port to use for local RPC")

	rootCmd.PersistentFlags().StringVarP(&telemetryName, "telemetry-name", "", "", "Unique name of this node to report")
	rootCmd.PersistentFlags().StringVarP(&telemetryURL, "telemetry-url", "", telemetry.DefaultURL, "Websocket endpoint for telemetry stats")

	rootCmd.PersistentFlags().UintVarP(&wasmHeapSize, "wasm-heap-size", "", wasm.DefaultHeapSizeKB, "Initial size for the WASM runtime heap (KB)")

	rootCmd.PersistentFlags().StringVarP(&chain, "chain", "", "main", "Use the chain specified, one of (dev, main) or custom '<chain>.json'")
	rootCmd.PersistentFlags().StringVarP(&clientID, "client-id", "", client.ID(), "The client/version identifier for the running node")
	rootCmd.PersistentFlags().StringVarP(&role, "role", "", "full", "Set either (full, light) as the type of role the node operates as")
}

// Execute ...
func Execute() {
	setup()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	logger.Set(logger.ContextHook{}, false)
	Execute()
}
