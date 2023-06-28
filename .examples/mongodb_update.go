package main

import (
	"fmt"

	mongodb "github.com/codelesshub/nanogo/config/database"
	"github.com/codelesshub/nanogo/config/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Connecta no mongodb
	mongodb.ConnectMongoDB()

	repository := NewMongoRepository("users")
	id, _ := primitive.ObjectIDFromHex("60d5ec9af682fbd39a8920") // Use the correct ID
	update := map[string]interface{}{
		"email": "joao_novo@gmail.com",
	}
	result, err := repository.Update(id, update)
	if err != nil {
		// Handle error
	}
	fmt.Println("Updated documents:", result.ModifiedCount)

}
