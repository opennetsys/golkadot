package wasm

// DefaultHeapSizeKB ...
// NOTE: Rust implementation starts with 8 pages, auto-grow in in-place
var DefaultHeapSizeKB uint = 8 * 64
