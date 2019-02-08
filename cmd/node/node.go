package main

import (
	"fmt"
	"os"

	"github.com/opennetsys/godot/client"
	"github.com/opennetsys/godot/client/p2p"
	"github.com/opennetsys/godot/client/rpc"
	"github.com/opennetsys/godot/client/telemetry"
	clienttypes "github.com/opennetsys/godot/client/types"
	"github.com/opennetsys/godot/client/wasm"
	"github.com/opennetsys/godot/common/db"
	"github.com/opennetsys/godot/logger"
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
	var p2pPort int
	var p2pMaxPeers int
	var p2pNoBootnodes bool

	var rpcPath string
	var rpcTypes []string
	var rpcPort int

	var telemetryName string
	var telemetryURL string

	var wasmHeapSize int

	var chain string
	var role string
	var clientID string

	rootCmd = &cobra.Command{
		Use:   "start",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
			config := &clienttypes.ConfigClient{}
			cl := client.NewClient()
			cl.Start(config)

		},
	}

	rootCmd.PersistentFlags().BoolVarP(&dbCompact, "db-compact", "", false, "Compact existing databases")
	rootCmd.PersistentFlags().BoolVarP(&dbSnapshot, "db-snapshot", "", false, "Create trie snapshot, drop & restore database from this")
	rootCmd.PersistentFlags().StringVarP(&dbPath, "db-path", "", db.DefaultPath, "Sets the path for all storage operations")
	rootCmd.PersistentFlags().StringVarP(&dbType, "db-type", "", db.DefaultType, "The type of database storage to use")

	rootCmd.PersistentFlags().StringVarP(&p2pAddress, "p2p-address", "", p2p.DefaultAddress, "The interface to bind to (p2p-port > 0)")
	rootCmd.PersistentFlags().StringArrayVarP(&p2pNodes, "p2p-nodes", "", nil, "Reserved nodes to make initial connections to")
	rootCmd.PersistentFlags().IntVarP(&p2pMaxPeers, "p2p-max-peers", "", p2p.DefaultMaxPeers, "The maximum allowed peers")
	rootCmd.PersistentFlags().IntVarP(&p2pPort, "p2p-port", "", p2p.DefaultPort, "Sets the peer-to-peer port, 0 for non-listening mode")
	rootCmd.PersistentFlags().BoolVarP(&p2pNoBootnodes, "p2p-no-bootnodes", "", false, "When specified, do not make connections to chain-specific bootnodes")

	rootCmd.PersistentFlags().StringVarP(&rpcPath, "rpc-path", "", rpc.DefaultPath, "Sets the endpoint for RPC POST requests")
	rootCmd.PersistentFlags().StringArrayVarP(&rpcTypes, "rpc-types", "", nil, "Sets the available RPC protocol types")
	rootCmd.PersistentFlags().IntVarP(&rpcPort, "rpc-port", "", p2p.DefaultPort, "Sets the port to use for local RPC")

	rootCmd.PersistentFlags().StringVarP(&telemetryName, "telemetry-name", "", "", "Unique name of this node to report")
	rootCmd.PersistentFlags().StringVarP(&telemetryURL, "telemetry-url", "", telemetry.DefaultURL, "Websocket endpoint for telemetry stats")

	rootCmd.PersistentFlags().IntVarP(&wasmHeapSize, "wasm-heap-size", "", wasm.DefaultHeapSizeKB, "Initial size for the WASM runtime heap (KB)")

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
