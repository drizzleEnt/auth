package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/drizzleent/auth/internal/repository"
	desc "github.com/drizzleent/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRpcServerV1 struct {
	desc.UnimplementedUserV1Server
	log *log.Logger
	db  repository.Authorisation //TODO REMOVE FROM HERE
}

func NewUserRpcsServer(log *log.Logger, db repository.Authorisation) *UserRpcServerV1 {
	return &UserRpcServerV1{
		log: log,
		db:  db,
	}
}

func (s *UserRpcServerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	resSrt := fmt.Sprintf("Received create:\n\tName: %v\n\tEmail: %v\n\tPassword:%v\n\tPassword confirm: %v\n\tRole: %v",
		req.Name,
		req.Email,
		req.Password,
		req.PasswordConfirm,
		req.Role)

	fmt.Println(resSrt)

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline %v\n", dline)
	}

	// user := auth.User{
	// 	Name:     req.Name,
	// 	Email:    req.Email,
	// 	Password: req.Password,
	// 	Role:     req.Role.String(),
	// }

	// id, err := s.db.CreateUser(ctx, user)

	// if err != nil {
	// 	s.log.Printf("cant create user in database %v\n", err.Error())
	// 	return nil, status.Errorf(codes.Internal, err.Error())
	// }

	return &desc.CreateResponse{
		Id: 0, //int64(id),
	}, nil
}
func (s *UserRpcServerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	s.log.Printf("Receive Get ID: %v", req.Id)

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline %v\n", dline)
	}

	resp, err := s.db.Get(ctx, int(req.Id))

	if err != nil {
		s.log.Printf("Error try get user whit id=%v %v\n", req.Id, err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &desc.GetResponse{
		Id:    0,  //int64(resp.Id),
		Name:  "", //resp.Name,
		Email: "", //resp.Email,
		Role:  0,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: resp.CreatedAt.Unix(),
			Nanos:   int32(resp.CreatedAt.Nanosecond()),
		},
		UpdatedAt: &timestamppb.Timestamp{},
	}, nil
}
func (s *UserRpcServerV1) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	s.log.Printf("Receive Update")

	if req.Name != nil {
		nameStr := fmt.Sprintf("Name: %v\n", req.GetName().GetValue())
		s.log.Printf(nameStr)
	}

	if req.Email != nil {
		nameStr := fmt.Sprintf("Email: %v\n", req.GetEmail().GetValue())
		s.log.Printf(nameStr)
	}

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline %v\n", dline)
	}

	return &emptypb.Empty{}, nil
}
func (s *UserRpcServerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline %v\n", dline)
	}

	s.log.Printf("Received Delete in ID: %v\n", req.GetId())

	return &emptypb.Empty{}, nil
}
