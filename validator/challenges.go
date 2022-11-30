package validator

import (
	"context"
	"fmt"
	"sync"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	statemanager "github.com/OffchainLabs/new-rollup-exploration/state-manager"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// TODO: A challenge ID in our challenge manager is a unique identifier for that challenge.
// Currently, it is using the parent assertion's state commitment hash as an ID,
// but this value may not work if we want to manage subchallenges in the same way.
type challengeID common.Hash

// Each challenge has a lifecycle we need to manage. A single challenge's entire lifecycle should
// be managed in a goroutine specific to that challenge. A challenge goroutine will exit if
//
// (a) A winner has been found (meaning all subchallenges are resolved), or
// (b) The chess clock for the challenge ends
//
// The validator has a challenge manager which is able to dispatch events from the global feed
// to specific challenge goroutines. A challenge goroutine is spawned upon receiving
// a ChallengeStarted event.
type challengeManager struct {
	lock                   sync.RWMutex
	chain                  protocol.ChainReadWriter
	stateManager           statemanager.Manager
	validatorAddress       common.Address
	validatorName          string
	challenges             map[challengeID]*challengeWorker
	challengeEventsBufSize int
}

type challengeWorker struct {
	lock             sync.RWMutex
	challenge        *protocol.Challenge
	validatorAddress common.Address
	validatorName    string
	createdVertices  []*protocol.ChallengeVertex
	events           chan protocol.ChallengeEvent
}

func newChallengeManager(
	chain protocol.ChainReadWriter,
	stateManager statemanager.Manager,
	validatorAddress common.Address,
	validatorName string,
) *challengeManager {
	return &challengeManager{
		chain:                  chain,
		stateManager:           stateManager,
		challenges:             make(map[challengeID]*challengeWorker),
		validatorAddress:       validatorAddress,
		validatorName:          validatorName,
		challengeEventsBufSize: 100, // TODO: Make configurable.
	}
}

func (c *challengeManager) numChallenges() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.challenges)
}

func (c *challengeManager) spawnChallenge(
	ctx context.Context,
	challenge *protocol.Challenge,
	vertex *protocol.ChallengeVertex,
) {
	c.lock.Lock()
	ch := make(chan protocol.ChallengeEvent, c.challengeEventsBufSize)
	c.chain.SubscribeChallengeEvents(ctx, ch)
	id := challenge.ParentStateCommitment().Hash()
	if _, ok := c.challenges[challengeID(id)]; ok {
		c.lock.Unlock()
		log.WithFields(logrus.Fields{
			"challengeID": fmt.Sprintf("%#x", id),
			"name":        c.validatorName,
		}).Error("Attempted to spawn challenge that is already in progress")
		return
	}
	worker := &challengeWorker{
		challenge:        challenge,
		createdVertices:  []*protocol.ChallengeVertex{vertex},
		validatorAddress: c.validatorAddress,
		validatorName:    c.validatorName,
		events:           ch,
	}
	c.challenges[challengeID(id)] = worker
	c.lock.Unlock()
	log.WithFields(logrus.Fields{
		"challengeID": fmt.Sprintf("%#x", id),
		"name":        c.validatorName,
	}).Info("Spawning challenge lifecycle manager")
	go worker.runChallengeLifecycle(ctx, c, ch)
}

func (w *challengeWorker) runChallengeLifecycle(
	ctx context.Context,
	manager *challengeManager,
	evs chan protocol.ChallengeEvent,
) {
	// TODO: Figure out if we are at a one-step fork, and then depending on who's turn it is,
	// spawn a subchallenge (BigStepChallenge).
	// defer close(evs)
	reachedOneStepFork := make(chan struct{})
	defer close(reachedOneStepFork)
	for {
		select {
		case genericEvent := <-evs:
			switch ev := genericEvent.(type) {
			case *protocol.ChallengeLeafEvent:
				go func() {
					if err := w.onChallengeLeafAdded(ctx, manager, ev); err != nil {
						log.WithError(err).Error("Could not process challenge leaf added event")
					}
				}()
			case *protocol.ChallengeBisectEvent:
				go func() {
					if err := w.onBisectionEvent(ctx, manager, ev); err != nil {
						log.WithError(err).Error("Could not process bisection event")
					}
				}()
			case *protocol.ChallengeMergeEvent:
				go func() {
					if err := w.onMergeEvent(ctx, ev); err != nil {
						log.WithError(err).Error("Could not process merge event")
					}
				}()
			default:
				log.WithField("ev", fmt.Sprintf("%+v", ev)).Error("Not a recognized challenge event")
			}
		case <-reachedOneStepFork:
			// TODO: Trigger subchallenge!
			return
		case <-ctx.Done():
			return
		}
	}
}

// If a leaf has been added, we check if we should add bisect.
func (w *challengeWorker) onChallengeLeafAdded(
	ctx context.Context, manager *challengeManager, ev *protocol.ChallengeLeafEvent,
) error {
	if isFromSelf(w.validatorAddress, ev.Challenger) {
		return nil
	}
	// If we have the history commitment the new leaf claims to have, we do not need to act.
	// TODO: Check if this is the correct assumption.
	// TODO: Should we just return nil here or do something else?
	if manager.stateManager.HasHistoryCommitment(ctx, ev.History) {
		return nil
	}

	w.lock.Lock()
	if len(w.createdVertices) == 0 {
		w.lock.Unlock()
		return nil
	}
	validatorChallengeVertex := w.createdVertices[len(w.createdVertices)-1]
	w.lock.Unlock()

	hasPresumptiveSuccessor := validatorChallengeVertex.IsPresumptiveSuccessor()
	currentVertex := validatorChallengeVertex
	var err error

	// TODO: Check if we are also at a one-step fork or not.
	for !hasPresumptiveSuccessor {
		currentVertex, err = w.bisect(ctx, ev.ParentStateCommit.Height, currentVertex, manager)
		if err != nil {
			return err
		}
		w.lock.Lock()
		w.createdVertices = append(w.createdVertices, currentVertex)
		w.lock.Unlock()
	}
	return nil
}

// If a bisection has occurred, we need to determine if we should merge or bisect ourselves.
func (w *challengeWorker) onBisectionEvent(
	ctx context.Context, manager *challengeManager, ev *protocol.ChallengeBisectEvent,
) error {
	if isFromSelf(w.validatorAddress, ev.Challenger) {
		return nil
	}
	// If we agree with the history commitment, we make a merge move on the vertex.
	log.WithFields(logrus.Fields{
		"name":   w.validatorName,
		"height": ev.History.Height,
		"merkle": fmt.Sprintf("%#x", ev.History.Merkle),
	}).Infof("Received bisection event from %#x", ev.Challenger)

	// Make a merge move.
	if manager.stateManager.HasHistoryCommitment(ctx, ev.History) {
		log.Info("Agreed with bisection event from other challenger, starting a merge move")
		return nil
	}
	w.lock.Lock()
	if len(w.createdVertices) == 0 {
		w.lock.Unlock()
		return nil
	}
	validatorChallengeVertex := w.createdVertices[len(w.createdVertices)-1]
	w.lock.Unlock()

	log.WithFields(logrus.Fields{
		"name":   w.validatorName,
		"height": validatorChallengeVertex.Commitment.Height,
		"merkle": fmt.Sprintf("%#x", validatorChallengeVertex.Commitment.Merkle),
	}).Info("Disagreed with bisection event from other challenger, bisecting")

	// Bisect.
	hasPresumptiveSuccessor := validatorChallengeVertex.IsPresumptiveSuccessor()
	currentVertex := validatorChallengeVertex
	var err error

	// TODO: Check if we are also at a one-step fork or not.
	for !hasPresumptiveSuccessor {
		currentVertex, err = w.bisect(ctx, ev.ParentStateCommit.Height, currentVertex, manager)
		if err != nil {
			return err
		}
		w.lock.Lock()
		w.createdVertices = append(w.createdVertices, currentVertex)
		w.lock.Unlock()
	}
	return nil
}

// If we have seen a merge event, what happens...??
func (w *challengeWorker) onMergeEvent(ctx context.Context, ev *protocol.ChallengeMergeEvent) error {
	return nil
}

func (w *challengeWorker) bisect(
	ctx context.Context,
	parentHeight uint64,
	validatorChallengeVertex *protocol.ChallengeVertex,
	manager *challengeManager,
) (*protocol.ChallengeVertex, error) {
	toHeight := validatorChallengeVertex.Commitment.Height
	bisectTo, err := util.BisectionPoint(parentHeight, toHeight)
	if err != nil {
		return nil, err
	}
	historyCommit, err := manager.stateManager.HistoryCommitmentUpTo(ctx, bisectTo)
	if err != nil {
		return nil, err
	}
	proof, err := manager.stateManager.PrefixProof(ctx, bisectTo, toHeight)
	if err != nil {
		return nil, err
	}
	if err := util.VerifyPrefixProof(historyCommit, validatorChallengeVertex.Commitment, proof); err != nil {
		return nil, err
	}
	// Otherwise, we must bisect to our own historical commitment and produce
	// a proof of the vertex we want to bisect to.
	var bisectedVertex *protocol.ChallengeVertex
	err = manager.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		bisectedVertex, err = validatorChallengeVertex.Bisect(tx, historyCommit, proof, w.validatorAddress)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"could not bisect vertex with sequence %d and challenger %#x to height %d with history %d and %#x",
			validatorChallengeVertex.SequenceNum,
			validatorChallengeVertex.Challenger,
			bisectTo,
			historyCommit.Height,
			historyCommit.Merkle,
		)
	}
	log.WithFields(logrus.Fields{
		"name":                   w.validatorName,
		"IsPresumptiveSuccessor": bisectedVertex.IsPresumptiveSuccessor(),
	}).Infof(
		"Successfully bisected to vertex with height %d and commit %#x",
		bisectedVertex.Commitment.Height,
		bisectedVertex.Commitment.Merkle,
	)
	return bisectedVertex, nil
}

// Process new challenge creation events from the protocol that were not initiated by self.
func (v *Validator) onChallengeStarted(ctx context.Context, ev *protocol.StartChallengeEvent) error {
	if ev == nil {
		return nil
	}
	// Ignore challenges initiated by self.
	if isFromSelf(v.address, ev.Challenger) {
		return nil
	}
	// Checks if the challenge has to do with a vertex we created.
	v.leavesLock.RLock()
	leaf, ok := v.createdLeaves[ev.ParentStateCommitment.StateRoot]
	if !ok {
		v.leavesLock.RUnlock()
		// TODO: Act on the honest vertices even if this challenge does not have to do with us by
		// keeping track of associated challenge vertices' clocks and acting if the associated
		// staker we agree with is not performing their responsibilities on time. As an honest
		// validator, we should participate in confirming valid assertions.
		return nil
	}
	v.leavesLock.RUnlock()
	challengerName := "unknown-name"
	if !leaf.Staker.IsEmpty() {
		if name, ok := v.knownValidatorNames[leaf.Staker.OpenKnownFull()]; ok {
			challengerName = name
		} else {
			challengerName = leaf.Staker.OpenKnownFull().Hex()
		}
	}
	log.WithFields(logrus.Fields{
		"name":                 v.name,
		"challenger":           challengerName,
		"challengingStateRoot": fmt.Sprintf("%#x", leaf.StateCommitment.StateRoot),
		"challengingHeight":    leaf.StateCommitment.Height,
	}).Warn("Received challenge for a created leaf")

	// TODO: Do we produce a historial commitment at the height == our latest height?
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return err
	}

	// We then add a leaf to the challenge using a historical commitment at our latest height.
	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, err := p.AssertionBySequenceNum(tx, ev.ParentSeqNum)
		if err != nil {
			return err
		}
		challenge, err = p.ChallengeByParentCommitmentHash(tx, parentAssertion.StateCommitment.Hash())
		if err != nil {
			return err
		}
		// TODO: What if the challenge already has a leaf we agree with?
		// TODO: Match on error ErrDuplicateLeaf to add nicer logs to the user.
		challengeVertex, err = challenge.AddLeaf(tx, parentAssertion, historyCommit, v.address)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "could not create challenge on leaf with sequence number: %d", ev.ParentSeqNum)
	}

	// Start tracking the challenge and created vertex using the challenge manager.
	v.challengeManager.spawnChallenge(ctx, challenge, challengeVertex)

	return nil
}

// Initiates a challenge on a created leaf event.
func (v *Validator) challengeLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	// Retrieves the parent assertion to begin the challenge on.
	var parentAssertion *protocol.Assertion
	var currentAssertion *protocol.Assertion
	var err error
	if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		parentAssertion, err = p.AssertionBySequenceNum(tx, ev.PrevSeqNum)
		if err != nil {
			return err
		}
		currentAssertion, err = p.AssertionBySequenceNum(tx, ev.SeqNum)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	logFields := logrus.Fields{}
	logFields["name"] = v.name
	logFields["disagreedHeight"] = ev.StateCommitment.Height
	logFields["disagreedLeaf"] = fmt.Sprintf("%#x", ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info("Initiating challenge on parent of leaf validator disagrees with")

	var challenge *protocol.Challenge
	var challengeVertex *protocol.ChallengeVertex
	err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challenge, err = parentAssertion.CreateChallenge(tx, ctx, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot make challenge")
		}
		return nil
	})
	switch {
	case errors.Is(err, protocol.ErrChallengeAlreadyExists):
		log.Info("Challenge on leaf already exists, reading existing challenge from protocol")
		if err = v.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
			challenge, err = p.ChallengeByParentCommitmentHash(tx, parentAssertion.StateCommitment.Hash())
			if err != nil {
				return errors.Wrap(err, "cannot make challenge")
			}
			return nil
		}); err != nil {
			return errors.Wrap(err, "could not get challenge by ID")
		}
	case err != nil:
		return errors.Wrapf(
			err,
			"could not initiate challenge on assertion with seq num %d",
			parentAssertion.SequenceNum,
		)
	default:
	}

	if challenge == nil {
		return errors.New("got nil challenge from protocol")
	}

	// We produce a historical commiment to add a leaf to the initiated challenge
	// by retrieving it from our local state manager.
	// TODO: Do we produce a historial commitment at the height == our latest height?
	historyCommit, err := v.stateManager.LatestHistoryCommitment(ctx)
	if err != nil {
		return err
	}

	if err = v.chain.Tx(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		challengeVertex, err = challenge.AddLeaf(tx, currentAssertion, historyCommit, v.address)
		if err != nil {
			return errors.Wrap(err, "cannot add leaf")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "could not add leaf to challenge")
	}

	// Start tracking the challenge and created vertex using the challenge manager.
	v.challengeManager.spawnChallenge(ctx, challenge, challengeVertex)

	logFields = logrus.Fields{}
	logFields["name"] = v.name
	logFields["parentAssertionSeqNum"] = parentAssertion.SequenceNum
	logFields["parentAssertionStateRoot"] = fmt.Sprintf("%#x", parentAssertion.StateCommitment.StateRoot)
	logFields["challengeID"] = fmt.Sprintf("%#x", parentAssertion.StateCommitment.Hash())
	log.WithFields(logFields).Info("Successfully created challenge and added leaf, now tracking events")

	return nil
}

// Prepares to defend a leaf that matches our local history and is part of a fork
// in the assertions tree. This leaf may be challenged and the local validator should
// be ready to perform proper challenge moves on the assertion if no one else is making them.
func (v *Validator) defendLeaf(ctx context.Context, ev *protocol.CreateLeafEvent) error {
	logFields := logrus.Fields{}
	if name, ok := v.knownValidatorNames[ev.Staker]; ok {
		logFields["createdBy"] = name
	}
	logFields["name"] = v.name
	logFields["height"] = ev.StateCommitment.Height
	logFields["stateRoot"] = fmt.Sprintf("%#x", ev.StateCommitment.StateRoot)
	log.WithFields(logFields).Info(
		"New leaf created by another validator matching local state has " +
			"forked the protocol, preparing to defend",
	)
	return nil
}
