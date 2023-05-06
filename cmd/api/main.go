package main

import (
	"log"

	"github.com/andrersp/go-api-template/internal/config"
	api "github.com/andrersp/go-api-template/internal/handlers/api"
	repository "github.com/andrersp/go-api-template/internal/repository/postgres"
)

func init() {
	config.SetConfig()
	err := repository.CreateSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = repository.AutoMigrate()
	if err != nil {
		log.Fatal()
	}

}

// func GenerateSecureToken(length int) string {
// 	b := make([]byte, length)
// 	if _, err := rand.Read(b); err != nil {
// 		return ""
// 	}
// 	return hex.EncodeToString(b)
// }

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @Accept json
// Produce json

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	// token := GenerateSecureToken(32)
	// fmt.Println(token)

	api.StartApiServer()

}
