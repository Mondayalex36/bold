// Package history defines the primitive HistoryCommitment type in the BOLD
// protocol.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package history

import (
	"errors"
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	state_hashes "github.com/OffchainLabs/bold/state-commitments/state-hashes"
	"sync"

	inclusionproofs "github.com/OffchainLabs/bold/state-commitments/inclusion-proofs"
	"github.com/ethereum/go-ethereum/common"
)

var (
	emptyCommit = History{}
)

// History defines a Merkle accumulator over a list of leaves, which
// are understood to be state roots in the goimpl. A history commitment contains
// a "height" value, which can refer to a height of an assertion in the assertions
// tree, or a "step" of WAVM states in a big step or small step subchallenge.
// A commitment contains a Merkle root over the list of leaves, and can optionally
// provide a proof that the last leaf in the accumulator Merkleizes into the
// specified root hash, which is required when verifying challenge creation invariants.
type History struct {
	Height         uint64
	Merkle         common.Hash
	FirstLeaf      common.Hash
	LastLeafProof  []common.Hash
	FirstLeafProof []common.Hash
	LastLeaf       common.Hash
}

func New(leaves *state_hashes.StateHashes) (History, error) {
	if leaves.Length() == 0 {
		return emptyCommit, errors.New("must commit to at least one leaf")
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	var firstLeafProof []common.Hash
	var err1 error
	go func() {
		defer waitGroup.Done()
		firstLeafProof, err1 = inclusionproofs.GenerateInclusionProof(leaves, 0)
	}()

	var lastLeafProof []common.Hash
	var err2 error
	go func() {
		defer waitGroup.Done()
		lastLeafProof, err2 = inclusionproofs.GenerateInclusionProof(leaves, leaves.Length()-1)
	}()

	var root common.Hash
	var err3 error
	go func() {
		defer waitGroup.Done()
		exp := prefixproofs.NewEmptyMerkleExpansion()
		for i := uint64(0); i < leaves.Length(); i++ {
			exp, err3 = prefixproofs.AppendLeaf(exp, leaves.At(i))
			if err3 != nil {
				return
			}
		}
		root, err3 = prefixproofs.Root(exp)
	}()
	waitGroup.Wait()

	if err1 != nil {
		return emptyCommit, err1
	}
	if err2 != nil {
		return emptyCommit, err2
	}
	if err3 != nil {
		return emptyCommit, err3
	}

	return History{
		Merkle:         root,
		Height:         leaves.Length() - 1,
		FirstLeaf:      leaves.At(0),
		LastLeaf:       leaves.At(leaves.Length() - 1),
		FirstLeafProof: firstLeafProof,
		LastLeafProof:  lastLeafProof,
	}, nil
}
