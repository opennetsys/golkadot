package clientchain

import (
	"fmt"
	"log"
	"math/big"

	clientchainloader "github.com/opennetsys/golkadot/client/chain/loader"
	clientchaintypes "github.com/opennetsys/golkadot/client/chain/types"
	clientdb "github.com/opennetsys/golkadot/client/db"
	clientruntime "github.com/opennetsys/golkadot/client/runtime"
	storagetypes "github.com/opennetsys/golkadot/client/storage/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	clientwasm "github.com/opennetsys/golkadot/client/wasm"
	"github.com/opennetsys/golkadot/common/crypto"
	"github.com/opennetsys/golkadot/common/hexutil"
	"github.com/opennetsys/golkadot/common/triehash"
	"github.com/opennetsys/golkadot/common/u8compact"
	"github.com/opennetsys/golkadot/common/u8util"
	"github.com/opennetsys/golkadot/logger"
)

// Chain ...
type Chain struct {
	Blocks   *clientdb.BlockDB
	Chain    *clientchaintypes.ChainJSON
	Executor *clientwasm.Executer
	Genesis  *clientchaintypes.ChainGenesis
	State    *clientdb.StateDB
}

// NewChain ...
func NewChain(config *clienttypes.ConfigClient) (*Chain, error) {
	var err error

	chain := clientchainloader.NewLoader(config)
	dbs := clientdb.NewDB(config, chain)

	c := &Chain{
		Chain:  chain.Chain,
		Blocks: dbs.Blocks(),
		State:  dbs.State(),
	}

	c.Genesis, err = c.InitGenesis()
	if err != nil {
		return nil, err
	}

	bestHash := c.Blocks.BestHash.Get(nil)
	bestNumber := c.Blocks.BestNumber.Get(nil)
	logGenesis := ""

	if bestNumber == nil {
		log.Fatal("[client/chain] bestNumber is nil")
	}

	if bestNumber.Cmp(big.NewInt(0)) != 0 {
		logGenesis = fmt.Sprintf("(genesis %s)", u8util.ToHex(c.Genesis.Block.Hash[:], 48, true))
	}

	fmt.Printf("%s, #%s, %s %s", c.Chain.Name, bestNumber.String(), u8util.ToHex(bestHash, 48, true), logGenesis)

	// NOTE: Snapshot _before_ we attach the runtime since it ties directly to the backing DBs
	dbs.Snapshot()

	runtime := clientruntime.NewRuntime(c.State.DB)
	c.Executor = clientwasm.NewExecuter(config, c.Blocks, c.State, runtime)

	return c, nil
}

// InitGenesis ...
func (c *Chain) InitGenesis() (*clientchaintypes.ChainGenesis, error) {
	bestHash := c.Blocks.BestHash.Get(nil)
	if bestHash == nil || len(bestHash) == 0 {
		return c.CreateGenesis()
	}

	bestBlock := c.GetBlock(bestHash)

	return c.InitGenesisFromBest(bestBlock.Header, true), nil
}

// InitGenesisFromBest ...
// NOTE: the default for rollback bool should be true
func (c *Chain) InitGenesisFromBest(bestHeader *clienttypes.Header, rollback bool) *clientchaintypes.ChainGenesis {
	if bestHeader.StateRoot == nil {
		// TODO: throw err
		logger.Error("[chain] state root is nil")
	}
	hexState := u8util.ToHex(bestHeader.StateRoot[:], 48, true)
	fmt.Printf("[chain] initializing from state %s\n", hexState)

	c.State.DB.SetRoot(bestHeader.StateRoot[:])

	if u8util.ToHex(c.State.DB.GetRoot(), 48, true) != hexState {
		log.Fatalf("[chain] unable to move state to %s\n", hexState)
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
		fmt.Printf("[chain] unable to validate root, moving to block #%d, %s\n", prevNumber, u8util.ToHex(prevHash, 48, true))

		prevBlock := c.GetBlock(prevHash)

		c.Blocks.BestHash.Set(prevHash, nil)
		c.Blocks.BestNumber.Set(prevBlock.Header.BlockNumber)

		return c.InitGenesisFromBest(prevBlock.Header, false)
	}

	log.Fatal("[chain] unable to retrieve genesis hash. aborting\n")
	return nil
}

// GetBlock ...
func (c *Chain) GetBlock(headerHash []uint8) *clienttypes.BlockData {
	data := c.Blocks.BlockData.Get(headerHash)

	if data == nil || len(data) == 0 {
		log.Fatalf("[chain] unable to retrieve block %s\n", u8util.ToHex(headerHash, -1, true))
	}

	return clienttypes.NewBlockData(data)
}

// GetCode ...
func (c *Chain) GetCode() []uint8 {
	_, decodedValue := u8compact.StripLength(storagetypes.Substrate.Code(nil), -1)

	code := c.State.DB.Get(decodedValue)

	if code == nil || len(code) == 0 {
		log.Fatal("[chain] unable to retrieve genesis code")
	}

	return code
}

// CreateGenesis ...
func (c *Chain) CreateGenesis() (*clientchaintypes.ChainGenesis, error) {
	c.CreateGenesisState()

	genesis, err := c.CreateGenesisBlock()
	if err != nil {
		return nil, err
	}

	c.Blocks.BestHash.Set(genesis.Block.Hash[:], nil)
	c.Blocks.BestNumber.Set(big.NewInt(0))
	c.Blocks.BlockData.Set(genesis.Block.ToU8a(), genesis.Block.Hash)
	c.Blocks.Hash.Set(genesis.Block.Hash[:], 0)

	return genesis, nil
}

// CreateGenesisBlock ...
func (c *Chain) CreateGenesisBlock() (*clientchaintypes.ChainGenesis, error) {
	header, err := clienttypes.NewHeader(nil, nil)
	if err != nil {
		return nil, err
	}
	header.SetStateRoot(crypto.NewBlake2b256(c.State.DB.GetRoot())[:])
	header.SetExtrinsicsRoot(crypto.NewBlake2b256(triehash.TrieRoot(nil))[:])
	header.SetParentHash(crypto.NewBlake2b256(make([]uint8, 32))[:])

	block := clienttypes.NewBlockData(map[string]interface{}{
		"hash":   header.Hash,
		"header": header,
	})

	return &clientchaintypes.ChainGenesis{
		Block: block,
		Code:  c.GetCode(),
	}, nil
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

// GetBestBlocksHash ...
func (c *Chain) GetBestBlocksHash() ([]byte, error) {
	return c.Blocks.BestHash.Get(nil), nil
}

// GetBestBlocksNumber ...
func (c *Chain) GetBestBlocksNumber() (*big.Int, error) {
	return c.Blocks.BestNumber.Get(), nil
}

// GetBlockDataByHash ...
func (c *Chain) GetBlockDataByHash(hash []byte) (*clienttypes.StateBlock, error) {
	// TODO
	return nil, nil
}

// GetGenesisHash ...
func (c *Chain) GetGenesisHash() ([]byte, error) {
	// TODO
	return nil, nil
}

// ImportBlock ...
func (c *Chain) ImportBlock(block *clienttypes.StateBlock) (bool, error) {
	// TODO
	return false, nil
}
