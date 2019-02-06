package chains

import (
	"fmt"
	"log"
	"math/big"

	clientchainloader "github.com/opennetsys/go-substrate/client/chains/loader"
	clientchaintypes "github.com/opennetsys/go-substrate/client/chains/types"
	clientdb "github.com/opennetsys/go-substrate/client/db"
	"github.com/opennetsys/go-substrate/client/runtime"
	storagetypes "github.com/opennetsys/go-substrate/client/storage/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/client/wasm"
	"github.com/opennetsys/go-substrate/common/crypto"
	"github.com/opennetsys/go-substrate/common/hexutil"
	"github.com/opennetsys/go-substrate/common/triehash"
	"github.com/opennetsys/go-substrate/common/u8compact"
	"github.com/opennetsys/go-substrate/common/u8util"
	"github.com/opennetsys/go-substrate/logger"
)

// Chain ...
type Chain struct {
	Blocks   *clientdb.BlockDB
	Chain    *clientchaintypes.ChainJSON
	Executor *wasm.Executer
	Genesis  *clientchaintypes.ChainGenesis
	State    *clientdb.StateDB
}

// NewChain ...
// TODO: configClient?
func NewChain(config *clienttypes.ConfigClient) *Chain {
	chain := clientchainloader.NewLoader(config)
	dbs := clientdb.NewDB(config, chain)

	c := &Chain{
		Chain:  chain.Chain,
		Blocks: dbs.Blocks(),
		State:  dbs.State(),
	}

	c.Genesis = c.InitGenesis()

	bestHash := c.Blocks.BestHash.Get()
	bestNumber := c.Blocks.BestNumber.Get()
	logGenesis := ""
	if bestNumber.Cmp(big.NewInt(0)) != 0 {
		logGenesis = fmt.Sprintf("(genesis %s)", u8util.ToHex(c.Genesis.Block.Hash[:], 48, true))
	}

	fmt.Printf("%s, #%s, %s %s", c.Chain.Name, bestNumber.String(), u8util.ToHex(bestHash, 48, true), logGenesis)

	// NOTE: Snapshot _before_ we attach the runtime since it ties directly to the backing DBs
	dbs.Snapshot()

	runtime := runtime.NewRuntime(c.State.DB)
	c.Executor = wasm.NewExecuter(config, c.Blocks, c.State, runtime)

	return c
}

// InitGenesis ...
func (c *Chain) InitGenesis() *clientchaintypes.ChainGenesis {
	bestHash := c.Blocks.BestHash.Get()
	if bestHash == nil || len(bestHash) == 0 {
		return c.CreateGenesis()
	}

	bestBlock := c.GetBlock(bestHash)

	return c.InitGenesisFromBest(bestBlock.Header, true)
}

// InitGenesisFromBest ...
// NOTE: the default for rollback bool should be true
func (c *Chain) InitGenesisFromBest(bestHeader *clienttypes.Header, rollback bool) *clientchaintypes.ChainGenesis {
	if bestHeader.StateRoot == nil {
		// TODO: throw err
		logger.Error("[chain] state root is nil")
	}
	hexState := u8util.ToHex(bestHeader.StateRoot[:], 48, true)
	fmt.Printf("Initializing from state %s", hexState)

	c.State.DB.SetRoot(bestHeader.StateRoot[:])

	if u8util.ToHex(c.State.DB.GetRoot(), 48, true) != hexState {
		log.Fatalf("Unable to move state to %s", hexState)
	}

	genesisHash := c.Blocks.Hash.Get(0)
	if genesisHash == nil || len(genesisHash) == 0 {
		return c.RollbackBlock(bestHeader, rollback)
	}

	genesisBlock := c.GetBlock(genesisHash)

	return &clientchaintypes.ChainGenesis{
		Block: genesisBlock,
		Code:  c.GetCode(),
	}
}

// RollbackBlock ...
// NOTE: the default for rollback bool should be true
func (c *Chain) RollbackBlock(bestHeader *clienttypes.Header, rollback bool) *clientchaintypes.ChainGenesis {
	prevHash := bestHeader.ParentHash[:]
	// TODO: use big.Int rather than Int64()?
	prevNumber := bestHeader.BlockNumber.Int64() - 1

	if rollback && prevNumber > 1 {
		fmt.Printf("Unable to validate root, moving to block #%d, %s\n", prevNumber, u8util.ToHex(prevHash, 48, true))

		prevBlock := c.GetBlock(prevHash)

		c.Blocks.BestHash.Set(prevHash)
		c.Blocks.BestNumber.Set(prevBlock.Header.BlockNumber)

		return c.InitGenesisFromBest(prevBlock.Header, false)
	}

	panic("Unable to retrieve genesis hash, aborting")
}

// GetBlock ...
func (c *Chain) GetBlock(headerHash []uint8) *clienttypes.BlockData {
	data := c.Blocks.BlockData.Get(headerHash)

	if data == nil || len(data) == 0 {
		log.Fatalf("Unable to retrieve block %s\n", u8util.ToHex(headerHash, -1, true))
	}

	return clienttypes.NewBlockData(data)
}

// GetCode ...
func (c *Chain) GetCode() []uint8 {
	_, decodedValue := u8compact.StripLength(storagetypes.Substrate.Code(), 32)
	code := c.State.DB.Get(decodedValue)

	if code == nil || len(code) == 0 {
		panic("Unable to retrieve genesis code")
	}

	return code
}

// CreateGenesis ...
func (c *Chain) CreateGenesis() *clientchaintypes.ChainGenesis {
	c.CreateGenesisState()

	genesis := c.CreateGenesisBlock()

	c.Blocks.BestHash.Set(genesis.Block.Hash[:])
	c.Blocks.BestNumber.Set(big.NewInt(0))
	c.Blocks.BlockData.Set(genesis.Block.ToU8a(), genesis.Block.Hash)
	c.Blocks.Hash.Set(genesis.Block.Hash[:], 0)

	return genesis
}

// CreateGenesisBlock ...
func (c *Chain) CreateGenesisBlock() *clientchaintypes.ChainGenesis {
	header := clienttypes.NewHeader()
	header.SetStateRoot(crypto.NewBlake2b256(c.State.DB.GetRoot()))
	header.SetExtrinsicsRoot(crypto.NewBlake2b256(triehash.TrieRoot(nil)))
	header.SetParentHash(crypto.NewBlake2b256(make([]uint8, 32)))

	block := clienttypes.NewBlockData(map[string]interface{}{
		"hash":   header.Hash,
		"header": header,
	})

	return &clientchaintypes.ChainGenesis{
		Block: block,
		Code:  c.GetCode(),
	}
}

// CreateGenesisState ...
func (c *Chain) CreateGenesisState() {
	chain := c.Chain
	raw := chain.Genesis.Raw

	if ok, err := c.State.DB.Transaction(func() bool {
		for key, value := range raw {
			k, err := hexutil.ToUint8Slice(key, -1)
			if err != nil {
				return false
			}
			v, err := hexutil.ToUint8Slice(value, -1)
			if err != nil {
				return false
			}
			c.State.DB.Put(k, v)
		}

		return true
	}); err != nil || !ok {
		// TODO: return err?
		logger.Errorf("[chain] statedb ok: %v, err:\n%v", ok, err)
	}
}
