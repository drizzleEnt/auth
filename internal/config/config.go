package config

import "github.com/subosito/gotenv"

func Load(path string) error {
	err := gotenv.Load(path)

	if err != nil {
		return err
	}

	return nil
}

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Address() string
}

type SwaggerConfig interface {
	Address() string
}

type PrometheusConfig interface {
	Address() string
}
