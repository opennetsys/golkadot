package db

// CachedValue ...
type CachedValue struct {
	Value []uint8
}

// LruMap ...
type LruMap map[string]*CachedValue

// LruDB ...
type LruDB struct {
	backing BaseDB
	lru     LruMap
}

var defaultItemCount = 4096

// NewLruDB ...
func NewLruDB(backing BaseDB, itemCount int) *LruDB {
	if itemCount == -1 {
		itemCount = defaultItemCount
	}

	return &LruDB{
		backing: backing,
		lru:     LruMap{},
	}
}

// Close ...
func (l *LruDB) Close() {
	l.lru = LruMap{}
	l.backing.Close()
}

// Open ...
func (l *LruDB) Open() {
	l.lru = LruMap{}
	l.backing.Open()
}

// Drop ...
func (l *LruDB) Drop() {
	l.backing.Drop()
}

// Empty ...
func (l *LruDB) Empty() {
	l.lru = LruMap{}
	l.backing.Empty()
}

// Rename ...
func (l *LruDB) Rename(base, file string) {
	l.backing.Rename(base, file)
}

// Maintain ...
func (l *LruDB) Maintain(fn *ProgressCB) error {
	return l.backing.Maintain(fn)
}

// Size ...
func (l *LruDB) Size() int {
	return l.backing.Size()
}

// Del ...
func (l *LruDB) Del(key []uint8) {
	keyStr := string(key)

	l.backing.Del(key)
	l.lru[keyStr] = &CachedValue{Value: nil}
}

// Get ...
func (l *LruDB) Get(key []uint8) []uint8 {
	keyStr := string(key)
	cached, found := l.lru[keyStr]
	if found {
		return cached.Value
	}

	value := l.backing.Get(key)
	l.lru[keyStr] = &CachedValue{Value: value}
	return value
}

// Put ...
func (l *LruDB) Put(key, value []uint8) {
	keyStr := string(key)

	l.backing.Put(key, value)
	l.lru[keyStr] = &CachedValue{Value: value}
}
