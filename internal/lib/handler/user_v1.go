package handler

import (
	"context"
	crypto "crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/brianvoe/gofakeit"
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
}

func NewUserRpcsServer(log *log.Logger) *UserRpcServerV1 {
	return &UserRpcServerV1{
		log: log,
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

	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100123))
	if err != nil {
		s.log.Printf("cant generate rand id %v\n", err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	id := safeNum.Int64()

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
func (s *UserRpcServerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	s.log.Printf("Receive Get ID: %v", req.Id)

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline %v\n", dline)
	}

	role := gofakeit.RandString([]string{"ADMIN", "USER"})

	resp := desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	respStr := fmt.Sprintf("Responce Get \n\tId: %v,\n\tName: %v,\n\tEmail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt)

	s.log.Println(respStr)

	return &resp, nil
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
