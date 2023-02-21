package goimpl

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/util"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
)

var (
	AssertionStake       = big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(100))
	ChallengeVertexStake = big.NewInt(params.Ether)

	ErrWrongChain             = errors.New("wrong chain")
	ErrParentDoesNotExist     = errors.New("assertion's parent does not exist on-chain")
	ErrInvalidOp              = errors.New("invalid operation")
	ErrChallengeAlreadyExists = errors.New("challenge already exists on leaf")
	ErrCannotChallengeOwnLeaf = errors.New("cannot challenge own leaf")
	ErrInvalidHeight          = errors.New("invalid block height")
	ErrVertexAlreadyExists    = errors.New("vertex already exists")
	ErrWrongState             = errors.New("vertex state does not allow this operation")
	ErrWrongPredecessorState  = errors.New("predecessor state does not allow this operation")
	ErrNotYet                 = errors.New("deadline has not yet passed")
	ErrNoWinnerYet            = errors.New("challenges does not yet have a winner assertion")
	ErrPastDeadline           = errors.New("deadline has passed")
	ErrInsufficientBalance    = errors.New("insufficient balance")
	ErrNotImplemented         = errors.New("not yet implemented")
	ErrNoLastLeafProof        = errors.New("history commitment must provide a last leaf proof")
	ErrWrongFirstLeaf         = errors.New("first leaf of history does not match required state root")
	ErrWrongLastLeaf          = errors.New("last leaf of history does not match required state root")
	ErrProofFailsToVerify     = errors.New("Merkle proof fails to verify for last state of history commitment")
)

// ChainReadWriter can make mutating and non-mutating calls to the blockchain.
type ChainReadWriter interface {
	ChainReader
	ChainWriter
	EventProvider
}

// ChainReader can make non-mutating calls to the on-chain goimpl. It provides
// an ActiveTx type which has the block number to use when making chain calls.
type ChainReader interface {
	Call(clo func(protocol.ActiveTx) error) error
}

// ChainWriter can make mutating calls to the on-chain goimpl.
type ChainWriter interface {
	Tx(clo func(protocol.ActiveTx) error) error
}

// EventProvider allows subscribing to chain events for the on-chain goimpl.
type EventProvider interface {
	SubscribeChainEvents(ctx context.Context, ch chan<- AssertionChainEvent)
	SubscribeChallengeEvents(ctx context.Context, ch chan<- ChallengeEvent)
}

type AssertionChain struct {
	mutex                         sync.RWMutex
	timeReference                 util.TimeReference
	challengePeriod               time.Duration
	latestConfirmed               protocol.AssertionSequenceNumber
	assertions                    []*Assertion
	assertionsBySeqNum            map[common.Hash]protocol.AssertionSequenceNumber
	challengeVerticesByCommitHash map[protocol.ChallengeHash]map[protocol.VertexHash]*ChallengeVertex
	challengesByCommitHash        map[protocol.ChallengeHash]*Challenge
	balances                      *util.MapWithDefault[common.Address, *big.Int]
	feed                          *EventFeed[AssertionChainEvent]
	challengesFeed                *EventFeed[ChallengeEvent]
	inbox                         *Inbox
	chainId                       uint64
}

const (
	DeadTxStatus = iota
	ReadOnlyTxStatus
	ReadWriteTxStatus
)

// ActiveTx is a transaction that is currently being processed.
type ActiveTx struct {
	TxStatus    int
	BlockNumber *big.Int // If nil, uses the latest block in the chain.
}

// verifyRead is a helper function to verify that the transaction is read-only.
func (tx *ActiveTx) VerifyRead() bool {
	if tx.TxStatus == DeadTxStatus {
		panic("tried to read chain after call ended")
	}
	return true
}

// verifyReadWrite is a helper function to verify that the transaction is read-write.
func (tx *ActiveTx) VerifyReadWrite() bool {
	if tx.TxStatus != ReadWriteTxStatus {
		panic("tried to modify chain in read-only call")
	}
	return true
}

func (tx *ActiveTx) FinalizedBlockNumber() *big.Int {
	return nil
}

func (tx *ActiveTx) HeadBlockNumber() *big.Int {
	return nil
}

func (tx *ActiveTx) ReadOnly() bool {
	return false
}

// Tx enables a mutating call to the on-chain goimpl.
func (chain *AssertionChain) Tx(clo func(tx *ActiveTx) error) error {
	chain.mutex.Lock()
	defer chain.mutex.Unlock()
	tx := &ActiveTx{TxStatus: ReadWriteTxStatus}
	err := clo(tx)
	tx.TxStatus = DeadTxStatus
	return err
}

// Call enables a non-mutating call to the on-chain goimpl.
func (chain *AssertionChain) Call(clo func(tx *ActiveTx) error) error {
	chain.mutex.RLock()
	defer chain.mutex.RUnlock()
	tx := &ActiveTx{TxStatus: ReadOnlyTxStatus}
	err := clo(tx)
	tx.TxStatus = DeadTxStatus
	return err
}

func (chain *AssertionChain) GetChallengeVerticesByCommitHashmap() map[protocol.ChallengeHash]map[protocol.VertexHash]*ChallengeVertex {
	return chain.challengeVerticesByCommitHash
}

func (chain *AssertionChain) GetChallengesByCommitHash() map[protocol.ChallengeHash]*Challenge {
	return chain.challengesByCommitHash
}

func (chain *AssertionChain) GetChallengesFeed() *EventFeed[ChallengeEvent] {
	return chain.challengesFeed
}

func (chain *AssertionChain) GetFeed() *EventFeed[AssertionChainEvent] {
	return chain.feed
}

func (chain *AssertionChain) SetLatestConfirmed(assertionSequenceNumber protocol.AssertionSequenceNumber) {
	chain.latestConfirmed = assertionSequenceNumber
}

const (
	PendingAssertionState = iota
	ConfirmedAssertionState
	RejectedAssertionState
)

// AssertionState is a type used to represent the state of an assertion.
type AssertionState int

// Assertion represents an assertion in the goimpl.
type Assertion struct {
	SequenceNum             protocol.AssertionSequenceNumber `json:"sequence_num"`
	StateCommitment         util.StateCommitment             `json:"state_commitment"`
	Staker                  util.Option[common.Address]
	Prev                    util.Option[*Assertion]
	challengeManager        protocol.ChallengeManager
	status                  AssertionState
	isFirstChild            bool
	firstChildCreationTime  util.Option[time.Time]
	secondChildCreationTime util.Option[time.Time]
	challenge               util.Option[*Challenge]
}

func (a *Assertion) Height() uint64 {
	return a.StateCommitment.Height
}

func (a *Assertion) SeqNum() protocol.AssertionSequenceNumber {
	return protocol.AssertionSequenceNumber(a.SequenceNum)
}

func (a *Assertion) PrevSeqNum() protocol.AssertionSequenceNumber {
	return protocol.AssertionSequenceNumber(a.Prev.Unwrap().SequenceNum)
}

func (a *Assertion) StateHash() common.Hash {
	return a.StateCommitment.StateRoot
}

// NewAssertionChainWithChainId creates a new AssertionChain with specified chainId.
func NewAssertionChainWithChainId(ctx context.Context, timeRef util.TimeReference, challengePeriod time.Duration, chainId uint64) *AssertionChain {
	genesis := &Assertion{
		challengeManager: nil,
		status:           ConfirmedAssertionState,
		SequenceNum:      0,
		StateCommitment: util.StateCommitment{
			Height:    0,
			StateRoot: common.Hash{},
		},
		Prev:                    util.None[*Assertion](),
		isFirstChild:            false,
		firstChildCreationTime:  util.None[time.Time](),
		secondChildCreationTime: util.None[time.Time](),
		challenge:               util.None[*Challenge](),
		Staker:                  util.None[common.Address](),
	}

	genesisKey := crypto.Keccak256Hash(
		binary.BigEndian.AppendUint64(
			genesis.StateCommitment.Hash().Bytes(),
			math.MaxUint64,
		),
	)
	assertionsSeen := map[common.Hash]protocol.AssertionSequenceNumber{
		genesisKey: 0,
	}
	chain := &AssertionChain{
		mutex:                         sync.RWMutex{},
		timeReference:                 timeRef,
		challengePeriod:               challengePeriod,
		challengesByCommitHash:        make(map[protocol.ChallengeHash]*Challenge),
		challengeVerticesByCommitHash: make(map[protocol.ChallengeHash]map[protocol.VertexHash]*ChallengeVertex),
		latestConfirmed:               0,
		assertions:                    []*Assertion{genesis},
		balances:                      util.NewMapWithDefaultAdvanced[common.Address, *big.Int](common.Big0, func(x *big.Int) bool { return x.Sign() == 0 }),
		feed:                          NewEventFeed[AssertionChainEvent](ctx),
		challengesFeed:                NewEventFeed[ChallengeEvent](ctx),
		inbox:                         NewInbox(ctx),
		assertionsBySeqNum:            assertionsSeen,
		chainId:                       chainId,
	}
	genesis.challengeManager = chain
	return chain
}

// NewAssertionChain creates a new AssertionChain with specified chainId.
func NewAssertionChain(ctx context.Context, timeRef util.TimeReference, challengePeriod time.Duration) *AssertionChain {
	return NewAssertionChainWithChainId(ctx, timeRef, challengePeriod, 0)
}

/* Assertion chain methods */

// TimeReference returns the time reference used by the chain.
func (chain *AssertionChain) TimeReference() util.TimeReference {
	return chain.timeReference
}

// Inbox returns the inbox used by the chain.
func (chain *AssertionChain) Inbox() *Inbox {
	return chain.inbox
}

// ChainId returns the chainId used by the chain.
func (chain *AssertionChain) ChainId() uint64 {
	return chain.chainId
}

// GetBalance returns the balance of the given address.
func (chain *AssertionChain) GetBalance(tx protocol.ActiveTx, addr common.Address) *big.Int {
	tx.VerifyRead()
	return chain.balances.Get(addr)
}

// SetBalance sets the balance of the given address.
func (chain *AssertionChain) SetBalance(tx protocol.ActiveTx, addr common.Address, balance *big.Int) {
	tx.VerifyReadWrite()
	oldBalance := chain.balances.Get(addr)
	chain.balances.Set(addr, balance)
	chain.feed.Append(&SetBalanceEvent{Addr: addr, OldBalance: oldBalance, NewBalance: balance})
}

// AddToBalance adds the given amount to the balance of the given address.
func (chain *AssertionChain) AddToBalance(tx protocol.ActiveTx, addr common.Address, amount *big.Int) {
	tx.VerifyReadWrite()
	chain.SetBalance(tx, addr, new(big.Int).Add(chain.GetBalance(tx, addr), amount))
}

// DeductFromBalance deducts the given amount from the balance of the given address.
func (chain *AssertionChain) DeductFromBalance(tx protocol.ActiveTx, addr common.Address, amount *big.Int) error {
	tx.VerifyReadWrite()
	balance := chain.GetBalance(tx, addr)
	if balance.Cmp(amount) < 0 {
		return errors.Wrapf(ErrInsufficientBalance, "%s < %s", balance.String(), amount.String())
	}
	chain.SetBalance(tx, addr, new(big.Int).Sub(balance, amount))
	return nil
}

// LatestConfirmed returns the latest confirmed assertion.
func (chain *AssertionChain) LatestConfirmed(ctx context.Context, tx protocol.ActiveTx) (protocol.Assertion, error) {
	tx.VerifyRead()
	return chain.assertions[chain.latestConfirmed], nil
}

// AssertionBySequenceNum returns the assertion with the given sequence number.
func (chain *AssertionChain) AssertionBySequenceNum(ctx context.Context, tx protocol.ActiveTx, seqNum protocol.AssertionSequenceNumber) (protocol.Assertion, error) {
	tx.VerifyRead()
	if seqNum >= protocol.AssertionSequenceNumber(len(chain.assertions)) {
		return nil, fmt.Errorf("assertion sequence out of range %d >= %d", seqNum, len(chain.assertions))
	}
	return chain.assertions[seqNum], nil
}

/* Challenge manager methods */
func (chain *AssertionChain) CurrentChallengeManager(ctx context.Context, tx protocol.ActiveTx) (protocol.ChallengeManager, error) {
	tx.VerifyRead()
	return chain, nil
}

func (chain *AssertionChain) ChallengePeriodSeconds(
	ctx context.Context, tx protocol.ActiveTx,
) (time.Duration, error) {
	return time.Second, nil
}

func (chain *AssertionChain) CalculateChallengeHash(
	ctx context.Context,
	tx protocol.ActiveTx,
	itemId common.Hash,
	challengeType protocol.ChallengeType,
) (protocol.ChallengeHash, error) {
	return protocol.ChallengeHash{}, nil
}

func (chain *AssertionChain) GetVertex(
	ctx context.Context,
	tx protocol.ActiveTx,
	vertexId protocol.VertexHash,
) (util.Option[protocol.ChallengeVertex], error) {
	return util.None[protocol.ChallengeVertex](), nil
}

func (chain *AssertionChain) GetChallenge(
	ctx context.Context,
	tx protocol.ActiveTx,
	challengeId protocol.ChallengeHash,
) (util.Option[protocol.Challenge], error) {
	return util.None[protocol.Challenge](), nil
}

func (v *ChallengeVertex) ChildrenAreAtOneStepFork(
	ctx context.Context,
	tx protocol.ActiveTx,
) (bool, error) {
	tx.VerifyRead()
	if vertexCommit.Height != vertexParentCommit.Height+1 {
		return false, nil
	}
	vertices, ok := chain.challengeVerticesByCommitHash[challengeCommitHash]
	if !ok {
		return false, fmt.Errorf("challenge vertices not found for assertion with state commit hash %#x", challengeCommitHash)
	}
	parentCommitHash := protocol.VertexHash(vertexParentCommit.Hash())
	return verticesContainOneStepFork(ctx, tx, vertices, parentCommitHash), nil
}

// Check if a vertices with a matching parent commitment hash are at a one-step-fork from their parent.
// First, we filter out vertices with the specified parent commit hash, then check that all of the
// matching vertices are one-step away from their parent.
func verticesContainOneStepFork(ctx context.Context, tx protocol.ActiveTx, vertices map[protocol.VertexHash]protocol.ChallengeVertex, parentCommitHash protocol.VertexHash) bool {
	if len(vertices) < 2 {
		return false
	}
	childVertices := make([]protocol.ChallengeVertex, 0)
	for _, v := range vertices {
		prev, _ := v.GetPrev(ctx, tx)
		if prev.IsNone() {
			continue
		}
		// We only check vertices that have a matching parent commit hash.
		commitment, _ := prev.Unwrap().GetCommitment(ctx, tx)
		vParentHash := protocol.VertexHash(commitment.Hash())
		if vParentHash == parentCommitHash {
			childVertices = append(childVertices, v)
		}
	}
	if len(childVertices) < 2 {
		return false
	}
	for _, vertex := range childVertices {
		if !isOneStepAwayFromParent(ctx, tx, vertex) {
			return false
		}
	}
	return true
}

func isOneStepAwayFromParent(ctx context.Context, tx protocol.ActiveTx, vertex protocol.ChallengeVertex) bool {
	prev, _ := vertex.GetPrev(ctx, tx)
	if prev.IsNone() {
		return false
	}
	prevCommitment, _ := prev.Unwrap().GetCommitment(ctx, tx)
	commitment, _ := vertex.GetCommitment(ctx, tx)
	return commitment.Height == prevCommitment.Height+1
}

// ChallengeVertexByCommitHash returns the challenge vertex with the given commit hash.
func (chain *AssertionChain) ChallengeVertexByCommitHash(
	tx *ActiveTx, challengeHash protocol.ChallengeHash, vertexHash protocol.VertexHash,
) (protocol.ChallengeVertex, error) {
	tx.verifyRead()
	vertices, ok := chain.challengeVerticesByCommitHash[challengeHash]
	if !ok {
		return nil, fmt.Errorf("challenge vertices not found for assertion with state commit hash %#x", challengeHash)
	}
	vertex, ok := vertices[vertexHash]
	if !ok {
		return nil, fmt.Errorf("challenge vertex with sequence number not found %#x", vertexHash)
	}
	return vertex, nil
}

// ChallengeByCommitHash returns the challenge with the given commit hash.
func (chain *AssertionChain) ChallengeByCommitHash(tx protocol.ActiveTx, commitHash protocol.ChallengeHash) (protocol.Challenge, error) {
	tx.verifyRead()
	chal, ok := chain.challengesByCommitHash[commitHash]
	if !ok {
		return nil, errors.Wrapf(ErrVertexAlreadyExists, fmt.Sprintf("Hash: %s", commitHash))
	}
	return chal, nil
}

// SubscribeChainEvents subscribes to chain events.
func (chain *AssertionChain) SubscribeChainEvents(ctx context.Context, ch chan<- AssertionChainEvent) {
	chain.feed.Subscribe(ctx, ch)
}

// SubscribeChallengeEvents subscribes to challenge events.
func (chain *AssertionChain) SubscribeChallengeEvents(ctx context.Context, ch chan<- ChallengeEvent) {
	chain.challengesFeed.Subscribe(ctx, ch)
}

func (chain *AssertionChain) Confirm(
	ctx context.Context,
	tx protocol.ActiveTx,
	blockHash,
	sendRoot common.Hash,
) error {
	return nil
}

func (chain *AssertionChain) Reject(
	ctx context.Context,
	tx protocol.ActiveTx,
	staker common.Address,
) error {
	return nil
}

// CreateLeaf creates a new leaf assertion.
func (chain *AssertionChain) CreateAssertion(
	ctx context.Context,
	tx protocol.ActiveTx,
	height uint64,
	prevAssertionId uint64,
	prevAssertionState *protocol.ExecutionState,
	postState *protocol.ExecutionState,
	prevInboxMaxCount *big.Int,
) (protocol.Assertion, error) {
	tx.VerifyReadWrite()
	if prev.challengeManager.ChainId() != chain.ChainId() {
		return nil, ErrWrongChain
	}
	if prev.StateCommitment.Height >= commitment.Height {
		return nil, ErrInvalidOp
	}

	// Ensure the assertion being created is not a duplicate.
	if _, ok := chain.assertionsBySeqNum[commitment.Hash()]; ok {
		return nil, ErrVertexAlreadyExists
	}

	if !prev.Prev.IsNone() {
		// The parent must exist on-chain.
		prevSeqNum, ok := chain.assertionsBySeqNum[prev.StateCommitment.Hash()]
		if !ok {
			return nil, ErrParentDoesNotExist
		}
		// Parent sequence number must be < the new assertion's assigned sequence number.
		if prevSeqNum >= AssertionSequenceNumber(uint64(len(chain.assertions))) {
			return nil, ErrParentDoesNotExist
		}
	}

	if err := prev.Staker.IfLet(
		func(oldStaker common.Address) error {
			if staker != oldStaker {
				if err := chain.DeductFromBalance(tx, staker, AssertionStake); err != nil {
					return err
				}
				chain.AddToBalance(tx, oldStaker, AssertionStake)
				prev.Staker = util.None[common.Address]()
			}
			return nil
		},
		func() error {
			if err := chain.DeductFromBalance(tx, staker, AssertionStake); err != nil {
				return err
			}
			return nil

		},
	); err != nil {
		return nil, err
	}

	leaf := &Assertion{
		challengeManager:        chain,
		status:                  PendingAssertionState,
		SequenceNum:             protocol.AssertionSequenceNumber(len(chain.assertions)),
		StateCommitment:         commitment,
		Prev:                    util.Some(prev),
		isFirstChild:            prev.firstChildCreationTime.IsNone(),
		firstChildCreationTime:  util.None[time.Time](),
		secondChildCreationTime: util.None[time.Time](),
		challenge:               util.None[*Challenge](),
		Staker:                  util.Some(staker),
	}
	if prev.firstChildCreationTime.IsNone() {
		prev.firstChildCreationTime = util.Some(chain.timeReference.Get())
	} else if prev.secondChildCreationTime.IsNone() {
		prev.secondChildCreationTime = util.Some(chain.timeReference.Get())
	}
	chain.assertions = append(chain.assertions, leaf)
	chain.assertionsBySeqNum[commitment.Hash()] = leaf.SequenceNum
	chain.feed.Append(&CreateLeafEvent{
		PrevStateCommitment: prev.StateCommitment,
		PrevSeqNum:          prev.SequenceNum,
		SeqNum:              leaf.SequenceNum,
		StateCommitment:     leaf.StateCommitment,
		Validator:           staker,
	})
	return leaf, nil
}

// RejectForPrev rejects the assertion and emits the information through feed. It moves assertion to `RejectedAssertionState` state.
func (a *Assertion) RejectForPrev(tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if a.status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("State: %d", a.status))
	}
	if a.Prev.IsNone() {
		return ErrInvalidOp
	}
	if a.Prev.Unwrap().status != RejectedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", a.Prev.Unwrap().status))
	}
	a.status = RejectedAssertionState
	a.challengeManager.GetFeed().Append(&RejectEvent{
		SeqNum: a.SequenceNum,
	})
	return nil
}

// RejectForLoss rejects the assertion and emits the information through feed. It moves assertion to `RejectedAssertionState` state.
func (a *Assertion) RejectForLoss(ctx context.Context, tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if a.status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("State: %d", a.status))
	}
	if a.Prev.IsNone() {
		return ErrInvalidOp
	}
	chal := a.Prev.Unwrap().challenge
	if chal.IsNone() {
		return util.ErrOptionIsEmpty
	}
	winner, err := chal.Unwrap().Winner(ctx, tx)
	if err != nil {
		return err
	}
	if winner == a {
		return ErrInvalidOp
	}
	a.status = RejectedAssertionState
	a.challengeManager.GetFeed().Append(&RejectEvent{
		SeqNum: a.SequenceNum,
	})
	return nil
}

// ConfirmNoRival confirms that there is no rival for the assertion and moves the assertion to `ConfirmedAssertionState` state.
func (a *Assertion) ConfirmNoRival(tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if a.status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("State: %d", a.status))
	}
	if a.Prev.IsNone() {
		return ErrInvalidOp
	}
	prev := a.Prev.Unwrap()
	if prev.status != ConfirmedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", a.Prev.Unwrap().status))
	}
	if !prev.secondChildCreationTime.IsNone() {
		return ErrInvalidOp
	}
	if !a.challengeManager.TimeReference().Get().After(prev.firstChildCreationTime.Unwrap().Add(a.challengeManager.ChallengePeriodLength(tx))) {
		return errors.Wrapf(ErrNotYet, fmt.Sprintf("%d > %d", a.challengeManager.TimeReference().Get().Unix(), prev.firstChildCreationTime.Unwrap().Add(a.challengeManager.ChallengePeriodLength(tx)).Unix()))
	}
	a.status = ConfirmedAssertionState
	a.challengeManager.SetLatestConfirmed(a.SequenceNum)
	a.challengeManager.GetFeed().Append(&ConfirmEvent{
		SeqNum: a.SequenceNum,
	})

	if !a.Staker.IsNone() && a.firstChildCreationTime.IsNone() {
		a.challengeManager.AddToBalance(tx, a.Staker.Unwrap(), AssertionStake)
		a.Staker = util.None[common.Address]()
	}
	return nil
}

// ConfirmForWin confirms that the assertion is the WinnerAssertion of the challenge and moves the assertion to `ConfirmedAssertionState` state.
func (a *Assertion) ConfirmForWin(ctx context.Context, tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if a.status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("State: %d", a.status))
	}
	if a.Prev.IsNone() {
		return ErrInvalidOp
	}
	prev := a.Prev.Unwrap()
	if prev.status != ConfirmedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", a.Prev.Unwrap().status))
	}
	if prev.challenge.IsNone() {
		return ErrWrongPredecessorState
	}
	winner, err := prev.challenge.Unwrap().Winner(ctx, tx)
	if err != nil {
		return err
	}
	if winner != a {
		return ErrInvalidOp
	}
	a.status = ConfirmedAssertionState
	a.challengeManager.SetLatestConfirmed(a.SequenceNum)
	a.challengeManager.GetFeed().Append(&ConfirmEvent{
		SeqNum: a.SequenceNum,
	})
	return nil
}

type ChallengeType uint

const (
	NoChallengeType    ChallengeType = iota
	BlockChallenge                   = 1
	BigStepChallenge                 = 2
	SmallStepChallenge               = 3
)

// Challenge created by an assertion.
type Challenge struct {
	rootAssertion          util.Option[*Assertion]
	WinnerAssertion        util.Option[*Assertion]
	WinnerV                util.Option[*ChallengeVertex]
	rootVertex             util.Option[*ChallengeVertex]
	leafVertexCount        uint64
	creationTime           time.Time
	includedHistories      map[common.Hash]bool
	currentVertexSeqNumber protocol.VertexSequenceNumber
	challengePeriod        time.Duration
	challengeType          protocol.ChallengeType
}

// CreateChallenge creates a challenge for the assertion and moves the assertion to `ChallengedAssertionState` state.
func (a *AssertionChain) CreateSuccessionChallenge(
	ctx context.Context,
	tx protocol.ActiveTx,
	seqNum protocol.AssertionSequenceNumber,
) (protocol.Challenge, error) {
	tx.VerifyReadWrite()
	if a.status != PendingAssertionState && a.challengeManager.LatestConfirmed(tx) != a {
		return nil, errors.Wrapf(ErrWrongState, fmt.Sprintf("State: %d, Confirmed status: %v", a.status, a.challengeManager.LatestConfirmed(tx) != a))
	}
	if !a.challenge.IsNone() {
		return nil, ErrChallengeAlreadyExists
	}
	if a.secondChildCreationTime.IsNone() {
		return nil, ErrInvalidOp
	}
	currSeqNumber := protocol.VertexSequenceNumber(0)
	rootVertex := &ChallengeVertex{
		Challenge:   util.None[protocol.Challenge](),
		SequenceNum: currSeqNumber,
		isLeaf:      false,
		Status:      ConfirmedAssertionState,
		Commitment: util.HistoryCommitment{
			Height: 0,
			Merkle: common.Hash{},
		},
		Prev:                 util.None[protocol.ChallengeVertex](),
		PresumptiveSuccessor: util.None[protocol.ChallengeVertex](),
		PsTimer:              util.NewCountUpTimer(a.challengeManager.TimeReference()),
		SubChallenge:         util.None[protocol.Challenge](),
	}

	chal := &Challenge{
		rootAssertion:     util.Some(a),
		WinnerAssertion:   util.None[*Assertion](),
		WinnerV:           util.None[*ChallengeVertex](),
		rootVertex:        util.Some(rootVertex),
		creationTime:      a.challengeManager.TimeReference().Get(),
		includedHistories: make(map[common.Hash]bool),
		challengePeriod:   a.challengeManager.ChallengePeriodLength(tx),
		challengeType:     BlockChallenge,
	}
	rootVertex.Challenge = util.Some(protocol.Challenge(chal))
	chal.includedHistories[rootVertex.Commitment.Hash()] = true
	a.challenge = util.Some(chal)
	parentStaker := common.Address{}
	if !a.Staker.IsNone() {
		parentStaker = a.Staker.Unwrap()
	}
	a.challengeManager.GetFeed().Append(&StartChallengeEvent{
		ParentSeqNum:          a.SequenceNum,
		ParentStateCommitment: a.StateCommitment,
		ParentStaker:          parentStaker,
		Validator:             validator,
	})

	challengeID := ChallengeCommitHash(a.StateCommitment.Hash())
	a.challengeManager.GetChallengesByCommitHash()[challengeID] = chal
	a.challengeManager.GetChallengeVerticesByCommitHashmap()[challengeID] = map[VertexCommitHash]protocol.ChallengeVertex{VertexCommitHash(rootVertex.Commitment.Hash()): rootVertex}

	return chal, nil
}

// Parentutil.StateCommitment returns the state commitment of the parent assertion.
func (c *Challenge) ParentStateCommitment(ctx context.Context, tx protocol.ActiveTx) (util.StateCommitment, error) {
	if c.rootAssertion.IsNone() {
		return util.StateCommitment{}, nil
	}
	return c.rootAssertion.Unwrap().StateCommitment, nil
}

// AssertionSeqNumber returns the sequence number of the assertion that created the challenge.
func (c *Challenge) AssertionSeqNumber(ctx context.Context, tx protocol.ActiveTx) (protocol.AssertionSequenceNumber, error) {
	return c.rootAssertion.Unwrap().SequenceNum, nil
}

// RootAssertion returns the root assertion of challenge
func (c *Challenge) RootAssertion(ctx context.Context, tx protocol.ActiveTx) (protocol.Assertion, error) {
	return c.rootAssertion.Unwrap(), nil
}

// RootVertex returns the root vertex of challenge
func (c *Challenge) RootVertex(ctx context.Context, tx protocol.ActiveTx) (protocol.ChallengeVertex, error) {
	return c.rootVertex.Unwrap(), nil
}

func (c *Challenge) WinnerVertex(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeVertex], error) {
	return c.WinnerV, nil
}

func (c *Challenge) WinningClaim() util.Option[protocol.AssertionHash] {
	return [32]byte{}
}

// AddLeaf adds a new leaf to the challenge.
func (c *Challenge) AddLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	assertion *Assertion,
	history util.HistoryCommitment,
	validator common.Address,
) (protocol.ChallengeVertex, error) {
	tx.verifyReadWrite()
	if assertion.Prev.IsNone() {
		return nil, ErrInvalidOp
	}
	prev := assertion.Prev.Unwrap()
	if prev != c.rootAssertion.Unwrap() {
		return nil, ErrInvalidOp
	}
	if completed, _ := c.Completed(ctx, tx); completed {
		return nil, ErrWrongState
	}
	if eligibleForNewSuccessor, _ := c.rootVertex.Unwrap().EligibleForNewSuccessor(ctx, tx); !eligibleForNewSuccessor {
		return nil, ErrPastDeadline
	}
	if c.includedHistories[history.Hash()] {
		return nil, errors.Wrapf(ErrVertexAlreadyExists, fmt.Sprintf("Hash: %s", history.Hash().String()))
	}
	if err := c.rootAssertion.Unwrap().challengeManager.DeductFromBalance(tx, validator, ChallengeVertexStake); err != nil {
		return nil, errors.Wrapf(ErrInsufficientBalance, err.Error())
	}

	if !historyProvidesLastLeafProof(history) {
		return nil, ErrNoLastLeafProof
	}

	// The first leaf in the history commitment must be the
	// same as the previous vertex's history state root.
	if prev.StateCommitment.Height != 0 && prev.StateCommitment.StateRoot != history.FirstLeaf {
		return nil, ErrWrongFirstLeaf
	}

	// The last leaf claimed in the history commitment must be the
	// state root of the assertion we are adding a leaf for.
	if assertion.StateCommitment.StateRoot != history.LastLeaf {
		return nil, ErrWrongLastLeaf
	}

	// Assert the history commitment's height is equal to the
	// assertion.height - assertion.prev.height
	if prev.StateCommitment.Height >= assertion.StateCommitment.Height {
		return nil, errors.Wrapf(
			ErrInvalidHeight,
			"previous assertion's height %d, must be less than %d",
			prev.StateCommitment.Height,
			assertion.StateCommitment.Height,
		)
	}
	expectedHeight := assertion.StateCommitment.Height - prev.StateCommitment.Height
	if history.Height != expectedHeight {
		return nil, errors.Wrapf(
			ErrInvalidHeight,
			"history height does not match expected value %d != %d",
			history.Height,
			expectedHeight,
		)
	}

	// The validator must provide a history commitment over
	// a series of states where the last state must be proven to be
	// one corresponding to the assertion specified.
	if err := util.VerifyPrefixProof(
		history.LastLeafPrefix.Unwrap(),
		history.Normalized().Unwrap(),
		history.LastLeafProof,
	); err != nil {
		return nil, ErrProofFailsToVerify
	}

	challengeManager := assertion.challengeManager
	timer := util.NewCountUpTimer(challengeManager.TimeReference())
	if assertion.isFirstChild {
		delta := prev.secondChildCreationTime.Unwrap().Sub(prev.firstChildCreationTime.Unwrap())
		timer.Set(delta)
	}
	nextSeqNumber := c.currentVertexSeqNumber + 1
	leaf := &ChallengeVertex{
		Challenge:            util.Some(protocol.Challenge(c)),
		SequenceNum:          nextSeqNumber,
		Validator:            validator,
		isLeaf:               true,
		Status:               PendingAssertionState,
		Commitment:           history,
		Prev:                 c.rootVertex,
		PresumptiveSuccessor: util.None[protocol.ChallengeVertex](),
		PsTimer:              timer,
		SubChallenge:         util.None[protocol.Challenge](),
		winnerIfConfirmed:    util.Some(assertion),
	}
	c.currentVertexSeqNumber = nextSeqNumber
	err := c.rootVertex.Unwrap().(*ChallengeVertex).maybeNewPresumptiveSuccessor(ctx, tx, leaf)
	if err != nil {
		return nil, err
	}
	parentSeqNum, _ := leaf.Prev.Unwrap().GetSequenceNum(ctx, tx)
	c.rootAssertion.Unwrap().challengeManager.GetChallengesFeed().Append(&ChallengeLeafEvent{
		ParentSeqNum:      parentSeqNum,
		SequenceNum:       leaf.SequenceNum,
		WinnerIfConfirmed: assertion.SequenceNum,
		History:           history,
		BecomesPS:         leaf.Prev.Unwrap().(*ChallengeVertex).PresumptiveSuccessor.Unwrap() == leaf,
		Validator:         validator,
	})
	c.includedHistories[history.Hash()] = true
	h := ChallengeCommitHash(c.rootAssertion.Unwrap().StateCommitment.Hash())
	c.rootAssertion.Unwrap().challengeManager.GetChallengesByCommitHash()[h] = c
	c.rootAssertion.Unwrap().challengeManager.GetChallengeVerticesByCommitHashmap()[h][VertexCommitHash(leaf.Commitment.Hash())] = leaf
	c.leafVertexCount++

	return leaf, nil
}

// AddBlockChallengeLeaf --
func (c *Challenge) AddBlockChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	assertion protocol.Assertion,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	return nil, errors.New("unimplemented")
}

// AddBigStepChallengeLeaf adds a big step leaf to a subchallenge.
func (c *Challenge) AddBigStepChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	vertex protocol.ChallengeVertex,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	return nil, errors.New("unimplemented")
}

// Completed returns true if the challenge is completed.
func (c *Challenge) Completed(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	tx.VerifyRead()
	return !c.WinnerAssertion.IsNone(), nil
}

// CreationTime returns the time the challenge was created.
func (c *Challenge) GetCreationTime(ctx context.Context, tx protocol.ActiveTx) (time.Time, error) {
	tx.VerifyRead()
	return c.creationTime, nil
}

func (c *Challenge) GetType() protocol.ChallengeType {
	return c.challengeType
}

// Winner returns the winning assertion if the challenge is completed.
func (c *Challenge) Winner(ctx context.Context, tx *ActiveTx) (*Assertion, error) {
	tx.verifyRead()
	if c.WinnerAssertion.IsNone() {
		return nil, ErrNoWinnerYet
	}
	return c.WinnerAssertion.Unwrap(), nil
}

// HasConfirmedSibling returns true if another sibling vertex is confirmed.
func (v *ChallengeVertex) HasConfirmedSibling(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	tx.verifyRead()

	if c.rootAssertion.IsNone() {
		return false, nil
	}
	parentStateCommitment, _ := c.ParentStateCommitment(ctx, tx)
	vertices, ok := c.rootAssertion.Unwrap().challengeManager.GetChallengeVerticesByCommitHashmap()[protocol.ChallengeHash(parentStateCommitment.Hash())]
	if !ok {
		return false, nil
	}

	prevVertex, _ := vertex.GetPrev(ctx, tx)
	if prevVertex.IsNone() {
		return false, nil
	}
	commitment, _ := prevVertex.Unwrap().GetCommitment(ctx, tx)
	parentHash := commitment.Hash()
	for _, v := range vertices {
		prev, err := v.GetPrev(ctx, tx)
		if err != nil {
			return false, err
		}
		if prev.IsNone() {
			continue
		}
		// We only check vertices that have a matching parent commit hash.
		if commitment.Hash() == parentHash {
			var status AssertionState
			status, err = v.GetStatus(ctx, tx)
			if err != nil {
				return false, err
			}
			if vertex != v && status == ConfirmedAssertionState {
				return true, nil
			}
		}
	}

	return false, nil
}

type ChallengeVertex struct {
	Commitment            util.HistoryCommitment
	Challenge             util.Option[protocol.Challenge]
	SequenceNumV          protocol.VertexSequenceNumber // unique within the challenge
	Validator             common.Address
	isLeaf                bool
	StatusV               protocol.AssertionState
	PrevV                 util.Option[protocol.ChallengeVertex]
	PresumptiveSuccessorV util.Option[protocol.ChallengeVertex]
	PsTimerV              *util.CountUpTimer
	SubChallenge          util.Option[protocol.Challenge]
	winnerIfConfirmed     util.Option[*Assertion]
}

func (v *ChallengeVertex) Id() [32]byte {
	return protocol.VertexHash{}
}

func (v *ChallengeVertex) SequenceNum(ctx context.Context, tx protocol.ActiveTx) (protocol.VertexSequenceNumber, error) {
	return v.SequenceNumV, nil
}

func (v *ChallengeVertex) Status(ctx context.Context, tx protocol.ActiveTx) (protocol.AssertionState, error) {
	return v.StatusV, nil
}

func (v *ChallengeVertex) Prev(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeVertex], error) {
	return v.PrevV, nil
}

func (v *ChallengeVertex) PsTimer(ctx context.Context, tx protocol.ActiveTx) (uint64, error) {
	return uint64(v.PsTimerV.Get().Seconds()), nil
}

func (v *ChallengeVertex) MiniStaker(ctx context.Context, tx protocol.ActiveTx) (common.Address, error) {
	return v.Validator, nil
}

func (v *ChallengeVertex) HistoryCommitment(ctx context.Context, tx protocol.ActiveTx) (util.HistoryCommitment, error) {
	return v.Commitment, nil
}

func (v *ChallengeVertex) PresumptiveSuccessor(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeVertex], error) {
	return v.PresumptiveSuccessorV, nil
}

// EligibleForNewSuccessor returns true if the vertex is eligible to have a new successor.
func (v *ChallengeVertex) EligibleForNewSuccessor(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	if v.PresumptiveSuccessor.IsNone() {
		return true, nil
	}
	presumptiveSuccessorPsTimer, _ := v.PresumptiveSuccessor.Unwrap().GetPsTimer(ctx, tx)
	return presumptiveSuccessorPsTimer.Get() <= v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.ChallengePeriodLength(tx), nil
}

// maybeNewPresumptiveSuccessor updates the presumptive successor if the given vertex is eligible.
func (v *ChallengeVertex) maybeNewPresumptiveSuccessor(ctx context.Context, tx protocol.ActiveTx, succ protocol.ChallengeVertex) error {
	successorCommitment, _ := succ.GetCommitment(ctx, tx)
	if !v.PresumptiveSuccessor.IsNone() {
		presumptiveSuccessorCommitment, _ := v.PresumptiveSuccessor.Unwrap().GetCommitment(ctx, tx)
		if successorCommitment.Height < presumptiveSuccessorCommitment.Height {
			presumptiveSuccessorPsTimer, _ := v.PresumptiveSuccessor.Unwrap().GetPsTimer(ctx, tx)
			presumptiveSuccessorPsTimer.Stop()
			v.PresumptiveSuccessor = util.None[protocol.ChallengeVertex]()
		}
	}

	if v.PresumptiveSuccessor.IsNone() {
		v.PresumptiveSuccessor = util.Some(succ)
		successorPsTimer, _ := succ.GetPsTimer(ctx, tx)
		successorPsTimer.Start()
	}
	return nil
}

// IsPresumptiveSuccessor returns true if the vertex is the presumptive successor of its parent.
func (v *ChallengeVertex) IsPresumptiveSuccessor(ctx context.Context, tx protocol.ActiveTx) (bool, error) {
	if v.Prev.IsNone() {
		return true, nil
	}
	prevPresumptiveSuccessor, _ := v.Prev.Unwrap().GetPresumptiveSuccessor(ctx, tx)
	return prevPresumptiveSuccessor.Unwrap() == v, nil
}

// requiredBisectionHeight returns the height of the history commitment that must be bisectioned to prove the vertex.
func (v *ChallengeVertex) requiredBisectionHeight(ctx context.Context, tx protocol.ActiveTx) (uint64, error) {
	prevCommitment, _ := v.Prev.Unwrap().GetCommitment(ctx, tx)
	return util.BisectionPoint(prevCommitment.Height, v.Commitment.Height)
}

func (v *ChallengeVertex) Bisect(
	ctx context.Context,
	tx protocol.ActiveTx,
	history util.HistoryCommitment,
	proof []common.Hash,
) (protocol.ChallengeVertex, error) {
	tx.VerifyReadWrite()
	isPresumptiveSuccessor, _ := v.IsPresumptiveSuccessor(ctx, tx)
	if isPresumptiveSuccessor {
		return nil, ErrWrongState
	}
	prevEligibleForNewSuccessor, _ := v.Prev.Unwrap().EligibleForNewSuccessor(ctx, tx)
	if !prevEligibleForNewSuccessor {
		return nil, ErrPastDeadline
	}
	if v.Challenge.Unwrap().(*Challenge).includedHistories[history.Hash()] {
		return nil, errors.Wrapf(ErrVertexAlreadyExists, fmt.Sprintf("Hash: %s", history.Hash().String()))
	}
	bisectionHeight, err := v.requiredBisectionHeight(ctx, tx)
	if err != nil {
		return nil, err
	}
	if bisectionHeight != history.Height {
		return nil, errors.Wrapf(ErrInvalidHeight, fmt.Sprintf("%d != %v", bisectionHeight, history))
	}
	if err = util.VerifyPrefixProof(history, v.Commitment, proof); err != nil {
		return nil, err
	}

	v.PsTimer.Stop()
	nextSeqNum := v.SequenceNum + 1
	newVertex := &ChallengeVertex{
		Challenge:            v.Challenge,
		SequenceNum:          nextSeqNum,
		Validator:            validator,
		isLeaf:               false,
		Commitment:           history,
		Prev:                 v.Prev,
		PresumptiveSuccessor: util.None[protocol.ChallengeVertex](),
		PsTimer:              v.PsTimer.Clone(),
	}
	v.SequenceNum = nextSeqNum
	err = newVertex.maybeNewPresumptiveSuccessor(ctx, tx, v)
	if err != nil {
		return nil, err
	}
	err = newVertex.Prev.Unwrap().(*ChallengeVertex).maybeNewPresumptiveSuccessor(ctx, tx, newVertex)
	if err != nil {
		return nil, err
	}
	newVertex.Challenge.Unwrap().(*Challenge).includedHistories[history.Hash()] = true

	v.Prev = util.Some[protocol.ChallengeVertex](newVertex)

	newVertex.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.GetChallengesFeed().Append(&ChallengeBisectEvent{
		FromSequenceNum: v.SequenceNum,
		SequenceNum:     newVertex.SequenceNum,
		ToHistory:       newVertex.Commitment,
		FromHistory:     v.Commitment,
		BecomesPS:       newVertex.Prev.Unwrap().(*ChallengeVertex).PresumptiveSuccessor.Unwrap() == newVertex,
		Validator:       validator,
	})
	commitHash := ChallengeCommitHash(newVertex.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().StateCommitment.Hash())
	newVertex.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.GetChallengeVerticesByCommitHashmap()[commitHash][VertexCommitHash(newVertex.Commitment.Hash())] = newVertex

	return newVertex, nil
}

// Merge merges the vertex with its presumptive successor.
func (v *ChallengeVertex) Merge(
	ctx context.Context,
	tx protocol.ActiveTx,
	mergingTo util.HistoryCommitment,
	proof []common.Hash,
) (protocol.ChallengeVertex, error) {
	tx.verifyReadWrite()
	eligibleForNewSuccessor, _ := mergingTo.EligibleForNewSuccessor(ctx, tx)
	if !eligibleForNewSuccessor {
		return ErrPastDeadline
	}
	// The vertex we are merging to should be the mandatory bisection point
	// of the current vertex's height and its parent's height.
	prevCommitment, _ := v.Prev.Unwrap().GetCommitment(ctx, tx)
	bisectionPoint, err := util.BisectionPoint(prevCommitment.Height, v.Commitment.Height)
	if err != nil {
		return err
	}
	mergingToCommitment, _ := mergingTo.GetCommitment(ctx, tx)
	if mergingToCommitment.Height != bisectionPoint {
		return errors.Wrapf(ErrInvalidHeight, "%d != %d", mergingToCommitment.Height, bisectionPoint)
	}
	if err = util.VerifyPrefixProof(mergingToCommitment, v.Commitment, proof); err != nil {
		return err
	}

	v.Prev = util.Some(mergingTo)
	mergingToPsTimer, _ := mergingTo.GetPsTimer(ctx, tx)
	mergingToPsTimer.Add(v.PsTimer.Get())
	err = mergingTo.(*ChallengeVertex).maybeNewPresumptiveSuccessor(ctx, tx, v)
	if err != nil {
		return err
	}
	v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.GetChallengesFeed().Append(&ChallengeMergeEvent{
		DeeperSequenceNum:    v.SequenceNum,
		ShallowerSequenceNum: mergingTo.(*ChallengeVertex).SequenceNum,
		BecomesPS:            mergingTo.(*ChallengeVertex).PresumptiveSuccessor.Unwrap() == v,
		ToHistory:            mergingTo.(*ChallengeVertex).Commitment,
		FromHistory:          v.Commitment,
		Validator:            validator,
	})
	return nil
}

// ConfirmForSubChallengeWin confirms the vertex as the winner of a sub-challenge.
func (v *ChallengeVertex) ConfirmForSubChallengeWin(ctx context.Context, tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if v.Status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("Status: %d", v.Status))
	}
	prevStatus, _ := v.Prev.Unwrap().GetStatus(ctx, tx)
	if prevStatus != ConfirmedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", prevStatus))
	}
	preSubChal, _ := v.Prev.Unwrap().GetSubChallenge(ctx, tx)
	if preSubChal.IsNone() {
		return ErrInvalidOp
	}
	preSubChalWinnerVertex, _ := preSubChal.Unwrap().GetWinnerVertex(ctx, tx)
	if preSubChalWinnerVertex.IsNone() {
		return ErrInvalidOp
	}
	winnerVertex := preSubChalWinnerVertex.Unwrap()
	if winnerVertex != v {
		return ErrInvalidOp
	}
	v._confirm(tx)
	return nil
}

// ConfirmForPsTimer confirms the vertex as the Winner of a PsTimer.
func (v *ChallengeVertex) ConfirmForPsTimer(ctx context.Context, tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if v.Status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("Status: %d", v.Status))
	}
	prevSubChallenge, _ := v.Prev.Unwrap().GetSubChallenge(ctx, tx)
	if !prevSubChallenge.IsNone() {
		return errors.Wrap(ErrInvalidOp, "predecessor contains sub-challenge")
	}
	prevStatus, _ := v.Prev.Unwrap().GetStatus(ctx, tx)
	if prevStatus != ConfirmedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", prevStatus))
	}
	if v.PsTimer.Get() <= v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.ChallengePeriodLength(tx) {
		return errors.Wrapf(
			ErrNotYet,
			fmt.Sprintf(
				"%d <= %d",
				v.PsTimer.Get(),
				v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.ChallengePeriodLength(tx)),
		)
	}
	v._confirm(tx)
	return nil
}

// ConfirmForChallengeDeadline confirms the vertex as the winner of a challenge deadline.
func (v *ChallengeVertex) ConfirmForChallengeDeadline(ctx context.Context, tx protocol.ActiveTx) error {
	tx.verifyReadWrite()
	if v.Status != PendingAssertionState {
		return errors.Wrapf(ErrWrongState, fmt.Sprintf("Status: %d", v.Status))
	}
	if !v.Prev.Unwrap().(*ChallengeVertex).SubChallenge.IsNone() {
		return errors.Wrap(ErrInvalidOp, "predecessor contains sub-challenge")
	}
	if v.Prev.Unwrap().(*ChallengeVertex).Status != ConfirmedAssertionState {
		return errors.Wrapf(ErrWrongPredecessorState, fmt.Sprintf("State: %d", v.Prev.Unwrap().(*ChallengeVertex).Status))
	}
	if v != v.Prev.Unwrap().(*ChallengeVertex).PresumptiveSuccessor.Unwrap() {
		return errors.Wrap(ErrInvalidOp, "Vertex is not the presumptive successor")
	}
	challengeManger := v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager
	chalPeriod := challengeManger.ChallengePeriodLength(tx)
	if !challengeManger.TimeReference().Get().After(v.Challenge.Unwrap().(*Challenge).creationTime.Add(2 * chalPeriod)) {
		return errors.Wrapf(
			ErrNotYet, fmt.Sprintf(
				"%d <= %d",
				challengeManger.TimeReference().Get().Unix(),
				v.Challenge.Unwrap().(*Challenge).creationTime.Add(2*chalPeriod).Unix(),
			),
		)
	}
	v._confirm(tx)
	return nil
}

func (v *ChallengeVertex) _confirm(tx *ActiveTx) {
	v.Status = ConfirmedAssertionState
	if v.isLeaf {
		refund := big.NewInt(0)
		leafCount := int64(v.Challenge.Unwrap().(*Challenge).leafVertexCount)
		refund.Mul(ChallengeVertexStake, big.NewInt(leafCount+1))
		refund.Div(refund, big.NewInt(2))
		v.Challenge.Unwrap().(*Challenge).rootAssertion.Unwrap().challengeManager.AddToBalance(tx, v.Validator, refund)
		v.Challenge.Unwrap().(*Challenge).WinnerAssertion = v.winnerIfConfirmed
	}
}

func (v *ChallengeVertex) GetPrev(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeVertex], error) {
	return v.Prev, nil
}
func (v *ChallengeVertex) GetStatus(ctx context.Context, tx protocol.ActiveTx) (AssertionState, error) {
	return v.Status, nil
}
func (v *ChallengeVertex) GetSubChallenge(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.Challenge], error) {
	return v.SubChallenge, nil
}
func (v *ChallengeVertex) GetPsTimer(ctx context.Context, tx protocol.ActiveTx) (*util.CountUpTimer, error) {
	return v.PsTimer, nil
}
func (v *ChallengeVertex) GetCommitment(ctx context.Context, tx protocol.ActiveTx) (util.HistoryCommitment, error) {
	return v.Commitment, nil
}
func (v *ChallengeVertex) GetValidator(ctx context.Context, tx protocol.ActiveTx) (common.Address, error) {
	return v.Validator, nil
}
func (v *ChallengeVertex) GetSequenceNum(ctx context.Context, tx protocol.ActiveTx) (protocol.VertexSequenceNumber, error) {
	return v.SequenceNum, nil
}
func (v *ChallengeVertex) GetPresumptiveSuccessor(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.ChallengeVertex], error) {
	return v.PresumptiveSuccessor, nil
}

func historyProvidesLastLeafProof(history util.HistoryCommitment) bool {
	return history.LastLeaf != (common.Hash{}) &&
		len(history.LastLeafProof) != 0 &&
		!history.LastLeafPrefix.IsNone() &&
		!history.Normalized().IsNone()
}
