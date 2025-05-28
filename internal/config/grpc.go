package config

import (
	"errors"
	"net"
	"os"
	"time"
)

const (
	grpcHostEnvName    = "GRPC_HOST"
	grpcPortEnvName    = "GRPC_PORT"
	grpcTimeoutEnvName = "GRPC_TIMEOUT"
)

type GRPCConfig interface {
	Address() string
	Timeout() time.Duration
}

type grpcConfig struct {
	host    string
	port    string
	timeout time.Duration
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}
	timeout := os.Getenv(grpcTimeoutEnvName)
	if len(timeout) == 0 {
		return nil, errors.New("grpc port not found")
	}

	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, errors.New("grpc timeout not found")
	}

	return &grpcConfig{
		host:    host,
		port:    port,
		timeout: t,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
func (cfg *grpcConfig) Timeout() time.Duration {
	return cfg.timeout
}
