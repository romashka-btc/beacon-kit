// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package blockchain

import (
	"context"
	"time"

	payloadtime "github.com/berachain/beacon-kit/mod/beacon/payload-time"
	engineerrors "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/errors"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
)

// VerifyIncomingBlock verifies the state root of an incoming block
// and logs the process.
func (s *Service[
	_, ConsensusBlockT, BeaconBlockT, _, _, _, _, _, ExecutionPayloadHeaderT, _, _,
]) VerifyIncomingBlock(
	ctx context.Context,
	blk ConsensusBlockT,
) error {
	var (
		beaconBlk     = blk.GetBeaconBlock()
		consensusTime = blk.GetConsensusTime()
	)

	// Grab a copy of the state to verify the incoming block.
	preState := s.storageBackend.StateFromContext(ctx)

	// Force a sync of the startup head if we haven't done so already.
	//
	// TODO: This is a super hacky. It should be handled better elsewhere,
	// ideally via some broader sync service.
	s.forceStartupSyncOnce.Do(func() { s.forceStartupHead(ctx, preState) })

	// If the block is nil or a nil pointer, exit early.
	if beaconBlk.IsNil() {
		s.logger.Warn(
			"Aborting block verification - beacon block not found in proposal",
		)
		return errors.WrapNonFatal(ErrNilBlk)
	}

	s.logger.Info(
		"Received incoming beacon block",
		"state_root", beaconBlk.GetStateRoot(), "slot", beaconBlk.GetSlot(),
	)

	// We purposefully make a copy of the BeaconState in order
	// to avoid modifying the underlying state, for the event in which
	// we have to rebuild a payload for this slot again, if we do not agree
	// with the incoming block.
	postState := preState.Copy()

	// Verify the state root of the incoming block.
	if err := s.verifyStateRoot(ctx, postState, blk); err != nil {
		s.logger.Error(
			"Rejecting incoming beacon block ❌ ",
			"state_root",
			beaconBlk.GetStateRoot(),
			"reason",
			err,
		)

		if s.shouldBuildOptimisticPayloads() {
			var lph ExecutionPayloadHeaderT
			lph, err = preState.GetLatestExecutionPayloadHeader()
			if err != nil {
				return err
			}

			go s.handleRebuildPayloadForRejectedBlock(
				ctx,
				preState,
				payloadtime.Next(
					consensusTime,
					lph.GetTimestamp(),
					true, // buildOptimistically
				),
			)
		}

		return err
	}

	s.logger.Info(
		"State root verification succeeded - accepting incoming beacon block",
		"state_root",
		beaconBlk.GetStateRoot(),
	)

	if s.shouldBuildOptimisticPayloads() {
		lph, err := postState.GetLatestExecutionPayloadHeader()
		if err != nil {
			return err
		}

		go s.handleOptimisticPayloadBuild(
			ctx,
			postState,
			beaconBlk,
			payloadtime.Next(
				consensusTime,
				lph.GetTimestamp(),
				true, // buildOptimistically
			),
		)
	}

	return nil
}

// verifyStateRoot verifies the state root of an incoming block.
func (s *Service[
	_, ConsensusBlockT, _, _, _, BeaconStateT, _, _, _, _, _,
]) verifyStateRoot(
	ctx context.Context,
	st BeaconStateT,
	blk ConsensusBlockT,
) error {
	startTime := time.Now()
	defer s.metrics.measureStateRootVerificationTime(startTime)
	_, err := s.stateProcessor.Transition(
		// We run with a non-optimistic engine here to ensure
		// that the proposer does not try to push through a bad block.
		&transition.Context{
			Context:                 ctx,
			OptimisticEngine:        false,
			SkipPayloadVerification: false,
			SkipValidateResult:      false,
			SkipValidateRandao:      false,
			ProposerAddress:         blk.GetProposerAddress(),
			ConsensusTime:           blk.GetConsensusTime(),
		},
		st, blk.GetBeaconBlock(),
	)
	if errors.Is(err, engineerrors.ErrAcceptedPayloadStatus) {
		// It is safe for the validator to ignore this error since
		// the state transition will enforce that the block is part
		// of the canonical chain.
		//
		// TODO: this is only true because we are assuming SSF.
		return nil
	}

	return err
}

// shouldBuildOptimisticPayloads returns true if optimistic
// payload builds are enabled.
func (s *Service[
	_, _, _, _, _, _, _, _, _, _, _,
]) shouldBuildOptimisticPayloads() bool {
	return s.optimisticPayloadBuilds && s.localBuilder.Enabled()
}
