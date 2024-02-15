package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	//? loads env variables from file in root
	godotenv.Load("../../.env")

	//? check for API key
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		log.Fatal("API_KEY is not found in the env")
	}
	return API_KEY
}
