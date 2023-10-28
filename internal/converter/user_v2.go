package converter

import (
	model1 "github.com/drizzleent/auth/internal/model"
	desc "github.com/drizzleent/auth/pkg/user_v2"
)

func ToModelFromCreateRequest(req *desc.CreateRequest) *model1.User {
	return &model1.User{
		Name:     req.User.Name.Value,
		Email:    req.User.Email.Value,
		Password: req.Password.Value,
		Role:     int(req.User.Role),
	}
}

func ToModelUserFromUpdateRequest(req *desc.UpdateRequest) *model1.User {
	var (
		name  string
		email string
		role  int
	)

	if req.User.Name != nil {
		name = req.User.Name.Value
	}

	if req.User.Email != nil {
		email = req.User.Email.Value
	}

	if req.User.Role != desc.Role_UNKNOWN {
		role = int(req.User.Role.Number())
	}

	return &model1.User{
		ID:    req.Id.Value,
		Name:  name,
		Email: email,
		Role:  role,
	}
}
