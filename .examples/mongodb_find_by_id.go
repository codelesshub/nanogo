package main

import (
	"fmt"

	mongodb "github.com/codelesshub/nanogo/config/database"
	"github.com/codelesshub/nanogo/config/env"
	"github.com/google/uuid"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Connecta no mongodb
	mongodb.ConnectMongoDB()

	repository := mongodb.NewMongoRepository("users")
	id, _ := uuid.Parse("3b241101-aa3d-4adb-8c92-fb4d4f6b1f16") // Use the correct UUID
	document, err := repository.FindById(id)
	if err != nil {
		// Handle error
	}
	fmt.Println("Found document:", document)

}
