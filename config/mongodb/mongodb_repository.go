package mongodb

import (
	"context"
	"reflect"
	"time"

	"github.com/codelesshub/nanogo/config/log"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// other imports
)

type MongoRepository struct {
	collection *mongo.Collection
	model      Model
}

func NewMongoRepository(collectionName string, model Model) *MongoRepository {
	db, err := ConnectMongoDB()

	if err != nil {
		log.Fatal(err)
	}

	collection := db.Collection(collectionName)

	return &MongoRepository{collection: collection, model: model}
}

func (r *MongoRepository) Insert(document Model) (Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uuid, _ := uuid.NewRandom()
	document.SetID(&uuid)
	_, err := r.collection.InsertOne(ctx, document)

	if err != nil {
		return nil, err
	}

	return document, err
}

func (r *MongoRepository) Update(document Model) (Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", document.GetID()}}
	updateDoc := bson.M{"$set": document}
	_, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (r *MongoRepository) Delete(document Model) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", document.GetID()}}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MongoRepository) Save(document Model) (Model, error) {
	if document.GetID() == nil {
		return r.Insert(document)
	} else {
		return r.Update(document)
	}
}

func (r *MongoRepository) FindById(id *uuid.UUID) (Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id}}
	var result map[string]interface{}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Convert _id field back to uuid.UUID
	if idField, ok := result["_id"].(primitive.Binary); ok {
		convertedUUID, err := uuid.FromBytes(idField.Data)
		if err != nil {
			return nil, err
		}
		result["id"] = convertedUUID
	}

	// We need a fresh instance for each document.
	outputModel := reflect.New(reflect.TypeOf(r.model).Elem()).Interface().(Model)
	err = mapstructure.Decode(result, &outputModel)
	if err != nil {
		return nil, err
	}

	return outputModel, nil
}

func (r *MongoRepository) FindAll() ([]Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	var outputModels []Model

	for _, result := range results {

		if idField, ok := result["_id"].(primitive.Binary); ok {
			convertedUUID, err := uuid.FromBytes(idField.Data)
			if err != nil {
				return nil, err
			}
			result["id"] = convertedUUID
		}

		model := reflect.New(reflect.TypeOf(r.model).Elem()).Interface().(Model)
		err = mapstructure.Decode(result, &model)
		if err != nil {
			return nil, err
		}

		outputModels = append(outputModels, model)
	}

	return outputModels, nil
}

func (r *MongoRepository) RawQuery(query bson.M) ([]Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	var outputModels []Model

	for _, result := range results {
		if idField, ok := result["_id"].(primitive.Binary); ok {
			convertedUUID, err := uuid.FromBytes(idField.Data)
			if err != nil {
				return nil, err
			}
			result["id"] = convertedUUID
		}

		model := reflect.New(reflect.TypeOf(r.model).Elem()).Interface().(Model)
		err = mapstructure.Decode(result, &model)
		if err != nil {
			return nil, err
		}

		outputModels = append(outputModels, model)
	}

	return outputModels, nil
}

func decode(document Model) (map[string]interface{}, error) {
	var mapInterface map[string]interface{}
	err := mapstructure.Decode(document, &mapInterface)
	if err != nil {
		return nil, err
	}
	return mapInterface, nil
}

func encode(inputMap map[string]interface{}, outputModel Model) error {
	return mapstructure.Decode(inputMap, outputModel)
}
