package login

import (
	"context"

	"github.com/drizzleent/auth/internal/converter"
	desc "github.com/drizzleent/auth/pkg/login_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	obj, err := i.loginService.Login(ctx, converter.ToUserClaimsFromLogin(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: obj,
	}, nil
}
