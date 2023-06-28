package main

import (
	"fmt"

	mongodb "github.com/codelesshub/nanogo/config/database"
	"github.com/codelesshub/nanogo/config/env"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Connecta no mongodb
	mongodb.ConnectMongoDB()

	repository := mongodb.NewMongoRepository("users")
	documents, err := repository.FindAll()
	if err != nil {
		// Handle error
	}
	for _, document := range documents {
		fmt.Println("Found document:", document)
	}

}
