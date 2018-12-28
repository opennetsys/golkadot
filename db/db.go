package db

// Trie ...
type Trie struct {
}

// TrieDB ...
type TrieDB MemoryDB

// NewTrie ...
func NewTrie() *Trie {
	/*
	   export default class Trie extends Impl implements TrieDb {
	     constructor (db: TxDb = new MemoryDb(), rootHash?: Uint8Array, codec?: Codec) {
	       super(db, rootHash, codec);
	*/

	return &Trie{}
}

// Transaction ...
func (t *Trie) Transaction(fn func() bool) bool {
	/*
	   try {
	     this.createCheckpoint();

	     const result = this.db.transaction(fn);

	     if (result) {
	       this.commitCheckpoint();
	     } else {
	       this.revertCheckpoint();
	     }

	     return result;
	   } catch (error) {
	     this.revertCheckpoint();

	     throw error;
	   }
	*/
	return true
}

// Open ...
func (t *Trie) Open() {
	//t.db.Open()
}

// Close ...
func (t *Trie) Close() {
	//t.db.Close()
}

// Empty ...
func (t *Trie) Empty() {
	//t.db.Empty()
}

// Drop ...
func (t *Trie) Drop() {
	//t.db.Drop()
}

// Maintain ...
func (t *Trie) Maintain(progressCb func()) {
	//t.db.Maintain(progressCb)
}

// Rename ...
func (t *Trie) Rename(base, file string) {
	//t.db.Rename(base, file)
}

// Size ...
func (t *Trie) Size() int {
	return 0
	//return t.db.Size()
}

// Delete ...
func (t *Trie) Delete(key []uint8) {
	/*
	   this._setRootNode(
	     this._del(
	       this._getNode(this.rootHash),
	       toNibbles(key)
	     )
	   );
	*/
}

// Get ...
func (t *Trie) Get(key []uint8) []uint8 {
	/*
	   return this._get(
	     this._getNode(this.rootHash),
	     toNibbles(key)
	   );
	*/
	return nil
}

// Put ...
func (t *Trie) Put(key, value []uint8) []uint8 {
	/*
	   this._setRootNode(
	     this._put(
	       this._getNode(this.rootHash),
	       toNibbles(key),
	       value
	     )
	   );
	*/
	return nil
}

// GetRoot ...
func (t *Trie) GetRoot() []uint8 {
	/*
		rootnode := this.GetNode()
		if isNull(rootNode) {
			return []uint8{}
		}

		return this.rootHash
	*/
	return nil
}

// Node ..
type Node struct{}

// GetNode ...
func (t *Trie) GetNode(hash []uint8) *Node {
	//return this._getNode(hash || this.rootHash);
	return nil
}

// SetRoot ...
func (t *Trie) SetRoot(rootHash []uint8) {
	//this.rootHash = rootHash
}

// Snapshot ...
func (t *Trie) Snapshot(dest *TrieDB, fn ProgressCB) int {
	/*
	   const start = Date.now();

	   l.log('creating current state snapshot');

	   const keys = this._snapshot(dest, fn, this.rootHash, 0, 0, 0);
	   const elapsed = (Date.now() - start) / 1000;

	   dest.setRoot(this.rootHash);

	   const newSize = dest.db.size();
	   const percentage = 100 * (newSize / this.db.size());
	   const sizeMB = newSize / (1024 * 1024);

	   l.log(`snapshot created in ${elapsed.toFixed(2)}s, ${(keys / 1000).toFixed(2)}k keys, ${sizeMB.toFixed(2)}MB (${percentage.toFixed(2)}%)`);

	   fn({
	     isCompleted: true,
	     keys,
	     percent: 100
	   });

	   return keys;
	*/
	return 0
}
