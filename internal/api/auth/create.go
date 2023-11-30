package auth

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/user_v2"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	id, err := i.authservice.Create(ctx, converter.ToUserFromDescCreate(req.GetInfo()))

	if err != nil {
		log.Printf("error in service layer: %v\n", err.Error())
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
