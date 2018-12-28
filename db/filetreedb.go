package db

// FileTreeDB ...
type FileTreeDB struct {
}

// NewFileTreeDBDB ...
func NewFileTreeDBDB(options *BaseOptions) *FileTreeDB {
	return &FileTreeDB{}
}

// Close ...
func (f *FileTreeDB) Close() {

}

// Open ...
func (f *FileTreeDB) Open() {

}

// Drop ...
func (f *FileTreeDB) Drop() {

}

// Empty ...
func (f *FileTreeDB) Empty() {

}

// Rename ...
func (f *FileTreeDB) Rename(base, file string) {

}

// Maintain ...
func (f *FileTreeDB) Maintain(fn *ProgressCB) error {
	return nil
}

// Size ...
func (f *FileTreeDB) Size() int {
	return 0
}

// Del ...
func (f *FileTreeDB) Del(key []uint8) {

}

// Get ...
func (f *FileTreeDB) Get(key []uint8) []uint8 {
	return nil
}

// Put ...
func (f *FileTreeDB) Put(key, value []uint8) {

}
