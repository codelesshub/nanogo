package mongodb

import (
	"context"
	"time"

	"github.com/codelesshub/nanogo/config/log"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// other imports
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collectionName string) *MongoRepository {
	db, err := ConnectMongoDB()

	if err != nil {
		log.Fatal(err)
	}

	collection := db.Collection(collectionName)

	return &MongoRepository{collection: collection}
}

func (r *MongoRepository) Insert(document map[string]interface{}) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Generate a new UUID
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	// Add the UUID to the document as a byte slice
	document["_id"] = uuid[:]

	_, err = r.collection.InsertOne(ctx, document)

	if err != nil {
		return nil, err
	}

	return &uuid, nil
}

func (r *MongoRepository) UpdateById(id *uuid.UUID, update map[string]interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id[:]}}
	updateDoc := bson.M{"$set": update}
	result, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository) DeleteById(id *uuid.UUID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id[:]}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository) FindById(id *uuid.UUID) (map[string]interface{}, error) {
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
			// handle error
		}
		result["_id"] = convertedUUID
	}

	return result, nil
}

func (r *MongoRepository) FindAll() ([]map[string]interface{}, error) {
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

	for i := range results {
		if idField, ok := results[i]["_id"].(primitive.Binary); ok {
			convertedUUID, err := uuid.FromBytes(idField.Data)
			if err != nil {
				// handle error
			}
			results[i]["_id"] = convertedUUID
		}
	}

	return results, nil
}

func (r *MongoRepository) RawQuery(query bson.M) ([]map[string]interface{}, error) {
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

	for i := range results {
		if idField, ok := results[i]["_id"].(primitive.Binary); ok {
			convertedUUID, err := uuid.FromBytes(idField.Data)
			if err != nil {
				// handle error
			}
			results[i]["_id"] = convertedUUID
		}
	}

	return results, nil
}
