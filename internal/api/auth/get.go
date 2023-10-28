package auth

import (
	"context"

	desc "github.com/drizzleent/auth/pkg/user_v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	i.log.Printf("Receive Get ID: %v", req.Id)

	if dline, ok := ctx.Deadline(); ok {
		i.log.Printf("Deadline %v\n", dline)
	}

	resp, err := i.authservice.Get(ctx, req.Id.GetValue())

	if err != nil {
		i.log.Printf("Error try get user whit id=%v %v\n", req.Id, err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	descUserInfo := &desc.UserInfo{
		Id: &wrapperspb.Int64Value{Value: resp.ID},
		User: &desc.User{
			Name: &wrapperspb.StringValue{
				Value: resp.Name,
			},
			Email: &wrapperspb.StringValue{
				Value: resp.Email,
			},
			Role: desc.Role(resp.Role),
		},
		CreatedAt: &timestamppb.Timestamp{
			Seconds: resp.CreatedAt.Unix(),
			Nanos:   int32(resp.CreatedAt.Nanosecond()),
		},
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: resp.UpdatedAt.Time.Unix(),
			Nanos:   int32(resp.UpdatedAt.Time.Nanosecond()),
		},
	}

	return &desc.GetResponse{
		UserInfo: descUserInfo,
	}, nil
}
