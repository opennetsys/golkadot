package clientdb

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	clientchainloader "github.com/opennetsys/golkadot/client/chain/loader"
	clientdbtypes "github.com/opennetsys/golkadot/client/db/types"
	clienttypes "github.com/opennetsys/golkadot/client/types"
	"github.com/opennetsys/golkadot/common/bnutil"
	db "github.com/opennetsys/golkadot/common/db"
	diskdb "github.com/opennetsys/golkadot/common/diskdb"
	triedb "github.com/opennetsys/golkadot/common/triedb"
	"github.com/opennetsys/golkadot/common/u8util"
	types "github.com/opennetsys/golkadot/types"
)

// Config ...
type Config struct {
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

// InterfaceChainDbs ...
type InterfaceChainDbs interface {
	Snapshot()
}

// StorageMethodU8a ...
type StorageMethodU8a struct {
	base      *Base
	createKey types.StorageFunction
}

// NewStorageMethodU8a ...
func NewStorageMethodU8a(dbs db.BaseDB, createKey types.StorageFunction) StorageMethodU8a {
	return StorageMethodU8a{
		base:      NewBase(dbs),
		createKey: createKey,
	}
}

// Del ...
func (s *StorageMethodU8a) Del(keyParam interface{}) {
	s.base.Del(s.createKey(keyParam))
}

// Get ...
func (s *StorageMethodU8a) Get(keyParam interface{}) []uint8 {
	return s.base.Get(s.createKey(keyParam))
}

// Set ...
func (s *StorageMethodU8a) Set(value []uint8, keyParam interface{}) {
	s.base.Set(s.createKey(keyParam), value)
}

// OnUpdate ...
func (s *StorageMethodU8a) OnUpdate(callback func(value []uint8)) {
	s.base.OnUpdate(callback)
}

// StorageMethodBn ...
type StorageMethodBn struct {
	base      *Base
	createKey types.StorageFunction
	bitLen    int
}

// NewStorageMethodBn ...
func NewStorageMethodBn(dbs db.BaseDB, createKey types.StorageFunction, bitLen int) StorageMethodBn {
	return StorageMethodBn{
		base:      NewBase(dbs),
		createKey: createKey,
		bitLen:    bitLen,
	}
}

// Del ...
func (s *StorageMethodBn) Del(keyParams ...interface{}) {
	s.base.Del(s.createKey(keyParams))
}

// Get ...
func (s *StorageMethodBn) Get(keyParams ...interface{}) *big.Int {
	return u8util.ToBN(s.base.Get(s.createKey(keyParams)), true)
}

// Set ...
func (s *StorageMethodBn) Set(value *big.Int, keyParams ...interface{}) {
	s.base.Set(s.createKey(keyParams), bnutil.ToUint8Slice(value, s.bitLen, true, false))
}

// OnUpdate ...
func (s *StorageMethodBn) OnUpdate(callback func(value []uint8)) {
	s.base.OnUpdate(callback)
}

// DB ...
type DB struct {
	BlocksDB *BlockDB
	StateDB  *StateDB
	BasePath string
	Config   *clientdbtypes.Config
}

// NewDB ...
func NewDB(config *clienttypes.ConfigClient, chain *clientchainloader.Loader) *DB {
	if config == nil {
		log.Fatal("config must not be nil")
	}

	if config.DB == nil {
		log.Fatal("config db must not be nil")
	}

	ret := &DB{}
	ret.Config = config.DB

	if config.DB.Type == "disk" {
		ret.BasePath = fmt.Sprintf("%s/chains/%s/%s", config.DB.Path, chain.ID, hex.EncodeToString(chain.GenesisRoot))
	}

	// NOTE: blocks compress very well
	ret.BlocksDB = NewBlockDB(ret.createBackingDB("block.db", true))

	basedb := ret.createBackingDB("state.db", false)
	codec := triedb.NewTrieCodec()
	// NOTE: state RLP does not compress well here
	ret.StateDB = createStateDB(
		triedb.NewTrieDB(
			db.TXDB(db.NewTransactionDB(&basedb)),
			nil,
			codec,
		),
	)

	ret.BlocksDB.DB.Open()
	ret.StateDB.DB.Open()

	return ret
}

// createBackingDB ...
func (c *DB) createBackingDB(name string, isCompressed bool) db.BaseDB {
	var dbs db.BaseDB
	if c.Config.Type == "disk" {
		diskdbs := diskdb.NewDiskDB(c.BasePath, name, &db.BaseDBOptions{
			IsCompressed: isCompressed,
		})
		dbs = db.BaseDB(diskdbs)
	} else {
		memdb := db.NewMemoryDB(nil)
		dbs = db.BaseDB(memdb)
	}

	return dbs
}

// Snapshot ...
func (c *DB) Snapshot() {
	if !c.Config.Snapshot {
		return
	}

	basedb := c.createBackingDB("state.db.snapshot", false)
	codec := triedb.NewTrieCodec()
	newDb := triedb.NewTrieDB(
		db.TXDB(db.NewTransactionDB(&basedb)),
		nil,
		codec,
	)

	newDb.Open()

	c.StateDB.DB.Snapshot(newDb, c.createProgress())
	c.StateDB.DB.Close()
	c.StateDB.DB.Rename(c.BasePath, fmt.Sprintf("state.db.backup-%d", time.Now().Unix()))

	newDb.Close()
	newDb.Rename(c.BasePath, "state.db")
	newDb.Open()
}

// createProgress ...
func (c *DB) createProgress() db.ProgressCB {
	var lastUpdate int64
	var spin int

	spinner := []string{"|", "/", "-", "\\"}
	prepend := strings.Repeat(" ", 37)

	return func(progress *db.ProgressValue) {
		now := time.Now().Unix()

		if (now - lastUpdate) > 200 {
			percent := fmt.Sprintf("      %v", progress.Percent)
			percent = percent[len(percent)-6 : len(percent)]
			keys := fmt.Sprintf("%d", progress.Keys)
			if progress.Keys > 9999 {
				keys = fmt.Sprintf("%vk", progress.Keys/1e3)
			}

			log.Printf("%s%s %s%%, %s keys\n", prepend, spinner[spin%len(spinner)], percent, keys)

			lastUpdate = now
			spin++
		}
	}
}

// Blocks ...
func (c *DB) Blocks() *BlockDB {
	return c.BlocksDB
}

// State ...
func (c *DB) State() *StateDB {
	return c.StateDB
}

// Base ...
type Base struct {
	db db.BaseDB
}

// NewBase ...
func NewBase(dbs db.BaseDB) *Base {
	return &Base{
		db: dbs,
	}
}

// Del ...
func (b *Base) Del(key []uint8) {
	b.db.Del(key)
}

// Get ...
func (b *Base) Get(key []uint8) []uint8 {
	value := b.db.Get(key)

	return value
}

// Set ...
func (b *Base) Set(key []uint8, value []uint8) []uint8 {
	b.db.Put(key, value)
	// b.subscribers.each(func(subscriber) {
	// subscriber(value)
	//})

	return value
}

// OnUpdate  ...
func (b *Base) OnUpdate(subscriber func(value []uint8)) {
	// b.subscribers = append(b.subscribers, subscriber)
}
