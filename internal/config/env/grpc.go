package env

import (
	"net"
	"os"

	"github.com/drizzleent/auth/internal/config"
	"github.com/pkg/errors"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST_AUTH"
	grpcPortEnvName = "GRPC_PORT_AUTH"
)

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)

	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)

	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
