package main

import (
	"fmt"

	"github.com/codelesshub/nanogo/config/env"
	"github.com/codelesshub/nanogo/config/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// INSERT
	repository := mongodb.NewMongoRepository("users")
	document := map[string]interface{}{
		"_id":   uuid.New(), // Use UUID for document ID
		"nome":  "João",
		"idade": 30,
		"email": "joao@gmail.com",
	}
	id, err := repository.Insert(document)
	if err != nil {
		// Handle error
	}
	fmt.Println("Inserted document with ID:", id)

	update := map[string]interface{}{
		"email": "joao_novo@gmail.com",
	}
	updateResult, err := repository.UpdateById(id, update)
	if err != nil {
		// Handle error
	}
	fmt.Println("Updated documents:", updateResult.ModifiedCount)

	// FIND BY ID
	foundDocument, err := repository.FindById(id)
	if err != nil {
		// Handle error
	}
	fmt.Println("Found document:", foundDocument)

	if email, ok := foundDocument["email"]; ok {
		fmt.Println("Email:", email)
	} else {
		fmt.Println("Email field does not exist")
	}

	result, err := repository.DeleteById(id)
	if err != nil {
		// Handle error
	}
	fmt.Println("Deleted documents:", result.DeletedCount)

	documents, err := repository.FindAll()
	if err != nil {
		// Handle error
	}
	for _, document := range documents {
		fmt.Println("Found document:", document)
	}

	query := bson.M{"idade": bson.M{"$gt": 20}} // Find users older than 20
	documentsRawQuery, err := repository.RawQuery(query)
	if err != nil {
		// Handle error
	}
	for _, document := range documentsRawQuery {
		fmt.Println("Found document:", document)
	}

}
