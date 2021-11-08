package rpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	ethpbservice "github.com/prysmaticlabs/prysm/proto/eth/service"
	"github.com/prysmaticlabs/prysm/validator/keymanager"
	"github.com/prysmaticlabs/prysm/validator/keymanager/derived"
	"github.com/prysmaticlabs/prysm/validator/keymanager/imported"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListKeystores implements the standard validator key management API.
func (s Server) ListKeystores(
	ctx context.Context, _ *empty.Empty,
) (*ethpbservice.ListKeystoresResponse, error) {
	if !s.walletInitialized {
		return nil, status.Error(codes.Internal, "Wallet not ready")
	}
	pubKeys, err := s.keymanager.FetchValidatingPublicKeys(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not list keystores: %v", err)
	}
	keystoreResponse := make([]*ethpbservice.ListKeystoresResponse_Keystore, len(pubKeys))
	for i := 0; i < len(pubKeys); i++ {
		keystoreResponse[i] = &ethpbservice.ListKeystoresResponse_Keystore{
			ValidatingPubkey: pubKeys[i][:],
		}
		if s.wallet.KeymanagerKind() == keymanager.Derived {
			keystoreResponse[i].DerivationPath = fmt.Sprintf(derived.ValidatingKeyDerivationPathTemplate, i)
		}
	}
	return &ethpbservice.ListKeystoresResponse{
		Keystores: keystoreResponse,
	}, nil
}

func (s *Server) DeleteKeystores(
	ctx context.Context, req *ethpbservice.DeleteKeystoresRequest,
) (*ethpbservice.DeleteKeystoresResponse, error) {
	km, ok := s.keymanager.(*imported.Keymanager)
	if !ok {
		return nil, nil
	}
	if err := km.DeleteAccounts(ctx, req.PublicKeys); err != nil {
		return nil, err
	}
	protection, err := s.ExportSlashingProtection(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	res := &ethpbservice.DeleteKeystoresResponse{
		Statuses: []*ethpbservice.DeletedKeystoreStatus{
			{
				Status:  ethpbservice.DeletedKeystoreStatus_DELETED,
				Message: "Did not work",
			},
		},
		SlashingProtection: protection.File,
	}
	return res, nil
}
