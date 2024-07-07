package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
