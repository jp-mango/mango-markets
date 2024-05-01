package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (API_KEY string, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY = os.Getenv("API_KEY")

	return API_KEY, err
}
