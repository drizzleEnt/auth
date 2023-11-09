package auth

import (
	"log"
	"os"

	"github.com/drizzleent/auth/internal/service"
	desc "github.com/drizzleent/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	log         *log.Logger
	authservice service.AuthService
}

func NewImplementation(service service.AuthService) *Implementation {
	log := log.New(os.Stdout, "INFO:", log.Flags())
	return &Implementation{
		log:         log,
		authservice: service,
	}
}
