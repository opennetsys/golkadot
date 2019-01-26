package sync

import (
	"context"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/c3systems/go-substrate/block"
	"github.com/c3systems/go-substrate/chain"
	"github.com/c3systems/go-substrate/client"
	"github.com/c3systems/go-substrate/logger"
	"github.com/c3systems/go-substrate/p2p/defaults"
	peertypes "github.com/c3systems/go-substrate/p2p/peer/types"
)

// New ...
func New(ctx context.Context, cfg *client.Config, chn chain.Interface) (*Sync, error) {
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
func (s *Sync) PeerRequests(pr peertypes.InterfacePeer) (Requests, error) {
	if pr == nil {
		return nil, errors.New("nil peer")
	}
	prID := pr.Config().ID
	if prID == nil {
		return nil, errors.New("nil perr id")
	}

	var ret Requests
	for k := range s.BlockRequests {
		sr, ok := s.BlockRequests[k]
		if !ok {
			continue
		}

		if sr.Peer != nil && sr.Peer.Config().ID != nil && *sr.Peer.Config().ID == *prID {
			ret = append(ret, sr)
		}
	}

	return ret, nil
}

func (s *Sync) processBlocks() {
	timeout := 1 * time.Millisecond
	hasOne, err := this.processBlock()
	if err != nil {
		logger.Errorf("[sync] err processing block\n%v", err)
	}

	if !hasOne {
		timeout = 100 * time.Millisecond
	}

	time.AfterFunc(timeout, s.processBlocks)
}

func (s *Sync) setStatus() {
	status := Idle
	if len(s.BlockQueue) > defaults.Defaults.MinIdleBlocks {
		status = Sync
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
		mod = mod.Mod(nextNumber, REPORT_COUNT)
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

	return false, hasImported

	// if (count) {
	//   l.log(`#${startNumber.toString()}- ${count} imported (${Date.now() - start}ms)`);
	// }
}

// TODO finish...
func (s *Sync) provideBlocks(pr peertypes.InterfacePeer, request *client.BlockRequest) error {
	current := request.FromValue
	best := s.Chain.GetBestBlocksNumber()
	// TODO: change...
	var blocks []interface{}

	// FIXME: Also send blocks starting with hash
	max := math.Min(request.Max || defaults.MaxRequestBlocks, defaults.MaxRequestBlocks)
	count := 0
	if u8autil.IsU8a(request.From) {
		count = max
	}

	// note: use enum?
	increment := big.NewInt(-1)
	if request.Direction.String() == "Ascending" {
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
		current = current.Add(increment)
	}

	pr.send(&client.BlockResponse{
		Blocks: blocks,
		ID:     request.ID,
	})
}

// QueueBlocks ...
func (s *Sync) QueueBlocks(pr peertypes.InterfacePeer, response *client.BlockResponse) error {
	if pr == nil {
		return errors.New("nil peer")
	}
	if response == nil {
		return errors.New("nil response")
	}

	request, ok := s.BlockRequests[pr.Cfg().ID]
	defer delete(s.BlockRequests, pr.Cfg().ID)

	if !ok {
		logger.Warnf("Unrequested response from %v", pr.Cfg().ShortID)
		return nil

	} else if response.Cfg().ID.Cmp(request.Cfg().ID) != 0 {
		//logger.Warnf("Mismatched response from %v", pr.Cfg().ShortID)
		//return nil
	}

	bestNumber := s.Chain.GetBestBlocksNumber()
	var (
		firstNumber, count int
	)

	for idx := range response.Cfg().Blocks {
		block := response.Cfg().Blocks[idx]
		dbBlock, err := s.Chain.GetBlockDataByHash(block.Hash)
		if err != nil {
			logger.Errorf("[sync] err getting block by hash\n%v", err)
			return err
		}

		header := block.Header
		queueNumber := header.BlockNumber.String()
		isImportable := len(dbBlock.length) != 0 || bestNumber.Cmp(header.BlockNumber) == -1
		_, ok = s.BlockQueue[queueNumber]
		canQueue := isImportable && !ok

		if canQueue {
			s.BlockQueue[queueNumber] = &StateBlock{
				Block: block,
				Peer:  pr,
			}
			if firstNumber == 0 {
				firstNumber = header.BlockNumber
			} else {
				// note: unecessary line of code?
				firstNumber = firstNumber
			}

			if s.BestQueued.Cmp(header.BlockNumber) == -1 {
				s.BestQueued = header.BlockNumber
			}

			count++
		}
	}

	if count != 0 && firstNumber != 0 {
		logger.Infof("Queued %d blocks from %s, %v", count, pr.Cfg().ShortID, firstNumber)
	}

	return nil
}

// RequestBlocks ...
func (s *Sync) RequestBlocks(pr peertypes.InterfacePeer) error {
	s.timeoutRequests()

	if !pr.isActive() {
		return
	}

	one := big.NewInt(1)
	nextNumber := s.Chain.GetBestBlocksNumber().Add(one)
	from := new(big.Int)
	if s.BestQueued.Cmp(nextNumber) == -1 {
		from.Set(nextNumber)
	} else {
		tmpBest := new(big.Int)
		tmpBest.Set(s.BestQueued)
		tmpMaxQueued := big.NewInt(defaults.MaxQueuedBlocks / 2)
		if tmpBest.Sub(nextNumber).Cmp(tmpMaxQueued) == -1 {
			s.BestQueued = s.BestQueued.Add(one)
		}
	}

	if pr.Cfg().BestNumber.Cmp(s.BestSeen) == 1 {
		s.BestSeen = pr.Cfg().BestNumber
	}

	// TODO: This assumes no stale block downloading
	_, ok := s.BlockRequests[pr.id]
	if ok || from == nil || from.Cmp(pr.Cfg().BestNumber) == 1 {
		return nil
	}

	logger.Infof("Requesting blocks from %v, %v", pr.Cfg().ShortID, from)

	timeout := Time.now().Add(REQUEST_TIMEOUT * time.Millisecond)
	request = &block.Request{
		From: from,
		ID:   pr.GetNextId(),
		Max:  defaults.MaxRequestBlocks,
	}

	s.BlockRequests[pr.Cfg().ID] = &StateRequest{
		Peer:    pr,
		Request: request,
		Timeout: timeout,
	}

	ok, err := pr.Send(request)
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
		if s.BlockRequests.Timeout <= now {
			delete(s.BlockRequests, k)
		}
	}
}
