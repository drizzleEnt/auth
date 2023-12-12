package access

import (
	"github.com/drizzleent/auth/internal/repository"
	"github.com/drizzleent/auth/internal/service"
)

type serviceAccess struct {
	accessRepository repository.AccessRepository
}

func NewAccessService(accessRepository repository.AccessRepository) service.AccessService {
	return &serviceAccess{
		accessRepository: accessRepository,
	}
}
