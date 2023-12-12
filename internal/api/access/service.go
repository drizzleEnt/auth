package access

import (
	"github.com/drizzleent/auth/internal/service"
	desc "github.com/drizzleent/auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(service service.AccessService) *Implementation {

	return &Implementation{
		accessService: service,
	}
}
