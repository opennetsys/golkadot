package clientdb

import (
	"github.com/c3systems/go-substrate/client"
	clientchainloader "github.com/c3systems/go-substrate/clientchain/loader"
	clientdbtypes "github.com/c3systems/go-substrate/clientdb/types"
	"github.com/c3systems/go-substrate/common/db"
	"github.com/c3systems/go-substrate/common/triedb"
	types "github.com/c3systems/go-substrate/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-db/src/index.ts

// DBConfig ...
type DBConfig struct {
	Compact  bool
	IsTrieDb bool
	Path     string
	Snapshot bool
	Type     string // DBConfigType
}

// BlockDB ...
type BlockDB struct {
	DB         db.BaseDB
	BestHash   StorageMethodU8a
	BestNumber StorageMethodBn
	BlockData  StorageMethodU8a
	Hash       StorageMethodU8a
	Header     StorageMethodU8a
}

// StateDB ...
type StateDB struct {
	DB   *triedb.TrieDB
	Code StorageMethodU8a
}

// InterfaceChainDbs ...
type InterfaceChainDbs interface {
	Snapshot()
}

// StorageMethodU8a ...
// TODO
type StorageMethodU8a struct {
	//Del(params ...interface{})
	//Get(params ...interface{}) []uint8
	//Set(value []uint8, params ...interface{})
	//OnUpdate(callback func(value []uint8))
}

// Del ...
func (s *StorageMethodU8a) Del(params ...interface{}) {
	// TODO
}

// Get ...
func (s *StorageMethodU8a) Get(params ...interface{}) []uint8 {
	// TODO
	return nil
}

// Set ...
func (s *StorageMethodU8a) Set(value []uint8, params ...interface{}) {
	// TODO
}

// OnUpdate ...
func (s *StorageMethodU8a) OnUpdate(callback func(value []uint8)) {
	// TODO
}

// StorageMethodBn ...
// TODO
type StorageMethodBn struct {
}

// Del ...
func (s *StorageMethodBn) Del(params ...interface{}) {
	// TODO
}

// Get ...
func (s *StorageMethodBn) Get(params ...interface{}) *types.Int {
	// TODO
	return nil
}

// Set ...
func (s *StorageMethodBn) Set(value *types.Int, params ...interface{}) {
	// TODO
}

// OnUpdate ...
func (s *StorageMethodBn) OnUpdate(callback func(value *types.Int)) {
	// TODO
}

// DB ...
type DB struct {
	blocks   *BlockDB
	state    *StateDB
	basePath string
	config   *clientdbtypes.InterfaceDBConfig
}

// NewDB ...
func NewDB(config *client.Config, chain *clientchainloader.Loader) *DB {
	d := &DB{}
	d.config = config.DB
	_ = d
	//blocks := NewBlockDB(config)
	/*
	   this.config = db;
	    this.basePath = db.type === 'disk'
	      ? path.join(db.path, 'chains', chain.id, u8aToHex(chain.genesisRoot))
	      : '';

	    // NOTE blocks compress very well
	    this.blocks = createBlockDb(
	      this.createBackingDb('block.db', true)
	    );
	    // NOTE state RLP does not compress well here
	    this.state = createStateDb(
	      new TrieDb(
	        this.createBackingDb('state.db', false)
	      )
	    );

	    this.blocks.db.open();
	    this.state.db.open();
	*/

	// TODO
	return &DB{
		blocks: nil,
		state:  nil,
	}
}

// Snapshot ...
func (c *DB) Snapshot() {
	// TODO
}

// Blocks ...
func (c *DB) Blocks() *BlockDB {
	return c.blocks
}

// State ...
func (c *DB) State() *StateDB {
	return c.state
}
