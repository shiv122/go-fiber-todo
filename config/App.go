package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	SecretKey string
	Port      string
}

var App AppConfig

func init() {
	a := new(AppConfig)
	a.loadEnv()
	App = AppConfig{
		SecretKey: a.getEnv("SECRET_KEY", "secret"),
		Port:      a.getEnv("PORT", "4000"),
	}
}

func (ac *AppConfig) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func (ac *AppConfig) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
