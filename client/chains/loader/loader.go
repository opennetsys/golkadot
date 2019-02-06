package chainloader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	clientchaintypes "github.com/opennetsys/go-substrate/client/chains/types"
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

var defaultChain = "../json/dev.json"

// NewLoader ...
func NewLoader(config *clienttypes.ConfigClient) *Loader {
	loader := &Loader{}
	chain := config.Chain

	if chain == nil {
		chain = &defaultChain
	}

	loader.Chain = LoadJSON(*chain)

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

// LoadJSON ...
func LoadJSON(path string) *clientchaintypes.ChainJSON {
	var chain *clientchaintypes.ChainJSON

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
