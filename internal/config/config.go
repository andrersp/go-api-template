package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ENVIRONMENT  string
	API_PORT     string
	SECRET_TOKEN string
)

func SetConfig() {

	ENVIRONMENT = os.Getenv("ENVIRONMENT")

	err := godotenv.Load()
	if err != nil && ENVIRONMENT == "LOCAL" {
		log.Fatal(err)
	}

	API_PORT = os.Getenv("API_PORT")
	SECRET_TOKEN = os.Getenv("SECRET_TOKEN")
}
