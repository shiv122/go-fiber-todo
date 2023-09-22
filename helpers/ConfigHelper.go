package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvHelper struct{}

func (eh *EnvHelper) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func (eh *EnvHelper) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
