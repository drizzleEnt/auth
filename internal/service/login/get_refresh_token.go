package login

import (
	"context"
	"os"

	"github.com/drizzleent/auth/internal/model"
	"github.com/drizzleent/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serviceLogin) GetRefreshToken(ctx context.Context, token string) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))

	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", nil
	}

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		UserName: claims.Username,
		Role:     r,
	}, []byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
