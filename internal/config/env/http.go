package env

import (
	"errors"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST_AUTH"
	httpPortEnvName = "HTTP_PORT_AUTH"
)

type httpConfig struct {
	host string
	port string
}

func NewHttpConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host nit found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
