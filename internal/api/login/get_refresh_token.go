package login

import (
	"context"

	desc "github.com/drizzleent/auth/pkg/login_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	obj, err := i.loginService.GetRefreshToken(ctx, req.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: obj,
	}, nil
}
