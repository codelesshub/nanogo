package main

import (
	"fmt"

	mongodb "github.com/codelesshub/nanogo/config/database"
	"github.com/codelesshub/nanogo/config/env"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Connecta no mongodb
	mongodb.ConnectMongoDB()

	repository := mongodb.NewMongoRepository("users")

	query := bson.M{"idade": bson.M{"$gt": 20}} // Find users older than 20
	documents, err := repository.RawQuery(query)
	if err != nil {
		// Handle error
	}
	for _, document := range documents {
		fmt.Println("Found document:", document)
	}

}
