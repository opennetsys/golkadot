package db

// DiskDB ...
type DiskDB struct {
	TransactionDB
}

// NewDiskDB ...
func NewDiskDB(base, name string, options *BaseDBOptions) *DiskDB {
	//flatdb := NewFileFlatDB(base, name, options)
	//lrudb : NewLruDB(flatdb)
	return &DiskDB{
		// db: lrudb,
	}
}
