package db

// BaseOptions ...
type BaseOptions struct {
}

// ProgressValue ...
type ProgressValue struct {
	IsCompleted bool
	Keys        int
	Percent     int
}

// ProgressCB ...
type ProgressCB func(*ProgressValue)

// BaseDBOptions ....
type BaseDBOptions struct {
	IsCompressed bool
}

// BaseDB ...
type BaseDB interface {
	Close()
	Open()
	Drop()
	Empty()
	Maintain(fn *ProgressCB) error
	Rename(base, file string)
	Size() int

	Del(key []uint8)
	Get(key []uint8) []uint8
	Put(key, value []uint8)
}

// TXDB ...
type TXDB interface {
	BaseDB
	Transaction(fn func() bool) bool
}
