package auth

import (
	"context"
	"fmt"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/user_v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	resSrt := fmt.Sprintf("Received create:\n\tName: %v\n\tEmail: %v\n\tPassword:%v\n\tPassword confirm: %v\n\tRole: %v",
		req.User.GetName().GetValue(),
		req.User.GetEmail().GetValue(),
		req.Password.GetValue(),
		req.PasswordConfirm.GetValue(),
		req.User.Role.String())

	fmt.Println(resSrt)

	if dline, ok := ctx.Deadline(); ok {
		i.log.Printf("Deadline %v\n", dline)
	}

	if err := validateUserCreateRequest(req); err != nil {
		return nil, err
	}

	id, err := i.authservice.Create(ctx, converter.ToModelFromCreateRequest(req))

	if err != nil {
		i.log.Printf("error in service layer: %v\n", err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.CreateResponse{
		Id: &wrapperspb.Int64Value{
			Value: int64(id),
		},
	}, nil
}

func validateUserCreateRequest(req *desc.CreateRequest) error {

	if req.User.Name == nil {
		return status.Error(codes.InvalidArgument, "Name is requared")
	}

	if req.User.Email == nil {
		return status.Error(codes.InvalidArgument, "Email is requared")
	}

	if req.Password == nil {
		return status.Error(codes.InvalidArgument, "Password is requared")
	}

	if req.PasswordConfirm == nil {
		return status.Error(codes.InvalidArgument, "Password Confirm is requared")
	}

	if req.PasswordConfirm.GetValue() != req.Password.GetValue() {
		return status.Error(codes.InvalidArgument, "Password not equal")
	}

	if req.User.Role == desc.Role_UNKNOWN {
		return status.Error(codes.InvalidArgument, "Role is reqared")
	}
	return nil
}
