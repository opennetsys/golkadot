package db

import (
	"errors"
)

// KV ...
type KV struct {
	Key   []uint8
	Value []uint8
}

// Overlay ...
type Overlay map[string]*KV

// TransactionDB ...
type TransactionDB struct {
	TXDB
	Backing   BaseDB
	txOverlay Overlay
	txStarted bool
}

// NewTransactionDB ...
func NewTransactionDB(backing *BaseDB) *TransactionDB {
	return &TransactionDB{
		Backing:   *backing,
		txOverlay: Overlay{},
		txStarted: false,
	}
}

// Transaction ...
func (t *TransactionDB) Transaction(fn func() bool) (bool, error) {
	t.CreateTx()
	result := fn()

	if result {
		if err := t.CommitTx(); err != nil {
			return false, t.RevertTx()
		}

		return result, nil
	}

	return false, t.RevertTx()
}

// Close ...
func (t *TransactionDB) Close() {
	t.Backing.Close()
}

// Open ...
func (t *TransactionDB) Open() {
	t.Backing.Open()
}

// Drop ...
func (t *TransactionDB) Drop() {
	t.Backing.Drop()
}

// Empty ...
func (t *TransactionDB) Empty() {
	t.Backing.Empty()
}

// Rename ...
func (t *TransactionDB) Rename(base, file string) {
	t.Backing.Rename(base, file)
}

// Maintain ...
func (t *TransactionDB) Maintain(fn *ProgressCB) error {
	if t.txStarted {
		return errors.New("cannot maintain inside an open transaction")
	}

	return t.Backing.Maintain(fn)
}

// Size ...
func (t *TransactionDB) Size() int {
	return t.Backing.Size()
}

// Del ...
func (t *TransactionDB) Del(key []uint8) {
	if t.txStarted {
		t.txOverlay[string(key)] = &KV{
			Key:   key,
			Value: nil,
		}
		return
	}

	t.Backing.Del(key)
}

// Get ...
func (t *TransactionDB) Get(key []uint8) []uint8 {
	if t.txStarted {
		value, found := t.txOverlay[string(key)]

		if found {
			return value.Value
		}
	}

	return t.Backing.Get(key)
}

// Put ...
func (t *TransactionDB) Put(key, value []uint8) {
	if t.txStarted {
		t.txOverlay[string(key)] = &KV{
			Key:   key,
			Value: value,
		}

		return
	}

	t.Backing.Put(key, value)
}

// CreateTx ...
func (t *TransactionDB) CreateTx() error {
	if t.txStarted {
		return errors.New("cannot create a transaction when one is already active")
	}

	t.txOverlay = Overlay{}
	t.txStarted = true
	return nil
}

// CommitTx ...
func (t *TransactionDB) CommitTx() error {
	if !t.txStarted {
		return errors.New("cannot commit when not in transaction")
	}

	for _, kv := range t.txOverlay {
		if kv.Value == nil {
			t.Backing.Del(kv.Key)
		} else {
			t.Backing.Put(kv.Key, kv.Value)
		}
	}

	t.txOverlay = Overlay{}
	t.txStarted = false
	return nil
}

// RevertTx ...
func (t *TransactionDB) RevertTx() error {
	if !t.txStarted {
		return errors.New("cannot revert when not in transaction")
	}

	t.txOverlay = Overlay{}
	t.txStarted = false
	return nil
}
