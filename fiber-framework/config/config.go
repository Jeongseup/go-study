package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error while loading .env file")
	}
	return os.Getenv(key)
}

// TODO : Auth with JWT
