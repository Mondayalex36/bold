package solimpl

import (
	"context"
	"math/big"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/challengeV2gen"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var ErrNoRootVertex = errors.New("root vertex not found")

func (c *Challenge) Id() protocol.ChallengeHash {
	return c.id
}

func (c *Challenge) Challenger(ctx context.Context, tx protocol.ActiveTx) (common.Address, error) {
	inner, err := c.inner(ctx, tx)
	if err != nil {
		return common.Address{}, err
	}
	return inner.Challenger, nil
}

func (c *Challenge) RootAssertion(
	ctx context.Context, tx protocol.ActiveTx,
) (protocol.Assertion, error) {
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return nil, err
	}
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	rootVertex, err := cManager.GetVertex(ctx, tx, cInner.RootId)
	if err != nil {
		return nil, err
	}
	if rootVertex.IsNone() {
		return nil, ErrNoRootVertex
	}
	root := rootVertex.Unwrap().(*ChallengeVertex)
	rootInner, err := root.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	assertionNum, err := c.chain.GetAssertionNum(ctx, tx, rootInner.ClaimId)
	if err != nil {
		return nil, err
	}
	assertion, err := c.chain.AssertionBySequenceNum(ctx, tx, assertionNum)
	if err != nil {
		return nil, err
	}
	return assertion, nil
}

// TopLevelClaimVertex gets the vertex at the BlockChallenge level that originated a subchallenge.
// For example, if two validators open a subchallenge S at vertex A in a BlockChallenge, the TopLevelClaimVertex
// of S is A. If two validators open a subchallenge S' at vertex B in BigStepChallenge, the TopLevelClaimVertex
// is vertex A.
func (c *Challenge) TopLevelClaimVertex(ctx context.Context, tx protocol.ActiveTx) (protocol.ChallengeVertex, error) {
	chalTyp, err := c.GetType(ctx, tx)
	if err != nil {
		return nil, err
	}
	if !chalTyp.IsSubChallenge() {
		return nil, errors.New("not a subchallenge")
	}
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return nil, err
	}
	rootId := cInner.RootId
	rootV, err := cManager.GetVertex(ctx, tx, rootId)
	if err != nil {
		return nil, err
	}
	if rootV.IsNone() {
		return nil, ErrNoRootVertex
	}
	root, ok := rootV.Unwrap().(*ChallengeVertex)
	if !ok {
		return nil, errors.New("root vertex is not *solimpl.ChallengeVertex type")
	}
	vInner, err := root.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	claimVertexV, err := cManager.GetVertex(ctx, tx, vInner.ClaimId)
	if err != nil {
		return nil, err
	}
	if claimVertexV.IsNone() {
		return nil, ErrNoRootVertex
	}
	claimVertex := claimVertexV.Unwrap()

	// If we are in a big step challenge, the claim vertex is the top-level vertex of the
	// corresponding BlockChallenge, so we are done.
	if chalTyp == protocol.BigStepChallenge {
		return claimVertex, nil
	}

	// Otherwise, a bit more work is required.
	// Get the root vertex of the BigStepChallenge claimVertex belongs to.
	claimVertexItem, ok := claimVertex.(*ChallengeVertex)
	if !ok {
		return nil, errors.New("claim vertex is not *solimpl.ChallengeVertex type")
	}
	claimInner, err := claimVertexItem.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	bigStepChallengeId := claimInner.ChallengeId
	bigStepC, err := cManager.GetChallenge(ctx, tx, bigStepChallengeId)
	if err != nil {
		return nil, err
	}
	bigStepChallenge, ok := bigStepC.Unwrap().(*Challenge)
	if !ok {
		return nil, errors.New("big challenge is not *solimpl.Challenge type")
	}
	bigStepChalInner, err := bigStepChallenge.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	bigStepRootV, err := cManager.GetVertex(ctx, tx, bigStepChalInner.RootId)
	if err != nil {
		return nil, err
	}
	if bigStepRootV.IsNone() {
		return nil, ErrNoRootVertex
	}
	bigStepRoot, ok := bigStepRootV.Unwrap().(*ChallengeVertex)
	if !ok {
		return nil, errors.New("big step root vertex is not *solimpl.ChallengeVertex type")
	}

	// Get the claim vertex of the BigStepChallenge's root vertex.
	bigStepRootInner, err := bigStepRoot.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	claimVertexV, err = cManager.GetVertex(ctx, tx, bigStepRootInner.ClaimId)
	if err != nil {
		return nil, err
	}
	if claimVertexV.IsNone() {
		return nil, errors.New("no claim vertex for BigStepChallenge found")
	}
	return claimVertexV.Unwrap(), nil
}

func (c *Challenge) RootVertex(
	ctx context.Context, tx protocol.ActiveTx,
) (protocol.ChallengeVertex, error) {
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return nil, err
	}
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return nil, err
	}
	rootId := cInner.RootId
	v, err := cManager.GetVertex(ctx, tx, rootId)
	if err != nil {
		return nil, err
	}
	return v.Unwrap(), nil
}

func (c *Challenge) WinningClaim(ctx context.Context, tx protocol.ActiveTx) (util.Option[protocol.AssertionHash], error) {
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return util.None[protocol.AssertionHash](), err
	}
	if cInner.WinningClaim == [32]byte{} {
		return util.None[protocol.AssertionHash](), nil
	}
	return util.Some[protocol.AssertionHash](cInner.WinningClaim), nil
}

func (c *Challenge) GetType(ctx context.Context, tx protocol.ActiveTx) (protocol.ChallengeType, error) {
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return 0, err
	}
	return protocol.ChallengeType(cInner.ChallengeType), nil
}

func (c *Challenge) GetCreationTime(
	ctx context.Context, tx protocol.ActiveTx,
) (time.Time, error) {
	return time.Time{}, errors.New("unimplemented")
}

func (c *Challenge) ParentStateCommitment(
	ctx context.Context, tx protocol.ActiveTx,
) (util.StateCommitment, error) {
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return util.StateCommitment{}, err
	}
	cInner, err := c.inner(ctx, tx)
	if err != nil {
		return util.StateCommitment{}, err
	}
	v, err := cManager.GetVertex(ctx, tx, cInner.RootId)
	if err != nil {
		return util.StateCommitment{}, err
	}
	if v.IsNone() {
		return util.StateCommitment{}, ErrNoRootVertex
	}
	concreteV, ok := v.Unwrap().(*ChallengeVertex)
	if !ok {
		return util.StateCommitment{}, errors.New("vertex is not expected concrete type")
	}
	concreteVInner, err := concreteV.inner(ctx, tx)
	if err != nil {
		return util.StateCommitment{}, err
	}
	assertionSeqNum, err := c.chain.rollup.GetAssertionNum(
		c.chain.callOpts, concreteVInner.ClaimId,
	)
	if err != nil {
		return util.StateCommitment{}, err
	}
	assertion, err := c.chain.AssertionBySequenceNum(ctx, tx, protocol.AssertionSequenceNumber(assertionSeqNum))
	if err != nil {
		return util.StateCommitment{}, err
	}
	height, err := assertion.Height()
	if err != nil {
		return util.StateCommitment{}, err
	}
	stateHash, err := assertion.StateHash()
	if err != nil {
		return util.StateCommitment{}, err
	}
	return util.StateCommitment{
		Height:    height,
		StateRoot: stateHash,
	}, nil
}

func (c *Challenge) WinnerVertex(
	ctx context.Context, tx protocol.ActiveTx,
) (util.Option[protocol.ChallengeVertex], error) {
	return util.None[protocol.ChallengeVertex](), errors.New("unimplemented")
}

func (c *Challenge) Completed(
	ctx context.Context, tx protocol.ActiveTx,
) (bool, error) {
	return false, errors.New("unimplemented")
}

// AddBlockChallengeLeaf vertex to a BlockChallenge using an assertion and a history commitment.
func (c *Challenge) AddBlockChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	assertion protocol.Assertion,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	flatLastLeafProof := make([]byte, 0, len(history.LastLeafProof)*32)
	lastLeafProof := make([][32]byte, len(history.LastLeafProof))
	for i, h := range history.LastLeafProof {
		var r [32]byte
		copy(r[:], h[:])
		flatLastLeafProof = append(flatLastLeafProof, r[:]...)
		lastLeafProof[i] = r
	}
	firstLeafProof := make([][32]byte, len(history.FirstLeafProof))
	for i, h := range history.FirstLeafProof {
		var r [32]byte
		copy(r[:], h[:])
		firstLeafProof[i] = r
	}
	callOpts := c.chain.callOpts
	assertionId, err := c.chain.rollup.GetAssertionId(callOpts, uint64(assertion.SeqNum()))
	if err != nil {
		return nil, err
	}
	leafData := challengeV2gen.AddLeafArgs{
		ChallengeId:            c.id,
		ClaimId:                assertionId,
		Height:                 big.NewInt(int64(history.Height)),
		HistoryRoot:            history.Merkle,
		FirstState:             history.FirstLeaf,
		FirstStatehistoryProof: firstLeafProof,
		LastState:              history.LastLeaf,
		LastStatehistoryProof:  lastLeafProof,
	}

	// Check the current mini-stake amount and transact using that as the value.
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return nil, err
	}
	miniStake, err := cManager.miniStakeAmount()
	if err != nil {
		return nil, err
	}
	opts := copyTxOpts(c.chain.txOpts)
	opts.Value = miniStake

	_, err = transact(ctx, c.chain.backend, c.chain.headerReader, func() (*types.Transaction, error) {
		return cManager.writer.AddLeaf(
			opts,
			leafData,
			flatLastLeafProof,
			make([]byte, 0), // Inbox proof
		)
	})
	if err != nil {
		return nil, err
	}

	vertexId, err := cManager.caller.CalculateChallengeVertexId(
		c.chain.callOpts,
		c.id,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	_, err = cManager.caller.GetVertex(
		c.chain.callOpts,
		vertexId,
	)
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		id:    vertexId,
		chain: c.chain,
	}, nil
}

// AddSubChallengeLeaf adds the appropriate leaf to the challenge based on a vertex and history commitment.
func (c *Challenge) AddSubChallengeLeaf(
	ctx context.Context,
	tx protocol.ActiveTx,
	vertex protocol.ChallengeVertex,
	history util.HistoryCommitment,
) (protocol.ChallengeVertex, error) {
	// Flatten the last leaf proof for submission to the chain.
	flatLastLeafProof := make([]byte, 0, len(history.LastLeafProof)*32)
	lastLeafProof := make([][32]byte, len(history.LastLeafProof))
	for i, h := range history.LastLeafProof {
		var r [32]byte
		copy(r[:], h[:])
		flatLastLeafProof = append(flatLastLeafProof, r[:]...)
		lastLeafProof[i] = r
	}

	firstLeafProof := make([][32]byte, len(history.FirstLeafProof))
	for i, h := range history.FirstLeafProof {
		var r [32]byte
		copy(r[:], h[:])
		firstLeafProof[i] = r
	}
	leafData := challengeV2gen.AddLeafArgs{
		ChallengeId:            c.id,
		ClaimId:                vertex.Id(),
		Height:                 big.NewInt(int64(history.Height)),
		HistoryRoot:            history.Merkle,
		FirstState:             history.FirstLeaf,
		FirstStatehistoryProof: firstLeafProof,
		LastState:              history.LastLeaf,
		LastStatehistoryProof:  lastLeafProof,
	}

	// Check the current mini-stake amount and transact using that as the value.
	cManager, err := c.manager(ctx, tx)
	if err != nil {
		return nil, err
	}
	miniStake, err := cManager.miniStakeAmount()
	if err != nil {
		return nil, err
	}
	opts := copyTxOpts(c.chain.txOpts)
	opts.Value = miniStake

	_, err = transact(ctx, c.chain.backend, c.chain.headerReader, func() (*types.Transaction, error) {
		return cManager.writer.AddLeaf(
			opts,
			leafData,
			flatLastLeafProof,
			flatLastLeafProof,
		)
	})
	if err != nil {
		return nil, err
	}

	vertexId, err := cManager.caller.CalculateChallengeVertexId(
		c.chain.callOpts,
		c.id,
		history.Merkle,
		big.NewInt(int64(history.Height)),
	)
	if err != nil {
		return nil, err
	}
	_, err = cManager.caller.GetVertex(
		c.chain.callOpts,
		vertexId,
	)
	if err != nil {
		return nil, err
	}
	return &ChallengeVertex{
		id:    vertexId,
		chain: c.chain,
	}, nil
}

func (c *Challenge) inner(ctx context.Context, tx protocol.ActiveTx) (challengeV2gen.Challenge, error) {
	manager, err := c.manager(ctx, tx)
	if err != nil {
		return challengeV2gen.Challenge{}, err
	}

	challengeInner, err := manager.caller.GetChallenge(c.chain.callOpts, c.id)
	if err != nil {
		return challengeV2gen.Challenge{}, err
	}
	return challengeInner, nil
}

func (c *Challenge) manager(ctx context.Context, tx protocol.ActiveTx) (*ChallengeManager, error) {
	manager, err := c.chain.CurrentChallengeManager(ctx, tx)
	if err != nil {
		return nil, err
	}
	challengeManager, ok := manager.(*ChallengeManager)
	if !ok {
		return nil, errors.New("challengemanager is not expected concrete type")
	}
	return challengeManager, nil
}
