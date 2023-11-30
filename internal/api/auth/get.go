package auth

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/user_v2"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	log.Printf("Receive Get ID: %v", req.Id)

	authResp, err := i.authservice.Get(ctx, req.GetId())

	if err != nil {
		log.Printf("Error try get user whit id=%v %v\n", req.Id, err)
		return nil, err
	}

	log.Printf("Id = %d", authResp.UserCreate.UserUpdate.ID)

	return &desc.GetResponse{
		User: converter.ToUserFromService(authResp),
	}, nil
}
