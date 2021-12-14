package relay

import (
	"context"

	eth2types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/proto/eth/service"
	ethpbv1 "github.com/prysmaticlabs/prysm/proto/eth/v1"
	ethpbv2 "github.com/prysmaticlabs/prysm/proto/eth/v2"
	ethpbalpha "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

type Api int

const (
	Prysm = iota
	Eth
)

// TODO: Visibility of field types
type Relayer struct {
	api                  Api
	PrysmValidatorClient ethpbalpha.BeaconNodeValidatorClient
	EthValidatorClient   service.BeaconValidatorClient
	EthBeaconChainClient service.BeaconChainClient
}

func New(api Api) *Relayer {
	return &Relayer{api: api}
}

func (r *Relayer) GetDuties(ctx context.Context, epoch eth2types.Epoch, publicKeys [][]byte) (*Duties, error) {
	if len(publicKeys) == 0 {
		// TODO: What to do? Return empty duties without error?
	}

	if r.api == Prysm {
		apiDuties, err := r.PrysmValidatorClient.GetDuties(ctx, &ethpbalpha.DutiesRequest{
			Epoch:      epoch,
			PublicKeys: publicKeys,
		})
		if err != nil {
			return nil, err
		}
		return convertPrysmDutiesResponse(apiDuties), nil
	} else {
		// TODO: What about things like:
		// https://github.com/prysmaticlabs/prysm/blob/52d8a1646fb49ab4680e55040103926dddee7a67/beacon-chain/rpc/prysm/v1alpha1/validator/assignments.go#L204

		vals, err := r.EthBeaconChainClient.ListValidators(ctx, &ethpbv1.StateValidatorsRequest{StateId: []byte("head"), Id: publicKeys})
		if err != nil {
			return nil, err
		}
		if vals == nil || len(vals.Data) == 0 {
			// TODO: What to do? Return empty duties without error?
		}
		indices := make([]eth2types.ValidatorIndex, len(vals.Data))
		for i, v := range vals.Data {
			indices[i] = v.Index
		}

		currEpochAttesterDuties, err := r.EthValidatorClient.GetAttesterDuties(ctx, &ethpbv1.AttesterDutiesRequest{
			Epoch: epoch,
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		currEpochProposerDuties, err := r.EthValidatorClient.GetProposerDuties(ctx, &ethpbv1.ProposerDutiesRequest{Epoch: epoch})
		if err != nil {
			return nil, err
		}
		currEpochSyncCommitteeDuties, err := r.EthValidatorClient.GetSyncCommitteeDuties(ctx, &ethpbv2.SyncCommitteeDutiesRequest{
			Epoch: epoch,
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		nextEpochAttesterDuties, err := r.EthValidatorClient.GetAttesterDuties(ctx, &ethpbv1.AttesterDutiesRequest{
			Epoch: epoch.Add(1),
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		nextEpochSyncCommitteeDuties, err := r.EthValidatorClient.GetSyncCommitteeDuties(ctx, &ethpbv2.SyncCommitteeDutiesRequest{
			Epoch: epoch.Add(1),
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		currEpochCommmittees, err := r.EthBeaconChainClient.ListCommittees(ctx, &ethpbv1.StateCommitteesRequest{
			StateId: []byte("head"),
			Epoch:   &epoch,
		})
		if err != nil {
			return nil, err
		}
		nextEpoch := epoch.Add(1)
		nextEpochCommmittees, err := r.EthBeaconChainClient.ListCommittees(ctx, &ethpbv1.StateCommitteesRequest{
			StateId: []byte("head"),
			Epoch:   &nextEpoch,
		})
		if err != nil {
			return nil, err
		}

		return r.convertEthDutiesResponse(
			vals.Data,
			currEpochAttesterDuties,
			nextEpochAttesterDuties,
			currEpochProposerDuties,
			currEpochSyncCommitteeDuties,
			nextEpochSyncCommitteeDuties,
			currEpochCommmittees.Data,
			nextEpochCommmittees.Data,
		)
	}
}
