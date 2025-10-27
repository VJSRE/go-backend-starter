package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetValue(key string) string {
	err := godotenv.Load("./credentials.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
