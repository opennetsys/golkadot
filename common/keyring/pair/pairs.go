package pair

import (
	"errors"

	"github.com/c3systems/go-substrate/common/keyring/address"
)

// NewPairs ...
func NewPairs() (*Pairs, error) {
	return &Pairs{
		PairMap: make(MapPair),
	}, nil
}

// Add ...
func (p *Pairs) Add(pair *Pair) (*Pair, error) {
	if pair == nil {
		return nil, errors.New("nil pair")
	}
	if pair.State == nil {
		return nil, errors.New("nil state")
	}
	// note: just make the map?
	if p.PairMap == nil {
		return nil, errors.New("nil pair map")
	}

	p.PairMap[string(pair.State.PublicKey[:])] = pair

	return pair, nil
}

// All ...
func (p *Pairs) All() ([]*Pair, error) {
	// note: just make the map?
	if p.PairMap == nil {
		return nil, errors.New("nil pair map")
	}

	var pairs []*Pair
	for k := range p.PairMap {
		pairs = append(pairs, p.PairMap[k])
	}

	return pairs, nil
}

// Get ...
func (p *Pairs) Get(addr []byte) (*Pair, error) {
	// note: just make the map?
	if p.PairMap == nil {
		return nil, errors.New("nil pair map")
	}

	decoded, err := address.Decode(string(addr), nil)
	if err != nil {
		return nil, err
	}

	pair, ok := p.PairMap[string(decoded)]
	if !ok {
		// note: or just return nil, nil?
		return nil, errors.New("no such pair")
	}

	return pair, nil
}

// Remove ...
func (p *Pairs) Remove(addr []byte) error {
	// note: just make the map?
	if p.PairMap == nil {
		return errors.New("nil pair map")
	}

	decoded, err := address.Decode(string(addr), nil)
	if err != nil {
		return err
	}

	delete(p.PairMap, string(decoded))

	return nil
}
