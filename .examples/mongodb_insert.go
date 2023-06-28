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
	document := map[string]interface{}{
		"nome":  "João",
		"idade": 30,
		"email": "joao@gmail.com",
	}
	result, err := repository.Insert(document)
	if err != nil {
		// Handle error
	}
	fmt.Println("Inserted document with ID:", result.InsertedID)

}
