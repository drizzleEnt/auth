package auth

import (
	"context"
	"log"

	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	log.Printf("Received Delete in ID: %v\n", req.GetId())

	if err := i.authservice.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
