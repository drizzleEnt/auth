package access

import (
	"context"
	"errors"
	"strings"

	"github.com/drizzleent/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const (
	prefix = "Bearer"

	accessTokenSecretKey = "AwdaWAdsIU8769iJBVFmxkslkcejcUajvueoJLnHf90a"
)

var accessibleRoles map[string]string

func (s *serviceAccess) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadat is not provided")
	}

	authHeader, ok := md["authtorization"]

	if !ok || len(authHeader) == 0 {
		return errors.New("authrization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], prefix) {
		return errors.New("invalid authrization header format")
	}

	accesToken := strings.TrimPrefix(authHeader[0], prefix)

	claims, err := utils.VerifyToken(accesToken, []byte(accessTokenSecretKey))

	if err != nil {
		return errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)

	if err != nil {
		return errors.New("failed to get accessible roles")
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
