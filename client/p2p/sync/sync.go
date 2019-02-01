package sync

import (
	"context"
	"errors"
	"math"
	"math/big"
	"strings"
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
		bestQueued:    new(big.Int),
		blockQueue:    make(clienttypes.StateBlockQueue),
		blockRequests: make(clienttypes.StateBlockRequests),
		chain:         chn,
		config:        cfg,
		ctx:           ctx,
		handlers:      make(map[synctypes.EventEnum]clienttypes.EventCallback),
		BestSeen:      new(big.Int),
		Status:        synctypes.Idle,
	}

	go s.processBlocks()

	return s, nil
}

// PeerRequests ...
func (s *Sync) PeerRequests(pr clienttypes.InterfacePeer) (clienttypes.Requests, error) {
	if pr == nil {
		return nil, errors.New("nil peer")
	}

	prID := pr.GetID()

	var ret clienttypes.Requests
	for k := range s.blockRequests {
		// note: don't need to check "ok" because using range
		sr, _ := s.blockRequests[k]

		if sr.Peer != nil && sr.Peer.GetID() != "" && sr.Peer.GetID() == prID {
			ret = append(ret, sr)
		}
	}

	return ret, nil
}

func (s *Sync) processBlocks() {
	for {
		select {
		case <-s.ctx.Done():
			{
				logger.Printf("[sync] processBlocks context done\n%v", s.ctx.Err())
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
	if uint(len(s.blockQueue)) > defaults.Defaults.MinIdleBlocks {
		status = synctypes.Sync
	}

	s.Status = status
}

func (s *Sync) processBlock() (bool, error) {
	// const start = Date.now();
	bestNumber, err := s.chain.GetBestBlocksNumber()
	if err != nil {
		logger.Errorf("[sync] err getting best chain blocks number")
		return false, err
	}
	nextNumber := big.NewInt(1)
	nextNumber = nextNumber.Add(bestNumber, nextNumber)
	hasImported := false

	s.setStatus()

	if block, ok := s.blockQueue[nextNumber.String()]; ok {
		logger.Infof("[sync ] importing block #%s", nextNumber.String())

		// TODO: executor?
		ok, err = s.chain.ImportBlock(block)
		if err != nil {
			logger.Errorf("[sync] err importing block\n%v", err)
			return false, err
		}
		if !ok {
			return false, nil
		}

		delete(s.blockQueue, nextNumber.String())

		mod := new(big.Int)
		mod = mod.Set(nextNumber)
		mod = mod.Mod(mod, big.NewInt(int64(REPORT_COUNT)))
		zero := big.NewInt(0)
		if mod.Cmp(zero) == 0 || len(s.blockQueue) < 10 {
			s.handleEvent(synctypes.Imported)
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

// ProvideBlocks ...
func (s *Sync) ProvideBlocks(pr clienttypes.InterfacePeer, request *clienttypes.BlockRequest) error {
	if pr == nil {
		return errors.New("nil peer")
	}
	if request == nil {
		return errors.New("nil request")
	}

	current := request.FromValue
	best, err := s.chain.GetBestBlocksNumber()
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
	if strings.ToUpper(request.Direction) == "ASCENDING" {
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

	request, ok := s.blockRequests[pr.GetID()]
	defer delete(s.blockRequests, pr.GetID())

	if !ok {
		// TODO: nil check
		logger.Warnf("Unrequested response from %v", pr.Cfg().Peer.ShortID)
		return nil

	} else if response.ID != request.ID {
		//logger.Warnf("Mismatched response from %v", pr.Cfg().ShortID)
		//return nil
	}

	bestNumber, err := s.chain.GetBestBlocksNumber()
	if err != nil {
		return err
	}

	var (
		firstNumber            *big.Int
		count                  int
		block, dbBlock         *clienttypes.StateBlock
		header                 *clienttypes.Header
		queueNumber            string
		isImportable, canQueue bool
	)
	for idx := range response.Blocks {
		block = response.Blocks[idx]
		if block == nil {
			continue
		}

		dbBlock, err = s.chain.GetBlockDataByHash(block.Block.Hash)
		if err != nil {
			logger.Errorf("[sync] err getting block by hash\n%v", err)
			return err
		}

		header = block.Block.Header
		queueNumber = header.BlockNumber.String()
		// TODO: why dbBlock == nil ?
		isImportable = dbBlock == nil || dbBlock.Block == nil || len(dbBlock.Block.Body) == 0 || bestNumber.Cmp(header.BlockNumber) == -1
		_, ok = s.blockQueue[queueNumber]
		canQueue = isImportable && !ok

		if canQueue {
			s.blockQueue[queueNumber] = &clienttypes.StateBlock{
				Block: block.Block,
				Peer:  pr,
			}
			if firstNumber == nil {
				firstNumber = header.BlockNumber
			}

			if s.bestQueued.Cmp(header.BlockNumber) == -1 {
				s.bestQueued = header.BlockNumber
			}

			count++
		}
	}

	if count != 0 && firstNumber != nil {
		logger.Infof("Queued %d blocks from %s, %s", count, pr.Cfg().Peer.ShortID, firstNumber.String())
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
	nextNumber, err := s.chain.GetBestBlocksNumber()
	if err != nil {
		return err
	}
	nextNumber = nextNumber.Add(nextNumber, one)
	from := new(big.Int)
	if s.bestQueued.Cmp(nextNumber) == -1 {
		from.Set(nextNumber)
	} else {
		tmpBest := new(big.Int)
		tmpBest.Set(s.bestQueued)
		tmpMaxQueued := big.NewInt(int64(defaults.Defaults.MaxQueuedBlocks / 2))
		tmpBest = tmpBest.Sub(tmpBest, nextNumber)
		if tmpBest.Cmp(tmpMaxQueued) == -1 {
			s.bestQueued = s.bestQueued.Add(s.bestQueued, one)
		}
	}

	if pr.Cfg().Peer.BestNumber.Cmp(s.BestSeen) == 1 {
		s.BestSeen = pr.Cfg().Peer.BestNumber
	}

	// TODO: This assumes no stale block downloading
	_, ok := s.blockRequests[pr.GetID()]
	if ok || from == nil || from.Cmp(pr.Cfg().Peer.BestNumber) == 1 {
		return nil
	}

	logger.Infof("Requesting blocks from %v, %v", pr.Cfg().Peer.ShortID, from.String())

	timeout := time.Now().Add(time.Duration(REQUEST_TIMEOUT) * time.Millisecond)
	nextID := pr.GetNextID()
	request := &clienttypes.BlockRequest{
		From: int(from.Int64()),
		ID:   uint64(nextID),
		Max:  int(defaults.Defaults.MaxRequestBlocks),
	}

	s.blockRequests[pr.GetID()] = &clienttypes.StateRequest{
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

	for k := range s.blockRequests {
		if s.blockRequests[k].Timeout <= now {
			delete(s.blockRequests, k)
		}
	}
}

// On ...
func (s *Sync) On(event synctypes.EventEnum, cb clienttypes.EventCallback) {
	s.handlers[event] = cb

	return
}

func (s *Sync) handleEvent(event synctypes.EventEnum) {
	if event == nil {
		logger.Info("nil event")
		return
	}

	cb, ok := s.handlers[event]
	if !ok {
		logger.Infof("[sync] no event for %s", event.String())
		return
	}

	iface, err := cb()
	logger.Infof("[sync] handled event %s\nresults:\n%v\n%v", event.String(), iface, err)
	return
}
