package mongodb

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/codelesshub/nanogo/config/log"
)

var (
	clientInstance *mongo.Client
	once           sync.Once
	err            error
)

func ConnectMongoDB() (*mongo.Database, error) {
	once.Do(func() {
		mongoURI := os.Getenv("MONGO_URI")
		var clientOptions *options.ClientOptions

		if mongoURI != "" {
			clientOptions = options.Client().ApplyURI(mongoURI)
		} else {
			dbname := os.Getenv("MONGO_DBNAME")
			username := os.Getenv("MONGO_USERNAME")
			password := os.Getenv("MONGO_PASSWORD")
			host := os.Getenv("MONGO_HOST")
			port, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
			if err != nil {
				port = 27017 // Default MongoDB Port
			}

			if dbname == "" || username == "" || password == "" || host == "" {
				err = fmt.Errorf("MongoDB environment variables are not set properly!")
				return
			}

			connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
			clientOptions = options.Client().ApplyURI(connectionURI).SetAuth(options.Credential{AuthSource: dbname})
		}

		clientInstance, err = mongo.Connect(context.Background(), clientOptions)
	})

	if err != nil {
		return nil, err
	}

	log.Info("Connected to MongoDB!")

	dbname := os.Getenv("MONGO_DBNAME")

	return clientInstance.Database(dbname), nil
}
