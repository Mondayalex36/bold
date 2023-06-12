package assertions

import (
	"context"
	"time"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	challengemanager "github.com/OffchainLabs/challenge-protocol-v2/challenge-manager"
	retry "github.com/OffchainLabs/challenge-protocol-v2/runtime"
	"github.com/OffchainLabs/challenge-protocol-v2/solgen/go/rollupgen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("prefix", "assertion-scanner")

type Scanner struct {
	chain            protocol.AssertionChain
	backend          bind.ContractBackend
	challengeManager challengemanager.ChallengeCreator
	pollInterval     time.Duration
	rollupAddr       common.Address
	validatorName    string
}

func (s *Scanner) Scan(ctx context.Context) {
	scanRange, err := retry.UntilSucceeds(ctx, func() (filterRange, error) {
		return s.getStartEndBlockNum(ctx)
	})
	if err != nil {
		log.Error(err)
		return
	}
	fromBlock := scanRange.startBlockNum
	toBlock := scanRange.endBlockNum

	// Do the initial scan...
	filterer, err := retry.UntilSucceeds(ctx, func() (*rollupgen.RollupUserLogicFilterer, error) {
		return rollupgen.NewRollupUserLogicFilterer(s.rollupAddr, s.backend)
	})
	if err != nil {
		log.Error(err)
		return
	}
	filterOpts := &bind.FilterOpts{
		Start:   fromBlock,
		End:     &toBlock,
		Context: ctx,
	}
	_, err = retry.UntilSucceeds(ctx, func() (bool, error) {
		return true, s.checkForAssertionAdded(ctx, filterer, filterOpts)
	})
	if err != nil {
		log.Error(err)
		return
	}

	fromBlock = toBlock
	ticker := time.NewTicker(s.pollInterval)
	defer ticker.Stop()
	for {
		for {
			select {
			case <-ticker.C:

				// Scan up up to the block delta once more.
				latestBlock, err := s.backend.HeaderByNumber(ctx, nil)
				if err != nil {
					log.Error(err)
					continue
				}
				if !latestBlock.Number.IsUint64() {
					log.Error("latest block header number is not a uint64")
					continue
				}
				toBlock := latestBlock.Number.Uint64()
				if fromBlock == toBlock {
					continue
				}

				// Do what we need to do...

				fromBlock = toBlock
			case <-ctx.Done():
				return
			}
		}
	}
}

type filterRange struct {
	startBlockNum uint64
	endBlockNum   uint64
}

// Gets the start and end block numbers for our filter queries, starting from the
// latest confirmed assertion's block number up to the latest block number.
func (s *Scanner) getStartEndBlockNum(ctx context.Context) (filterRange, error) {
	latestConfirmed, err := s.chain.LatestConfirmed(ctx)
	if err != nil {
		return filterRange{}, err
	}
	firstBlock, err := latestConfirmed.CreatedAtBlock()
	if err != nil {
		return filterRange{}, err
	}
	startBlock := firstBlock
	header, err := s.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return filterRange{}, err
	}
	if !header.Number.IsUint64() {
		return filterRange{}, errors.New("header number is not a uint64")
	}
	return filterRange{
		startBlockNum: startBlock,
		endBlockNum:   header.Number.Uint64(),
	}, nil
}

func (s *Scanner) checkForAssertionAdded(
	ctx context.Context,
	filterer *rollupgen.RollupUserLogicFilterer,
	filterOpts *bind.FilterOpts,
) error {
	it, err := filterer.FilterAssertionCreated(filterOpts, nil, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err = it.Close(); err != nil {
			log.WithError(err).Error("Could not close filter iterator")
		}
	}()
	for it.Next() {
		if it.Error() != nil {
			return errors.Wrapf(
				err,
				"got iterator error when scanning assertion creations from block %d to %d",
				filterOpts.Start,
				*filterOpts.End,
			)
		}
		_, processErr := retry.UntilSucceeds(ctx, func() (bool, error) {
			return true, s.processAssertionCreation(ctx, it.Event)
		})
		if processErr != nil {
			return processErr
		}
	}
	return nil
}

func (s *Scanner) processAssertionCreation(
	ctx context.Context,
	event *rollupgen.RollupUserLogicAssertionCreated,
) error {
	log.WithFields(logrus.Fields{
		"validatorName": s.validatorName,
	}).Info("Processed assertion creation event")
	assertion, err := s.chain.GetAssertion(ctx, event.AssertionHash)
	if err != nil {
		return err
	}
	isFirstChild, err := assertion.IsFirstChild()
	if err != nil {
		return err
	}
	if isFirstChild {
		return nil
	}
	return s.challengeManager.ChallengeAssertion(ctx, assertion.Id())
}
