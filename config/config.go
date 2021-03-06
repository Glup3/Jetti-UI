package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get return env secret for given key
func Get(key string) string {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return os.Getenv(key)
}
