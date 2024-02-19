package login

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/utils"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 10 * time.Minute
)

func (s *serviceLogin) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")

	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		UserName: info.Username,
		Password: info.Password,
		Role:     r,
	}, []byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
