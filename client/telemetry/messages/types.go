package messages

// Level ...
type Level string

//type Level = "INFO"

// Message ...
type Message string

// type Message = 'system.connected' | 'system.interval' | 'node.start' | 'block.import';

// SyncStatus ...
type SyncStatus int

// BaseJSON ...
type BaseJSON struct {
	level Level
	msg   Message
	ts    string
}

// BlockJSON ...
type BlockJSON struct {
	level  Level
	msg    Message
	ts     string
	best   string
	height int
}

// ConnectedJSON ...
type ConnectedJSON struct {
	level          Level
	msg            Message
	ts             string
	chain          string
	config         string
	implementation string
	name           string
	version        string
}

// IntervalJSON ...
type IntervalJSON struct {
	peers   int
	status  SyncStatus
	txcount int
}
