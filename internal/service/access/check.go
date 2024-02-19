package access

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/drizzleent/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const (
	prefix = "Bearer "
)

var accessibleRoles map[string]string

func (s *serviceAccess) Check(ctx context.Context, endpointAddress string) error {
	accessTokenSecretKey := os.Getenv("accessTokenSecretKey")

	accesToken, err := accessToken(ctx)
	if err != nil {
		return errors.New("access token is invalid: " + err.Error())
	}

	claims, err := utils.VerifyToken(accesToken, []byte(accessTokenSecretKey))

	if err != nil {
		return errors.New("access token is invalid " + err.Error())
	}

	accessibleMap, err := s.accessibleRoles(ctx)

	if err != nil {
		return errors.New("failed to get accessible roles " + err.Error())
	}

	role, ok := accessibleMap[endpointAddress]

	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access dinied")
}

func (s *serviceAccess) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		Roles, err := s.accessRepository.Check(ctx)
		if err != nil {
			return nil, nil
		}

		accessibleRoles = Roles
	}

	return accessibleRoles, nil
}

func accessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadat is not provided")
	}
	authHeader, ok := md["authorization"]

	if !ok || len(authHeader) == 0 {
		return "", errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], prefix) {
		return "", errors.New("invalid authrization header format")
	}

	accesToken := strings.TrimPrefix(authHeader[0], prefix)
	return accesToken, nil
}
