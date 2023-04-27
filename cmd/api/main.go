package main

import (
	"log"

	api "github.com/andrersp/go-api-template/internal/api"
	"github.com/andrersp/go-api-template/internal/config"
)

func init() {
	config.SetConfig()
	err := config.CreateSQLiteConn()
	if err != nil {
		log.Fatal(err)
	}

	err = config.AutoMigrate()
	if err != nil {
		log.Fatal()
	}
}
func main() {

	api.StartApiServer()

}
