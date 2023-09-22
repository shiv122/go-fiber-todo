package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfiguration struct {
	Database string
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

var DB DatabaseConfiguration

func init() {
	dbc := new(DatabaseConfiguration)
	dbc.loadEnv()
	DB = DatabaseConfiguration{
		Database: dbc.getEnv("DATABASE", "mysql"),
		Username: dbc.getEnv("DB_USERNAME", "root"),
		Password: dbc.getEnv("DB_PASSWORD", ""),
		Host:     dbc.getEnv("DB_HOST", "127.0.0.1"),
		Port:     dbc.getEnv("DB_PORT", "3306"),
		DBName:   dbc.getEnv("DB_NAME", ""),
	}
}

func (dc *DatabaseConfiguration) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func (dc *DatabaseConfiguration) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
