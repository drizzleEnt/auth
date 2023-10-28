package auth

import (
	"context"
	"fmt"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	i.log.Printf("Receive Update")

	if dline, ok := ctx.Deadline(); ok {
		i.log.Printf("Deadline %v\n", dline)
	}

	if err := validateUserUpdateRequest(req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err := i.authservice.Update(ctx, converter.ToModelUserFromUpdateRequest(req)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func validateUserUpdateRequest(req *desc.UpdateRequest) error {
	if req.Id == nil {
		return fmt.Errorf("id is requared")
	}

	if req.User == nil {
		return fmt.Errorf("user is requared")
	}

	return nil
}
