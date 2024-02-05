package configurations

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func CreateClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(LoadEnv()))

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
		return err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging to MongoDB: ", err)
		return err
	}

	log.Println("Connected to MongoDB!")
	Client = client
	return nil
}
