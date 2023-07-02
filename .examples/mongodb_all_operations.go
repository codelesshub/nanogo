package main

import (
	"fmt"

	"github.com/codelesshub/nanogo/config/env"
	"github.com/codelesshub/nanogo/config/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID    *uuid.UUID `bson:"_id,omitempty"`
	Nome  string     `bson:"nome"`
	Idade int        `bson:"idade"`
	Email string     `bson:"email"`
}

func (u *User) GetID() *uuid.UUID {
	return u.ID
}

func (u *User) SetID(id *uuid.UUID) {
	u.ID = id
}

func main() {
	// Carrega o arquivo .env
	env.LoadEnv()

	// Cria um novo usuário
	user := &User{
		Nome:  "João",
		Idade: 30,
		Email: "joao@gmail.com",
	}

	// Cria o repositório passando o user como model
	repository := mongodb.NewMongoRepository("users", user)

	// INSERT
	createdUser, err := repository.Insert(user)
	if err != nil {
		// Handle error
	}
	fmt.Println("Inserted document with ID:", createdUser.GetID())

	// Update email
	user.Email = "joao_novo@gmail.com"
	updatedUser, err := repository.Update(user)
	if err != nil {
		// Handle error
	}
	fmt.Println("Updated document:", updatedUser)

	// FIND BY ID
	resultUser, err := repository.FindById(createdUser.GetID())
	if err != nil {
		// Handle error
	}
	fmt.Println("Found document:", resultUser)

	// FIND ALL
	users, err := repository.FindAll()
	if err != nil {
		// Handle error
	}
	for _, user := range users {
		fmt.Println("Found all document:", user)
	}

	// RAW QUERY
	query := bson.M{"idade": bson.M{"$gt": 20}} // Find users older than 20
	usersRawQuery, err := repository.RawQuery(query)
	if err != nil {
		// Handle error
	}
	for _, user := range usersRawQuery {
		fmt.Println("Found raw document", user)
	}

	// DELETE
	deleted, err := repository.Delete(user)
	if err != nil {
		// Handle error
	}
	fmt.Println("Deleted documents:", deleted)

}
