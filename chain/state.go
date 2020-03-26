package chain

import (
	"sync"

	"github.com/olympus-protocol/ogen/db/blockdb"
	"github.com/olympus-protocol/ogen/logger"
	"github.com/olympus-protocol/ogen/params"
	"github.com/olympus-protocol/ogen/primitives"
	"github.com/olympus-protocol/ogen/utils/chainhash"
)

// StateService keeps track of the blockchain and its state. This is where pruning should eventually be implemented to
// get rid of old states.
type StateService struct {
	log    *logger.Logger
	lock   sync.RWMutex
	params params.ChainParams

	View     *ChainView
	stateMap map[chainhash.Hash]primitives.State

	sync bool
}

func (s *StateService) IsSync() bool {
	return s.sync
}

func (s *StateService) SetSyncStatus(sync bool) {
	s.sync = sync
	return
}

func (s *StateService) initChainState(db blockdb.DB, params params.ChainParams) error {
	// Get the state snap from db dbindex and deserialize
	s.log.Info("loading chain state...")

	genesisBlock := primitives.GetGenesisBlock(params)

	// load chain state
	loc, err := db.AddRawBlock(&genesisBlock)
	if err != nil {
		return err
	}
	view, err := NewChainView(genesisBlock.Header, *loc)
	if err != nil {
		return err
	}
	s.View = view

	// TODO: load block index
	return nil
}

func (s *StateService) TipState() primitives.State {
	tip := s.View.Tip()
	return s.stateMap[tip.Hash]
}

func (s *StateService) GetStateForHash(hash chainhash.Hash) (*primitives.State, bool) {
	s.lock.RLock()
	oldState, found := s.stateMap[hash]
	s.lock.RUnlock()
	return &oldState, found
}

func (s *StateService) Add(block *primitives.Block, location blockdb.BlockLocation, newTip bool, newState *primitives.State) error {
	row, err := s.View.Add(block.Header, location)
	if err != nil {
		return err
	}
	rowHash := row.Header.Hash()
	s.lock.Lock()
	s.stateMap[rowHash] = *newState
	s.lock.Unlock()
	if newTip {
		err = s.View.SetTip(rowHash)
		if err != nil {
			return err
		}
	}
	return err
}

func NewStateService(log *logger.Logger, params params.ChainParams, db blockdb.DB) (*StateService, error) {
	genesisBlock := primitives.GetGenesisBlock(params)
	genesisHash := genesisBlock.Hash()
	ss := &StateService{
		params: params,
		log:    log,
		sync:   false,
		stateMap: map[chainhash.Hash]primitives.State{
			genesisHash: {
				UtxoState: primitives.UtxoState{
					UTXOs: make(map[chainhash.Hash]primitives.Utxo),
				},
				GovernanceState: primitives.GovernanceState{
					Proposals: make(map[chainhash.Hash]primitives.GovernanceProposal),
				},
				UserState: primitives.UserState{
					Users: make(map[chainhash.Hash]primitives.User),
				},
				WorkerState: primitives.WorkerState{
					Workers: make(map[chainhash.Hash]primitives.Worker),
				},
			},
		},
	}
	err := ss.initChainState(db, params)
	if err != nil {
		return nil, err
	}
	return ss, nil
}
