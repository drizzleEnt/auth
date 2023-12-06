package auth

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	log.Printf("Receive Update")

	if err := i.authservice.Update(ctx, converter.ToUserFromDescUpdate(req.GetInfo())); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
