package wasm

import (
	clientdb "github.com/c3systems/go-substrate/client/db"
	"github.com/c3systems/go-substrate/client/runtime"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/types"
)

// Executer ...
type Executer struct {
	BlockDB *clientdb.BlockDB
	Config  *clienttypes.ConfigClient
	Runtime runtime.Interface
	StateDB *clientdb.StateDB
}

// NewExecuter ...
// TODO: config client?
func NewExecuter(config *clienttypes.ConfigClient, blockDB *clientdb.BlockDB, stateDB *clientdb.StateDB, runtime runtime.Interface) *Executer {
	return &Executer{
		BlockDB: blockDB,
		Config:  config,
		StateDB: stateDB,
		Runtime: runtime,
	}
}

// ExecuteBlock ...
func (e *Executer) ExecuteBlock(blockData *clienttypes.BlockData, forceNew bool) bool {
	// TODO
	/*
		start := time.Now().Unix()

		importBlock := NewImportBlock(blockData)
		buffer := importBlock.Uint8Slice()
		result := e.Call("Core_execute_block", forceNew)(buffer)

		return result.bool
	*/
	return false
}

// ImportBlock ...
func (e *Executer) ImportBlock(blockData *types.BlockData) bool {
	// TODO
	/*
	   const start = Date.now();
	   const { blockNumber, hash } = blockData.header;
	   const shortHash = u8aToHex(hash, 48);

	   try {
	     this.stateDb.db.transaction(() =>
	       this.executeBlock(blockData)
	     );
	   } catch (error) {
	     l.error(`Failed importing #${blockNumber}, ${shortHash}`);

	     throw error;
	   }

	   this.blockDb.bestHash.set(hash);
	   this.blockDb.bestNumber.set(blockNumber);
	   this.blockDb.blockData.set(blockData.toU8a(), hash);
	   this.blockDb.hash.set(hash, blockNumber);

	   return false;
	*/
	return false
}

// Call ...
type Call struct{}

// Call ...
func (e *Executer) Call(name string, forceNew bool) *Call {
	// TODO

	/*
	   const code = this.stateDb.db.get(CODE_KEY);

	   assert(code, 'Expected to have code available in runtime');

	   // @ts-ignore code check above
	   const instance = createWasm({ config: this.config, l }, this.runtime, code, proxy, forceNew);
	   const { heap } = this.runtime.environment;

	   return (...data: Array<Uint8Array>): CallResult => {
	     const start = Date.now();

	     // runtime.instrument.start();

	     const params = data.reduce((params, data) => {
	       params.push(heap.set(heap.allocate(data.length), data));
	       params.push(data.length);

	       return params;
	     }, ([] as number[]));

	     const lo: number = instance[name].apply(null, params);
	     const hi: number = instance['get_result_hi']();

	     return {
	       bool: hi === 0 && lo === 1,
	       hi,
	       lo
	     };
	   };
	*/
	return nil
}
