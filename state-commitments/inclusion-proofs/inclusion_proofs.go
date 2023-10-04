// Package inclusionproofs defines a series of utilities for generating and verifying
// traditional Merkle proofs of data.
//
// Copyright 2023, Offchain Labs, Inc.
// For license information, see https://github.com/offchainlabs/bold/blob/main/LICENSE
package inclusionproofs

import (
	prefixproofs "github.com/OffchainLabs/bold/state-commitments/prefix-proofs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	ErrProofTooLong  = errors.New("merkle proof too long")
	ErrInvalidLeaves = errors.New("invalid number of leaves for merkle tree")
)

// FullTree generates a Merkle tree from a list of leaves.
func FullTree(leaves []common.Hash) ([][]common.Hash, error) {
	msb, err := prefixproofs.MostSignificantBit(uint64(len(leaves)))
	if err != nil {
		return nil, err
	}
	lsb, err := prefixproofs.LeastSignificantBit(uint64(len(leaves)))
	if err != nil {
		return nil, err
	}
	maxLevel := msb + 1
	if msb == lsb {
		maxLevel = msb
	}

	layers := make([][]common.Hash, maxLevel+1)
	layers[0] = leaves
	l := uint64(1)

	prevLayer := leaves
	for len(prevLayer) > 1 {
		nextLayer := make([]common.Hash, (len(prevLayer)+1)/2)
		for i := 0; i < len(nextLayer); i++ {
			if 2*i+1 < len(prevLayer) {
				nextLayer[i] = crypto.Keccak256Hash(prevLayer[2*i].Bytes(), prevLayer[2*i+1].Bytes())
			} else {
				nextLayer[i] = crypto.Keccak256Hash(prevLayer[2*i].Bytes(), (common.Hash{}).Bytes())
			}
		}
		layers[l] = nextLayer
		prevLayer = nextLayer
		l++
	}
	return layers, nil
}

// GenerateInclusionProof from a list of Merkle leaves at a specified index.
func GenerateInclusionProof(leaves []common.Hash, idx uint64) ([]common.Hash, error) {
	if len(leaves) == 0 {
		return nil, ErrInvalidLeaves
	}
	if idx >= uint64(len(leaves)) {
		return nil, ErrInvalidLeaves
	}
	if len(leaves) == 1 {
		return make([]common.Hash, 0), nil
	}
	rehashed := make([]common.Hash, len(leaves))
	for i, r := range leaves {
		rehashed[i] = crypto.Keccak256Hash(r.Bytes())
	}

	fullT, err := FullTree(rehashed)
	if err != nil {
		return nil, err
	}
	maxLevel, err := prefixproofs.MostSignificantBit(uint64(len(rehashed)) - 1)
	if err != nil {
		return nil, err
	}
	proof := make([]common.Hash, maxLevel+1)

	for level := uint64(0); level <= maxLevel; level++ {
		levelIndex := idx >> level
		counterpartIndex := levelIndex ^ 1
		layer := fullT[level]
		counterpart := common.Hash{}
		if counterpartIndex <= uint64(len(layer))-1 {
			counterpart = layer[counterpartIndex]
		}
		proof[level] = counterpart
	}

	return proof, nil
}

// CalculateRootFromProof calculates a Merkle root from a Merkle proof, index, and leaf.
func CalculateRootFromProof(proof []common.Hash, index uint64, leaf common.Hash) (common.Hash, error) {
	if len(proof) > 256 {
		return common.Hash{}, ErrProofTooLong
	}
	h := crypto.Keccak256Hash(leaf[:])
	for i := 0; i < len(proof); i++ {
		node := proof[i]
		if index&(1<<i) == 0 {
			h = crypto.Keccak256Hash(h[:], node[:])
		} else {
			h = crypto.Keccak256Hash(node[:], h[:])
		}
	}
	return h, nil
}

func GenerateInclusionProofForFirstElement(me prefixproofs.MerkleExpansion) []common.Hash {
	var proof []common.Hash

	// Iterate over the MerkleExpansion levels
	for i := 0; i < len(me)-1; i++ {
		// Since we are generating the proof for the first element,
		// it's always on the left side, so we add the right siblings to the proof.
		// However, we need to ensure that the sibling (right side) exists and is not empty.
		if (i+1) < len(me) && me[i+1] != (common.Hash{}) {
			proof = append(proof, me[i+1])
		}
	}

	return proof
}

func GenerateInclusionProofForLastElement(me prefixproofs.MerkleExpansion) ([]common.Hash, error) {
	var proof []common.Hash

	if len(me) == 0 {
		return nil, errors.New("empty MerkleExpansion")
	}

	// Find the position of the last element in the tree
	lastElementIndex := uint64(len(me)) - 1

	// Iterate over the MerkleExpansion levels in reverse to find the siblings of the last element
	for i := len(me) - 1; i >= 0; i-- {
		// If the element is on the right side (which the last element often is), add the left sibling to the proof
		// However, we need to ensure that the sibling (left side) exists and is not empty
		if i > 0 && me[i-1] != (common.Hash{}) {
			proof = append(proof, me[i-1])
		}
		// Update the position to the parent's index
		lastElementIndex = lastElementIndex / 2
	}

	return proof, nil
}
