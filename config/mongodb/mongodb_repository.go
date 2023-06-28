package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *MongoRepository) Insert(document map[string]interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (r *MongoRepository) Update(id primitive.ObjectID, update map[string]interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id}}
	updateDoc := bson.M{"$set": update} // Convert map to bson.M
	result, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return nil, err
	}

	return result, nil
}


func (r *MongoRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository) FindById(id primitive.ObjectID) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"_id", id}}
	var result bson.M
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MongoRepository) FindAll() ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func (r *MongoRepository) RawQuery(query bson.M) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
