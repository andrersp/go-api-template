package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	API_PORT     string
	SECRET_TOKEN string
)

func SetConfig() {
	API_PORT = "8080"

	godotenv.Load()
	SECRET_TOKEN = os.Getenv("SECRET_TOKEN")
}
