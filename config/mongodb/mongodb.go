package mongodb

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/codelesshub/nanogo/config/env"
	"github.com/codelesshub/nanogo/config/log"
)

var (
	clientInstance *mongo.Client
	once           sync.Once
	err            error
)

func ConnectMongoDB() (*mongo.Database, error) {
	once.Do(func() {
		mongoURI := env.GetEnv("MONGO_URI", "")
		var clientOptions *options.ClientOptions

		if mongoURI != "" {
			clientOptions = options.Client().ApplyURI(mongoURI)
		} else {
			dbnameAuth := env.GetEnv("MONGO_AUTH_DBNAME", "admin")
			username := env.GetEnv("MONGO_USERNAME")
			password := env.GetEnv("MONGO_PASSWORD")
			host := env.GetEnv("MONGO_HOST")
			port := env.GetEnv("MONGO_PORT", "27017")

			connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
			clientOptions = options.Client().ApplyURI(connectionURI).SetAuth(options.Credential{AuthSource: dbnameAuth})
		}

		clientInstance, err = mongo.Connect(context.Background(), clientOptions)
	})

	if err != nil {
		return nil, err
	}

	log.Info("Connected to MongoDB!")

	dbname := env.GetEnv("MONGO_DBNAME")

	return clientInstance.Database(dbname), nil
}
