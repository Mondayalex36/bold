package validator

import (
	"context"
	"go/types"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Each challenge has a lifecycle we need to manage. A single challenge's entire lifecycle should
// be managed in a goroutine specific to that challenge. A challenge goroutine will exit if
//
// - A winner has been found (meaning all subchallenges are resolved), or
// - The validator's chess clock times out
//
// The validator has is able to dispatch events from the global feed
// to specific challenge goroutines. A challenge goroutine is spawned upon receiving
// a ChallengeStarted event. Each challenge goroutine is managed by a challenge worker struct
// which has enough information about the challenge to make respective moves on its
// associated events coming from the protocol.
type blockChallengeWorker struct {
	challenge          *protocol.Challenge
	validatorAddress   common.Address
	reachedOneStepFork chan struct{}
	validatorName      string
	createdVertices    *util.ThreadSafeSlice[*protocol.ChallengeVertex]
	events             chan protocol.ChallengeEvent
}

// Performs the actions required by a validator when a ChallengeEvent is fired during
// a BlockChallenge. The basic algorith is as follows:
//
// 1. We fetch our last created vertex that is higher than the commitment of the vertex seen
// during the challenge event.
// 2. If we have the same history commitment, and the vertex is not ours, we make a merge move.
// 3. While we are not presumptive, we keep trying to merge until we are presumptive or we reach a one-step-fork.
// /
func (w *blockChallengeWorker) actOnBlockChallenge(
	ctx context.Context,
	validator *Validator,
	eventActor common.Address,
	eventHistoryCommit util.HistoryCommitment,
	eventSequenceNum protocol.VertexSequenceNumber,
) error {
	if isFromSelf(w.validatorAddress, eventActor) {
		return nil
	}
	if w.createdVertices.Empty() {
		return nil
	}
	// Go down the tree to find the first vertex we created that has a commitment height >
	// the vertex seen from the merge event.
	vertexToActUpon := w.loadVertexToActUpon(eventHistoryCommit.Height)

	// If we should make a merge move, do that first.
	if w.shouldMakeMergeMove(ctx, validator, eventHistoryCommit, vertexToActUpon.Commitment) {
		challengeID := protocol.CommitHash(w.challenge.ParentStateCommitment().Hash())
		mergingFrom := vertexToActUpon
		mergingTo, err := w.vertexToMergeInto(validator, challengeID, eventSequenceNum)
		if err != nil {
			return err
		}
		if err := validator.merge(ctx, mergingTo, mergingFrom); err != nil {
			// TODO: Find a better way to exit if a merge is invalid than showing a scary log to the user.
			// Validators currently try to make merge moves they should not during the challenge game.
			if errors.Is(err, protocol.ErrInvalidOp) {
				return nil
			}
			return errors.Wrap(err, "failed to merge")
		}
	}

	// While we do not have the presumptive successor, we keep trying to bisect and
	// until we either reach a one-step-fork, or we have the presumptive successor, in which
	// case other validators should be the ones to act next.
	return w.bisectWhileNonPresumptive(ctx, validator, vertexToActUpon)
}

// Gets the vertex a validator has created that we should act upon when receiving
// a ChallengeEvent during a BlockChallenge goroutine. For example, if we see a bisection event from
// another validator, we should find our last created vertex that we should act upon in response.
// Currently, this retrieves the last vertex we created with a height > event.HistoryCommitment.Height
// NOTE: This expects the list of created vertices we be non-empty.
// TODO: Determine if this is a good idea.
func (w *blockChallengeWorker) loadVertexToActUpon(eventHeight uint64) *protocol.ChallengeVertex {
	// Go down the tree to find the first vertex we created that has a
	// HistoryCommit.Height > event.HistoryCommit.Height.
	vertexToActUpon := w.createdVertices.Last().Unwrap()
	numVertices := w.createdVertices.Len()
	for i := numVertices - 1; i > 0; i-- {
		vertex := w.createdVertices.Get(i).Unwrap()
		if vertex.Commitment.Height > eventHeight {
			return vertex
		}
	}
	return vertexToActUpon
}

// If the event is firing a history commitment that we have, and the event is not a merge
// move to one of our vertices, we should attempt to merge to that event's vertex.
func (w *blockChallengeWorker) shouldMakeMergeMove(
	ctx context.Context,
	validator *Validator,
	incomingEventCommit,
	ourVertexCommit util.HistoryCommitment,
) bool {
	mergedToOurs := incomingEventCommit.Hash() == ourVertexCommit.Hash()
	if mergedToOurs {
		log.WithFields(logrus.Fields{
			"name":                w.validatorName,
			"mergedHeight":        incomingEventCommit.Height,
			"mergedHistoryMerkle": incomingEventCommit.Merkle,
		}).Info("Other validator merged to our vertex")
	}
	return validator.stateManager.HasHistoryCommitment(ctx, incomingEventCommit) && !mergedToOurs
}

// Obtains a challenge vertex we should perform a merge move into given its corresponding challenge ID
// and the sequence number of the vertex itself (obtained through a fired event).
func (w *blockChallengeWorker) vertexToMergeInto(
	validator *Validator,
	vertexChallengeID protocol.CommitHash,
	vertexSeqNum protocol.VertexSequenceNumber,
) (*protocol.ChallengeVertex, error) {
	var mergingTo *protocol.ChallengeVertex
	var err error
	if err = validator.chain.Call(func(tx *protocol.ActiveTx, p protocol.OnChainProtocol) error {
		mergingTo, err = p.ChallengeVertexBySequenceNum(tx, vertexChallengeID, vertexSeqNum)
		if err != nil {
			return err
		}
		return nil

	}); err != nil {
		return nil, err
	}
	return mergingTo, nil
}

// When a validator needs to make a bisection move, it should continue making bisection moves until
// it either has the presumptive successor vertex or it reaches a one-step fork. If it has the presumptive
// successor for the challenge, it will yield action to other validators in the protocol. If it reaches
// a one-step-fork, it will fire an event over a channel and exit. This one-step-fork should
// then be handled by a corresponding BigStepChallenge.
func (w *blockChallengeWorker) bisectWhileNonPresumptive(
	ctx context.Context,
	validator *Validator,
	startVertex *protocol.ChallengeVertex,
) error {
	current := startVertex
	hasPresumptiveSuccessor := startVertex.IsPresumptiveSuccessor()
	// While we do not have the presumptive successor, we keep trying to bisect and
	// break from the loop if reach a one step fork.
	for !hasPresumptiveSuccessor {
		bisectedVertex, err := validator.bisect(ctx, current)
		if err != nil {
			// TODO: Find another way of cleanly ending the bisection process so that we do not
			// end on a scary "state did not allow this operation" log.
			if errors.Is(err, protocol.ErrWrongState) {
				log.WithError(err).Debug("State incorrect for bisection")
				return nil
			}
			if errors.Is(err, protocol.ErrVertexAlreadyExists) {
				return nil
			}
			return err
		}
		current = bisectedVertex
		hasPresumptiveSuccessor = current.IsPresumptiveSuccessor()
		w.createdVertices.Append(current)

		// If we have reached a one-step-fork, we send a notification to a channel.
		if current.Commitment.Height == current.Prev.Commitment.Height+1 {
			w.reachedOneStepFork <- struct{}{}
			return nil
		}
	}
	return nil
}
