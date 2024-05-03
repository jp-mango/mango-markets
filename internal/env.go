package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (API_KEY, DB_CONN string, err error) {
	err = godotenv.Load()
	//* (UNCOMMENT FOR DEBUGGING) //err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY = os.Getenv("API_KEY")
	DB_CONN = os.Getenv("DB_DSN")

	return API_KEY, DB_CONN, err
}
