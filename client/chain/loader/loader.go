package chainloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	chainjson "github.com/opennetsys/go-substrate/client/chain/json"
	clientchaintypes "github.com/opennetsys/go-substrate/client/chain/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
	"github.com/opennetsys/go-substrate/common/triehash"
	"github.com/opennetsys/go-substrate/common/u8util"
)

// Loader ...
type Loader struct {
	Chain       *clientchaintypes.ChainJSON
	ID          string
	GenesisRoot []uint8
}

var defaultChain = "dev"

// NewLoader ...
func NewLoader(config *clienttypes.ConfigClient) *Loader {
	loader := &Loader{}

	if config.Chain == "" || config.Chain == "dev" {
		loader.Chain = loadJSON([]byte(chainjson.Dev))
	} else {
		loader.Chain = loadJSONPath(config.Chain)
	}

	loader.GenesisRoot = loader.CalculateGenesisRoot()
	loader.ID = loader.Chain.ID

	return loader
}

// CalculateGenesisRoot ...
func (loader *Loader) CalculateGenesisRoot() []uint8 {
	chain := loader.Chain
	raw := chain.Genesis.Raw

	var pairs []*triehash.TriePair
	for k, v := range raw {
		pairs = append(pairs, &triehash.TriePair{
			K: u8util.FromHex(k),
			V: u8util.FromHex(v),
		})
	}

	return triehash.TrieRoot(pairs)
}

// loadJSONPath...
func loadJSONPath(path string) *clientchaintypes.ChainJSON {
	var chain *clientchaintypes.ChainJSON

	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("file at path %q does not exist", path)
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = json.Unmarshal(data, &chain)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return chain
}

// loadJSON ...
func loadJSON(data []byte) *clientchaintypes.ChainJSON {
	var chain *clientchaintypes.ChainJSON
	err := json.Unmarshal(data, &chain)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return chain
}
