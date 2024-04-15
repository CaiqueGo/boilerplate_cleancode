package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	dbName string
	dbHost string
	dbUser string
	dbPass string
	dbPort string
}

func Load() (Environment, error) {
	if os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			return Environment{}, err
		}
	}

	return Environment{
		dbName: os.Getenv("DB_NAME"),
		dbHost: os.Getenv("DB_HOST"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASS"),
		dbPort: os.Getenv("DB_PORT"),
	}, nil
}
