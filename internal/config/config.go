package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func Environment() (string, error) {
	e := os.Getenv(env)
	if len(e) == 0 {
		return "", errors.New("environment not found")
	}
	return e, nil
}
