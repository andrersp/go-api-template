package config

var (
	API_PORT     string
	SECRET_TOKEN string
)

func SetConfig() {
	API_PORT = "8080"
}
