package rpc

import (
	"context"

	ethpbservice "github.com/prysmaticlabs/prysm/proto/eth/service"
)

func (s *Server) DeleteKeystores(
	ctx context.Context, req *ethpbservice.DeleteKeystoresRequest,
) (*ethpbservice.DeleteKeystoresResponse, error) {
	res := &ethpbservice.DeleteKeystoresResponse{
		Statuses: []*ethpbservice.DeletedKeystoreStatus{
			{
				Status:  ethpbservice.DeletedKeystoreStatus_DELETED,
				Message: "Did not work",
			},
		},
		SlashingProtection: "",
	}
	return res, nil
}
