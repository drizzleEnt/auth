package login

import (
	"context"

	desc "github.com/drizzleent/auth/pkg/login_v1"
)

func (i *Implementation) GetAccesToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	obj, err := i.loginService.GetAccessToken(ctx, req.GetRefreshToken())

	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: obj,
	}, err
}
