package chainspec

// Genesis structure
type Genesis struct {
	Raw Raw `json:"raw"`
}

// Raw structure
type Raw map[string]string

// Chainspec structure
type Chainspec struct {
	BootNodes    []string          `json:"bootNodes"`
	ID           string            `json:"id"`
	Genesis      Genesis           `json:"genesis"`
	GenesisRoot  string            `json:"genesisRoot"`
	Name         string            `json:"name"`
	Properties   map[string]string `json:"properties"`
	ProtocolID   *string           `json:"protocolId"`
	TelemetryURL string            `json:"telemetryUrl"`
}
