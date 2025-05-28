package config

import (
	"errors"
	"os"
)

const env = "ENV"

type LoggerConfig interface {
	Environment() string
}

type loggerConfig struct {
	env string
}

func NewLoggerConfig() (LoggerConfig, error) {
	e := os.Getenv(env)
	if len(e) == 0 {
		return nil, errors.New("environment not found")
	}

	return &loggerConfig{
		env: e,
	}, nil
}

func (l loggerConfig) Environment() string {
	return l.env
}
