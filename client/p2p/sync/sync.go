package sync

import (
	"context"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/c3systems/go-substrate/client/p2p/defaults"
	synctypes "github.com/c3systems/go-substrate/client/p2p/sync/types"
	clienttypes "github.com/c3systems/go-substrate/client/types"
	"github.com/c3systems/go-substrate/common/u8util"
	"github.com/c3systems/go-substrate/logger"
)

// note: ensure the struct emplements the interface
var _ clienttypes.InterfaceSync = (*Sync)(nil)

// New ...
func New(ctx context.Context, cfg *clienttypes.ConfigClient, chn clienttypes.InterfaceChains) (*Sync, error) {
	if cfg == nil {
		return nil, ErrNilConfig
	}
	if chn == nil {
		return nil, ErrNilChain
	}

	s := &Sync{
		Chain:  chn,
		Config: cfg,
	}

	go s.processBlocks(ctx)

	return s, nil
}

// PeerRequests ...
func (s *Sync) PeerRequests(pr clienttypes.InterfacePeer) (clienttypes.Requests, error) {
	if pr == nil {
		return nil, errors.New("nil peer")
	}

	prID := pr.GetID()

	var ret clienttypes.Requests
	for k := range s.BlockRequests {
		sr, ok := s.BlockRequests[k]
		if !ok {
			continue
		}

		if sr.Peer != nil && sr.Peer.GetID() != "" && sr.Peer.GetID() == prID {
			ret = append(ret, sr)
		}
	}

	return ret, nil
}

func (s *Sync) processBlocks(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			{
				logger.Printf("[sync] processBlocks context done\n%v", ctx.Err())
				return
			}
		default:
			{
				timeout := 1 * time.Millisecond
				hasOne, err := s.processBlock()
				if err != nil {
					logger.Errorf("[sync] err processing block\n%v", err)
				}

				if !hasOne {
					timeout = 100 * time.Millisecond
				}

				time.Sleep(timeout)
			}
		}
	}
}

func (s *Sync) setStatus() {
	status := synctypes.Idle
	if uint(len(s.BlockQueue)) > defaults.Defaults.MinIdleBlocks {
		status = synctypes.Sync
	}

	s.Status = status
}

func (s *Sync) processBlock() (bool, error) {
	// const start = Date.now();
	bestNumber, err := s.Chain.GetBestBlocksNumber()
	if err != nil {
		logger.Errorf("[sync] err getting best chain blocks number")
		return false, err
	}
	nextNumber := big.NewInt(1)
	nextNumber = nextNumber.Add(bestNumber, nextNumber)
	hasImported := false

	s.setStatus()

	if block, ok := s.BlockQueue[nextNumber.String()]; ok {
		logger.Infof("Importing block #%s", nextNumber.String())

		ok, err := s.Chain.ImportBlock(block)
		if err != nil {
			logger.Errorf("[sync] err importing block\n%v", err)
			return false, err
		}
		if !ok {
			return false, nil
		}

		delete(s.BlockQueue, nextNumber.String())

		mod := new(big.Int)
		mod = mod.Set(nextNumber)
		mod = mod.Mod(nextNumber, big.NewInt(int64(REPORT_COUNT)))
		zero := big.NewInt(0)
		if mod.Cmp(zero) == 0 || len(s.BlockQueue) < 10 {
			// TODO...
			//this.emit('imported');
		}

		hasImported = true

		// if (this.lastPeer !== peer || !queueLength) {
		//   if (this.lastPeer !== null || !queueLength) {
		//     this.requestBlocks(peer);
		//   }

		//   this.lastPeer = peer;
		// }
	}

	return hasImported, nil

	// if (count) {
	//   l.log(`#${startNumber.toString()}- ${count} imported (${Date.now() - start}ms)`);
	// }
}

// TODO finish...
func (s *Sync) provideBlocks(pr clienttypes.InterfacePeer, request *clienttypes.BlockRequest) error {
	current := request.FromValue
	best, err := s.Chain.GetBestBlocksNumber()
	if err != nil {
		return err
	}

	var blocks []*clienttypes.StateBlock

	// FIXME: Also send blocks starting with hash
	maxReq := uint(request.Max)
	if maxReq == 0 {
		maxReq = defaults.Defaults.MaxRequestBlocks
	}

	max := math.Min(float64(maxReq), float64(defaults.Defaults.MaxRequestBlocks))
	count := 0.0
	if u8util.IsU8a(request.From) {
		count = max
	}

	// note: use enum?
	increment := big.NewInt(-1)
	if request.Direction == "Ascending" {
		increment = big.NewInt(1)
	}

	zero := big.NewInt(0)
	for ; count < max && current.Cmp(best) == -1 && current.Cmp(zero) == -1; count++ {
		// const hash = this.chain.state.blockHashAt.get(current);
		//
		// blocks.push(
		//   this.getBlockData(request.fields.values, hash)
		// );

		count++
		current = current.Add(current, increment)
	}

	ok, err := pr.Send(&clienttypes.BlockResponse{
		Blocks: blocks,
		ID:     request.ID,
	})

	if err != nil {
		return err
	}
	if !ok {
		return errors.New("send not ok")
	}

	return nil
}

// QueueBlocks ...
func (s *Sync) QueueBlocks(pr clienttypes.InterfacePeer, response *clienttypes.BlockResponse) error {
	if pr == nil {
		return errors.New("nil peer")
	}
	if response == nil {
		return errors.New("nil response")
	}

	request, ok := s.BlockRequests[pr.GetID()]
	defer delete(s.BlockRequests, pr.GetID())

	if !ok {
		logger.Warnf("Unrequested response from %v", pr.Cfg().ShortID)
		return nil

	} else if response.ID != request.ID {
		//logger.Warnf("Mismatched response from %v", pr.Cfg().ShortID)
		//return nil
	}

	bestNumber, err := s.Chain.GetBestBlocksNumber()
	if err != nil {
		return err
	}

	var (
		firstNumber *big.Int
		count       int
	)

	for idx := range response.Blocks {
		block := response.Blocks[idx]
		// TODO: why dbBlock not used?!
		//dbBlock, err := s.Chain.GetBlockDataByHash(block.Block.Hash)
		if err != nil {
			logger.Errorf("[sync] err getting block by hash\n%v", err)
			return err
		}

		header := block.Block.Header
		queueNumber := header.BlockNumber.String()
		// TODO: len?
		isImportable := bestNumber.Cmp(header.BlockNumber) == -1
		_, ok = s.BlockQueue[queueNumber]
		canQueue := isImportable && !ok

		if canQueue {
			s.BlockQueue[queueNumber] = &clienttypes.StateBlock{
				Block: block.Block,
				Peer:  pr,
			}
			if firstNumber == nil {
				firstNumber = header.BlockNumber
			}

			if s.BestQueued.Cmp(header.BlockNumber) == -1 {
				s.BestQueued = header.BlockNumber
			}

			count++
		}
	}

	if count != 0 && firstNumber != nil {
		logger.Infof("Queued %d blocks from %s, %s", count, pr.Cfg().ShortID, firstNumber.String())
	}

	return nil
}

// RequestBlocks ...
func (s *Sync) RequestBlocks(pr clienttypes.InterfacePeer) error {
	s.timeoutRequests()

	isActive, err := pr.IsActive()
	if err != nil {
		return err
	}

	if !isActive {
		return errors.New("peer is not active")
	}

	one := big.NewInt(1)
	nextNumber, err := s.Chain.GetBestBlocksNumber()
	if err != nil {
		return err
	}
	nextNumber = nextNumber.Add(nextNumber, one)
	from := new(big.Int)
	if s.BestQueued.Cmp(nextNumber) == -1 {
		from.Set(nextNumber)
	} else {
		tmpBest := new(big.Int)
		tmpBest.Set(s.BestQueued)
		tmpMaxQueued := big.NewInt(int64(defaults.Defaults.MaxQueuedBlocks / 2))
		tmpBest = tmpBest.Sub(tmpBest, nextNumber)
		if tmpBest.Cmp(tmpMaxQueued) == -1 {
			s.BestQueued = s.BestQueued.Add(s.BestQueued, one)
		}
	}

	if pr.Cfg().BestNumber.Cmp(s.BestSeen) == 1 {
		s.BestSeen = pr.Cfg().BestNumber
	}

	// TODO: This assumes no stale block downloading
	_, ok := s.BlockRequests[pr.GetID()]
	if ok || from == nil || from.Cmp(pr.Cfg().BestNumber) == 1 {
		return nil
	}

	logger.Infof("Requesting blocks from %v, %v", pr.Cfg().ShortID, from)

	timeout := time.Now().Add(time.Duration(REQUEST_TIMEOUT) * time.Millisecond)
	nextID, err := pr.GetNextID()
	if err != nil {
		return err
	}
	request := &clienttypes.Request{
		From: from,
		ID:   nextID,
		Max:  uint64(defaults.Defaults.MaxRequestBlocks),
	}

	s.BlockRequests[pr.GetID()] = &clienttypes.StateRequest{
		Peer:    pr,
		Request: request,
		Timeout: timeout.UnixNano() / int64(time.Millisecond),
	}

	ok, err = pr.Send(request)
	if err != nil {
		logger.Errorf("[sync] err sending peer message %v\n%v", *request, err)
		return err
	}
	if !ok {
		logger.Errorf("[sync] peer could not send message: %v", *request)
		// note: return err here?
	}

	return nil
}

// TODO We can probably use a package with a timeout like an LRU
func (s *Sync) timeoutRequests() {
	// note: get time in ms
	now := time.Now().UnixNano() / int64(time.Millisecond)

	for k := range s.BlockRequests {
		if s.BlockRequests[k].Timeout <= now {
			delete(s.BlockRequests, k)
		}
	}
}

// ProvideBlocks ...
// TODO ...
func (s *Sync) ProvideBlocks() {

}

// On ...
// TODO ...
func (s *Sync) On(event synctypes.EventEnum, cb clienttypes.EventCallback) (interface{}, error) {
	return nil, nil
}
