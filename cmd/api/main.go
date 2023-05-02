package main

import (
	"log"

	"github.com/andrersp/go-api-template/internal/adapters"
	api "github.com/andrersp/go-api-template/internal/adapters/api"
	repository "github.com/andrersp/go-api-template/internal/adapters/repository/postgres"
)

func init() {
	adapters.SetConfig()
	err := repository.CreateSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = repository.AutoMigrate()
	if err != nil {
		log.Fatal()
	}

}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @Accept json
// Produce json

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {

	api.StartApiServer()

}
