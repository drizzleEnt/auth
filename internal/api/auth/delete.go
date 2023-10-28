package auth

import (
	"context"

	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	if dline, ok := ctx.Deadline(); ok {
		i.log.Printf("Deadline %v\n", dline)
	}

	i.log.Printf("Received Delete in ID: %v\n", req.GetId())

	if err := i.authservice.Delete(ctx, req.Id.Value); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user %v", err)
	}

	return &emptypb.Empty{}, nil
}
