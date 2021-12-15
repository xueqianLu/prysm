package relay

import (
	types "github.com/prysmaticlabs/eth2-types"
	ethpbv1 "github.com/prysmaticlabs/prysm/proto/eth/v1"
	ethpbv2 "github.com/prysmaticlabs/prysm/proto/eth/v2"
	ethpbalpha "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

func convertPrysmDutiesResponse(resp *ethpbalpha.DutiesResponse) *Duties {
	currEpoch := make([]*Duty, len(resp.CurrentEpochDuties))
	for i, d := range resp.CurrentEpochDuties {
		currEpoch[i] = &Duty{
			Committee:       d.Committee,
			CommitteeIndex:  d.CommitteeIndex,
			AttesterSlot:    d.AttesterSlot,
			ProposerSlots:   d.ProposerSlots,
			PublicKey:       d.PublicKey,
			Status:          convertPrysmValidatorStatus(d.Status),
			ValidatorIndex:  d.ValidatorIndex,
			IsSyncCommittee: d.IsSyncCommittee,
		}
	}
	nextEpoch := make([]*Duty, len(resp.NextEpochDuties))
	for i, d := range resp.NextEpochDuties {
		nextEpoch[i] = &Duty{
			Committee:       d.Committee,
			CommitteeIndex:  d.CommitteeIndex,
			AttesterSlot:    d.AttesterSlot,
			ProposerSlots:   d.ProposerSlots,
			PublicKey:       d.PublicKey,
			Status:          convertPrysmValidatorStatus(d.Status),
			ValidatorIndex:  d.ValidatorIndex,
			IsSyncCommittee: d.IsSyncCommittee,
		}
	}
	return &Duties{
		CurrentEpochDuties: currEpoch,
		NextEpochDuties:    nextEpoch,
	}
}

func (r *Relayer) convertEthDutiesResponse(
	vals []*ethpbv1.ValidatorContainer,
	currEpochAttesterResp *ethpbv1.AttesterDutiesResponse,
	nextEpochAttesterResp *ethpbv1.AttesterDutiesResponse,
	currEpochProposerResp *ethpbv1.ProposerDutiesResponse,
	currEpochSyncCommitteeResp *ethpbv2.SyncCommitteeDutiesResponse,
	nextEpochSyncCommitteeResp *ethpbv2.SyncCommitteeDutiesResponse,
	currEpochCommmittees []*ethpbv1.Committee,
	nextEpochCommmittees []*ethpbv1.Committee,
) (*Duties, error) {
	currEpochAttesterMap := make(map[types.ValidatorIndex]*ethpbv1.AttesterDuty, len(currEpochAttesterResp.Data))
	nextEpochAttesterMap := make(map[types.ValidatorIndex]*ethpbv1.AttesterDuty, len(nextEpochAttesterResp.Data))
	currEpochProposerMap := make(map[types.ValidatorIndex][]*ethpbv1.ProposerDuty, len(currEpochProposerResp.Data))
	currEpochSyncCommitteeMap := make(map[types.ValidatorIndex]*ethpbv2.SyncCommitteeDuty, len(currEpochSyncCommitteeResp.Data))
	nextEpochSyncCommitteeMap := make(map[types.ValidatorIndex]*ethpbv2.SyncCommitteeDuty, len(nextEpochSyncCommitteeResp.Data))

	for _, d := range currEpochAttesterResp.Data {
		currEpochAttesterMap[d.ValidatorIndex] = d
	}
	for _, d := range nextEpochAttesterResp.Data {
		nextEpochAttesterMap[d.ValidatorIndex] = d
	}
	for _, d := range currEpochProposerResp.Data {
		if currEpochProposerMap[d.ValidatorIndex] == nil {
			currEpochProposerMap[d.ValidatorIndex] = make([]*ethpbv1.ProposerDuty, 0)
		}
		currEpochProposerMap[d.ValidatorIndex] = append(currEpochProposerMap[d.ValidatorIndex], d)
	}
	for _, d := range currEpochSyncCommitteeResp.Data {
		currEpochSyncCommitteeMap[d.ValidatorIndex] = d
	}
	for _, d := range nextEpochSyncCommitteeResp.Data {
		nextEpochSyncCommitteeMap[d.ValidatorIndex] = d
	}

	indices := make([]types.ValidatorIndex, len(vals))
	for i, v := range vals {
		indices[i] = v.Index
	}

	currEpochDuties := make([]*Duty, 0)
	nextEpochDuties := make([]*Duty, 0)
	for i, idx := range indices {
		hasCurrentDuty := currEpochAttesterMap[idx] != nil || currEpochProposerMap[idx] != nil || currEpochSyncCommitteeMap[idx] != nil
		hasNextDuty := nextEpochAttesterMap[idx] != nil || nextEpochSyncCommitteeMap[idx] != nil

		if hasCurrentDuty {
			var attSlot types.Slot
			committee := make([]types.ValidatorIndex, 0)
			var committeeIndex types.CommitteeIndex
			attDuty, ok := currEpochAttesterMap[idx]
			if ok {
				attSlot = attDuty.Slot

			currEpoch:
				for _, c := range currEpochCommmittees {
					for _, idx := range c.Validators {
						if idx == types.ValidatorIndex(i) {
							committee = c.Validators
							committeeIndex = c.Index
							break currEpoch
						}
					}
				}
			}

			proposerSlots := make([]types.Slot, 0)
			proposerDuties, ok := currEpochProposerMap[idx]
			if ok {
				for _, d := range proposerDuties {
					proposerSlots = append(proposerSlots, d.Slot)
				}
			}

			isSyncCommittee := false
			_, ok = currEpochSyncCommitteeMap[idx]
			if ok {
				isSyncCommittee = true
			}

			duty := &Duty{
				Committee:       committee,
				CommitteeIndex:  committeeIndex,
				AttesterSlot:    attSlot,
				ProposerSlots:   proposerSlots,
				PublicKey:       vals[i].Validator.Pubkey,
				Status:          convertEthValidatorStatus(vals[i].Status),
				ValidatorIndex:  idx,
				IsSyncCommittee: isSyncCommittee,
			}
			currEpochDuties = append(currEpochDuties, duty)
		}
		if hasNextDuty {
			var attSlot types.Slot
			committee := make([]types.ValidatorIndex, 0)
			var committeeIndex types.CommitteeIndex
			attDuty, ok := nextEpochAttesterMap[idx]
			if ok {
				attSlot = attDuty.Slot

			nextEpoch:
				for _, c := range nextEpochCommmittees {
					for _, idx := range c.Validators {
						if idx == types.ValidatorIndex(i) {
							committee = c.Validators
							committeeIndex = c.Index
							break nextEpoch
						}
					}
				}
			}

			isSyncCommittee := false
			_, ok = nextEpochSyncCommitteeMap[idx]
			if ok {
				isSyncCommittee = true
			}

			duty := &Duty{
				Committee:       committee,
				CommitteeIndex:  committeeIndex,
				AttesterSlot:    attSlot,
				ProposerSlots:   make([]types.Slot, 0),
				PublicKey:       vals[i].Validator.Pubkey,
				Status:          convertEthValidatorStatus(vals[i].Status),
				ValidatorIndex:  idx,
				IsSyncCommittee: isSyncCommittee,
			}
			nextEpochDuties = append(nextEpochDuties, duty)
		}
	}

	return &Duties{
		CurrentEpochDuties: currEpochDuties,
		NextEpochDuties:    nextEpochDuties,
	}, nil
}

func convertPrysmValidatorStatus(status ethpbalpha.ValidatorStatus) ValidatorStatus {
	switch status {
	case ethpbalpha.ValidatorStatus_UNKNOWN_STATUS:
		return ValidatorStatus_Unknown
	case ethpbalpha.ValidatorStatus_DEPOSITED:
		return ValidatorStatus_Deposited
	case ethpbalpha.ValidatorStatus_PENDING:
		return ValidatorStatus_Pending
	case ethpbalpha.ValidatorStatus_ACTIVE:
		return ValidatorStatus_Active
	case ethpbalpha.ValidatorStatus_EXITING:
		return ValidatorStatus_Exiting
	case ethpbalpha.ValidatorStatus_SLASHING:
		return ValidatorStatus_Slashing
	case ethpbalpha.ValidatorStatus_EXITED:
		return ValidatorStatus_Exited
	case ethpbalpha.ValidatorStatus_INVALID:
		return ValidatorStatus_Invalid
	case ethpbalpha.ValidatorStatus_PARTIALLY_DEPOSITED:
		return ValidatorStatus_PartiallyDeposited
	default:
		return ValidatorStatus_Unknown
	}
}

func convertEthValidatorStatus(status ethpbv1.ValidatorStatus) ValidatorStatus {
	switch status {
	case ethpbv1.ValidatorStatus_PENDING_INITIALIZED, ethpbv1.ValidatorStatus_PENDING_QUEUED:
		return ValidatorStatus_Pending
	case ethpbv1.ValidatorStatus_ACTIVE_ONGOING, ethpbv1.ValidatorStatus_ACTIVE_SLASHED, ethpbv1.ValidatorStatus_ACTIVE_EXITING:
		return ValidatorStatus_Active
	case ethpbv1.ValidatorStatus_EXITED_UNSLASHED,
		ethpbv1.ValidatorStatus_EXITED_SLASHED,
		ethpbv1.ValidatorStatus_WITHDRAWAL_POSSIBLE,
		ethpbv1.ValidatorStatus_WITHDRAWAL_DONE:
		return ValidatorStatus_Exited
	default:
		return ValidatorStatus_Unknown
	}
}
