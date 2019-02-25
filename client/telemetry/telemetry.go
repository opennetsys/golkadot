package telemetry

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	ws "github.com/gorilla/websocket"
	chains "github.com/opennetsys/golkadot/client/chain"
	clientdb "github.com/opennetsys/golkadot/client/db"
	synctypes "github.com/opennetsys/golkadot/client/p2p/sync/types"
	messages "github.com/opennetsys/golkadot/client/telemetry/messages"
	clienttypes "github.com/opennetsys/golkadot/client/types"
)

// Telemetry ...
type Telemetry struct {
	Blocks    *clientdb.BlockDB
	IsActive  bool
	Chain     string
	Name      string
	URL       string
	Websocket *ws.Conn
}

// NewTelemetry ...
func NewTelemetry(config *clienttypes.ConfigClient, chain chains.Chain) *Telemetry {
	tel := config.Telemetry
	name := strings.TrimSpace(tel.Name)

	isActive := len(name) > 0 && len(tel.URL) > 0

	return &Telemetry{
		Blocks:   chain.Blocks,
		IsActive: isActive,
		Chain:    chain.Chain.Name,
		Name:     name,
		URL:      tel.URL,
	}
}

// Start ...
func (t *Telemetry) Start() {
	if !t.IsActive {
		return
	}

	t.Connect()
}

// Connect ...
func (t *Telemetry) Connect() {
	log.Printf("Connecting to telemtry, url=%s, name=%s\n", t.URL, t.Name)

	wsconn, _, err := ws.DefaultDialer.Dial(t.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	wsconn.SetCloseHandler(func(code int, text string) error {
		log.Println("Disconnected from telemetry")
		t.Websocket = nil

		time.AfterFunc(5*time.Second, t.Connect)

		return nil
	})

	log.Println("Connected to telemetry")
	t.Websocket = wsconn
	t.SendInitial()
}

// BlockImported ...
func (t *Telemetry) BlockImported() {
	bestHash := t.Blocks.BestHash.Get(nil)
	bestNumber := t.Blocks.BestNumber.Get(nil)
	t.Send(messages.NewBlockImport(bestHash, bestNumber))
}

// IntervalInfo ...
func (t *Telemetry) IntervalInfo(peers int, status synctypes.StatusEnum) {
	bestHash := t.Blocks.BestHash.Get(nil)
	bestNumber := t.Blocks.BestNumber.Get(nil)
	t.Send(messages.NewInterval(bestHash, bestNumber, peers, status))
}

// SendInitial ...
func (t *Telemetry) SendInitial() {
	bestHash := t.Blocks.BestHash.Get(nil)
	bestNumber := t.Blocks.BestNumber.Get(nil)
	t.Send(messages.NewStarted(bestHash, bestNumber))
	t.Send(messages.NewConnected(t.Chain, t.Name))
}

// Send ...
func (t *Telemetry) Send(message messages.Base) {
	if t.Websocket == nil {
		return
	}

	b, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sending %s\n", string(b))
	err = t.Websocket.WriteMessage(ws.TextMessage, b)
	if err != nil {
		log.Fatal(err)
	}
}
