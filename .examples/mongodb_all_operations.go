package main

import (
	"fmt"

	"github.com/codelesshub/nanogo/config/env"
	"github.com/codelesshub/nanogo/config/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// INSERT
	repository := mongodb.NewMongoRepository("users")
	document := map[string]interface{}{
		"nome":  "João",
		"idade": 30,
		"email": "joao@gmail.com",
	}
	insertResult, err := repository.Insert(document)
	if err != nil {
		// Handle error
	}
	fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	// Convert InsertedID to ObjectID
	id, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		// Handle error if conversion is not successful
	}

	update := map[string]interface{}{
		"email": "joao_novo@gmail.com",
	}
	updateResult, err := repository.Update(id, update)
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

	result, err := repository.Delete(id)
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
