package relay

import (
	"context"

	"github.com/pkg/errors"
	eth2types "github.com/prysmaticlabs/eth2-types"
	"github.com/prysmaticlabs/prysm/beacon-chain/blockchain"
	"github.com/prysmaticlabs/prysm/encoding/bytesutil"
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
	prysmValidatorClient ethpbalpha.BeaconNodeValidatorClient
	ethValidatorClient   service.BeaconValidatorClient
	headFetcher          blockchain.HeadFetcher
}

func New(api Api) *Relayer {
	return &Relayer{api: api}
}

func (r *Relayer) GetDuties(ctx context.Context, epoch eth2types.Epoch, publicKeys [][]byte) (*Duties, error) {
	var duties Duties

	if r.api == Prysm {
		apiDuties, err := r.prysmValidatorClient.GetDuties(ctx, &ethpbalpha.DutiesRequest{
			Epoch:      epoch,
			PublicKeys: publicKeys,
		})
		if err != nil {
			return nil, err
		}
		return ConvertPrysmDutiesResponse(apiDuties), nil
	} else {
		s, err := r.headFetcher.HeadState(ctx)
		if err != nil {
			return nil, err
		}
		indices := make([]eth2types.ValidatorIndex, len(publicKeys))
		for i, pk := range publicKeys {
			idx, ok := s.ValidatorIndexByPubkey(bytesutil.ToBytes48(pk))
			if !ok {
				return nil, errors.New("could not get validator index")
			}
			indices[i] = idx
		}
		currEpochAttesterDuties, err := r.ethValidatorClient.GetAttesterDuties(ctx, &ethpbv1.AttesterDutiesRequest{
			Epoch: epoch,
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		currEpochProposerDuties, err := r.ethValidatorClient.GetProposerDuties(ctx, &ethpbv1.ProposerDutiesRequest{Epoch: epoch})
		if err != nil {
			return nil, err
		}
		currEpochSyncCommitteeDuties, err := r.ethValidatorClient.GetSyncCommitteeDuties(ctx, &ethpbv2.SyncCommitteeDutiesRequest{
			Epoch: epoch,
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		nextEpochAttesterDuties, err := r.ethValidatorClient.GetAttesterDuties(ctx, &ethpbv1.AttesterDutiesRequest{
			Epoch: epoch.Add(1),
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
		nextEpochProposerDuties, err := r.ethValidatorClient.GetProposerDuties(ctx, &ethpbv1.ProposerDutiesRequest{Epoch: epoch.Add(1)})
		if err != nil {
			return nil, err
		}
		nextEpochSyncCommitteeDuties, err := r.ethValidatorClient.GetSyncCommitteeDuties(ctx, &ethpbv2.SyncCommitteeDutiesRequest{
			Epoch: epoch.Add(1),
			Index: indices,
		})
		if err != nil {
			return nil, err
		}
	}

	return &duties, nil
}
