package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ENVIROMMENT  string
	API_PORT     string
	SECRET_TOKEN string
)

func SetConfig() {
	API_PORT = os.Getenv("API_PORT")
	ENVIROMMENT = os.Getenv("ENVIROMMENT")

	err := godotenv.Load()

	if err != nil && ENVIROMMENT == "LOCAL" {
		log.Fatal(err)
	}
	SECRET_TOKEN = os.Getenv("SECRET_TOKEN")
}
