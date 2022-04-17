package config

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnvVariable(key string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
