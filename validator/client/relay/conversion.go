package relay

import (
	ethpbv1 "github.com/prysmaticlabs/prysm/proto/eth/v1"
	ethpbv2 "github.com/prysmaticlabs/prysm/proto/eth/v2"
	ethpbalpha "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

func ConvertPrysmDutiesResponse(resp *ethpbalpha.DutiesResponse) *Duties {
	currEpoch := make([]*Duty, len(resp.CurrentEpochDuties))
	for i, d := range resp.CurrentEpochDuties {
		currEpoch[i] = &Duty{
			Committee:       d.Committee,
			CommitteeIndex:  d.CommitteeIndex,
			AttesterSlot:    d.AttesterSlot,
			ProposerSlots:   d.ProposerSlots,
			PublicKey:       d.PublicKey,
			Status:          ConvertPrysmValidatorStatus(d.Status),
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
			Status:          ConvertPrysmValidatorStatus(d.Status),
			ValidatorIndex:  d.ValidatorIndex,
			IsSyncCommittee: d.IsSyncCommittee,
		}
	}
	return &Duties{
		CurrentEpochDuties: currEpoch,
		NextEpochDuties:    nextEpoch,
	}
}

func ConvertEthDutiesResponse(
	currEpochAttesterResp ethpbv1.AttesterDutiesResponse,
	nextEpochAttesterResp ethpbv1.AttesterDutiesResponse,
	currEpochProposerResp ethpbv1.ProposerDutiesResponse,
	nextEpochProposerResp ethpbv1.ProposerDutiesResponse,
	currEpochSyncCommitteeResp ethpbv2.SyncCommitteeDutiesResponse,
	nextEpochSyncCommitteeResp ethpbv2.SyncCommitteeDutiesResponse,
) {
	duties := &Duties{
		CurrentEpochDuties: make([]*Duty, 0),
		NextEpochDuties:    make([]*Duty, 0),
	}
}

func ConvertPrysmValidatorStatus(status ethpbalpha.ValidatorStatus) ValidatorStatus {
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
